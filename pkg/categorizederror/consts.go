package categorizederror

// ErrorSubCode used to assigning errorSubcode to apierror.SubCode
type ErrorSubCode string

const (
	Unknown         ErrorSubCode = "Unknown"
	EOF             ErrorSubCode = "EOF"
	Conflict        ErrorSubCode = "Conflict"
	ContextCanceled ErrorSubCode = "ContextCanceled"

	FailedPrecondition            ErrorSubCode = "FailedPrecondition"
	NilSubscription               ErrorSubCode = "NilSubscription"
	EmptyLBList                   ErrorSubCode = "EmptyLBList"
	EnsureLoadBalancersExistError ErrorSubCode = "EnsureLoadBalancersExistError"
	FailedGetVMSSName             ErrorSubCode = "FailedGetVMSSName"
	NilImageReference             ErrorSubCode = "NilImageReference"
	EmptyNodeResourceGroup        ErrorSubCode = "EmptyNodeResourceGroup"

	//overlaymgr
	Timedout                              ErrorSubCode = "Timedout"
	OverlayIDNotFound                     ErrorSubCode = "OverlayIDNotFound"
	FailedGetOverlay                      ErrorSubCode = "FailedGetOverlay"
	FailedReconcileOverlay                ErrorSubCode = "FailedReconcileOverlay"
	FailedReconcileAgentPoolConfigmap     ErrorSubCode = "FailedReconcileAgentPoolConfigmap"
	FailedReconcileCloudConfigSecret      ErrorSubCode = "FailedReconcileCloudConfigSecret"
	FailedReconcileNetworkAssociationLink ErrorSubCode = "FailedReconcileNetworkAssociationLink"
	FailedWaitForReadiness                ErrorSubCode = "FailedWaitForReadiness"
	FailedCleanupPrivateConnectResources  ErrorSubCode = "FailedCleanupPrivateConnectResources"

	NoResourceReference                             ErrorSubCode = "NoResourceReference"
	InvalidResourceReference                        ErrorSubCode = "InvalidResourceReference"
	NilManagedCluster                               ErrorSubCode = "NilManagedCluster"
	NilManagedClusterProperties                     ErrorSubCode = "NilManagedClusterProperties"
	FailedGetKubeconfigFromCCP                      ErrorSubCode = "FailedGetKubeconfigFromCCP"
	NilAgentPool                                    ErrorSubCode = "NilAgentPool"
	NilAgentPoolProperties                          ErrorSubCode = "NilAgentPoolProperties"
	NilAgentPoolSinglePlacementGroup                ErrorSubCode = "NilAgentPoolSinglePlacementGroup"
	FailedGetAgentVMSSName                          ErrorSubCode = "FailedGetAgentVMSSName"
	FailedGetLatestImageReference                   ErrorSubCode = "FailedGetLatestImageReference"
	FailedGetSIGAzureCloudSpecConfig                ErrorSubCode = "FailedGetSIGAzureCloudSpecConfig"
	NoOrchestratorVersion                           ErrorSubCode = "NoOrchestratorVersion"
	EmptyPrivateIP                                  ErrorSubCode = "EmptyPrivateIP"
	EmptyVMSSNameForVMSSReconcilerGoal              ErrorSubCode = "EmptyVMSSNameForVMSSReconcilerGoal"
	NotOneAgentPoolProfileForVMSSReconciler         ErrorSubCode = "NotOneAgentPoolProfileForVMSSReconciler"
	AgentPoolNameNotMatched                         ErrorSubCode = "AgentPoolNameNotMatched"
	FailedGetVMSSInstanceIDsAndNodeNameForScaleDown ErrorSubCode = "FailedGetVMSSInstanceIDsAndNodeNameForScaleDown"
	FailedPopulateAgentPoolImageReference           ErrorSubCode = "FailedPopulateAgentPoolImageReference"
	FailedPrepareContainerDatamodel                 ErrorSubCode = "FailedPrepareContainerDatamodel"
	FailedConstructVMSSObject                       ErrorSubCode = "FailedConstructVMSSObject"
	TimedoutWaitingForVMSSCreation                  ErrorSubCode = "TimedoutWaitingForVMSSCreation"
	VMSSNotFoundAfterCreating                       ErrorSubCode = "VMSSNotFoundAfterCreating"
	VMSSInstancesCSEFailure                         ErrorSubCode = "VMSSInstancesCSEFailure"
	UnexpectedVMSSProvisioningState                 ErrorSubCode = "UnexpectedVMSSProvisioningState"
	NotFinalStateOfVMSSProvisioningState            ErrorSubCode = "NotFinalStateOfVMSSProvisioningState"
	FailedPreserveLoadBalancerBackendAddressPoolIDs ErrorSubCode = "FailedPreserveLoadBalancerBackendAddressPoolIDs"

	// availability set reconciler
	UnexpectedAgentPoolForAvailabilitySet ErrorSubCode = "UnexpectedAgentPoolForAvailabilitySet"
	FailedCreateAvailabilitySetObject     ErrorSubCode = "FailedCreateAvailabilitySetObject"

	// outbound reconciler
	EmptyGoalForOutboundReconciler                     ErrorSubCode = "EmptyGoalForOutboundReconciler"
	EmptyStandardLoadBalancerGoalForOutboundReconciler ErrorSubCode = "EmptyStandardLoadBalancerGoalForOutboundReconciler"
	EmptyNATGatewayGoalForOutboundReconciler           ErrorSubCode = "EmptyNATGatewayGoalForOutboundReconciler"
	NotStandardLoadBalancer                            ErrorSubCode = "NotStandardLoadBalancer"
	NoOutboundRulesForSLB                              ErrorSubCode = "NoOutboundRulesForSLB"
	NoExactOneOutboundRuleForSLB                       ErrorSubCode = "NoExactOneOutboundRuleForSLB"
	NoExactTwoOutboundRulesForSLBWithDualStack         ErrorSubCode = "NoExactTwoOutboundRulesForSLBWithDualStack"
	NilPublicIPAddressFields                           ErrorSubCode = "NilPublicIPAddressFields"
	NilPublicIPPrefix                                  ErrorSubCode = "NilPublicIPPrefixFields"
	NoOutboundPublicIPOrPrefix                         ErrorSubCode = "NoOutboundPublicIPOrPrefix"
	NoOutboundPublicIPOrPrefixForLB                    ErrorSubCode = "NoOutboundPublicIPOrPrefixForLB"
	NoIDForOutboundIP                                  ErrorSubCode = "NoIDForOutboundIP"
	NoIPAddressForOutboundIP                           ErrorSubCode = "NoIPAddressForOutboundIP"
	LBNotInTerminatingState                            ErrorSubCode = "LBNotInTerminatingState"
	FailedListLB                                       ErrorSubCode = "FailedListLB"
	NotStandardNATGateway                              ErrorSubCode = "NotStandardNATGateway"
	NotRegionalNATGateway                              ErrorSubCode = "NotRegionalNATGateway"
	UnexpectedNATGatewayLocation                       ErrorSubCode = "UnexpectedNATGatewayLocation"
	NATGatewayNotInTerminatingState                    ErrorSubCode = "NATGatewayNotInTerminatingState"

	// route table reconciler
	MissingManagedClusterProperties ErrorSubCode = "MissingManagedClusterProperties"
	FailedParseResourceID           ErrorSubCode = "FailedParseResourceID"
	NilGoalForRouteTableReconciler  ErrorSubCode = "NilGoalForRouteTableReconciler"
	NilRouteTableAfterPut           ErrorSubCode = "NilRouteTableAfterPut"
	UnexpectedRouteTable            ErrorSubCode = "UnexpectedRouteTable"

	// billing usage reconciler
	NoControlPlaneID                           ErrorSubCode = "NoControlPlaneID"
	ResourceReferenceAndBillingUsageNotMatched ErrorSubCode = "ResourceReferenceAndBillingUsageNotMatched"
	UnexpectedBillingUsage                     ErrorSubCode = "UnexpectedBillingUsage"
	NilGoalForBillingUsageReconciler           ErrorSubCode = "NilGoalForBillingUsageReconciler"

	// private endpoint reconciler
	NilCCPWrapper                         ErrorSubCode = "NilCCPWrapper"
	NilCCPWrapperProperties               ErrorSubCode = "NilCCPWrapperProperties"
	NilUnderlay                           ErrorSubCode = "NilUnderlay"
	IncorrectFormatOfPrivateLinkServiceID ErrorSubCode = "IncorrectFormatOfPrivateLinkServiceID"
	PrivateEndpointIsUnderUpdating        ErrorSubCode = "PrivateEndpointIsUnderUpdating"
	UnexpectedPrivateEndpoint             ErrorSubCode = "UnexpectedPrivateEndpoint"
	ScopeLocked                           ErrorSubCode = "ScopeLocked"

	// vmas reconciler
	NilManagedClusterHostMasterProfile ErrorSubCode = "NilManagedClusterHostMasterProfile"
	NilAPITracking                     ErrorSubCode = "NilAPITracking"
	FailedGetVMClient                  ErrorSubCode = "FailedGetVMClient"
	UnexpectedVMASProvisioningState    ErrorSubCode = "UnexpectedVMASProvisioningState"
	FailedGetCurrentAgentPoolState     ErrorSubCode = "FailedGetCurrentAgentPoolState"
	FailedCloneManagedCluster          ErrorSubCode = "FailedCloneManagedCluster"
	NilTargetAgentPoolProfile          ErrorSubCode = "NilTargetAgentPoolProfile"
	UnexpectedGoalAgentCountForVMAS    ErrorSubCode = "UnexpectedGoalAgentCountForVMAS"
	FailedGetNetworkInterfaceClient    ErrorSubCode = "FailedGetNetworkInterfaceClient"
	NilVHDForOSDisk                    ErrorSubCode = "NilVHDForOSDisk"
	FailedGetNicName                   ErrorSubCode = "FailedGetNicName"
	FailedParseVHDURI                  ErrorSubCode = "FailedParseVHDURI"
	FailedGetAccountNameFromVHDURL     ErrorSubCode = "FailedGetAccountNameFromVHDURL"
	FailedGetStorageDataClient         ErrorSubCode = "FailedGetStorageDataClient"
	FailedDeleteBlob                   ErrorSubCode = "FailedDeleteBlob"
	FailedGetDiskClient                ErrorSubCode = "FailedGetDiskClient"
	EmptyVMName                        ErrorSubCode = "EmptyVMName"
	FailedGetVMExtensionClient         ErrorSubCode = "FailedGetVMExtensionClient"
	FailedGetK8sVMName                 ErrorSubCode = "FailedGetK8sVMName"
	NotAvailabilitySet                 ErrorSubCode = "NotAvailabilitySet"
	NoPoolToUpgrade                    ErrorSubCode = "NoPoolToUpgrade"
	WindowsVMASNotSupported            ErrorSubCode = "WindowsVMASNotSupported"
	FailedGetK8sLinuxVMName            ErrorSubCode = "FailedGetK8sLinuxVMName"
	FailedGetKubeClient                ErrorSubCode = "FailedGetKubeClient"
	AgentCountNotMatch                 ErrorSubCode = "AgentCountNotMatch"
	FailedGetResourcesClient           ErrorSubCode = "FailedGetResourcesClient"

	// retrievers
	EmptyControlPlaneID           ErrorSubCode = "EmptyControlPlaneID"
	FailedSetCustomHyperKubeImage ErrorSubCode = "FailedSetCustomHyperKubeImage"
	FailedSetCustomKubeProxyImage ErrorSubCode = "FailedSetCustomKubeProxyImage"
	FailedSetCustomKubeBinaryURL  ErrorSubCode = "FailedSetCustomKubeBinaryURL"

	// vnet reconciler
	PutVnetFailedVnetNil        ErrorSubCode = "PutVnetFailedVnetNil"
	PutVnetFailedUnexpectedVnet ErrorSubCode = "PutVnetFailedUnexpectedVnet"

	// blob storage reconciler
	NilBlobClient ErrorSubCode = "NilBlobClient"

	// pls reconciler
	NilPrivateLinkServiceSettings ErrorSubCode = "NilPrivateLinkServiceSettings"

	// subnet reconciler
	SubnetPrefixMisMatch        ErrorSubCode = "SubnetPrefixMisMatch"
	SubnetInDeletingState       ErrorSubCode = "SubnetInDeletingState"
	EmptyILBServiceUID          ErrorSubCode = "EmptyILBServiceUID"
	EmptyUnderlaySubscriptionID ErrorSubCode = "EmptyUnderlaySubscriptionID"
	EmptyUnderlayTenantID       ErrorSubCode = "EmptyUnderlayTenantID"
	EmptyUnderlayName           ErrorSubCode = "EmptyUnderlayName"
	UnexpectedVnetNumber        ErrorSubCode = "UnexpectedVnetNumber"
	InvalidVnet                 ErrorSubCode = "InvalidVnet"

	// subnet delegation
	FailedToRegisterWithTokenService        ErrorSubCode = "FailedToRegisterWithTokenService"
	FailedToReadFromResponse                ErrorSubCode = "FailedToReadFromResponse"
	FailedToGetVNetClient                   ErrorSubCode = "FailedToGetVNetClient"
	FailedToGetSubnetClient                 ErrorSubCode = "FailedToGetSubnetClient"
	SubnetNotFound                          ErrorSubCode = "SubnetNotFound"
	FailedToPostSubnetToken                 ErrorSubCode = "FailedToPostSubnetToken"
	EmptySubnetNameOrGUID                   ErrorSubCode = "EmptySubnetNameOrGUID"
	FailedToUnregisterWithTokenService      ErrorSubCode = "FailedToUnregisterWithTokenService"
	FailedToDeleteSubnetFromTokenService    ErrorSubCode = "FailedToDeleteSubnetFromTokenService"
	ErrorRetrievingNodeNetworkConfiguration ErrorSubCode = "ErrorRetrievingNodeNetworkConfiguration"
	MoreThanOneNodeNetworkConfiguration     ErrorSubCode = "MoreThanOneNodeNetworkConfiguration"

	// private dns reconciler
	PrivateDNSRecordMetadataNotMatchError ErrorSubCode = "PrivateDNSRecordMetadataNotMatchError"
	PrivateIPNotExist                     ErrorSubCode = "PrivateIPNotExist"

	// dns reconciler
	EmptyControlPlaneIP         ErrorSubCode = "EmptyControlPlaneIP"
	EmptyIngressFQDN            ErrorSubCode = "EmptyIngressFQDN"
	NilPrivateClusterProperties ErrorSubCode = "NilPrivateClusterProperties"

	// ccp reconciler
	NilControlPlane ErrorSubCode = "NilControlPlane"
	// resolvers
	FailedGenerateOpenVPNProfile           ErrorSubCode = "FailedGenerateOpenVPNProfile"
	FailedGenerateAllAccessProfiles        ErrorSubCode = "FailedGenerateAllAccessProfiles"
	FailedGetLogAnalyticsWorkspaceKey      ErrorSubCode = "FailedGetLogAnalyticsWorkspaceKey"
	FailedGetLogAnalyticsWorkspaceID       ErrorSubCode = "FailedGetLogAnalyticsWorkspaceID"
	EmptyCPWFQDN                           ErrorSubCode = "EmptyCPWFQDN"
	FailedGeneratePortalFQDN               ErrorSubCode = "FailedGeneratePortalFQDN"
	FailedGenerateOIDCProfile              ErrorSubCode = "FailedGenerateOIDCProfile"
	FailedGeneratePodIdentityV2Profile     ErrorSubCode = "FailedGeneratePodIdentityV2Profile"
	FailedGenerateGatekeeperV3Profile      ErrorSubCode = "FailedGenerateGatekeeperV3Profile"
	FailedGenerateAzurePolicyV2Profile     ErrorSubCode = "FailedGenerateAzurePolicyV2Profile"
	FailedGenerateCertificates             ErrorSubCode = "FailedGenerateCertificates"
	FailedGenerateWebhookCertificates      ErrorSubCode = "FailedGenerateWebhookCertificates"
	FailedGenerateAggregatorCertificates   ErrorSubCode = "FailedGenerateAggregatorCertificates"
	FailedGenerateETCDCertificates         ErrorSubCode = "FailedGenerateETCDCertificates"
	FailedGenerateCBCEncryptionConfig      ErrorSubCode = "FailedGenerateCBCEncryptionConfig"
	FailedGenerateDefaultEncryptionConfig  ErrorSubCode = "FailedGenerateDefaultEncryptionConfig"
	FailedRotateETCDBackupEncryptionConfig ErrorSubCode = "FailedRotateETCDBackupEncryptionConfig"
	FailedGenerateACIConnectorCertificates ErrorSubCode = "FailedGenerateACIConnectorCertificates"
	NilUnderlays                           ErrorSubCode = "NilUnderlays"
	NilFilteredUnderlays                   ErrorSubCode = "NilFilteredUnderlays"
	NilOverlaymgrClient                    ErrorSubCode = "NilOverlaymgrClient"
	FailedOverlaymgrHealthCheck            ErrorSubCode = "FailedOverlaymgrHealthCheck"
	ValidateUnderlayPanic                  ErrorSubCode = "ValidateUnderlayPanic"
	TimedoutValidateUnderlay               ErrorSubCode = "TimedoutValidateUnderlay"
	NoSuitableUnderlay                     ErrorSubCode = "NoSuitableUnderlay"
	OverConstrainedUnderlay                ErrorSubCode = "OverConstrainedUnderlay"
	FailedGetSuitableUnderlay              ErrorSubCode = "FailedGetSuitableUnderlay"
	InvalidTenantID                        ErrorSubCode = "InvalidTenantID"
	OverlayWithoutPrivateConnectIP         ErrorSubCode = "OverlayWithoutPrivateConnectIP"
	OverlayWithoutServiceIP                ErrorSubCode = "OverlayWithoutServiceIP"
	OverlayWithoutPrivateLinkServiceName   ErrorSubCode = "OverlayWithoutPrivateLinkServiceName"
	CCPWithoutServiceIP                    ErrorSubCode = "CCPWithoutServiceIP"
	EmptyUnderlayKubeConfig                ErrorSubCode = "EmptyUnderlayKubeConfig"
	ServiceAssociationLinkDetailNotExist   ErrorSubCode = "ServiceAssociationLinkDetailNotExist"
	EmptyPrimaryContextID                  ErrorSubCode = "EmptyPrimaryContextID"
	EmptyVNetResourceGUID                  ErrorSubCode = "EmptyVNetResourceGUID"
	NilPrivateEndpointSetting              ErrorSubCode = "NilPrivateEndpointSetting"
	NilPrivateLinkProfile                  ErrorSubCode = "NilPrivateLinkProfile"
	FailedGetPrivateLinkServiceName        ErrorSubCode = "FailedGetPrivateLinkServiceName"
	InvalidPrivateEndpoint                 ErrorSubCode = "InvalidPrivateEndpoint"
	InvalidNetworkInterface                ErrorSubCode = "InvalidNetworkInterface"
	PrivateEndpointNeedOneNIC              ErrorSubCode = "PrivateEndpointNeedOneNIC"
	NICNeedOneIPConfiguration              ErrorSubCode = "NICNeedOneIPConfiguration"
	InvalidIPConfiguration                 ErrorSubCode = "InvalidIPConfiguration"
	FailedGenerateCCPWindowsGmsaProfile    ErrorSubCode = "FailedGenerateCCPWindowsGmsaProfile"

	// nsg reconciler
	NilNSGAfterCreation            ErrorSubCode = "NilNSGAfterCreation"
	UnexpectedNSGProvisioningState ErrorSubCode = "UnexpectedNSGProvisioningState"

	// private dns zone reconciler
	UnexpectedPrivateDNSZoneProvisioningState ErrorSubCode = "UnexpectedPrivateDNSZoneProvisioningState"

	// resource provider registration reconciler
	NilResourceProvider ErrorSubCode = "NilResourceProvider"

	// msi credential store
	FailedToMarshalMSICredential   ErrorSubCode = "FailedToMarshalMSICredential"
	FailedToUnmarshalMSICredential ErrorSubCode = "FailedToUnmarshalMSICredential"
	InvalidMSICredentialParameter  ErrorSubCode = "InvalidMSICredentialParameter"

	// msi credential reconciler
	UnexpectedIdentityType ErrorSubCode = "UnexpectedIdentityType"

	// service associate link reconciler
	VNetNotFound                             ErrorSubCode = "VNetNotFound"
	NilVNetPropertiesFormat                  ErrorSubCode = "NilVNetPropertiesFormat"
	EmptyResourceGUIDForVNetPropertiesFormat ErrorSubCode = "EmptyResourceGUIDForVNetPropertiesFormat"

	// role assignment reconciler
	MissingRequiredProperty    ErrorSubCode = "MissingRequiredProperty"
	FailedDecodeARMError       ErrorSubCode = "FailedDecodeARMError"
	NilServicePrincipalProfile ErrorSubCode = "NilServicePrincipalProfile"
	/* #nosec */
	FailedLinearRetryRefreshToken ErrorSubCode = "FailedLinearRetryRefreshToken"
	ServicePrincipalNotFound      ErrorSubCode = "ServicePrincipalNotFound"
	FailedConstructOauthToken     ErrorSubCode = "FailedConstructOauthToken"
	FailedConstructSPToken        ErrorSubCode = "FailedConstructSPToken"
	FailedRefreshToken            ErrorSubCode = "FailedRefreshToken"
	FailedGetRoleAssignmentClient ErrorSubCode = "FailedGetRoleAssignmentClient"

	// reconcilers related to overlaymgr
	NilDiagnosticSetting          ErrorSubCode = "NilDiagnosticSetting"
	EmptyUnderlayKubeconfig       ErrorSubCode = "EmptyUnderlayKubeconfig"
	EmptyCPWKubeConfig            ErrorSubCode = "EmptyCPWKubeConfig"
	FailedDecodeCPWKubeConfig     ErrorSubCode = "FailedDecodeCPWKubeConfig"
	RetryNeededForDeletingOverlay ErrorSubCode = "RetryNeededForDeletingOverlay"

	// private dns vlink reconciler
	UnexpectedPrivateDNSVNetLinkProvisioningState ErrorSubCode = "UnexpectedPrivateDNSVNetLinkProvisioningState"

	// connectivity errors
	ConnectionResetByPeer   ErrorSubCode = "ConnectionResetByPeer"
	ConnectionRefused       ErrorSubCode = "ConnectionRefused"
	ContextDeadlineExceeded ErrorSubCode = "ContextDeadlineExceeded"
	ConnectionTimedout      ErrorSubCode = "ConnectionTimedout"
	TLSHandshakeTimedout    ErrorSubCode = "TLSHandshakeTimedout"
	IOTimedout              ErrorSubCode = "IOTimedout"
	LookupNoSuchHost        ErrorSubCode = "LookupNoSuchHost"

	// private connect reconciler
	NotPrivateConnectCluster    ErrorSubCode = "NotPrivateConnectCluster"
	UnexpectedStateLoadBalancer ErrorSubCode = "UnexpectedStateLoadBalancer"
	NilLoadBalancer             ErrorSubCode = "NilLoadBalancer"

	// OIDC profile reconciler
	FailedGenerateOpenIDMetadata ErrorSubCode = "FailedGenerateOpenIDMetadata"

	FailedPrepareRequest ErrorSubCode = "FailedPrepareRequest"

	// vmss related
	FailedValidateVMSSParameters           ErrorSubCode = "FailedValidateVMSSParameters"
	FailedValidateVMSSRunCommandParameters ErrorSubCode = "FailedValidateVMSSRunCommandParameters"
	VMSSNotFound                           ErrorSubCode = "VMSSNotFound"
	VMSSNilSKUOrCapacity                   ErrorSubCode = "VMSSNilSKUOrCapacity"
	VMSSCreationFailed                     ErrorSubCode = "VMSSCreationFailed"

	// resource group client
	FailedValidateRGName       ErrorSubCode = "FailedValidateRGName"
	FailedValidateRGParameters ErrorSubCode = "FailedValidateRGParameters"

	// Key vault secret client
	FailedValidateSecretName          ErrorSubCode = "FailedValidateSecretName"
	FailedValidateSecretSetParameters ErrorSubCode = "FailedValidateSecretSetParameters"
	FailedValidateSecretMaxResults    ErrorSubCode = "FailedValidateSecretMaxResults"

	// blob storage client
	UnexpectedStatusCodeCheckContainerExist ErrorSubCode = "UnexpectedStatusCodeCheckContainerExist"
	UnexpectedStatusCodeDeleteContainer     ErrorSubCode = "UnexpectedStatusCodeDeleteContainer"
	UnexpectedStatusCodeCreateContainer     ErrorSubCode = "UnexpectedStatusCodeCreateContainer"
	UnexpectedStatusCodePutBlob             ErrorSubCode = "UnexpectedStatusCodePutBlob"
	UnexpectedStatusCodeCheckBlobExist      ErrorSubCode = "UnexpectedStatusCodeCheckBlobExist"
	UnexpectedStatusCodeAcquireLease        ErrorSubCode = "UnexpectedStatusCodeAcquireLease"
	FailedGetResponse                       ErrorSubCode = "FailedGetResponse"

	// client errors
	RequestDisallowedByPolicy ErrorSubCode = "RequestDisallowedByPolicy"
	PolicyViolation           ErrorSubCode = "PolicyViolation"
	DiskEncryptionSetError    ErrorSubCode = "DiskEncryptionSetError"
	InvalidParameter          ErrorSubCode = "InvalidParameter"

	// k8s client related
	UnexpectedNilRetryResult  ErrorSubCode = "UnexpectedNilRetryResult"
	UnexpectedRetryResultType ErrorSubCode = "UnexpectedRetryResultType"
	WaitForDeleteTimedOut     ErrorSubCode = "WaitForDeleteTimedOut"
	DrainDidNotComplete       ErrorSubCode = "DrainDidNotComplete"
	K8SFailedGetNode          ErrorSubCode = "K8SFailedGetNode"
	K8SFailedDrainNode        ErrorSubCode = "K8SFailedDrainNode"
	K8SFailedSetUnschedulable ErrorSubCode = "K8SFailedSetUnschedulable"
	K8SFailedUncordon         ErrorSubCode = "K8SFailedUncordon"

	// failed get client
	FailedGetResourceGroupClient                ErrorSubCode = "FailedGetResourceGroupClient"
	FailedGetDenyAssignmentClient               ErrorSubCode = "FailedGetDenyAssignmentClient"
	FailedGetResourceProviderClient             ErrorSubCode = "FailedGetResourceProviderClient"
	FailedGetMSICredentialClient                ErrorSubCode = "FailedGetMSICredentialClient"
	FailedGetUserAssignedIdentityClient         ErrorSubCode = "FailedGetUserAssignedIdentityClient"
	FailedGetLoadBalancerClient                 ErrorSubCode = "FailedGetLoadBalancerClient"
	FailedGetAzureResourcesClient               ErrorSubCode = "FailedGetAzureResourcesClient"
	FailedGetDNSClient                          ErrorSubCode = "FailedGetDNSClient"
	FailedGetAvailabilitySetClient              ErrorSubCode = "FailedGetAvailabilitySetClient"
	FailedGetVMSSClient                         ErrorSubCode = "FailedGetVMSSClient"
	FailedGetSecurityGroupClient                ErrorSubCode = "FailedGetSecurityGroupClient"
	FailedGetRouteTableClient                   ErrorSubCode = "FailedGetRouteTableClient"
	FailedGetVNetClient                         ErrorSubCode = "FailedGetVNetClient"
	FailedGetSubnetClient                       ErrorSubCode = "FailedGetSubnetClient"
	RouteTableNotFound                          ErrorSubCode = "RouteTableNotFound"
	FailedGetPrivateEndpointClient              ErrorSubCode = "FailedGetPrivateEndpointClient"
	FailedGetPrivateClusterDNSResourceReference ErrorSubCode = "FailedGetPrivateClusterDNSResourceReference"
	FailedGetPrivateDNSClient                   ErrorSubCode = "FailedGetPrivateDNSClient"
	FailedGetControlPlaneV1                     ErrorSubCode = "FailedGetControlPlaneV1"
	FailedGetPrivateLinkServiceClient           ErrorSubCode = "FailedGetPrivateLinkServiceClient"
	FailedGetNetworkSecurityGroupClient         ErrorSubCode = "FailedGetNetworkSecurityGroupClient"
)

