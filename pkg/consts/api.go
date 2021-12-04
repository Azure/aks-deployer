//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package consts

import "fmt"

// Case insensitive literals
const (
	SubscriptionsLiteral                       = "{sb:(?i)subscriptions}"
	ResourceGroupsLiteral                      = "{rg:(?i)resourcegroups}"
	ProvidersLiteral                           = "{pv:(?i)providers}"
	ContainerServicesLiteral                   = "{cs:(?i)containerservices}"
	ControlPlanesLiteral                       = "{cp:(?i)controlplanes}"
	LocationsLiteral                           = "{lc:(?i)locations}"
	OperationResultsLiteral                    = "{or:(?i)operationresults}"
	OperationsLiteral                          = "{op:(?i)operations}"
	DeploymentsLiteral                         = "{dp:(?i)deployments}"
	PreflightLiteral                           = "{pf:(?i)preflight}"
	InternalLiteral                            = "{in:(?i)internal}"
	ManagedClustersLiteral                     = "{mc:(?i)managedclusters}"
	VirtualNetworksLiteral                     = "{vn:(?i)virtualnetworks}"
	AgentPoolsLiteral                          = "{ap:(?i)agentPools}"
	SubnetsLiteral                             = "{sn:(?i)subnets}"
	AvailableAgentPoolVersionsLiteral          = "{ap:(?i)availableagentpoolversions}"
	OrchestratorsLiteral                       = "{or:(?i)orchestrators}"
	OSOptionsLiteral                           = "{os:(?i)osOptions}"
	UpgradeProfilesLiteral                     = "{us:(?i)upgradeprofiles}"
	DiagnosticsStateLiteral                    = "{ds:(?i)diagnosticsstate}"
	AccessProfilesLiteral                      = "{ap:(?i)accessprofiles}"
	AdminLiteral                               = "{ad:(?i)admin}"
	APILiteral                                 = "{ad:(?i)api}"
	PodsLiteral                                = "{po:(?i)pods}"
	LogLiteral                                 = "{lo:(?i)log}"
	EventLiteral                               = "{ev:(?i)events}"
	KubectlLiteral                             = "{ku:(?i)kubectl}"
	ContainersLiteral                          = "{co:(?i)containers}"
	UnderlaysLiteral                           = "{un:(?i)underlays}"
	ExpandUnderlayCapacityLiteral              = "{euc:(?i)expandunderlaycapacity}"
	OverlayLiteral                             = "{ov:(?i)overlay}"
	DefaultLiteral                             = "{up:(?i)default}"
	ListCredentialLiteral                      = "{li:(?i)listcredential}"
	ContainerServiceProviderLiteral            = "{mcs:(?i)microsoft.containerservice}"
	NetworkProviderLiteral                     = "{np:(?i)microsoft.network}"
	NodeLiteral                                = "{no:(?i)nodes}"
	ActionsLiteral                             = "{ac:(?i)actions}"
	ListClusterAdminCredentialLiteral          = "{lcac:(?i)listclusteradmincredential}"
	ListClusterUserCredentialLiteral           = "{lcuc:(?i)listclusterusercredential}"
	ListClusterMonitoringUserCredentialLiteral = "{lcuc:(?i)listclustermonitoringusercredential}"
	RunCommandLiteral                          = "{rc:(?i)runcommand}"
	CommandResultLiteral                       = "{cr:(?i)commandresults}"
	NotifyLiteral                              = "{ntf:(?i)notify}"
	DetectorsLiteral                           = "{de:(?i)detectors}"
	MigrateCustomerControlPlaneLiteral         = "{mccp:(?i)migrateccp}"
	DeallocateControlPlaneLiteral              = "{daccp:(?i)deallocateccp}"
	DrainCustomerControlPlanesLiteral          = "{dccps:(?i)drainccps}"
	UnderlayDataSourcesLiteral                 = "{da:(?i)datasources}"
	PrivateEndpointConnectionsLiteral          = "{pec:(?i)privateEndpointConnections}"
	PrivateLinkResourcesLiteral                = "{plr:(?i)privateLinkResources}"
	ResolvePrivateLinkServiceIDLiteral         = "{rpls:(?i)resolvePrivateLinkServiceId}"
	BackfillManagedClusterLiteral              = "{bfmc:(?i)backfillmanagedcluster}"
	ReimageManagedClusterLiteral               = "{rmc:(?i)reimageManagedCluster}"
	DelegateSubnetLiteral                      = "{ds:(?i)delegateSubnet}"
	UnDelegateSubnetLiteral                    = "{ds:(?i)undelegateSubnet}"
	ServiceOutboundIPRangesLiteral             = "{soip:(?i)serviceoutboundipranges}"
	MaintenanceConfigurationsLiteral           = "{mtc:(?i)maintenanceconfigurations}"
	JsonPatchLiteral                           = "{jp:(?i)jsonpatch}"
	ExtensionAddonsLiteral                     = "{ea:(?i)extensionaddons}"
	OutboundNetworkDependenciesEndpoints       = "{ond:(?i)outboundNetworkDependenciesEndpoints}"
	SnapshotsLiteral                           = "{ss:(?i)snapshots}"
	MigrateClusterV2Literal                    = "{mcv:(?i)migrateClusterV2}"
)

const (
	//PathunVersionedApiVersionParameter is the path parameter used in routing for apiversion
	PathunVersionedApiVersionParameter = "unVersioned"
	// PathSubscriptionIDParameter is the path parameter name used in routing for the subscription id
	PathSubscriptionIDParameter = "subscriptionId"
	// PathUnderlayQuarantinedParameter is the path parameter name used in routing for the underlay quarantined
	PathUnderlayActionParameter = "action"
	// PathResourceGroupNameParameter is the path parameter name used in routing for the resource group name
	PathResourceGroupNameParameter = "resourceGroupName"
	// PathResourceGroupNameParameter is the path parameter name used in routing for the control plane name
	PathControlPlaneParameter = "ControlPlaneId"
	// PathResourceNameParameter is the path parameter name used in routing for the resource name
	PathResourceNameParameter = "resourceName"
	// PathAgentPoolNameParameter is the agentpool parameter name used in route for agentpool name.
	PathAgentPoolNameParameter = "agentPoolName"
	// PathMaintenanceConfigurationNameParameter is the maintenance configuration parameter name used in route for maintenance configuration name.
	PathMaintenanceConfigurationNameParameter = "maintenanceConfigurationName"
	// PathVirtualNetworkNameParameter is the subnet parameter name used in route for virtual network name.
	PathVirtualNetworkNameParameter = "virtualNetworkName"
	// PathSubnetNameParameter is the subnet parameter name used in route for subnet name.
	PathSubnetNameParameter = "subnetName"
	// PathPodNameParameter is the path parameter name used in routing for the pod name
	PathPodNameParameter = "podName"
	// PathContainerNameParameter is the path parameter name used in routing for the container name
	PathContainerNameParameter = "containerName"
	// PathVMNameParameter is the path parameter name used in routing for the vm name
	PathVMNameParameter = "vmName"
	// PathActionNameParameter is the path parameter name used in routing for the action name
	PathActionNameParameter = "actionName"
	// PathDetectorNameParameter is the path parameter name used in routing for app lens detector name
	PathDetectorNameParameter = "detectorName"
	// PathRunCommandIdParameter is the path parameter name used in routing for command result request.
	PathRunCommandIdParameter = "commandId"
	// BodyKubectlCommandParameter is the body parameter used in routing for the kubectl command
	BodyKubectlCommandParameter = "kubectlcommand"
	// BodyRunCommandParameter is the body parameter name used in routing for run command
	BodyRunCommandParameter = "runcommand"
	// PathLocationParameter is the path parameter name used in routing for the location
	PathLocationParameter = "location"
	// PathOperationIDParameter is the path parameter name used in routing for the operation id
	PathOperationIDParameter = "operationId"
	// PathDeploymentNameParameter is the path parameter name used in routing for the deployment name
	PathDeploymentNameParameter = "deploymentName"
	// PathAccessProfileParameter is the path parameter name used in routing for the accessProfile role name
	PathAccessProfileParameter = "accessProfile"
	// PathUnderlayNameParameter is the path parameter name used in routing for the underlay name
	PathUnderlayNameParameter = "underlayName"
	// PathExtensionProviderParameter is the path parameter arm extension resource provider
	PathExtensionProviderParameter = "extensionProvider"
	// PathExtensionResourceTypeParameter is the path parameter arm extension resource type
	PathExtensionResourceTypeParameter = "extensionResourceType"
	// PathExtensionResourceNameParameter is the path parameter arm extension resource name
	PathExtensionResourceNameParameter = "extensionResourceName"
	// PathUnderlayDataSourceParameter is the path parameter name used in routing for the underlay data source
	PathUnderlayDataSourceParameter = "datasource"
	// PathPrivateEndpointConnectionNameParameter is the path parameter name used in routing for the private endpoint connection
	PathPrivateEndpointConnectionNameParameter = "privateEndpointConnectionName"
	// PathServiceNameParameter is the path parameter name used in routing for service name
	PathServiceNameParameter = "serviceName"
	// BodyIPRangesParameter is the body parameter name used in routing for IP ranges
	BodyIPRangesParameter = "ipRanges"

	// BodyUnderlayConfigParameter is the body parameter for underlay config
	BodyUnderlayConfigParameter = "underlayconfig"

	// ContainerServiceNamespace is the ARM namespace for ACS
	ContainerServiceNamespace = "Microsoft.ContainerService"
	// InsightsRPNamespace is ARM namespace for Insights
	InsightsRPNamespace = "Microsoft.Insights"
	// DiagnosticSettingResourceType is resource type name for diagnosticSetting
	DiagnosticSettingResourceType = "diagnosticSettings"
	// ManagedClusterResourceTypeName is name of ManagedCluster resource type
	ManagedClusterResourceTypeName = "ManagedClusters"
	// ContainerServiceType is the ARM type for ACS
	ContainerServiceType = ContainerServiceNamespace + "/ContainerServices"
	// ManagedClusterType is the ARM type for ACS
	ManagedClusterType = ContainerServiceNamespace + "/" + ManagedClusterResourceTypeName
	// AccessProfileType is a managed clusters proxy resource type
	AccessProfileType = ManagedClusterType + "/AccessProfiles"
	// DestinationUnderlay is the name of target or destination underlay (for ccpMigration)
	DestinationUnderlay = "DestinationUnderlay"
	// MigrateCustomerControlPlaneWithForce allows caller to override errors and force migration off of an unhealthy underlay
	MigrateCustomerControlPlaneWithForce = "MigrateCustomerControlPlaneWithForce"
	// CustomerControlPlaneID is id of CCP (customer control plane).
	CustomerControlPlaneID = "CustomerControlPlaneID"
	// SourceUnderlay is the name of source underlay (for ccpMigration)
	SourceUnderlay = "SourceUnderlay"

	// ExtensionAddonType is the ARM type for extension addons
	ExtensionAddonType = ManagedClusterType + "/ExtensionAddons"
	// PathExtensionAddonNameParameter is the extension addon parameter name used in route for extension addons
	PathExtensionAddonNameParameter = "extensionAddonName"

	// SnaptshotType is the ARM type for Snapshots
	SnapshotType = ContainerServiceNamespace + "/Snapshots"

	// MigrateClusterV2ApiserverSubnetId is the key for propertyBag for ApiserverSubnetId
	MigrateClusterV2ApiserverSubnetId = "MigrateClusterV2ApiserverSubnetId"
)

