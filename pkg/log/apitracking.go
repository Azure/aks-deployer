// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This file is for code about storing and retrieving api tracking
// info from a context struct

package log

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"

	"github.com/Azure/aks-deployer/pkg/consts"
)

type Tracking interface {
}

type trackingKey string

const apiTrackingKey trackingKey = "AcsApiTracking"

const (
	// AcceptLanguageHeader is the standard http header name used so that we don't have to pass in the http request
	AcceptLanguageHeader = "Accept-Language"
	// HostHeader is the standard http header Host used to indicate the target host name
	HostHeader = "Host"
	// ContentTypeHeader is the standard http header Content-Type
	ContentTypeHeader = "Content-Type"
	// RequestAPIVersionParameterName is the query string parameter name ARM adds for the api version
	RequestAPIVersionParameterName = "api-version"
	// RequestResourceTypeParameterName is the query string parameter name to distinguish between cluster types
	RequestResourceTypeParameterName = "resource-type"
	// RequestCorrelationIDHeader is the http header name ARM adds for the correlationID
	RequestCorrelationIDHeader = "x-ms-correlation-request-id"
	// RequestARMClientRequestIDHeader is the http header name ARM adds for the client request id
	RequestARMClientRequestIDHeader = "x-ms-client-request-id"
	// RequestGraphClientRequestIDHeader is the http header name graph adds for the client request id
	RequestGraphClientRequestIDHeader = "client-request-id"
	// RequestClientSessionIDHeader is the http header name ARM adds for the client session id
	RequestClientSessionIDHeader = "x-ms-client-session-id"
	// RequestClientApplicationIDHeader is the http header name ARM adds for the client app id
	RequestClientApplicationIDHeader = "x-ms-client-app-id"
	// RequestClientPrincipalNameHeader is the http header name ARM adds for the client principal name
	RequestClientPrincipalNameHeader = "x-ms-client-principal-name"
	// ResponseARMRequestIDHeader is the http header name our rp adds to uniquely identify this request, used by ARM
	ResponseARMRequestIDHeader = "x-ms-request-id"
	// ResponseGraphRequestIDHeader is the http header name our rp adds to uniquely identify this request, used by graph
	ResponseGraphRequestIDHeader = "request-id"
	// CredentialFormatParameterName is the query string parameter name, optional
	CredentialFormatParameterName = "format"
	// LoginMethodParameterName is the query string parameter name, optional
	LoginMethodParameterName = "login"
	// CredentialServerFQDNFormatParameterName is the query string parameter name, optional
	CredentialServerFQDNFormatParameterName = "server-fqdn"
	// AADTenantIDHeader is the header used by ARM to pass the requesting entity's AAD tenant ID to RP.
	AADTenantIDHeader = "x-ms-client-tenant-id"
	// LocationHeader is the http header name required for Accepted status
	LocationHeader = "Location"
	// AzureAsyncOperationHeader is the http header name required for async operation
	AzureAsyncOperationHeader = "Azure-AsyncOperation"
	// RetryAfterHeader is the http header name for client's back off duration
	RetryAfterHeader = "Retry-After"
	// UserAgentHeader is the http header name for the user agent
	UserAgentHeader = "User-Agent"
	// ForwardedForHeader is the http header name for identifying the original client IP address
	ForwardedForHeader = "X-Forwarded-For"
	// RealIPHeader is the http header name for identifying the client IP address
	RealIPHeader = "X-Real-Ip"
	// RetryAttemptHeader is the grpc header name for retry attempt count
	RetryAttemptHeader = "x-retry-attempty"
)

const (
	// RequestAcsOperationIDHeader is the http header name ACS RP adds for it's operation ID
	RequestAcsOperationIDHeader = "x-ms-acs-operation-id"
	//RequestClientCertThumbprintHeader contains the sha-1 thumbprint hash for the client cert, added by the nginx proxy in production clusters
	RequestClientCertThumbprintHeader = "x-ssl-cert-thumbprint"
)

const (
	// CorrelationID is the id that correlation an RP operation with ARM operation/request
	CorrelationID = "correlationID"
	// OperationID is the operation id to track a request received by ACS
	OperationID = "operationID"
	// HCPOperationID is the operation id of operation being processed by HCP service
	HCPOperationID = "hcpOperationID"
	// HCPControlPlaneID is the ID of the control-plane as set by HCP
	HCPControlPlaneID = "hcpControlPlaneID"
	// OperationName is the name of operation invoked by a request received by ACS
	OperationName = "operationName"
	// SubOperationName is the name of suboperation invoked by a request received by ACS
	SubOperationName = "subOperationName"
	// AADTenantID is the ID of the customer's AAD tenant
	AADTenantID = "aadTenantID"
	// SubscriptionID identifies a customer/an organization
	SubscriptionID = "subscriptionID"
	// ResourceGroupName contains a group of ARM (RP) resources
	ResourceGroupName = "resourceGroupName"
	// ResourceName is the name of ACS resource
	ResourceName = "resourceName"
	// DetectorName is the name of app lens detector
	DetectorName = "detectorName"
	// AgentPoolName is the name of the agentpool in the managed cluster.
	AgentPoolName = "agentPoolName"
	// MaintenanceConfigurationName is the name of the maintenanceConfiguration in the managed cluster.
	MaintenanceConfigurationName = "maintenanceConfiguration"
	// APIVersion is api version on which the request was sent to ACS
	APIVersion = "apiVersion"
	// AcceptLanguage is the language of request
	AcceptLanguage = "acceptLanguage"
	// ClientApplicationID is the value that came in http header for the client app id
	ClientApplicationID = "clientApplicationID"
	// ClientPrincipalName is the client principal name
	ClientPrincipalName = "clientPrincipalName"
	// ClientRequestID is the client request id
	ClientRequestID = "clientRequestID"
	// UserAgent is user agent of the request
	UserAgent = "userAgent"
	// MessageID is the messageID of a queue message
	MessageID = "messageID"
	// PopReceipt is the pop receipt of a queue message
	PopReceipt = "popReceipt"
	// InsertionTime is the insertion time of a queue message
	InsertionTime = "insertionTime"
	// ExpirationTime is the expiration time of a queue message
	ExpirationTime = "expirationTime"
	// DequeueCount is the dequeue count of a queue message
	DequeueCount = "dequeueCount"
	// TimeNextVisible is the time a queue message will become visible
	TimeNextVisible = "timeNextVisible"
	// OperationTimeout is the operation execution timeout
	OperationTimeout = "operationTimeout"
	// DelayStartInSeconds indicate worker not to pick up a message for this amount of time initially.
	DelayStartInSeconds = "delayStartInSeconds"
	// RegionName is the region of the service serving traffic
	RegionName = "region"
	// PropertiesBag is the field to put the properties
	PropertiesBag = "propertiesBag"
	//k8sCurrentVersion is the current k8s version
	k8sCurrentVersion = "k8sCurrentVersion"
	//k8sGoalVersion is the goal k8s version
	k8sGoalVersion = "k8sGoalVersion"
	// HCPUnderlayName is the name of the selected underlay for the hosted control-plane pods
	HCPUnderlayName = "hcpUnderlayName"
	//ControlPlaneAZEnabled is the propertyname which we set in propertybag for the existence az for the control plane
	ControlPlaneAZEnabled = "enableControlPlaneAZ"
	//AgentPoolAZEnabled is part of the propertyname which we set in propertybag for the existence az for the agentpool
	AgentPoolAZEnabled = "enableAZ"
	// InternalSubscription is the parameters map key used to indicate if the corresponding subscription ID in the
	// map is an internal subscription ID.
	InternalSubscription = "internalSubscription"
	// agentPoolCount provides a hint to the async side to how many agentPools are expected to be in the Database.
	agentPoolCount = "agentpoolcount"

	// ExtensionAddonName is the name of the extension addon in the managed cluster.
	ExtensionAddonName = "extensionAddonName"

	spanContext = "spanContext"

	agentPoolLabelsToDelete = "agentPoolLabelsToDelete"
)

