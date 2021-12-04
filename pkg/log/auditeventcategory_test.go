package log

import (
	"testing"

	"github.com/Azure/aks-deployer/pkg/consts"
	. "github.com/onsi/gomega"
)

func TestGetOperationCategory(t *testing.T) {
	g := NewGomegaWithT(t)
	expectedMapping := map[string]AuditEventCategory{
		"": Other,
		"SomeOperation.That.DoesNotExist.ButFollowsADotPattern":      Other,
		consts.DeleteAgentPoolOperationName:                          ResourceManagement,
		consts.PatchManagedClusterOperationName:                      ResourceManagement,
		consts.AdminDeleteUnderlayOperationName:                      ResourceManagement, // admin route
		consts.PostKubectlCommandOperationName:                       ApplicationManagement,
		consts.ListManagedClusterClusterAdminCredentialOperationName: UserManagement,
	}
	for k, v := range expectedMapping {
		g.Expect(getOperationCategory(k)).To(Equal(v), "when testing %v category", k)
	}
}
