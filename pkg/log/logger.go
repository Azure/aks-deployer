// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// logger.go contains the code for making a logger, storing it on the context,
// and pulling it off the context for use. It also has some helpers for adding
// tracking fields on to the logger

package log

import (
	"context"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"github.com/Azure/aks-deployer/pkg/deploy"
	"github.com/Azure/aks-deployer/pkg/version"
)

type loggingKey string

const loggerKey loggingKey = "AcsLogger"

// Logger is a type that can log trace logs or
type Logger struct {
	TraceLogger *logrus.Entry
	QosLogger   *logrus.Entry
}

// Exported field names to use in the structured logger
const (
	CorrelationIDFieldName     = "correlationID"
	ClientSessionIDFieldName   = "clientSessionID"
	OperationIDFieldName       = "operationID"
	SourceFieldName            = "source"
	ClientRequestIDFieldName   = "clientRequestID"
	LatencyFieldName           = "latency"
	SubscriptionIDFieldName    = "subscriptionID"
	IsInternalSubFieldName     = "isInternalSub"
	ResourceGroupNameFieldName = "resourceGroupName"
	ResourceNameFieldName      = "resourceName"
	AgentPoolNameFieldName     = "agentPoolName"
	OperationNameFieldName     = "operationName"
	RegionFieldName            = "region"
	InsertionTimeFieldName     = "insertionTime"
	RetryAttemptFieldName      = "retryAttempt"

	SuboperationNameFieldName = "suboperationName"
	TruncatedLogFieldName     = "truncated"

	// user-agent field is already used by grpc request so a different namespace is needed
	UserAgentGRPC = "user-agent-grpc"
	HostNameGRPC  = "host-name-grpc"
)

// Field names to use in the structured logger
const (
	stackTraceFieldName     = "stacktrace"
	serviceBuildFieldName   = "serviceBuild"
	clientApplicationID     = "clientApplicationID"
	clientPrincipalName     = "clientPrincipalName"
	httpMethodFieldName     = "httpMethod"
	targetURIFieldName      = "targetURI"
	userAgentFieldName      = "userAgent"
	acceptLanguageFieldName = "acceptLanguage"

	fileNameFieldName   = "fileName"
	lineNumberFieldName = "lineNumber"
	apiVersionFieldName = "apiVersion"
	hostNameFieldName   = "hostName"

	accessProfileRoleName = "accessProfileRoleName"

	clientRemoteAddrFieldName = "clientRemoteAddr"
	xForwardedForFieldName    = "xForwardedFor"
	xRealIPFieldName          = "xRealIP"
	routeNameFieldName        = "routeName"

	// Extra environment fields for IFx audit logs
	epochFieldName = "env_epoch"

	eventCategoryFieldName = "Category"
)

// Event sources the logger uses that connect to the mdsd.d config to control the tables in geneva this gets uploaded to
const (
	QosSource   = "AcsQOSLog"
	TraceSource = "AcsTraceLog"

	contextlessTraceSource = "AcsContextlessTraceLog"

	outgoingRequestTraceSource = "OutgoingRequestTraceLog"
	incomingRequestTraceSource = "IncomingRequestTraceLog"

	latencyTraceEventSource      = "LatencyTraceEvent"
	queueWatcherTraceSource      = "QueueWatcherLog"
	msiCredentialRefresherSource = "MSICredentialRefresherLog"
	msiConnectorSource           = "MSIConnectorLog"
	addonTokenReconcilerSource   = "AddonTokenReconcilerLog"
)

const (
	NanoSecondToMillisecondConversionFactor = 1000 * 1000
	TimeFormat                              = "2006-01-02T15:04:05.999999999Z07:00"
)

var (
	logger *logrus.Entry
)

// GetGlobalLogger gets the global logger
func GetGlobalLogger() *logrus.Entry {
	epoch, _ := getEpochRandomString()
	return logger.WithField(epochFieldName, epoch)
}

// GetContextlessTraceLogger returns a logger that can be used outside of a context
func GetContextlessTraceLogger() *Logger {
	return &Logger{
		TraceLogger: GetGlobalLogger().WithField(SourceFieldName, contextlessTraceSource),
	}
}

// TODO(thgamble) - 5/5/20 - possible move? only used in ./common/cmd/queuewatcher/queuewatcher.go
// GetQueueWatcherLogger returns a logger to use in the queue watcher
func GetQueueWatcherLogger() *Logger {
	return &Logger{
		TraceLogger: GetGlobalLogger().WithField(SourceFieldName, queueWatcherTraceSource),
	}
}