const (
	// DefaultSubOperationName is default suboperation name
	DefaultSubOperationName = "None"
	// LowerTrueString is the bool value = true as string used at multiple places
	LowerTrueString = "true"
)

const (
	//DefaultMessageTTL : set default message TTL for 3 days
	DefaultMessageTTL = time.Duration(72) * time.Hour

	//DefaultOperationTimeout : default timeout for completing an async operation (PUT/DELETE). The RP will
	// timeout the operation if long running operation does not complete in 120 minutes
	DefaultOperationTimeout = time.Duration(120) * time.Minute
)

// APITracking is a type for storing all the context info needed to track an api call
type APITracking struct {
	correlationID                uuid.UUID
	operationID                  uuid.UUID
	hcpOperationID               string
	hcpUnderlayName              string
	hcpControlPlaneID            string
	operationName                string
	subOperationName             string
	aadTenantID                  string
	subscriptionID               uuid.UUID
	resourceGroupName            string
	resourceName                 string
	detectorName                 string
	agentPoolName                string
	maintenanceConfigurationName string
	extensionAddonName           string
	resourceType                 string
	apiVersion                   string
	acceptLanguage               string
	host                         string
	referer                      string
	clientApplicationID          string
	clientPrincipalName          string
	clientRequestID              uuid.UUID
	clientSessionID              string
	userAgent                    string
	accessProfileRoleName        string
	podName                      string
	containerName                string
	virtualNetworkName           string
	subnetName                   string
	vmName                       string
	actionName                   string
	region                       string
	resultCodeDependency         []string
	propertiesBag                map[string]string
	auxiliaryToken               string
	httpMethod                   string
	targetURI                    string
	fqdn                         string
	cxUnderlayName               string
	controlPlaneID               string

	// error categorization related fields
	errorSubcode    string
	errorCategory   string
	errorDependency string
	errorAKSTeam    string

	// to receive linked notification on diagnosticSettings
	extensionProviderName string
	extensionResourceType string
	extensionResourceName string

	// isInternalSubscription indicates if the subscription ID is an internal subscription.
	isInternalSubscription bool

	// async properties
	messageID      int
	popReceipt     string
	insertionTime  time.Time
	expirationTime time.Time

	dequeueCount int

	// operationTimeout, indicate worker manager to kill a worker(operation) if it exceed this time.
	// we usually keep extend this (if the process is still responsive)
	operationTimeout time.Duration

	// timeNextVisible, tells us at what time next worker going to pick up this message.
	timeNextVisible time.Time

	// delayStartInSeconds, indicate worker not to pick up this message for certain time initially.
	// but after first pickup, timeNextVisible should be relied on.
	delayStartInSeconds uint

	// messageTTL, indicate at what time to drop a message.
	// no matter how many time it retried, or how long it has been running.
	messageTTL time.Duration

	// private link
	privateEndpointConnectionName string

	// list credential
	credentialFormat           string
	loginMethod                string
	credentialServerFQDNFormat string
}

// NewAPITracking creates a new APITracking struct with data from the http request and returns a pointer to it
func NewAPITracking(pathParameters map[string]string, request *http.Request, routePath string) *APITracking {
	return NewAPITrackingWithRegion(pathParameters, request, routePath, "")
}

