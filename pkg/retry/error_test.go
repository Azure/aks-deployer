package retry

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("NewTerminalError", func() {
	It("should return the expected error", func() {
		e := NewTerminalError("foo %s", "bar")
		Expect(e.Retriable).To(BeFalse())
		Expect(e.Error.Error()).To(Equal("foo bar"))
	})
})

var _ = Describe("NewRetriableError", func() {
	It("should return the expected error", func() {
		e := NewRetriableError("foo %s", "bar")
		Expect(e.Retriable).To(BeTrue())
		Expect(e.Error.Error()).To(Equal("foo bar"))
	})
})

var _ = Describe("ErrorStringerImplementation", func() {
	It("should return the expected string", func() {
		e := &Error{true, errors.New("dummy error")}
		Expect(fmt.Sprintf("%s", e)).To(Equal(fmt.Sprintf("{Retriable: %t Error: %s}", e.Retriable, e.Error)))
	})
})

var _ = Describe("hasRetryableErrorMessage", func() {
	DescribeTable("should set retryable for given error messages",
		func(err error, retryable bool) {
			Expect(hasRetryableErrorMessage(err)).To(Equal(retryable))
		},
		Entry("Error msg contains retryable error", newRetryableError(), true),
		Entry("An internal execution error occurred. Please retry later.", newRetryableError(), true),
		Entry("concurrent error", errors.New("Code=\"Conflict\" Message=\"The request failed due to conflict with a concurrent request. To resolve it, please refer to https://aka.ms/activitylog to get more details on the conflicting requests.\""), true),
		Entry("Error msg does not contains retryable error", errors.New("dummy error"), false))
})

var _ = Describe("isRetriableADALError", func() {
	It("EOF", func() {
		err := GetError(nil, fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/fad79396-0443-4fec-bb05-2d89298c67e3/resourcegroups/MC_cluster-amazon-us_cluster-amazon-mendel-us_westus?api-version=2018-05-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/3fe2b629-26bc-4fb2-89d1-1bcdf9eddaf3/oauth2/token?api-version=1.0\": EOF"))
		Expect(err.Retriable).To(BeTrue())
	})

	It("ConnectionResetByPeer", func() {
		err := GetError(nil, fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/fad79396-0443-4fec-bb05-2d89298c67e3/resourcegroups/MC_cluster-amazon-us_cluster-amazon-mendel-us_westus?api-version=2018-05-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/3fe2b629-26bc-4fb2-89d1-1bcdf9eddaf3/oauth2/token?api-version=1.0\": connection reset by peer"))
		Expect(err.Retriable).To(BeTrue())
	})

	It("EOF for azure.multiTenantSPTAuthorizer", func() {
		err := GetError(nil, fmt.Errorf("azure.multiTenantSPTAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/fad79396-0443-4fec-bb05-2d89298c67e3/resourcegroups/MC_cluster-amazon-us_cluster-amazon-mendel-us_westus?api-version=2018-05-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/3fe2b629-26bc-4fb2-89d1-1bcdf9eddaf3/oauth2/token?api-version=1.0\": EOF"))
		Expect(err.Retriable).To(BeTrue())
	})
})

func newRetryableError() error {
	return errors.New("Code=\"RetryableError\" Message=\"A retryable error occurred.\" Details=[{\"code\":\"RetryableErrorDueToAnotherOperation\",\"message\":\"Operation AllocateTenantNetworkResourcesOperation (29cb5ee9-ffaf-4950-a5c5-2774b29e254a) is updating resource 8d419d30-f984-425c-baa0-478d259d6890. The call can be retried in 12 seconds")
}

func newRetryableGRPCError() error {
	return status.Error(codes.Unknown, "Code=\"RetryableError\" Message=\"A retryable error occurred.\" Details=[{\"code\":\"RetryableErrorDueToAnotherOperation\",\"message\":\"Operation AllocateTenantNetworkResourcesOperation (29cb5ee9-ffaf-4950-a5c5-2774b29e254a) is updating resource 8d419d30-f984-425c-baa0-478d259d6890. The call can be retried in 12 seconds")
}

func newNonRetryableError() error {
	return errors.New("Code=\"OperationNotAllowed\" Message=\"Operation could not be completed as it results in exceeding approved Total Regional Cores quota. Additional details - Deployment Model: Resource Manager, Location: westeurope, Current Limit: 10, Current Usage: 28, Additional Required: 8, (Minimum) New Limit Required: 36. Submit a request for Quota increase at https://aka.ms/ProdportalCRP/?#create/Microsoft.Support/Parameters/%7B%22subId%22:%22a92e3685-71e2-4968-8e40-dbcad424b510%22,%22pesId%22:%2206bfd9d3-516b-d5c6-5802-169c800dec89%22,%22supportTopicId%22:%22e12e3d1d-7fa0-af33-c6d0-3c50df9658a3%22%7D by specifying parameters listed in the ‘Details’ section for deployment to succeed. Please read more about quota limits at https://docs.microsoft.com/en-us/azure/azure-supportability/regional-quota-requests.")
}

