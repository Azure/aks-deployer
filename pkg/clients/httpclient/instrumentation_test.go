package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	"github.com/sirupsen/logrus"

	"github.com/Azure/aks-deployer/pkg/log"
)

var _ = Describe("instrumented HTTP round tripper", func() {
	It("should log the expected output given a successful request - ARM", func() {
		// Store log entries to assert on later
		hooks := &loggerTestHooks{}
		logger := logrus.New()
		logger.Hooks.Add(hooks)
		logger.SetFormatter(&logrus.TextFormatter{DisableColors: true})

		serviceRequestID := uuid.NewV4().String()
		correlationID := uuid.NewV4()
		content := "hello world"

		// Run a fake HTTP server to handle test request(s)
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := json.Marshal(content)

			header := w.Header()
			header.Set("Content-Type", "application/json")
			header.Set("x-ms-request-id", serviceRequestID)
			header.Set("x-ms-correlation-request-id", correlationID.String())

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, string(body))
		}))
		defer svr.Close()

		// Prove a tracing span is created by getting it off of
		// the context from the next round tripper
		hookInvoked := false
		next := &roundTripperProxy{
			Hook: func(req *http.Request) {
				hookInvoked = true
			},
			Next: http.DefaultTransport,
		}

		// Build an instrumented round tripper with the test logger
		region := "test-region"
		roundTripper := Instrument(next, region)
		roundTripperImpl := roundTripper.(*instrumentedRoundTripper)
		roundTripperImpl.loggerConstructor = func(a *log.APITracking, request *http.Request) *log.Logger {
			return log.NewOutgoingRequestLoggerForLogrus(logger.WithField("source", "test"), a, request)
		}

		operationID := uuid.NewV4()
		subscriptionID := uuid.NewV4()
		resourceGroupName := "testRG"
		resourceName := "testResource"
		operationName := "testOp"
		subOperationName := "testSubOp"

		ctx := log.WithLogger(context.Background(), log.InitializeTestLogger())
		fields := map[string]interface{}{
			"operationID":       operationID,
			"subscriptionID":    subscriptionID,
			"resourceGroupName": resourceGroupName,
			"resourceName":      resourceName,
			"operationName":     operationName,
			"subOperationName":  subOperationName,
			"correlationID":     uuid.NewV4().String(),
		}
		apiTracking := log.NewAPITrackingFromParametersMap(fields)
		ctx = log.WithAPITracking(ctx, apiTracking)

		// Issue the request
		req, err := http.NewRequest("GET", svr.URL, nil)
		req = req.WithContext(ctx)
		clientRequestID := uuid.NewV4().String()
		req.Header.Add("x-ms-client-request-id", clientRequestID)
		req.Header.Add("x-ms-client-session-id", operationID.String())

		Expect(err).ToNot(HaveOccurred())
		res, err := roundTripper.RoundTrip(req)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.StatusCode).To(Equal(http.StatusOK))
		Expect(hookInvoked).To(BeTrue())

		// Two log entries should be published: one for start and one for end
		Expect(hooks.Entries).To(HaveLen(2))

		// Prove the expected fields were provided on both log entries
		entry, err := hooks.Entries[0].String()
		Expect(err).ToNot(HaveOccurred())
		Expect(entry).To(ContainSubstring("httpMethod=GET"))
		Expect(entry).To(ContainSubstring(`targetURI="http://127.0.0.1`))
		Expect(entry).To(ContainSubstring(`hostName="127.0.0.1:`))
		Expect(entry).To(ContainSubstring("region=" + region))
		Expect(entry).To(ContainSubstring("source=OutgoingRequestTraceLog"))
		Expect(entry).To(ContainSubstring("clientRequestID=" + clientRequestID))
		Expect(entry).To(ContainSubstring("operationID=" + operationID.String()))
		Expect(entry).To(ContainSubstring("subscriptionID=" + subscriptionID.String()))
		Expect(entry).To(ContainSubstring("resourceGroupName=" + resourceGroupName))
		Expect(entry).To(ContainSubstring("resourceName=" + resourceName))
		Expect(entry).To(ContainSubstring("operationName=" + operationName))
		Expect(entry).To(ContainSubstring("suboperationName=" + subOperationName))
		Expect(entry).To(ContainSubstring("clientSessionID=" + operationID.String()))

		entry, err = hooks.Entries[1].String()
		Expect(err).ToNot(HaveOccurred())
		Expect(entry).To(ContainSubstring("httpMethod=GET"))
		Expect(entry).To(ContainSubstring(`targetURI="http://127.0.0.1`))
		Expect(entry).To(ContainSubstring(`hostName="127.0.0.1:`))
		Expect(entry).To(ContainSubstring("region=" + region))
		Expect(entry).To(ContainSubstring("source=OutgoingRequestTraceLog"))
		Expect(entry).To(ContainSubstring("clientRequestID=" + clientRequestID))
		Expect(entry).To(ContainSubstring("statusCode=200"))
		Expect(entry).To(ContainSubstring("serviceRequestID=" + serviceRequestID))
		Expect(entry).To(ContainSubstring("correlationID=" + correlationID.String()))
		Expect(entry).To(ContainSubstring("contentLength=14"))
		Expect(entry).To(ContainSubstring("operationID=" + operationID.String()))
		Expect(entry).To(ContainSubstring("subscriptionID=" + subscriptionID.String()))
		Expect(entry).To(ContainSubstring("resourceGroupName=" + resourceGroupName))
		Expect(entry).To(ContainSubstring("resourceName=" + resourceName))
		Expect(entry).To(ContainSubstring("operationName=" + operationName))
		Expect(entry).To(ContainSubstring("suboperationName=" + subOperationName))
		Expect(entry).To(ContainSubstring("durationInMilliseconds="))
		Expect(entry).To(ContainSubstring("clientSessionID=" + operationID.String()))
		// no connectivity instrument, value should be empty
		Expect(entry).To(ContainSubstring("connReused= "))
		Expect(entry).To(ContainSubstring("remoteAddr= "))
		Expect(entry).To(ContainSubstring("localAddr= "))
	})

	It("should log the expected output given a successful request - Graph", func() {
		// Store log entries to assert on later
		hooks := &loggerTestHooks{}
		logger := logrus.New()
		logger.Hooks.Add(hooks)
		logger.SetFormatter(&logrus.TextFormatter{DisableColors: true})

		serviceRequestID := uuid.NewV4().String()
		content := "hello world"

		// Run a fake HTTP server to handle test request(s)
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := json.Marshal(content)

			header := w.Header()
			header.Set("Content-Type", "application/json")
			header.Set("request-id", serviceRequestID)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, string(body))
		}))
		defer svr.Close()

		// Prove a tracing span is created by getting it off of
		// the context from the next round tripper
		hookInvoked := false
		next := &roundTripperProxy{
			Hook: func(req *http.Request) {
				hookInvoked = true
			},
			Next: http.DefaultTransport,
		}

		// Build an instrumented round tripper with the test logger
		region := "test-region"

		// Instrument with connection info
		roundTripper := InstrumentWithConnection(next, region)
		roundTripperImpl := roundTripper.(*instrumentedRoundTripper)
		roundTripperImpl.loggerConstructor = func(a *log.APITracking, request *http.Request) *log.Logger {
			return log.NewOutgoingRequestLoggerForLogrus(logger.WithField("source", "test"), a, request)
		}

		operationID := uuid.NewV4()
		subscriptionID := uuid.NewV4()
		resourceGroupName := "testRG"
		resourceName := "testResource"
		operationName := "testOp"
		subOperationName := "testSubOp"

		ctx := log.WithLogger(context.Background(), log.InitializeTestLogger())
		fields := map[string]interface{}{
			"operationID":       operationID,
			"subscriptionID":    subscriptionID,
			"resourceGroupName": resourceGroupName,
			"resourceName":      resourceName,
			"operationName":     operationName,
			"subOperationName":  subOperationName,
		}
		apiTracking := log.NewAPITrackingFromParametersMap(fields)
		ctx = log.WithAPITracking(ctx, apiTracking)

		// Issue the request
		req, err := http.NewRequest("GET", svr.URL, nil)
		req = req.WithContext(ctx)
		clientRequestID := uuid.NewV4().String()
		req.Header.Add("client-request-id", clientRequestID)

		Expect(err).ToNot(HaveOccurred())
		res, err := roundTripper.RoundTrip(req)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.StatusCode).To(Equal(http.StatusOK))
		Expect(hookInvoked).To(BeTrue())

		// Two log entries should be published: one for start and one for end
		Expect(hooks.Entries).To(HaveLen(2))

		// Prove the expected fields were provided on both log entries
		entry, err := hooks.Entries[0].String()
		Expect(err).ToNot(HaveOccurred())
		Expect(entry).To(ContainSubstring("httpMethod=GET"))
		Expect(entry).To(ContainSubstring(`targetURI="http://127.0.0.1`))
		Expect(entry).To(ContainSubstring(`hostName="127.0.0.1:`))
		Expect(entry).To(ContainSubstring("region=" + region))
		Expect(entry).To(ContainSubstring("source=OutgoingRequestTraceLog"))
		Expect(entry).To(ContainSubstring("clientRequestID=" + clientRequestID))
		Expect(entry).To(ContainSubstring("operationID=" + operationID.String()))
		Expect(entry).To(ContainSubstring("subscriptionID=" + subscriptionID.String()))
		Expect(entry).To(ContainSubstring("resourceGroupName=" + resourceGroupName))
		Expect(entry).To(ContainSubstring("resourceName=" + resourceName))
		Expect(entry).To(ContainSubstring("operationName=" + operationName))
		Expect(entry).To(ContainSubstring("suboperationName=" + subOperationName))

		entry, err = hooks.Entries[1].String()
		Expect(err).ToNot(HaveOccurred())
		Expect(entry).To(ContainSubstring("httpMethod=GET"))
		Expect(entry).To(ContainSubstring(`targetURI="http://127.0.0.1`))
		Expect(entry).To(ContainSubstring(`hostName="127.0.0.1:`))
		Expect(entry).To(ContainSubstring("region=" + region))
		Expect(entry).To(ContainSubstring("source=OutgoingRequestTraceLog"))
		Expect(entry).To(ContainSubstring("clientRequestID=" + clientRequestID))
		Expect(entry).To(ContainSubstring("statusCode=200"))
		Expect(entry).To(ContainSubstring("serviceRequestID=" + serviceRequestID))
		Expect(entry).To(ContainSubstring("contentLength=14"))
		Expect(entry).To(ContainSubstring("operationID=" + operationID.String()))
		Expect(entry).To(ContainSubstring("subscriptionID=" + subscriptionID.String()))
		Expect(entry).To(ContainSubstring("resourceGroupName=" + resourceGroupName))
		Expect(entry).To(ContainSubstring("resourceName=" + resourceName))
		Expect(entry).To(ContainSubstring("operationName=" + operationName))
		Expect(entry).To(ContainSubstring("suboperationName=" + subOperationName))
		Expect(entry).To(ContainSubstring("durationInMilliseconds="))

		// with connection instrument, should tell remote IP and local port
		Expect(entry).To(ContainSubstring("remoteAddr=\"127.0.0.1:"))
		Expect(entry).To(ContainSubstring("localAddr=\"127.0.0.1:"))
		Expect(entry).To(ContainSubstring("connReused=false"))
	})
})

type loggerTestHooks struct {
	Entries []*logrus.Entry
}

func (l *loggerTestHooks) Levels() []logrus.Level { return logrus.AllLevels }

func (l *loggerTestHooks) Fire(e *logrus.Entry) error {
	l.Entries = append(l.Entries, e)
	return nil
}

type roundTripperProxy struct {
	Hook func(*http.Request)
	Next http.RoundTripper
}

func (r *roundTripperProxy) RoundTrip(req *http.Request) (*http.Response, error) {
	r.Hook(req)
	return r.Next.RoundTrip(req)
}
