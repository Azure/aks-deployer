//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package armerror

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/go-autorest/autorest"
)

// ParseRawError is trying to deduce error category, error code, related resource type all from err message body
func ParseRawError(err error) (apierror.ErrorCategory, apierror.ErrorCode, string) {
	// set default
	category := apierror.InternalError
	code := apierror.InternalOperationError
	message := err.Error()

	if de, ok := err.(autorest.DetailedError); ok {
		if de.Original != nil {
			message = de.Original.Error()
		}
	}

	if reason, isClientError := IsErrorStringClientError(message); isClientError {
		category = apierror.ClientError
		code = apierror.ErrorCode(reason)
		return category, code, message
	}

	if reason, isVMSSProvisionError := IsErrorStringVMSSProvisionError(message); isVMSSProvisionError {
		code = apierror.ErrorCode(reason)
		return category, code, message
	}

	if reason, isVMExtensionError := IsErrorStringVMExtensionError(message); isVMExtensionError {
		code = apierror.ErrorCode(reason)
		if IsVmExtensionErrorClientError(reason) {
			category = apierror.ClientError
		}
	}

	if isPolicyDisallowedError := IsErrorPolicyDisallowedError(message); isPolicyDisallowedError {
		category := apierror.ClientError
		code := apierror.RequestDisallowedByPolicy
		return category, code, message
	}

	if isPolicyConflictError := IsErrorPolicyViolationError(message); isPolicyConflictError {
		category = apierror.ClientError
		code = apierror.PolicyViolation
		return category, code, message
	}

	// purposefully don't set error category as it is an opinionated business decision
	if isConflictError := IsErrorStringConflictError(message); isConflictError {
		code = apierror.Conflict
		return category, code, message
	}

	if IsErrorStringDiskEncryptionSetError(message) {
		category = apierror.ClientError
		return category, apierror.DiskEncryptionSetError, message
	}

	if IsErrorRoleAssignmentLimitExceededError(message) {
		category = apierror.ClientError
		code = apierror.RoleAssignmentLimitExceeded
		return category, code, message
	}

	if IsErrorObjectIsBeingDeletedError(message) {
		code = apierror.ObjectIsBeingDeleted
		return category, code, message
	}

	if IsErrorObjectIsDeletedButRecoverableError(message) {
		code = apierror.ObjectIsDeletedButRecoverable
		return category, code, message
	}

	//TODO some of these already return api errors. Make sure they all do then remove this function and let pass through below handle them
	if readinessErrCode, isReadinessCheckError := isErrorStringReadinessCheckError(message); isReadinessCheckError {
		code = readinessErrCode
	}

	//pass through apierrors
	var apierr *apierror.Error
	if errors.As(err, &apierr) {
		return apierr.Category, apierr.Code, apierr.Message
	}

	return category, code, message
}

// GetErrorCategoryfromArmResponse returns the error category for an ARM Response Error Code
func GetErrorCategoryfromArmResponse(parsedErrorCode ARMOrResourceProviderErrorCode) apierror.ErrorCategory {
	if _, isClientError := IsErrorStringClientError(string(parsedErrorCode)); isClientError {
		return apierror.ClientError
	}
	return apierror.InternalError
}

// IsErrorStringClientError determines an error string from azure sdk for go is client error or not, using default Arm list
func IsErrorStringClientError(err string) (string, bool) {
	// Checking both Regex patterns for matches
	//because sometimes we can get an error string with multiple occurrences of
	// either `Code="ErrorCode"` or `"code":"ErrorCode"`, and we want to catch as much as possible
	var foundErrorCodes []string
	// do we find something that matches the format: Code="ErrorCode"?
	if clientErrorCodeEqualsMessageRegex.MatchString(err) {
		foundErrorCodes = parseErrorCodesFromMatches(clientErrorCodeEqualsMessageRegex.FindAllStringSubmatch(err, -1), foundErrorCodes)
	}
	// or do we find something that matches the format: "code": "ErrorCode"?
	if clientErrorCodeJSONMessageRegex.MatchString(err) {
		foundErrorCodes = parseErrorCodesFromMatches(clientErrorCodeJSONMessageRegex.FindAllStringSubmatch(err, -1), foundErrorCodes)
	}
	// if no error codes found by regex, default to the naive way
	if len(foundErrorCodes) == 0 {
		return naiveClientErrorCodeFind(err)
	}

	//sanity check, make sure that one of the errorCodes we parse is actually a client error
	for _, code := range foundErrorCodes {
		if _, ok := clientErrorCodes[code]; ok {
			return code, true
		}
		if _, ok := extraCheckErrorCodes[code]; ok {
			return naiveClientErrorCodeFind(err)
		}
	}

	return "", false
}