func newLogger(formatter logrus.Formatter) *logrus.Entry {
	logger := logrus.New()
	logger.Formatter = formatter
	return logger.WithField(serviceBuildFieldName, version.String())
}

// NewLoggerWithCustomWriter used for testing to overwrite the logger output used by codebase
// for verification of log fields
func NewLoggerWithCustomWriter(formatter logrus.Formatter, customWriter io.Writer) *Logger {
	entryLogger := logrus.New()
	entryLogger.Formatter = formatter
	entryLogger.Out = customWriter
	logger = logrus.NewEntry(entryLogger)
	return &Logger{
		TraceLogger: logger.WithField(serviceBuildFieldName, version.String()),
	}
}

func withCallerInfo(logger *logrus.Entry) *logrus.Entry {
	_, file, line, _ := runtime.Caller(3)
	fields := make(map[string]interface{})
	fields[fileNameFieldName] = file
	fields[lineNumberFieldName] = line
	return logger.WithFields(fields)
}

// addAPITrackingToLogger adds the api tracking info to a logger
func addAPITrackingToLogger(logger *logrus.Entry, apiTracking *APITracking) *logrus.Entry {
	fields := make(map[string]interface{})

	fields[OperationIDFieldName] = apiTracking.GetOperationID().String()
	fields[CorrelationIDFieldName] = apiTracking.GetCorrelationID().String()
	fields[ClientRequestIDFieldName] = apiTracking.GetClientRequestID().String()
	fields[ClientSessionIDFieldName] = apiTracking.GetClientSessionID()
	fields[clientApplicationID] = apiTracking.GetClientAppID()
	fields[clientPrincipalName] = apiTracking.GetClientPrincipalName()
	fields[userAgentFieldName] = apiTracking.GetUserAgent()
	fields[acceptLanguageFieldName] = apiTracking.GetAcceptLanguage()
	fields[SubscriptionIDFieldName] = apiTracking.GetSubscriptionID()
	fields[ResourceGroupNameFieldName] = apiTracking.GetResourceGroupName()
	fields[ResourceNameFieldName] = apiTracking.GetResourceName()
	fields[AgentPoolNameFieldName] = apiTracking.GetAgentPoolName()
	fields[OperationNameFieldName] = apiTracking.GetOperationName()
	fields[eventCategoryFieldName] = apiTracking.GetOperationCategory()
	fields[SuboperationNameFieldName] = apiTracking.GetSubOperationName()
	fields[apiVersionFieldName] = apiTracking.GetAPIVersion()
	fields[hostNameFieldName] = apiTracking.GetHost()
	fields[accessProfileRoleName] = apiTracking.GetAccessProfileRoleName()
	fields[targetURIFieldName] = apiTracking.GetTargetURI()
	fields[httpMethodFieldName] = apiTracking.GetHttpMethod()
	fields[RegionFieldName] = deploy.GetLoggingRegion(apiTracking.GetRegion())

	return logger.WithFields(fields)
}

func appendOutgoingRequestFields(logger *logrus.Entry, apiTracking *APITracking, request *http.Request) *logrus.Entry {
	fields := map[string]interface{}{}

	if apiTracking != nil {
		fields = map[string]interface{}{
			OperationIDFieldName:       apiTracking.GetOperationID().String(),
			SubscriptionIDFieldName:    apiTracking.GetSubscriptionID().String(),
			ResourceGroupNameFieldName: apiTracking.GetResourceGroupName(),
			ResourceNameFieldName:      apiTracking.GetResourceName(),
			AgentPoolNameFieldName:     apiTracking.GetAgentPoolName(),
			OperationNameFieldName:     apiTracking.GetOperationName(),
			SuboperationNameFieldName:  apiTracking.GetSubOperationName(),
			ClientRequestIDFieldName:   apiTracking.GetClientRequestID().String(),
			ClientSessionIDFieldName:   request.Header.Get(RequestClientSessionIDHeader),
			userAgentFieldName:         apiTracking.GetUserAgent(),
			hostNameFieldName:          apiTracking.GetHost(),
			targetURIFieldName:         apiTracking.GetTargetURI(),
			httpMethodFieldName:        apiTracking.GetHttpMethod(),
			RegionFieldName:            deploy.GetLoggingRegion(apiTracking.GetRegion()),
		}
	}

	return logger.WithFields(fields)
}

