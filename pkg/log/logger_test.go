package log

import (
	"bytes"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("get json logger", func() {
		apiTracking := APITracking{}
		apiTracking.operationID = expected
		apiTracking.correlationID = expected

		logger := addAPITrackingToLogger(GetGlobalLogger(), &apiTracking)

		Expect(logger.Data[OperationIDFieldName].(string)).Should(Equal(expected.String()), "The OperationID was not properly placed on the log.Entry")
		Expect(logger.Data[CorrelationIDFieldName].(string)).Should(Equal(expected.String()), "The CorrelationID was not properly placed on the log.Entry")
	})

	It("loggers roundtrip", func() {
		apiTracking := APITracking{operationID: expected}
		inputLogger := logrus.New()
		inputLogger.Out = &bytes.Buffer{}
		logger := NewLogger(&apiTracking)

		observed := logger.TraceLogger

		Expect(observed.Data[OperationIDFieldName].(string)).Should(Equal(expected.String()), "operationID does not match.")
		Expect(observed.Data[SourceFieldName].(string)).Should(Equal(TraceSource), "operationID does not match.")

		observed = logger.QosLogger

		Expect(observed.Data[OperationIDFieldName].(string)).Should(Equal(expected.String()), "operationID does not match.")
		Expect(observed.Data[SourceFieldName].(string)).Should(Equal(QosSource), "operationID does not match.")

		ctx := context.Background()

		ctx = WithLogger(ctx, logger)
		observed = GetLogger(ctx).TraceLogger

		Expect(observed.Data[OperationIDFieldName].(string)).Should(Equal(expected.String()), "operationID does not match.")
		Expect(observed.Data[SourceFieldName].(string)).Should(Equal(TraceSource), "operationID does not match.")
	})

	It("trace info", func() {
		logger, buf := initializeLogger()
		logger.TraceInfo(testMessage)

		Expect(buf.String()).Should(ContainSubstring(testMessage), "TraceInfo didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("info"), "TraceInfo didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "TraceInfo didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "TraceInfo didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "TraceInfo didn't set lineNumber field properly")
	})

	It("trace infof", func() {
		logger, buf := initializeLogger()
		logger.Infof(nil, testFormat, testFormatArg1, testFormatArg2)

		Expect(buf.String()).Should(ContainSubstring(expectedFormattedString), "traceInfof didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("info"), "traceInfof didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceInfof didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceInfof didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceInfof didn't set lineNumber field properly")
	})

	It("trace warning", func() {
		logger, buf := initializeLogger()
		logger.Warning(nil, testMessage)

		Expect(buf.String()).Should(ContainSubstring(testMessage), "traceWarning didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("warning"), "traceWarning didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceWarning didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceWarning didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceWarning didn't set lineNumber field properly")
	})

	It("trace warningf", func() {
		logger, buf := initializeLogger()
		logger.Warningf(nil, testFormat, testFormatArg1, testFormatArg2)

		Expect(buf.String()).Should(ContainSubstring(expectedFormattedString), "traceWarningf didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("warning"), "traceWarningf didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceWarningf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceWarningf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceWarningf didn't set lineNumber field properly")
	})

	It("trace debug", func() {
		logrus.SetLevel(logrus.DebugLevel)
		logger, buf := initializeLogger()
		logger.Debug(nil, testMessage)

		Expect(buf.String()).Should(ContainSubstring(testMessage), "traceDebug didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("debug"), "traceDebug didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceDebug didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceDebug didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceDebug didn't set lineNumber field properly")
	})

	It("trace debugf", func() {
		logrus.SetLevel(logrus.DebugLevel)
		logger, buf := initializeLogger()
		logger.Debugf(nil, testFormat, testFormatArg1, testFormatArg2)

		Expect(buf.String()).Should(ContainSubstring(expectedFormattedString), "traceDebugf didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("debug"), "traceDebugf didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceDebugf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceDebugf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceDebugf didn't set lineNumber field properly")
	})

	It("trace error", func() {
		logger, buf := initializeLogger()
		logger.Error(nil, testMessage)

		Expect(buf.String()).Should(ContainSubstring(testMessage), "traceError didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("error"), "traceError didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceError didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceError didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceError didn't set lineNumber field properly")
	})

	It("trace errorf", func() {
		logger, buf := initializeLogger()
		logger.Errorf(nil, testFormat, testFormatArg1, testFormatArg2)

		Expect(buf.String()).Should(ContainSubstring(expectedFormattedString), "traceErrorf didn't place message in the log")
		Expect(buf.String()).Should(ContainSubstring("error"), "traceErrorf didn't level itself to correctly")
		Expect(buf.String()).Should(ContainSubstring("fileName"), "traceErrorf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("logger_test.go"), "traceErrorf didn't set fileName field properly")
		Expect(buf.String()).Should(ContainSubstring("lineNumber"), "traceErrorf didn't set lineNumber field properly")
	})

	It("Get logger with an empty context", func() {
		logger := GetLogger(context.Background())
		Expect(logger).NotTo(BeNil())
	})

	Context("can add fields on the fly", func() {
		It("With Qos And Trace logger", func() {
			logger, buf := initializeLogger()
			loggerWithFields := logger.WithFields(map[string]interface{}{
				"additionalTestField": "something",
			})
			loggerWithFields.TraceInfo("Some important log")
			Expect(buf.String()).To(ContainSubstring("additionalTestField=something"))
		})

		It("With one logger nil", func() {
			logger, buf := initializeLogger()
			loggerWithFields := logger.WithFields(map[string]interface{}{
				"additionalTestField": "something",
			})
			loggerWithFields.TraceInfo("Some important log")
			Expect(buf.String()).To(ContainSubstring("additionalTestField=something"))
		})

		It("It does not affect the original logger instance", func() {
			logger, buf := initializeLogger()
			logger.WithFields(map[string]interface{}{
				"additionalTestField": "something",
			})
			logger.TraceInfo("Some important log")
			Expect(buf.String()).ToNot(ContainSubstring("additionalTestField=something"))
		})
	})
})

const testMessage = "testing"
const testFormat = "test formating %s %s"
const testFormatArg1 = "arg1"
const testFormatArg2 = "arg2"
const expectedFormattedString = "test formating arg1 arg2"

func initializeLogger() (*Logger, *bytes.Buffer) {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	lctx := internalNewLogger(logger.WithField("key", "value"), &APITracking{})

	return lctx, buf
}