// NewAPITrackingWithRegion creates a new APITracking struct with data from the http request and returns a pointer to it
func NewAPITrackingWithRegion(pathParameters map[string]string, request *http.Request, routePath, region string) *APITracking {
	query := request.URL.Query()
	method := request.Method
	headers := request.Header

	retVal := &APITracking{}

	retVal.correlationID = uuid.FromStringOrNil(headers.Get(RequestCorrelationIDHeader))

	acsOpIDInHcpRequest := request.Header.Get(http.CanonicalHeaderKey(RequestAcsOperationIDHeader))
	if acsOpIDInHcpRequest != "" {
		retVal.operationID = uuid.FromStringOrNil(acsOpIDInHcpRequest)
	} else {
		retVal.operationID = uuid.NewV4()
	}

	retVal.setOperationName(method, routePath)
	retVal.SetSubOperationName(DefaultSubOperationName)

	if subIDString, ok := pathParameters[consts.PathSubscriptionIDParameter]; ok {
		retVal.subscriptionID, _ = uuid.FromString(subIDString) // nolint
	}

	if rGroupName, ok := pathParameters[consts.PathResourceGroupNameParameter]; ok {
		retVal.resourceGroupName = rGroupName
	}
	if rName, ok := pathParameters[consts.PathResourceNameParameter]; ok {
		retVal.resourceName = rName
	}
	if value, ok := pathParameters[consts.PathDetectorNameParameter]; ok {
		retVal.detectorName = value
	}
	if apName, ok := pathParameters[consts.PathAgentPoolNameParameter]; ok {
		retVal.agentPoolName = apName
	}
	if mtcName, ok := pathParameters[consts.PathMaintenanceConfigurationNameParameter]; ok {
		retVal.maintenanceConfigurationName = mtcName
	}
	if extensionAddonName, ok := pathParameters[consts.PathExtensionAddonNameParameter]; ok {
		retVal.extensionAddonName = extensionAddonName
	}
	if rAccessProfileRoleName, ok := pathParameters[consts.PathAccessProfileParameter]; ok {
		retVal.accessProfileRoleName = rAccessProfileRoleName
	}
	if pName, ok := pathParameters[consts.PathPodNameParameter]; ok {
		retVal.podName = pName
	}
	if cName, ok := pathParameters[consts.PathContainerNameParameter]; ok {
		retVal.containerName = cName
	}
	if cName, ok := pathParameters[consts.PathVirtualNetworkNameParameter]; ok {
		retVal.virtualNetworkName = cName
	}
	if cName, ok := pathParameters[consts.PathSubnetNameParameter]; ok {
		retVal.subnetName = cName
	}
	if vmName, ok := pathParameters[consts.PathVMNameParameter]; ok {
		retVal.vmName = vmName
	}
	if aName, ok := pathParameters[consts.PathActionNameParameter]; ok {
		retVal.actionName = aName
	}

	if extProviderName, ok := pathParameters[consts.PathExtensionProviderParameter]; ok {
		retVal.extensionProviderName = extProviderName
	}

	if extResourceType, ok := pathParameters[consts.PathExtensionResourceTypeParameter]; ok {
		retVal.extensionResourceType = extResourceType
	}

	if extResourceName, ok := pathParameters[consts.PathExtensionResourceNameParameter]; ok {
		retVal.extensionResourceName = extResourceName
	}

	if pecName, ok := pathParameters[consts.PathPrivateEndpointConnectionNameParameter]; ok {
		retVal.privateEndpointConnectionName = pecName
	}

	if conPlID, ok := pathParameters[consts.PathControlPlaneParameter]; ok {
		retVal.controlPlaneID = conPlID
	}

	retVal.httpMethod = method
	retVal.targetURI = request.URL.String()
	retVal.apiVersion = query.Get(RequestAPIVersionParameterName)
	retVal.resourceType = query.Get(RequestResourceTypeParameterName)
	retVal.acceptLanguage = headers.Get(AcceptLanguageHeader)
	retVal.clientApplicationID = headers.Get(RequestClientApplicationIDHeader)
	retVal.aadTenantID = headers.Get(AADTenantIDHeader)
	cpn := headers.Get(RequestClientPrincipalNameHeader)
	if retVal.IsAdminOperation() {
		retVal.clientPrincipalName = cpn
	} else if strings.Contains(cpn, "@") {
		retVal.clientPrincipalName = cpn[strings.Index(cpn, "@"):]
	}

	clientRequestID := headers.Get(RequestARMClientRequestIDHeader)
	if clientRequestID == "" {
		clientRequestID = headers.Get(RequestGraphClientRequestIDHeader)
	}
	retVal.clientRequestID = uuid.FromStringOrNil(clientRequestID)

	retVal.clientSessionID = headers.Get(RequestClientSessionIDHeader)
	retVal.userAgent = request.UserAgent()
	retVal.region = region
	retVal.referer = request.Referer()
	// Note. For incoming request, Host is promoted to request.Host, and removed from header map
	retVal.host = request.Host
	if retVal.host == "" && request.URL != nil {
		retVal.host = request.URL.Host
	}
	if retVal.host == "" {
		// For tests, Host header is not promoted
		retVal.host = headers.Get(HostHeader)
	}

	retVal.messageTTL = DefaultMessageTTL
	retVal.operationTimeout = DefaultOperationTimeout

	// init properties bag
	retVal.propertiesBag = make(map[string]string)

	// list credential
	retVal.credentialFormat = query.Get(CredentialFormatParameterName)
	retVal.loginMethod = query.Get(LoginMethodParameterName)
	retVal.credentialServerFQDNFormat = query.Get(CredentialServerFQDNFormatParameterName)

	return retVal
}

// NewAPITrackingFromOutgoingRequest creates a new APITracking struct with data from the out going request (extracting fields out of context)
// and returns a pointer to it
func NewAPITrackingFromOutgoingRequest(pathParameters map[string]string, request *http.Request, routePath, region string) (*APITracking, error) {
	retVal := NewAPITrackingWithRegion(pathParameters, request, routePath, region)

	ctx := request.Context()
	apiTrackingInCtx, hasApiTracking := GetAPITracking(ctx)
	if hasApiTracking {
		retVal.operationID = apiTrackingInCtx.operationID
		retVal.operationName = apiTrackingInCtx.operationName
		retVal.subOperationName = apiTrackingInCtx.subOperationName
		retVal.subscriptionID = apiTrackingInCtx.subscriptionID
		retVal.resourceGroupName = apiTrackingInCtx.resourceGroupName
		retVal.resourceName = apiTrackingInCtx.resourceName
		retVal.clientSessionID = apiTrackingInCtx.clientSessionID
		retVal.correlationID = apiTrackingInCtx.correlationID
	}

	return retVal, nil
}

// NewHealthAPITracking creates a new APITracking struct for a health check call
func NewHealthAPITracking(region string) *APITracking {
	apiTracking := &APITracking{}
	apiTracking.operationID = uuid.NewV4()
	apiTracking.operationName = consts.HealthCheckOperationName
	apiTracking.subOperationName = DefaultSubOperationName
	apiTracking.region = region
	return apiTracking
}