// subscription and common routes shared by containerservice and managedcluster resource.
const (
	// SubscriptionsURLPrefix is the base route prefix for all subscription based operations.
	SubscriptionsURLPrefix = "/" + SubscriptionsLiteral

	// InternalURLPrefix is the base route prefix for all internal operations.
	InternalURLPrefix = "/" + InternalLiteral

	// HealthCheckRoute is the base route for the health check
	HealthCheckRoute = "/healthz"

	// SubscriptionResourceOperationRoute is the route used to perform PUT/GET on one Subscription resource
	// /{subscriptionId}
	SubscriptionResourceOperationRoute = "/{" +
		PathSubscriptionIDParameter + "}"

	// SubscriptionResourceFullPath is the full path, prefix and operation route for
	// actions performed on subscription resources
	// /subscriptions/{subscriptionId}
	SubscriptionResourceFullPath = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute

	// ContainerServiceProviderFullPath is re-usable part to build other routes
	// providers/Microsoft.ContainerService
	ContainerServiceProviderPath = ProvidersLiteral + "/" + ContainerServiceProviderLiteral

	// OperationResultsResourceOperationRoute is the route used to perform GET on one asynchronous operation
	// using 202 Accepted and Location Headers to indicate progress
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#202-accepted-and-location-headers
	// {subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/operationresults/{operationId}
	OperationResultsResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral +
		"/" + LocationsLiteral + "/{" + PathLocationParameter +
		"}/" + OperationResultsLiteral + "/{" + PathOperationIDParameter + "}"

	// OperationResultsResourceFullPath is the full path, prefix and operation route for
	// actions performed on Operation resources
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/operationresults/{operationId}
	OperationResultsResourceFullPath = SubscriptionsURLPrefix + OperationResultsResourceOperationRoute

	// OperationStatusResourceOperationRoute is the route used to perform GET on one operation
	// {subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/operations/{operationId}
	OperationStatusResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral +
		"/" + LocationsLiteral + "/{" + PathLocationParameter +
		"}/" + OperationsLiteral + "/{" + PathOperationIDParameter + "}"

	// OperationStatusResourceFullPath is the full path, prefix and operation route for
	// actions performed on Operation resources
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/operations/{operationId}
	OperationStatusResourceFullPath = SubscriptionsURLPrefix + OperationStatusResourceOperationRoute

	// DeploymentResourceFullPath is the full path, prefix and operation route for
	// actions performed on Deployment preflight
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/deployments/{deploymentName}/preflight
	DeploymentResourceFullPath = SubscriptionsURLPrefix + DeploymentPreflightOperationRoute

	// DeploymentPreflightOperationRoute is the route used to perform POST on one deployment
	// /{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ContainerService/deployments/{deploymentName}/preflight
	DeploymentPreflightOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + DeploymentsLiteral + "/{" +
		PathDeploymentNameParameter + "}/" + PreflightLiteral

	// DeploymentPreflightFullPath is the full path, prefix and operation route for preflight operation
	// /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ContainerService/deployments/{deploymentName}/preflight
	DeploymentPreflightFullPath = SubscriptionsURLPrefix + DeploymentPreflightOperationRoute

	// AdminURLPrefix is the base path of admin operations
	AdminURLPrefix = "/" + AdminLiteral
	// APIV1Prefix is the base path of v1 api operations
	APIV1Prefix = "/" + APILiteral + "/" + ApiVersionV1

	// /providers/Microsoft.ContainerService/operations
	GetAvailableOperationsFullPath = "/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + OperationsLiteral
)

// containerservices resource operation routes
const (
	// ContainerServiceResourceOperationRoute is the route used to perform PUT/GET/DELETE on one container service resource
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}
	ContainerServiceResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ContainerServicesLiteral + "/{" +
		PathResourceNameParameter + "}"

	// ContainerServiceResourceFullPath is the full path, prefix and operation route for
	// actions performed on ContainerService resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}
	ContainerServiceResourceFullPath = SubscriptionsURLPrefix + ContainerServiceResourceOperationRoute

	// ListContainerServiceResourcesBySubscriptionOperationRoute is the route used to perform GET to list containerServices in a subscription
	// /{subscriptionId}/providers/Microsoft.ContainerService/containerServices
	ListContainerServiceResourcesBySubscriptionOperationRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral +
		"/" + ContainerServicesLiteral

	// ListContainerServiceResourcesByResourceGroupOperationRoute is the route used to perform GET to list containerServices in a resourcegroup
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices
	ListContainerServiceResourcesByResourceGroupOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ContainerServicesLiteral

	// ListContainerServiceResourcesBySubscriptionFullPath is the full path, prefix and operation route for
	// list actions performed on ContainerService under subscription scope
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices
	ListContainerServiceResourcesBySubscriptionFullPath = SubscriptionsURLPrefix + ListContainerServiceResourcesBySubscriptionOperationRoute

	// ListContainerServiceResourcesByResourceGroupFullPath is the full path, prefix and operation route for
	// list actions performed on ContainerService under subscription/resourceGroup scope
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices
	ListContainerServiceResourcesByResourceGroupFullPath = SubscriptionsURLPrefix + ListContainerServiceResourcesByResourceGroupOperationRoute

	// ListOrchestratorsBySubscriptionRoute is the route used to GET available orchestrators in a subscription
	// /{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators
	ListOrchestratorsBySubscriptionRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral + "/" + LocationsLiteral + "/{" + PathLocationParameter +
		"}/" + OrchestratorsLiteral

	// GetOSOptionsBySubscriptionRoute is the route used to GET os options in a subscription
	// /{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/osOptions
	GetOSOptionsBySubscriptionRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral + "/" + LocationsLiteral + "/{" + PathLocationParameter +
		"}/" + OSOptionsLiteral + "/" + DefaultLiteral

	// GetContainerServiceUpgradeProfileRoute is the route used to GET available upgrade versions for container service
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}/upgradeprofiles/default
	GetContainerServiceUpgradeProfileRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ContainerServicesLiteral + "/{" +
		PathResourceNameParameter + "}" + "/" + UpgradeProfilesLiteral + "/" + DefaultLiteral

	// ListOrchestratorsBySubscriptionFullPath is the full path, prefix and operation route for list of supported orchestrators
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators
	ListOrchestratorsBySubscriptionFullPath = SubscriptionsURLPrefix + ListOrchestratorsBySubscriptionRoute

	// GetOSOptionsBySubscriptionFullPath is the full path, prefix and operation route for get os options
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/osOptions
	GetOSOptionsBySubscriptionFullPath = SubscriptionsURLPrefix + GetOSOptionsBySubscriptionRoute

	// GetContainerServiceUpgradeProfileFullPath is the full path, prefix and operation route for list of available orchestrators upgrades
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}/upgradeprofiles/upgradeprofile
	GetContainerServiceUpgradeProfileFullPath = SubscriptionsURLPrefix + GetContainerServiceUpgradeProfileRoute

	// ListManagedClusterCredentialFullPath is the full path, prefix and operation route for list credential
	ListManagedClusterCredentialFullPath                      = SubscriptionsURLPrefix + ListManagedClusterCredentialRoute
	ListManagedClusterClusterAdminCredentialFullPath          = SubscriptionsURLPrefix + ListManagedClusterClusterAdminCredentialRoute
	ListManagedClusterClusterUserCredentialFullPath           = SubscriptionsURLPrefix + ListManagedClusterClusterUserCredentialRoute
	ListManagedClusterClusterMonitoringUserCredentialFullPath = SubscriptionsURLPrefix + ListManagedClusterClusterMonitoringUserCredentialRoute
)

const (
	// InternalSubscriptionResourceFullPath is the full path, prefix and operation route for
	// Internal actions performed on subscription resources
	// /internal/subscriptions/{subscriptionId}
	InternalSubscriptionResourceFullPath = InternalURLPrefix + SubscriptionsURLPrefix + SubscriptionResourceOperationRoute

	// InternalContainerServiceResourceFullPath is the full path, prefix and operation route for
	// Internal actions performed on ContainerService resources
	// /internal/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}
	InternalContainerServiceResourceFullPath = InternalURLPrefix + SubscriptionsURLPrefix + ContainerServiceResourceOperationRoute

	// InternalOperationStatusResourceFullPath is the full path, prefix and operation route for
	// Internalactions performed on Operation resources
	// /internal/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}/operations/{operationId}
	InternalOperationStatusResourceFullPath = InternalURLPrefix + SubscriptionsURLPrefix + InternalOperationStatusResourceOperationRoute

	// InternalOperationStatusResourceOperationRoute is the route used to perform PUT on one operation
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{resourceName}/operations/{operationId}
	InternalOperationStatusResourceOperationRoute = SubscriptionResourceOperationRoute + "/" +
		ResourceGroupsLiteral + "/{" + PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral +
		"/" + ContainerServicesLiteral + "/{" + PathResourceNameParameter + "}" +
		"}/" + OperationsLiteral + "/{" + PathOperationIDParameter + "}"
)