// ConnectivityErrors define mapping between "connectivity keyword error" and ErrorSubCode
var ConnectivityErrors = map[string]ErrorSubCode{
	"context deadline exceeded": ContextDeadlineExceeded,
	"connection refused":        ConnectionRefused,
	"connection reset by peer":  ConnectionResetByPeer,
	"connection timed out":      ConnectionTimedout,
	"TLS handshake timeout":     TLSHandshakeTimedout,
	"i/o timeout":               IOTimedout,
	"no such host":              LookupNoSuchHost,
}

// ResourceType
type ResourceType string

const (
	// ARM resources dependency
	ADAL                          ResourceType = "ADAL Azure Directory Authentication Library"
	ARM                           ResourceType = "Azure Resource Manager"
	ResourceProvider              ResourceType = "Azure Resource Provider"
	NetworkUsage                  ResourceType = "Azure Network Usage"
	ResourceGroup                 ResourceType = "Azure Resource Group"
	AzureResouce                  ResourceType = "Azure Resource"
	AADGraph                      ResourceType = "AAD Graph"
	MSICredential                 ResourceType = "MSI Credential"
	BlobStorage                   ResourceType = "Blob Storage"
	VMSS                          ResourceType = "Microsoft.Compute/VirtualMachineScaleSet"
	AvailabilitySets              ResourceType = "Microsoft.Compute/AvailabilitySets"
	Disks                         ResourceType = "Microsoft.Compute/Disks"
	ResouceSKU                    ResourceType = "Microsoft.Compute/Skus"
	VirtualMachines               ResourceType = "Microsoft.Compute/virtualMachines"
	VirtualMachineExtensions      ResourceType = "Microsoft.Compute/virtualMachines/extensions"
	ProximityPlacementGroups      ResourceType = "Microsoft.Compute/ProximityPlacementGroups"
	LB                            ResourceType = "Microsoft.Network/LoadBalancers"
	ApplicationGateways           ResourceType = "Microsoft.Network/ApplicationGateways"
	DNSZones                      ResourceType = "Microsoft.Network/DNSZones"
	NetworkInterfaces             ResourceType = "Microsoft.Network/NetworkInterfaces"
	RouteTables                   ResourceType = "Microsoft.Network/routeTables"
	NetworkSecurityGroups         ResourceType = "Microsoft.Network/networkSecurityGroups"
	ServiceAssociationLinks       ResourceType = "Microsoft.Network/virtualNetworks/serviceAssociationLinks"
	PrivateDnsZones               ResourceType = "Microsoft.Network/PrivateDnsZones"
	PrivateEndpoints              ResourceType = "Microsoft.Network/PrivateEndpoints"
	PrivateLinkServices           ResourceType = "Microsoft.Network/PrivateLinkServices"
	PublicIPAddresses             ResourceType = "Microsoft.Network/PublicIPAddresses"
	PublicIPPrefixes              ResourceType = "Microsoft.Network/PublicIPPrefixes"
	VirtualNetworks               ResourceType = "Microsoft.Network/virtualNetworks"
	Subnets                       ResourceType = "Microsoft.Network/virtualNetworks/subnets"
	NATGateways                   ResourceType = "Microsoft.Network/natGateways"
	AuthorizationCheckAccess      ResourceType = "Microsoft.Authorization/CheckAccess"
	Lock                          ResourceType = "Microsoft.Authorization/Locks"
	RoleAssignments               ResourceType = "Microsoft.Authorization/roleAssignments"
	DenyAssignments               ResourceType = "Microsoft.Authorization/denyAssignments"
	DiagnosticSetting             ResourceType = "Microsoft.Insights/DiagnosticSettings"
	Deployment                    ResourceType = "Microsoft.Resources/Deployments"
	StorageAccounts               ResourceType = "Microsoft.Storage/storageAccounts"
	ManagedIdentity               ResourceType = "Microsoft.ManagedIdentity"
	UserAssignedIdentities        ResourceType = "Microsoft.ManagedIdentity/userAssignedIdentities"
	Solutions                     ResourceType = "Microsoft.OperationsManagement/solutions"
	OperationalInsightsWorkspaces ResourceType = "Microsoft.OperationalInsights/workspaces"
	FeatureProvider               ResourceType = "Microsoft.Features/providers/features"

	KeyVaultSecret ResourceType = "KeyVaultSecret"
	DNC            ResourceType = "DNC"

	// Other components' dependency
	//
	// These are internal to AKS. When failure is in one of our internal components, it should count
	// as AKS failure when calculating QoS. We have a stored function, isAksServiceFailure, in our Kusto
	// database to determine if an error is AKS's or not. If new internal dependency is added here,
	// that function also need to be updated accordingly.
	HCP           ResourceType = "HostedControlPlaneDataStore"
	KubernetesAPI ResourceType = "Kubernetes API"
	Overlaymgr    ResourceType = "Overlaymgr"
)

var NonReconcileSubcodes = map[string]bool{
	string(NilSubscription): true,
}
