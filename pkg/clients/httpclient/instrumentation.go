// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package httpclient

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"time"

	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type instrumentedRoundTripper struct {
	Next                    http.RoundTripper
	Region                  string
	enableConnectivityTrace bool
	loggerConstructor       func(*log.APITracking, *http.Request) *log.Logger
}

type requestConnectionInfo struct {
	connReused string
	remoteAddr string
	localAddr  string

	// use string to contain error case.
	connLatency  string
	dnsLatency   string
	tlsLatency   string
	totalLatency string
}

// Instrument logs and traces requests handled by the returned round tripper.
//
// The logic was derived from @xiazhan's work in logging_send_decorator.go
// originally committed in 74df1ef as an autorest send decorator. This adoption
// allows it to be used by non-autorest HTTP requests.
func Instrument(next http.RoundTripper, region string) http.RoundTripper {
	return &instrumentedRoundTripper{
		Next:                    next,
		Region:                  region,
		enableConnectivityTrace: false,
		loggerConstructor: func(a *log.APITracking, req *http.Request) *log.Logger {
			return log.NewOutgoingRequestLogger(a, req)
		},
	}
}

func InstrumentWithConnection(next http.RoundTripper, region string) http.RoundTripper {
	return &instrumentedRoundTripper{
		Next:                    next,
		Region:                  region,
		enableConnectivityTrace: true,
		loggerConstructor: func(a *log.APITracking, req *http.Request) *log.Logger {
			return log.NewOutgoingRequestLogger(a, req)
		},
	}
}

func (i *instrumentedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	// Get fields to log for this particular request
	vars := mux.Vars(req)
	apiTracking, err := log.NewAPITrackingFromOutgoingRequest(vars, req, req.RequestURI, i.Region)
	if err != nil {
		return nil, fmt.Errorf("failed to build apitracking: %v", err)
	}

	var getConn, dnsStart, connStart, tlsStart *time.Time
	var totalLatency, dnsLatency, connLatency, tlsLatency string
	var reqConnInfo *httptrace.GotConnInfo
	if i.enableConnectivityTrace {
		trace := &httptrace.ClientTrace{
			GetConn: func(hostPort string) {
				getConn = timeNowPtr()
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				if getConn != nil {
					totalLatency = fmt.Sprintf("%dms", time.Now().Sub(*getConn).Milliseconds())
				}

				reqConnInfo = &connInfo
			},
			DNSStart: func(_ httptrace.DNSStartInfo) {
				dnsStart = timeNowPtr()
			},
			DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
				if dnsInfo.Err == nil {
					if dnsStart != nil {
						dnsLatency = fmt.Sprintf("%dms", time.Now().Sub(*dnsStart).Milliseconds())
					}
				} else {
					dnsLatency = dnsInfo.Err.Error()
				}
			},
			ConnectStart: func(_, _ string) {
				connStart = timeNowPtr()
			},
			ConnectDone: func(_, _ string, err error) {
				if err == nil {
					if connStart != nil {
						connLatency = fmt.Sprintf("%dms", time.Now().Sub(*connStart).Milliseconds())
					}
				} else {
					connLatency = err.Error()
				}
			},
			TLSHandshakeStart: func() {
				tlsStart = timeNowPtr()
			},
			TLSHandshakeDone: func(_ tls.ConnectionState, err error) {
				if err == nil {
					if tlsStart != nil {
						tlsLatency = fmt.Sprintf("%dms", time.Now().Sub(*tlsStart).Milliseconds())
					}
				} else {
					tlsLatency = err.Error()
				}
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	}

	// Construct the outbound request logger
	logger := i.loggerConstructor(apiTracking, req)

	var res *http.Response
	defer func() {
		var extendedConnInfo *requestConnectionInfo = &requestConnectionInfo{}
		if i.enableConnectivityTrace {
			extendedConnInfo.connLatency = connLatency
			extendedConnInfo.dnsLatency = dnsLatency
			extendedConnInfo.tlsLatency = tlsLatency
			extendedConnInfo.totalLatency = totalLatency
			if reqConnInfo != nil {
				extendedConnInfo.remoteAddr = reqConnInfo.Conn.RemoteAddr().String()
				extendedConnInfo.localAddr = reqConnInfo.Conn.LocalAddr().String()
				extendedConnInfo.connReused = strconv.FormatBool(reqConnInfo.Reused)
			}
		}

		logger.TraceLogger = logger.TraceLogger.WithFields(getFieldsForResponse(res, startTime, extendedConnInfo))
		logger.Info(req.Context(), "HttpRequestEnd")
	}()

	logger.Info(req.Context(), "HttpRequestStart")
	res, err = i.Next.RoundTrip(req)
	return res, err
}

func getFieldsForResponse(res *http.Response, startTime time.Time, extendedConnInfo *requestConnectionInfo) logrus.Fields {
	const nanoToMilliSecondFactor = 1e6
	duration := int(time.Now().Sub(startTime).Nanoseconds() / nanoToMilliSecondFactor)
	if res == nil {
		return logrus.Fields{
			"statusCode":             "0",
			"contentLength":          "-1",
			"durationInMilliseconds": strconv.Itoa(duration),
		}
	}

	serviceRequestID := res.Header.Get(log.ResponseARMRequestIDHeader)
	if serviceRequestID == "" {
		serviceRequestID = res.Header.Get(log.ResponseGraphRequestIDHeader)
	}

	return logrus.Fields{
		"statusCode":             strconv.Itoa(res.StatusCode),
		"serviceRequestID":       serviceRequestID,
		"correlationID":          res.Header.Get(log.RequestCorrelationIDHeader),
		"contentLength":          strconv.FormatInt(res.ContentLength, 10),
		"durationInMilliseconds": strconv.Itoa(duration),
		"connReused":             extendedConnInfo.connReused,
		"remoteAddr":             extendedConnInfo.remoteAddr,
		"localAddr":              extendedConnInfo.localAddr,
		"totalConnLatency":       extendedConnInfo.totalLatency,
		"dnsLatency":             extendedConnInfo.dnsLatency,
		"connLatency":            extendedConnInfo.connLatency,
		"tlsLatency":             extendedConnInfo.tlsLatency,
	}
}

func timeNowPtr() *time.Time {
	t := time.Now()
	return &t
}