// managedclusters resource operation routes
const (
	// ManagedClusterResourceOperationRoute is the route used to perform PUT/GET/DELETE on one hosted control plane resource
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}
	ManagedClusterResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}"

	// ManagedClusterResourceFullPath is the full path, prefix and operation route for
	// actions performed on ContainerService resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}
	ManagedClusterResourceFullPath = SubscriptionsURLPrefix + ManagedClusterResourceOperationRoute

	ManagedClusterRunCommandRoute          = ManagedClusterResourceOperationRoute + "/" + RunCommandLiteral
	ManagedClusterRunCommandResultRoute    = ManagedClusterResourceOperationRoute + "/" + CommandResultLiteral + "/{" + PathRunCommandIdParameter + "}"
	ManagedClusterRunCommandFullPath       = SubscriptionsURLPrefix + ManagedClusterRunCommandRoute
	ManagedClusterRunCommandResultFullPath = SubscriptionsURLPrefix + ManagedClusterRunCommandResultRoute

	StopManagedClusterOperationRoute  = ManagedClusterResourceOperationRoute + "/stop"
	StartManagedClusterOperationRoute = ManagedClusterResourceOperationRoute + "/start"
	StopManagedClusterFullPath        = SubscriptionsURLPrefix + StopManagedClusterOperationRoute
	StartManagedClusterFullPath       = SubscriptionsURLPrefix + StartManagedClusterOperationRoute

	//MaintenanceConfigurationOperationRoute is the route used to perform PUT/GET/DELETE operation on the maintenance configuration.
	MaintenanceConfigurationOperationRoute = ManagedClusterResourceOperationRoute + "/" + MaintenanceConfigurationsLiteral + "/{" +
		PathMaintenanceConfigurationNameParameter + "}"

	// MaintenanceConfigurationFullPath is the full path, prefix and operation route for
	// actions performed on MaintenanceConfiguration resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/maintenanceConfigurations/{maintenanceConfigurationName}
	MaintenanceConfigurationFullPath = SubscriptionsURLPrefix + MaintenanceConfigurationOperationRoute

	//AgentPoolResourceOperationRoute is the route used to perform PUT/GET/DELETE operation on the agentpool.
	AgentPoolResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" +
		ManagedClustersLiteral + "/{" + PathResourceNameParameter + "}/" + AgentPoolsLiteral + "/{" +
		PathAgentPoolNameParameter + "}"

	// AgentPoolResourceFullPath is the full path, prefix and operation route for
	// actions performed on AgentPool resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools/{agentPoolName}
	AgentPoolResourceFullPath = SubscriptionsURLPrefix + AgentPoolResourceOperationRoute

	//UpgradeAgentPoolNodeImageOperationRoute is the route used to perform POST upgrade node image version operation on the agentpool.
	UpgradeAgentPoolNodeImageOperationRoute = AgentPoolResourceOperationRoute + "/upgradeNodeImageVersion"

	// UpgradeAgentPoolNodeImageFullPath is the full path, prefix and operation route for node image upgrade operations.
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools/{agentPoolName}/upgradeNodeImageVersion
	UpgradeAgentPoolNodeImageFullPath = SubscriptionsURLPrefix + UpgradeAgentPoolNodeImageOperationRoute

	// LinkedNotificationOperationRoute is routing path to linked notification for extension resource to managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/providers/{extensionProvider}/{extensionResourceType}/{extensionResourceName}/providers/microsoft.ContainerService/notify
	LinkedNotificationOperationRoute = ManagedClusterResourceOperationRoute + "/" +
		ProvidersLiteral + "/{" + PathExtensionProviderParameter + "}/{" + PathExtensionResourceTypeParameter + "}/{" + PathExtensionResourceNameParameter + "}/" +
		ContainerServiceProviderPath + "/" + NotifyLiteral

	// LinkedNotificationFullPath is full path to linked notification for extension resource to managed cluster
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/providers/{extensionProvider}/{extensionResourceType}/{extensionResourceName}/providers/microsoft.ContainerService/notify
	LinkedNotificationFullPath = SubscriptionsURLPrefix + LinkedNotificationOperationRoute

	// GetAgentPoolUpgradeProfileRoute is the route used to GET available upgrade versions for agent pool
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeprofiles/default
	GetAgentPoolUpgradeProfileRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" +
		ManagedClustersLiteral + "/{" + PathResourceNameParameter + "}/" +
		AgentPoolsLiteral + "/{" + PathAgentPoolNameParameter + "}/" + UpgradeProfilesLiteral + "/" + DefaultLiteral

	// GetAgentPoolUpgradeProfileRouteFullPath is the full path, prefix and operation route for
	// actions performed on getting upgrade profiles for AgentPool resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools/{agentPoolName}/upgradeprofiles/default
	GetAgentPoolUpgradeProfileFullPath = SubscriptionsURLPrefix + GetAgentPoolUpgradeProfileRoute

	// ListAgentPoolAvailableVersionsRoute is the route used to GET available upgrade versions for agent pool
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableagentpoolversions
	ListAgentPoolAvailableVersionsRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" +
		ManagedClustersLiteral + "/{" + PathResourceNameParameter + "}/" + AvailableAgentPoolVersionsLiteral

	// ListAgentPoolAvailableVersionsFullPath is the full path, prefix and operation route for
	// actions performed on getting orchestrators for AgentPool resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/availableagentpoolversions
	ListAgentPoolAvailableVersionsFullPath = SubscriptionsURLPrefix + ListAgentPoolAvailableVersionsRoute

	// ListManagedClusterResourcesBySubscriptionOperationRoute is the route used to perform GET to list managedclusters in a subscription
	// {subscriptionId}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesBySubscriptionOperationRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral +
		"/" + ManagedClustersLiteral

	// ListManagedClusterResourcesByResourceGroupOperationRoute is the route used to perform GET to list ManagedClusters in a resourcegroup
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesByResourceGroupOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral

	// ListManagedClusterResourcesBySubscriptionFullPath is the full path, prefix and operation route for
	// list actions performed on ManagedCluster under subscription scope
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesBySubscriptionFullPath = SubscriptionsURLPrefix + ListManagedClusterResourcesBySubscriptionOperationRoute

	// ListManagedClusterResourcesByResourceGroupFullPath is the full path, prefix and operation route for
	// list actions performed on ManagedCluster under subscription/resourceGroup scope
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesByResourceGroupFullPath = SubscriptionsURLPrefix + ListManagedClusterResourcesByResourceGroupOperationRoute

	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools
	ListAgentPoolsByClusterOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" +
		ManagedClustersLiteral + "/{" + PathResourceNameParameter + "}/" + AgentPoolsLiteral

	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools
	ListAgentPoolsByClusterFullPath = SubscriptionsURLPrefix + ListAgentPoolsByClusterOperationRoute

	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/maintenanceConfigurations
	ListMaintenanceConfigurationsByManagedClusterOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" +
		ManagedClustersLiteral + "/{" + PathResourceNameParameter + "}/" + MaintenanceConfigurationsLiteral

	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/maintenanceConfigurations
	ListMaintenanceConfigurationsByManagedClusterFullPath = SubscriptionsURLPrefix + ListMaintenanceConfigurationsByManagedClusterOperationRoute

	// GetManagedClusterUpgradeProfileRoute is the route used to GET available upgrade versions for managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeprofiles/default
	GetManagedClusterUpgradeProfileRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + UpgradeProfilesLiteral + "/" + DefaultLiteral

	// GetManagedClusterUpgradeProfileFullPath is the full path, prefix and operation route for list of available orchestrators upgrades
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeprofiles/default
	GetManagedClusterUpgradeProfileFullPath = SubscriptionsURLPrefix + GetManagedClusterUpgradeProfileRoute

	// GetManagedClusterDiagnosticsStateRoute is the route used to GET the diagnostics state for the managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/diagnosticsstate/default
	GetManagedClusterDiagnosticsStateRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + DiagnosticsStateLiteral + "/" + DefaultLiteral

	// GetManagedClusterDiagnosticsStateFullPath is the full path, prefix and operation route for diagnostics state for the managed cluster
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/diagnosticsstate/default
	GetManagedClusterDiagnosticsStateFullPath = SubscriptionsURLPrefix + GetManagedClusterDiagnosticsStateRoute

	// GetManagedClusterAccessProfileRoute is the route used to GET accessProfile for the roleName for managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/accessProfiles/{roleName}
	GetManagedClusterAccessProfileRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + AccessProfilesLiteral + "/{" + PathAccessProfileParameter + "}"

	// GetManagedClusterAccessProfileFullPath is the full path, prefix and operation for GET accessProfile for the roleName for managed cluster
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/accessProfiles/{roleName}
	GetManagedClusterAccessProfileFullPath = SubscriptionsURLPrefix + GetManagedClusterAccessProfileRoute

	// ListManagedClusterCredentialRoute is the route used to GET accessProfile for the roleName of a managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/accessProfiles/{roleName}/listCredential
	ListManagedClusterCredentialRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral +
		"/{" + PathResourceGroupNameParameter + "}/" + ProvidersLiteral +
		"/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral +
		"/{" + PathResourceNameParameter + "}/" + AccessProfilesLiteral +
		"/{" + PathAccessProfileParameter + "}/" + ListCredentialLiteral

	// ListManagedClusterClusterAdminCredentialRoute is the route used to GET accessProfile for clusterAdmin role of a managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/listClusterAdminCredential
	ListManagedClusterClusterAdminCredentialRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral +
		"/{" + PathResourceGroupNameParameter + "}/" + ProvidersLiteral +
		"/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral +
		"/{" + PathResourceNameParameter + "}/" + ListClusterAdminCredentialLiteral

	// ListManagedClusterClusterUserCredentialRoute is the route used to GET accessProfile for clusterUser role of a managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/listClusterUserCredential
	ListManagedClusterClusterUserCredentialRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral +
		"/{" + PathResourceGroupNameParameter + "}/" + ProvidersLiteral +
		"/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral +
		"/{" + PathResourceNameParameter + "}/" + ListClusterUserCredentialLiteral

	// ListManagedClusterClusterMonitoringUserCredentialRoute is the route used to GET accessProfile for clusterMonitoringUser role of a managed cluster
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/listClusterMonitoringUserCredential
	ListManagedClusterClusterMonitoringUserCredentialRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral +
		"/{" + PathResourceGroupNameParameter + "}/" + ProvidersLiteral +
		"/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral +
		"/{" + PathResourceNameParameter + "}/" + ListClusterMonitoringUserCredentialLiteral

	// ResetServicePrincipalProfileRoute is the route used to update the service principal profile for a managed cluster.
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/resetServicePrincipalProfile
	ResetServicePrincipalProfileRoute = ManagedClusterResourceOperationRoute + "/resetServicePrincipalProfile"

	// ResetServicePrincipalProfileFullPath is the full path, prefix and operation route for service principal profile update operations.
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/resetServicePrincipalProfile
	ResetServicePrincipalProfileFullPath = SubscriptionsURLPrefix + ResetServicePrincipalProfileRoute

	// ResetAADProfileRoute is the route used to update the AAD server profile for a managed cluster.
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/resetAADProfile
	ResetAADProfileRoute = ManagedClusterResourceOperationRoute + "/resetAADProfile"

	// ResetServicePrincipalProfileFullPath is the full path, prefix and operation route for AAD profile update operations.
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/resetAADProfile
	ResetAADProfileFullPath = SubscriptionsURLPrefix + ResetAADProfileRoute

	// ListDetectorsRoute is the route used to list app lens detectors
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/detectors
	ListDetectorsRoute    = ManagedClusterResourceOperationRoute + "/" + DetectorsLiteral
	ListDetectorsFullPath = SubscriptionsURLPrefix + ListDetectorsRoute

	// GetDetectorRoute is the route used to get app lens detector
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/detectors/{detectorName}
	GetDetectorRoute    = ManagedClusterResourceOperationRoute + "/" + DetectorsLiteral + "/{" + PathDetectorNameParameter + "}"
	GetDetectorFullPath = SubscriptionsURLPrefix + GetDetectorRoute

	// RotateClusterCertificates is the route used to rotate cluster certificates for a managed cluster.
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/rotateClusterCertificates
	RotateClusterCertificatesRoute = ManagedClusterResourceOperationRoute + "/rotateClusterCertificates"

	// RotateClusterCertificatesFullPath is the full path, prefix and operation route for cluster certificates rotate operations.
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/rotateClusterCertificates
	RotateClusterCertificatesFullPath = SubscriptionsURLPrefix + RotateClusterCertificatesRoute

	// ListPrivateEndpointConnectionsRoute is the route used to list private endpoint connections
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections
	ListPrivateEndpointConnectionsRoute       = ManagedClusterResourceOperationRoute + "/" + PrivateEndpointConnectionsLiteral
	ListPrivateLinkServiceConnectionsFullPath = SubscriptionsURLPrefix + ListPrivateEndpointConnectionsRoute

	// PrivateEndpointConnectionsRoute is the route used to update private endpoint connection
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections/{privateendpointconnectionName}
	PrivateEndpointConnectionRoute       = ManagedClusterResourceOperationRoute + "/" + PrivateEndpointConnectionsLiteral + "/{" + PathPrivateEndpointConnectionNameParameter + "}"
	PrivateLinkServiceConnectionFullPath = SubscriptionsURLPrefix + PrivateEndpointConnectionRoute

	// ListPrivateLinkResourcesRoute is the route used to list private link resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateLinkResources
	ListPrivateLinkResourcesRoute    = ManagedClusterResourceOperationRoute + "/" + PrivateLinkResourcesLiteral
	ListPrivateLinkResourcesFullPath = SubscriptionsURLPrefix + ListPrivateLinkResourcesRoute

	// ResolvePrivateLinkServiceIDRoute is the route used to list private link resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/resolvePrivateLinkServiceId
	ResolvePrivateLinkServiceIDRoute    = ManagedClusterResourceOperationRoute + "/" + ResolvePrivateLinkServiceIDLiteral
	ResolvePrivateLinkServiceIDFullPath = SubscriptionsURLPrefix + ResolvePrivateLinkServiceIDRoute

	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/extensionaddons
	ListExtensionAddonsByManagedClusterRoute    = ManagedClusterResourceOperationRoute + "/" + ExtensionAddonsLiteral
	ListExtensionAddonsByManagedClusterFullPath = SubscriptionsURLPrefix + ListExtensionAddonsByManagedClusterRoute

	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/extensionaddons/{extensionAddonName}
	ExtensionAddonOperationsRoute = ManagedClusterResourceOperationRoute + "/" + ExtensionAddonsLiteral + "/{" +
		PathExtensionAddonNameParameter + "}"
	ExtensionAddonFullPath = SubscriptionsURLPrefix + ExtensionAddonOperationsRoute

	// ListOutboundNetworkDependenciesEndpointsRoute is the route used to list all outboundNetworkDependenciesEndpoints
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/outboundNetworkDependenciesEndpoints
	ListOutboundNetworkDependenciesEndpointsRoute    = ManagedClusterResourceOperationRoute + "/" + OutboundNetworkDependenciesEndpoints
	ListOutboundNetworkDependenciesEndpointsFullPath = SubscriptionsURLPrefix + ListOutboundNetworkDependenciesEndpointsRoute

	// SnapshotResourceOperationRoute is the route used to perform PUT/GET/DELETE on one snapshot resource
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/Snapshots/{resourceName}
	SnapshotResourceOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + SnapshotsLiteral + "/{" +
		PathResourceNameParameter + "}"

	// SnapshotResourceFullPath is the full path, prefix and operation route for actions performed on Snapshot resources
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/Snapshots/{resourceName}
	SnapshotResourceFullPath = SubscriptionsURLPrefix + SnapshotResourceOperationRoute

	// ListSnapshotResourcesBySubscriptionOperationRoute is the route used to perform GET to list Snapshots in a subscription
	// {subscriptionId}/providers/Microsoft.ContainerService/Snapshots
	ListSnapshotResourcesBySubscriptionOperationRoute = SubscriptionResourceOperationRoute + "/" + ProvidersLiteral + "/" +
		ContainerServiceProviderLiteral +
		"/" + SnapshotsLiteral

	// ListSnapshotResourcesByResourceGroupOperationRoute is the route used to perform GET to list Snapshots in a resourcegroup
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/Snapshots
	ListSnapshotResourcesByResourceGroupOperationRoute = SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter +
		"}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + SnapshotsLiteral

	// ListSnapshotResourcesBySubscriptionFullPath is the full path, prefix and operation route for
	// list actions performed on Snapshot under subscription scope
	// /subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/Snapshots
	ListSnapshotResourcesBySubscriptionFullPath = SubscriptionsURLPrefix + ListSnapshotResourcesBySubscriptionOperationRoute

	// ListSnapshotResourcesByResourceGroupFullPath is the full path, prefix and operation route for
	// list actions performed on Snapshot under subscription/resourceGroup scope
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/Snapshots
	ListSnapshotResourcesByResourceGroupFullPath = SubscriptionsURLPrefix + ListSnapshotResourcesByResourceGroupOperationRoute

	// GetSnapshotResourceAdminOperationRoute is the route used to get snapshot from acis
	GetSnapshotResourceAdminOperationRoute = SubscriptionsURLPrefix + SnapshotResourceOperationRoute

	// GetSnapshotResourceAdminOperationFullPath is the full path to get snapshot from acis
	// /admin/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/Snapshots/{resourceName}
	GetSnapshotResourceAdminOperationFullPath = AdminURLPrefix + GetSnapshotResourceAdminOperationRoute

	// MigrateClusterV2OperationRoute is the route used to migrate cluster to v2
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/migrateClusterV2
	MigrateClusterV2OperationRoute         = ManagedClusterResourceOperationRoute + "/" + MigrateClusterV2Literal
	MigrateClusterV2OperationRouteFullPath = SubscriptionsURLPrefix + MigrateClusterV2OperationRoute
)

