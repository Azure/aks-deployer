package consts

type VMSSExtension struct {
	Name      string
	Publisher string
	Type      string
}

// VMSSLinuxExtensionAllowList defines the Linux extensions to be retained.
var VMSSLinuxExtensionAllowList = map[string]VMSSExtension{
	// The data plane team uses a custom vm extension for monitoring that's deployed out of band
	// from AKS RP.
	"akslinuxextension": {
		Name:      "AKSLinuxExtension",
		Publisher: "Microsoft.AKS",
		Type:      "Compute.AKS.Linux.AKSNode",
	},
	// The cosmic team's
	"cosmic.compute.linux.node": {
		Name:      "Cosmic.Compute.Linux.Node",
		Publisher: "Microsoft.M365",
		Type:      "Cosmic.Compute.Linux.Node",
	},
	// Two extensions for AzSecPack
	"microsoft.azure.security.monitoring.azuresecuritylinuxagent": {
		Name:      "Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent",
		Publisher: "Microsoft.Azure.Security.Monitoring",
		Type:      "AzureSecurityLinuxAgent",
	},
	"microsoft.azure.monitor.azuremonitorlinuxagent": {
		Name:      "Microsoft.Azure.Monitor.AzureMonitorLinuxAgent",
		Publisher: "Microsoft.Azure.Monitor",
		Type:      "AzureMonitorLinuxAgent",
	},
}

// VMSSWindowsExtensionAllowList defines the Windows extensions to be retained.
var VMSSWindowsExtensionAllowList = map[string]VMSSExtension{
	// The cosmic team's
	"cosmic.compute.windows.node": {
		Name:      "Cosmic.Compute.Windows.Node",
		Publisher: "Microsoft.M365",
		Type:      "Cosmic.Compute.Windows.Node",
	},
	// AzSecPack's Windows extension
	"microsoft.azure.geneva.genevamonitoring": {
		Name:      "Microsoft.Azure.Geneva.GenevaMonitoring",
		Publisher: "Microsoft.Azure.Geneva",
		Type:      "GenevaMonitoring",
	},
	// KeyVaultForWindows extension
	"microsoft.azure.keyvault.keyvaultforwindows": {
		Name:      "Microsoft.Azure.KeyVault.KeyVaultForWindows",
		Publisher: "Microsoft.Azure.KeyVault",
		Type:      "KeyVaultForWindows",
	},
	// CAPZ extension - DO NOT USE, Added temporarily to unblock Falcon team, can be removed without notice
	"microsoft.azure.containerupstream.capz.windows.bootstrapping": {
		Name:      "Microsoft.Azure.ContainerUpstream.CAPZ.Windows.Bootstrapping",
		Publisher: "Microsoft.Azure.ContainerUpstream",
		Type:      "CAPZ.Windows.Bootstrapping",
	},
}
