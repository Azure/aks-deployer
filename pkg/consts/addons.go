package consts

import (
	"strings"
)

// Hard coded key values, versions, and images for Addons
const (
	AddonNameOmsAgent                 = "omsagent"
	AddonConfigLogAnalyticsResourceID = "LogAnalyticsWorkspaceResourceID"
	AddonConfigOmsWorkspaceID         = "OmsWorkspaceID"
	AddonConfigOmsWorkspaceKey        = "OmsWorkspaceKey"
	AddonConfigOmsUseAADAuth          = "useAADAuth"
	AddonConfigAksClusterName         = "AksClusterName"
	AddonConfigAksResourceID          = "AksResourceID"
	AddonConfigAksNodeResourceGroup   = "AksNodeResourceGroup"
	AddonConfigAksRegion              = "AksRegion"

	AddonNameHTTPApplicationRouting                      = "HTTPApplicationRouting"
	AddonConfigHTTPApplicationRoutingZoneName            = "HTTPApplicationRoutingZoneName"
	AddonConfigHTTPApplicationRoutingExternalDNSImage    = "HTTPApplicationRoutingExternalDNSImage"
	AddonConfigHTTPApplicationRoutingDefaultBackendImage = "HTTPApplicationRoutingDefaultBackendImage"
	AddonConfigHTTPApplicationRoutingNginxIngressImage   = "HTTPApplicationRoutingNginxIngressImage"

	DefaultHTTPApplicationRoutingExternalDNSImage    = "mcr.microsoft.com/oss/kubernetes/external-dns:v0.4.8"
	DefaultHTTPApplicationRoutingDefaultBackendImage = "mcr.microsoft.com/oss/kubernetes/defaultbackend:1.4"
	DefaultHTTPApplicationRoutingNginxIngressImage   = "mcr.microsoft.com/oss/kubernetes/ingress/nginx-ingress-controller:0.19.0"

	AddonNameACIConnectorLinux        = "ACIConnectorLinux"
	AddonConfigACIConnectorSubnetName = "SubnetName"
	ACIConnectorLinuxOSType           = "Linux"

	DefaultACIConnectorLinuxImage = "mcr.microsoft.com/oss/virtual-kubelet/virtual-kubelet:latest"

	AddonNameKubeDashboard = "KubeDashboard"

	ACIConnectorWindowsOSType = "Windows"

	DefaultCalicoBlockSize = 26

	AddonNameAzurePolicy = "azurepolicy"

	AddonNameGitOps = "gitops"

	AddonNameExtensionManager = "extensionManager"

	AddonIngressApplicationGatewayAddonName               = "ingressApplicationGateway"
	AddonConfigIngressApplicationGatewayID                = "applicationGatewayId"
	AddonConfigIngressApplicationGatewayName              = "applicationGatewayName"
	AddonConfigIngressEffectiveApplicationGatewayID       = "effectiveApplicationGatewayId"
	AddonConfigIngressDefaultApplicationGatewayName       = "applicationgateway"
	AddonConfigIngressApplicationGatewaySubnetPrefix      = "subnetPrefix"
	AddonConfigIngressApplicationGatewaySubnetCIDR        = "subnetCidr"
	AddonConfigIngressApplicationGatewaySubnetID          = "subnetId"
	AddonConfigIngressApplicationGatewayShared            = "shared"
	AddonConfigIngressApplicationGatewayCreatedByTagKey   = "created-by"
	AddonConfigIngressApplicationGatewayCreatedByTagValue = "ingress-appgw"

	AddonOpenServiceMeshAddonName = "openServiceMesh"

	AddonNameACCSGXDevicePlugin         = "ACCSGXDevicePlugin"
	AddonConfigACCSGXQuoteHelperEnabled = "ACCSGXQuoteHelperEnabled"

	AddonNameAzureKeyvaultSecretsProvider = "azureKeyvaultSecretsProvider"

	AddonNameWindowsGmsa = "WindowsGmsa"

	AddonManagerModeReconcile    = "Reconcile"
	AddonManagerModeEnsureExists = "EnsureExists"

	DefaultCCPWebhookMutateNamespaces          = "kube-system,gatekeeper-system"
	DefaultGatekeeperWebhookExcludedNamespaces = "kube-system,gatekeeper-system,aks-periscope"
	CalicoNamespaces                           = "tigera-operator,calico-system"

	ExtensionTypeDapr = "Microsoft.Dapr"
)

var AddonNameToControlPlaneAddonName = map[string]string{
	strings.ToLower(AddonNameKubeDashboard):                  AddonNameKubeDashboard,
	strings.ToLower(AddonNameAzurePolicy):                    AddonNameAzurePolicy,
	strings.ToLower(AddonNameGitOps):                         AddonNameGitOps,
	strings.ToLower(AddonNameExtensionManager):               AddonNameExtensionManager,
	strings.ToLower(AddonIngressApplicationGatewayAddonName): AddonIngressApplicationGatewayAddonName,
	strings.ToLower(AddonOpenServiceMeshAddonName):           AddonOpenServiceMeshAddonName,
	strings.ToLower(AddonNameACCSGXDevicePlugin):             AddonNameACCSGXDevicePlugin,
	strings.ToLower(AddonNameWindowsGmsa):                    AddonNameWindowsGmsa,
}