// this is where admin operation routes are defined
const (
	// Agent Pool Operations
	// GetAgentPoolResourceAdminOperationRoute is the route used to get agent pool resource from acis
	GetAgentPoolResourceAdminOperationRoute = SubscriptionsURLPrefix + AgentPoolResourceOperationRoute

	// ListAgentPoolsByClusterAdminOperationRoute is the route used to list agent pool resources by cluster from acis
	ListAgentPoolsByClusterAdminOperationRoute = SubscriptionsURLPrefix + ListAgentPoolsByClusterOperationRoute

	// PutAgentPoolResourceAdminOperationRoute is the route used to reconcile agent pool resource from acis
	PutAgentPoolResourceAdminOperationRoute = SubscriptionsURLPrefix + AgentPoolResourceOperationRoute

	// GetAgentPoolResourcesAdminOperationFullPath is the full path to get agent pool from acis
	GetAgentPoolResourcesAdminOperationFullPath = AdminURLPrefix + GetAgentPoolResourceAdminOperationRoute

	// ListAgentPoolsByClusterAdminOperationFullPath is the full path to list agent pools by cluster from acis
	ListAgentPoolsByClusterAdminOperationFullPath = AdminURLPrefix + ListAgentPoolsByClusterAdminOperationRoute

	// Managed Cluster Operations
	// GetManagedClusterResourceAdminOperationRoute is the route used to get managed cluster from acis
	GetManagedClusterResourceAdminOperationRoute = SubscriptionsURLPrefix + ManagedClusterResourceOperationRoute

	// PutManagedClusterResourceAdminOperationRoute is the route used to put managed cluster from acis
	PutManagedClusterResourceAdminOperationRoute = SubscriptionsURLPrefix + ManagedClusterResourceOperationRoute

	// BackfillManagedClusterAdminOperationRoute is the route used to back fill a managed cluster from Azure resources.
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/backfillmanagedcluster
	BackfillManagedClusterAdminOperationRoute = SubscriptionsURLPrefix + ManagedClusterResourceOperationRoute + "/" + BackfillManagedClusterLiteral

	// OperationStatusResourceAdminOperationRoute is the route used to get operation status from acis
	OperationStatusResourceAdminOperationRoute = SubscriptionsURLPrefix + OperationStatusResourceOperationRoute

	// ListManagedClusterResourcesBySubscriptionAdminOperationRoute is the route used to perform list managedclusters in a subscription
	// /{subscriptionId}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesBySubscriptionAdminOperationRoute = SubscriptionsURLPrefix + ListManagedClusterResourcesBySubscriptionOperationRoute

	// ListManagedClusterResourcesByResourceGroupAdminOperationRoute is the route used to perform list managedclusters in a resource group
	// /{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters
	ListManagedClusterResourcesByResourceGroupAdminOperationRoute = SubscriptionsURLPrefix + ListManagedClusterResourcesByResourceGroupOperationRoute

	// ListCustomerControlPlanePodsOperationRoute is the route used to list the customer control plane pods
	ListCustomerControlPlanePodsOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + PodsLiteral

	// GetCustomerControlPlanePodLogOperationRoute is the route used to get customer control plane pod log
	GetCustomerControlPlanePodLogOperationRoute = ListCustomerControlPlanePodsOperationRoute + "/{" + PathPodNameParameter + "}/" + ContainersLiteral +
		"/{" + PathContainerNameParameter + "}/" + LogLiteral

	//	GetCustomerControlPlaneEventsOperationRoute is the route used to list the customer control plane events
	GetCustomerControlPlaneEventsOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + EventLiteral

	// ListUnderlaysOperationRoute is the route used to list the customer control plane events
	ListUnderlaysOperationRoute = "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral

	// PostUnderlayKubectlCommandOperationRoute is the route used to post generic kubectl command given an underlay
	PostUnderlayKubectlCommandOperationRoute = "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral +
		"/{" + PathUnderlayNameParameter + "}/" + KubectlLiteral

	// JsonPatchAgentPoolCommandOperationRoute is the route to change the agent pool json properties
	JsonPatchAgentPoolCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + AgentPoolsLiteral + "/{" + PathAgentPoolNameParameter + "}/" + JsonPatchLiteral

	// JsonPatchAgentPoolCommandOperationRoute is the route to change the agent pool json properties
	JsonPatchManagedClusterCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + JsonPatchLiteral

	// PostKubectlCommandOperationRoute is the route used to post generic kubectl command
	PostKubectlCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + KubectlLiteral

	// PostOverlayKubectlCommandOperationRoute is the route used to post generic kubectl command
	PostOverlayKubectlCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + OverlayLiteral + "/" + KubectlLiteral

	// PostVirtualMachineRunCommandOperationRoute is the route to post run command on vm
	PostVirtualMachineRunCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + NodeLiteral + "/{" + PathVMNameParameter + "}/" + ActionsLiteral + "/{" + PathActionNameParameter + "}"

	// PostVirtualMachineGenericRunCommandOperationRoute is the route to post generic run command on vm
	PostVirtualMachineGenericRunCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + NodeLiteral + "/{" + PathVMNameParameter + "}/" + RunCommandLiteral

	// ReimageManagedClusterLiteralOperationRoute is the route used reimage a customer pool
	ReimageManagedClusterCommandOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + AgentPoolsLiteral + "/{" + PathAgentPoolNameParameter + "}/" + ReimageManagedClusterLiteral

	// DelegateSubnetOperationRoute is the route used to delegate a subnet to AKS
	DelegateSubnetOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + NetworkProviderLiteral + "/" + VirtualNetworksLiteral + "/{" +
		PathVirtualNetworkNameParameter + "}/" + SubnetsLiteral + "/{" + PathSubnetNameParameter + "}/" + DelegateSubnetLiteral

	// UnDelegateSubnetOperationRoute is the route used to undelegate a subnet from AKS
	UnDelegateSubnetOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + NetworkProviderLiteral + "/" + VirtualNetworksLiteral + "/{" +
		PathVirtualNetworkNameParameter + "}/" + SubnetsLiteral + "/{" + PathSubnetNameParameter + "}/" + UnDelegateSubnetLiteral

	// GetCustomerAgentNodesStatusOperationRoute is the route to get customer agent nodes status
	GetCustomerAgentNodesStatusOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + NodeLiteral

	// MigrateCustomerControlPlaneOperationRoute is the route used to migrate a CCP between underlays
	MigrateCustomerControlPlaneOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + MigrateCustomerControlPlaneLiteral

	// DeallocateControlPlaneOperationRoute is the route used to deallocate a CCP
	DeallocateControlPlaneOperationRoute = SubscriptionsURLPrefix + SubscriptionResourceOperationRoute + "/" + ResourceGroupsLiteral + "/{" +
		PathResourceGroupNameParameter + "}/" + ProvidersLiteral + "/" + ContainerServiceProviderLiteral + "/" + ManagedClustersLiteral + "/{" +
		PathResourceNameParameter + "}/" + DeallocateControlPlaneLiteral

	// DrainCustomerControlPlanesOperationRoute is the route used to drain all CCPs off of an underlay
	DrainCustomerControlPlanesOperationRoute = "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral + "/{" + PathUnderlayNameParameter + "}/" + DrainCustomerControlPlanesLiteral

	// OperationStatusResourceAdminOperationFullPath is the full path to get operation status from acis
	OperationStatusResourceAdminOperationFullPath = AdminURLPrefix + OperationStatusResourceAdminOperationRoute

	// ListManagedClusterResourcesBySubscriptionAdminOperationFullPath is the full path to list managed cluster by subscription from acis
	ListManagedClusterResourcesBySubscriptionAdminOperationFullPath = AdminURLPrefix + ListManagedClusterResourcesBySubscriptionAdminOperationRoute

	// ListManagedClusterResourcesByResourceGroupAdminOperationFullPath is the full path to list managed cluster by resource group from acis
	ListManagedClusterResourcesByResourceGroupAdminOperationFullPath = AdminURLPrefix + ListManagedClusterResourcesByResourceGroupAdminOperationRoute

	// GetManagedClusterResourceAdminOperationFullPath is the full path to get managed cluster from acis
	GetManagedClusterResourceAdminOperationFullPath = AdminURLPrefix + GetManagedClusterResourceAdminOperationRoute

	// ListCustomerControlPlanePodsOperationFullPath is the full path to list customer control plane pods
	ListCustomerControlPlanePodsOperationFullPath = AdminURLPrefix + ListCustomerControlPlanePodsOperationRoute

	// GetCustomerControlPlanePodLogOperationFullPath is the full path to get customer control plane pod log
	GetCustomerControlPlanePodLogOperationFullPath = AdminURLPrefix + GetCustomerControlPlanePodLogOperationRoute

	// GetCustomerControlPlaneEventsOperationFullPath is the full path to list customer control plane events
	GetCustomerControlPlaneEventsOperationFullPath = AdminURLPrefix + GetCustomerControlPlaneEventsOperationRoute

	// JsonPatchAgentPoolAdminOperationFullPath is the full path to change the json properties of a Agent Pool
	JsonPatchAgentPoolAdminOperationFullPath = AdminURLPrefix + JsonPatchAgentPoolCommandOperationRoute

	// JsonPatchManagedClusterAdminOperationFullPath is the full path to change the json properties of a Managed Cluster
	JsonPatchManagedClusterAdminOperationFullPath = AdminURLPrefix + JsonPatchManagedClusterCommandOperationRoute

	// BackfillManagedClusterOperationFullPath is the full path to back fill managed cluster from Azure resource
	BackfillManagedClusterOperationFullPath = AdminURLPrefix + BackfillManagedClusterAdminOperationRoute

	// PostKubectlCommandOperationFullPath is the full path to post generic kubectl command
	PostKubectlCommandOperationFullPath = AdminURLPrefix + PostKubectlCommandOperationRoute

	// PostOverlayKubectlCommandOperationFullPath is the full path to post generic kubectl command
	PostOverlayKubectlCommandOperationFullPath = AdminURLPrefix + PostOverlayKubectlCommandOperationRoute

	// ListUnderlaysOperationFullPath is the full path to post generic kubectl command given an underlay
	ListUnderlaysOperationFullPath = AdminURLPrefix + ListUnderlaysOperationRoute

	// PostUnderlayKubectlCommandOperationFullPath is the full path to post generic kubectl command given an underlay
	PostUnderlayKubectlCommandOperationFullPath = AdminURLPrefix + PostUnderlayKubectlCommandOperationRoute

	// GetSubscriptionResourceOperationFullPath is the full path used to get subscription resource
	GetSubscriptionResourceOperationFullPath = AdminURLPrefix + SubscriptionResourceFullPath

	// PostVirtualMachineRunCommandOperationFullPath is the full path to post vm runCommand
	PostVirtualMachineRunCommandOperationFullPath = AdminURLPrefix + PostVirtualMachineRunCommandOperationRoute

	// ReimageManagedClusterAdminOperationFullPath is the full path to reimage a managed cluster's vm
	ReimageManagedClusterAdminOperationFullPath = AdminURLPrefix + ReimageManagedClusterCommandOperationRoute

	// DelegateSubnetAdminOperationFullPath is the full path to delegate a subnet
	DelegateSubnetAdminOperationFullPath = AdminURLPrefix + DelegateSubnetOperationRoute

	// UnDelegateSubnetAdminOperationFullPath is the full path to undelegate a subnet
	UnDelegateSubnetAdminOperationFullPath = AdminURLPrefix + UnDelegateSubnetOperationRoute

	// PostVirtualMachineGenericRunCommandOperationFullPath is the full path to post generic vm runCommand
	PostVirtualMachineGenericRunCommandOperationFullPath = AdminURLPrefix + PostVirtualMachineGenericRunCommandOperationRoute

	// GetCustomerAgentNodesStatusOperationFullPath is the full path to get customer agent nodes status
	GetCustomerAgentNodesStatusOperationFullPath = AdminURLPrefix + GetCustomerAgentNodesStatusOperationRoute

	// AdminLinkedNotificationOperationRoute is admin routing to reconcile diagnosticSettings for managedCluster
	// We re-used linkedNotificationOperation
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/providers/{extensionProvider}/{extensionResourceType}/{extensionResourceName}/providers/microsoft.ContainerService/notify
	AdminLinkedNotificationOperationRoute = SubscriptionsURLPrefix + LinkedNotificationOperationRoute

	// AdminLinkedNotificationFullPath is full path to admin operation to reconcile diagnosticSettings for managedCluster
	// /admin/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/providers/{extensionProvider}/{extensionResourceType}/{extensionResourceName}/providers/microsoft.ContainerService/notify
	AdminLinkedNotificationFullPath = AdminURLPrefix + AdminLinkedNotificationOperationRoute

	// JsonPatchControlPlaneCommandOperationRoute is the route to change the agent pool json properties
	JsonPatchControlPlaneCommandOperationRoute = APIV1Prefix + "/" + ControlPlanesLiteral + "/{" + PathControlPlaneParameter + "}/" + JsonPatchLiteral
	// JsonPatchControlPlaneAdminOperationFullPath is the full path to change the json properties of a Control Plane
	JsonPatchControlPlaneAdminOperationFullPath = AdminURLPrefix + JsonPatchControlPlaneCommandOperationRoute

	// AdminListUnderlayRoute is the admin routing to list underlay configs or put underlay config
	AdminListUnderlayRoute = APIV1Prefix + "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral
	// AdminListUnderlayFullPath is the full path to admin operation to list underlay configs or put underlay config
	AdminListUnderlayFullPath = AdminURLPrefix + AdminListUnderlayRoute

	// AdminExpandUnderlayCapacityRoute is the admin routing to expand regional underlay capacity
	AdminExpandUnderlayCapacityRoute = APIV1Prefix + "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + ExpandUnderlayCapacityLiteral
	// AdminExpandUnderlayCapacityFullPath is the full path to admin operation to expand regional underlay capacity
	AdminExpandUnderlayCapacityFullPath = AdminURLPrefix + AdminExpandUnderlayCapacityRoute

	// AdminUnderlayRoute is the admin routing to put/delete underlay config by underlay name
	AdminUnderlayRoute = APIV1Prefix + "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral + "/{" + PathUnderlayNameParameter + "}"
	// AdminUnderlayFullPath is the full path to admin operation to put/delete underlay config by underlay name
	AdminUnderlayFullPath = AdminURLPrefix + AdminUnderlayRoute

	// AdminGetUnderlayRoute is the admin routing to get underlay config by underlay name
	AdminGetUnderlayRoute = APIV1Prefix + "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral + "/{" + PathUnderlayNameParameter + "}/" + UnderlayDataSourcesLiteral + "/{" + PathUnderlayDataSourceParameter + "}"
	// AdminGetUnderlayFullPath is the full path to admin operation to get config by underlay name
	AdminGetUnderlayFullPath = AdminURLPrefix + AdminGetUnderlayRoute

	// AdminPostUnderlayActionRoute is the admin routing to update a underlay config quarantined value
	AdminPostUnderlayActionRoute = APIV1Prefix + "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + UnderlaysLiteral +
		"/{" + PathUnderlayNameParameter + "}/{" + PathUnderlayActionParameter + "}"
	// AdminPostUnderlayActionFullPath is the full path to admin operation to update a underlay config quarantined value
	AdminPostUnderlayActionFullPath = AdminURLPrefix + AdminPostUnderlayActionRoute
	// MigrateCustomerControlPlaneAdminOperationFullPath is the full path to migrate a CCP between underlays
	MigrateCustomerControlPlaneAdminOperationFullPath = AdminURLPrefix + MigrateCustomerControlPlaneOperationRoute
	// DeallocateControlPlaneAdminOperationFullPath is the full path to deallocate a CCP
	DeallocateControlPlaneAdminOperationFullPath = AdminURLPrefix + DeallocateControlPlaneOperationRoute
	// DrainCustomerControlPlanesAdminOperationFullPath is the full path to drain a CCP between underlays
	DrainCustomerControlPlanesAdminOperationFullPath = AdminURLPrefix + DrainCustomerControlPlanesOperationRoute

	// ServiceOutboundIPRangesRoute is the route to GET/PUT/DELETE the service outbound IP ranges by service name
	ServiceOutboundIPRangesRoute = "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + ServiceOutboundIPRangesLiteral + "/{" + PathServiceNameParameter + "}"
	// ServiceOutboundIPRangesFullPath is the full path to admin operation to GET/PUT/DELETE the service outbound IP ranges by service name
	ServiceOutboundIPRangesFullPath = AdminURLPrefix + ServiceOutboundIPRangesRoute
	// ListServiceOutboundIPRangesRoute is the route to list all the service outbound IP ranges records
	ListServiceOutboundIPRangesRoute = "/" + LocationsLiteral + "/{" + PathLocationParameter + "}/" + ServiceOutboundIPRangesLiteral
	// ListServiceOutboundIPRangesFullPath is the full path to admin operation to list all the service outbound IP ranges records
	ListServiceOutboundIPRangesFullPath = AdminURLPrefix + ListServiceOutboundIPRangesRoute
)

