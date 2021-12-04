//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package armerror

import (
	"fmt"
	"time"
)

// ARMOrResourceProviderErrorCode represents the message returned from ARM or a
// Resource Provider
type ARMOrResourceProviderErrorCode string

const (
	ResourceDeploymentFailure                                                        ARMOrResourceProviderErrorCode = "ResourceDeploymentFailure"
	DNSRecordInUse                                                                   ARMOrResourceProviderErrorCode = "DNSRecordInUse"
	MaxStorageAccountsCountPerSubscriptionExceeded                                   ARMOrResourceProviderErrorCode = "MaxStorageAccountsCountPerSubscriptionExceeded"
	OverconstrainedAllocationRequest                                                 ARMOrResourceProviderErrorCode = "OverconstrainedAllocationRequest"
	OverconstrainedZonalAllocationRequest                                            ARMOrResourceProviderErrorCode = "OverconstrainedZonalAllocationRequest"
	TooManyRequests                                                                  ARMOrResourceProviderErrorCode = "TooManyRequests"
	Conflict                                                                         ARMOrResourceProviderErrorCode = "Conflict"
	SkuNotAvailable                                                                  ARMOrResourceProviderErrorCode = "SkuNotAvailable"
	RequestDisallowedByPolicy                                                        ARMOrResourceProviderErrorCode = "RequestDisallowedByPolicy"
	StorageAccountNotRecognized                                                      ARMOrResourceProviderErrorCode = "StorageAccountNotRecognized"
	MissingSubscriptionRegistration                                                  ARMOrResourceProviderErrorCode = "MissingSubscriptionRegistration"
	SubscriptionNotRegistered                                                        ARMOrResourceProviderErrorCode = "SubscriptionNotRegistered"
	MissingRegistrationForType                                                       ARMOrResourceProviderErrorCode = "MissingRegistrationForType"
	ResourceGroupBeingDeleted                                                        ARMOrResourceProviderErrorCode = "ResourceGroupBeingDeleted"
	ResourceGroupNotFound                                                            ARMOrResourceProviderErrorCode = "ResourceGroupNotFound"
	ResourceGroupQuotaExceeded                                                       ARMOrResourceProviderErrorCode = "ResourceGroupQuotaExceeded"
	AuthorizationFailed                                                              ARMOrResourceProviderErrorCode = "AuthorizationFailed"
	PublicIPCountLimitReached                                                        ARMOrResourceProviderErrorCode = "PublicIPCountLimitReached"
	StandardSkuPublicIPCountLimitReached                                             ARMOrResourceProviderErrorCode = "StandardSkuPublicIPCountLimitReached"
	VnetCountLimitReached                                                            ARMOrResourceProviderErrorCode = "VnetCountLimitReached"
	VnetAddressSpaceOverlapsWithAlreadyPeeredVnet                                    ARMOrResourceProviderErrorCode = "VnetAddressSpaceOverlapsWithAlreadyPeeredVnet"
	VnetAddressSpacesOverlap                                                         ARMOrResourceProviderErrorCode = "VnetAddressSpacesOverlap"
	OperationNotAllowed                                                              ARMOrResourceProviderErrorCode = "OperationNotAllowed"
	PrivateIPAddressNotInSubnet                                                      ARMOrResourceProviderErrorCode = "PrivateIPAddressNotInSubnet"
	InvalidParameter                                                                 ARMOrResourceProviderErrorCode = "InvalidParameter"
	InvalidResourceReference                                                         ARMOrResourceProviderErrorCode = "InvalidResourceReference"
	ResizeDiskError                                                                  ARMOrResourceProviderErrorCode = "ResizeDiskError"
	ResourceGroupDeletionTimeout                                                     ARMOrResourceProviderErrorCode = "ResourceGroupDeletionTimeout"
	ScopeLocked                                                                      ARMOrResourceProviderErrorCode = "ScopeLocked"
	PublicIPAddressCannotBeDeleted                                                   ARMOrResourceProviderErrorCode = "PublicIPAddressCannotBeDeleted"
	MgmtDeletionDependency                                                           ARMOrResourceProviderErrorCode = "MgmtDeletionDependency"
	BadRequest                                                                       ARMOrResourceProviderErrorCode = "BadRequest"
	ApplicationGatewaySubnetCannotBeUsedByOtherResources                             ARMOrResourceProviderErrorCode = "ApplicationGatewaySubnetCannotBeUsedByOtherResources"
	GatewaySubnet                                                                    ARMOrResourceProviderErrorCode = "GatewaySubnet"
	InUseSubnetCannotBeDeleted                                                       ARMOrResourceProviderErrorCode = "InUseSubnetCannotBeDeleted"
	InUseNetworkSecurityGroupCannotBeDeleted                                         ARMOrResourceProviderErrorCode = "InUseNetworkSecurityGroupCannotBeDeleted"
	InUseRouteTableCannotBeDeleted                                                   ARMOrResourceProviderErrorCode = "InUseRouteTableCannotBeDeleted"
	AllocationFailed                                                                 ARMOrResourceProviderErrorCode = "AllocationFailed"
	QuotaExceeded                                                                    ARMOrResourceProviderErrorCode = "QuotaExceeded"
	AppendPoliciesFieldsExist                                                        ARMOrResourceProviderErrorCode = "AppendPoliciesFieldsExist"
	VMExtensionProvisioningError                                                     ARMOrResourceProviderErrorCode = "VMExtensionProvisioningError"
	SubnetIsFull                                                                     ARMOrResourceProviderErrorCode = "SubnetIsFull"
	OperationPreempted                                                               ARMOrResourceProviderErrorCode = "OperationPreempted"
	SubnetWithExternalResourcesCannotBeUsedByOtherResources                          ARMOrResourceProviderErrorCode = "SubnetWithExternalResourcesCannotBeUsedByOtherResources"
	ApplianceBeingDeleted                                                            ARMOrResourceProviderErrorCode = "ApplianceBeingDeleted"
	CustomerDeniedServicePlanAccess                                                  ARMOrResourceProviderErrorCode = "CustomerDeniedServicePlanAccess"
	ResourcePurchaseValidationFailed                                                 ARMOrResourceProviderErrorCode = "ResourcePurchaseValidationFailed"
	MarketplacePurchaseEligibilityFailed                                             ARMOrResourceProviderErrorCode = "MarketplacePurchaseEligibilityFailed"
	FailedRoleAssignmentCreation                                                     ARMOrResourceProviderErrorCode = "FailedRoleAssignmentCreation"
	PrincipalNotFound                                                                ARMOrResourceProviderErrorCode = "PrincipalNotFound"
	Canceled                                                                         ARMOrResourceProviderErrorCode = "Canceled"
	CanceledAndSupersededDueToAnotherOperation                                       ARMOrResourceProviderErrorCode = "CanceledAndSupersededDueToAnotherOperation"
	CertificateNotFound                                                              ARMOrResourceProviderErrorCode = "CertificateNotFound"
	ResourceNotPermittedOnDelegatedSubnet                                            ARMOrResourceProviderErrorCode = "ResourceNotPermittedOnDelegatedSubnet"
	ReferencedResourceNotProvisioned                                                 ARMOrResourceProviderErrorCode = "ReferencedResourceNotProvisioned"
	PrivateEndpointAndServiceEndpointsCannotCoexistInASubnet                         ARMOrResourceProviderErrorCode = "PrivateEndpointAndServiceEndpointsCannotCoexistInASubnet"
	ResourcesBeingMoved                                                              ARMOrResourceProviderErrorCode = "ResourcesBeingMoved"
	CannotUpdatePlan                                                                 ARMOrResourceProviderErrorCode = "CannotUpdatePlan"
	ConflictingAppendPolicies                                                        ARMOrResourceProviderErrorCode = "ConflictingAppendPolicies"
	InvalidPolicyParameters                                                          ARMOrResourceProviderErrorCode = "InvalidPolicyParameters"
	PublicIPCountLimitExceededByVMScaleSet                                           ARMOrResourceProviderErrorCode = "PublicIPCountLimitExceededByVMScaleSet"
	InvalidGatewayHost                                                               ARMOrResourceProviderErrorCode = "InvalidGatewayHost"
	ReservedResourceName                                                             ARMOrResourceProviderErrorCode = "ReservedResourceName"
	KeyVaultAccessForbidden                                                          ARMOrResourceProviderErrorCode = "KeyVaultAccessForbidden"
	DiskEncryptionSet                                                                ARMOrResourceProviderErrorCode = "DiskEncryptionSet"
	PolicyViolation                                                                  ARMOrResourceProviderErrorCode = "policy violation"
	SpecifiedAllocatedOutboundPortsForOutboundRuleExceedsTotalNumberOfAvailablePorts ARMOrResourceProviderErrorCode = "SpecifiedAllocatedOutboundPortsForOutboundRuleExceedsTotalNumberOfAvailablePorts"
	NetworkSecurityGroupInUseByVirtualMachineScaleSet                                ARMOrResourceProviderErrorCode = "NetworkSecurityGroupInUseByVirtualMachineScaleSet"
	LoadBalancerInUseByVirtualMachineScaleSet                                        ARMOrResourceProviderErrorCode = "LoadBalancerInUseByVirtualMachineScaleSet"
	ReadOnlyDisabledSubscription                                                     ARMOrResourceProviderErrorCode = "ReadOnlyDisabledSubscription"
	Throttled                                                                        ARMOrResourceProviderErrorCode = "Throttled"
	SubscriptionRequestsThrottled                                                    ARMOrResourceProviderErrorCode = "SubscriptionRequestsThrottled"
	RoleAssignmentLimitExceeded                                                      ARMOrResourceProviderErrorCode = "RoleAssignmentLimitExceeded"
	ObjectIsDeletedButRecoverable                                                    ARMOrResourceProviderErrorCode = "ObjectIsDeletedButRecoverable"
	ObjectIsBeingDeleted                                                             ARMOrResourceProviderErrorCode = "ObjectIsBeingDeleted"
	NetworkInterfaceAndLoadBalancerAreInDifferentAvailabilitySets                    ARMOrResourceProviderErrorCode = "NetworkInterfaceAndLoadBalancerAreInDifferentAvailabilitySets "
)

