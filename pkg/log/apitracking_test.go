package log

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/Azure/aks-deployer/pkg/consts"
)

var expected = uuid.FromStringOrNil("337a0c51-7d39-46f9-a4b7-ca73f54079f8")

const testUnderlayName string = "test-underlay-c14"
const testControlPlaneID string = "507c7f79bcf86cd7994f6c0e"

var _ = Describe("Api Tracking", func() {
	It("new api tracking positive case", func() {
		request, err := http.NewRequest(http.MethodGet, "http://localhost/subscription/subID?api-version=2016-09-30&server-fqdn=public", nil)
		// Expect(err).ShouldNot(HaveOccurred())
		if err != nil {
			Skip("Unable to talk to http://localhost/subscription/subID?api-version=2016-09-30")
		}

		request.Header.Set(RequestCorrelationIDHeader, expected.String())
		request.Header.Set(RequestARMClientRequestIDHeader, expected.String())
		request.Header.Set(RequestClientApplicationIDHeader, "app-id")
		request.Header.Set(AcceptLanguageHeader, "es")
		request.Header.Set(RequestClientPrincipalNameHeader, "user@site.com")
		request.Header.Set("User-Agent", "agent")

		pathParameters := make(map[string]string)
		pathParameters[consts.PathSubscriptionIDParameter] = expected.String()
		pathParameters[consts.PathResourceGroupNameParameter] = "groupName"
		pathParameters[consts.PathResourceNameParameter] = "resource"
		pathParameters[consts.PathPrivateEndpointConnectionNameParameter] = "privateEndpointConnection"
		pathParameters[consts.PathControlPlaneParameter] = "12345"
		apiTracking := NewAPITracking(pathParameters, request, consts.SubscriptionResourceFullPath)
		Expect(apiTracking.GetCorrelationID().String()).Should(Equal(expected.String()))
		Expect(apiTracking.GetClientRequestID().String()).Should(Equal(expected.String()))
		Expect(apiTracking.GetOperationID()).Should(Not(BeEmpty()))
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetSubscriptionOperationName))
		Expect(apiTracking.GetSubscriptionID()).Should(Equal(expected))
		Expect(apiTracking.GetResourceGroupName()).Should(Equal("groupName"))
		Expect(apiTracking.GetResourceName()).Should(Equal("resource"))
		Expect(apiTracking.GetAPIVersion()).Should(Equal("2016-09-30"))
		Expect(apiTracking.GetAcceptLanguage()).Should(Equal("es"))
		Expect(apiTracking.GetClientAppID()).Should(Equal("app-id"))
		Expect(apiTracking.GetClientPrincipalName()).Should(Equal("@site.com"))
		Expect(apiTracking.GetUserAgent()).Should(Equal("agent"))
		Expect(apiTracking.GetPrivateEndpointConnectionName()).Should(Equal("privateEndpointConnection"))
		Expect(apiTracking.GetDelayStartInSeconds()).To(Equal(uint(0)))
		Expect(apiTracking.GetControlPlaneID()).Should(Equal("12345"))
		Expect(apiTracking.GetCredentialServerFQDNFormat()).To(Equal("public"))
	})

	It("should honor DelayStartInSeconds if set", func() {
		// should init with 0
		apiTracking := NewAPITrackingFromParametersMap(map[string]interface{}{})
		Expect(apiTracking.GetDelayStartInSeconds()).To(Equal(uint(0)))

		// should honor it if set
		apiTracking.SetDelayStartInSeconds(60)
		Expect(apiTracking.GetDelayStartInSeconds()).To(Equal(uint(60)))

		// init through parameterMap
		paramMap := map[string]interface{}{
			DelayStartInSeconds: uint(120),
		}
		apiTracking2 := NewAPITrackingFromParametersMap(paramMap)
		Expect(apiTracking2.GetDelayStartInSeconds()).To(Equal(uint(120)))
	})

	It("grab host from URL.Host when request.Host is empty", func() {
		request, _ := http.NewRequest(http.MethodGet, "http://localhost/subscription/subID?api-version=2016-09-30", nil)
		pathParameters := make(map[string]string)
		apiTracking := NewAPITracking(pathParameters, request, consts.SubscriptionResourceFullPath)
		Expect(apiTracking.GetHost()).Should(Equal("localhost"))

		request.Host = ""
		apiTracking = NewAPITracking(pathParameters, request, consts.SubscriptionResourceFullPath)
		Expect(apiTracking.GetHost()).Should(Equal("localhost"))
	})

	It("new api tracking missing parameter case", func() {
		request, err := http.NewRequest(http.MethodGet, "http://api", nil)
		Expect(err).ShouldNot(HaveOccurred())
		pathParameters := make(map[string]string)

		apiTracking := NewAPITracking(pathParameters, request, "api")
		Expect(apiTracking.GetCorrelationID().String()).Should(Equal("00000000-0000-0000-0000-000000000000"))
	})

	It("new api tracking bad sub id should be ignored", func() {
		request, err := http.NewRequest(http.MethodGet, "http://api", nil)
		Expect(err).ShouldNot(HaveOccurred())
		pathParameters := make(map[string]string)
		pathParameters[consts.PathSubscriptionIDParameter] = "subid"

		apiTracking := NewAPITracking(pathParameters, request, "api")
		Expect(apiTracking).ShouldNot(BeNil())
	})

	It("api tracking not found case", func() {
		ctx := context.Background()
		observed, ok := GetAPITracking(ctx)
		Expect(observed).To(BeNil())
		Expect(ok).To(Equal(false))
	})

	It("api tracking roundtrip", func() {
		ctx := context.Background()
		apiTracking := &APITracking{}
		apiTracking.correlationID = expected
		ctx = WithAPITracking(ctx, apiTracking)
		observed, ok := GetAPITracking(ctx)

		Expect(observed.GetCorrelationID().String()).To(Equal(expected.String()))
		Expect(ok).To(Equal(true))
	})

	It("Multi Agent Pool api tracking operation name", func() {
		apiTracking := &APITracking{}
		apiTracking.setOperationName(http.MethodGet, consts.AgentPoolResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetAgentPoolOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPut, consts.AgentPoolResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutAgentPoolOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodDelete, consts.AgentPoolResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DeleteAgentPoolOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListAgentPoolsByClusterFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListAgentPoolsByClusterOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.GetAgentPoolUpgradeProfileFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetAgentPoolUpgradeProfileOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListAgentPoolAvailableVersionsFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListAgentPoolAvailableVersionsOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())
	})

	It("api tracking operation name", func() {
		apiTracking := &APITracking{}

		apiTracking.setOperationName(http.MethodGet, consts.OperationResultsResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetOperationResultsOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.OperationStatusResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetOperationStatusOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.SubscriptionResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetSubscriptionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPut, consts.SubscriptionResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutSubscriptionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ContainerServiceResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetContainerServiceOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPut, consts.ContainerServiceResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutContainerServiceOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodDelete, consts.ContainerServiceResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DeleteContainerServiceOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.GetManagedClusterResourceAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetManagedClusterAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPut, consts.GetManagedClusterResourceAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutManagedClusterAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.OperationStatusResourceAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetOperationStatusAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.ListManagedClusterResourcesBySubscriptionAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListManagedClustersBySubscriptionAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.ListManagedClusterResourcesByResourceGroupAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListManagedClustersByResourceGroupAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.ResetAADProfileFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ResetAADProfileOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPost, consts.ResetServicePrincipalProfileFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ResetServicePrincipalProfileOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.GetDetectorFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetDetectorOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListDetectorsFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListDetectorsOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.AdminListUnderlayFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminListUnderlaysOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.AdminGetUnderlayFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminGetUnderlayOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPut, consts.AdminUnderlayFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminPutUnderlayOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodDelete, consts.AdminUnderlayFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminDeleteUnderlayOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.AdminPostUnderlayActionFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminPostUnderlayActionOperationName))
		apiTracking.setOperationName(http.MethodPost, consts.MigrateCustomerControlPlaneAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.MigrateCustomerControlPlaneOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.GetAgentPoolResourcesAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetAgentPoolAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.ListAgentPoolsByClusterAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListAgentPoolsByClusterAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPut, consts.GetAgentPoolResourcesAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutAgentPoolAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.ReimageManagedClusterAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ReimageManagedClusterAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.JsonPatchManagedClusterAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.JsonPatchManagedClusterAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.JsonPatchAgentPoolAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.JsonPatchAgentPoolAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.JsonPatchControlPlaneAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.JsonPatchControlPlaneAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.DelegateSubnetAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DelegateSubnetAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.UnDelegateSubnetAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.UnDelegateSubnetAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.ListManagedClusterClusterAdminCredentialFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListManagedClusterClusterAdminCredentialOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPost, consts.ListManagedClusterClusterUserCredentialFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListManagedClusterClusterUserCredentialOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPost, consts.ListManagedClusterClusterMonitoringUserCredentialFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListManagedClusterClusterMonitoringUserCredentialOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListPrivateLinkServiceConnectionsFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListPrivateEndpointConnectionsOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPut, consts.PrivateLinkServiceConnectionFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutPrivateEndpointConnectionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.PrivateLinkServiceConnectionFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetPrivateEndpointConnectionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodDelete, consts.PrivateLinkServiceConnectionFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DeletePrivateEndpointConnectionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ServiceOutboundIPRangesFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetServiceOutboundIPRangesOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPut, consts.ServiceOutboundIPRangesFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutServiceOutboundIPRangesOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodDelete, consts.ServiceOutboundIPRangesFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DeleteServiceOutboundIPRangesOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.ListServiceOutboundIPRangesFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListServiceOutboundIPRangesOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodGet, consts.SnapshotResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetSnapshotOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodDelete, consts.SnapshotResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.DeleteSnapshotOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodPut, consts.SnapshotResourceFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.PutSnapshotOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListSnapshotResourcesBySubscriptionFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListSnapshotsBySubscriptionOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.ListSnapshotResourcesByResourceGroupFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.ListSnapshotsByResourceGroupOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())

		apiTracking.setOperationName(http.MethodGet, consts.GetSnapshotResourceAdminOperationFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.GetSnapshotAdminOperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeTrue())

		apiTracking.setOperationName(http.MethodPost, consts.MigrateClusterV2OperationRouteFullPath)
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.MigrateClusterV2OperationName))
		Expect(apiTracking.IsAdminOperation()).Should(BeFalse())
	})

	It("should inject the hcp underlay name into the context's logger", func() {
		// setup the context and logger
		ctx := context.Background()
		apitracking := NewAPITrackingFromParametersMap(nil)
		apitracking.SetHCPUnderlayName(testUnderlayName)
		ctx = WithAPITracking(ctx, apitracking)

		// extract it back out of the context
		apiTracking, ok := GetAPITracking(ctx)
		Expect(ok).To(Equal(true))
		uName := apiTracking.GetHCPUnderlayName()
		Expect(uName).To(Equal(testUnderlayName))
	})

	It("should inject the control plane ID into the context's logger", func() {
		// setup the context and logger
		ctx := context.Background()
		apitracking := NewAPITrackingFromParametersMap(nil)
		apitracking.SetHCPControlPlaneID(testControlPlaneID)
		ctx = WithAPITracking(ctx, apitracking)

		// extract it back out of the context
		apiTracking, ok := GetAPITracking(ctx)
		Expect(ok).To(Equal(true))
		cpid := apiTracking.GetHCPControlPlaneID()
		Expect(cpid).To(Equal(testControlPlaneID))
	})

	It("should return correct resultCodeDependency", func() {
		apitracking := NewAPITrackingFromParametersMap(nil)
		apitracking.AddResultCodeDependency("Microsoft.Foo/Bar")
		apitracking.AddResultCodeDependency("Microsoft.Bar/Foo")
		apitracking.AddResultCodeDependency("Microsoft.bar/foo")
		res := apitracking.GetResultCodeDependency()
		Expect(res).To(Equal("microsoft.bar/foo,microsoft.foo/bar"))
	})

	It("should keep clientPrincipalName for admin operation", func() {
		request, err := http.NewRequest(http.MethodGet, "http://localhost/admin/api/v1/locations/westus2/underlays?api-version=2016-09-30", nil)
		// Expect(err).ShouldNot(HaveOccurred())
		if err != nil {
			Skip("Unable to talk to http://localhost/admin/api/v1/locations/westus2/underlays/?api-version=2016-09-30")
		}

		request.Header.Set(RequestCorrelationIDHeader, expected.String())
		request.Header.Set(RequestARMClientRequestIDHeader, expected.String())
		request.Header.Set(RequestClientApplicationIDHeader, "app-id")
		request.Header.Set(AcceptLanguageHeader, "es")
		request.Header.Set(RequestClientPrincipalNameHeader, "user@site.com")
		request.Header.Set("User-Agent", "agent")

		pathParameters := make(map[string]string)
		pathParameters[consts.PathSubscriptionIDParameter] = expected.String()
		apiTracking := NewAPITracking(pathParameters, request, consts.AdminListUnderlayFullPath)
		Expect(apiTracking.GetCorrelationID().String()).Should(Equal(expected.String()))
		Expect(apiTracking.GetClientRequestID().String()).Should(Equal(expected.String()))
		Expect(apiTracking.GetOperationID()).Should(Not(BeEmpty()))
		Expect(apiTracking.GetOperationName()).Should(Equal(consts.AdminListUnderlaysOperationName))
		Expect(apiTracking.GetSubscriptionID()).Should(Equal(expected))
		Expect(apiTracking.GetAPIVersion()).Should(Equal("2016-09-30"))
		Expect(apiTracking.GetAcceptLanguage()).Should(Equal("es"))
		Expect(apiTracking.GetClientAppID()).Should(Equal("app-id"))
		Expect(apiTracking.GetClientPrincipalName()).Should(Equal("user@site.com"))
		Expect(apiTracking.GetUserAgent()).Should(Equal("agent"))
	})

	It("should serialize span context", func() {
		_, span := StartSpan(context.Background(), "name", AKSTeamUnknown)
		defer span.End()
		at := NewAPITrackingFromParametersMap(make(map[string]interface{}))
		at.SetSpanContext(span.SpanContext())
		sc, ok := at.GetSpanContext()
		Expect(ok).To(BeTrue())
		Expect(sc).To(BeEquivalentTo(span.SpanContext()))
	})

	It("should return correct fields for categorizedError", func() {
		apiTracking := NewAPITrackingFromParametersMap(nil)
		apiTracking.SetErrorCategory("InternalError")
		apiTracking.SetErrorSubcode("ServiceUnavailable")
		apiTracking.SetErrorDependency("Microsoft.Compute/VirtualMachineScaleSet")
		apiTracking.SetErrorAKSTeam("NodeProvisioning")
		Expect(apiTracking.GetErrorCategory()).To(Equal("InternalError"))
		Expect(apiTracking.GetErrorSubcode()).To(Equal("ServiceUnavailable"))
		Expect(apiTracking.GetErrorDependency()).To(Equal("Microsoft.Compute/VirtualMachineScaleSet"))
		Expect(apiTracking.GetErrorAKSTeam()).To(Equal("NodeProvisioning"))
	})

	It("set labels and get labels", func() {
		at := NewAPITrackingFromParametersMap(make(map[string]interface{}))
		agentPoolLabels := make(map[string]map[string]string)
		labels := make(map[string]string)
		labels["foo"] = "bar"
		agentPoolLabels["pool1"] = labels
		err := at.SetToBeDeletedAgentPoolLabels(agentPoolLabels)
		Expect(err).To(BeNil())
		getLabels, err := at.GetToBeDeletedAgentPoolLabels()
		Expect(err).To(BeNil())
		Expect(len(getLabels)).To(Equal(1))
		Expect(len(getLabels["pool1"])).To(Equal(1))
		Expect(getLabels["pool1"]["foo"]).To(Equal("bar"))
	})

	It("set empty labels string and get labels", func() {
		at := NewAPITrackingFromParametersMap(make(map[string]interface{}))
		agentPoolLabels := make(map[string]map[string]string)
		err := at.SetToBeDeletedAgentPoolLabels(agentPoolLabels)
		Expect(err).To(BeNil())
		getLabels, err := at.GetToBeDeletedAgentPoolLabels()
		Expect(err).To(BeNil())
		Expect(len(getLabels)).To(Equal(0))
	})

	It("not set labels and get labels", func() {
		at := NewAPITrackingFromParametersMap(make(map[string]interface{}))
		getLabels, err := at.GetToBeDeletedAgentPoolLabels()
		Expect(err).To(BeNil())
		Expect(len(getLabels)).To(Equal(0))
	})
})