const (
	// GetSubscriptionHandlerName is the constant logged for get subscription calls
	GetSubscriptionHandlerName = "GetSubscriptionHandler"
	// PutSubscriptionHandlerName is the constant logged for put subscription calls
	PutSubscriptionHandlerName = "PutSubscriptionHandler"

	// GetContainerServiceHandlerName is the constant logged for get container service calls
	GetContainerServiceHandlerName = "GetLegacyContainerServiceHandler"
	// PutContainerServiceHandlerName is the constant logged for put container service calls
	PutContainerServiceHandlerName = "PutLegacyContainerServiceHandler"
	// DeleteContainerServiceHandlerName is the constant logged for delete container service calls
	DeleteContainerServiceHandlerName = "DeleteLegacyContainerServiceHandler"
	// ListContainerServicesBySubscriptionHandlerName is the constant logged for list container services by subscription calls
	ListContainerServicesBySubscriptionHandlerName = "ListLegacyContainerServicesBySubscriptionHandler"
	// ListContainerServicesByResourceGroupHandlerName is the constant logged for list container services by subscription calls
	ListContainerServicesByResourceGroupHandlerName = "ListLegacyContainerServicesByResourceGroupHandler"

	// ListOrchestratorsHandlerName is the constant logged for list supported orchestrators
	ListOrchestratorsHandlerName = "ListOrchestratorsHandler"
	// GetOSOptionsHandlerName is the constant logged for get os options
	GetOSOptionsHandlerName = "GetOSOptionsHandler"
	// GetContainerServiceUpgradeProfileHandlerName is the constant logged for list available upgrade versions
	GetContainerServiceUpgradeProfileHandlerName = "GetContainerServiceUpgradeProfileHandler"
	// GetManagedClusterUpgradeProfileHandlerName is the constant logged for list available upgrade versions
	GetManagedClusterUpgradeProfileHandlerName = "GetManagedClusterUpgradeProfileHandler"
	//GetManagedClusterDiagnosticsStateHandlerName is the constant logged for get diagnostics state calls
	GetManagedClusterDiagnosticsStateHandlerName = "GetManagedClusterDiagnosticsStateHandler"
	// GetOperationResultsHandlerName is the constant logged for get operation results calls
	GetOperationResultsHandlerName = "GetOperationResultsHandler"

	// GetOperationStatusHandlerName is the constant logged for get operation calls
	GetOperationStatusHandlerName = "GetOperationStatusHandler"

	// PostDeploymentPreflightHandlerName is the constant logged for post deployment preflight calls
	PostDeploymentPreflightHandlerName = "PostDeploymentPreflightHandler"

	// MockHandlerName is the constant used for to log mock calls
	MockHandlerName = "MockHandler"

	// InternalPutContainerServiceHandlerName is the constant logged for internal put container service calls
	InternalPutContainerServiceHandlerName = "InternalPutContainerServiceHandler"

	// InternalPutSubscriptionHandlerName is the constant logged for internal put subscription calls
	InternalPutSubscriptionHandlerName = "InternalPutSubscriptionHandler"

	// InternalPutOperationStatusHandlerName is the constant logged for internal put operation calls
	InternalPutOperationStatusHandlerName = "InternalPutOperationStatusHandler"

	//HealthCheckHandlerName is a constant logged for health check calls
	HealthCheckHandlerName = "HealthCheckHandler"

	// GetAvailableOperationsHandlerName is the constant logged for get available operations calls
	GetAvailableOperationsHandlerName = "GetAvailableOperationsHandler"

	// LinkedNotificationHandlerName is the constant logged for Linked Notification calls
	LinkedNotificationHandlerName = "LinkedNotificationHandler"

	// AdminListUnderlaysHandlerName is the constant logged for list underlays calls
	AdminListUnderlaysHandlerName = "AdminListUnderlaysHandler"

	// AdminGetUnderlayHandlerName is the constant logged for get underlay calls
	AdminGetUnderlayHandlerName = "AdminGetUnderlayHandler"

	// AdminExpandUnderlayCapacityHandlerName is the constant logged for expand regional underlay capacity calls
	AdminExpandUnderlayCapacityHandlerName = "AdminExpandUnderlayCapacityHandler"

	// AdminPutUnderlayHandlerName is the constant logged for put underlay calls
	AdminPutUnderlayHandlerName = "AdminPutUnderlayHandler"

	// AdminDeleteUnderlayHandlerName is the constant logged for delete underlay calls
	AdminDeleteUnderlayHandlerName = "AdminDeleteUnderlayHandler"

	// AdminPostUnderlayActionHandlerName is the constant logged for post underlay action calls
	AdminPostUnderlayActionHandlerName = "AdminPostUnderlayActionHandler"

	// GetServiceOutboundIPRangesHandlerName is the constant logged for get service outbound IP ranges calls
	GetServiceOutboundIPRangesHandlerName = "GetServiceOutboundIPRangesHandler"
	// PutServiceOutboundIPRangesHandlerName is the constant logged for put service outbound IP ranges calls
	PutServiceOutboundIPRangesHandlerName = "PutServiceOutboundIPRangesHandler"
	// DeleteServiceOutboundIPRangesHandlerName is the constant logged for delete service outbound IP ranges calls
	DeleteServiceOutboundIPRangesHandlerName = "DeleteServiceOutboundIPRangesHandler"
	// ListServiceOutboundIPRangesHandlerName is the constant logged for list service outbound IP ranges calls
	ListServiceOutboundIPRangesHandlerName = "ListServiceOutboundIPRangesHandler"
)

