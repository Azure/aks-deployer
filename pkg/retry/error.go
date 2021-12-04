package retry

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Azure/aks-deployer/pkg/armerror"
)

var statusCodesForRetry []int
var retryableErrorMessages = [...]string{
	"A retryable error occurred",
	"RetryableError",
	"Etag mismatch",
	"etag not expected",
	"stream error",
	"http2: server sent GOAWAY and closed the connection", // we got http connectivity error.
	"http: can't write HTTP request on broken connection",
	"connection reset by peer",
	"use of closed network connection",
	"Please retry later.",
	"missing content-type field",
	"The request failed due to conflict with a concurrent request",
}

// grpc status codes we retry for
var grpcStatusCodesForRetry = []codes.Code{
	codes.Canceled,
	codes.DeadlineExceeded,
	codes.Unknown,
	codes.Internal,
	codes.Unavailable,
	codes.ResourceExhausted,
}

// NewError creates an error
func NewError(retriable bool, err error) *Error {
	return &Error{
		Retriable: retriable,
		Error:     err,
	}
}

// NewTerminalError returns a new non-retriable error with an fmt formatted message.
func NewTerminalError(msg string, args ...interface{}) *Error {
	return &Error{
		Retriable: false,
		Error:     fmt.Errorf(msg, args...),
	}
}

// NewRetriableError returns a new retriable error with an fmt formatted message.
func NewRetriableError(msg string, args ...interface{}) *Error {
	return &Error{
		Retriable: true,
		Error:     fmt.Errorf(msg, args...),
	}
}

func init() {
	// in go-autorest SDK vendor\github.com\Azure\go-autorest\autorest\sender.go Ln 242
	// if ARM returns http.StatusTooManyRequests, the sender doesn't increase the retry attempt count,
	// so the arm_client will keep retrying forever until it get a status code other than 429
	for _, code := range autorest.StatusCodesForRetry {
		if code != http.StatusTooManyRequests {
			statusCodesForRetry = append(statusCodesForRetry, code)
		}
	}
}

// Error represents an error returned by APIs
type Error struct {
	Retriable bool
	Error     error
}

// GetError gets common error based on error and http response
func GetError(response *http.Response, err error, extraRetryCodes ...int) *Error {
	if err == nil {
		return nil
	}

	// In both cases response is cracked or not cracked, we need to double check error string to find out retry or not.
	if hasRetryableErrorMessage(err) {
		return NewError(true, err)
	}

	if response != nil { // if StatusCode exist, then statusCode take priority.
		retriableStatusCode := append(statusCodesForRetry, extraRetryCodes...)
		retryable := autorest.ResponseHasStatusCode(response, retriableStatusCode...)
		return NewError(retryable, err)
	}

	if _, isClientError := armerror.IsErrorStringClientError(err.Error()); isClientError {
		return NewError(false, err)
	}

	// we got http connectivity error.
	// this part is copied from /k8s.io/apimachinery/pkg/util/net/http.go IsProbableEOF
	if errors.Is(err, io.EOF) {
		return NewError(true, err)
	}

	// more details on how this error is propagated from the client: https://stackoverflow.com/questions/51105792/golang-unexpected-eof
	if errors.Is(err, io.ErrUnexpectedEOF) {
		return NewError(true, err)
	}

	if isADALRetriableError(err) {
		return NewError(true, err)
	}

	return NewError(false, err)
}

// GetGRPCError gets common error based on error and grpc status codes
func GetGRPCError(err error, extraRetryableCodes ...codes.Code) *Error {
	if err == nil {
		return nil
	}

	// In both cases response is cracked or not cracked, we need to double check error string to find out retry or not.
	if hasRetryableErrorMessage(err) {
		return NewError(true, err)
	}

	// with REST, response status code dictates retryable or not but we always have status codes with GRPC
	if isGRPCStatusCodeRetryable(status.Code(err), extraRetryableCodes...) {
		return NewError(true, err)
	}

	// we got http connectivity error.
	// this part is copied from /k8s.io/apimachinery/pkg/util/net/http.go IsProbableEOF
	if errors.Is(err, io.EOF) {
		return NewError(true, err)
	}

	if _, isClientError := armerror.IsErrorStringClientError(err.Error()); isClientError {
		return NewError(false, err)
	}

	if isADALRetriableError(err) {
		return NewError(true, err)
	}

	return NewError(false, err)
}

func isGRPCStatusCodeRetryable(code codes.Code, extraRetryableCodes ...codes.Code) bool {
	allRetryableCodes := append(grpcStatusCodesForRetry, extraRetryableCodes...)
	for _, retryCode := range allRetryableCodes {
		if retryCode == code {
			return true
		}
	}

	return false
}

func isADALRetriableError(err error) bool {
	if !strings.Contains(err.Error(), armerror.ADALResourceWithinError) {
		return false
	}

	var connectivityErrors = []string{
		"context deadline exceeded",
		"connection refused",
		"connection reset by peer",
		"connection timed out",
		"TLS handshake timeout",
		"i/o timeout",
		"no such host",
		"EOF",
		"context canceled",
	}

	for _, e := range connectivityErrors {
		if strings.Contains(err.Error(), e) {
			return true
		}
	}

	return false
}

func hasRetryableErrorMessage(err error) bool {
	errMsg := err.Error()
	for _, retryable := range retryableErrorMessages {
		if strings.Contains(errMsg, retryable) {
			return true
		}
	}
	return false
}

// Status get status from an error
func (err *Error) Status() Status {
	if err.Retriable {
		return NeedRetry
	}
	return Failed
}

func (err Error) String() string {
	return fmt.Sprintf("{Retriable: %t Error: %s}", err.Retriable, err.Error)
}
