package log

import (
	"context"
	"net/http"
	"runtime"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
	"go.opencensus.io/trace"
	"github.com/Azure/aks-deployer/pkg/apierror"
)

// AKSTeam is an AKS sub-team that owns certain AKS code or component. This info is
// imbedded in the trace spans as an attribute. Because we log the span ID for each log message,
// with the team info found in the span, for each error in the log we now know which sub-team
// owns it.
//
// The ownership info is obtained from the following team WiKi page:
// https://msazure.visualstudio.com/CloudNativeCompute/_wiki/wikis/CloudNativeCompute.wiki/16868/AKS-Service-Overview
type AKSTeam string

// teamContextKey is used to save team info into the context. We are doing this in addition to add
// team info to the span because turns out there's no API to get the attribute out of an existing
// span.
type teamContextKey struct{}

func NewContextWithTeam(ctx context.Context, team AKSTeam) context.Context {
	return context.WithValue(ctx, teamContextKey{}, team)
}

func GetTeamFromContext(ctx context.Context) AKSTeam {
	// Defaults to Unknown unless we succeed in getting the team info from the context
	team := AKSTeamUnknown

	teamFromContext := ctx.Value(teamContextKey{})
	if teamFromContext != nil {
		team = teamFromContext.(AKSTeam)
	}

	return team
}

const (
	aksTeamAttributeKey = "team"

	// AKSTeamSameAsParent means the team owner of this child span is the same as its parent span.
	AKSTeamSameAsParent AKSTeam = ""
	AKSTeamUnknown      AKSTeam = "Unknown"

	// See the code comment of
	//    type AKSTeam string
	// to find the link to the WiKi page that contains mapping from the teams defined here to
	// the actual email group/engineers to contact.
	AKSTeamAddon             AKSTeam = "Addon"
	AKSTeamBilling           AKSTeam = "Billing"
	AKSTeamCCPPoolController AKSTeam = "CCPPoolController"
	AKSTeamKonnectivityVPN   AKSTeam = "Konnectivity_VPN"
	AKSTeamETCD              AKSTeam = "ETCD"
	AKSTeamEventGrid         AKSTeam = "EventGrid"
	AKSTeamHCPDataAccess     AKSTeam = "HCP_DataAccess"
	AKSTeamJITController     AKSTeam = "JITController"
	AKSTeamMSIAuth           AKSTeam = "MSI_Auth"
	AKSTeamNetworking        AKSTeam = "Networking"
	AKSTeamNodeProvisioning  AKSTeam = "NodeProvisioning"
	AKSTeamWindows           AKSTeam = "Windows"
	AKSTeamOverlayManager    AKSTeam = "OverlayManager"
	AKSTeamPrivateCluster    AKSTeam = "PrivateCluster"
	AKSTeamOperationQueue    AKSTeam = "OperationQueue"
	AKSTeamRegionalLooper    AKSTeam = "RegionalLooper"
	AKSTeamRP                AKSTeam = "RP"
	AKSTeamUpgrade           AKSTeam = "Upgrade"
)

// ProfileSpan creates a new span for provided function, sets errors on exit
func ProfileSpan(ctx context.Context, spanName string, team AKSTeam, executeFunc func(ctx context.Context) *apierror.Error) *apierror.Error {
	spanCtx, span := StartSpan(ctx, spanName, team)
	defer span.End()

	putErr := executeFunc(spanCtx)

	if putErr != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: putErr.Error()})
	}

	return putErr
}

// StartSpan wrapper when there are no tags
func StartSpan(ctx context.Context, name string, team AKSTeam) (context.Context, *trace.Span) {
	return startSpanWithTagsInternal(ctx, name, team, nil)
}

// getCommonTags returns common tags which needs to be appended to traces
func getCommonTags(ctx context.Context) map[string]string {
	tags := make(map[string]string)
	if apitracking, ok := GetAPITracking(ctx); ok {
		tags["subscription"] = apitracking.GetSubscriptionID().String()
		tags["resourceGroupName"] = apitracking.GetResourceGroupName()
		tags["resourceName"] = apitracking.GetResourceName()
		tags["operationID"] = apitracking.GetOperationID().String()
		tags["operationName"] = apitracking.GetOperationName()
		tags["subOperationName"] = apitracking.GetSubOperationName()
	}
	return tags
}

// StartSpanWithTags wrapper around trace.StartSpan to add tags as attributes and also add additional apitracking fields
func StartSpanWithTags(ctx context.Context, name string, team AKSTeam, tags map[string]string) (context.Context, *trace.Span) {
	return startSpanWithTagsInternal(ctx, name, team, tags)
}

// StackAnnotationKey allows the kusto exporter to identify the stack annotation and extract it into columns
const StackAnnotationKey string = "stackinfo"

// startSpanWithTagsInternal ensures correct call stack level when calling for logging purposes
func startSpanWithTagsInternal(ctx context.Context, name string, team AKSTeam, tags map[string]string) (context.Context, *trace.Span) {

	// Don't assume that the stack frame we want is at a fixed offset
	// Instead, walk up the stack to find the first frame that's outside this file
	_, traceFile, _, _ := runtime.Caller(0)
	frame := 2 // By default skip any public methods calling this internal one
	var file string
	var line int
	for {
		_, file, line, _ = runtime.Caller(frame)
		if file != traceFile {
			break
		}

		frame++
	}

	var attributes []trace.Attribute

	for key, value := range getCommonTags(ctx) {
		attributes = append(attributes, trace.StringAttribute(key, value))
	}

	for key, value := range tags {
		attributes = append(attributes, trace.StringAttribute(key, value))
	}

	if team == AKSTeamSameAsParent {
		team = GetTeamFromContext(ctx)
	}

	attributes = append(attributes, trace.StringAttribute(aksTeamAttributeKey, string(team)))
	ctx = NewContextWithTeam(ctx, team)

	ctx, span := trace.StartSpan(ctx, name)
	span.Annotate(
		[]trace.Attribute{
			trace.StringAttribute(fileNameFieldName, file),
			trace.Int64Attribute(lineNumberFieldName, int64(line)),
		}, StackAnnotationKey)

	span.AddAttributes(attributes...)

	return ctx, span
}

// WithTags adds additional tags as attributes to the trace.Span
func WithTags(span *trace.Span, tags map[string]string) {
	var attributes []trace.Attribute

	for key, value := range tags {
		attributes = append(attributes, trace.StringAttribute(key, value))
	}

	span.AddAttributes(attributes...)
}

// SpanContextFromRequest extracts a span context from a given request [used for testing, use Propagator middleware]
func SpanContextFromRequest(req *http.Request) (trace.SpanContext, bool) {
	format := tracecontext.HTTPFormat{}
	return format.SpanContextFromRequest(req)
}

// SpanContextToRequest is a wrapper to serialize span information from a given context into an http request
func SpanContextToRequest(ctx context.Context, req *http.Request) {
	format := tracecontext.HTTPFormat{}
	format.SpanContextToRequest(trace.FromContext(ctx).SpanContext(), req)
}

// PropagatorMiddleware will deserialize span information from a given http request and set it in the http context
func PropagatorMiddleware(inner http.Handler) http.Handler {

	return &ochttp.Handler{
		Handler:     inner,
		Propagation: &tracecontext.HTTPFormat{},
	}

}