// NewAPITrackingFromParametersMap creates a new APITracking struct with data pass in a map
// The intended use of this function is mainly for creating API tracking for queuemessageprocessor
// and unit testing.
func NewAPITrackingFromParametersMap(m map[string]interface{}) *APITracking {
	retVal := &APITracking{}

	if operationID, ok := m[OperationID].(uuid.UUID); ok {
		retVal.operationID = operationID
	}

	if correlationID, ok := m[CorrelationID].(uuid.UUID); ok {
		retVal.correlationID = correlationID
	}

	if operationName, ok := m[OperationName].(string); ok {
		retVal.operationName = operationName
	}

	if subOperationName, ok := m[SubOperationName].(string); ok {
		retVal.subOperationName = subOperationName
	}

	if aadTenantID, ok := m[AADTenantID].(string); ok {
		retVal.aadTenantID = aadTenantID
	}

	if subscriptionID, ok := m[SubscriptionID].(uuid.UUID); ok {
		retVal.subscriptionID = subscriptionID
	}

	if v, ok := m[InternalSubscription]; ok {
		if vBool, ok := v.(bool); ok {
			retVal.isInternalSubscription = vBool
		}
	}

	if resourceGroupName, ok := m[ResourceGroupName].(string); ok {
		retVal.resourceGroupName = resourceGroupName
	}

	if resourceName, ok := m[ResourceName].(string); ok {
		retVal.resourceName = resourceName
	}

	if clientApplicationID, ok := m[ClientApplicationID].(string); ok {
		retVal.clientApplicationID = clientApplicationID
	}

	if name, ok := m[DetectorName].(string); ok {
		retVal.detectorName = name
	}

	if apName, ok := m[AgentPoolName].(string); ok {
		retVal.agentPoolName = apName
	}

	if mtcName, ok := m[MaintenanceConfigurationName].(string); ok {
		retVal.maintenanceConfigurationName = mtcName
	}

	if extensionAddonName, ok := m[ExtensionAddonName].(string); ok {
		retVal.extensionAddonName = extensionAddonName
	}

	if messageID, ok := m[MessageID].(int); ok {
		retVal.messageID = messageID
	}

	if popReceipt, ok := m[PopReceipt].(string); ok {
		retVal.popReceipt = popReceipt
	}

	if dequeueCount, ok := m[DequeueCount].(int); ok {
		retVal.dequeueCount = dequeueCount
	}

	if insertionTime, ok := m[InsertionTime].(time.Time); ok {
		retVal.insertionTime = insertionTime
	}

	if expirationTime, ok := m[ExpirationTime].(time.Time); ok {
		retVal.expirationTime = expirationTime
	}

	if timeNextVisible, ok := m[TimeNextVisible].(time.Time); ok {
		retVal.timeNextVisible = timeNextVisible
	}

	if delayStartInSeconds, ok := m[DelayStartInSeconds].(uint); ok {
		retVal.delayStartInSeconds = delayStartInSeconds
	}

	if operationTimeout, ok := m[OperationTimeout].(time.Duration); ok {
		retVal.operationTimeout = operationTimeout
	}

	if language, ok := m[AcceptLanguage].(string); ok {
		retVal.acceptLanguage = language
	}

	if hcpOperationID, ok := m[HCPOperationID].(string); ok {
		retVal.hcpOperationID = hcpOperationID
	}

	if hcpControlPlaneID, ok := m[HCPControlPlaneID].(string); ok {
		retVal.hcpControlPlaneID = hcpControlPlaneID
	}

	if hcpUnderlayName, ok := m[HCPUnderlayName].(string); ok {
		retVal.hcpUnderlayName = hcpUnderlayName
	}

	if region, ok := m[RegionName].(string); ok {
		retVal.region = region
	}

	if propertiesBag, ok := m[PropertiesBag].(map[string]string); ok {
		retVal.propertiesBag = propertiesBag
	} else {
		retVal.propertiesBag = map[string]string{}
	}

	if extProviderName, ok := m[consts.PathExtensionProviderParameter].(string); ok {
		retVal.extensionProviderName = extProviderName
	}

	if extResourceType, ok := m[consts.PathExtensionResourceTypeParameter].(string); ok {
		retVal.extensionResourceType = extResourceType
	}

	if extResourceName, ok := m[consts.PathExtensionResourceNameParameter].(string); ok {
		retVal.extensionResourceName = extResourceName
	}

	if privateEndpointConnectionName, ok := m[consts.PathPrivateEndpointConnectionNameParameter].(string); ok {
		retVal.privateEndpointConnectionName = privateEndpointConnectionName
	}

	if apiVersion, ok := m[APIVersion].(string); ok {
		retVal.apiVersion = apiVersion
	}

	if credentialFormat, ok := m[CredentialFormatParameterName].(string); ok {
		retVal.credentialFormat = credentialFormat
	}

	if loginMethod, ok := m[LoginMethodParameterName].(string); ok {
		retVal.loginMethod = loginMethod
	}

	if credentialServerFQDNFormat, ok := m[CredentialServerFQDNFormatParameterName].(string); ok {
		retVal.credentialServerFQDNFormat = credentialServerFQDNFormat
	}

	// set default message TTL for 72 hour
	retVal.messageTTL = DefaultMessageTTL

	return retVal
}

// GetCorrelationID retrieves the correlation id from the APITracking struct
func (at *APITracking) GetCorrelationID() uuid.UUID {
	return at.correlationID
}

// GetOperationID retrieves the operation id from the APITracking struct
func (at *APITracking) GetOperationID() uuid.UUID {
	return at.operationID
}

// GetHCPOperationID retrieves the operation id from the APITracking struct
func (at *APITracking) GetHCPOperationID() string {
	return at.hcpOperationID
}

// GetRegion returns a request's target region
func (at *APITracking) GetRegion() string {
	return at.region
}

// SetRegion sets the region of target resource
func (at *APITracking) SetRegion(region string) {
	at.region = region
}