const (
	// GetManagedClusterHandlerName is the constant logged for get hosted control plane calls
	GetManagedClusterHandlerName = "GetManagedClusterHandler"
	// GetAgentPoolHandlerName is the const logged for get agent pool calls
	GetAgentPoolHandlerName = "GetAgentPoolHandler"
	// GetAgentPoolUpgradeProfileHandlerName is the const logged for get agent pool upgrade profile
	GetAgentPoolUpgradeProfileHandlerName = "GetAgentPoolUpgradeProfileHandler"
	// ListAgentPoolAvailableVersionsHandlerName is the constant logged for listing available agent pool versions
	ListAgentPoolAvailableVersionsHandlerName = "ListAgentPoolAvailableVersionsHandler"
	// GetManagedClusterAccessProfileHandlerName is the constant logged for get accessProfile for roleName hosted control plane calls
	GetManagedClusterAccessProfileHandlerName = "GetManagedClusterAccessProfileHandler"
	// PutManagedClusterHandlerName is the constant logged for put hosted control plane calls
	PutManagedClusterHandlerName = "PutManagedClusterHandler"
	//PutAgentPoolHandlerName is the constant logged for put agentpool calls
	PutAgentPoolHandlerName = "PutAgentPoolHandler"
	// JsonPatchAgentPoolhandlerName is the const logged for Patching the Agent Pool Properties
	JsonPatchAgentPoolHandlerName = "JsonPatchAgentPoolhandler"
	// JsonPatchManagedClusterHandlerName is the const logged for Patching the Managed Cluster Properties
	JsonPatchManagedClusterHandlerName = "JsonPatchManaggedClusterhandlerName"
	// JsonPatchManagedClusterHandlerName is the const logged for Patching the Managed Cluster Properties
	JsonPatchControlPlaneHandlerName = "JsonPatchControlPlaneHandlerName"
	//UpgradeNodeImageAgentPoolHandlerName is the constant logged for upgrade node image version calls
	UpgradeNodeImageAgentPoolHandlerName = "UpgradeNodeImageAgentPoolHandler"
	// PatchManagedClusterHandlerName is the constant logged for patch hosted control plane calls
	PatchManagedClusterHandlerName = "PatchManagedClusterHandler"
	// DeleteManagedClusterHandlerName is the constant logged for delete hosted control plane calls
	DeleteManagedClusterHandlerName = "DeleteManagedClusterHandler"
	// StopManagedClusterHandlerName is the constant logged for stop managed cluster calls
	StopManagedClusterHandlerName = "StopManagedClusterHandler"
	// StartManagedClusterHandlerName is the constant logged for start managed cluster calls
	StartManagedClusterHandlerName = "StartManagedClusterHandler"
	// DeleteAgentPoolHandlerName is the constant logged for agent pool level operations.
	DeleteAgentPoolHandlerName = "DeleteAgentPoolHandler"
	// ListManagedClustersBySubscriptionHandlerName is the constant logged for list container services by subscription calls
	ListManagedClustersBySubscriptionHandlerName = "ListManagedClustersBySubscriptionHandler"
	// ListManagedClustersByResourceGroupHandlerName is the constant logged for list container services by subscription calls
	ListManagedClustersByResourceGroupHandlerName = "ListManagedClustersByResourceGroupHandler"
	// ListCustomerControlPlanePodsHandlerName is the constant logged for list customer control plane pods calls
	ListCustomerControlPlanePodsHandlerName = "ListCustomerControlPlanePodsHandler"
	// ListAgentPoolsByClusterHandlerName - List the agentpools in a cluster
	ListAgentPoolsByClusterHandlerName = "ListAgentPoolsByClusterHandler"
	// GetCustomerControlPlanePodLogHandlerName is the constant logged for get customer control plane pod log calls
	GetCustomerControlPlanePodLogHandlerName = "GetCustomerControlPlanePodLogHandler"
	//GetCustomerControlPlaneEventsHandlerName is the constant logged for get customer control plane event calls
	GetCustomerControlPlaneEventsHandlerName = "GetCustomerControlPlaneEventsHandler"
	// PostKubectlCommandHandlerName is the const logged for post kubectl command calls
	PostKubectlCommandHandlerName = "PostKubectlCommandHandler"
	// PostOverlayKubectlCommandHandlerName is the const logged for post kubectl command calls
	PostOverlayKubectlCommandHandlerName = "PostOverlayKubectlCommandHandler"
	// BackfillManagedClusterHandlerName is the const logged for back fill managed cluster call
	BackfillManagedClusterHandlerName = "BackfillManagedClusterHandlerName"
	// ReimageManagedClusterHandlerName is the const logged for back fill managed cluster call
	ReimageManagedClusterHandlerName = "ReimageManagedClusterHandlerName"
	// DelegateSubnetHandlerName is the const logged for delegate subnet call
	DelegateSubnetHandlerName = "DelegateSubnetHandlerName"
	// UnDelegateSubnetHandlerName is the const logged for undelegate subnet call
	UnDelegateSubnetHandlerName = "UnDelegateSubnetHandlerName"
	// ListUnderlaysHandlerName is the const logged for list underlays calls
	ListUnderlaysHandlerName = "ListUnderlaysHandler"
	// PostUnderlayKubectlCommandHandlerName is the const logged for post kubectl command calls given an underlay
	PostUnderlayKubectlCommandHandlerName = "PostUnderlayKubectlCommandHandler"
	// ListManagedClusterCredentialHandlerName is the const logged for list credentials of a managed cluster
	ListManagedClusterCredentialHandlerName = "ListManagedClusterCredentialHandlerName"
	// PostVirtualMachineRunCommandHandlerName is the const logged for post runCommand of a vm
	PostVirtualMachineRunCommandHandlerName = "PostVirtualMachineRunCommandHandler"
	// PostVirtualMachineGenericRunCommandHandlerName is the const logged for post generic runCommand of a vm
	PostVirtualMachineGenericRunCommandHandlerName = "PostVirtualMachineGenericRunCommandHandler"
	// GetCustomerAgentNodesStatusHandlerName is the const logged for get customer agent nodes status
	GetCustomerAgentNodesStatusHandlerName = "GetCustomerAgentNodesStatusHandler"
	// ListManagedClusterClusterAdminCredentialHandlerName is the const logged for list clusterAdmin credential of a managed cluster
	ListManagedClusterClusterAdminCredentialHandlerName = "ListManagedClusterClusterAdminCredentialHandlerName"
	// ListManagedClusterClusterUserCredentialHandlerName is the const logged for list clusterUser credential of a managed cluster
	ListManagedClusterClusterUserCredentialHandlerName = "ListManagedClusterClusterUserCredentialHandlerName"
	// ListManagedClusterClusterMonitoringUserCredentialHandlerName is the const logged for list clusterMonitoringUser credential of a managed cluster
	ListManagedClusterClusterMonitoringUserCredentialHandlerName = "ListManagedClusterClusterMonitoringUserCredentialHandlerName"
	// ResetServicePrincipalProfileHandlerName is the name of ResetServicePrincipalProfileHandler
	ResetServicePrincipalProfileHandlerName = "ResetServicePrincipalProfileHandler"
	// ResetAADProfileHandlerName is the name of ResetAADProfileHandler
	ResetAADProfileHandlerName = "ResetAADProfileHandler"
	// ListDetectorsHandlerName is the name of ListDetectors
	ListDetectorsHandlerName = "ListDetectorsHandler"
	// GetDetectorHandlerName is the name of GetDetector
	GetDetectorHandlerName = "GetDetectorHandler"
	// GetMaintenanceConfigurationHandlerName is the name of GetMaintenanceConfiguration
	GetMaintenanceConfigurationHandlerName = "GetMaintenanceConfigurationHandler"
	// PutMaintenanceConfigurationHandlerName is the name of PutMaintenanceConfiguration
	PutMaintenanceConfigurationHandlerName = "PutMaintenanceConfigurationHandler"
	// DeleteMaintenanceConfigurationHandlerName is the name of DeleteMaintenanceConfiguration
	DeleteMaintenanceConfigurationHandlerName = "DeleteMaintenanceConfigurationHandler"
	// ListMaintenanceConfigurationsByManagedClusterHandlerName is the name of ListMaintenanceConfigurationsByManagedCluster
	ListMaintenanceConfigurationsByManagedClusterHandlerName = "ListMaintenanceConfigurationsByManagedCluster"
	// MigrateCustomerControlPlaneHandlerName is the name of the MigrateCustomerControlPlane handler
	MigrateCustomerControlPlaneHandlerName = "MigrateCustomerControlPlaneHandler"
	// DeallocateControlPlaneHandlerName is the name of the DeallocateControlPlane handler
	DeallocateControlPlaneHandlerName = "DeallocateControlPlaneHandler"
	// DrainCustomerControlPlanesHandlerName is the name of the DrainCustomerControlPlanes handler
	DrainCustomerControlPlanesHandlerName = "DrainCustomerControlPlanesHandler"
	// RotateClusterCertificatesHandlerName is the name of RotateClusterCertificatesHandler
	RotateClusterCertificatesHandlerName = "RotateClusterCertificatesHandler"
	// RunCommandHandlerName is the name of the RunCommand handler
	RunCommandHandlerName = "RunCommandHandler"
	// RunCommandResultHandlerName is the name of the RunCommandResult handler
	RunCommandResultHandlerName = "RunCommandResultHandler"
	// ListPrivateEndpointConnectionsHandlerName is the name of ListPrivateEndpointConnectionsHandler
	ListPrivateEndpointConnectionsHandlerName = "ListPrivateEndpointConnectionsHandler"
	// PutPrivateEndpointConnectionHandlerName  is the name of PutPrivateEndpointConnectionHandler
	PutPrivateEndpointConnectionHandlerName = "PutPrivateEndpointConnectionHandler"
	// GetPrivateEndpointConnectionHandlerName  is the name of GetPrivateEndpointConnectionHandler
	GetPrivateEndpointConnectionHandlerName = "GetPrivateEndpointConnectionHandler"
	// DeletePrivateEndpointConnectionHandlerName  is the name of DeletePrivateEndpointConnectionHandler
	DeletePrivateEndpointConnectionHandlerName = "DeletePrivateEndpointConnectionHandler"
	// ListPrivateLinkResourcesHandlerName is the name of ListPrivateLinkResourcesHandler
	ListPrivateLinkResourcesHandlerName = "ListPrivateLinkResourcesHandler"
	// ResolvePrivateLinkServiceIDHandlerName is the name of ResolvePrivateLinkServiceIDHandler
	ResolvePrivateLinkServiceIDHandlerName = "ResolvePrivateLinkServiceIDHandler"

	// GetExtensionAddonHandlerName is the name of GetExtensionAddonHandler
	GetExtensionAddonHandlerName = "GetExtensionAddonHandler"
	// PutExtensionAddonHandlerName is the name of PutExtensionAddonHandler
	PutExtensionAddonHandlerName = "PutExtensionAddonHandler"
	// DeleteExtensionAddonHandlerName is the name of DeleteExtensionAddonHandler
	DeleteExtensionAddonHandlerName = "DeleteExtensionAddonHandler"
	// ListExtensionAddonsByManagedClusterHandlerName is the name of ListExtensionAddonsByManagedClusterHandler
	ListExtensionAddonsByManagedClusterHandlerName = "ListExtensionAddonsByManagedClusterHandler"
	// ListOutboundNetworkDependenciesEndpointsHandlerName is the name of ListOutboundNetworkDependenciesEndpointsHandler
	ListOutboundNetworkDependenciesEndpointsHandlerName = "ListOutboundNetworkDependenciesEndpointsHandler"

	// GetSnapshotHandlerName is the constant logged for get snapshot calls
	GetSnapshotHandlerName = "GetSnapshotHandler"
	// PutSnapshotHandlerName is the constant logged for put snapshot calls
	PutSnapshotHandlerName = "PutSnapshotHandler"
	// DeleteSnapshotHandlerName is the constant logged for delete snapshot calls
	DeleteSnapshotHandlerName = "DeleteSnapshotHandler"
	// ListSnapshotsBySubscriptionHandlerName is the constant logged for list Snapshots by subscription calls
	ListSnapshotsBySubscriptionHandlerName = "ListSnapshotsBySubscriptionHandler"
	// ListSnapshotsByResourceGroupHandlerName is the constant logged for list Snapshots by resource group calls
	ListSnapshotsByResourceGroupHandlerName = "ListSnapshotsByResourceGroupHandler"
	// GetSnapshotAdminHandlerName is the constant logged for get snapshot admin calls
	GetSnapshotAdminHandlerName = "GetSnapshotAdminHandler"

	// MigrateClusterV2HandlerName is the constant logged for migrating private cluster
	MigrateClusterV2HandlerName = "MigrateClusterV2Handler"
)

