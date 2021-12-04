package log

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.opencensus.io/trace"
)

var _ = Describe("Logging with span id", func() {
	It("can still log without context", func() {
		l, buf := initializeLogger()
		l.TraceInfo("testMySimpleLog")
		Expect(buf.String()).To(ContainSubstring("testMySimpleLog"))
	})
	Context("context has no spanContext", func() {
		It("logs without extra info when nil", func() {
			l, buf := initializeLogger()
			l.Info(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("testMySimpleLog"))
		})
		It("logs without extra info when no span", func() {
			l, buf := initializeLogger()
			l.Info(context.TODO(), "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("testMySimpleLog"))
		})
	})
	Context("context has spancontext", func() {
		It("logs with span info", func() {
			l, buf := initializeLogger()
			ctx := context.TODO()
			ctx, span := trace.StartSpan(ctx, "some trace")
			l.Info(ctx, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("spanID=%s", span.SpanContext().SpanID))
			Expect(buf.String()).To(ContainSubstring("traceID=%s", span.SpanContext().TraceID))
		})
	})

	Context("test level logging", func() {
		It("logs debug", func() {
			l, buf := initializeLogger()
			l.Debug(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=debug"))
		})
		It("logs info", func() {
			l, buf := initializeLogger()
			l.Info(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=info"))
		})
		It("logs warnings", func() {
			l, buf := initializeLogger()
			l.Warning(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=warning"))
		})
		It("logs error", func() {
			l, buf := initializeLogger()
			l.Error(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=error"))
		})
		It("logs debug", func() {
			l, buf := initializeLogger()
			l.Debugf(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=debug"))
		})
		It("logs info", func() {
			l, buf := initializeLogger()
			l.Infof(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=info"))
		})
		It("logs warnings", func() {
			l, buf := initializeLogger()
			l.Warningf(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=warning"))
		})
		It("logs error", func() {
			l, buf := initializeLogger()
			l.Errorf(nil, "testMySimpleLog")
			Expect(buf.String()).To(ContainSubstring("level=error"))
		})
	})

	Context("test source file", func() {
		It("should use the correct caller source file", func() {
			l, buf := initializeLogger()
			l.Debug(nil, "testMySimpleLog")
			// Should log this file, spanlogger_test.go, as source file.
			Expect(buf.String()).To(ContainSubstring("spanlogger_test.go"))
		})
	})
})