// setOperationName sets the operation name in the APITracking based on the method and path passed in
func (at *APITracking) setOperationName(method, routePath string) {
	at.operationName = "Unknown"
	switch routePath {
	case consts.SubscriptionResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetSubscriptionOperationName
		case http.MethodPut:
			at.operationName = consts.PutSubscriptionOperationName
		}
	case consts.ContainerServiceResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetContainerServiceOperationName
		case http.MethodPut:
			at.operationName = consts.PutContainerServiceOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteContainerServiceOperationName
		}
	case consts.OperationResultsResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetOperationResultsOperationName
		}
	case consts.OperationStatusResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetOperationStatusOperationName
		}
	case consts.ListContainerServiceResourcesBySubscriptionFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListContainerServicesBySubscriptionOperationName
		}
	case consts.ListContainerServiceResourcesByResourceGroupFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListContainerServicesByResourceGroupOperationName
		}
	case consts.DeploymentPreflightFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostDeploymentPreflightOperationName
		}
	case consts.InternalContainerServiceResourceFullPath:
		switch method {
		case http.MethodPut:
			at.operationName = consts.InternalPutContainerServiceOperationName
		}
	case consts.InternalSubscriptionResourceFullPath:
		switch method {
		case http.MethodPut:
			at.operationName = consts.InternalPutSubscriptionOperationName
		}
	case consts.InternalOperationStatusResourceFullPath:
		switch method {
		case http.MethodPut:
			at.operationName = consts.InternalPutOperationStatusOperationName
		}

	case consts.AgentPoolResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetAgentPoolOperationName
		case http.MethodPut:
			at.operationName = consts.PutAgentPoolOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteAgentPoolOperationName
		}

	case consts.MaintenanceConfigurationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetMaintenanceConfigurationOperationName
		case http.MethodPut:
			at.operationName = consts.PutMaintenanceConfigurationOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteMaintenanceConfigurationOperationName
		}

	case consts.ExtensionAddonFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetExtensionAddonOperationName
		case http.MethodPut:
			at.operationName = consts.PutExtensionAddonOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteExtensionAddonOperationName
		}

	case consts.ListExtensionAddonsByManagedClusterFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListExtensionAddonsByManagedClusterOperationName
		}

	case consts.UpgradeAgentPoolNodeImageFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.UpgradeNodeImageAgentPoolOperationName
		}

	case consts.GetAgentPoolUpgradeProfileFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetAgentPoolUpgradeProfileOperationName
		}

	case consts.ListAgentPoolAvailableVersionsFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListAgentPoolAvailableVersionsOperationName
		}

	case consts.ListAgentPoolsByClusterFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListAgentPoolsByClusterOperationName
		}

	case consts.ListMaintenanceConfigurationsByManagedClusterFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListMaintenanceConfigurationsByManagedClusterOperationName
		}

	case consts.GetAgentPoolResourcesAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetAgentPoolAdminOperationName
		case http.MethodPut:
			at.operationName = consts.PutAgentPoolAdminOperationName
		}
	case consts.ListAgentPoolsByClusterAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListAgentPoolsByClusterAdminOperationName
		}
	case consts.ManagedClusterResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetManagedClusterOperationName
		case http.MethodPut:
			at.operationName = consts.PutManagedClusterOperationName
		case http.MethodPatch:
			at.operationName = consts.PatchManagedClusterOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteManagedClusterOperationName
		}
	case consts.BackfillManagedClusterOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.BackfillManagedClusterOperationName
		}
	case consts.ReimageManagedClusterAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ReimageManagedClusterAdminOperationName
		}
	case consts.DelegateSubnetAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.DelegateSubnetAdminOperationName
		}
	case consts.UnDelegateSubnetAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.UnDelegateSubnetAdminOperationName
		}
	case consts.JsonPatchAgentPoolAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.JsonPatchAgentPoolAdminOperationName
		}
	case consts.JsonPatchControlPlaneAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.JsonPatchControlPlaneAdminOperationName
		}
	case consts.JsonPatchManagedClusterAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.JsonPatchManagedClusterAdminOperationName
		}
	case consts.ListManagedClusterResourcesBySubscriptionFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListManagedClustersBySubscriptionOperationName
		}
	case consts.ListManagedClusterResourcesByResourceGroupFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListManagedClustersByResourceGroupOperationName
		}
	case consts.ListOrchestratorsBySubscriptionFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListOrchestratorsOperationName
		}
	case consts.GetOSOptionsBySubscriptionFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetOSOptionsOperationName
		}
	case consts.GetContainerServiceUpgradeProfileFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetContainerServiceUpgradeProfileOperationName
		}
	case consts.GetManagedClusterUpgradeProfileFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetManagedClusterUpgradeProfileOperationName
		}
	case consts.GetManagedClusterDiagnosticsStateFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetManagedClusterDiagnosticsStateOperationName
		}
	case consts.GetManagedClusterAccessProfileFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetManagedClusterAccessProfileOperationName
		}
	case consts.OperationStatusResourceAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetOperationStatusAdminOperationName
		}
	case consts.ListManagedClusterResourcesBySubscriptionAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListManagedClustersBySubscriptionAdminOperationName
		}
	case consts.ListManagedClusterResourcesByResourceGroupAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListManagedClustersByResourceGroupAdminOperationName
		}
	case consts.GetManagedClusterResourceAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetManagedClusterAdminOperationName
		case http.MethodPut:
			at.operationName = consts.PutManagedClusterAdminOperationName
		}
	case consts.ListCustomerControlPlanePodsOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListCustomerControlPlanePodsOperationName
		}
	case consts.GetCustomerControlPlanePodLogOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetCustomerControlPlanePodLogOperationName
		}
	case consts.GetCustomerControlPlaneEventsOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetCustomerControlPlaneEventsOperationName
		}
	case consts.PostKubectlCommandOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostKubectlCommandOperationName
		}
	case consts.PostOverlayKubectlCommandOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostOverlayKubectlCommandOperationName
		}
	case consts.ListUnderlaysOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListUnderlaysOperationName
		}
	case consts.PostUnderlayKubectlCommandOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostUnderlayKubectlCommandOperationName
		}
	case consts.ListManagedClusterCredentialFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ListManagedClusterCredentialOperationName
		}
	case consts.GetAvailableOperationsFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetAvailableOperationsOperationName
		}
	case consts.GetSubscriptionResourceOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetSubscriptionResourceOperationName
		}
	case consts.PostVirtualMachineRunCommandOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostVirtualMachineRunCommandOperationName
		}
	case consts.PostVirtualMachineGenericRunCommandOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.PostVirtualMachineGenericRunCommandOperationName
		}
	case consts.GetCustomerAgentNodesStatusOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetCustomerAgentNodesStatusOperationName
		}
	case consts.ListManagedClusterClusterAdminCredentialFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ListManagedClusterClusterAdminCredentialOperationName
		}
	case consts.ListManagedClusterClusterUserCredentialFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ListManagedClusterClusterUserCredentialOperationName
		}
	case consts.ListManagedClusterClusterMonitoringUserCredentialFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ListManagedClusterClusterMonitoringUserCredentialOperationName
		}
	case consts.ResetAADProfileFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ResetAADProfileOperationName
		}
	case consts.ResetServicePrincipalProfileFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ResetServicePrincipalProfileOperationName
		}
	case consts.AdminLinkedNotificationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.LinkedNotificationAdminOperationName
		}
	case consts.LinkedNotificationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.LinkedNotificationOperationName
		}
	case consts.GetDetectorFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetDetectorOperationName
		}
	case consts.ListDetectorsFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListDetectorsOperationName
		}
	case consts.AdminListUnderlayFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.AdminListUnderlaysOperationName
		}
	case consts.AdminExpandUnderlayCapacityFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.AdminExpandUnderlayCapacityOperationName
		}
	case consts.AdminGetUnderlayFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.AdminGetUnderlayOperationName
		}
	case consts.AdminUnderlayFullPath:
		switch method {
		case http.MethodPut:
			at.operationName = consts.AdminPutUnderlayOperationName
		case http.MethodDelete:
			at.operationName = consts.AdminDeleteUnderlayOperationName
		}
	case consts.AdminPostUnderlayActionFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.AdminPostUnderlayActionOperationName
		}
	case consts.MigrateCustomerControlPlaneAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.MigrateCustomerControlPlaneOperationName
		}
	case consts.DeallocateControlPlaneAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.DeallocateControlPlaneOperationName
		}
	case consts.RotateClusterCertificatesFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.RotateClusterCertificatesOperationName
		}
	case consts.DrainCustomerControlPlanesAdminOperationFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.DrainCustomerControlPlanesOperationName
		}
	case consts.ListPrivateLinkServiceConnectionsFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListPrivateEndpointConnectionsOperationName
		}
	case consts.PrivateLinkServiceConnectionFullPath:
		switch method {
		case http.MethodPut:
			at.operationName = consts.PutPrivateEndpointConnectionOperationName
		case http.MethodGet:
			at.operationName = consts.GetPrivateEndpointConnectionOperationName
		case http.MethodDelete:
			at.operationName = consts.DeletePrivateEndpointConnectionOperationName
		}

	case consts.ListPrivateLinkResourcesFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListPrivateLinkResourcesOperationName
		}

	case consts.ResolvePrivateLinkServiceIDFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ResolvePrivateLinkServiceIDOperationName
		}

	case consts.StopManagedClusterFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.StopManagedClusterOperationName
		}

	case consts.StartManagedClusterFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.StartMangedClusterOperationName
		}

	case consts.ManagedClusterRunCommandFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.ManagedClusterRunCommandOperationName
		}

	case consts.ManagedClusterRunCommandResultFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ManagedClusterRunCommandResultOperationName
		}

	case consts.ServiceOutboundIPRangesFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetServiceOutboundIPRangesOperationName
		case http.MethodPut:
			at.operationName = consts.PutServiceOutboundIPRangesOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteServiceOutboundIPRangesOperationName
		}
	case consts.ListServiceOutboundIPRangesFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListServiceOutboundIPRangesOperationName
		}
	case consts.ListOutboundNetworkDependenciesEndpointsFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListOutboundNetworkDependenciesEndpointsOperationName
		}
	case consts.SnapshotResourceFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetSnapshotOperationName
		case http.MethodPut:
			at.operationName = consts.PutSnapshotOperationName
		case http.MethodDelete:
			at.operationName = consts.DeleteSnapshotOperationName
		}
	case consts.ListSnapshotResourcesBySubscriptionFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListSnapshotsBySubscriptionOperationName
		}
	case consts.ListSnapshotResourcesByResourceGroupFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.ListSnapshotsByResourceGroupOperationName
		}
	case consts.GetSnapshotResourceAdminOperationFullPath:
		switch method {
		case http.MethodGet:
			at.operationName = consts.GetSnapshotAdminOperationName
		}
	case consts.MigrateClusterV2OperationRouteFullPath:
		switch method {
		case http.MethodPost:
			at.operationName = consts.MigrateClusterV2OperationName
		}
	}
}

