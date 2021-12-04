package log

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pkg/errors"
)

var _ = Describe("Test Get Stack Trace from Errors", func() {
	It("should handle nil error", func() {
		stackStr := getStackTraceStr(nil)
		Expect(stackStr).To(BeEmpty())
	})

	It("should handle error without stacktrace", func() {
		err := fmt.Errorf("this is not a pkg/errors")
		stackStr := getStackTraceStr(err)
		Expect(stackStr).To(BeEmpty())
	})

	It("should handle using new error", func() {
		err := errors.New("testing stack error")
		stackStr := getStackTraceStr(err)
		Expect(stackStr).ToNot(BeEmpty())
	})

	It("should handle wrapped error", func() {
		err := errors.New("testing stack error")
		err = errors.Wrap(err, "testing wrapping")

		stackStr := getStackTraceStr(err)
		Expect(stackStr).ToNot(BeEmpty())
	})

	It("should handle multi-wrapped  error", func() {
		err := errors.New("testing stack error")
		err = errors.Wrap(err, "testing wrapping")
		err = errors.WithStack(err)

		stackStr := getStackTraceStr(err)
		Expect(stackStr).ToNot(BeEmpty())
	})
})
