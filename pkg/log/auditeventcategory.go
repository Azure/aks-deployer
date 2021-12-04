package log

import (
	"strings"

	"github.com/Azure/aks-deployer/pkg/consts"
)

// AuditEventCategory represents the category of an operation as per IfxAudit Event format
// https://genevamondocs.azurewebsites.net/collect/instrument/audit/onboarding.html
// https://microsoft.sharepoint.com/:w:/r/teams/WAG/EngSys/Monitor/_layouts/15/Doc.aspx?sourcedoc=%7B5B6CCA37-C4BF-477C-BD25-2BB98D2B83E3%7D&file=Cloud%20Services-%20Part-B%20-%20Audit%20Event%20Schema.docx&action=default&mobileredirect=true
type AuditEventCategory string

// audit event category types as defined in the Ifx documentation.
// little to no details on what they represent, so it is up to our interpretation
// we can request new types to be added.
const (
	Authentication        = AuditEventCategory("Authentication")
	Authorization         = AuditEventCategory("Authorization")
	UserManagement        = AuditEventCategory("UserManagement")
	GroupManagement       = AuditEventCategory("GroupManagement")
	RoleManagement        = AuditEventCategory("RoleManagement")
	ApplicationManagement = AuditEventCategory("ApplicationManagement")
	KeyManagement         = AuditEventCategory("KeyManagement")
	DirectoryManagement   = AuditEventCategory("DirectoryManagement")
	ResourceManagement    = AuditEventCategory("ResourceManagement")
	PolicyManagement      = AuditEventCategory("PolicyManagement")
	DeviceManagement      = AuditEventCategory("DeviceManagement")
	EntitlementManagement = AuditEventCategory("EntitlementManagement")
	PasswordManagement    = AuditEventCategory("PasswordManagement")
	ObjectManagement      = AuditEventCategory("ObjectManagement")
	IdentityProtection    = AuditEventCategory("IdentityProtection")
	Other                 = AuditEventCategory("Other")
)