const (
	// Snippet of error message to explain the error cause.
	ResourceGroupDeletionBlocked = "ResourceGroupDeletionBlocked"

	// resource type within error, used to identify ADAL resource type
	// adal package has two packageType "azure.multiTenantSPTAuthorizer" and "azure.BearerAuthorizer"
	// and both error have "WithAuthorization" method embedded. So we can safely check this message.
	ADALResourceWithinError = "WithAuthorization"
)

type ARMOrResourceProviderErrorMessage string

var NRPHostNameResolutionFailedMessage ARMOrResourceProviderErrorMessage = "Host name resolution failed for 'Microsoft.Network'"

// ARMResourceType represents the different types of resources available in ARM
type ARMResourceType string

const (
	// Represent all resources in ARM regardless of the resource provider

	AllResources             ARMResourceType = "AllResources"
	PublicIPAddresses        ARMResourceType = "Microsoft.Network/publicIPAddresses"
	NetworkInterfaces        ARMResourceType = "Microsoft.Network/networkInterfaces"
	NetworkSecurityGroups    ARMResourceType = "Microsoft.Network/networkSecurityGroups"
	LoadBalancers            ARMResourceType = "Microsoft.Network/loadBalancers"
	VirtualNetworks          ARMResourceType = "Microsoft.Network/virtualNetworks"
	RouteTables              ARMResourceType = "Microsoft.Network/routeTables"
	PrivateEndpoints         ARMResourceType = "Microsoft.Network/privateEndpoints"
	PrivateLinkServices      ARMResourceType = "Microsoft.Network/privateLinkServices"
	VirtualMachineScaleSets  ARMResourceType = "Microsoft.Compute/virtualMachineScaleSets"
	VirtualMachines          ARMResourceType = "Microsoft.Compute/virtualMachines"
	Disks                    ARMResourceType = "Microsoft.Compute/disks"
	AvailabilitySet          ARMResourceType = "Microsoft.Compute/availabilitySets"
	Snapshots                ARMResourceType = "Microsoft.Compute/snapshots"
	StorageAccounts          ARMResourceType = "Microsoft.Storage/storageAccounts"
	StorageSyncServices      ARMResourceType = "Microsoft.StorageSync/storageSyncServices"
	VirtualMachineExtensions ARMResourceType = "Microsoft.Compute/virtualMachines/extensions"
	ManagedApplications      ARMResourceType = "Microsoft.Solutions/applications"
	OperationalInsights      ARMResourceType = "Microsoft.OperationalInsights/workspaces"
	RoleAssignments          ARMResourceType = "Microsoft.Authorization/roleAssignments"
	ResourceGroups           ARMResourceType = "Microsoft.Resources/resourceGroups"
	Certificates             ARMResourceType = "Microsoft.KeyVault/certificates"
	ContainerInsights        ARMResourceType = "Microsoft.OperationsManagement/solutions"
	PolicyAssignments        ARMResourceType = "Microsoft.Authorization/policyAssignments"
	DNSZones                 ARMResourceType = "Microsoft.Network/dnsZones"
	NATGateway               ARMResourceType = "Microsoft.Network/natGateways"
)