// Path operation names
const (
	GetSubscriptionOperationName = GetSubscriptionHandlerName + ".GET"
	PutSubscriptionOperationName = PutSubscriptionHandlerName + ".PUT"

	GetContainerServiceOperationName                  = GetContainerServiceHandlerName + ".GET"
	PutContainerServiceOperationName                  = PutContainerServiceHandlerName + ".PUT"
	DeleteContainerServiceOperationName               = DeleteContainerServiceHandlerName + ".DELETE"
	ListContainerServicesBySubscriptionOperationName  = ListContainerServicesBySubscriptionHandlerName + ".GET"
	ListContainerServicesByResourceGroupOperationName = ListContainerServicesByResourceGroupHandlerName + ".GET"

	GetManagedClusterOperationName                  = GetManagedClusterHandlerName + ".GET"
	GetManagedClusterAccessProfileOperationName     = GetManagedClusterAccessProfileHandlerName + ".GET"
	PutManagedClusterOperationName                  = PutManagedClusterHandlerName + ".PUT"
	PatchManagedClusterOperationName                = PatchManagedClusterHandlerName + ".PATCH"
	DeleteManagedClusterOperationName               = DeleteManagedClusterHandlerName + ".DELETE"
	ListManagedClustersBySubscriptionOperationName  = ListManagedClustersBySubscriptionHandlerName + ".GET"
	ListManagedClustersByResourceGroupOperationName = ListManagedClustersByResourceGroupHandlerName + ".GET"

	StopManagedClusterOperationName = StopManagedClusterHandlerName + ".POST"
	StartMangedClusterOperationName = StartManagedClusterHandlerName + ".POST"

	ManagedClusterRunCommandOperationName       = RunCommandHandlerName + ".POST"
	ManagedClusterRunCommandResultOperationName = RunCommandResultHandlerName + ".GET"

	GetAgentPoolOperationName                   = GetAgentPoolHandlerName + ".GET"
	PutAgentPoolOperationName                   = PutAgentPoolHandlerName + ".PUT"
	DeleteAgentPoolOperationName                = DeleteAgentPoolHandlerName + ".DELETE"
	ListAgentPoolsByClusterOperationName        = ListAgentPoolsByClusterHandlerName + ".GET"
	GetAgentPoolUpgradeProfileOperationName     = GetAgentPoolUpgradeProfileHandlerName + ".GET"
	ListAgentPoolAvailableVersionsOperationName = ListAgentPoolAvailableVersionsHandlerName + ".GET"
	UpgradeNodeImageAgentPoolOperationName      = UpgradeNodeImageAgentPoolHandlerName + ".POST"

	GetMaintenanceConfigurationOperationName                   = GetMaintenanceConfigurationHandlerName + ".GET"
	PutMaintenanceConfigurationOperationName                   = PutMaintenanceConfigurationHandlerName + ".PUT"
	DeleteMaintenanceConfigurationOperationName                = DeleteMaintenanceConfigurationHandlerName + ".DELETE"
	ListMaintenanceConfigurationsByManagedClusterOperationName = ListMaintenanceConfigurationsByManagedClusterHandlerName + ".GET"

	GetOperationResultsOperationName     = GetOperationResultsHandlerName + ".GET"
	GetOperationStatusOperationName      = GetOperationStatusHandlerName + ".GET"
	PostDeploymentPreflightOperationName = PostDeploymentPreflightHandlerName + ".POST"

	InternalPutContainerServiceOperationName = InternalPutContainerServiceHandlerName + ".PUT"
	InternalPutSubscriptionOperationName     = InternalPutSubscriptionHandlerName + ".PUT"
	InternalPutOperationStatusOperationName  = InternalPutOperationStatusHandlerName + ".PUT"

	HealthCheckOperationName = HealthCheckHandlerName + ".GET"

	ListOrchestratorsOperationName                                 = ListOrchestratorsHandlerName + ".GET"
	GetOSOptionsOperationName                                      = GetOSOptionsHandlerName + ".GET"
	GetContainerServiceUpgradeProfileOperationName                 = GetContainerServiceUpgradeProfileHandlerName + ".GET"
	GetManagedClusterUpgradeProfileOperationName                   = GetManagedClusterUpgradeProfileHandlerName + ".GET"
	GetManagedClusterDiagnosticsStateOperationName                 = GetManagedClusterDiagnosticsStateHandlerName + ".GET"
	ListManagedClusterCredentialOperationName                      = ListManagedClusterCredentialHandlerName + ".POST"
	GetAvailableOperationsOperationName                            = GetAvailableOperationsHandlerName + ".GET"
	ListManagedClusterClusterAdminCredentialOperationName          = ListManagedClusterClusterAdminCredentialHandlerName + ".POST"
	ListManagedClusterClusterUserCredentialOperationName           = ListManagedClusterClusterUserCredentialHandlerName + ".POST"
	ListManagedClusterClusterMonitoringUserCredentialOperationName = ListManagedClusterClusterMonitoringUserCredentialHandlerName + ".POST"

	ResetAADProfileOperationName              = ResetAADProfileHandlerName + ".POST"
	ResetServicePrincipalProfileOperationName = ResetServicePrincipalProfileHandlerName + ".POST"
	LinkedNotificationOperationName           = LinkedNotificationHandlerName + ".POST"

	ListDetectorsOperationName = ListDetectorsHandlerName + ".GET"
	GetDetectorOperationName   = GetDetectorHandlerName + ".GET"

	RotateClusterCertificatesOperationName = RotateClusterCertificatesHandlerName + ".POST"

	ListPrivateEndpointConnectionsOperationName  = ListPrivateEndpointConnectionsHandlerName + ".GET"
	PutPrivateEndpointConnectionOperationName    = PutPrivateEndpointConnectionHandlerName + ".PUT"
	GetPrivateEndpointConnectionOperationName    = GetPrivateEndpointConnectionHandlerName + ".GET"
	DeletePrivateEndpointConnectionOperationName = DeletePrivateEndpointConnectionHandlerName + ".DELETE"

	ListPrivateLinkResourcesOperationName    = ListPrivateLinkResourcesHandlerName + ".GET"
	ResolvePrivateLinkServiceIDOperationName = ResolvePrivateLinkServiceIDHandlerName + ".POST"

	MigrateClusterV2OperationName = MigrateClusterV2HandlerName + ".POST"

	// Admin operation names

	AdminOperationPrefix = "Admin."

	GetAgentPoolAdminOperationName                        = AdminOperationPrefix + GetAgentPoolOperationName
	ListAgentPoolsByClusterAdminOperationName             = AdminOperationPrefix + ListAgentPoolsByClusterOperationName
	PutAgentPoolAdminOperationName                        = AdminOperationPrefix + PutAgentPoolOperationName
	GetManagedClusterAdminOperationName                   = AdminOperationPrefix + GetManagedClusterOperationName
	PutManagedClusterAdminOperationName                   = AdminOperationPrefix + PutManagedClusterOperationName
	GetOperationStatusAdminOperationName                  = AdminOperationPrefix + GetOperationStatusOperationName
	ListManagedClustersBySubscriptionAdminOperationName   = AdminOperationPrefix + ListManagedClustersBySubscriptionOperationName
	ListManagedClustersByResourceGroupAdminOperationName  = AdminOperationPrefix + ListManagedClustersByResourceGroupOperationName
	ListCustomerControlPlanePodsOperationName             = AdminOperationPrefix + ListCustomerControlPlanePodsHandlerName + ".GET"
	GetCustomerControlPlanePodLogOperationName            = AdminOperationPrefix + GetCustomerControlPlanePodLogHandlerName + ".GET"
	GetCustomerControlPlaneEventsOperationName            = AdminOperationPrefix + GetCustomerControlPlaneEventsHandlerName + ".GET"
	PostKubectlCommandOperationName                       = AdminOperationPrefix + PostKubectlCommandHandlerName + ".POST"
	PostOverlayKubectlCommandOperationName                = AdminOperationPrefix + PostOverlayKubectlCommandHandlerName + ".POST"
	BackfillManagedClusterOperationName                   = AdminOperationPrefix + BackfillManagedClusterHandlerName + ".POST"
	ReimageManagedClusterAdminOperationName               = AdminOperationPrefix + ReimageManagedClusterHandlerName + ".POST"
	DelegateSubnetAdminOperationName                      = AdminOperationPrefix + DelegateSubnetHandlerName + ".POST"
	UnDelegateSubnetAdminOperationName                    = AdminOperationPrefix + UnDelegateSubnetHandlerName + ".POST"
	ListUnderlaysOperationName                            = AdminOperationPrefix + ListUnderlaysHandlerName + ".GET"
	PostUnderlayKubectlCommandOperationName               = AdminOperationPrefix + PostUnderlayKubectlCommandHandlerName + ".POST"
	GetSubscriptionResourceOperationName                  = AdminOperationPrefix + GetSubscriptionHandlerName + ".GET"
	PostVirtualMachineRunCommandOperationName             = AdminOperationPrefix + PostVirtualMachineRunCommandHandlerName + ".POST"
	PostVirtualMachineGenericRunCommandOperationName      = AdminOperationPrefix + PostVirtualMachineGenericRunCommandHandlerName + ".POST"
	GetCustomerAgentNodesStatusOperationName              = AdminOperationPrefix + GetCustomerAgentNodesStatusHandlerName + ".GET"
	LinkedNotificationAdminOperationName                  = AdminOperationPrefix + LinkedNotificationHandlerName + ".POST"
	AdminListUnderlaysOperationName                       = AdminOperationPrefix + AdminListUnderlaysHandlerName + ".GET"
	AdminGetUnderlayOperationName                         = AdminOperationPrefix + AdminGetUnderlayHandlerName + ".GET"
	AdminExpandUnderlayCapacityOperationName              = AdminOperationPrefix + AdminExpandUnderlayCapacityHandlerName + ".POST"
	AdminPutUnderlayOperationName                         = AdminOperationPrefix + AdminPutUnderlayHandlerName + ".PUT"
	AdminDeleteUnderlayOperationName                      = AdminOperationPrefix + AdminDeleteUnderlayHandlerName + ".DELETE"
	AdminPostUnderlayActionOperationName                  = AdminOperationPrefix + AdminPostUnderlayActionHandlerName + ".POST"
	MigrateCustomerControlPlaneOperationName              = AdminOperationPrefix + MigrateCustomerControlPlaneHandlerName + ".POST"
	DrainCustomerControlPlanesOperationName               = AdminOperationPrefix + DrainCustomerControlPlanesHandlerName + ".POST"
	DeallocateControlPlaneOperationName                   = AdminOperationPrefix + DeallocateControlPlaneHandlerName + ".POST"
	GetServiceOutboundIPRangesOperationName               = AdminOperationPrefix + GetServiceOutboundIPRangesHandlerName + ".GET"
	PutServiceOutboundIPRangesOperationName               = AdminOperationPrefix + PutServiceOutboundIPRangesHandlerName + ".PUT"
	DeleteServiceOutboundIPRangesOperationName            = AdminOperationPrefix + DeleteServiceOutboundIPRangesHandlerName + ".DELETE"
	ListServiceOutboundIPRangesOperationName              = AdminOperationPrefix + ListServiceOutboundIPRangesHandlerName + ".GET"
	GetExtensionAddonOperationName                        = GetExtensionAddonHandlerName + ".GET"
	PutExtensionAddonOperationName                        = PutExtensionAddonHandlerName + ".PUT"
	DeleteExtensionAddonOperationName                     = DeleteExtensionAddonHandlerName + ".DELETE"
	ListExtensionAddonsByManagedClusterOperationName      = ListExtensionAddonsByManagedClusterHandlerName + ".GET"
	ListOutboundNetworkDependenciesEndpointsOperationName = ListOutboundNetworkDependenciesEndpointsHandlerName + ".GET"
	JsonPatchAgentPoolAdminOperationName                  = AdminOperationPrefix + JsonPatchAgentPoolHandlerName + ".POST"
	JsonPatchManagedClusterAdminOperationName             = AdminOperationPrefix + JsonPatchManagedClusterHandlerName + ".POST"
	JsonPatchControlPlaneAdminOperationName               = AdminOperationPrefix + JsonPatchControlPlaneHandlerName + ".POST"
	GetSnapshotOperationName                              = GetSnapshotHandlerName + ".GET"
	PutSnapshotOperationName                              = PutSnapshotHandlerName + ".PUT"
	DeleteSnapshotOperationName                           = DeleteSnapshotHandlerName + ".DELETE"
	ListSnapshotsBySubscriptionOperationName              = ListSnapshotsBySubscriptionHandlerName + ".GET"
	ListSnapshotsByResourceGroupOperationName             = ListManagedClustersByResourceGroupHandlerName + ".GET"
	GetSnapshotAdminOperationName                         = AdminOperationPrefix + GetSnapshotOperationName
)

