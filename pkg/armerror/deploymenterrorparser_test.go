package armerror

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/consts/provisioningstate"
	"github.com/Azure/aks-deployer/pkg/log"
)

var (
	testLogger *log.Logger
)

var _ = BeforeSuite(func() {
	testLogger = log.InitializeTestLogger()
})

var _ = Describe("Test ToErrorResponse", func() {
	DescribeTable("Test ToErrorResponse", func(detailsCode string, detailsMessage string, statusCode string, statusMessage string, expectedErrorCode string, expectedMessage string) {
		state := string(provisioningstate.Failed)
		statusmsg := make(map[string]interface{})
		details := map[string]interface{}{
			"code":    detailsCode,
			"message": detailsMessage,
		}

		statusmsg["error"] = map[string]interface{}{
			"code":    statusCode,
			"message": statusMessage,
			"details": []map[string]interface{}{details},
		}
		operation := resources.DeploymentOperation{}
		operation.Properties = &resources.DeploymentOperationProperties{
			ProvisioningState: &state,
			StatusMessage:     &statusmsg,
		}

		errresp, err := ToErrorResponse(context.Background(), testLogger, operation)
		Expect(err).Should(BeNil())
		Expect(errresp.Body.Code).To(Equal(apierror.ErrorCode(expectedErrorCode)))
		Expect(errresp.Body.Message).To(Equal(expectedMessage))
	},
		Entry("Test SubscriptionNotRegistered", string(VMExtensionProvisioningError), "vm_error_msg", string(SubscriptionNotRegistered), "sub_error_msg", string(SubscriptionNotRegistered), "sub_error_msg"),
		Entry("Test ResourceDeploymentFailure", string(VMExtensionProvisioningError), "vm_error_msg", string(ResourceDeploymentFailure), "resource_error_msg", string(VMExtensionProvisioningError), "vm_error_msg"))
})

var _ = Describe("Test isCodeForClientError", func() {
	DescribeTable("Test whether code is client error", func(resourceType ARMResourceType, errorCode ARMOrResourceProviderErrorCode, matcher types.GomegaMatcher) {
		Expect(isCodeForClientError(context.Background(), testLogger, resourceType, errorCode)).To(matcher)
	},
		Entry(string(DNSRecordInUse), PublicIPAddresses, DNSRecordInUse, BeTrue()),
		Entry(string(SubscriptionNotRegistered), VirtualMachines, SubscriptionNotRegistered, BeTrue()),
		Entry(string(ResourceGroupBeingDeleted), AllResources, ResourceGroupBeingDeleted, BeTrue()),
		Entry(string(ResourceGroupBeingDeleted), ARMResourceType("fakeresource"), ResourceGroupBeingDeleted, BeTrue()),
		Entry(string(SubnetIsFull), NetworkInterfaces, SubnetIsFull, BeTrue()),
		Entry(string(SubnetWithExternalResourcesCannotBeUsedByOtherResources), NetworkInterfaces, SubnetWithExternalResourcesCannotBeUsedByOtherResources, BeTrue()),
		Entry(string(ResourceGroupQuotaExceeded), ResourceGroups, ResourceGroupQuotaExceeded, BeTrue()),
		Entry(string(GatewaySubnet), NetworkInterfaces, GatewaySubnet, BeTrue()),
		Entry("fakeerror", ARMResourceType("fakeresource"), ARMOrResourceProviderErrorCode("fackeerror"), BeFalse()),
		Entry("Empty ARMOrResourceProviderErrorCode", ARMResourceType(""), ARMOrResourceProviderErrorCode(""), BeFalse()),
		Entry(string(VnetCountLimitReached), PublicIPAddresses, VnetCountLimitReached, BeFalse()),
		Entry(string(SkuNotAvailable), AllResources, SkuNotAvailable, BeFalse()),
		Entry(string(AppendPoliciesFieldsExist), AllResources, AppendPoliciesFieldsExist, BeTrue()),
	)
})

var _ = Describe("Test IsErrorCodeRetryable", func() {
	DescribeTable("Test IsErrorCodeRetryable", func(resourceType ARMResourceType, errorCode ARMOrResourceProviderErrorCode, matcher types.GomegaMatcher) {
		Expect(IsErrorCodeRetryable(context.Background(), testLogger, resourceType, errorCode)).To(matcher)
	},
		Entry("Test StorageAccountNotRecognized - StorageAccounts", StorageAccounts, StorageAccountNotRecognized, BeTrue()),
		Entry("Test MissingRegistrationForType", ARMResourceType(AllResourcesProviders+"/type"), MissingRegistrationForType, BeTrue()),
		Entry("Test fakeresource", ARMResourceType("fakeresource"), ARMOrResourceProviderErrorCode("fackeerror"), BeFalse()),
		Entry("Test empty string", ARMResourceType(""), ARMOrResourceProviderErrorCode(""), BeFalse()),
		Entry("Test VnetCountLimitReached", StorageAccounts, VnetCountLimitReached, BeFalse()),
		Entry("Test StorageAccountNotRecognized - VirtualMachines", VirtualMachines, StorageAccountNotRecognized, BeFalse()))
})

var _ = Describe("Test GetErrorCategoryfromArmResponse function", func() {
	DescribeTable("Test GetErrorCategoryFromArmResponse", func(errorCode ARMOrResourceProviderErrorCode, expectedCategory apierror.ErrorCategory) {
		actualErrorCategory := GetErrorCategoryfromArmResponse(errorCode)
		Expect(actualErrorCategory).To(Equal(expectedCategory))
	},
		Entry("Test DnsRecordInUse", DNSRecordInUse, apierror.ClientError),
		Entry("Test SubscriptionNotRegistered", SubscriptionNotRegistered, apierror.ClientError),
		Entry("Test QuotaExceeded", QuotaExceeded, apierror.ClientError),
		Entry("Test fakeerror", ARMOrResourceProviderErrorCode("fakeerror"), apierror.InternalError))
})

var _ = Describe("Test GetErrorCategoryfromArmResponse function", func() {
	DescribeTable("Test GetErrorCategoryfromArmResponse", func(errorCode string, errorMessage string, resourceType ARMResourceType, expectedSubcode VMExtensionErrorCode) {
		errresp := apierror.ErrorResponse{
			Body: apierror.Error{
				Code:    apierror.ErrorCode(errorCode),
				Message: errorMessage,
			},
		}
		parseErrorMessageForSubCode(context.Background(), testLogger, &errresp, resourceType)
		Expect(VMExtensionErrorCode(errresp.Body.Subcode)).To(Equal(expectedSubcode))
	},
		Entry("Test VMextensionProvisioningError", "VMExtensionProvisioningError", "VM has reported a failure when processing extension 'cse-agent-2'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=9\n[stdout]\n\n[stderr]\n\".", VirtualMachineExtensions, AptInstallTimeoutVMExtensionError),
		Entry("Test not set Subcode - VirtualMachines", "VMExtensionProvisioningError", "VM has reported a failure when processing extension 'cse-agent-2'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=9\n[stdout]\n\n[stderr]\n\".", VirtualMachines, VMExtensionErrorCode("")),
		Entry("Test not set Subcode - VirtualMachineExtensions", "OverconstrainedAllocationRequest", "VM has reported a failure when processing extension 'cse-agent-2'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=9\n[stdout]\n\n[stderr]\n\".", VirtualMachineExtensions, VMExtensionErrorCode("")))
})