// NewGRPCOutgoingRequestLogger creates an outgoing request logger for gRPC requests
func NewGRPCOutgoingRequestLogger(ctx context.Context, fullMethod string) *Logger {
	logrus := GetGlobalLogger()

	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		fields := map[string]interface{}{
			OperationIDFieldName:       getFromMetadata(md, RequestAcsOperationIDHeader),
			SubscriptionIDFieldName:    getFromMetadata(md, SubscriptionIDFieldName),
			ResourceGroupNameFieldName: getFromMetadata(md, ResourceGroupNameFieldName),
			ResourceNameFieldName:      getFromMetadata(md, ResourceNameFieldName),
			AgentPoolNameFieldName:     getFromMetadata(md, AgentPoolNameFieldName),
			OperationNameFieldName:     getFromMetadata(md, OperationNameFieldName),
			SubOperationName:           getFromMetadata(md, SubOperationName),
			ClientRequestIDFieldName:   getFromMetadata(md, RequestARMClientRequestIDHeader),
			ClientSessionIDFieldName:   getFromMetadata(md, RequestClientSessionIDHeader),
			ClientApplicationID:        getFromMetadata(md, ClientApplicationID),
			CorrelationIDFieldName:     getFromMetadata(md, RequestCorrelationIDHeader),
			userAgentFieldName:         getFromMetadata(md, UserAgentGRPC),
			xForwardedForFieldName:     getFromMetadata(md, ForwardedForHeader),
			xRealIPFieldName:           getFromMetadata(md, RealIPHeader),
			targetURIFieldName:         getFromMetadata(md, targetURIFieldName),
			hostNameFieldName:          getFromMetadata(md, HostNameGRPC),
			RetryAttemptFieldName:      getFromMetadata(md, RetryAttemptHeader),
			clientRemoteAddrFieldName:  getClientRemoteAddr(ctx),
			routeNameFieldName:         fullMethod,
		}
		logrus = logrus.WithFields(fields)
	}

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, outgoingRequestTraceSource),
	}
}

// NewOutgoingRequestLogger creates an outgoing request logger
func NewOutgoingRequestLogger(apiTracking *APITracking, request *http.Request) *Logger {
	logrus := GetGlobalLogger()
	return NewOutgoingRequestLoggerForLogrus(logrus, apiTracking, request)
}

// NewOutgoingRequestLoggerForLogrus creates an outgoing request logger from a
// Logrus logger.
func NewOutgoingRequestLoggerForLogrus(logrus *logrus.Entry, apiTracking *APITracking, request *http.Request) *Logger {
	logrus = appendOutgoingRequestFields(logrus, apiTracking, request)

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, outgoingRequestTraceSource),
	}
}

// NewIncomingRequestLogger creates an incoming request logger
func NewIncomingRequestLogger(apiTracking *APITracking, req *http.Request, routeName string) *Logger {
	logrus := GetGlobalLogger()

	if apiTracking != nil {
		fields := map[string]interface{}{
			OperationIDFieldName:       apiTracking.GetOperationID().String(),
			SubscriptionIDFieldName:    apiTracking.GetSubscriptionID().String(),
			ResourceGroupNameFieldName: apiTracking.GetResourceGroupName(),
			ResourceNameFieldName:      apiTracking.GetResourceName(),
			AgentPoolNameFieldName:     apiTracking.GetAgentPoolName(),
			OperationNameFieldName:     apiTracking.GetOperationName(),
			SuboperationNameFieldName:  apiTracking.GetSubOperationName(),
			ClientRequestIDFieldName:   apiTracking.GetClientRequestID().String(),
			RegionFieldName:            deploy.GetLoggingRegion(apiTracking.GetRegion()),
			ClientSessionIDFieldName:   apiTracking.GetClientSessionID(),
			clientApplicationID:        apiTracking.GetClientAppID(),
			CorrelationIDFieldName:     apiTracking.GetCorrelationID().String(),
		}

		logrus = logrus.WithFields(fields)
	}

	if req != nil {
		fields := map[string]interface{}{
			httpMethodFieldName:       req.Method,
			targetURIFieldName:        req.URL.String(),
			hostNameFieldName:         req.Host,
			clientRemoteAddrFieldName: req.RemoteAddr,
			userAgentFieldName:        req.Header.Get(UserAgentHeader),
			xForwardedForFieldName:    req.Header.Get(ForwardedForHeader),
			xRealIPFieldName:          req.Header.Get(RealIPHeader),
			routeNameFieldName:        routeName,
		}

		logrus = logrus.WithFields(fields)
	}

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, incomingRequestTraceSource),
	}
}

