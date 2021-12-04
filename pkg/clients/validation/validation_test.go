package validation_test

import (
	om "github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/Azure/aks-deployer/pkg/clients/validation"
)

var _ = Describe("Validation", func() {
	Describe("ResourceGroup Validation", func() {
		It("validation.ValidateResourceGroupName", func() {
			err := validation.ValidateResourceGroupName("")
			Expect(err).NotTo(BeNil())

			err = validation.ValidateResourceGroupName("te:s;t")
			Expect(err).NotTo(BeNil())

			err = validation.ValidateResourceGroupName("abcdefghijklmlopqrstuvwxyzabcdefghijklmlopqrstuvwxyzabcdefghijklmlopqrstuvwxyzabcdefghijklmlopqrstuvwxyz")
			Expect(err).NotTo(BeNil())

			err = validation.ValidateResourceGroupName("test-resource-group")
			Expect(err).To(BeNil())
		})

		It("ValidateResourceGroupParameters", func() {
			parameters := &resources.Group{}
			err := validation.ValidateResourceGroupParameters(parameters)
			Expect(err).NotTo(BeNil())

			location := "westus"
			parameters = &resources.Group{
				Location: &location,
			}
			err = validation.ValidateResourceGroupParameters(parameters)
			Expect(err).To(BeNil())
		})
	})

	Describe("Deployment Validation", func() {
		It("validation.ValidateDeploymentName", func() {
			err := validation.ValidateDeploymentName("abcdefghijklmlopqrstuvwxyzabcdefghijklmlopqrstuvwxyz0000000000abcdefghi")
			Expect(err).NotTo(BeNil())

			err = validation.ValidateDeploymentName("test-deployment-name")
			Expect(err).To(BeNil())
		})

		It("validation.ValidateDeploymentParameters", func() {
			parameters := resources.Deployment{}
			err := validation.ValidateDeploymentParameters(parameters)
			Expect(err).NotTo(BeNil())

			parameters = resources.Deployment{
				Properties: &resources.DeploymentProperties{
					TemplateLink: &resources.TemplateLink{
						URI: nil,
					},
				},
			}
			err = validation.ValidateDeploymentParameters(parameters)
			Expect(err).NotTo(BeNil())

			parameters = resources.Deployment{
				Properties: &resources.DeploymentProperties{
					ParametersLink: &resources.ParametersLink{
						URI: nil,
					},
				},
			}
			err = validation.ValidateDeploymentParameters(parameters)
			Expect(err).NotTo(BeNil())

			parameters = resources.Deployment{
				Properties: &resources.DeploymentProperties{
					ParametersLink: nil,
					TemplateLink:   nil,
				},
			}
			err = validation.ValidateDeploymentParameters(parameters)
			Expect(err).To(BeNil())
		})
	})

	Describe("Solution Validation", func() {
		It("validation.ValidateSolutionParameters", func() {
			parameters := &om.Solution{
				Name:     to.StringPtr("test"),
				Type:     to.StringPtr("Microsoft.OperationsManagement/solutions"),
				Location: to.StringPtr("eastus"),
				Plan: &om.SolutionPlan{
					Name:          to.StringPtr("test"),
					Publisher:     to.StringPtr("Microsoft"),
					PromotionCode: to.StringPtr(""),
					Product:       to.StringPtr("OMSGallery/ContainerInsights"),
				},
				Properties: &om.SolutionProperties{
					WorkspaceResourceID: nil,
				},
			}

			err := validation.ValidateSolutionParameters(parameters)
			Expect(err).NotTo(BeNil())

			parameters.Properties.WorkspaceResourceID = to.StringPtr("/subscriptions/f3b504bb-826e-46c7-a1b7-674a5a0ae43a/resourceGroups/aks-cit/providers/Microsoft.OperationalInsights/workspaces/cit-workspace")
			err = validation.ValidateSolutionParameters(parameters)
			Expect(err).To(BeNil())
		})
	})
})