func parseErrorCodesFromMatches(matches [][]string, errCodes []string) []string {
	for _, match := range matches {
		errCodes = append(errCodes, match[1])
	}
	return errCodes
}

func naiveClientErrorCodeFind(err string) (string, bool) {
	for code := range clientErrorCodes {
		if strings.Contains(err, code) {
			return code, true
		}
	}

	return "", false
}

// IsErrorStringVMSSProvisionError determines an error string is VMSSProvisionError or not
func IsErrorStringVMSSProvisionError(err string) (string, bool) {
	matches := vmssProvisionErrorCodeRegex.FindStringSubmatch(err)
	if len(matches) >= 2 {
		vmssProvisionError := matches[1]
		return vmssProvisionError, true
	}
	return "", false
}

// IsErrorStringVMExtensionError determines an error string is VMExtensionError or not
func IsErrorStringVMExtensionError(err string) (string, bool) {
	matchs := vmExtensionErrorMessageRegex.FindStringSubmatch(err)
	if len(matchs) >= 2 {
		vmExtnError := matchs[1]
		vmExtnErrorCode, err := strconv.Atoi(vmExtnError)
		if err == nil {
			if errorCodeString, ok := vmExtensionErrorCodes[vmExtnErrorCode]; ok {
				return string(errorCodeString), true
			}
		}
	}
	return "", false
}

// IsErrorPolicyViolationError determines if error string is a policy violation error or not
func IsErrorPolicyViolationError(err string) bool {
	if strings.Contains(err, string(PolicyViolation)) {
		return true
	}

	return false
}

// IsErrorStringConflictError determines if error string is a ConflictError or not
func IsErrorStringConflictError(err string) bool {
	if strings.Contains(err, string(Conflict)) {
		return true
	}

	return false
}

// IsVmExtensionErrorClientError determines an vm extension error is client error
func IsVmExtensionErrorClientError(code string) bool {
	for _, v := range vmExtensionClientErrorCodes {
		if strings.EqualFold(code, string(v)) {
			return true
		}
	}
	return false
}

// IsErrorStringDiskEncryptionSetError determines if error string is related to DiskEncryptionSet error
//   including following errors: KeyVaultAccessForbidden, DiskEncryptionSet not found
func IsErrorStringDiskEncryptionSetError(err string) bool {
	if strings.Contains(err, string(KeyVaultAccessForbidden)) {
		return true
	}

	if strings.Contains(err, string(DiskEncryptionSet)) {
		return true
	}
	return false
}

// isErrorStringReadinessCheckError determines if error string is related to readiness check
func isErrorStringReadinessCheckError(err string) (apierror.ErrorCode, bool) {
	knownReadinessCheckError := [...]apierror.ErrorCode{apierror.AgentCountNotMatch, apierror.NodesNotReady, apierror.ControlPlaneAPIServerNotReady}
	for _, knownErr := range knownReadinessCheckError {
		if strings.Contains(err, string(knownErr)) {
			return knownErr, true
		}
	}
	return apierror.ErrorCode(""), false
}

// IsErrorPolicyDisallowedError checks whether ARM error string contains RequestDisallowedByPolicy.
func IsErrorPolicyDisallowedError(err string) bool {
	if strings.Contains(err, string(RequestDisallowedByPolicy)) {
		return true
	}

	return false
}

// IsErrorRoleAssignmentLimitExceededError checks whether ARM error string contains RoleAssignmentLimitExceeded
func IsErrorRoleAssignmentLimitExceededError(err string) bool {
	if strings.Contains(err, string(RoleAssignmentLimitExceeded)) {
		return true
	}
	return false
}

// IsErrorObjectIsDeletedButRecoverableError checks whether ARM error string contains ObjectIsDeletedButRecoverable
func IsErrorObjectIsDeletedButRecoverableError(err string) bool {
	if strings.Contains(err, string(ObjectIsDeletedButRecoverable)) {
		return true
	}
	return false
}

// IsErrorObjectIsBeingDeletedError checks whether ARM error string contains ObjectIsBeingDeleted
func IsErrorObjectIsBeingDeletedError(err string) bool {
	if strings.Contains(err, string(ObjectIsBeingDeleted)) {
		return true
	}
	return false
}