// NewIncomingGRPCRequestLogger creates a request logger for GRPC
func NewIncomingGRPCRequestLogger(ctx context.Context, req interface{}, fullMethod, region string) *Logger {
	logrus := GetGlobalLogger()

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fields := map[string]interface{}{
			OperationIDFieldName:       getFromMetadata(md, RequestAcsOperationIDHeader),
			SubscriptionIDFieldName:    getFromMetadata(md, SubscriptionIDFieldName),
			ResourceGroupNameFieldName: getFromMetadata(md, ResourceGroupNameFieldName),
			ResourceNameFieldName:      getFromMetadata(md, ResourceNameFieldName),
			AgentPoolNameFieldName:     getFromMetadata(md, AgentPoolNameFieldName),
			OperationNameFieldName:     getFromMetadata(md, OperationNameFieldName),
			SubOperationName:           getFromMetadata(md, SubOperationName),
			ClientRequestIDFieldName:   getFromMetadata(md, RequestARMClientRequestIDHeader),
			ClientSessionIDFieldName:   getFromMetadata(md, RequestClientSessionIDHeader),
			ClientApplicationID:        getFromMetadata(md, ClientApplicationID),
			CorrelationIDFieldName:     getFromMetadata(md, RequestCorrelationIDHeader),
			userAgentFieldName:         getFromMetadata(md, UserAgentGRPC),
			xForwardedForFieldName:     getFromMetadata(md, ForwardedForHeader),
			xRealIPFieldName:           getFromMetadata(md, RealIPHeader),
			targetURIFieldName:         getFromMetadata(md, targetURIFieldName),
			hostNameFieldName:          getFromMetadata(md, HostNameGRPC),
			RetryAttemptFieldName:      getFromMetadata(md, RetryAttemptHeader),
			clientRemoteAddrFieldName:  getClientRemoteAddr(ctx),
			RegionFieldName:            region,
			routeNameFieldName:         fullMethod,
		}
		logrus = logrus.WithFields(fields)
	}

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, incomingRequestTraceSource),
	}
}

// getFromMetadata gets a value out of gRPC context metadata or returns an empty string
func getFromMetadata(md metadata.MD, k string) string {
	val := md.Get(k)
	if len(val) > 0 {
		return val[0]
	}

	return ""
}

// getClientRemoteAddr returns the gRPC Remote Client IP Address
func getClientRemoteAddr(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if ok {
		return p.Addr.String()
	}

	return ""
}

// NewServiceLogger creates a logger for a specific service
func NewServiceLogger(sourceName string, fields map[string]interface{}) *Logger {
	logrus := GetGlobalLogger()

	logrus = logrus.WithFields(fields)

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, sourceName),
	}
}

// NewMSICredentialRefresherLogger creates a logger supposed to be used by msi credential refresher
func NewMSICredentialRefresherLogger() *Logger {
	return &Logger{
		TraceLogger: GetGlobalLogger().WithField(SourceFieldName, msiCredentialRefresherSource),
	}
}

// NewMSIConnectorLogger creates a logger supposed to be used by msi connector
func NewMSIConnectorLogger(apiTracking *APITracking) *Logger {
	logrus := GetGlobalLogger()
	if apiTracking != nil {
		fields := map[string]interface{}{
			SubscriptionIDFieldName:    apiTracking.GetSubscriptionID().String(),
			ResourceGroupNameFieldName: apiTracking.GetResourceGroupName(),
			ResourceNameFieldName:      apiTracking.GetResourceName(),
			ClientRequestIDFieldName:   apiTracking.GetClientRequestID().String(),
			RegionFieldName:            deploy.GetLoggingRegion(apiTracking.GetRegion()),
		}
		logrus = logrus.WithFields(fields)
	}

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, msiConnectorSource),
	}
}

// NewAddonTokenReconcilerLogger creates a logger supposed to be used by addon token reconciler
func NewAddonTokenReconcilerLogger(apiTracking *APITracking) *Logger {
	logrus := GetGlobalLogger()
	if apiTracking != nil {
		fields := map[string]interface{}{
			SubscriptionIDFieldName:    apiTracking.GetSubscriptionID().String(),
			ResourceGroupNameFieldName: apiTracking.GetResourceGroupName(),
			ResourceNameFieldName:      apiTracking.GetResourceName(),
			ClientRequestIDFieldName:   apiTracking.GetClientRequestID().String(),
			RegionFieldName:            deploy.GetLoggingRegion(apiTracking.GetRegion()),
		}
		logrus = logrus.WithFields(fields)
	}

	return &Logger{
		TraceLogger: logrus.WithField(SourceFieldName, addonTokenReconcilerSource),
	}
}

