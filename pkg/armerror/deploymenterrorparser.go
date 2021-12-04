//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package armerror

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/log"
)

const (
	vmExtensionErrorMessageFormat = apierror.VMExtensionErrorMessageExitStatus + "([0-9]+)"
	// Regex that looks for the pattern: vmssInstanceErrorCode=SomeErrorCode
	vmssProvisionErrorCodeFormat = apierror.VMSSInstanceErrorCode + "([[:alpha:]]+)"
	// Regex that looks for the pattern: Code="SomeErrorCode"
	clientErrorCodeEqualsMessageFormat = "Code=\"([[:alpha:]]+)\""
	// Regex that looks for the pattern: "code": "ErrorCode"
	clientErrorCodeJSONMessageFormat = "\"code\":[[:space:]]*\"([[:alpha:]]+)\""
)

var vmExtensionErrorMessageRegex *regexp.Regexp
var vmssProvisionErrorCodeRegex *regexp.Regexp
var clientErrorCodeEqualsMessageRegex *regexp.Regexp
var clientErrorCodeJSONMessageRegex *regexp.Regexp

func init() {
	vmExtensionErrorMessageRegex = regexp.MustCompile(vmExtensionErrorMessageFormat)
	vmssProvisionErrorCodeRegex = regexp.MustCompile(vmssProvisionErrorCodeFormat)
	clientErrorCodeEqualsMessageRegex = regexp.MustCompile(clientErrorCodeEqualsMessageFormat)
	clientErrorCodeJSONMessageRegex = regexp.MustCompile(clientErrorCodeJSONMessageFormat)
}

// ToErrorResponse parses ARM DeploymentOperationProperties object and
// determines if its an internal or client error.
func ToErrorResponse(ctx context.Context, logger *log.Logger, operation resources.DeploymentOperation) (apierror.ErrorResponse, error) {
	parentErrresp := apierror.ErrorResponse{}
	var resourceType string
	if operation.Properties != nil && operation.Properties.StatusMessage != nil {
		b, err := json.MarshalIndent(operation.Properties.StatusMessage, "", "  ")
		if err != nil {
			logger.Errorf(ctx,
				"Error occurred marshalling JSON - Error = '%v'", err)
			return parentErrresp, err
		}

		if err := json.Unmarshal(b, &parentErrresp); err != nil {
			logger.Errorf(ctx,
				"Error occurred unmarshalling JSON - Error = '%v' JSON = '%s'", err, string(b))
			return parentErrresp, err
		}
	}

	if operation.Properties != nil && operation.Properties.TargetResource != nil && operation.Properties.TargetResource.ResourceType != nil {
		resourceType = *operation.Properties.TargetResource.ResourceType
	}

	errresp := parentErrresp
	// In some cases ARM returns  "ResourceDeploymentFailure" as the top error code and error
	// code returned by the resource provider is in the child array "details". The following code
	// covers both of those scenarios. Also note, that details array is recursive i.e. the objects
	// in details array can contain another details array with no max depth of this tree enforced
	// by ARM. The following code only evaluates the first level property of "StatusMessage": "error"
	// and its child array: "details" for error codes.

	// If error code is ResourceDeploymentFailure then RP error is defined in the child object field: "details
	if parentErrresp.Body.Code == apierror.ErrorCode(ResourceDeploymentFailure) {
		// StatusMessage.error.details array supports multiple errors but in this particular case
		// DeploymentOperationProperties contains error from one specific resource type so the
		// chances of multiple deployment errors being returned for a single resource type is slim
		// (but possible) based on current error/QoS analysis. In those cases where multiple errors
		// are returned ACS will pick the first error code for determining whether this is an internal
		// or a client error. This can be reevaluated later based on practical experience.
		// However, note that customer will be returned the entire contents of "StatusMessage" object
		// (like before) so they have access to all the errors returned by ARM.
		logger.Infof(ctx,
			"Found ResourceDeploymentFailure error code - error response = '%+v'", parentErrresp)
		details := parentErrresp.Body.Details
		if len(details) > 0 {
			errresp.Body.Code = details[0].Code
			errresp.Body.Message = details[0].Message
			errresp.Body.Target = details[0].Target

			if resourceType != "" {
				parseErrorMessageForSubCode(ctx, logger, &errresp, ARMResourceType(resourceType))
			}
		}
	}

	if isCodeForClientError(ctx, logger, ARMResourceType(resourceType), ARMOrResourceProviderErrorCode(errresp.Body.Code)) {
		errresp.Body.Category = apierror.ClientError
	} else {
		errresp.Body.Category = apierror.InternalError
	}

	return errresp, nil
}

// isCodeForClientError checks the error code returned by ARM (from ARM/Resource Provider) and determines
// if the error should be returned as client error (in most cases this is because of a bad customer input
// or a bad request from RP standpoint, e.g. DnsRecordInUse)
func isCodeForClientError(ctx context.Context, logger *log.Logger, resourceType ARMResourceType, errorCode ARMOrResourceProviderErrorCode) bool {
	if codes, ok := resourceTypeClientErrorCodes[resourceType]; ok {
		if _, ok := codes[errorCode]; ok {
			logger.Infof(ctx, "Found client error code %s for resource type %s", errorCode, resourceType)
			return true
		}
	}

	if codes, ok := resourceTypeClientErrorCodes[AllResources]; ok {
		if _, ok := codes[errorCode]; ok {
			logger.Infof(ctx, "Found client error code %s for resource type %s", errorCode, resourceType)
			return true
		}
	}

	return false
}

// IsErrorCodeRetryable checks the error code returned by ARM (from ARM/Resource Provider) and determines
// if the error should be returned as retryable error so that the caller can retry ARM deployment operation
func IsErrorCodeRetryable(ctx context.Context, logger *log.Logger, resourceType ARMResourceType, errorCode ARMOrResourceProviderErrorCode) bool {
	parts := strings.Split(string(resourceType), "/")
	if len(parts) == 0 {
		logger.Errorf(ctx, "Unable to parse resource provider for resource type %s", resourceType)
		return false
	}

	resourceProvider := ARMResourceProvider(parts[0])
	if codes, ok := retryableErrorCodes[errorCode]; ok {
		if _, ok := codes[resourceProvider]; ok {
			logger.Infof(ctx, "Found retryable error code %s for resource type %s", errorCode, resourceType)
			return true
		}

		if _, ok := codes[AllResourcesProviders]; ok {
			logger.Infof(ctx, "Found retryable error code %s for resource type %s", errorCode, resourceType)
			return true
		}
	}

	return false
}

func parseErrorMessageForSubCode(ctx context.Context, logger *log.Logger, errresp *apierror.ErrorResponse, resourceType ARMResourceType) {
	if errresp == nil {
		return
	}
	if resourceType == VirtualMachineExtensions && ARMOrResourceProviderErrorCode(errresp.Body.Code) == VMExtensionProvisioningError {
		logger.Infof(ctx, "Found internal error code %s for resource type %s", errresp.Body.Code, resourceType)
		if subcode, ok := IsErrorStringVMExtensionError(errresp.Body.Message); ok {
			errresp.Body.Subcode = string(subcode)
		}
	}
}