// GetOperationCategory returns the AuditEventCategory mapped to the current operation
func (at *APITracking) GetOperationCategory() AuditEventCategory {
	return getOperationCategory(at.operationName)
}

// GetOperationName returns the operation name
func (at *APITracking) GetOperationName() string {
	return at.operationName
}

// SetSubOperationName sets the suboperation name
func (at *APITracking) SetSubOperationName(subOperationName string) {
	at.subOperationName = subOperationName
}

// GetSubOperationName returns the suboperation name
func (at *APITracking) GetSubOperationName() string {
	return at.subOperationName
}

// SetHCPUnderlayName sets the HCP Underlay Name
func (at *APITracking) SetHCPUnderlayName(underlayName string) {
	at.hcpUnderlayName = underlayName
}

// GetHCPUnderlayName gets the HCP Underlay Name
func (at *APITracking) GetHCPUnderlayName() string {
	return at.hcpUnderlayName
}

// SetAuxiliaryToken sets aux token
func (at *APITracking) SetAuxiliaryToken(token string) {
	at.auxiliaryToken = token
}

// GetAuxiliaryToken gets aux token
func (at *APITracking) GetAuxiliaryToken() string {
	return at.auxiliaryToken
}

// SetHCPControlPlaneID sets the HCP Control Plane ID
func (at *APITracking) SetHCPControlPlaneID(cpid string) {
	at.hcpControlPlaneID = cpid
}

// GetHCPControlPlaneID gets the HCP Control Plane ID
func (at *APITracking) GetHCPControlPlaneID() string {
	return at.hcpControlPlaneID
}

// GetAADTenantID returns the AAD tenant id if it was provided by ARM
func (at *APITracking) GetAADTenantID() string {
	return at.aadTenantID
}

// GetSubscriptionID returns the subscription id if it was on the request else empty string
func (at *APITracking) GetSubscriptionID() uuid.UUID {
	return at.subscriptionID
}

func (at *APITracking) SetControlPlaneID(controlPlane string) {
	at.controlPlaneID = controlPlane
}

// GetSubscriptionIDString returns the subscription ID as a string. This avoids coupling the caller
// to a specific version of the UUID package as would occur when using GetSubscriptionID.
func (at *APITracking) GetSubscriptionIDString() string {
	return at.subscriptionID.String()
}

// GetAPIVersion returns the api version on the request
func (at *APITracking) GetAPIVersion() string {
	return at.apiVersion
}

// GetAcceptLanguage returns the accept language from the header on the request
func (at *APITracking) GetAcceptLanguage() string {
	return at.acceptLanguage
}

// GetHost returns the host name from the header on the request
func (at *APITracking) GetHost() string {
	return at.host
}

// GetReferer returns the referer on the request
func (at *APITracking) GetReferer() string {
	return at.referer
}

// GetResourceGroupName returns the resourceGroupName if it was on the request else empty string
func (at *APITracking) GetResourceGroupName() string {
	return at.resourceGroupName
}

// GetResourceName returns the resourceName if it was on the request else empty string
func (at *APITracking) GetResourceName() string {
	return at.resourceName
}

// GetDetectorName returns the detectorName if it was on the request else empty string
func (at *APITracking) GetDetectorName() string {
	return at.detectorName
}

// GetAgentPoolName returns the agentpoolName if it was an agentpool request, else empty string.
func (at *APITracking) GetAgentPoolName() string {
	return at.agentPoolName
}

// GetMaintenanceConfigurationName returns the maintenance configuration if it was a maintenance configuration request, else empty string.
func (at *APITracking) GetMaintenanceConfigurationName() string {
	return at.maintenanceConfigurationName
}

// GetExtensionAddonName returns the extension addon name if it was a extension addon request, else empty string.
func (at *APITracking) GetExtensionAddonName() string {
	return at.extensionAddonName
}