// NewLogger makes a new Logger that can log trace and qos events
func NewLogger(apiTracking *APITracking) *Logger {
	logrus := GetGlobalLogger()
	return internalNewLogger(logrus, apiTracking)
}

func internalNewLogger(logrus *logrus.Entry, apiTracking *APITracking) *Logger {
	loggerEntry := addAPITrackingToLogger(logrus, apiTracking)
	return &Logger{
		loggerEntry.WithField(SourceFieldName, TraceSource),
		loggerEntry.WithField(SourceFieldName, QosSource),
	}
}

func internalNewDefaultLogger() *Logger {
	return &Logger{
		logrus.WithField(SourceFieldName, TraceSource),
		logrus.WithField(SourceFieldName, QosSource),
	}
}

// WithLogger takes a context and logger then returns a child context with the logger attached
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLoggerWithFallback(ctx context.Context, fallback func() *Logger) *Logger {
	retVal, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return fallback()
	}
	return retVal
}

// GetLogger pulls a logger off the context and returns it
func GetLogger(ctx context.Context) *Logger {
	retVal, ok := ctx.Value(loggerKey).(*Logger)

	if !ok {
		retVal = internalNewDefaultLogger()
		retVal.Warning(ctx, "Can't get logger. Use a default one instead.")
	}

	return retVal
}

// NewLoggerWithFields returns a new logger with additional fields
func NewLoggerWithFields(logger *Logger, fields map[string]interface{}) *Logger {
	return &Logger{
		TraceLogger: logger.TraceLogger.WithFields(fields),
	}
}

// WithField returns a new logger with the additional field, allowing to add context information on the fly
func (logger *Logger) WithField(key string, value interface{}) *Logger {
	return logger.WithFields(map[string]interface{}{key: value})
}

// WithFields returns a new logger with additional fields, allowing to add context information on the fly
func (logger *Logger) WithFields(fields map[string]interface{}) *Logger {
	newLogger := &Logger{}
	if logger.QosLogger != nil {
		newLogger.QosLogger = logger.QosLogger.WithFields(fields)
	}
	if logger.TraceLogger != nil {
		newLogger.TraceLogger = logger.TraceLogger.WithFields(fields)
	}
	return newLogger
}

// TraceInfo logs a trace info line containing the message with a field name msg
func (logger *Logger) TraceInfo(msg string) {
	logger.traceInfo(msg)
}

func (logger *Logger) traceInfo(msg string) {
	withCallerInfo(logger.TraceLogger).Info(msg)
}

// TraceInfof logs a trace info line containing the formated string with a field name msg
func (logger *Logger) traceInfof(fmt string, args ...interface{}) {
	withCallerInfo(logger.TraceLogger).Infof(fmt, args...)
}

// TraceInfow logs a trace info line with some extra fields
func (logger *Logger) traceInfow(msg string, keysAndValues ...interface{}) {
	loggerWithCallerInfo := withCallerInfo(logger.TraceLogger)
	fields := logger.sweetenFields(keysAndValues)
	loggerWithCallerInfo.WithFields(fields).Infof(msg)
}

func (logger *Logger) sweetenFields(context []interface{}) map[string]interface{} {
	if len(context) == 0 {
		return nil
	}

	// Since the caller to this "sweetenFields" func already sets up span info,
	// the logging calls inside this "sweetenFields" func don't need to provide
	// a context with span info again.
	fields := map[string]interface{}{}
	for i := 0; i < len(context); {
		if context[i] == nil {
			logger.Error(nil, "one field is nil in log fields")
		}
		if v, ok := context[i].(string); ok {
			if i == len(context)-1 {
				// this means the caller input a variable, but not give us the value, we set it to an empty value.
				logger.Warning(nil, "one field is empty in log fields")
			} else {
				fields[v] = context[i+1]
			}
		}
		i += 2
	}
	return fields
}