var (
	PublicIPAddressesError   = fmt.Errorf(string(PublicIPAddresses))
	LoadBalancerError        = fmt.Errorf(string(LoadBalancers))
	PrivateLinkServicesError = fmt.Errorf(string(PrivateLinkServices))
	PrivateEndpointsError    = fmt.Errorf(string(PrivateEndpoints))
	NATGatewayError          = fmt.Errorf(string(NATGateway))
)

// ARMResourceProvider represents the different provider types available in ARM
type ARMResourceProvider string

const (
	AllResourcesProviders ARMResourceProvider = "AllResourcesProviders"
	StorageProvider       ARMResourceProvider = "Microsoft.Storage"
	NetworkProvider       ARMResourceProvider = "Microsoft.Network"
	ComputeProvider       ARMResourceProvider = "Microsoft.Compute"
)

var resourceTypeClientErrorCodes = map[ARMResourceType]map[ARMOrResourceProviderErrorCode]bool{
	PublicIPAddresses:       {DNSRecordInUse: true, PublicIPCountLimitReached: true, StandardSkuPublicIPCountLimitReached: true, PublicIPAddressCannotBeDeleted: true},
	VirtualNetworks:         {VnetCountLimitReached: true, InUseSubnetCannotBeDeleted: true, VnetAddressSpaceOverlapsWithAlreadyPeeredVnet: true, VnetAddressSpacesOverlap: true, ReferencedResourceNotProvisioned: true},
	NetworkInterfaces:       {PrivateIPAddressNotInSubnet: true, InvalidResourceReference: true, ApplicationGatewaySubnetCannotBeUsedByOtherResources: true, SubnetIsFull: true, SubnetWithExternalResourcesCannotBeUsedByOtherResources: true, GatewaySubnet: true},
	LoadBalancers:           {PrivateIPAddressNotInSubnet: true, LoadBalancerInUseByVirtualMachineScaleSet: true},
	RouteTables:             {InUseRouteTableCannotBeDeleted: true},
	PrivateEndpoints:        {PrivateEndpointAndServiceEndpointsCannotCoexistInASubnet: true},
	NetworkSecurityGroups:   {InUseNetworkSecurityGroupCannotBeDeleted: true, NetworkSecurityGroupInUseByVirtualMachineScaleSet: true},
	StorageAccounts:         {MaxStorageAccountsCountPerSubscriptionExceeded: true},
	VirtualMachineScaleSets: {OverconstrainedAllocationRequest: true, SkuNotAvailable: true, SubscriptionNotRegistered: true, ScopeLocked: true, TooManyRequests: true, PublicIPCountLimitExceededByVMScaleSet: true, OverconstrainedZonalAllocationRequest: true, SpecifiedAllocatedOutboundPortsForOutboundRuleExceedsTotalNumberOfAvailablePorts: true, NetworkInterfaceAndLoadBalancerAreInDifferentAvailabilitySets: true},
	AvailabilitySet:         {SubscriptionNotRegistered: true},
	VirtualMachines:         {InvalidParameter: true, OperationNotAllowed: true, OverconstrainedAllocationRequest: true, SkuNotAvailable: true, SubscriptionNotRegistered: true, ResizeDiskError: true, AllocationFailed: true, QuotaExceeded: true, OverconstrainedZonalAllocationRequest: true},
	Snapshots:               {ResourceGroupDeletionTimeout: true},
	AllResources:            {RequestDisallowedByPolicy: true, ResourceGroupBeingDeleted: true, ResourceGroupNotFound: true, AuthorizationFailed: true, MissingSubscriptionRegistration: true, ScopeLocked: true, OperationPreempted: true, Canceled: true, CanceledAndSupersededDueToAnotherOperation: true, ResourceNotPermittedOnDelegatedSubnet: true, SubnetIsFull: true, ReservedResourceName: true, AppendPoliciesFieldsExist: true, PolicyViolation: true, ReadOnlyDisabledSubscription: true, SubscriptionRequestsThrottled: true},
	StorageSyncServices:     {MgmtDeletionDependency: true, BadRequest: true},
	ManagedApplications:     {ApplianceBeingDeleted: true, CustomerDeniedServicePlanAccess: true, ResourcePurchaseValidationFailed: true, MarketplacePurchaseEligibilityFailed: true},
	OperationalInsights:     {BadRequest: true},
	RoleAssignments:         {FailedRoleAssignmentCreation: true, PrincipalNotFound: true},
	ResourceGroups:          {ResourceGroupQuotaExceeded: true},
	Certificates:            {CertificateNotFound: true},
	ContainerInsights:       {CannotUpdatePlan: true, ResourcesBeingMoved: true, InvalidPolicyParameters: true},
	PolicyAssignments:       {ConflictingAppendPolicies: true},
	DNSZones:                {Throttled: true},
}