func newNonRetryableGRPCError() error {
	return status.Error(codes.Unauthenticated, "Code=\"OperationNotAllowed\" Message=\"Operation could not be completed as it results in exceeding approved Total Regional Cores quota. Additional details - Deployment Model: Resource Manager, Location: westeurope, Current Limit: 10, Current Usage: 28, Additional Required: 8, (Minimum) New Limit Required: 36. Submit a request for Quota increase at https://aka.ms/ProdportalCRP/?#create/Microsoft.Support/Parameters/%7B%22subId%22:%22a92e3685-71e2-4968-8e40-dbcad424b510%22,%22pesId%22:%2206bfd9d3-516b-d5c6-5802-169c800dec89%22,%22supportTopicId%22:%22e12e3d1d-7fa0-af33-c6d0-3c50df9658a3%22%7D by specifying parameters listed in the ‘Details’ section for deployment to succeed. Please read more about quota limits at https://docs.microsoft.com/en-us/azure/azure-supportability/regional-quota-requests.")
}

var _ = Describe("GetError", func() {
	When("Use common retry code", func() {
		It("should return retry code with expected retriable decision", func() {
			for _, code := range statusCodesForRetry {
				retryCodeGetErrorTestHelper(code, true)
			}

			retryCodeGetErrorTestHelper(http.StatusNetworkAuthenticationRequired, false)
		})
	})

	When("Error code isn't retryable but msg is", func() {
		It("should return retryable", func() {
			response := &http.Response{
				StatusCode: http.StatusProxyAuthRequired, //non-retryable code
			}

			err := GetError(response, newRetryableError())
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("With extra retry code", func() {
		It("should return retry code with expected retriable decision", func() {
			extraRetriableCodes := []int{http.StatusPreconditionFailed}
			for _, code := range statusCodesForRetry {
				retryCodeGetErrorTestHelper(code, true, extraRetriableCodes...)
			}

			retryCodeGetErrorTestHelper(http.StatusNetworkAuthenticationRequired, false, extraRetriableCodes...)
			retryCodeGetErrorTestHelper(http.StatusPreconditionFailed, true, extraRetriableCodes...)
		})

		It("409 for extra retry", func() {
			resp := &http.Response{
				StatusCode: 409,
			}
			err := GetError(resp, fmt.Errorf("conflict error"), 409)
			Expect(err.Retriable).To(BeTrue())
		})

		It("409 not for extra retry", func() {
			resp := &http.Response{
				StatusCode: 409,
			}
			err := GetError(resp, fmt.Errorf("conflict error"))
			Expect(err.Retriable).To(BeFalse())
		})
	})

	When("given EOF", func() {
		It("should be retriable", func() {
			err := GetError(nil, io.EOF)
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given wrapped EOF", func() {
		It("should be retriable", func() {
			err := GetError(nil, fmt.Errorf("foo %w", io.EOF))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given unexpected EOF", func() {
		It("should be retriable", func() {
			err := GetError(nil, io.ErrUnexpectedEOF)
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given wrapped unexpected EOF", func() {
		It("should be retriable", func() {
			err := GetError(nil, fmt.Errorf("foo %w", io.ErrUnexpectedEOF))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("Connectivity error happens", func() {
		It("should return as retry-able error", func() {
			err := GetError(nil, errors.New("http2: server sent GOAWAY and closed the connection"))
			Expect(err.Retriable).To(BeTrue())
		})

		// we found GOAWAY error with NO_ERROR, and we do receive a 200 response
		// http2: server sent GOAWAY and closed the connection; LastStreamID=1999, ErrCode=NO_ERROR, debug=""
		It("should return as retry-able error, even with valid response", func() {
			err := GetError(&http.Response{StatusCode: http.StatusOK}, errors.New("http2: server sent GOAWAY and closed the connection"))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("GetControlPlaneV1 error happens", func() {
		It("should return as retry-able error", func() {
			err := GetError(nil, errors.New("Resolve control plane goal failed: '{Retriable: false Error: GetControlPlaneV1 returned error: stream error: stream ID 613; INTERNAL_ERROR}'"))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("Status code isn't retryable", func() {
		It("StatusCode is client error related", func() {
			response := &http.Response{
				StatusCode: http.StatusConflict, //non-retryable code
			}

			err := GetError(response, newNonRetryableError())
			Expect(err.Retriable).NotTo(BeTrue())
		})

		It("Http response is nil, but have non-retryable error message", func() {
			err := GetError(nil, newNonRetryableError())
			Expect(err.Retriable).NotTo(BeTrue())
		})

		It("Http response is nil, but have retryable error message", func() {
			err := GetError(nil, newRetryableError())
			Expect(err.Retriable).To(BeTrue())
		})

		It("Http response is nil, also have empty message, not retry", func() {
			err := GetError(nil, fmt.Errorf(""))
			Expect(err.Retriable).NotTo(BeTrue())
		})
	})

	When("Status code is retryable", func() {
		It("StatusCode is not client error, but retryable status code, retry", func() {
			response := &http.Response{
				StatusCode: http.StatusInternalServerError, //retryable code
			}

			err := GetError(response, newNonRetryableError())
			Expect(err.Retriable).To(BeTrue())
		})

		It("StatusCode is client error, but retryable statusCode , retry.", func() {
			response := &http.Response{
				StatusCode: http.StatusRequestTimeout, //retryable code
			}

			err := GetError(response, newNonRetryableError())
			Expect(err.Retriable).To(BeTrue())
		})

		It("StatusCode is retryable client error, with empty error string, retry", func() {
			response := &http.Response{
				StatusCode: http.StatusPreconditionFailed, //non-etryable code
			}
			err := GetError(response, fmt.Errorf(""), http.StatusPreconditionFailed)
			Expect(err.Retriable).To(BeTrue())
		})
	})
})

var _ = Describe("GetGRPCError", func() {
	When("Use common retry code", func() {
		It("should return retry code with expected retryable decision", func() {
			for _, code := range grpcStatusCodesForRetry {
				retryCodeGetGRPCErrorTestHelper(code, true)
			}

			retryCodeGetGRPCErrorTestHelper(codes.NotFound, false)
		})
	})

	When("Error code isn't retryable but msg is", func() {
		It("should return retryable", func() {
			err := GetGRPCError(newRetryableGRPCError())
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("With extra retry code", func() {
		It("should return retry code with expected retriable decision", func() {
			extraRetriableCodes := []codes.Code{codes.NotFound}
			for _, code := range grpcStatusCodesForRetry {
				retryCodeGetGRPCErrorTestHelper(code, true, extraRetriableCodes...)
			}

			retryCodeGetGRPCErrorTestHelper(codes.FailedPrecondition, false, extraRetriableCodes...)
			retryCodeGetGRPCErrorTestHelper(codes.NotFound, true, extraRetriableCodes...)
		})
	})

	When("given EOF", func() {
		It("should be retriable", func() {
			err := GetGRPCError(status.Error(codes.Unknown, io.EOF.Error()))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given wrapped EOF", func() {
		It("should be retriable", func() {
			err := GetGRPCError(status.Error(codes.Unknown, fmt.Errorf("foo %w", io.EOF).Error()))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given unexpected EOF", func() {
		It("should be retriable", func() {
			err := GetGRPCError(status.Error(codes.Unknown, io.ErrUnexpectedEOF.Error()))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("given wrapped unexpected EOF", func() {
		It("should be retriable", func() {
			err := GetGRPCError(status.Error(codes.Unknown, fmt.Errorf("foo %w", io.ErrUnexpectedEOF).Error()))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("Connectivity error happens", func() {
		It("should return as retry-able error", func() {
			err := GetGRPCError(status.Error(codes.Unavailable, "http2: server sent GOAWAY and closed the connection"))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("GetControlPlaneV1 error happens", func() {
		It("should return as retry-able error", func() {
			err := GetGRPCError(status.Error(codes.Internal, "Resolve control plane goal failed: '{Retriable: false Error: GetControlPlaneV1 returned error: stream error: stream ID 613; INTERNAL_ERROR}'"))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("Status code isn't retryable", func() {
		It("StatusCode is client error related", func() {
			err := GetGRPCError(newNonRetryableGRPCError())
			Expect(err.Retriable).NotTo(BeTrue())
		})

		It("have non-retryable error message", func() {
			err := GetGRPCError(newNonRetryableGRPCError())
			Expect(err.Retriable).NotTo(BeTrue())
		})

		It("have retryable error message", func() {
			err := GetGRPCError(newRetryableGRPCError())
			Expect(err.Retriable).To(BeTrue())
		})

		It("have empty message, should retry", func() {
			err := GetGRPCError(fmt.Errorf(""))
			Expect(err.Retriable).To(BeTrue())
		})
	})

	When("Status code is retryable", func() {
		It("StatusCode is not client error, but retryable status code, retry", func() {
			err := GetGRPCError(status.Error(codes.AlreadyExists, newRetryableGRPCError().Error()))
			Expect(err.Retriable).To(BeTrue())
		})

		It("StatusCode is client error, but retryable statusCode , retry.", func() {
			err := GetGRPCError(status.Error(codes.Canceled, newRetryableGRPCError().Error()))
			Expect(err.Retriable).To(BeTrue())
		})
	})
})

func retryCodeGetErrorTestHelper(responseCode int, expectedRetriable bool, extraRetriableCodes ...int) {
	response := &http.Response{
		StatusCode: responseCode,
	}

	err := GetError(response, errors.New("foo bar error happened"), extraRetriableCodes...)
	Expect(err.Retriable).To(Equal(expectedRetriable), fmt.Sprintf("%d should be retriable: %t", responseCode, expectedRetriable))
}

func retryCodeGetGRPCErrorTestHelper(code codes.Code, expectedRetryable bool, extraRetryableCodes ...codes.Code) {
	err := GetGRPCError(status.Error(code, "foo bar error happened"), extraRetryableCodes...)
	Expect(err.Retriable).To(Equal(expectedRetryable), fmt.Sprintf("%d should be retriable: %t", code, expectedRetryable))
}