func (logger *Logger) withSpanInfo(ctx context.Context) *Logger {
	if ctx == nil {
		return logger
	}
	span := trace.FromContext(ctx)
	if span == nil {
		return logger
	}
	return logger.WithFields(map[string]interface{}{
		"spanID":  span.SpanContext().SpanID.String(),
		"traceID": span.SpanContext().TraceID.String(),
	})
}

// TraceDebug logs a trace debug line containing the message with a field name msg
func (logger *Logger) traceDebug(msg string) {
	withCallerInfo(logger.TraceLogger).Debug(msg)
}

// TraceDebugf logs a trace debug line containing the formatted string with a field name msg
func (logger *Logger) traceDebugf(format string, args ...interface{}) {
	withCallerInfo(logger.TraceLogger).Debugf(format, args...)
}

// TraceWarning logs a trace warning line containing the message with a field name msg
func (logger *Logger) traceWarning(msg string) {
	withCallerInfo(logger.TraceLogger).Warn(msg)
}

// TraceWarningf logs a trace warning line containing the formatted string with a field name msg
func (logger *Logger) traceWarningf(fmt string, args ...interface{}) {
	withCallerInfo(logger.TraceLogger).Warnf(fmt, args...)
}

// TraceWarningw logs a trace info line with some extra fields
func (logger *Logger) traceWarningw(msg string, keysAndValues ...interface{}) {
	loggerWithCallerInfo := withCallerInfo(logger.TraceLogger)
	fields := logger.sweetenFields(keysAndValues)
	loggerWithCallerInfo.WithFields(fields).Warnf(msg)
}

// TraceError logs a trace error line containing the message with a field name msg
func (logger *Logger) traceError(msg string) {
	withCallerInfo(logger.TraceLogger).Error(msg)
}

// TraceErrorf logs a trace error line containing the formatted string with a field name msg
func (logger *Logger) traceErrorf(fmt string, args ...interface{}) {
	withCallerInfo(logger.TraceLogger).Errorf(fmt, args...)
}

// TraceErrorw logs a trace info line with some extra fields
func (logger *Logger) traceErrorw(msg string, keysAndValues ...interface{}) {
	loggerWithCallerInfo := withCallerInfo(logger.TraceLogger)
	fields := logger.sweetenFields(keysAndValues)
	loggerWithCallerInfo.WithFields(fields).Errorf(msg)
}

// TraceErrorWithStack logs a trace error along with a stacktrace if exists on the error
func (logger *Logger) traceErrorWithStack(err error) {
	withCallerInfo(logger.TraceLogger).WithField(stackTraceFieldName, getStackTraceStr(err)).Error(err.Error())
}

// TraceFatal logs a trace fatal line containing the message with a field name msg and kills process
func (logger *Logger) traceFatal(msg string) {
	withCallerInfo(logger.TraceLogger).Fatal(msg)
}

// TraceFatalf logs a trace fatal line containing the formatted string with a field name msg and kills process
func (logger *Logger) traceFatalf(fmt string, args ...interface{}) {
	withCallerInfo(logger.TraceLogger).Fatalf(fmt, args...)
}

// TraceFatalw logs a trace info line with some extra fields
func (logger *Logger) traceFatalw(msg string, keysAndValues ...interface{}) {
	loggerWithCallerInfo := withCallerInfo(logger.TraceLogger)
	fields := logger.sweetenFields(keysAndValues)
	loggerWithCallerInfo.WithFields(fields).Fatalf(msg)
}

// TracePanicw logs a trace info line with some extra fields
func (logger *Logger) tracePanicw(msg string, keysAndValues ...interface{}) {
	loggerWithCallerInfo := withCallerInfo(logger.TraceLogger)
	fields := logger.sweetenFields(keysAndValues)
	loggerWithCallerInfo.WithFields(fields).Panicf(msg)
}

// TraceInfofWithLatency logs a trace info line containing the formatted string with latency in operation fields
func (logger *Logger) traceInfofWithLatency(latency time.Duration, fmt string, args ...interface{}) {
	operationFields := map[string]interface{}{
		LatencyFieldName: latency.Nanoseconds() / NanoSecondToMillisecondConversionFactor,
	}

	withCallerInfo(logger.TraceLogger).WithFields(operationFields).Infof(fmt, args...)
}

func (logger *Logger) QOSEvent(fields map[string]interface{}, placeHolder string) {

	qosLogger := logger.QosLogger.WithFields(fields)

	// This placeHolder is useless but the tech we are using to forward the log
	// to mdsd isn't recognizing the log as json if the message is empty
	qosLogger.Info(placeHolder)
}
