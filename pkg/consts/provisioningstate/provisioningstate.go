package provisioningstate

// ProvisioningState is the ARM resource provisioning state
type ProvisioningState string

const (
	// Creating means resource is being created.
	Creating ProvisioningState = "Creating"
	// Updating means an existing resource is being updated
	Updating ProvisioningState = "Updating"
	// Scaling means an existing resource is being scaled only
	Scaling ProvisioningState = "Scaling"
	// Canceled means the resource provisioning has been cancelled
	Canceled ProvisioningState = "Canceled"
	// Failed means resource is in failed state
	Failed ProvisioningState = "Failed"
	// Succeeded means resource created succeeded during last create/update
	Succeeded ProvisioningState = "Succeeded"
	// Deleting means resource is in the process of being deleted
	Deleting ProvisioningState = "Deleting"
	// Migrating means resource is being migrated from one subscription or
	// resource group to another
	Migrating ProvisioningState = "Migrating"
	// Upgrading means an existing ContainerService resource is being upgraded
	Upgrading ProvisioningState = "Upgrading"
	// RefreshingServicePrincipalProfile means the cluster resource service principal is being refreshed
	RefreshingServicePrincipalProfile ProvisioningState = "RefreshingServicePrincipalProfile"
	// RefreshingAADProfile means the cluster resource AAD profile is being refreshed
	RefreshingAADProfile ProvisioningState = "RefreshingAADProfile"
	// MaintenanceMode is internal operation (not customer initiated) block update to managedCluster
	MaintenanceMode ProvisioningState = "Maintenance"
	// RotatingClusterCertificates means the cluster certificates are being rotated
	RotatingClusterCertificates ProvisioningState = "RotatingClusterCertificates"
	// ReconcilingClusterCertificates means the cluster certificates are being reconciling
	ReconcilingClusterCertificates ProvisioningState = "ReconcilingClusterCertificates"
	// RotatingUserProfileTokens means the cluster user profile tokens are being rotated
	RotatingUserProfileTokens ProvisioningState = "RotatingClusterStaticTokens"
)

// AgentPoolProvisioningState is the ARM resource provisioning state only for agent pool
type AgentPoolProvisioningState string

const (
	// UpgradingNodeImageVersion means the agent pool's node image version is being upgraded
	UpgradingNodeImageVersion AgentPoolProvisioningState = "UpgradingNodeImageVersion"
)

// IsFinal returns true for -ED states (Canceled, Failed, Succeeded)
// it returns false all for -ING states
func (state ProvisioningState) IsFinal() bool {
	return state == Succeeded || state == Failed || state == Canceled
}

func (state ProvisioningState) IsMaintenance() bool {
	return state == MaintenanceMode
}