var retryableErrorCodes = map[ARMOrResourceProviderErrorCode]map[ARMResourceProvider]bool{
	MissingRegistrationForType:      {AllResourcesProviders: true},
	MissingSubscriptionRegistration: {AllResourcesProviders: true},
	StorageAccountNotRecognized:     {StorageProvider: true},
}

// used for double check those implicit client error.
var extraCheckErrorCodes = map[string]bool{
	ResourceGroupDeletionBlocked: true,
}

type VMExtensionErrorCode string

const (
	SystemctlEnableFailVMExtensionError            VMExtensionErrorCode = "SystemctlEnableFailVMExtensionError"            // Service could not be enabled by systemctl -- DEPRECATED
	SystemctlStartFailVMExtensionError             VMExtensionErrorCode = "SystemctlStartFailVMExtensionError"             // Service could not be started or enabled by systemctl
	CloudInitTimeoutVMExtensionError               VMExtensionErrorCode = "CloudInitTimeoutVMExtensionError"               // Timeout waiting for cloud-init runcmd to complete
	FileWatchTimeoutVMExtensionError               VMExtensionErrorCode = "FileWatchTimeoutVMExtensionError"               // Timeout waiting for a file
	HoldWalinuxagentVMExtensionError               VMExtensionErrorCode = "HoldWalinuxagentVMExtensionError"               // Unable to place walinuxagent apt package on hold during install
	ReleaseHoldWalinuxagentVMExtensionError        VMExtensionErrorCode = "ReleaseHoldWalinuxagentVMExtensionError"        // Unable to release hold on walinuxagent apt package after install
	AptInstallTimeoutVMExtensionError              VMExtensionErrorCode = "AptInstallTimeoutVMExtensionError"              // Timeout installing required apt packages
	EtcdDataDirNotFoundVMExtensionError            VMExtensionErrorCode = "EtcdDataDirNotFoundVMExtensionError"            // Etcd data dir not found
	EtcdRunningTimeoutVMExtensionError             VMExtensionErrorCode = "EtcdRunningTimeoutVMExtensionError"             // Timeout waiting for etcd to be accessible
	EtcdDownloadTimeoutVMExtensionError            VMExtensionErrorCode = "EtcdDownloadTimeoutVMExtensionError"            // Timeout waiting for etcd to download
	EtcdVolMountFailVMExtensionError               VMExtensionErrorCode = "EtcdVolMountFailVMExtensionError"               // Unable to mount etcd disk volume
	EtcdStartTimeoutVMExtensionError               VMExtensionErrorCode = "EtcdStartTimeoutVMExtensionError"               // Unable to start etcd runtime
	EtcdConfigFailVMExtensionError                 VMExtensionErrorCode = "EtcdConfigFailVMExtensionError"                 // Unable to configure etcd cluster
	DockerInstallTimeoutVMExtensionError           VMExtensionErrorCode = "DockerInstallTimeoutVMExtensionError"           // Timeout waiting for docker install
	DockerDownloadTimeoutVMExtensionError          VMExtensionErrorCode = "DockerDownloadTimeoutVMExtensionError"          // Timout waiting for docker download(s)
	DockerKeyDownloadTimeoutVMExtensionError       VMExtensionErrorCode = "DockerKeyDownloadTimeoutVMExtensionError"       // Timeout waiting to download docker repo key
	DockerAptKeyTimeoutVMExtensionError            VMExtensionErrorCode = "DockerAptKeyTimeoutVMExtensionError"            // Timeout waiting for docker apt-key
	DockerStartFailVMExtensionError                VMExtensionErrorCode = "DockerStartFailVMExtensionError"                // Docker could not be started by systemctl
	K8SRunningTimeoutVMExtensionError              VMExtensionErrorCode = "K8SRunningTimeoutVMExtensionError"              // Timeout waiting for k8s cluster to be healthy
	K8SDownloadTimeoutVMExtensionError             VMExtensionErrorCode = "K8SDownloadTimeoutVMExtensionError"             // Timeout waiting for Kubernetes download(s)
	KubectlNotFoundVMExtensionError                VMExtensionErrorCode = "KubectlNotFoundVMExtensionError"                // kubectl client binary not found on local disk
	ImgDownloadTimeoutVMExtensionError             VMExtensionErrorCode = "ImgDownloadTimeoutVMExtensionError"             // Timeout waiting for img utility download
	ContainerImageDownloadTimeoutVMExtensionError  VMExtensionErrorCode = "ContainerImageDownloadTimeoutVMExtensionError"  // Timeout waiting for a container image to pull
	KubeletStartFailVMExtensionError               VMExtensionErrorCode = "KubeletStartFailVMExtensionError"               // kubelet could not be started by systemctl
	CniDownloadTimeoutVMExtensionError             VMExtensionErrorCode = "CniDownloadTimeoutVMExtensionError"             // Timeout waiting for CNI download(s)
	MSProdDebDownloadError                         VMExtensionErrorCode = "MSProdDebDownloadError"                         // Timeout waiting for https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb
	MSProdDebPkgAddError                           VMExtensionErrorCode = "MSProdDebPkgAddError"                           // Failed to add repo pkg file
	FlexvolumeDownloadError                        VMExtensionErrorCode = "FlexvolumeDownloadError"                        // Failed to download flexvol drivers -- DEPRECATED
	ModprobeError                                  VMExtensionErrorCode = "ModprobeError"                                  // Unable to load a kernel module using modprobe
	SystemdInstallError                            VMExtensionErrorCode = "SystemdInstallError"                            // Failed to install systemd
	OutboundConnFailVMExtensionError               VMExtensionErrorCode = "OutboundConnFailVMExtensionError"               // Unable to establish outbound connection
	K8SAPIServerConnFailVMExtensionError           VMExtensionErrorCode = "K8SAPIServerConnFailVMExtensionError"           // Unable to establish connection to k8s api server
	K8SAPIServerDNSLookupFailVMExtensionError      VMExtensionErrorCode = "K8SAPIServerDNSLookupFailVMExtensionError"      // Unable to resolve k8s api server name
	K8SAPIServerAzureDNSLookupFailVMExtensionError VMExtensionErrorCode = "K8SAPIServerAzureDNSLookupFailVMExtensionError" // Unable to resolve k8s api server name due to Azure DNS issue
	KataKeyDownloadTimeoutVMExtensionError         VMExtensionErrorCode = "KataKeyDownloadTimeoutVMExtensionError"         // Timeout waiting to download kata repo key
	KataAptKeyTimeoutVMExtensionError              VMExtensionErrorCode = "KataAptKeyTimeoutVMExtensionError"              // Timeout waiting for kata apt-key
	KataInstallTimeoutVMExtensionError             VMExtensionErrorCode = "KataInstallTimeoutVMExtensionError"             // Timeout waiting for kata install
	ContainerdTimeoutVMExtensionError              VMExtensionErrorCode = "ContainerdTimeoutVMExtensionError"              // Timeout waiting for containerd download(s)
	CustomSearchDomainsFailVMExtensionError        VMExtensionErrorCode = "CustomSearchDomainsFailVMExtensionError"        // Unable to configure custom search domains
	GpuDriversStartFailVMExtensionError            VMExtensionErrorCode = "GpuDriversStartFailVMExtensionError"            // nvidia-modprobe could not be started by systemctl
	GpuDriversInstallTimeoutVMExtensionError       VMExtensionErrorCode = "GpuDriversInstallTimeoutVMExtensionError"       // Timeout waiting for GPU drivers install
	AptDailyTimeoutVMExtensionError                VMExtensionErrorCode = "AptDailyTimeoutVMExtensionError"                // Timeout waiting for apt daily updates
	AptUpdateTimeoutVMExtensionError               VMExtensionErrorCode = "AptUpdateTimeoutVMExtensionError"               // Timeout waiting for apt-get update to complete
	CseProvisionScriptNotReadyVMExtensionError     VMExtensionErrorCode = "CseProvisionScriptNotReadyVMExtensionError"     // Timeout waiting for cloud-init to place custom script on the vm
	MobyAptListVMExtensionError                    VMExtensionErrorCode = "MobyAptListVMExtensionError"                    // Timeout waiting to get moby-required apt repos
	MobyGPGKeyDownloadVMExtensionError             VMExtensionErrorCode = "MobyGPGKeyDownloadVMExtensionError"             // Timeout waiting to get moby-required repo GPG public key
	MobyInstallVMExtensionError                    VMExtensionErrorCode = "MobyInstallVMExtensionError"                    // Timeout waiting to install moby
)