//GetVMName returns the vmName if it was on the request
func (at *APITracking) GetVMName() string {
	return at.vmName
}

// GetActionName returns the actionName if it was on the request
func (at *APITracking) GetActionName() string {
	return at.actionName
}

// GetResourceType returns the resourceType if it was on the request else empty string
func (at *APITracking) GetResourceType() string {
	return at.resourceType
}

// GetClientAppID returns the client app id
func (at *APITracking) GetClientAppID() string {
	return at.clientApplicationID
}

// GetUserAgent returns the user agent on the request
func (at *APITracking) GetUserAgent() string {
	return at.userAgent
}

// GetTargetURI returns the target URI on the request
func (at *APITracking) GetTargetURI() string {
	return at.targetURI
}

// Get GetHttpMethod the http method on the request
func (at *APITracking) GetHttpMethod() string {
	return at.httpMethod
}

// GetClientPrincipalName returns the client principal id
func (at *APITracking) GetClientPrincipalName() string {
	return at.clientPrincipalName
}

// GetClientRequestID returns the client request id
func (at *APITracking) GetClientRequestID() uuid.UUID {
	return at.clientRequestID
}

// GetClientSessionID returns the client session id
func (at *APITracking) GetClientSessionID() string {
	return at.clientSessionID
}

// GetMessageID returns the message id of a queue message
func (at *APITracking) GetMessageID() int {
	return at.messageID
}

// GetPopReceipt returns the pop receipt of a queue message
func (at *APITracking) GetPopReceipt() string {
	return at.popReceipt
}

// GetInsertionTime returns the insertion time of a queue message
func (at *APITracking) GetInsertionTime() time.Time {
	return at.insertionTime
}

// SetInsertionTime sets queue message insertion time
func (at *APITracking) SetInsertionTime(d time.Time) {
	at.insertionTime = d
}

// GetExpirationTime returns the expiration time of a queue message
func (at *APITracking) GetExpirationTime() time.Time {
	return at.expirationTime
}

// GetTimeNextVisible returns the next visible time of a queue message
func (at *APITracking) GetTimeNextVisible() time.Time {
	return at.timeNextVisible
}

// GetMessageTTLSec returns the message TTL in seconds
func (at *APITracking) GetMessageTTLSec() uint {
	return uint(at.messageTTL / time.Second)
}

// SetMessageTTL sets message TTL
func (at *APITracking) SetMessageTTL(ttl time.Duration) {
	at.messageTTL = ttl
}

// GetDelayStartInSeconds returns the message start delay in seconds
func (at *APITracking) GetDelayStartInSeconds() uint {
	return uint(at.delayStartInSeconds)
}

// SetDelayStartInSeconds set the message start delay in seconds
func (at *APITracking) SetDelayStartInSeconds(delayStartInSeconds uint) {
	at.delayStartInSeconds = delayStartInSeconds
}

// GetOperationTimeout returns operation execution time limit in seconds
func (at *APITracking) GetOperationTimeout() time.Duration {
	return at.operationTimeout
}

// SetOperationTimeout sets operation execution time limit
func (at *APITracking) SetOperationTimeout(d time.Duration) {
	at.operationTimeout = d
}

// SetHost sets operation host
func (at *APITracking) SetHost(h string) {
	at.host = h
}

// GetDequeueCount returns the dequeue count of a queue message
func (at *APITracking) GetDequeueCount() int {
	return at.dequeueCount
}

// GetPodName returns the podName if it was on the request else empty string
func (at *APITracking) GetPodName() string {
	return at.podName
}

// GetContainerName returns the container name if it was on the request else empty string
func (at *APITracking) GetContainerName() string {
	return at.containerName
}

// GetVirtualNetworkName returns the virtual network name if it was on the request else empty string
func (at *APITracking) GetVirtualNetworkName() string {
	return at.virtualNetworkName
}

// GetSubnetName returns the subnet name if it was on the request else empty string
func (at *APITracking) GetSubnetName() string {
	return at.subnetName
}

// GetAPITracking retrieves the APITracking struct pointer out of a context struct
func GetAPITracking(ctx context.Context) (*APITracking, bool) {
	retVal, bool := ctx.Value(apiTrackingKey).(*APITracking)
	return retVal, bool
}

// WithAPITracking appends the operation.APITracking onto a copy of the passed in context and returns the copy
func WithAPITracking(ctx context.Context, apiTracking *APITracking) context.Context {
	return context.WithValue(ctx, apiTrackingKey, apiTracking)
}

// GetAccessProfileRoleName returns the roleName
func (at *APITracking) GetAccessProfileRoleName() string {
	return at.accessProfileRoleName
}

// IsAdminOperation returns whether this operation is admin operation
func (at *APITracking) IsAdminOperation() bool {
	return strings.HasPrefix(at.operationName, consts.AdminOperationPrefix)
}

// GetPropertiesBag returns propertiesbag
func (at *APITracking) GetPropertiesBag() map[string]string {
	return at.propertiesBag
}

// ConvertPropertiesBagtoString is the function used to convert propertiesbag to string
func (at *APITracking) ConvertPropertiesBagtoString() string {
	if at.propertiesBag != nil {
		ret, err := json.Marshal(at.propertiesBag)
		if err != nil {
			return "Failed to convert propertiesBag into json."
		}
		return string(ret)
	}
	return ""
}

//GetAgentPoolPropertyFromPropertyBag Find the property value for the corresponding agentpool, return empty if property is not present
func (at *APITracking) getAgentPoolPropertyFromPropertyBag(agentPoolName, property string) string {
	// Check whether agentPool name is in the property bag as a value if yes it is an agentPool if not it is a managedCluster
	agentpoolPrefix := "agentpool"
	nameProperty := "name"
	delimiter := "_"
	var agentPoolIndex string
	for k, v := range at.propertiesBag {
		if v == agentPoolName {
			agentPoolIndex = strings.TrimPrefix(k, agentpoolPrefix+delimiter+nameProperty+delimiter)
			break
		}
	}
	return at.GetPropertiesBag()[agentpoolPrefix+delimiter+property+delimiter+agentPoolIndex]
}

//IsAgentPoolAZEnabled checks whether the agentPool is az enabled or not this will have a valid value for agent pool operations
func (at *APITracking) IsAgentPoolAZEnabled() bool {
	value := at.getAgentPoolPropertyFromPropertyBag(at.resourceName, AgentPoolAZEnabled)
	if len(value) == 0 {
		return false
	}
	result, _ := strconv.ParseBool(value)
	return result
}

