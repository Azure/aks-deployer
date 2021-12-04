// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package apierror

// ErrorCategory indicates the kind of error
type ErrorCategory string

const (
	// ClientError is expected error
	ClientError ErrorCategory = "ClientError"
	// InternalError is system or internal error
	InternalError ErrorCategory = "InternalError"
)

// ErrorCode represents the more detailed error message
type ErrorCode string

const (
	// These codes are from Microsoft.Azure.ResourceProvider.API.ErrorCode.
	// They do not have individual documentation because they are just mirroring
	// existing codes

	InvalidParameter                                        ErrorCode = "InvalidParameter"
	BadRequest                                              ErrorCode = "BadRequest"
	NotFound                                                ErrorCode = "NotFound"
	Conflict                                                ErrorCode = "Conflict"
	PreconditionFailed                                      ErrorCode = "PreconditionFailed"
	OperationNotAllowed                                     ErrorCode = "OperationNotAllowed"
	OperationPreempted                                      ErrorCode = "OperationPreempted"
	PropertyChangeNotAllowed                                ErrorCode = "PropertyChangeNotAllowed"
	InternalOperationError                                  ErrorCode = "InternalOperationError"
	TooManyRequests                                         ErrorCode = "TooManyRequests"
	InvalidSubscriptionStateTransition                      ErrorCode = "InvalidSubscriptionStateTransition"
	UnregisterWithResourcesNotAllowed                       ErrorCode = "UnregisterWithResourcesNotAllowed"
	InvalidParameterConflictingProperties                   ErrorCode = "InvalidParameterConflictingProperties"
	SubscriptionNotRegistered                               ErrorCode = "SubscriptionNotRegistered"
	SubscriptionNotFound                                    ErrorCode = "SubscriptionNotFound"
	ConflictingUserInput                                    ErrorCode = "ConflictingUserInput"
	ProvisioningInternalError                               ErrorCode = "ProvisioningInternalError"
	ProvisioningFailed                                      ErrorCode = "ProvisioningFailed"
	NetworkingInternalOperationError                        ErrorCode = "NetworkingInternalOperationError"
	QuotaExceeded                                           ErrorCode = "QuotaExceeded"
	Unauthorized                                            ErrorCode = "Unauthorized"
	ResourcesOverConstrained                                ErrorCode = "ResourcesOverConstrained"
	ControlPlaneProvisioningInternalError                   ErrorCode = "ControlPlaneProvisioningInternalError"
	ControlPlaneProvisioningSyncError                       ErrorCode = "ControlPlaneProvisioningSyncError"
	ControlPlaneInternalError                               ErrorCode = "ControlPlaneInternalError"
	ControlPlaneUnexpectedValue                             ErrorCode = "ControlPlaneUnexpectedValue"
	ControlPlaneNotFound                                    ErrorCode = "ControlPlaneNotFound"
	ControlPlaneCloudProviderNotSet                         ErrorCode = "ControlPlaneCloudProviderNotSet"
	ControlPlaneNotAvailable                                ErrorCode = "ControlPlaneNotAvailable"
	ControlPlaneVersionImagesError                          ErrorCode = "ControlPlaneVersionImagesError"
	GetControlPlaneFailed                                   ErrorCode = "GetControlPlaneFailed"
	PostOperationError                                      ErrorCode = "PostOperationError"
	PostOperationNilResponse                                ErrorCode = "PostOperationNilResponse"
	RegionOverConstrained                                   ErrorCode = "RegionOverConstrained"
	DeleteResourceGroupInternalOperationError               ErrorCode = "DeleteResourceGroupInternalOperationError"
	DeleteDNSZoneError                                      ErrorCode = "DeleteDNSZoneError"
	CreateDNSZoneError                                      ErrorCode = "CreateDNSZoneError"
	DeleteDNSZoneSegmentError                               ErrorCode = "DeleteDNSZoneSegmentError"
	AddCNAMEDNSZoneError                                    ErrorCode = "AddCNAMEDNSZoneError"
	DeleteDNSRecordError                                    ErrorCode = "DeleteDNSRecordError"
	PutDNSRecordError                                       ErrorCode = "PutDNSRecordError"
	GetPrivateDNSRecordError                                ErrorCode = "GetPrivateDNSRecordError"
	PrivateDNSRecordMetadataNotMatchError                   ErrorCode = "PrivateDNSRecordMetadataNotMatchError"
	PrivateDNSUnauthorizedClientError                       ErrorCode = "unauthorized_client"
	PrivateDNSLinkedAuthorizationFailedError                ErrorCode = "LinkedAuthorizationFailed"
	DeleteResourceGroupError                                ErrorCode = "DeleteResourceGroupError"
	PutResourceGroupError                                   ErrorCode = "PutResourceGroupError"
	PatchResourceGroupError                                 ErrorCode = "PatchResourceGroupError"
	ResourceGroupRequestInternalError                       ErrorCode = "ResourceGroupRequestInternalError"
	ReconcileResourceGroupError                             ErrorCode = "ReconcileResourceGroupError"
	ReconcileResourceProviderError                          ErrorCode = "ReconcileResourceProviderError"
	SearchServicePrincipalError                             ErrorCode = "SearchServicePrincipalError"
	SPClientSecretTooLongError                              ErrorCode = "SPClientSecretTooLongError"
	DNSZoneNotSet                                           ErrorCode = "DNSZoneNotSet"
	CreateManagedAppError                                   ErrorCode = "CreateManagedAppError"
	CreateVirtualNetworkPeeringError                        ErrorCode = "VirtualNetworkPeeringError"
	DelegateSubnetError                                     ErrorCode = "DelegateSubnetError"
	UnDelegateSubnetError                                   ErrorCode = "UnDelegateSubnetError"
	InvalidPeerVnetID                                       ErrorCode = "InvalidPeerVnetID"
	InvalidPeerVnetSubscriptionID                           ErrorCode = "InvalidPeerVnetSubscriptionID"
	TemplateDeploymentTimedoutError                         ErrorCode = "TemplateDeploymentTimedoutError"
	EntityStoreOperationError                               ErrorCode = "EntityStoreOperationError"
	PrivateClusterPublicFQDNUnexpectedEmptyError            ErrorCode = "PrivateClusterPublicFQDNUnexpectedEmptyError"
	BackfillOperationError                                  ErrorCode = "BackfillOperationError"
	AgentPoolOperationError                                 ErrorCode = "AgentPoolOperationError"
	FQDNNotAvailable                                        ErrorCode = "FQDNNotAvailable"
	FQDNInvalid                                             ErrorCode = "FQDNInvalid"
	ResourceGroupLocked                                     ErrorCode = "ResourceGroupLocked"
	ResourceLocked                                          ErrorCode = "ResourceLocked"
	CheckResourceGroupLevelLocksError                       ErrorCode = "CheckResourceGroupLevelLocksError"
	CreateClusterKVError                                    ErrorCode = "CreateClusterKVError"
	ValidationError                                         ErrorCode = "ValidationError"
	TemplateValidationError                                 ErrorCode = "TemplateValidationError"
	InvalidManagedAppResourceGroupID                        ErrorCode = "InvalidManagedAppResourceGroupID"
	CreateResourceGroupError                                ErrorCode = "CreateResourceGroupError"
	CreateRoleDefinitionFailed                              ErrorCode = "CreateRoleDefinitionFailed"
	UpdateNotAllowed                                        ErrorCode = "UpdateNotAllowed"
	AdminUpdateNotAllowed                                   ErrorCode = "AdminUpdateNotAllowed"
	AdminUpdateNonExistentCluster                           ErrorCode = "AdminUpdateNonExistentCluster"
	PluginCreationFailed                                    ErrorCode = "PluginCreationFailed"
	PluginContextCreationFailed                             ErrorCode = "PluginContextCreationFailed"
	AgentPoolAPIsNotSupported                               ErrorCode = "AgentPoolAPIsNotSupported"
	PreviewFeatureNotRegistered                             ErrorCode = "PreviewFeatureNotRegistered"
	WindowsAgentPoolNameTooLong                             ErrorCode = "WindowsAgentPoolNameTooLong"
	WindowsAgentPoolAPIsNotSupported                        ErrorCode = "WindowsAgentPoolAPIsNotSupported"
	WindowsProfileLicenseTypeInvalid                        ErrorCode = "WindowsProfileLicenseTypeInvalid"
	WindowsProfileMissing                                   ErrorCode = "WindowsProfileMissing"
	WindowsProfileUsernameInvalid                           ErrorCode = "WindowsProfileUsernameInvalid"
	WindowsProfilePasswordInvalid                           ErrorCode = "WindowsProfilePasswordInvalid"
	InvalidHourSlot                                         ErrorCode = "InvalidHourSlot"
	NeedAtLeastOneHourPerWeekForUpdate                      ErrorCode = "NeedAtLeastOneHourPerWeekForUpdate"
	TooManyTimeInWeek                                       ErrorCode = "TooManyTimeInWeek"
	TooManyNotAllowedTime                                   ErrorCode = "TooManyNotAllowedTime"
	TooManyHourSlots                                        ErrorCode = "TooManyHourSlots"
	TimeInWeekIsNil                                         ErrorCode = "TimeInWeekIsNil"
	NotAllowedTimeIsNil                                     ErrorCode = "NotAllowedTimeIsNil"
	InvalidStartEndTime                                     ErrorCode = "InvalidStartEndTime"
	AzureCNIOnlyForWindows                                  ErrorCode = "AzureCNIOnlyForWindows"
	AgentPoolMaxCountExceeded                               ErrorCode = "AgentPoolMaxCountExceeded"
	MaxAgentPoolCountReached                                ErrorCode = "MaxAgentPoolCountReached"
	NotAllAgentPoolOrchestratorVersionSpecifiedAndUnchanged ErrorCode = "NotAllAgentPoolOrchestratorVersionSpecifiedAndUnchanged"
	NodePoolMcVersionIncompatible                           ErrorCode = "NodePoolMcVersionIncompatible"
	GetAgentPoolUpgradeProfileError                         ErrorCode = "GetAgentPoolUpgradeProfileError"
	GetLatestNodeImageVersionError                          ErrorCode = "GetLatestNodeImageVersionError"
	GetSIGAzureCloudSpecConfigError                         ErrorCode = "GetSIGAzureCloudSpecConfigError"
	AgentPoolUpgradeVersionNotAllowed                       ErrorCode = "AgentPoolUpgradeVersionNotAllowed"
	AgentPoolK8sVersionNotSupported                         ErrorCode = "AgentPoolK8sVersionNotSupported"
	VMSSPlatformFaultDomainCountZeroInvalid                 ErrorCode = "VMSSPlatformFaultDomainCountZeroInvalid"
	AgentPoolNotFound                                       ErrorCode = "AgentPoolNotFound"
	SpotPoolUpgradeNotAllowedDuringClusterUpgradeAll        ErrorCode = "SpotPoolUpgradeNotAllowedDuringClusterUpgradeAll"
	VMSizeNotSupportedForWindowsAgentPool                   ErrorCode = "VMSizeNotSupportedForWindowsAgentPool"
	SystemPoolHasRestrictedTaint                            ErrorCode = "SystemPoolHasRestrictedTaint"
	MustDefineAtLeastOneSystemPool                          ErrorCode = "MustDefineAtLeastOneSystemPool"
	LowPriPoolDeprecated                                    ErrorCode = "LowPriPoolDeprecated"
	AutoScalingNotEnabled                                   ErrorCode = "AutoScalingNotEnabled"
	MixedTypeNotSupported                                   ErrorCode = "MixedTypeNotSupported"
	ObjectCloneError                                        ErrorCode = "ObjectCloneError"
	CCPProxySvcGetError                                     ErrorCode = "CCPProxySvcGetError"
	CCPProxySvcPendingError                                 ErrorCode = "CCPProxySvcPendingError"
	CCPAPIGatewaySvcGetError                                ErrorCode = "CCPAPIGatewaySvcGetError"
	CCPAPIGatewaySvcPendingError                            ErrorCode = "CCPAPIGatewaySvcPendingError"
	FailedToConvertToVersionedEntity                        ErrorCode = "FailedToConvertToVersionedEntity"
	MustNotBeTheSystemPool                                  ErrorCode = "MustNotBeTheSystemPool"
	ScopeLocked                                             ErrorCode = "ScopeLocked"
	FailedToRemoveCertificateFromKubeconfig                 ErrorCode = "FailedToRemoveCertificateFromKubeconfig"
	// RetryNeeded indicates that this operation needs to be retried
	RetryNeeded ErrorCode = "RetryNeeded"
	// Canceled indicates the operation is Canceled
	Canceled ErrorCode = "Canceled"
	// ScaleDownInternalError is from Microsoft.WindowsAzure.ContainerService.API.AcsErrorCode
	ScaleDownInternalError ErrorCode = "ScaleDownInternalError"

	// Custom error codes specific to the underlay

	// PreconditionCheckTimeOut indicates that the check for all preconditions
	// before provisioning a hosted control plane timed out
	PreconditionCheckTimeOut ErrorCode = "PreconditionCheckTimeOut"
	// UpgradeFailed indicates that an upgrade of a hosted control plane has
	// failed without a specific indication as to why
	UpgradeFailed ErrorCode = "UpgradeFailed"
	// ScaleError indicates that there was a failure in scaling customer nodes
	ScaleError ErrorCode = "ScaleError"
	// CreateRoleAssignmentError indicates that we could not assign a resource
	// to a service principal
	CreateRoleAssignmentError ErrorCode = "CreateRoleAssignmentError"
	// CreateRoleAssignmentOnMSIError indicates that we could not assign a resource
	// to a managed identity.
	CreateRoleAssignmentOnMSIError ErrorCode = "CreateRoleAssignmentOnMSIError"
	// ServicePrincipalNotFound indicates that the Service Principal either does
	// not exist or we were not able to find it
	ServicePrincipalNotFound ErrorCode = "ServicePrincipalNotFound"
	// ClusterResourceGroupNotFound indicates that the Resource Group for the
	// cluster either does not exist or we were not able to find it
	ClusterResourceGroupNotFound ErrorCode = "ClusterResourceGroupNotFound"
	// NodeResourceGroupAlreadyExists indicates that the Resource Group for the cluster
	// already exists
	NodeResourceGroupAlreadyExists ErrorCode = "NodeResourceGroupAlreadyExists"
	// NodeResourceGroupSameAsResourceGroup indicates that the NodeResourceGroup is same as
	// ResourceGroup
	NodeResourceGroupSameAsResourceGroup ErrorCode = "NodeResourceGroupSameAsResourceGroup"
	// NodeResourceGroupManagedByAnotherResource indicates that the NodeResourceGroup is managed by
	// another resource
	NodeResourceGroupManagedByAnotherResource ErrorCode = "NodeResourceGroupManagedByAnotherResource"
	// KubeConfigError indicates that an error occurred when trying to generate
	// or fetch a kubeconfig for access to a managed cluster
	KubeConfigError ErrorCode = "KubeConfigError"
	// KubeConfigExecPluginCoversionError indicates that an error occurred when trying to convert
	// kubeconfig for exec plugin
	KubeConfigExecPluginCoversionError ErrorCode = "KubeConfigExecPluginCoversionError"
	// RegionNotSupported indicates that the given region does not support
	// managed clusters
	RegionNotSupported ErrorCode = "RegionNotSupported"
	// ControlPlaneAddOnsNotReady indicates that the Kubernetes cluster addons
	// in the hosted control plane are not in a ready state
	ControlPlaneAddOnsNotReady ErrorCode = "ControlPlaneAddOnsNotReady"
	// ControlPlaneAddOnsValidationError indicates that service was not able to
	// query add-on pods state
	ControlPlaneAddOnsValidationError ErrorCode = "ControlPlaneAddOnsValidationError"
	// ControlPlaneAPIServerNotReady indicates that we could not connect to customer's API server
	ControlPlaneAPIServerNotReady ErrorCode = "ControlPlaneAPIServerNotReady"
	// NodesNotReady indicates that nodes in a managed cluster are not in a
	// ready state as far as `kubectl get node`` is concerned
	NodesNotReady ErrorCode = "NodesNotReady"
	// PodDrainFailure means that we failed to evict a pod from a node before scaledown/reimage
	// Often this means there's a pod disruption budget that is blocking eviction
	PodDrainFailure ErrorCode = "PodDrainFailure"
	// UncordonNodeFailure indicates that we have failed to uncordon nodes
	UncordonNodeFailure ErrorCode = "UncordonNodeFailure"
	// ControlPlaneProvisioningTimeout indicates that a timeout was reached when
	// attempting to provision a hosted control plane
	ControlPlaneProvisioningTimeout ErrorCode = "ControlPlaneProvisioningTimeout"
	// AgentCountNotMatch indicates the number of VM created does not
	// match the number of specified agent count
	AgentCountNotMatch ErrorCode = "AgentCountNotMatch"
	// SSHKeyGeneratingFailed indicates a failure when attempting to generate an
	// SSH key for a managed node
	SSHKeyGeneratingFailed ErrorCode = "SSHKeyGeneratingFailed"

	// Error codes returned by a hosted control plane service

	// UnderlayDecode indicates that an underlay json string cannot be decoded
	UnderlayDecode ErrorCode = "UnderlayCannotBeDecoded"
	// UnderlayNotFound indicates that an underlay either does not exist or we
	// were not able to find it
	UnderlayNotFound ErrorCode = "UnderlayNotFound"
	// UnderlayNotSet indicates that an underlay ID was not given
	UnderlayNotSet ErrorCode = "UnderlayNotSet"
	// UnderlaysOverConstrained indicates that an operation is unable to be
	// completed due to high load on the underlays
	UnderlaysOverConstrained ErrorCode = "UnderlaysOverConstrained"
	// AZUnderlaysOverConstrained indicates that an operation is unable to be
	// completed due to high load on the az underlays
	AZUnderlaysOverConstrained  ErrorCode = "AZUnderlaysOverConstrained"
	Microsoft365OverConstrained ErrorCode = "Microsoft365OverConstrained"
	// GRPC subcode for get underlay not found
	GetUnderlayV1NotFound ErrorCode = "GetUnderlayV1_NotFound"
	// UnexpectedUnderlayCount indicates that an unexpected number of underlays
	// was found
	UnexpectedUnderlayCount ErrorCode = "UnexpectedUnderlayCount"
	// ControlPlaneNotReady indicates that the hosted control plane components are
	// not in a ready state
	ControlPlaneNotReady ErrorCode = "ControlPlaneNotReady"
	// TillerRecordNotFound indicates that the Tiller release record for a hosted
	// control plane either does not exist or we were not able to find it
	TillerRecordNotFound ErrorCode = "TillerRecordNotFound"
	// RegionCreatesDisabled indicates that creation of new hosted control planes
	// is currently disabled in the given region
	RegionCreatesDisabled ErrorCode = "RegionCreatesDisabled"
	// RegionUpdatesDisabled indicates that updates of hosted control planes
	// is currently disabled in the given region
	RegionUpdatesDisabled ErrorCode = "RegionUpdatesDisabled"
	// AADNotSupported error code is returned when AAD capable AKS cluster creation is turned off
	AADNotSupported ErrorCode = "AADNotSupported"
	// GetUpgradeVersionsError error code is returned when getting the list of upgrades fails
	GetUpgradeVersionsError ErrorCode = "GetUpgradeVersionsError"

	// Geneva Action related error codes

	// GenevaActionInternalError is a generic error indicating an issue with Geneva
	GenevaActionInternalError ErrorCode = "GenevaActionInternalError"

	// Error codes related to Addon

	// AddonInvalid is a generatic error indicating that a configured addon
	// definition is invalid
	AddonInvalid ErrorCode = "AddonInvalid"
	// AddonAlreadyEnabledAsExtensionAddon indicates the addon is already enabled as an extension addon
	AddonAlreadyEnabledAsExtensionAddon ErrorCode = "AddonAlreadyEnabledAsExtensionAddon"
	// HTTPApplicationRoutingAddonConfigInvalid indicates that the application
	// router configuration is invalid
	HTTPApplicationRoutingAddonConfigInvalid ErrorCode = "HTTPApplicationRoutingConfigInvalid"
	// LogAnalyticsWorkspaceNotSet indicates the Log Analytics workspace resource ID is not set
	LogAnalyticsWorkspaceNotSet ErrorCode = "LogAnalyticsWorkspaceNotSet"
	// OmsagentInvalidAADAuthInput indicates the input of useAADAuth in omsagent addon config is invalid
	OmsagentInvalidAADAuthInput ErrorCode = "OmsagentInvalidAADAuthInput"
	// LogAnalyticsWorkspaceNotSetForAzureDefender indicates the Log Analytics workspace resource ID is not set for Azure Defender addon
	LogAnalyticsWorkspaceNotSetForAzureDefender ErrorCode = "LogAnalyticsWorkspaceNotSetForAzureDefender"
	// GetLogAnalyticsWorkspaceError indicates a failure when attempting to
	// fetch information on an analytics workspace
	GetLogAnalyticsWorkspaceError ErrorCode = "GetLogAnalyticsWorkspaceError"
	// AddContainerInsightsSolutionError indicates a failure when attempting to add
	// ContainerInsights solution to a log analytics workspace
	AddContainerInsightsSolutionError                     ErrorCode = "AddContainerInsightsSolutionError"
	LogAnalyticsWorkspaceKeyAccessDeniedCrossSubscription ErrorCode = "LogAnalyticsWorkspaceKeyAccessDeniedCrossSubscription"
	LogAnalyticsWorkspaceKeyAccessDenied                  ErrorCode = "LogAnalyticsWorkspaceKeyAccessDenied"
	LogAnalyticsWorkspaceNotFound                         ErrorCode = "LogAnalyticsWorkspaceNotFound"
	LogAnalyticsWorkspaceNotFoundForAzureDefender         ErrorCode = "LogAnalyticsWorkspaceNotFoundForAzureDefender"

	// IngressAppGwAddonConfigInvalid indicates that the application gateway ingress controller
	// configuration is invalid
	IngressAppGwAddonConfigInvalid                                                   ErrorCode = "IngressAppGwAddonConfigInvalid"
	IngressAppGwAddonConfigInvalidApplicationGatewayID                               ErrorCode = "IngressAppGwAddonConfigInvalidApplicationGatewayId"
	IngressAppGwAddonConfigGetApplicationGatewayError                                ErrorCode = "IngressAppGwAddonConfigGetApplicationGatewayError"
	IngressAppGwAddonConfigApplicationGatewayNotFound                                ErrorCode = "IngressAppGwAddonConfigApplicationGatewayNotFound"
	IngressAppGwAddonConfigApplicationGatewayAccessDenied                            ErrorCode = "IngressAppGwAddonConfigApplicationGatewayAccessDenied"
	IngressAppGwAddonConfigUnsupportedApplicationGatewaySku                          ErrorCode = "IngressAppGwAddonConfigUnsupportedApplicationGatewaySku"
	IngressAppGwAddonConfigCannotUseApplicationGatewayInDeletingState                ErrorCode = "IngressAppGwAddonConfigCannotUseApplicationGatewayInDeletingState"
	IngressAppGwAddonConfigProvideEitherApplicationGatewayIDOrSubnet                 ErrorCode = "IngressAppGwAddonConfigProvideEitherApplicationGatewayIDOrSubnet"
	IngressAppGwAddonConfigProvideEitherApplicationGatewayIDOrApplicationGatewayName ErrorCode = "IngressAppGwAddonConfigProvideEitherApplicationGatewayIDOrApplicationGatewayName"
	IngressAppGwAddonConfigProvideEitherSubnetCIDROrSubnetID                         ErrorCode = "IngressAppGwAddonConfigProvideEitherSubnetCIDROrSubnetId"
	IngressAppGwAddonConfigProvideUseSubnetCIDRInsteadOfSubnetPrefix                 ErrorCode = "IngressAppGwAddonConfigProvideUseSubnetCIDRInsteadOfSubnetPrefix"
	IngressAppGwAddonConfigInvalidSubnetCIDR                                         ErrorCode = "IngressAppGwAddonConfigInvalidSubnetCIDR"
	IngressAppGwAddonConfigInvalidSubnetCIDRCannotBeModified                         ErrorCode = "IngressAppGwAddonConfigInvalidSubnetCIDRCannotBeModified"
	IngressAppGwAddonConfigInvalidSubnetCIDRNotContainedWithinVirtualNetwork         ErrorCode = "IngressAppGwAddonConfigInvalidSubnetCIDRNotContainedWithinVirtualNetwork"
	IngressAppGwAddonConfigInvalidSubnetCIDROverlapWithSubnet                        ErrorCode = "IngressAppGwAddonConfigInvalidSubnetCIDROverlapWithSubnet"
	IngressAppGwAddonConfigInvalidSubnetID                                           ErrorCode = "IngressAppGwAddonConfigInvalidSubnetId"
	IngressAppGwAddonConfigCannotUseSubnet                                           ErrorCode = "IngressAppGwAddonConfigCannotUseSubnet"
	IngressAppGwAddonConfigGetSubnetError                                            ErrorCode = "IngressAppGwAddonConfigGetSubnetError"
	IngressAppGwAddonConfigSubnetNotFound                                            ErrorCode = "IngressAppGwAddonConfigSubnetNotFound"
	IngressAppGwAddonConfigSubnetAccessDenied                                        ErrorCode = "IngressAppGwAddonConfigSubnetAccessDenied"
	IngressAppGwAddonConfigSubnetNotInSucceededState                                 ErrorCode = "IngressAppGwAddonConfigSubnetNotInSucceededState"
	IngressAppGwAddonConfigInvalidSharedFlag                                         ErrorCode = "IngressAppGwAddonConfigInvalidSharedFlag"
	IngressAppGwAddonConfigSharingNotAllowedForNewGateway                            ErrorCode = "IngressAppGwAddonConfigSharingNotAllowedForNewGateway"
	IngressAppGwAddonConfigSubnetNotAllowedWhenUsingExistingGateway                  ErrorCode = "IngressAppGwAddonConfigSubnetNotAllowedWhenUsingExistingGateway"
	IngressAppGwAddonDeleteApplicationGatewayError                                   ErrorCode = "IngressAppGwAddonDeleteApplicationGatewayError"

	// Error codes related to Open Service Mesh (OSM) addon

	// OpenServiceMeshAddonFeatureFlagNotEnabled indicates that the Open Service Mesh addon feature flag hasn't been enabled
	OpenServiceMeshAddonFeatureFlagNotEnabled               ErrorCode = "OpenServiceMeshAddonFeatureFlagNotEnabled"
	OpenServiceMeshAddonNotSupportedWithOrchestratorVersion ErrorCode = "OpenServiceMeshAddonNotSupportedWithOrchestratorVersion"

	// Error codes related to Azure Keyvault Secrets Provider addon
	/* #nosec */
	// AzureKeyvaultSecretsProviderAddonFeatureFlagNotEnabled indicates that the Azure Keyvault Secrets Provider addon feature flag hasn't been enabled
	AzureKeyvaultSecretsProviderAddonFeatureFlagNotEnabled ErrorCode = "AzureKeyvaultSecretsProviderAddonFeatureFlagNotEnabled"
	/* #nosec */
	AzureKeyvaultSecretsProviderAddonNotSupportedWithOrchestratorVersion ErrorCode = "AzureKeyvaultSecretsProviderNotSupportedWithOrchestratorVersion"

	DaprExtensionFeatureFlagNotEnabled               ErrorCode = "DaprExtensionFeatureFlagNotEnabled"
	DaprExtensionNotSupportedWithOrchestratorVersion ErrorCode = "DaprExtensionNotSupportedWithOrchestratorVersion"

	// AzureDefenderFeatureFlagNotEnabled indicates that the Azure Defender feature flag hasn't been enabled
	AzureDefenderFeatureFlagNotEnabled ErrorCode = "AzureDefenderFeatureFlagNotEnabled"

	// Error codes related pod identity addon

	// PodIdentityAddonFeatureFlagNotEnabled indicates that the pod identity feature flag hasn't been enabled.
	PodIdentityAddonFeatureFlagNotEnabled ErrorCode = "PodIdentityAddonFeatureFlagNotEnabled"
	// PodIdentityAddonUserAssignedIdentitiesNotAllowedInCreation indicates that the user assigned identity assignments are disallowed in cluster creation.
	PodIdentityAddonUserAssignedIdentitiesNotAllowedInCreation ErrorCode = "PodIdentityAddonUserAssignedIdentitiesNotAllowedInCreation"
	// PodIdentityAddonIdentitiesAssignmentError indicates that the identity assignment failed.
	PodIdentityAddonIdentitiesAssignmentError ErrorCode = "PodIdentityAddonIdentitiesAssignmentError"
	// PodIdentityAddonWithKubeNetNotAllowed indicates that the cluster is not allowed to use pod identity with kubenet network plugin.
	PodIdentityAddonWithKubeNetNotAllowed ErrorCode = "PodIdentityAddonWithKubeNetNotAllowed"
	// PodIdentityAddonWithKubeNetNeedsConsent indicates that the cluster needs customer consent to enable pod identity with kubenet network plugin.
	PodIdentityAddonWithKubeNetNeedsConsent ErrorCode = "PodIdentityAddonWithKubeNetNeedsConsent"
	ReconcilePodIdentityAddonError          ErrorCode = "ReconcilePodIdentityAddonError"

	// Error codes related to custom route table

	// TooManyCustomRouteTables indicates a customer has configured a cluster with more than one route table
	TooManyCustomRouteTables ErrorCode = "TooManyCustomRouteTables"
	// CustomRouteTableParseError indicates that a route table ID couldn't be parsed
	CustomRouteTableParseError ErrorCode = "CustomRouteTableParseError"
	// UpdateCustomRouteTableError indicates a failure to update a route table
	UpdateCustomRouteTableError ErrorCode = "UpdateCustomRouteTableError"
	// CustomRouteTableNotFound indicates a route table could not be found
	CustomRouteTableNotFound ErrorCode = "CustomRouteTableNotFound"
	// GetCustomRouteTableByIDError failed to get route table by ID
	GetCustomRouteTableByIDError ErrorCode = "GetCustomRouteTableByIDError"
	// CustomRouteTableUnsupportedK8SVersion indicates a customer didn't provide the minimum required kubernetes version (1.15.0)
	CustomRouteTableUnsupportedK8SVersion ErrorCode = "CustomRouteTableUnsupportedK8SVersion"
	//SemverOperationFailed indicates that a semver operation (create,check etc.) failed
	SemverOperationFailed ErrorCode = "SemverOperationFailed"
	// SubnetsMissingCustomRouteTableAssociation indicates that not all of the subnet's have a route table association
	SubnetsMissingCustomRouteTableAssociation ErrorCode = "SubnetMissingRouteTableAssociation"
	// AgentPoolsWithEmptySubnetError indicates that one or more agent pools are missing a subnetID
	AgentPoolsWithEmptySubnetError ErrorCode = "AgentPoolsWithEmptySubnetError"
	// CustomRouteTableWithMSINotSupported indicates that customer tried to use custom route table and MSI simultaneously
	CustomRouteTableWithMSINotSupported ErrorCode = "CustomRouteTableWithMSINotSupported"
	// CustomRouteTableCannotCrossSubscription indicates that a custom table is not in the same subscription as the AKS cluster
	CustomRouteTableCannotCrossSubscription ErrorCode = "CustomRouteTableCannotCrossSubscription"
	// CustomRouteTableMissingServicePrincipalPermission indicates that the customer's service principal is missing read/write permission to the route table
	CustomRouteTableMissingServicePrincipalPermission ErrorCode = "CustomRouteTableMissingServicePrincipalPermission"
	// CustomRouteTableConflict indicates a conflict occurred while reading or writing to the custom route table
	CustomRouteTableConflict ErrorCode = "CustomRouteTableConflict"
	// CustomRouteTablePreConditionFailed indicates a pre-condition failed  while reading or writing to the custom route table
	CustomRouteTablePreConditionFailed ErrorCode = "CustomRouteTablePreConditionFailed"
	// CustomRouteTableTooManyRequests indicates that throttling occurred when reading or writing a custom route table
	CustomRouteTableTooManyRequests ErrorCode = "CustomRouteTableTooManyRequests"
	// CustomRouteTableInvalidUpdateAttempt indicates that a customer attempted to update their custom route table which is immutable
	CustomRouteTableInvalidUpdateAttempt ErrorCode = "CustomRouteTableInvalidUpdateAttempt"

	// Error codes related to Custom VNET

	// Error codes related to private dns zone
	CustomPrivateDNSZoneMissingPermissionError ErrorCode = "CustomPrivateDNSZoneMissingPermissionError"

	// ServiceCidrOverlapping indicates that overlay service CIDR block overlaps with underlay service CIDR
	ServiceCidrOverlapping ErrorCode = "ServiceCidrOverlapping"
	// SubnetCidrOverlapping indicates that overlay subnet CIDR block overlaps with underlay service CIDR
	SubnetCidrOverlapping ErrorCode = "SubnetCidrOverlapping"
	// PodCidrOverlapExistingSubnetsCidr indicates that overlay pod CIDR block overlaps with existing subnets' CIDR
	PodCidrOverlapExistingSubnetsCidr ErrorCode = "PodCidrOverlapExistingSubnetsCidr"
	// VnetPeerCidrOverlapping indicates that OSA cluster vnet overlaps with peerVnetID
	VnetPeerCidrOverlapping ErrorCode = "VnetPeerCidrOverlapping"
	// ServiceCidrOverlapExistingSubnetsCidr indicates a new specified service CIDR is conflicted with an existing subnets' CIDR
	ServiceCidrOverlapExistingSubnetsCidr ErrorCode = "ServiceCidrOverlapExistingSubnetsCidr"
	// SubnetCidrOverlapExistingSubnetsCidr indicates a new specified subnet CIDR is conflicted with an existing subnets' CIDR
	SubnetCidrOverlapExistingSubnetsCidr ErrorCode = "SubnetCidrOverlapExistingSubnetsCidr"

	// SubnetCidrOverlapServiceCidr indicates that overlay subnet CIDR block overlaps with overlay service CIDR
	SubnetCidrOverlapServiceCidr ErrorCode = "SubnetCidrOverlapServiceCidr"
	// PodCidrOverlapServiceCidr indicates that overlay pod CIDR block overlaps with overlay service CIDR
	PodCidrOverlapServiceCidr ErrorCode = "PodCidrOverlapServiceCidr"

	// DockerBridgeInSubnetCidr indicates that docker bridge IP is in overlay subnet CIDR block
	DockerBridgeInSubnetCidr ErrorCode = "DockerBridgeInSubnetCidr"
	// DockerBridgeInPodCidr indicates that docker bridge IP is in overlay pod CIDR block
	DockerBridgeInPodCidr ErrorCode = "DockerBridgeInPodCidr"
	// DockerBridgeInServiceCidr indicates that docker bridge IP is in overlay serviceCIDR block
	DockerBridgeInServiceCidr ErrorCode = "DockerBridgeInServiceCidr"
	// DockerBridgeCidrInvalid indicates that docker bridge IP is not valid
	DockerBridgeCidrInvalid ErrorCode = "DockerBridgeCidrInvalid"

	// SubnetIsDelegated indicates that the subnet is delegated to other services so cannot be used for the agentpool
	SubnetIsDelegated ErrorCode = "SubnetIsDelegated"

	// InsufficientSubnetSize indicates that the pre-allocated IPs exceeds the size of subnet CIDR block
	InsufficientSubnetSize ErrorCode = "InsufficientSubnetSize"
	// InsufficientPodCidr indicates that the pre-allocated IPs exceeds the size of pod CIDR block
	InsufficientPodCidr ErrorCode = "InsufficientPodCidr"

	// MultipleAddressPrefixesSubnetNotSupported indicates that we don't support multiple address prefixes on AKS used subnets
	MultipleAddressPrefixesSubnetNotSupported ErrorCode = "MultipleAddressPrefixesSubnetNotSupported"
	// IPv6NotSupported indicates that CIDR is IPv6 which is unsupported
	IPv6NotSupported ErrorCode = "IPv6NotSupported"
	// CidrStringParseError indicates that an invalid CIDR block string was given
	CidrStringParseError ErrorCode = "CidrStringParseError"
	// NetworkPolicyNotSupported indicates that network policy is not supported
	NetworkPolicyNotSupported ErrorCode = "NetworkPolicyNotSupported"
	// NetworkPolicyAzureNotSupportedForKubenet indicates that network policy azure is not supported for kubenet plugin.
	NetworkPolicyAzureNotSupportedForKubenet ErrorCode = "NetworkPolicyAzureNotSupportedForKubenet"
	// NetworkPolicyNotSupportedForNone indicates that no network policy is allowed for network plugin none.
	NetworkPolicyNotSupportedForNone ErrorCode = "NetworkPolicyNotSupportedForNone"
	// NetworkModeNotSupported indicates that network mode is not supported
	NetworkModeNotSupported ErrorCode = "NetworkModeNotSupported"
	// NetworkModeInvalid indicates that network mode is invalid
	NetworkModeInvalid ErrorCode = "NetworkModeInvalid"
	// GetVnetSubnetCidrError indicates that an error in getting vnet and subnet Cidr
	GetVnetSubnetCidrError ErrorCode = "GetVnetSubnetCidrError"
	// GetVnetCidrError indicates that an error in getting vnet cidr
	GetVnetCidrError ErrorCode = "GetVnetCidrError"
	// GetSubnetCidrError indicates that an error in getting subnet cidr
	GetSubnetCidrError ErrorCode = "GetSubnetCidrError"
	// GetAllSubnetsError indicates there is an error in getting all subnets from VNet
	GetAllSubnetsError ErrorCode = "GetAllSubnetsError"
	// AzureCNISubnetNotDefined indicates that a subnet is not defined when Azure CNI network plugin is used
	AzureCNISubnetNotDefined ErrorCode = "AzureCNISubnetNotDefined"

	// UnsupportedOutboundType indicates that the outbound type has been set to a value that is not acceptable
	UnsupportedOutboundType ErrorCode = "UnsupportedOutboundType"
	// UpdatingOutboundTypeNotAllowed indicates that the outbound type has been set to a value that is not acceptable
	UpdatingOutboundTypeNotAllowed ErrorCode = "UpdatingOutboundTypeNotAllowed"
	// InvalidUserDefinedRoutingWithLoadBalancerProfile indicates that you cannot have both a UDR and a load balancer profile for outbound connectivity
	InvalidUserDefinedRoutingWithLoadBalancerProfile ErrorCode = "InvalidUserDefinedRoutingWithLoadBalancerProfile"

	// InvalidLoadBalancerSku indicates that load balancer SKU is not a valid value
	InvalidLoadBalancerSku ErrorCode = "InvalidLoadBalancerSku"
	// UnexpectedLoadBalancerSkuForCurrentOutboundConfiguration indicates that load balancer SKU must be Standard when load balancer profile is provided or outbound type is user defined routing or NAT gateway
	UnexpectedLoadBalancerSkuForCurrentOutboundConfiguration ErrorCode = "UnexpectedLoadBalancerSkuForCurrentOutboundConfiguration"
	// InvalidLoadBalancerProfile indicates that load balancer profile is not a valid value
	InvalidLoadBalancerProfile ErrorCode = "InvalidLoadBalancerProfile"
	// InvalidLoadBalancerProfileIPProvider indicates that load balancer profile has the wrong IP provider
	InvalidLoadBalancerProfileIPProvider ErrorCode = "InvalidLoadBalancerProfileIPProvider"
	// InvalidLoadBalancerProfileIPPrefixProvider indicates that load balancer profile has the wrong IPPrefix provider
	InvalidLoadBalancerProfileIPPrefixProvider ErrorCode = "InvalidLoadBalancerProfileIPPrefixProvider"
	// InvalidLoadBalancerProfileIdleTimeoutInMinutes indicates that load balancer profile idle timeout is not a valid value
	InvalidLoadBalancerProfileIdleTimeoutInMinutes ErrorCode = "InvalidLoadBalancerProfileIdleTimeoutInMinutes"
	// InvalidLoadBalancerProfileAllocatedOutboundPorts indicates that load balancer profile allocated ports is not a valid value
	InvalidLoadBalancerProfileAllocatedOutboundPorts ErrorCode = "InvalidLoadBalancerProfileAllocatedOutboundPorts"
	// LoadBalancerProfileAllocatedOutboundPortsCannotBeNegative indicates that load balancer profile allocated ports cannot be a negative value
	LoadBalancerProfileAllocatedOutboundPortsCannotBeNegative ErrorCode = "LoadBalancerProfileAllocatedOutboundPortsCannotBeNegative"
	// LoadBalancerProfileAllocatedOutboundPortsMustBeMultipleOfEight indicates that load balancer profile allocated ports must be a multiple of 8
	LoadBalancerProfileAllocatedOutboundPortsMustBeMultipleOfEight ErrorCode = "LoadBalancerProfileAllocatedOutboundPortsMustBeMultipleOfEight"
	// InsufficientOutboundPorts indicates that the outbound ports is insufficient given the requested allocated outbound ports in load balancer profile
	InsufficientOutboundPorts ErrorCode = "InsufficientOutboundPorts"
	// InvalidLoadBalancerProfileManagedOutboundIPCount indicates that managed outbound IP count in load balancer profile is not a valid value
	InvalidLoadBalancerProfileManagedOutboundIPCount ErrorCode = "InvalidLoadBalancerProfileManagedOutboundIPCount"
	// ManagedOutboundIPCountExceedsQuotaLimit indicates that managed outbound IP count in load balancer profile or NAT gateway profile exceeds quota limit
	ManagedOutboundIPCountExceedsQuotaLimit ErrorCode = "ManagedOutboundIPCountExceedsQuotaLimit"
	// InvalidLoadBalancerProfileOutboundIPPrefixes indicates that outbound IP prefixes in load balancer profile is not a valid value
	InvalidLoadBalancerProfileOutboundIPPrefixes ErrorCode = "InvalidLoadBalancerProfileOutboundIPPrefixes"
	// InvalidLoadBalancerProfileOutboundIPs indicates that outbound IPs in load balancer profile is not a valid value
	InvalidLoadBalancerProfileOutboundIPs ErrorCode = "InvalidLoadBalancerProfileOutboundIPs"
	// LoadBalancerProfileOutboundIPPrefixNotExist indicates that the at least one of the outbound IP prefixes in load balancer profile does not exist
	LoadBalancerProfileOutboundIPPrefixNotExist ErrorCode = "LoadBalancerProfileOutboundIPPrefixNotExist"
	// LoadBalancerProfileOutboundIPNotExist indicates that at least one of the outbound IPs in load balancer profile does not exist
	LoadBalancerProfileOutboundIPNotExist ErrorCode = "LoadBalancerProfileOutboundIPNotExist"
	// IPParseError indicates that an invalid IP string was given
	IPParseError ErrorCode = "IPParseError"
	// IPPrefixParseError indicates that an invalid IPPrefix string was given
	IPPrefixParseError ErrorCode = "IPPrefixParseError"
	// LoadBalancerProfileOutboundIPPrefixProvisioningNotSucceeded indicates that at least one of the outbound IP prefixes in load balancer profile does not have succeeded provisioning state
	LoadBalancerProfileOutboundIPPrefixProvisioningNotSucceeded ErrorCode = "LoadBalancerProfileOutboundIPPrefixProvisioningNotSucceeded"
	// LoadBalancerProfileOutboundIPProvisioningNotSucceeded indicates that at least one of the outbound IPs in load balancer profile does not have succeeded provisioning state
	LoadBalancerProfileOutboundIPProvisioningNotSucceeded ErrorCode = "LoadBalancerProfileOutboundIPProvisioningNotSucceeded"
	// LoadBalancerProfileOutboundIPPrefixNotInTheSameSubscription indicates that at least one of the outbound IP prefixes in load balancer profile is not in the same subscription as cluster
	LoadBalancerProfileOutboundIPPrefixNotInTheSameSubscription ErrorCode = "LoadBalancerProfileOutboundIPPrefixNotInTheSameSubscription"
	// LoadBalancerProfileOutboundIPPrefixNotInTheSameRegion indicates that at least one of the outbound IP prefixes in load balancer profile is not in the same region as cluster
	LoadBalancerProfileOutboundIPPrefixNotInTheSameRegion ErrorCode = "LoadBalancerProfileOutboundIPPrefixNotInTheSameRegion"
	// LoadBalancerProfileOutboundIPNotInTheSameRegion indicates that at least one of the outbound IPs in load balancer profile is not in the same region as cluster
	LoadBalancerProfileOutboundIPNotInTheSameRegion ErrorCode = "LoadBalancerProfileOutboundIPNotInTheSameRegion"
	// LoadBalancerProfileOutboundIPNotStandardSKU indicates that at least one of the outbound IPs in load balancer profile is not in standard SKU
	LoadBalancerProfileOutboundIPNotStandardSKU ErrorCode = "LoadBalancerProfileOutboundIPNotStandardSKU"
	// LoadBalancerProfileOutboundIPHasManagedOutboundIPTag indicates that at least one of the outbound IPs in load balancer profile has
	// managed outbound IP tag, so it cannot be used as customer provided outbound IP
	LoadBalancerProfileOutboundIPHasManagedOutboundIPTag ErrorCode = "LoadBalancerProfileOutboundIPHasManagedOutboundIPTag"
	// LoadBalancerProfileOutboundIPAttachedToAnotherIPConfiguration indicates that at least one of the outbound IPs in load balancer profile is
	// already attached to another IP configuration, so it cannot be used as customer provided outbound IP
	LoadBalancerProfileOutboundIPAttachedToAnotherIPConfiguration ErrorCode = "LoadBalancerProfileOutboundIPAttachedToAnotherIPConfiguration"
	// AzureSLBNotSupportedByOrchestrator indicates that Azure SLB is not supported by the selected orchestrator version
	AzureSLBNotSupportedByOrchestrator ErrorCode = "AzureSLBNotSupportedByOrchestrator"

	// PrivateClusterNotSupportedWithBasicLoadBalancer indicates that the setting of enablePrivateCluster is not supported
	PrivateClusterNotSupportedWithBasicLoadBalancer ErrorCode = "PrivateClusterNotSupportedWithBasicLoadBalancer"

	// InvalidNATGatewayProfileManagedOutboundIPCount indicates that managed outbound IP count in the NAT gateway profile is not in the valid range
	InvalidNATGatewayProfileManagedOutboundIPCount ErrorCode = "InvalidNATGatewayProfileManagedOutboundIPCount"
	// InvalidNATGatewayProfileIdleTimeoutInMinutes indicates that idleTimeoutInMinutes in the NAT gateway profile is not in the valid range
	InvalidNATGatewayProfileIdleTimeoutInMinutes ErrorCode = "InvalidNATGatewayProfileIdleTimeoutInMinutes"
	// ManagedNATGatewayWithCustomVNetNotAllowed indicates that outbound type is managedNATGateway but agent pool is using custom VNet which is not allowed
	ManagedNATGatewayWithCustomVNetNotAllowed ErrorCode = "ManagedNATGatewayWithCustomVNetNotAllowed"
	// UserAssignedNATGatewayWithManagedVNetNotAllowed indicates that outbound type is userAssignedNATGateway but agent pool is not using custom VNet
	UserAssignedNATGatewayWithManagedVNetNotAllowed ErrorCode = "UserAssignedNATGatewayWithManagedVNetNotAllowed"
	// SubnetNotAssociatedWithNATGateway indicates that the subnet is not associated with a NAT gateway
	SubnetNotAssociatedWithNATGateway ErrorCode = "SubnetNotAssociatedWithNATGateway"
	// NATGatewayNotAssociatedWithPublicIPOrIPPrefix indicates that the NAT gateway is not associated with any public IP or IP prefix
	NATGatewayNotAssociatedWithPublicIPOrIPPrefix ErrorCode = "NATGatewayNotAssociatedWithPublicIPOrIPPrefix"

	// ExistingRouteTableInSubnetNotSupported indicates there is already a route table in subnet
	ExistingRouteTableInSubnetNotSupported ErrorCode = "ExistingRouteTableInSubnetNotSupported"
	// ExistingRouteTableNotAssociatedWithSubnet indicates that there is not a route table properly associated with the subnet
	ExistingRouteTableNotAssociatedWithSubnet ErrorCode = "ExistingRouteTableNotAssociatedWithSubnet"
	// GetNetworkUsagesError indicates an error when getting network usages
	GetNetworkUsagesError ErrorCode = "GetNetworkUsagesError"
	// GetRouteTableError indicates an error when getting route table
	GetRouteTableError ErrorCode = "GetRouteTableError"
	// GetVnetError indicates an error in getting a VNet object
	GetVnetError ErrorCode = "GetVnetError"
	// ListVnetError indicates an error in list all vnets
	ListVnetError ErrorCode = "ListVnetError"
	// SetResourceOwnershipError indicates an error when setting resource ownership
	SetResourceOwnershipError ErrorCode = "SetResourceOwnershipError"
	// GetSubnetError indicates an error when getting subnet
	GetSubnetError ErrorCode = "GetSubnetError"
	// GetSubnetClientError indicates an error when getting a subnet client
	GetSubnetClientError ErrorCode = "GetSubnetClientError"
	// SubnetNotFound indicates an error when we could not find the subnet in the vnet
	SubnetNotFound ErrorCode = "SubnetNotFound"
	// UDRWithNodePublicIPNotAllowed indicates UDR can not be combined with node public IP feature
	UDRWithNodePublicIPNotAllowed ErrorCode = "UDRWithNodePublicIPNotAllowed"
	// GetSubnetClientFailed unable to retrieve subnet client
	GetSubnetClientFailed ErrorCode = "GetSubnetClientFailed"
	// GetRouteTableClientFailed unable to retrieve route table client
	GetRouteTableClientFailed ErrorCode = "GetRouteTableClientFailed"
	// RouteTableMissingDefaultRouteError indicates that the default route of 0.0.0.0/0 that is required for UDR is missing
	RouteTableMissingDefaultRouteError ErrorCode = "RouteTableMissingDefaultRouteError"
	// RouteTableInvalidNextHop indicates that the next hop for the route table is not set to an NVA or gateway
	RouteTableInvalidNextHop ErrorCode = "RouteTableInvalidNextHop"
	// VnetCIDRInvalid indicates an error when the CIDR of the vnet is invalid
	VnetCIDRInvalid ErrorCode = "VnetCIDRInvalid"
	// AssociateRouteTableToSubnetError indicates an error when associating route table to subnet
	AssociateRouteTableToSubnetError ErrorCode = "AssociateRouteTableToSubnetError"
	// DissociateRouteTableOrNetworkSecurityGroupFromSubnetError indicates an error when dissociate route table or NSG from subnet
	DissociateRouteTableOrNetworkSecurityGroupFromSubnetError ErrorCode = "DissociateRouteTableOrNetworkSecurityGroupFromSubnetError"
	// InvalidSubnetSourceID indicates the subnet sourceID is invalid
	InvalidSubnetSourceID ErrorCode = "InvalidSubnetSourceID"
	// InvalidVNetSourceID indicates the VNet sourceID is invalid
	InvalidVNetSourceID ErrorCode = "InvalidVNetSourceID"
	// EmptySubnetError indicates the subnet is empty
	EmptySubnetError ErrorCode = "EmptySubnetError"
	// ListAllSubnetsError indicates failure to list all subnets of one VNet
	ListAllSubnetsError ErrorCode = "ListAllSubnetsError"
	// GetNATGatewayError indicates an error when getting a NAT gateway
	GetNATGatewayError ErrorCode = "GetNATGatewayError"
	// InvalidNATGatewayResourceID indicates the NAT gateway resource ID is invalid
	InvalidNATGatewayResourceID ErrorCode = "InvalidNATGatewayResourceID"
	// NATGatewayNotFound indicates that the NAT gateway is not found
	NATGatewayNotFound ErrorCode = "NATGatewayNotFound"

	// DnsServiceIpParseError indicates that an invalid DNS Service IP string was given
	DnsServiceIpParseError ErrorCode = "DnsServiceIpParseError"
	// DnsServiceIpOutOfServiceCidr indicates that the DNS Service IP string is not in the range of service CIDR
	DnsServiceIpOutOfServiceCidr ErrorCode = "DnsServiceIpOutOfServiceCidr"
	// DnsServiceIpConflicting indicates that the customers specified DNS IP address is conflicted with predefined API server IP
	DnsServiceIpConflicting ErrorCode = "DnsServiceIpConflicting"
	// DnsBroadcastAddressConflicting indicates that the customers specified DNS IP address is conflicted with the broadcast address
	DnsBroadcastAddressConflicting ErrorCode = "DnsBroadcastAddressConflicting"

	// InvalidIPRangeCIDR indicates the CIDR is not a correct IP address range CIDR, but a single IP CIDR."
	InvalidIPRangeCIDR ErrorCode = "InvalidIPRangeCIDR"

	// UnexpectedIPVersionCIDR indicates a CIDR was parsed successfully but the expected IP version did not match the CIDR IP version
	UnexpectedIPVersionCIDR ErrorCode = "UnexpectedIPVersionCIDR"

	// SingleIPAddressCIDR indicates this CIDR is a valid CIDR but contains only a single IP address.
	SingleIPAddressCIDR ErrorCode = "SingleIPAddressCIDR"

	// VMSSRequiredForAvailabilityZone indicates Virtual Machine Scale Set is required for availability zone.
	VMSSRequiredForAvailabilityZone ErrorCode = "VMSSRequiredForAvailabilityZone"

	// SLBRequiredForAvailabilityZone indicates Standard Load Balancer is required for availability zone.
	SLBRequiredForAvailabilityZone ErrorCode = "SLBRequiredForAvailabilityZone"

	// AvailabilityZoneNotSupported indicates Availability zone is not supported in this region.
	AvailabilityZoneNotSupported ErrorCode = "AvailabilityZoneNotSupported"

	// InvalidAvailabilityZoneFormat indicates that Availability Zone should be inputted in numeric format.
	InvalidAvailabilityZoneFormat ErrorCode = "InvalidAvailabilityZoneFormat"

	// SLBRequiredForMultipleAgentPools indicates Standard Load Balancer is required for multiple agent pools.
	SLBRequiredForMultipleAgentPools ErrorCode = "SLBRequiredForMultipleAgentPools"

	// ClusterKubernetesFailure indicates failure to communicate to Kubernetes.
	ClusterKubernetesFailure ErrorCode = "ClusterKubernetesFailure"

	// AgentPoolProfileIsNil indicates that an unexpected nil agentpool profile was received
	AgentPoolProfileIsNil ErrorCode = "AgentPoolProfileIsNil"

	// Error Subcode strings add more information about specific error codes

	// UnhandledError indicates that the error is not handled well
	UnhandledError = "UnhandledError"
	// ErrorWhilePolling indicates that there is an error during polling
	ErrorWhilePolling = "ErrorWhilePolling"
	// OperationMissingError indicates that an operation with unexpected status is missing error
	OperationMissingError = "OperationMissingError"
	// DeleteResourceGroupFailed indicates that there was a failure in deleting
	// a resource group
	DeleteResourceGroupFailed = "DeleteResourceGroupFailed"
	// DeleteEntityFailed indicates that we were unable to delete from the
	// entity store
	DeleteEntityFailed = "DeleteEntityFailed"
	// GetEntityFailed indicates that we were unable to fetch data from the
	// entity store
	GetEntityFailed = "GetEntityFailed"
	// OperationTimeout is a generic error that indicates a timeout occurred
	// when trying to perform and operation
	OperationTimeout = "OperationTimeout"
	// OperationTimeout is a generic error that indicates a timeout occurred
	// when trying to perform and operation
	OperationMaxRetryExceeded = "OperationMaxRetryExceeded"
	// PanicCaught indicates that the code caught a panic during an operation
	PanicCaught = "PanicCaught"
	// GetTokenFailed indicates that we were unable to fetch token from ARM
	GetTokenFailed = "GetTokenFailed"
	// InvalidDequeueOperation indicates that there is an operation that the service isn't setup to handle
	InvalidDequeueOperation = "InvalidDequeueOperation"
	// FailedEnqueueOperation indicates that we failed to add the operation to the queue
	FailedEnqueueOperation = "FailedEnqueueOperation"
	// UnderlayOverConstrainedMessageInternal indicates underlay over constrained
	UnderlayOverConstrainedMessageInternal = "underlay is overconstrained"
	// AZUnderlayOverConstrainedMessageInternal indicates underlay over constrained
	AZUnderlayOverConstrainedMessageInternal = "az underlay is overconstrained"
	// VMSSAgentBakerNotEnabledError indicates that the vmss agent baker togen has not been enabled
	VMSSAgentBakerNotEnabled = "VMSSAgentBakerNotEnabled"
	// SetAgentPoolCustomHyperkubeImageError indicates that hitting error when setting agent pool custom hyperkube image
	SetAgentPoolCustomHyperkubeImageError = "SetAgentPoolCustomHyperkubeImageError"
	// SetAgentPoolCustomKubeProxyImageError indicates that hitting error when setting agent pool custom kube-proxy image
	SetAgentPoolCustomKubeProxyImageError = "SetAgentPoolCustomKubeProxyImageError"
	// SetAgentPoolCustomKubeBinaryURLError indicates that hitting error when setting agent pool custom kube binary URL
	SetAgentPoolCustomKubeBinaryURLError = "SetAgentPoolCustomKubeBinaryURLError"
	// AADAppSecretTooLong indicates that the secret provided for the
	// AAD App Secret is too long
	AADAppSecretTooLong ErrorCode = "AADAppSecretTooLong"
	// RBACNotSupportedWithOrchestratorVersion indicates that RBAC feature is not supported with this version of K8s
	RBACNotSupportedWithOrchestratorVersion ErrorCode = "RBACNotSupportedWithOrchestratorVersion"
	// PodSecurityPolicyNotSupportedWithOrchestratorVersion indicates that PodSecurityPolicy is not supported with this version of K8s
	PodSecurityPolicyNotSupportedWithOrchestratorVersion ErrorCode = "PodSecurityPolicyNotSupportedWithOrchestratorVersion"
	// PodSecurityPolicyRequiresRBAC indicates that PodSecurityPolicy is not supported when RBAC is not enabled
	PodSecurityPolicyRequiresRBAC ErrorCode = "PodSecurityPolicyRequiresRBAC"
	// PodSecurityPolicyNotSupportedWithAzurePolicyAddon indicates that PodSecurityPolicy is not supported with azure policy addon
	PodSecurityPolicyNotSupportedWithAzurePolicyAddon ErrorCode = "PodSecurityPolicyNotSupportedWithAzurePolicyAddon"
	// APIServerWhitelistNotAllowedWithoutSingleIPPerCCP indicates that SingleIPPerCCP feature is not enabled for this cluster and API server whitelisting is requested
	APIServerWhitelistNotAllowedWithoutSingleIPPerCCP ErrorCode = "APIServerWhitelistNotAllowedWithoutSingleIPPerCCP"

	//TemplateDeploymentCanceled indicates the user canceled the operation
	TemplateDeploymentCanceled ErrorCode = "TemplateDeploymentCanceled"
	// ResourceGroupBeingDeleted indicates the MC_ resource group is being deleted
	ResourceGroupBeingDeleted ErrorCode = "ResourceGroupBeingDeleted"
	// ResourceGroupNotFound means exactly what it sounds like
	ResourceGroupNotFound ErrorCode = "ResourceGroupNotFound"
	// ResourceGroupCheckError means there is an internal error to check resource group but not NotFound
	ResourceGroupCheckError ErrorCode = "ResourceGroupCheckError"
	// AuthorizationFailed indicates the service principal doesn't have permission on target resource
	AuthorizationFailed ErrorCode = "AuthorizationFailed"
	// KeyvaultGetFailed indicates RP wasn't able to read from KV
	KeyvaultGetFailed ErrorCode = "KeyvaultGetFailed"
	// UnexpectedJWTTokenValue indicates JWT token value was nil or unexpected
	UnexpectedJWTTokenValue ErrorCode = "UnexpectedJWTTokenValue"
	InvalidAPIVersion       ErrorCode = "InvalidAPIVersion"
	UnmarshalError          ErrorCode = "UnmarshalError"

	GetControlPlaneV1Error                ErrorCode = "GetControlPlaneV1Error"
	GetControlPlaneV1UnexpectedStatusCode ErrorCode = "GetControlPlaneV1UnexpectedStatusCode"
	PutControlPlaneV1Error                ErrorCode = "PutControlPlaneV1Error"
	PutControlPlaneV1UnexpectedStatusCode ErrorCode = "PutControlPlaneV1UnexpectedStatusCode"

	GetGetUnderlayV1Error ErrorCode = "GetGetUnderlayV1Error"

	FindVMSSNameFailed                  ErrorCode = "FindVMSSNameFailed"
	TemplateDeploymentFailed            ErrorCode = "TemplateDeploymentFailed"
	ScaleVMASAgentPoolFailed            ErrorCode = "ScaleVMASAgentPoolFailed"
	UpgradeVMASAgentPoolFailed          ErrorCode = "UpgradeVMASAgentPoolFailed"
	CreateVMASAgentPoolFailed           ErrorCode = "CreateVMASAgentPoolFailed"
	CreateAvailabilitySetFailed         ErrorCode = "CreateAvailabilitySetFailed"
	DeleteVMSSAgentPoolFailed           ErrorCode = "DeleteVMSSAgentPoolFailed"
	ScaleVMSSAgentPoolFailed            ErrorCode = "ScaleVMSSAgentPoolFailed"
	CreateVMSSAgentPoolFailed           ErrorCode = "CreateVMSSAgentPoolFailed"
	UpgradeVMSSAgentPoolFailed          ErrorCode = "UpgradeVMSSAgentPoolFailed"
	ReconcileVMSSAgentPoolFailed        ErrorCode = "ReconcileVMSSAgentPoolFailed"
	UpdateVMSSAgentPoolFailed           ErrorCode = "UpdateVMSSAgentPoolFailed"
	ACIConnectorRequiresAzureNetworking ErrorCode = "ACIConnectorRequiresAzureNetworking"
	ACIConnectorRegionNotAvailable      ErrorCode = "ACIConnectorRegionNotAvailable"
	InvalidACIConnectorSubnetName       ErrorCode = "InvalidACIConnectorSubnetName"
	RBACNotEnabledForAAD                ErrorCode = "RBACNotEnabledForAAD"
	HealthCheckFailed                   ErrorCode = "HealthCheckFailed"
	UpdateInternalError                 ErrorCode = "UpdateInternalError"
	AdminUpdateInternalError            ErrorCode = "AdminUpdateInternalError"

	SDKValidationError                ErrorCode = "SDKValidationError"
	URLError                          ErrorCode = "URLError"
	ResourceGroupCreationError        ErrorCode = "ResourceGroupCreationError"
	GenerateConfigError               ErrorCode = "GenerateConfigError"
	ValidateConfigError               ErrorCode = "ValidateConfigError"
	ARMTemplateGenerationError        ErrorCode = "ARMTemplateGenerationError"
	ReplaceVMSSBootstrappingDataError ErrorCode = "ReplaceVMSSBootstrappingDataError"
	ReplaceVMASBootstrappingDataError ErrorCode = "ReplaceVMASBootstrappingDataError"

	UnknownNetworkPlugin          ErrorCode = "UnknownNetworkPlugin"
	GetAzureBearerAuthorizerError ErrorCode = "GetAzureBearerAuthorizerError"
	CertificateValidationError    ErrorCode = "CertificateValidationError"
	NilCertificateProfile         ErrorCode = "NilCertificateProfile"

	// Error codes related to Windows GMSA
	ReconcileCcpWindowsGmsaProfileError            ErrorCode = "ReconcileCcpWindowsGmsaProfileError"
	WindowsGmsaFeatureFlagNotEnabled               ErrorCode = "WindowsGmsaFeatureFlagNotEnabled"
	WindowsGmsaNotSupportedWithOrchestratorVersion ErrorCode = "WindowsGmsaNotSupportedWithOrchestratorVersion"
	WindowsGmsaFeatureNotEnabled                   ErrorCode = "WindowsGmsaFeatureNotEnabled"

	GetCheckAccessClientError                                      ErrorCode = "GetCheckAccessClientError"
	CheckAccessError                                               ErrorCode = "CheckAccessError"
	CustomRouteTableMissingPermission                              ErrorCode = "CustomRouteTableMissingPermission"
	CustomRouteTableWithUnsupportedMSIType                         ErrorCode = "CustomRouteTableWithUnsupportedMSIType"
	ClusterDeploymentError                                         ErrorCode = "ClusterDeploymentError"
	ClusterInitializationFailed                                    ErrorCode = "ClusterInitializationFailed"
	UpdateBlobInitializationFailed                                 ErrorCode = "UpdateBlobInitializationFailed"
	ClientCreationError                                            ErrorCode = "ClientCreationError"
	ConsoleNotReady                                                ErrorCode = "ConsoleNotReady"
	InfraAPIError                                                  ErrorCode = "InfraAPIError"
	InfraDaemonSetsNotReady                                        ErrorCode = "InfraDaemonSetsNotReady"
	InfraStatefulSetsNotReady                                      ErrorCode = "InfraStatefulSetsNotReady"
	InfraDeploymentsNotReady                                       ErrorCode = "InfraDeploymentsNotReady"
	MasterScaleSetHashError                                        ErrorCode = "MasterScaleSetHashError"
	WorkerScaleSetHashError                                        ErrorCode = "WorkerScaleSetHashError"
	MasterUpdateReadBlobError                                      ErrorCode = "MasterUpdateReadBlobError"
	WorkerUpdateReadBlobError                                      ErrorCode = "WorkerUpdateReadBlobError"
	MasterReadBlobError                                            ErrorCode = "MasterReadBlobError"
	UpdateMasterDrainError                                         ErrorCode = "UpdateMasterDrainError"
	UpdateMasterDeallocateError                                    ErrorCode = "UpdateMasterDeallocateError"
	UpdateMasterUpdateVMsError                                     ErrorCode = "UpdateMasterUpdateVMsError"
	UpdateMasterReimageError                                       ErrorCode = "UpdateMasterReimageError"
	UpdateMasterStartError                                         ErrorCode = "UpdateMasterStartError"
	UpdateMasterWaitForReadyError                                  ErrorCode = "UpdateMasterWaitForReadyError"
	UpdateMasterUpdateBlobError                                    ErrorCode = "UpdateMasterUpdateBlobError"
	UpdateWorkerListScaleSetsError                                 ErrorCode = "UpdateWorkerListScaleSetsError"
	UpdateWorkerReadBlobError                                      ErrorCode = "UpdateWorkerReadBlobError"
	UpdateWorkerDrainError                                         ErrorCode = "UpdateWorkerDrainError"
	UpdateWorkerCreateScaleSetError                                ErrorCode = "UpdateWorkerCreateScaleSetError"
	UpdateWorkerUpdateScaleSetError                                ErrorCode = "UpdateWorkerUpdateScaleSetError"
	UpdateWorkerDeleteScaleSetError                                ErrorCode = "UpdateWorkerDeleteScaleSetError"
	UpdateWorkerWaitForReadyError                                  ErrorCode = "UpdateWorkerWaitForReadyError"
	UpdateWorkerUpdateBlobError                                    ErrorCode = "UpdateWorkerUpdateBlobError"
	UpdateWorkerDeleteVMError                                      ErrorCode = "UpdateWorkerDeleteVMError"
	InvalidateClusterSecretsError                                  ErrorCode = "InvalidateClusterSecretsError"
	RegenerateClusterSecretsError                                  ErrorCode = "RegenerateClusterSecretsError"
	AdminUpdateFailedError                                         ErrorCode = "AdminUpdateFailedError"
	PluginAuthError                                                ErrorCode = "PluginAuthError"
	AzureClientError                                               ErrorCode = "AzureClientError"
	InsufficientMaxPods                                            ErrorCode = "InsufficientMaxPods"
	MaxPodsLimitExceeded                                           ErrorCode = "MaxPodsLimitExceeded"
	InsufficientAgentPoolMaxPodsPerAgentPool                       ErrorCode = "InsufficientAgentPoolMaxPodsPerAgentPool"
	InvalidClusterVersion                                          ErrorCode = "InvalidClusterVersion"
	AzurePolicyNotSupportedWithOrchestratorVersion                 ErrorCode = "AzurePolicyNotSupportedWithOrchestratorVersion"
	GitOpsFeatureFlagNotEnabled                                    ErrorCode = "GitOpsFeatureFlagNotEnabled"
	GitOpsRegionNotAvailable                                       ErrorCode = "GitOpsRegionNotAvailable"
	NodePublicIPRequiresVMSS                                       ErrorCode = "NodePublicIPRequiresVMSS"
	NodePublicIPCanNotBeFalseWithIPPrefixSet                       ErrorCode = "NodePublicIPCanNotBeFalseWithIPPrefixSet"
	NodePublicIPFeatureFlagNotEnabled                              ErrorCode = "NodePublicIPFeatureFlagNotEnabled"
	ResourceNotPermittedOnDelegatedSubnetError                     ErrorCode = "ResourceNotPermittedOnDelegatedSubnet"
	SubnetIsFull                                                   ErrorCode = "SubnetIsFull"
	ReconcileResourceProviderRegistrationError                     ErrorCode = "ReconcileResourceProviderRegistrationError"
	ReconcileBillingUsageError                                     ErrorCode = "ReconcileBillingUsageError"
	ReconcileVNetError                                             ErrorCode = "ReconcileVNetError"
	ReconcileRouteTableError                                       ErrorCode = "ReconcileRouteTableError"
	ReconcileNetworkSecurityGroupError                             ErrorCode = "ReconcileNetworkSecurityGroupError"
	GetLoadBalancerError                                           ErrorCode = "GetLoadBalancerError"
	ListLoadBalancerError                                          ErrorCode = "ListLoadBalancerError"
	NoLoadBalancerFound                                            ErrorCode = "NoLoadBalancerFound"
	ReconcileStandardLoadBalancerError                             ErrorCode = "ReconcileStandardLoadBalancerError"
	StandardLoadBalancerNotFoundError                              ErrorCode = "StandardLoadBalancerNotFoundError"
	StandardLoadBalancerWithoutOutboundRuleError                   ErrorCode = "StandardLoadBalancerWithoutOutboundRuleError"
	StandardLoadBalancerNotHavingExactlyOneOutboundRuleError       ErrorCode = "StandardLoadBalancerNotHavingExactlyOneOutboundRuleError"
	StandardLoadBalancerWithNoEffectiveOutboundIPOrIPPrefixError   ErrorCode = "StandardLoadBalancerWithNoEffectiveOutboundIPOrIPPrefixError"
	StandardLoadBalancerWithoutBackendPoolError                    ErrorCode = "StandardLoadBalancerWithoutBackendPoolError"
	NilOutboundInfoError                                           ErrorCode = "NilOutboundInfoError"
	ReconcileDiagnosticSettingsError                               ErrorCode = "ReconcileDiagnosticSettingsError"
	ReconcileUnderlayProfileError                                  ErrorCode = "ReconcileUnderlayProfileError"
	ReconcilePrivateLinkServiceError                               ErrorCode = "ReconcilePrivateLinkServiceError"
	ReconcilePrivateEndpointConnectionsError                       ErrorCode = "ReconcilePrivateEndpointConnectionsError"
	ReconcilePrivateLinkProfileError                               ErrorCode = "ReconcilePrivateLinkProfileError"
	ReconcilePrivateLinkUnderlaySubnetError                        ErrorCode = "ReconcilePrivateLinkUnderlaySubnetError"
	ReconcilePrivateEndpointSubnetError                            ErrorCode = "ReconcilePrivateEndpointSubnetError"
	ReconcilePrivateEndpointSubnetWithNATGatewayError              ErrorCode = "ReconcilePrivateEndpointSubnetWithNATGatewayError"
	ReconcilePrivateEndpointSubnetWithServiceEndpointPoliciesError ErrorCode = "ReconcilePrivateEndpointSubnetWithServiceEndpointPoliciesError"
	ReconcilePrivateEndpoint                                       ErrorCode = "ReconcilePrivateEndpoint"
	ReconcilePrivateDNS                                            ErrorCode = "ReconcilePrivateDNS"
	ReconcilePrivateClusterPublicDNSError                          ErrorCode = "ReconcilePrivateClusterPublicDNSError"
	ReconcilePrivateConnectBalancerError                           ErrorCode = "ReconcilePrivateConnectBalancerError"
	ReconcilePrivateConnectProfileError                            ErrorCode = "ReconcilePrivateConnectProfileError"
	ReconcilePrivateConnectIPError                                 ErrorCode = "ReconcilePrivateConnectIPError"
	ReconcilePrivateConnectCleanupError                            ErrorCode = "ReconcilePrivateConnectCleanupError"
	ReconcileNetworkAssociationLinkError                           ErrorCode = "ReconcileNetworkAssociationLinkError"
	ReconcileServiceAssociationLinkError                           ErrorCode = "ReconcileServiceAssociationLinkError"
	GetPrivateLinkServiceError                                     ErrorCode = "GetPrivateLinkServiceError"
	ListPrivateLinkServiceError                                    ErrorCode = "ListPrivateLinkServiceError"
	GetPrivateEndpointConnectionError                              ErrorCode = "GetPrivateEndpointConnectionError"
	GetPrivateEndpointSettingError                                 ErrorCode = "GetPrivateEndpointSettingError"
	PrivateEndpointConnectionOperationError                        ErrorCode = "PrivateEndpointConnectionOperationError"
	PrivateEndpointConnectionNotInSucceededStateError              ErrorCode = "PrivateEndpointConnectionNotInSucceededStateError"
	AppLensDetectorError                                           ErrorCode = "AppLensDetectorError"
	GetDiagnosticSettingsError                                     ErrorCode = "GetDiagnosticSettingsError"
	GetDiagnosticSettingsClientFailed                              ErrorCode = "GetDiagnosticSettingsClientFailed"
	GetResourceProviderClientFailed                                ErrorCode = "GetResourceProviderClientFailed"
	GetPublicIPAddressClientFailed                                 ErrorCode = "GetPublicIPAddressClientFailed"
	GetNetworkUsagesClientFailed                                   ErrorCode = "GetNetworkUsagesClientFailed"
	GetPublicIPPrefixClientFailed                                  ErrorCode = "GetPublicIPPrefixClientFailed"
	GetLoadBalancerClientFailed                                    ErrorCode = "GetLoadBalancerClientFailed"
	GetNATGatewayClientFailed                                      ErrorCode = "GetNATGatewayClientFailed"
	ErrorValidatingLoadBalancerProfileOutboundIPs                  ErrorCode = "ErrorValidatingLoadBalancerProfileOutboundIPs"
	ErrorValidatingLoadBalancerProfileOutboundIPPrefixes           ErrorCode = "ErrorValidatingLoadBalancerProfileOutboundIPPrefixes"
	UpdateNodeSPError                                              ErrorCode = "UpdateNodeSPError"
	GenerateTemplateError                                          ErrorCode = "GenerateTemplateError"
	UpdateTemplateError                                            ErrorCode = "UpdateTemplateError"
	DeployTemplateError                                            ErrorCode = "DeployTemplateError"
	InvalidTemplateDeploymentError                                 ErrorCode = "InvalidTemplateDeploymentError"
	UnknownDeploymentError                                         ErrorCode = "UnknownDeploymentError"
	PrepareContainerServiceDataModelError                          ErrorCode = "PrepareContainerServiceDataModelError"
	ReconcileEtcdBackupBlobStorageError                            ErrorCode = "ReconcileEtcdBackupBlobStorageError"
	PrepareLargeClusterSettingError                                ErrorCode = "PrepareLargeClusterSettingError"
	UpdateVMSSSinglePlacementGroupError                            ErrorCode = "UpdateVMSSSinglePlacementGroupError"
	GetGraphClientError                                            ErrorCode = "GetGraphClientError"
	InvalidManagedIdentityResourceID                               ErrorCode = "InvalidManagedIdentityResourceID"
	GetVMSSClientError                                             ErrorCode = "GetVMSSClientError"
	GetHcpAPIClientError                                           ErrorCode = "GetHcpAPIClientError"
	GetAvailabilitySetError                                        ErrorCode = "GetAvailabilitySetError"
	ReconcileOpenVPNError                                          ErrorCode = "ReconcileOpenVPNError"
	ReconcileUsersError                                            ErrorCode = "ReconcileUsersError"
	ReconcileProbeBearerTokenError                                 ErrorCode = "ReconcileProbeBearerTokenError"
	ReconcileAccessProfilesError                                   ErrorCode = "ReconcileAccessProfilesError"
	DeleteEtcdBackupContainerError                                 ErrorCode = "DeleteEtcdBackupContainerError"
	CreateEtcdBackupContainerError                                 ErrorCode = "CreateEtcdBackupContainerError"
	ControlPlaneIDNotSet                                           ErrorCode = "ControlPlaneIDNotSet"
	ReconcileAzurePolicyAddonError                                 ErrorCode = "ReconcileAzurePolicyAddonProfileError"
	ReconcileSecurityProfileError                                  ErrorCode = "ReconcileSecurityProfileError"
	ReconcileControlPlaneCertificatesError                         ErrorCode = "ReconcileControlPlaneCertificatesError"
	ReconcileAPIServerCertificatesError                            ErrorCode = "ReconcileAPIServerCertificatesError"
	ReconcileAADAuthProfileError                                   ErrorCode = "ReconcileAADAuthProfileError"
	ReconcileAADWebhookCertificatesError                           ErrorCode = "ReconcileAADWebhookCertificatesError"
	ReconcileAggregatorCertificatesError                           ErrorCode = "ReconcileAggregatorCertificatesError"
	ReconcileEtcdCertificatesError                                 ErrorCode = "ReconcileEtcdCertificatesError"
	ReconcileEtcdBackupEncryptionProfileError                      ErrorCode = "ReconcileEtcdBackupEncryptionProfileError"
	ReconcileACIConnectorCertificatesError                         ErrorCode = "ReconcileACIConnectorCertificatesError"
	ReconcileMSICredentialError                                    ErrorCode = "ReconcileMSICredentialError"
	MSICredentialStoreOperationError                               ErrorCode = "MSICredentialStoreOperationError"
	GetMSICredentialError                                          ErrorCode = "GetMSICredentialError"
	MSICredentialNotExistError                                     ErrorCode = "MSICredentialNotExistError"
	StorageAccountProfilePreparationError                          ErrorCode = "StorageAccountProfilePreparationError"
	EmptyStorageAccountProfileError                                ErrorCode = "EmptyStorageAccountProfileError"
	OverlaymgrReconcileError                                       ErrorCode = "OverlaymgrReconcileError"
	AgentPoolConfigmapReconcileError                               ErrorCode = "AgentPoolConfigmapReconcileError"
	ReconcileOverlayServiceIPError                                 ErrorCode = "ReconcileOverlayServiceIPError"
	DeleteControlPlaneRecordFailed                                 ErrorCode = "DeleteControlPlaneRecordFailed"
	PrivateClusterCapacityOverConstrained                          ErrorCode = "PrivateClusterCapacityOverConstrained"
	PrivateClusterWithAZOverConstrained                            ErrorCode = "PrivateClusterWithAZOverConstrained"
	PrivateClusterNotAllowedInLocation                             ErrorCode = "PrivateClusterNotAllowedInLocation"
	PrivateClusterNotSupportedWithPaidSKU                          ErrorCode = "PrivateClusterNotSupportedWithPaidSKU"
	InvalidPrivateDNSZoneMode                                      ErrorCode = "InvalidPrivateDNSZoneMode"
	GetPrivateDNSZoneClientFailed                                  ErrorCode = "GetPrivateDNSZoneClientFailed"
	GetPrivateClusterDNSResourceReferenceError                     ErrorCode = "GetPrivateClusterDNSResourceReferenceError"
	InvalidPrivateDNSZoneResourceID                                ErrorCode = "InvalidPrivateDNSZoneResourceID"
	PrivateDNSZoneResourceNotFoundError                            ErrorCode = "PrivateDNSZoneResourceNotFoundError"
	SystemAssignedIdentityNotSupportedForCustomPrivateDNSZone      ErrorCode = "SystemAssignedIdentityNotSupportedForCustomPrivateDNSZone"
	PrivateDNSZoneAuthorizationFailed                              ErrorCode = "PrivateDNSZoneAuthorizationFailed"
	FQDNSubdomainRecordAlreadyExistsError                          ErrorCode = "FQDNSubdomainRecordAlreadyExistsError"
	NonePrivateDNSZonePublicFQDNDisabledError                      ErrorCode = "NonePrivateDNSZonePublicFQDNDisabledError"
	GetPublicIPAddressByResourceIDError                            ErrorCode = "GetPublicIPAddressByResourceIDError"
	GetPublicIPPrefixByResourceIDError                             ErrorCode = "GetPublicIPPrefixByResourceIDError"
	InvalidPublicIPPrefixType                                      ErrorCode = "InvalidPublicIPPrefixType"
	InvalidPublicIPPrefixDifferentSub                              ErrorCode = "InvalidPublicIPPrefixDifferentSub"
	InvalidPublicIPPrefixDifferentLocation                         ErrorCode = "InvalidPublicIPPrefixDifferentLocation"
	PublicIPPrefixIsAttachedToLoadBalancer                         ErrorCode = "PublicIPPrefixIsAttachedToLoadBalancer"
	GetClusterStateError                                           ErrorCode = "GetClusterStateError"
	DeletingControlPlaneError                                      ErrorCode = "DeletingControlPlaneError"
	ProvisioningControlPlaneError                                  ErrorCode = "ProvisioningControlPlaneError"
	ReconcileAgentPoolConfigmapError                               ErrorCode = "ReconcileAgentPoolConfigmapError"
	ReconcileAddonIdentityError                                    ErrorCode = "ReconcileAddonIdentityError"
	ReconcileKubeletIdentityError                                  ErrorCode = "ReconcileKubeletIdentityError"
	ReconcileCloudConfigSecretError                                ErrorCode = "ReconcileCloudConfigSecretError"
	GetUserAssignedIdentityResourceError                           ErrorCode = "GetUserAssignedIdentityResourceError"
	ErrorCodeEmptyPayload                                          ErrorCode = "EmptyPayload"
	ErrorCodeCCPMigrationError                                     ErrorCode = "CCPMigrationError"
	ErrorCodeCCPDeallocationError                                  ErrorCode = "CCPDeallocationError"
	ErrorCodeFailedToStopManagedCluster                            ErrorCode = "ManagedClusterStopError"
	ErrorCodeFailedToStartManagedCluster                           ErrorCode = "ManagedClusterStartError"
	ErrorCodeDrainUnderlayError                                    ErrorCode = "DrainUnderlayError"
	ErrorCodeRotateClusterCertificates                             ErrorCode = "RotateClusterCertificatesError"
	ErrorCodeServicePrincipalSecretUpdate                          ErrorCode = "ServicePrincipalSecretUpdateError"
	ErrorCodeAadProfileUpdate                                      ErrorCode = "AadProfileUpdateError"
	ErrorCodeUpdateAccessProfiles                                  ErrorCode = "UpdateAccessProfilesError"
	ErrorCodeEnableEncryptionAtHost                                ErrorCode = "EnableEncryptionAtHostError"
	ErrorCodeEnableSwiftNetworking                                 ErrorCode = "EnableSwiftNetworkingError"
	ErrorCodeEnableUltraSSD                                        ErrorCode = "EnableUltraSSDError"
	ErrorCodeGetVMSize                                             ErrorCode = "GetVMSizeError"
	InvalidImageRef                                                ErrorCode = "InvalidImageRef"
	InvalidGalleryImageRef                                         ErrorCode = "InvalidGalleryImageRef"
	InvalidOSSKU                                                   ErrorCode = "InvalidOSSKU"
	InvalidAdminGroupObjectID                                      ErrorCode = "InvalidAdminGroupObjectID"
	InvalidCustomizedUbuntu                                        ErrorCode = "InvalidCustomizedUbuntu"
	InvalidCustomizedWindows                                       ErrorCode = "InvalidCustomizedWindows"
	RequestDisallowedByPolicy                                      ErrorCode = "RequestDisallowedByPolicy"
	SystemPoolMustBeRegular                                        ErrorCode = "SystemPoolMustBeRegular"
	SystemPoolMustBeLinux                                          ErrorCode = "SystemPoolMustBeLinux"
	SystemPoolSkuTooLow                                            ErrorCode = "SystemPoolSkuTooLow"
	ReconcileAddonIdentityRoleAssignmentError                      ErrorCode = "ReconcileAddonIdentityRoleAssignmentError"
	ReconcileApiserverProxyProfileError                            ErrorCode = "ReconcileApiserverProxyProfileError"
	ErrorCodeUnsupportedGen2VMSize                                 ErrorCode = "UnsupportedGen2VMSize"
	UpdateAgentPoolNodeImageReferenceFailed                        ErrorCode = "UpdateAgentPoolNodeImageReferenceFailed"
	DeleteAgentPoolNodeLabelsFailed                                ErrorCode = "DeleteAgentPoolNodeLabelsFailed"
	ServicePrincipalValidationClientError                          ErrorCode = "ServicePrincipalValidationClientError"
	ServicePrincipalNotFoundTimedoutError                          ErrorCode = "ServicePrincipalNotFoundTimedoutError"
	ReconcilePortalDNSError                                        ErrorCode = "ReconcilePortalDNSError"
	InvalidPPGID                                                   ErrorCode = "InvalidPPGID"
	InvalidPPGSubscription                                         ErrorCode = "InvalidPPGSubscription"
	InvalidPPGRegion                                               ErrorCode = "InvalidPPGRegion"
	PPGNotSupportMultipleZones                                     ErrorCode = "PPGNotSupportMultipleZones"
	PPGSupportsVMSSOnly                                            ErrorCode = "PPGSupportsVMSSOnly"
	PPGSupportsSLBOnly                                             ErrorCode = "PPGSupportsSLBOnly"
	PPGNotFound                                                    ErrorCode = "PPGNotFound"
	GetPPGError                                                    ErrorCode = "GetPPGError"
	GetPPGClientError                                              ErrorCode = "GetPPGClientError"
	ErrorCodeUnsupportedGPUDedicatedVHDVMSize                      ErrorCode = "UnsupportedGPUDedicatedVHDVMSize"
	UbuntuDistroNotSupportedForContainerRuntime                    ErrorCode = "UbuntuDistroNotSupportedForContainerRuntime"
	UnsupportedCustomContainerRuntime                              ErrorCode = "UnsupportedCustomContainerRuntime"
	ContainerdSupportsVMSSOnly                                     ErrorCode = "ContainerdSupportsVMSSOnly"
	InvalidEphemeralOSDiskHeaderError                              ErrorCode = "InvalidEphemeralOSDiskHeaderError"
	InvalidOSDiskTypeSettingError                                  ErrorCode = "InvalidOSDiskTypeSettingError"
	VMSizeDoesNotSupportEphemeralOS                                ErrorCode = "VMSizeDoesNotSupportEphemeralOS"
	VMSizeDoesNotSupportEncryptionAtHost                           ErrorCode = "VMSizeDoesNotSupportEncryptionAtHost"
	SubscriptionNotEnabledEncryptionAtHost                         ErrorCode = "SubscriptionNotEnabledEncryptionAtHost"
	CheckEncryptionAtHostFeatureRegisterationFailure               ErrorCode = "CheckEncryptionAtHostFeatureRegisterationFailure"
	EncryptionAtHostNotAllowedOnVMAS                               ErrorCode = "EncryptionAtHostNotAllowedOnVMAS"
	VMSizeNotSupported                                             ErrorCode = "VMSizeNotSupported"
	VMCacheSizeTooSmall                                            ErrorCode = "VMCacheSizeTooSmall"
	VMTemporaryDiskTooSmall                                        ErrorCode = "VMTemporaryDiskTooSmall"
	EphemeralOSAndBYOKNotSupported                                 ErrorCode = "EphemeralOSAndBYOKNotSupported"
	DiskEncryptionSetError                                         ErrorCode = "DiskEncryptionSetError"
	CannotChangeOSDiskType                                         ErrorCode = "CannotChangeOSDiskType"
	MissingAgentPoolStorageProfile                                 ErrorCode = "MissingAgentPoolStorageProfile"
	ReconcileKonnectivityProfileError                              ErrorCode = "ReconcileKonnectivityProfileError"
	NodeLabelKeyNotAllowed                                         ErrorCode = "NodeLabelKeyNotAllowed"
	InvalidAgentPoolOrchestratorVersion                            ErrorCode = "InvalidAgentPoolOrchestratorVersion"
	MaxDiskCacheBytesParseError                                    ErrorCode = "MaxDiskCacheBytesParseError"
	GetSKUStoreError                                               ErrorCode = "GetSKUStoreError"
	PolicyViolation                                                ErrorCode = "PolicyViolation"
	CustomNodeConfigFeatureNotSupported                            ErrorCode = "CustomNodeConfigFeatureNotSupported"
	CustomKubeletConfigOrCustomLinuxOSConfigCanNotBeChanged        ErrorCode = "CustomKubeletConfigOrCustomLinuxOSConfigCanNotBeChanged"
	InvalidCustomKubeletConfig                                     ErrorCode = "InvalidCustomKubeletConfig"
	InvalidCustomLinuxOSConfig                                     ErrorCode = "InvalidCustomLinuxOSConfig"
	CustomLinuxOSConfigInternalError                               ErrorCode = "CustomLinuxOSConfigInternalError"
	InvalidEnableACRTeleportHeaderError                            ErrorCode = "InvalidEnableACRTeleportHeaderError"
	UnsupportedContainerRuntimeForACRTeleport                      ErrorCode = "UnsupportedContainerRuntimeForACRTeleport"
	AppendPoliciesFieldsExist                                      ErrorCode = "AppendPoliciesFieldsExist"
	InvalidKubeletDiskSettingError                                 ErrorCode = "InvaliKubeletDiskSettingError"
	CannotChangeKubeletDisk                                        ErrorCode = "CannotChangeKubeletDisk"
	WindowsSharedImageGalleryNotAllowed                            ErrorCode = "WindowsSharedImageGalleryNotAllowed"
	ExtensionManagerFeatureFlagNotEnabled                          ErrorCode = "ExtensionManagerFeatureFlagNotEnabled"
	ExtensionManagerRegionNotAvailable                             ErrorCode = "ExtensionManagerRegionNotAvailable"
	ExtensionAddonNotFound                                         ErrorCode = "ExtensionAddonNotFound"
	ExtensionAlreadyEnabledInAKSAddons                             ErrorCode = "ExtensionAlreadyEnabledInAKSAddons"
	ReconcileAgentPoolIdentityError                                ErrorCode = "ReconcileAgentPoolIdentityError"
	ReconcileExtensionManagerError                                 ErrorCode = "ReconcileExtensionManagerError"
	OmsagentUseAADAuthFeatureFlagNotEnabled                        ErrorCode = "OmsagentUseAADAuthFeatureFlagNotEnabled"
	RoleAssignmentLimitExceeded                                    ErrorCode = "RoleAssignmentLimitExceeded"
	HTTPProxyConfigFeatureNotSupported                             ErrorCode = "HTTPProxyConfigFeatureNotSupported"
	HTTPProxyConfigCADecodeError                                   ErrorCode = "HTTPProxyConfigCADecodeError"
	HTTPProxyConfigInputError                                      ErrorCode = "HTTPProxyConfigInputError"
	HTTPProxyWrongUrlError                                         ErrorCode = "HTTPProxyWrongUrlError"
	ObjectIsDeletedButRecoverable                                  ErrorCode = "ObjectIsDeletedButRecoverable"
	ObjectIsBeingDeleted                                           ErrorCode = "ObjectIsBeingDeleted"
	EnableNamespaceResourcesNotSupported                           ErrorCode = "EnableNamespaceResourcesNotSupported"

	// error code relates to windows customized container runtime
	WindowsCustomContainerRuntimeUnsupported                         ErrorCode = "WindowsCustomContainerRuntimeUnsupported"
	WindowsCustomContainerRuntimeNotSupportedWithOrchestratorVersion ErrorCode = "WindowsCustomContainerRuntimeNotSupportedWithOrchestratorVersion"

	// Error codes related to byo kubelet identity
	CustomKubeletIdentityMissingPermissionError                ErrorCode = "CustomKubeletIdentityMissingPermissionError"
	CustomKubeletIdentityOnlySupportedOnUserAssignedMSICluster ErrorCode = "CustomKubeletIdentityOnlySupportedOnUserAssignedMSICluster"
	CustomKubeletIdentityPropertyError                         ErrorCode = "CustomKubeletIdentityPropertyError"
	InvalidWorkloadRuntimeSettingError                         ErrorCode = "InvalidWorkloadRuntimeSettingError"
	// NRP internal server error
	NRPInternalServerError                            ErrorCode = "InternalServerError"
	ReservedResourceName                              ErrorCode = "ReservedResourceName"
	ResourceGroupDeletionBlocked                      ErrorCode = "ResourceGroupDeletionBlocked"
	NetworkSecurityGroupInUseByVirtualMachineScaleSet ErrorCode = "NetworkSecurityGroupInUseByVirtualMachineScaleSet"
	LoadBalancerInUseByVirtualMachineScaleSet         ErrorCode = "LoadBalancerInUseByVirtualMachineScaleSet"
	PublicIPAddressCannotBeDeleted                    ErrorCode = "PublicIPAddressCannotBeDeleted"
	ReadOnlyDisabledSubscription                      ErrorCode = "ReadOnlyDisabledSubscription"
	//VMSS internal error
	NoVMSSInstanceView            ErrorCode = "NoVMSSInstanceView"
	NoVMSSInstanceViewExtension   ErrorCode = "NoVMSSInstanceViewExtension"
	NoVMSSInstanceExtensionStatus ErrorCode = "NoVMSSInstanceExtensionStatus"
	NoVMSSInstanceStatusMessage   ErrorCode = "NoVMSSInstanceStatusMessage"
	NoVMSSInstanceCSEExtension    ErrorCode = "NoVMSSInstanceCSEExtension"

	// Error codes related to OIDC
	ReconcileOIDCProfileError         ErrorCode = "ReconcileOIDCProfileError"
	ReconcileOIDCCertificatesError    ErrorCode = "ReconcileOIDCCertificatesError"
	OIDCIssuerFeatureFlagNotEnabled   ErrorCode = "OIDCIssuerFeatureFlagNotEnabled"
	OIDCIssuerFeatureCannotBeDisabled ErrorCode = "OIDCIssuerCannotBeDisabled"

	// private connect cluster
	PrivateClusterV2CapacityOverConstrained ErrorCode = "PrivateClusterV2CapacityOverConstrained"
	PrivateClusterV2WithAZOverConstrained   ErrorCode = "PrivateClusterV2WithAZOverConstrained"
	PrivateClusterV2NotAllowedInLocation    ErrorCode = "PrivateClusterV2NotAllowedInLocation"
	ApiserverSubnetConfigError              ErrorCode = "ApiserverSubnetConfigError"

	//Publisher failed to publish an event
	PublisherFailedToPublishEventError ErrorCode = "PublisherFailedToPublishEventError"

	// Multi-instance GPU related errors
	EmptyGPUInstanceProfile       ErrorCode = "EmptyGPUInstanceProfile"
	UnsupportedGPUInstanceProfile ErrorCode = "UnsupportedGPUInstanceProfile"
	VMSizeDoesNotSupportMIG       ErrorCode = "VMSizeDoesNotSupportMIG"

	GetPrimaryScaleSetNameFailed ErrorCode = "GetPrimaryScaleSetNameFailed"
	EmptyPrimaryScaleSetName     ErrorCode = "EmptyPrimaryScaleSetName"

	GetSubscriptionError ErrorCode = "GetSubscriptionError"

	EnsureLoadBalancersExistError  ErrorCode = "EnsureLoadBalancersExistError"
	FailedListLoadBalancer         ErrorCode = "FailedListLoadBalancer"
	FailedUpdateNodeImageReference ErrorCode = "FailedUpdateNodeImageReference"

	// Dual-stack networking related errors
	DualStackNotAllowedInLocation             ErrorCode = "DualStackNotAllowedInLocation"
	InvalidNetworkingConfig                   ErrorCode = "InvalidNetworkingConfig"
	UnsupportedDualStackKubernetesVersion     ErrorCode = "UnsupportedDualStackKubernetesVersion"
	UnsupportedDualStackNetworkPlugin         ErrorCode = "UnsupportedDualStackNetworkPlugin"
	UnsupportedDualStackNetworkPolicy         ErrorCode = "UnsupportedDualStackNetworkPolicy"
	VnetSubnetMissingIPv6                     ErrorCode = "VnetSubnetMissingIPv6"
	InvalidServiceCidrIPv6MaskSize            ErrorCode = "InvalidServiceCidrIPv6MaskSize"
	SingleStackVnetSubnetWithDualStackCluster ErrorCode = "SingleStackVnetSubnetWithDualStackCluster"

	// Capacity Reservation Group related error code
	CRGConflictWithVMAS          ErrorCode = "CRGConflictWithVMAS"
	CRGConflictWithUltraSSD      ErrorCode = "CRGConflictWithUltraSSD"
	CRGConflictWithPPG           ErrorCode = "CRGConflictWithPPG"
	CRGConflictWithSpotVM        ErrorCode = "CRGConflictWithSpotVM"
	CRGConflictWithDedicatedHost ErrorCode = "CRGConflictWithDedicatedHost"
	CRGConflictWithSPG           ErrorCode = "CRGConflictWithSPG"
)

const (
	// VMExtensionErrorMessageExitStatus is used for vmss cse extension error in vmss agentpool reconciler
	VMExtensionErrorMessageExitStatus string = "exit status="
	// VMSSInstanceErrorCode is used for vmss instance error in vmss agentpool reconciler
	VMSSInstanceErrorCode string = "vmssInstanceErrorCode="
)
