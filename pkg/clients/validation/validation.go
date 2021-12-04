package validation

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	om "github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/validation"
)

func ValidateResourceGroupName(resourceGroupName string) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateResourceGroupName", err.Error())
	}

	return nil
}

func ValidateResourceGroupParameters(parameters *resources.Group) error {
	if parameters == nil {
		return validation.NewError("validation", "ValidateResourceGroupParameters", "nil resource group parameters")
	}

	if err := validation.Validate([]validation.Validation{
		{TargetValue: *parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateResourceGroupParameters", err.Error())
	}

	return nil
}

func ValidateServicePrincipalsParameters(parameters graphrbac.ServicePrincipalCreateParameters) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AppID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateServicePrincipalsParameters", err.Error())
	}

	return nil
}

func ValidatePublicIPAddressParameters(parameters *network.PublicIPAddress) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return validation.NewError("validation", "ValidatePublicIPAddressParameters", err.Error())
	}

	return nil
}

func ValidateDeploymentName(deploymentName string) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deploymentName,
			Constraints: []validation.Constraint{{Target: "deploymentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "deploymentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "deploymentName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateDeploymentName", err.Error())
	}

	return nil
}

func ValidateDeploymentParameters(parameters resources.Deployment) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink.URI", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.Properties.ParametersLink", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.ParametersLink.URI", Name: validation.Null, Rule: true, Chain: nil}}}}}}}}); err != nil {
		return validation.NewError("validation", "ValidateDeploymentParameters", err.Error())
	}

	return nil
}

func ValidateVirtualMachineScaleSetParameters(parameters *compute.VirtualMachineScaleSet) error {
	if parameters == nil {
		return validation.NewError("validation", "ValidateVirtualMachineScaleSetParameters", "nil virtual machine scale set parameters")
	}

	if err := validation.Validate([]validation.Validation{
		{TargetValue: *parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxBatchInstancePercent", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxBatchInstancePercent", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
								{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxBatchInstancePercent", Name: validation.InclusiveMinimum, Rule: 5, Chain: nil},
							}},
							{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyInstancePercent", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyInstancePercent", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
									{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyInstancePercent", Name: validation.InclusiveMinimum, Rule: 5, Chain: nil},
								}},
							{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyUpgradedInstancePercent", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyUpgradedInstancePercent", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
									{Target: "parameters.VirtualMachineScaleSetProperties.UpgradePolicy.RollingUpgradePolicy.MaxUnhealthyUpgradedInstancePercent", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
								}},
						}},
					}},
				}}}}}); err != nil {
		return validation.NewError("validation", "ValidateVirtualMachineScaleSetParameters", err.Error())
	}

	return nil
}

func ValidateVirtualMachineRunCommandParameters(parameters compute.RunCommandInput) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.CommandID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateVirtualMachineRunCommandParameters", err.Error())
	}

	return nil
}

func ValidateSolutionParameters(parameters *om.Solution) error {
	if parameters == nil {
		return validation.NewError("validation", "ValidateSolutionParameters", "nil solution parameters")
	}

	if err := validation.Validate([]validation.Validation{
		{TargetValue: *parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.WorkspaceResourceID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return validation.NewError("validation", "ValidateSolutionParameters", err.Error())
	}

	return nil
}

func ValidateSecretName(secretName string) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: secretName,
			Constraints: []validation.Constraint{{Target: "secretName", Name: validation.Pattern, Rule: `^[0-9a-zA-Z-]+$`, Chain: nil}}},
	}); err != nil {
		return validation.NewError("validation", "ValidateSecretName", err.Error())
	}

	return nil
}

func ValidateSecretSetParameters(parameters keyvault.SecretSetParameters) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Value", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return validation.NewError("validation", "ValidateSecretSetParameters", err.Error())
	}

	return nil
}

func ValidateSecretMaxResults(maxResults *int32) error {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(25), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return validation.NewError("validation", "ValidateSecretMaxResults", err.Error())
	}
	return nil
}