// TODO(sterbrec): find a way to make this list dynamic, or test for completeness
var opHandlerCategories = map[string]AuditEventCategory{
	consts.GetSubscriptionHandlerName:                                   ResourceManagement,
	consts.PutSubscriptionHandlerName:                                   ResourceManagement,
	consts.GetContainerServiceHandlerName:                               ResourceManagement,
	consts.PutContainerServiceHandlerName:                               ResourceManagement,
	consts.DeleteContainerServiceHandlerName:                            ResourceManagement,
	consts.ListContainerServicesBySubscriptionHandlerName:               ResourceManagement,
	consts.ListContainerServicesByResourceGroupHandlerName:              ResourceManagement,
	consts.ListOrchestratorsHandlerName:                                 ResourceManagement,
	consts.GetOSOptionsHandlerName:                                      ResourceManagement,
	consts.GetContainerServiceUpgradeProfileHandlerName:                 ResourceManagement,
	consts.GetManagedClusterUpgradeProfileHandlerName:                   ResourceManagement,
	consts.BackfillManagedClusterHandlerName:                            ResourceManagement,
	consts.ReimageManagedClusterHandlerName:                             ResourceManagement,
	consts.GetManagedClusterDiagnosticsStateHandlerName:                 ResourceManagement,
	consts.GetOperationResultsHandlerName:                               ResourceManagement,
	consts.GetOperationStatusHandlerName:                                ResourceManagement,
	consts.PostDeploymentPreflightHandlerName:                           ResourceManagement,
	consts.JsonPatchManagedClusterHandlerName:                           ResourceManagement,
	consts.JsonPatchAgentPoolHandlerName:                                ResourceManagement,
	consts.JsonPatchControlPlaneHandlerName:                             ResourceManagement,
	consts.MockHandlerName:                                              Other,
	consts.InternalPutContainerServiceHandlerName:                       ResourceManagement,
	consts.InternalPutSubscriptionHandlerName:                           ResourceManagement,
	consts.InternalPutOperationStatusHandlerName:                        ResourceManagement,
	consts.HealthCheckHandlerName:                                       Other,
	consts.GetAvailableOperationsHandlerName:                            Other,
	consts.LinkedNotificationHandlerName:                                Other,
	consts.DelegateSubnetHandlerName:                                    Other,
	consts.UnDelegateSubnetHandlerName:                                  Other,
	consts.AdminListUnderlaysHandlerName:                                Other,
	consts.AdminGetUnderlayHandlerName:                                  Other,
	consts.AdminPutUnderlayHandlerName:                                  ResourceManagement,
	consts.AdminDeleteUnderlayHandlerName:                               ResourceManagement,
	consts.AdminPostUnderlayActionHandlerName:                           ResourceManagement,
	consts.GetManagedClusterHandlerName:                                 ResourceManagement,
	consts.GetAgentPoolHandlerName:                                      ResourceManagement,
	consts.GetAgentPoolUpgradeProfileHandlerName:                        ResourceManagement,
	consts.ListAgentPoolAvailableVersionsHandlerName:                    ResourceManagement,
	consts.GetManagedClusterAccessProfileHandlerName:                    ResourceManagement,
	consts.PutManagedClusterHandlerName:                                 ResourceManagement,
	consts.PutAgentPoolHandlerName:                                      ResourceManagement,
	consts.UpgradeNodeImageAgentPoolHandlerName:                         ResourceManagement,
	consts.PatchManagedClusterHandlerName:                               ResourceManagement,
	consts.DeleteManagedClusterHandlerName:                              ResourceManagement,
	consts.DeleteAgentPoolHandlerName:                                   ResourceManagement,
	consts.ListManagedClustersBySubscriptionHandlerName:                 ResourceManagement,
	consts.ListManagedClustersByResourceGroupHandlerName:                ResourceManagement,
	consts.ListCustomerControlPlanePodsHandlerName:                      ApplicationManagement,
	consts.ListAgentPoolsByClusterHandlerName:                           ResourceManagement,
	consts.GetMaintenanceConfigurationHandlerName:                       ResourceManagement,
	consts.PutMaintenanceConfigurationHandlerName:                       ResourceManagement,
	consts.DeleteMaintenanceConfigurationHandlerName:                    ResourceManagement,
	consts.ListMaintenanceConfigurationsByManagedClusterHandlerName:     ResourceManagement,
	consts.GetExtensionAddonHandlerName:                                 Other,
	consts.PutExtensionAddonHandlerName:                                 Other,
	consts.DeleteExtensionAddonHandlerName:                              Other,
	consts.ListExtensionAddonsByManagedClusterHandlerName:               Other,
	consts.GetCustomerControlPlanePodLogHandlerName:                     ApplicationManagement,
	consts.GetCustomerControlPlaneEventsHandlerName:                     ApplicationManagement,
	consts.PostKubectlCommandHandlerName:                                ApplicationManagement,
	consts.PostOverlayKubectlCommandHandlerName:                         ApplicationManagement,
	consts.ListUnderlaysHandlerName:                                     ApplicationManagement,
	consts.PostUnderlayKubectlCommandHandlerName:                        Other,
	consts.ListManagedClusterCredentialHandlerName:                      UserManagement,
	consts.PostVirtualMachineRunCommandHandlerName:                      ResourceManagement,
	consts.PostVirtualMachineGenericRunCommandHandlerName:               ResourceManagement,
	consts.GetCustomerAgentNodesStatusHandlerName:                       ResourceManagement,
	consts.ListManagedClusterClusterAdminCredentialHandlerName:          UserManagement,
	consts.ListManagedClusterClusterUserCredentialHandlerName:           UserManagement,
	consts.ListManagedClusterClusterMonitoringUserCredentialHandlerName: UserManagement,
	consts.ResetServicePrincipalProfileHandlerName:                      UserManagement,
	consts.ResetAADProfileHandlerName:                                   UserManagement,
	consts.ListDetectorsHandlerName:                                     Other,
	consts.GetDetectorHandlerName:                                       Other,
	consts.MigrateCustomerControlPlaneHandlerName:                       ApplicationManagement,
	consts.DeallocateControlPlaneHandlerName:                            ResourceManagement,
	consts.DrainCustomerControlPlanesHandlerName:                        ResourceManagement,
	consts.RotateClusterCertificatesHandlerName:                         KeyManagement,
	consts.GetServiceOutboundIPRangesHandlerName:                        Other,
	consts.PutServiceOutboundIPRangesHandlerName:                        Other,
	consts.DeleteServiceOutboundIPRangesHandlerName:                     Other,
	consts.ListServiceOutboundIPRangesHandlerName:                       Other,
	consts.RunCommandHandlerName:                                        ApplicationManagement,
	consts.RunCommandResultHandlerName:                                  ApplicationManagement,
	consts.GetSnapshotHandlerName:                                       ResourceManagement,
	consts.PutSnapshotHandlerName:                                       ResourceManagement,
	consts.DeleteSnapshotHandlerName:                                    ResourceManagement,
	consts.ListSnapshotsBySubscriptionHandlerName:                       ResourceManagement,
	consts.ListSnapshotsByResourceGroupHandlerName:                      ResourceManagement,
	consts.GetSnapshotAdminHandlerName:                                  ResourceManagement,
}

func getOperationCategory(operationName string) AuditEventCategory {
	parts := strings.Split(operationName, ".")
	if len(parts) < 2 {
		return Other
	}
	// operation name is [Admin].OperationHandlerName.VERB,
	// so we just want the HandlerName to map to the category.
	handlerName := parts[len(parts)-2]
	if category, ok := opHandlerCategories[handlerName]; ok {
		return category
	}
	return Other
}