// Valid load balancer sku values
const (
	StandardLoadBalancerSku string = "standard"
	BasicLoadBalancerSku    string = "basic"
)

// Valid Outbound types for load balancer
const (
	// OutboundTypeLoadBalancer represents outbound connection type SLB
	OutboundTypeLoadBalancer string = "loadBalancer"

	// OutboundTypeUserDefinedRouting represents outbound connection type is defined from a UDR set by the customer
	// this can be in addition to a public IP
	OutboundTypeUserDefinedRouting string = "userDefinedRouting"

	// OutboundTypeManagedNATGateway represents outbound connection type is NAT gateway managed by AKS
	OutboundTypeManagedNATGateway string = "managedNATGateway"

	// OutboundTypeUserAssignedNATGateway represents outbound connection type is NAT gateway assigned to the subnets by the customer
	OutboundTypeUserAssignedNATGateway string = "userAssignedNATGateway"
)

// HCP Apiserver Claim Names
const (
	ClaimControlLoop    = "controlloop"
	ClaimACSRP          = "acsrp"
	ClaimRegionalLooper = "regionallooper"
	ClaimGenevaAction   = "genevaaction"
	ClaimOverlaymgr     = "overlaymgr"
	ClaimCPMonitor      = "cpmonitor"
	ClaimRemediator     = "remediator"
	ClaimReaderWriter   = "readerwriter"
	ClaimReader         = "reader"
	ClaimMCReconcile    = "mcreconcile"
	ClaimMCBackfill     = "mcbackfill"
	ClaimJITHandler     = "jithandler"
	ClaimJITController  = "jitcontroller"
	ClaimAutoUpgrader   = "autoupgrader"
)

// HCP API
const (
	// Non-terminal states
	StateAccepted   string = "Accepted"
	StateCreating   string = "Creating"
	StateUpdating   string = "Updating"
	StateDeleting   string = "Deleting"
	StateInProgress string = "In-Progress"

	// Operation priority
	PriorityHigh int = 0
	PriorityLow  int = 10
)

// Control Plane V1
const (
	ApiVersionV1 = "v1"
)

const (
	// SLBName the SLB is shared with ccp cloud provider that assumes the SLB name is the same as the cluster name.
	// Until we set it to a different one(through --cluster-name to controller manager), it is using the default value "kubernetes"
	SLBName     = "kubernetes"
	SLBNameIPv6 = "kubernetes-ipv6"

	// AKSManagedTagPrefix is the prefix of managed tags
	AKSManagedTagPrefix = "aks-managed-"

	// SLBManagedOutboundIPTypeTagName is the SLB managed outbound IP type tag name
	SLBManagedOutboundIPTypeTagName = "type"
	// SLBManagedOutboundIPTypeTagValue is the SLB managed outbound IP type tag value
	SLBManagedOutboundIPTypeTagValue = "aks-slb-managed-outbound-ip"

	// SlbOutboundBackendPoolName is the AKS SLB outbound backend pool name
	SlbOutboundBackendPoolName = "aksOutboundBackendPool"
	// SlbOutboundBackendPoolNameIPv6 is the AKS SLB outbound backend pool name for IPv6 traffic
	SlbOutboundBackendPoolNameIPv6 = "aksOutboundBackendPool-ipv6"
	// SlbOutboundRuleName is the AKS SLB outbound rule name
	SlbOutboundRuleName = "aksOutboundRule"

	// SlbDefaultOutboundRuleAllocatedOutboundPorts is the default outbound rule allocated outbound ports
	SlbDefaultOutboundRuleAllocatedOutboundPorts = 0
	// SlbDefaultOutboundRuleIdleTimeoutInMinutes is the default outbound rule idle timeout in minutes
	SlbDefaultOutboundRuleIdleTimeoutInMinutes = 30
	// SlbOutboundRuleProtocol is the outbound rule protocol
	SlbOutboundRuleProtocol = "All"
	// SlbManagedOutboundIPIdleTimeoutInMinutes is the managed outbound ip idle timeout in minutes
	SlbManagedOutboundIPIdleTimeoutInMinutes = 30

	// CloudProviderManagedServiceIPTagName is the cloud provider managed service IP tag name
	CloudProviderManagedServiceIPTagName = "service"

	// InternalLBNameSuffix is the suffix of the names of managed internal LBs
	InternalLBNameSuffix = "-internal"
)

var (
	// SLBManagedOutboundIPOwnerTagName is the SLB managed outbound IP owner tag key,
	// this is only useful when multi-slb is enabled.
	SLBManagedOutboundIPOwnerTagName = fmt.Sprintf("%s%s", AKSManagedTagPrefix, "slb-outbound-ip-owner")

	PrefixedSLBManagedOutboundIPTypeTagName = fmt.Sprintf("%s%s", AKSManagedTagPrefix, "type")
)

const (
	// NATGatewayManagedOutboundIPTypeTagName is the NAT gateway managed outbound IP type tag name
	NATGatewayManagedOutboundIPTypeTagName = "aks-managed-outbound-ip-type"

	// NATGatewayManagedOutboundIPTypeTagValue is the NAT gateway managed outbound IP type tag value
	NATGatewayManagedOutboundIPTypeTagValue = "nat-gateway"

	// NATGatewayManagedOutboundIPOwnerTagName is the managed outbound IP owner tag name
	NATGatewayManagedOutboundIPOwnerTagName = "aks-managed-outbound-ip-owner"

	// NatGatewayDefaultManagedOutboundIPCount is the default managed outbound IP count
	NatGatewayDefaultManagedOutboundIPCount = 1

	// NatGatewayDefaultIdleTimeoutInMinutes is the default value of idle timeout in minutes for NAT gateway
	NatGatewayDefaultIdleTimeoutInMinutes = 4

	// NatGatewayManagedOutboundIPIdleTimeoutInMinutes is the managed outbound IP idle timeout in minutes
	NatGatewayManagedOutboundIPIdleTimeoutInMinutes = 4
)

const (
	// CustomizedUbuntuAks1804 specifies AKS Ubuntu 1804 for customized Ubuntu
	CustomizedUbuntuAks1804 = "aks-ubuntu-1804"
)

const (
	// CustomizedWindowsAks2019PIR specifies AKS Windows 2019 PIR for customized Windows
	// SIG image is enabled by default for Windows so we use this to test Windows PIR images
	CustomizedWindowsAks2019PIR = "aks-windows-2019-pir"
)

const (
	// LongKubernetesNodeStatusUpdateFrequency (5m instead of 10s)
	LongKubernetesNodeStatusUpdateFrequency = "5m"
)

const (
	// DefaultHostedProfileMasterName specifies the 3 character orchestrator code of the clusters with hosted master profiles.
	DefaultHostedProfileMasterName = "aks"
)

const (
	// DefaultTempDiskDataDirPath is the default path for the container data directory when using temp disk.
	// It is used by the enable-container-data-dir-temp-disk toggle.
	DefaultTempDiskDataDirPath = "/mnt/containers"
)