// DefaultARMOperationTimeout defines a default (permissive) ARM operation timeout
const DefaultARMOperationTimeout = 150 * time.Minute

var vmExtensionErrorCodes = map[int]VMExtensionErrorCode{
	3:   SystemctlEnableFailVMExtensionError,
	4:   SystemctlStartFailVMExtensionError,
	5:   CloudInitTimeoutVMExtensionError,
	6:   FileWatchTimeoutVMExtensionError,
	7:   HoldWalinuxagentVMExtensionError,
	8:   ReleaseHoldWalinuxagentVMExtensionError,
	9:   AptInstallTimeoutVMExtensionError,
	10:  EtcdDataDirNotFoundVMExtensionError,
	11:  EtcdRunningTimeoutVMExtensionError,
	12:  EtcdDownloadTimeoutVMExtensionError,
	13:  EtcdVolMountFailVMExtensionError,
	14:  EtcdStartTimeoutVMExtensionError,
	15:  EtcdConfigFailVMExtensionError,
	20:  DockerInstallTimeoutVMExtensionError,
	21:  DockerDownloadTimeoutVMExtensionError,
	22:  DockerKeyDownloadTimeoutVMExtensionError,
	23:  DockerAptKeyTimeoutVMExtensionError,
	24:  DockerStartFailVMExtensionError,
	25:  MobyAptListVMExtensionError,
	26:  MobyGPGKeyDownloadVMExtensionError,
	27:  MobyInstallVMExtensionError,
	30:  K8SRunningTimeoutVMExtensionError,
	31:  K8SDownloadTimeoutVMExtensionError,
	32:  KubectlNotFoundVMExtensionError,
	33:  ImgDownloadTimeoutVMExtensionError,
	34:  KubeletStartFailVMExtensionError,
	35:  ContainerImageDownloadTimeoutVMExtensionError,
	41:  CniDownloadTimeoutVMExtensionError,
	42:  MSProdDebDownloadError,
	43:  MSProdDebPkgAddError,
	44:  FlexvolumeDownloadError,
	48:  SystemdInstallError,
	49:  ModprobeError,
	50:  OutboundConnFailVMExtensionError,
	51:  K8SAPIServerConnFailVMExtensionError,
	52:  K8SAPIServerDNSLookupFailVMExtensionError,
	53:  K8SAPIServerAzureDNSLookupFailVMExtensionError,
	60:  KataKeyDownloadTimeoutVMExtensionError,
	61:  KataAptKeyTimeoutVMExtensionError,
	62:  KataInstallTimeoutVMExtensionError,
	70:  ContainerdTimeoutVMExtensionError,
	80:  CustomSearchDomainsFailVMExtensionError,
	84:  GpuDriversStartFailVMExtensionError,
	85:  GpuDriversInstallTimeoutVMExtensionError,
	98:  AptDailyTimeoutVMExtensionError,
	99:  AptUpdateTimeoutVMExtensionError,
	100: CseProvisionScriptNotReadyVMExtensionError,
}

var vmExtensionClientErrorCodes = []VMExtensionErrorCode{
	OutboundConnFailVMExtensionError,
	K8SAPIServerConnFailVMExtensionError,
	K8SAPIServerDNSLookupFailVMExtensionError,
}

var clientErrorCodes = make(map[string]bool)

func init() {
	for _, v := range resourceTypeClientErrorCodes {
		for code := range v {
			clientErrorCodes[string(code)] = true
		}
	}
}