//IsControlPlaneAZEnabled checks whether the cluster is az enabled or not
func (at *APITracking) IsControlPlaneAZEnabled() bool {
	value := at.propertiesBag[ControlPlaneAZEnabled]
	result, _ := strconv.ParseBool(value)
	return result
}

// GetK8sCurrentVersion is the function used to get k8s current Version
func (at *APITracking) GetK8sCurrentVersion() string {
	if val, ok := at.propertiesBag[k8sCurrentVersion]; ok {
		return val
	}
	return ""
}

// GetK8sGoalVersion is the function used to get k8s goal Version
func (at *APITracking) GetK8sGoalVersion() string {
	if val, ok := at.propertiesBag[k8sGoalVersion]; ok {
		return val
	}
	return ""
}

// SetPropertiesBag set the properties bag
func (at *APITracking) SetPropertiesBag(key, value string) {
	if at.propertiesBag == nil {
		at.propertiesBag = map[string]string{}
	}
	at.propertiesBag[key] = value
}

// AddResultCodeDependency sets the result code dependency
func (at *APITracking) AddResultCodeDependency(value string) {
	at.resultCodeDependency = append(at.resultCodeDependency, value)
}

// GetResultCodeDependency gets the result code dependency in comma separated string
func (at *APITracking) GetResultCodeDependency() string {
	m := map[string]int{}

	// de-duplicate
	for _, v := range at.resultCodeDependency {
		if len(v) > 0 {
			m[strings.ToLower(v)] = 1
		}
	}

	var res []string
	for k := range m {
		res = append(res, k)
	}

	// sort so we have consistent ordering
	sort.Strings(res)

	return strings.Join(res, ",")
}

// IsInternalSubscription returns true if the subscription ID is an internal subscription ID.
func (at *APITracking) IsInternalSubscription() bool {
	return at.isInternalSubscription
}

// GetExtensionProviderName return providerName of the extension resource
func (at *APITracking) GetExtensionProviderName() string {
	return at.extensionProviderName
}

// GetExtensionResourceType return type of the extension resource
func (at *APITracking) GetExtensionResourceType() string {
	return at.extensionResourceType
}

// GetExtensionResourceName return name of the extension resource
func (at *APITracking) GetExtensionResourceName() string {
	return at.extensionResourceName
}

// SetAgentPoolName sets the agentpoolName
func (at *APITracking) SetAgentPoolName(name string) {
	at.agentPoolName = name
}

// SetMaintenanceConfigurationName sets the maintenanceConfigurationName
func (at *APITracking) SetMaintenanceConfigurationName(name string) {
	at.maintenanceConfigurationName = name
}

// SetExtensionAddonName sets the extensionAddonName
func (at *APITracking) SetExtensionAddonName(name string) {
	at.extensionAddonName = name
}

// GetFQDN returns the fqdn
func (at *APITracking) GetFQDN() string {
	return at.fqdn
}

// SetFQDN sets the fqdn
func (at *APITracking) SetFQDN(fqdn string) {
	at.fqdn = fqdn
}

// GetCxUnderlayName returns the customer underlay name
func (at *APITracking) GetCxUnderlayName() string {
	return at.cxUnderlayName
}

// SetCxUnderlayName sets the customer underlay name
func (at *APITracking) SetCxUnderlayName(cxUnderlayName string) {
	at.cxUnderlayName = cxUnderlayName
}

// GetPrivateEndpointConnectionName returns the privateEndpointConnection if it was on the request else empty string
func (at *APITracking) GetPrivateEndpointConnectionName() string {
	return at.privateEndpointConnectionName
}

// GetCredentialFormat returns the credential format if it was on the request else empty string
func (at *APITracking) GetCredentialFormat() string {
	return at.credentialFormat
}

// GetLoginMethod returns the login method if it was on the request else empty string
func (at *APITracking) GetLoginMethod() string {
	return at.loginMethod
}

// GetControlPlaneID returns the control plane Id if it was on the request else empty string
func (at *APITracking) GetControlPlaneID() string {
	return at.controlPlaneID

}

// GetCredentialServerFQDNFormat returns the kubeconfig server type. It can be public or private for private cluster.
func (at *APITracking) GetCredentialServerFQDNFormat() string {
	return at.credentialServerFQDNFormat
}

// GetLabels returns the custom node labels if it was on the request else nil
func (at *APITracking) SetToBeDeletedAgentPoolLabels(labels map[string]map[string]string) error {
	labelsBytes, err := json.Marshal(labels)
	if err != nil {
		return err
	}
	at.propertiesBag[agentPoolLabelsToDelete] = string(labelsBytes)
	return nil
}

// GetLabels returns the custom node labels if it was on the request else nil
func (at *APITracking) GetToBeDeletedAgentPoolLabels() (map[string]map[string]string, error) {
	var labels map[string]map[string]string
	labelsString, ok := at.propertiesBag[agentPoolLabelsToDelete]
	if !ok {
		return labels, nil
	}
	err := json.Unmarshal([]byte(labelsString), &labels)
	return labels, err
}

func (at *APITracking) SetAgentPoolCount(count int) {
	at.propertiesBag[agentPoolCount] = strconv.Itoa(count)
}

func (at *APITracking) GetAgentPoolCount() (int, error) {
	return strconv.Atoi(at.propertiesBag[agentPoolCount])
}

// SetSpanContext propagates the SpanContext from sync to async
func (at *APITracking) SetSpanContext(sc trace.SpanContext) {
	bin := propagation.Binary(sc)
	at.propertiesBag[spanContext] = string(bin)
}

// GetSpanContext retrieves the SpanContext when handling a message
func (at *APITracking) GetSpanContext() (trace.SpanContext, bool) {
	bin := at.propertiesBag[spanContext]
	if bin != "" {
		return propagation.FromBinary([]byte(bin))
	}
	return trace.SpanContext{}, false
}

func (at *APITracking) SetErrorSubcode(subcode string) {
	at.errorSubcode = subcode
}

func (at *APITracking) GetErrorSubcode() string {
	return at.errorSubcode
}

func (at *APITracking) SetErrorCategory(category string) {
	at.errorCategory = category
}

func (at *APITracking) GetErrorCategory() string {
	return at.errorCategory
}

func (at *APITracking) SetErrorDependency(dep string) {
	at.errorDependency = dep
}

func (at *APITracking) GetErrorDependency() string {
	return at.errorDependency
}

func (at *APITracking) SetErrorAKSTeam(aksTeam string) {
	at.errorAKSTeam = aksTeam
}

func (at *APITracking) GetErrorAKSTeam() string {
	return at.errorAKSTeam
}
