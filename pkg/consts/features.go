//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package consts

// Case insensitive literals
const (
	// TunnelVersionInHeader reprents the string in the http request header
	TunnelVersionInHeader = "TunnelVersion"
	// OpenVPNFeatureType represents the feature flag in subscription
	OpenVPNFeatureType = "Microsoft.ContainerService/OpenVPN"
	// KonnectivityFeatureType represents the feature flag in subscription
	KonnectivityFeatureType = "Microsoft.ContainerService/Konnectivity"
	// NoKonnectivityFeatureType represents the feature flag in the http request header
	NoKonnectivityFeatureType = "Microsoft.ContainerService/NoKonnectivity"
	// ControlPlaneUnderlayType represents the feature flag in subscription
	ControlPlaneUnderlayType = "Microsoft.ContainerService/ControlPlaneUnderlay"
	// ControlPlaneUnderlayHeader represents the feature flag in the http request header
	ControlPlaneUnderlayHeader = "ControlPlaneUnderlay"
	// VMSSPreview represents the feature flag for enabling VMSS agent nodes feature
	VMSSPreview = "Microsoft.ContainerService/VMSSPreview"
	// MultiAgentPoolPreview represents the flags for enabling the multi agentpool and operations at agentpool level feature.
	MultiAgentpoolPreview = "Microsoft.ContainerService/MultiAgentpoolPreview"
	// PodSecurityPolicyPreview represents the feature flag for enabling PodSecurityPolicy feature
	PodSecurityPolicyPreview = "Microsoft.ContainerService/PodSecurityPolicyPreview"
	// AKSHTTPCustomFeatures represents the feature flag for enabling reading custom features from HTTP header (mainly restricted for internal usage)
	AKSHTTPCustomFeatures = "Microsoft.ContainerService/AKSHTTPCustomFeatures"
	// AKSHTTPCustomFeaturesHeader represents the custom features in the http request header
	AKSHTTPCustomFeaturesHeader = "AKSHTTPCustomFeatures"
	// StaticTokenPreview represents the feature flag for enabling static token feature
	StaticTokenPreview = "Microsoft.ContainerService/StaticTokenPreview"
	// UseRawUbuntuVHD represents the feature flag in the http request header
	UseRawUbuntuVHD = "Microsoft.ContainerService/UseRawUbuntuVHD"
	// AKSNetworkModePreview represents the feature flag for enabling network mode feature
	AKSNetworkModePreview = "Microsoft.ContainerService/AKSNetworkModePreview"
	// UseCustomizedOSImage represents the feature flag in the http request header
	UseCustomizedOSImage = "Microsoft.ContainerService/UseCustomizedOSImage"
	// OSImageNameHeader represents OSImageName header name
	OSImageNameHeader = "OSImageName"
	// OSImageResourceGroupHeader represents OSImageResourceGroup header name
	OSImageResourceGroupHeader = "OSImageResourceGroup"
	// OSImageSubscriptionIDHeader represents OSImageSubscriptionID header name
	OSImageSubscriptionIDHeader = "OSImageSubscriptionID"
	// OSImageGalleryHeader represents OSImageGallery header name
	OSImageGalleryHeader = "OSImageGallery"
	// OSImageVersionHeader represents OSImageVersion header name
	OSImageVersionHeader = "OSImageVersion"
	// OSSKUHeader represents OSSKU header name
	OSSKUHeader = "OSSKU"
	// UseCustomizedWindowsOSImage represents the feature flag in the http request header
	UseCustomizedWindowsOSImage = "Microsoft.ContainerService/UseCustomizedWindowsOSImage"
	// WindowsOSImageNameHeader represents OSImageName header name
	WindowsOSImageNameHeader = "WindowsOSImageName"
	// WindowsOSImageResourceGroupHeader represents WindowsOSImageResourceGroup header name
	WindowsOSImageResourceGroupHeader = "WindowsOSImageResourceGroup"
	// WindowsOSImageSubscriptionIDHeader represents WindowsOSImageSubscriptionID header name
	WindowsOSImageSubscriptionIDHeader = "WindowsOSImageSubscriptionID"
	// WindowsOSImageGalleryHeader represents WindowsOSImageGallery header name
	WindowsOSImageGalleryHeader = "WindowsOSImageGallery"
	// WindowsOSImageVersionHeader represents WindowsOSImageVersion header name
	WindowsOSImageVersionHeader = "WindowsOSImageVersion"
	// CustomizedUbuntuHeader represents CustomizedUbuntu http header name
	CustomizedUbuntuHeader = "CustomizedUbuntu"
	// CustomizedWindowsHeader represents CustomizedWindows http header name
	CustomizedWindowsHeader = "CustomizedWindows"
	// UseCustomizedUbuntuPreview represents the feature flag for enabling UseCustomizedUbuntuPreview
	UseCustomizedUbuntuPreview = "Microsoft.ContainerService/UseCustomizedUbuntuPreview"
	// UseCustomizedWindowsPreview represents the feature flag for enabling UseCustomizedWindowsPreview
	UseCustomizedWindowsPreview = "Microsoft.ContainerService/UseCustomizedWindowsPreview"
	// CrossTenantVnet represents the feature flag for cross tenant vnet
	CrossTenantVnet = "Microsoft.ContainerService/CrossTenantVnet"
	// CrossTenantRouteTable represents the feature flag for cross tenant route table
	CrossTenantRouteTable = "Microsoft.ContainerService/CrossTenantRouteTable"
	// Gen2VMPreview represents the feature flag for enabling Gen2 VM
	Gen2VMPreview = "Microsoft.ContainerService/Gen2VMPreview"
	// UseGen2VMHeader represents UseGen2VMHeader http header name
	UseGen2VMHeader = "UseGen2VM"
	// UserAssignedIdentityPreview represents the feature flag for using user assigned identity in CCP for MSI cluster
	UserAssignedIdentityPreview = "Microsoft.ContainerService/UserAssignedIdentityPreview"
	// EnableAzureDiskFileCSIDriverHeader represents EnableAzureDiskFileCSIDriver header name
	EnableAzureDiskFileCSIDriverHeader = "EnableAzureDiskFileCSIDriver"
	// EnableAzureDiskFileCSIDriver represents the feature flag for enabling EnableAzureDiskFileCSIDriver
	EnableAzureDiskFileCSIDriver = "Microsoft.ContainerService/EnableAzureDiskFileCSIDriver"
	// EnableEncryptionAtHostHeader represents EnableEncryptionAtHost header name
	EnableEncryptionAtHostHeader = "EnableEncryptionAtHost"
	// EnableSwiftNetworkingHeader represents EnableSwiftNetworking header name
	EnableSwiftNetworkingHeader = "EnableSwiftNetworking"
	// EnableNetworkPluginNoneHeader represents EnableNetworkPluginNone header name
	EnableNetworkPluginNoneHeader = "EnableNetworkPluginNone"
	// EnableAzureDefender represents the feature flag for enabling AzureDefender
	EnableAzureDefender = "Microsoft.ContainerService/AKS-AzureDefender"
	// GPUDedicatedVHDPreview represents the feature flag for enabling GPU native VHD (with drivers and device plugin)
	GPUDedicatedVHDPreview = "Microsoft.ContainerService/GPUDedicatedVHDPreview"
	// UseGPUDedicatedVHDHeader represents UseGPUDedicatedVHD http header name
	UseGPUDedicatedVHDHeader = "UseGPUDedicatedVHD"
	// UseCustomizedContainerRuntime is a feature flag to specify a custom container runtime for agent nodes
	UseCustomizedContainerRuntime = "Microsoft.ContainerService/UseCustomizedContainerRuntime"
	// ContainerRuntime represents ContainerRuntime header name. Currently only containerd is supported.
	ContainerRuntime = "ContainerRuntime"
	// EnableUltraSSDHeader represents EnableUltraSSD header name
	EnableUltraSSDHeader = "EnableUltraSSD"
	// EnableEphemeralOSDiskHeader represents EnableEphemeralOSDisk header name
	EnableEphemeralOSDiskHeader = "EnableEphemeralOSDisk"
	// EnableEphemeralOSDiskPreview represents the feature flag for enabling EnableEphemeralOSDiskPreview
	EnableEphemeralOSDiskPreview = "Microsoft.ContainerService/EnableEphemeralOSDiskPreview"
	// EnableCloudControllerManagerHeader represents EnableCloudControllerManager header name
	EnableCloudControllerManagerHeader = "EnableCloudControllerManager"
	// EnableCloudControllerManager represents the feature flag for enabling EnableCloudControllerManager
	EnableCloudControllerManager = "Microsoft.ContainerService/EnableCloudControllerManager"
	// EnableMultipleStandardLoadBalancers represents the feature flag of EnableMultipleStandardLoadBalancers
	EnableMultipleStandardLoadBalancers = "Microsoft.ContainerService/EnableMultipleStandardLoadBalancers"
	// EnableMultipleStandardLoadBalancersHeader represents the EnableMultipleStandardLoadBalancer header name
	EnableMultipleStandardLoadBalancersHeader = "EnableMultipleStandardLoadBalancers"
	// EnablePodIdentityPreview represents the feature flag for enabling Pod Identity
	EnablePodIdentityPreview = "Microsoft.ContainerService/EnablePodIdentityPreview"
	// EnablePodIdentityV2Preview represents the feature flag for enabling Pod Identity v2
	EnablePodIdentityV2Preview = "Microsoft.ContainerService/EnablePodIdentityV2Preview"
	// MigrateToMSIClusterPreview represents the feature flag for migrating SPN cluster to MSI cluster
	MigrateToMSIClusterPreview = "Microsoft.ContainerService/MigrateToMSIClusterPreview"
	// PodSubnetPreview represents the feature flag for pod subnet preview
	PodSubnetPreview = "Microsoft.ContainerService/PodSubnetPreview"
	// EnableACRTeleportPreview represents the feature flag for enable acr teleport plugin
	EnableACRTeleportPreview = "Microsoft.ContainerService/EnableACRTeleport"
	// EnableACRTeleportHeader represents the EnableACRTeleport Header name
	EnableACRTeleportHeader = "EnableACRTeleport"
	// CustomNodeConfigPreview represents the feature flag for custom node config
	CustomNodeConfigPreview = "Microsoft.ContainerService/CustomNodeConfigPreview"
	// UseCustomNodeConfigHeader represents UseCustomNodeConfig http header name
	UseCustomNodeConfigHeader = "UseCustomNodeConfig"
	// EnablePrivateClusterFQDNSubdomain
	EnablePrivateClusterFQDNSubdomain = "Microsoft.ContainerService/EnablePrivateClusterFQDNSubdomain"
	// AKSConfidentialComputingAddon represents the feature flag to allow AKS cluster in a subscription to enable ACC confidential computing addon
	AKSConfidentialComputingAddon = "Microsoft.ContainerService/AKS-ConfidentialComputingAddon"
	// KubeletDisk represents the feature flag for using temp disk for KubeletDisk.
	KubeletDisk = "Microsoft.ContainerService/KubeletDisk"
	// WasmNodePoolPreview represents the feature flag to allow running WASM workloads with krustlet.
	WasmNodePoolPreview = "Microsoft.ContainerService/WasmNodePoolPreview"
	// CustomKubeletIdentityPreview represents the feature flag for using own kubelet identity
	CustomKubeletIdentityPreview = "Microsoft.ContainerService/CustomKubeletIdentityPreview"
	// EnableAKSWindowsCalico represents the feature flag to allow AKS cluster in a subscription to enable using Calico for all Windows agent pools
	EnableAKSWindowsCalico = "Microsoft.ContainerService/EnableAKSWindowsCalico"
	// EnableAgentNodeKubeletClientTLSBootstrapHeader represents the EnableAgentNodeKubeletClientTLSBootstrapHeader header name.
	EnableAgentNodeKubeletClientTLSBootstrapHeader = "EnableAgentNodeKubeletClientTLSBootstrap"
	// DisableLocalAccountsPreview represents the feature flag for disabling all local accounts
	DisableLocalAccountsPreview = "Microsoft.ContainerService/DisableLocalAccountsPreview"
	// AKSCBLMariner represents the feature flag for accessing the CBLMariner OSSKU
	AKSCBLMariner = "Microsoft.ContainerService/AKSCBLMariner"
	// WindowsContainerRuntime represents WindowsContainerRuntime header name. Currently only containerd is supported.
	WindowsContainerRuntime = "WindowsContainerRuntime"
	// HTTPProxyConfigPreview represents the feature flag for http proxy config
	HTTPProxyConfigPreview = "Microsoft.ContainerService/HTTPProxyConfigPreview"
	// EnablePrivateClusterPublicFQDN represents the feature to add additional public FQDN for private cluster
	EnablePrivateClusterPublicFQDN = "Microsoft.ContainerService/EnablePrivateClusterPublicFQDN"
	// EnablePrivateClusterV2 represents the feature flag to allow creating private cluster v2 in a subscription
	EnablePrivateClusterV2 = "Microsoft.ContainerService/EnablePrivateClusterV2"
	// EnablePrivateClusterSubZone represents the feature to use sub zone for byo private dns zone scenario
	EnablePrivateClusterSubZone = "Microsoft.ContainerService/EnablePrivateClusterSubZone"
	// ForceUpdateProvisioningState allows PUT to update the cluster's provisioning state to the desired value
	ForceUpdateProvisioningState = "ForceUpdateProvisioningState"
	// CreateMcReconcileRequest will create a mc reconcile request if true
	CreateMcReconcileRequest = "CreateMcReconcileRequest"
	// OmsagentAADAuthPreview is the feature flag for using AAD authentication for omsagent addon
	OmsagentAADAuthPreview = "Microsoft.ContainerService/OmsagentAADAuthPreview"
	// AKSNodelessPreview represents the feature flag for previewing empty input agentpool
	AKSNodelessPreview = "Microsoft.ContainerService/AKSNodelessPreview"
	// PreviewStartStopAgentPool represents the feature flag for agent pool start and stop
	PreviewStartStopAgentPool = "Microsoft.ContainerService/PreviewStartStopAgentPool"
	// InternalAzureEdgeZones is a feature flag to indicate that resources can be deployed to internal microsoft edge zones
	InternalAzureEdgeZones = "Microsoft.Resources/InternalAzureEdgeZones"
	// AzureEdgeZones is a feature flag to indicate that resources can be deployed to any microsoft edge zones.
	AzureEdgeZones = "Microsoft.Resources/azureedgezones"
	// EnableWindowsPIRToSIGUpgradeForVMSS enables upgrading Windows agentpool PIR images to SIG images
	// This is not a feature flag, but a header value introduced for testing
	EnableWindowsPIRToSIGUpgradeForVMSS = "EnableWindowsPIRToSIGUpgradeForVMSS"
	// ScaleDownModePreview determines the behaviour when scaling up and down Agent Pools
	ScaleDownModePreview = "Microsoft.ContainerService/AKS-ScaleDownModePreview"
	// SnapshotPreview represents the feature flag for previewing Nodepool Snapshot
	SnapshotPreview = "Microsoft.ContainerService/SnapshotPreview"
	// AKSSnapshotIdHeader represents AKSSnapshotIdHeader http header name
	AKSSnapshotIdHeader = "AKSSnapshotId"
	// NATGatewayPreview represents the feature flag to allow managedNATGateway and userAssignedNATGateway outbound types
	NATGatewayPreview = "Microsoft.ContainerService/AKS-NATGatewayPreview"
	// AKSWindowsGmsaPreview represents the feature flag to enable Windows GMSA
	AKSWindowsGmsaPreview = "Microsoft.ContainerService/AKSWindowsGmsaPreview"
	// EnableWinDSR represents EnableWinDSR header name.
	// This is not a feature flag, but a header value introduced for testing
	EnableWinDSR = "EnableWinDSR"
	// NTOE: EnableWindowsGmsa, WindowsGmsaDnsServer and WindowsGmsaRootDomainName are only used for test in custom header
	// EnableWindowsGmsa represents EnableWindowsGmsa header name.
	EnableWindowsGmsa = "EnableWindowsGmsa"
	// WindowsGmsaDnsServer represents WindowsGmsaDnsServer header name.
	WindowsGmsaDnsServer = "WindowsGmsaDnsServer"
	// WindowsGmsaRootDomainName represents WindowsGmsaRootDomainName header name.
	WindowsGmsaRootDomainName = "WindowsGmsaRootDomainName"
	// EnableOIDCIssuerPreview represents the feature flag for enabling OIDC issuer.
	EnableOIDCIssuerPreview = "Microsoft.ContainerService/EnableOIDCIssuerPreview"
	// AKSEnableDualStack represents the feature flag for dual-stack networking for AKS clusters
	AKSEnableDualStack = "Microsoft.ContainerService/AKS-EnableDualStack"
	// EnableNamespaceResourcesPreview represents the feature flag to enable Namespace Resources
	EnableNamespaceResourcesPreview = "Microsoft.ContainerService/EnableNamespaceResourcesPreview"

	// AddonManagerV2Preview represents the feature flag in the http request header
	AddonManagerV2Preview = "AddonManagerV2Preview"

	// UseE2EAgentPoolNameSeed is the feature flag used in E2E test to hard code the name seed value in
	// the agent pool file. This makes the generated VMSS name predictable. E2E needs this in order to
	// access and sometimes modify the VMSS directly.
	UseE2EAgentPoolNameSeed = "Microsoft.ContainerService/UseE2EAgentPoolNameSeed"
)

// AddonFeatureFlagMap is the mapping for addon to its respective feature flag
var AddonFeatureFlagMap = map[string]string{
	AddonNameGitOps:           "Microsoft.ContainerService/AKS-GitOps",
	AddonNameExtensionManager: "Microsoft.ContainerService/AKS-ExtensionManager",
}

// ExtensionFeatureFlagMap is the mapping for extension to its respective feature flag
var ExtensionFeatureFlagMap = map[string]string{
	ExtensionTypeDapr: "Microsoft.ContainerService/AKS-Dapr",
}
