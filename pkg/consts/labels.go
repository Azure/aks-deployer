package consts

import "regexp"

// From https://github.com/Azure/aks-engine/blob/f8f373d241780a7945e40cdc9dc0d4b326d334b2/pkg/api/vlabs/validate.go#L95
const (
	LabelKeyPrefixMaxLength = 253
	LabelKeyNameMaxLength   = 63
	LabelValueMaxLength     = 63
	LabelKeyFormat          = "^(([a-zA-Z0-9-]+[.])*[a-zA-Z0-9-]+[/])?([A-Za-z0-9][-A-Za-z0-9_.]{0,61})?[A-Za-z0-9]$"
	LabelValueFormat        = "^([A-Za-z0-9][-A-Za-z0-9_.]{0,61})?[A-Za-z0-9]?$"
)

var (
	LabelKeyRegex   = regexp.MustCompile(LabelKeyFormat)
	LabelValueRegex = regexp.MustCompile(LabelValueFormat)
)

// Taints, Labels, Annotations: https://dev.azure.com/msazure/CloudNativeCompute/_wiki/wikis/CloudNativeCompute.wiki/51748/Taints-Labels-Annotations-(TLA)
const (
	// RP and overlaymgr need this for label validation and to pass to agentbaker +  webhook configmap
	AgentpoolLabelKey      = "agentpool" // Label used by aks-rp, but added by AgentBaker
	StorageProfileLabelKey = "storageprofile"
	StorageTierLabelKey    = "storagetier"

	// Special case system label users can set within the versioned datamodel (mode in the datamodel maps to this label)
	SystemPoolLabelKey   = AKSPrefix + "mode"
	SystemPoolLabelValue = "system"
	UserPoolLabelValue   = "user"

	// Labels with associated VMSS tag used in reconciler
	NodeImageVersionLabelKey        = AKSPrefix + "node-image-version"
	ServicePrincipalVersionLabelKey = AKSPrefix + "service-principal-version"
	WindowsPasswordVersionLabelKey  = AKSPrefix + "windows-password-version"
	WindowsGmsaVersionLabelKey      = AKSPrefix + "windows-gmsa-version"

	// General system labels
	PodNetworkSubscriptionLabelKey   = AKSPrefix + "podnetwork-subscription"
	PodNetworkResourceGroupLabelKey  = AKSPrefix + "podnetwork-resourcegroup"
	PodNetworkNameLabelKey           = AKSPrefix + "podnetwork-name"
	PodNetworkSubnetLabelKey         = AKSPrefix + "podnetwork-subnet"
	PodNetworkDelegationGuidLabelKey = AKSPrefix + "podnetwork-delegationguid"
	NetworkNameLabelKey              = AKSPrefix + "network-name"
	NetworkSubnetLabelKey            = AKSPrefix + "network-subnet"
	NetworkSubscriptionLabelKey      = AKSPrefix + "network-subscription"
	NetworkResourceGroupLabelKey     = AKSPrefix + "network-resourcegroup"
	NetworkVnetGuidLabelKey          = AKSPrefix + "nodenetwork-vnetguid"
	EnableAcrTeleportPluginLabelKey  = AKSPrefix + "enable-acr-teleport-plugin"
	ScaleSetPriorityLabelKey         = AKSPrefix + "scalesetpriority"
)

func GetK8sSystemLabelKeys() []string {
	return []string{
		"beta.kubernetes.io/arch",
		"beta.kubernetes.io/instance-type",
		"beta.kubernetes.io/os",
		"failure-domain.beta.kubernetes.io/region",
		"failure-domain.beta.kubernetes.io/zone",
		"failure-domain.kubernetes.io/zone",
		"failure-domain.kubernetes.io/region",
		"kubernetes.io/arch",
		"kubernetes.io/hostname",
		"kubernetes.io/os",
		"kubernetes.io/role",
		"kubernetes.io/instance-type",
		"node.kubernetes.io/instance-type",
		"topology.kubernetes.io/region",
		"topology.kubernetes.io/zone",
	}
}

func GetAgentBakerGeneratedLabelKeys() []string {
	return []string{
		"kubernetes.azure.com/role",
		"node-role.kubernetes.io/agent",
		"kubernetes.io/role",
		AgentpoolLabelKey,
		StorageProfileLabelKey,
		StorageTierLabelKey,
		"accelerator",
		"kubernetes.azure.com/fips_enabled",
		"kubernetes.azure.com/os-sku",
		"kubernetes.azure.com/cluster",
	}
}
