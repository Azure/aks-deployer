package categorizederror

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/armerror"
	"github.com/Azure/aks-deployer/pkg/consts"
	"github.com/Azure/aks-deployer/pkg/log"
)

const (
	ADALErrorStatusCodeFormat = " StatusCode=([0-9]+) "
	// when above format returned 0 after doing regular express matching, we need to check extraStatusCode
	ADALErrorExtraStatusCodeFormat = "Status Code = '([0-9]+)'"
	ADALErrorDescriptionCodeFormat = "\"error_description\":\"([a-zA-Z0-9]+):"
)

const ServiceErrCodeUnknown = "Unknown"

var ADALErrorStatusCodeRegex *regexp.Regexp
var ADALErrorDescriptionCodeRegex *regexp.Regexp
var ADALErrorExtraStatusCodeRegex *regexp.Regexp

func init() {
	ADALErrorStatusCodeRegex = regexp.MustCompile(ADALErrorStatusCodeFormat)
	ADALErrorDescriptionCodeRegex = regexp.MustCompile(ADALErrorDescriptionCodeFormat)
	ADALErrorExtraStatusCodeRegex = regexp.MustCompile(ADALErrorExtraStatusCodeFormat)
}

func getDependencyFromError(err error) ResourceType {
	if strings.Contains(err.Error(), armerror.ADALResourceWithinError) {
		return ADAL
	}

	return ARM
}

// define or extract subcode for adal error.
// when it is connectivity error, we define subcode and return as internal error
// when it is 500+ error, we return http corresponding error subcode and internal error. eg: 500, we return serviceUnavailable
// when it is 400+ error, we find the AADSTSxxx code and return it as subcode and left cateogry blank and let caller to handle it.
func getCategoryAndSubCodeFromADALError(err error) (apierror.ErrorCategory, ErrorSubCode) {
	errString := err.Error()
	if strings.Contains(errString, "EOF") {
		return apierror.InternalError, EOF
	}

	if strings.Contains(errString, "context canceled") {
		return apierror.InternalError, ContextCanceled
	}

	// read statusCode and if it is 500+, then return the code defined by http standard eg: 503 return "ServiceUnavailable"
	matchs := ADALErrorStatusCodeRegex.FindStringSubmatch(errString)
	if len(matchs) >= 2 {
		adalStatusCode := matchs[1]
		statusCode, err := strconv.Atoi(adalStatusCode)
		if err == nil {
			if statusCode >= 500 {
				return apierror.InternalError, ErrorSubCode(strings.ReplaceAll(http.StatusText(statusCode), " ", ""))
			} else if statusCode == 0 {
				matchs = ADALErrorExtraStatusCodeRegex.FindStringSubmatch(errString)
				if len(matchs) >= 2 {
					innerAdalStatusCode := matchs[1]
					statusCode, err = strconv.Atoi(innerAdalStatusCode)
					if err == nil {
						if statusCode >= 500 {
							return apierror.InternalError, ErrorSubCode(strings.ReplaceAll(http.StatusText(statusCode), " ", ""))
						}
					}
				}
			}
		}
	}

	// if we can get "AADSTSXXXXX" defined by ADAL and return it as subcode eg: "error_description":"AADSTS700016 return "AADSTS700016"
	// with this "AADSTSXXXXX", we can check the code definition by https://login.microsoftonline.com/error
	// Since we cannot make sure if the error returned is client error or our aks error, so just left it as blank
	matchs = ADALErrorDescriptionCodeRegex.FindStringSubmatch(errString)
	if len(matchs) >= 2 {
		adaldescriptionCode := matchs[1]
		if strings.HasPrefix(adaldescriptionCode, "AADSTS") {
			return apierror.ErrorCategory(""), ErrorSubCode(adaldescriptionCode)
		}
	}

	return apierror.InternalError, Unknown
}

// find generic category and subcode from error returned from armclient, not including adal error.
func getCategoryAndSubCodeFromError(resp *http.Response, err error) (apierror.ErrorCategory, ErrorSubCode) {
	if errors.Is(err, io.EOF) {
		return apierror.InternalError, EOF
	}

	if errors.Is(err, context.Canceled) {
		return apierror.InternalError, ContextCanceled
	}

	// check if connectivity errors
	for errString, errSubCode := range ConnectivityErrors {
		if strings.Contains(err.Error(), errString) && !strings.Contains(err.Error(), "VMExtensionProvisioningError") {
			return apierror.InternalError, errSubCode
		}
	}

	// VMExtensionProvisioningError special handling
	if strings.Contains(err.Error(), "VMExtensionProvisioningError") {
		if customerProvidedVMExtension(err) {
			return apierror.ClientError, ErrorSubCode("VMExtensionProvisioningError")
		}

		if strings.Contains(err.Error(), "VMExtensionCSEWindowsTroubleshoot") {
			return apierror.InternalError, ErrorSubCode("VMExtensionProvisioningError_Windows")
		}
	}

	// VMExtensionHandlerNonTransientError special handling
	if strings.Contains(err.Error(), "VMExtensionHandlerNonTransientError") {
		if allowedVMSSLinuxExtension(err) {
			return apierror.InternalError, ErrorSubCode("VMExtensionHandlerNonTransientError")
		}

		if allowedVMSSWindowsExtension(err) {
			return apierror.InternalError, ErrorSubCode("VMExtensionHandlerNonTransientError_Windows")
		}

		return apierror.ClientError, ErrorSubCode("VMExtensionHandlerNonTransientError")
	}

	// check client or server error via response
	if resp != nil {
		category := apierror.InternalError
		var code ErrorSubCode
		if resp.StatusCode >= 500 {
			code = ErrorSubCode(strings.ReplaceAll(http.StatusText(resp.StatusCode), " ", ""))
		} else if resp.StatusCode >= 400 {
			category = apierror.ClientError
			code = ErrorSubCode(strings.ReplaceAll(http.StatusText(resp.StatusCode), " ", ""))
		}

		if errCode := getErrorCode(err); errCode != "" {
			code = ErrorSubCode(errCode)
		}
		return category, code
	}

	// check if client errors
	if reason, isClientError := armerror.IsErrorStringClientError(err.Error()); isClientError {
		return apierror.ClientError, ErrorSubCode(reason)
	}

	if reason, isVMExtensionError := armerror.IsErrorStringVMExtensionError(err.Error()); isVMExtensionError {
		subCode := ErrorSubCode(reason)
		if armerror.IsVmExtensionErrorClientError(reason) {
			return apierror.ClientError, subCode
		}
		return apierror.InternalError, subCode
	}

	if reason, isVMSSProvisionError := armerror.IsErrorStringVMSSProvisionError(err.Error()); isVMSSProvisionError {
		subCode := ErrorSubCode(reason)
		return apierror.InternalError, subCode
	}

	if armerror.IsErrorPolicyDisallowedError(err.Error()) {
		return apierror.ClientError, RequestDisallowedByPolicy
	}

	if armerror.IsErrorPolicyViolationError(err.Error()) {
		return apierror.ClientError, PolicyViolation
	}

	if armerror.IsErrorStringDiskEncryptionSetError(err.Error()) {
		return apierror.ClientError, DiskEncryptionSetError
	}

	// extract error code from error msg when it is not known client error or internal error. eg: VMStartTimedout
	// By default we set error as internalError, we can set Client Error for some special cases when needed at IsErrorStringClientError()
	code := getErrorCode(err)
	if code != "" {
		return apierror.InternalError, ErrorSubCode(code)
	}
	return apierror.InternalError, Unknown
}

func getErrorCode(err error) string {
	var sErr *azure.RequestError
	var svErr *azure.ServiceError
	if errors.As(err, &sErr) {
		svErr = sErr.ServiceError
	}

	var code string
	if svErr != nil || errors.As(err, &svErr) {
		if svErr.Code != "" && svErr.Code != ServiceErrCodeUnknown {
			code = svErr.Code
		}

		// Get error code from Details best effort
		for _, detailMap := range svErr.Details {
			if detailCode, ok := detailMap["code"]; ok {
				code = appendSubCode(code, detailCode)
			}
		}
	}

	return code
}

func appendSubCode(code string, subcode interface{}) string {
	subcodeStr, ok := subcode.(string)
	if !ok || subcodeStr == "" || strings.EqualFold(subcodeStr, "null") || strings.Contains(strings.ToLower(code), strings.ToLower(subcodeStr)) {
		return code
	}

	if code == "" {
		return subcodeStr
	}

	code = fmt.Sprintf("%s_%s", code, subcodeStr)
	return code
}

// HandleErrorToCategorizedError defines categorizedError for error returned from arm client, we categorize it by finding category, subcode and dependency.
// including handling adal error.
func HandleErrorToCategorizedError(ctx context.Context, resp *http.Response, err error) *CategorizedError {
	var sErr *azure.RequestError // RequestError contains DetailedError and ServiceError
	var svErr *azure.ServiceError
	var dErr *autorest.DetailedError

	if errors.As(err, &sErr) {
		dErr = &sErr.DetailedError
		svErr = sErr.ServiceError
		log.GetLogger(ctx).Infof(ctx, "error can cast to requestError: %s", err)
	} else {
		log.GetLogger(ctx).Infof(ctx, "error cannot cast to requestError: %s", err)
	}

	if dErr != nil || errors.As(err, &dErr) {
		resp = dErr.Response
		log.GetLogger(ctx).Infof(ctx, "error can cast to detailedError: %s", dErr)
	} else {
		log.GetLogger(ctx).Infof(ctx, "error cannot cast to detailedError: %s", err)
	}

	// right now for logging purpose and then deciding what is the value.
	if svErr != nil || errors.As(err, &svErr) {
		log.GetLogger(ctx).Infof(ctx, "error can cast to serviceError: %s", svErr)
	} else {
		log.GetLogger(ctx).Infof(ctx, "error cannot cast to serviceError: %s", err)
	}

	var category apierror.ErrorCategory
	var subCode ErrorSubCode
	dep := getDependencyFromError(err)
	if dep == ADAL {
		category, subCode = getCategoryAndSubCodeFromADALError(err)
		category, subCode = treatInvalidParameterAsInternalError(category, subCode, err)
		if subCode != Unknown {
			return setRetriableBasedOnCategorizedError(NewCategorizedError(ctx, category, subCode, dep, err))
		}
	}

	category, subCode = getCategoryAndSubCodeFromError(resp, err)
	category, subCode = treatInvalidParameterAsInternalError(category, subCode, err)
	return setRetriableBasedOnCategorizedError(NewCategorizedError(ctx, category, subCode, dep, err))
}

// For ClientError, if the error sub-code is "InvalidParameter", we'll treat it as InternalError.
// This case means it is most likely a bug in our codes: either in the place we call the other service, or in
// our validation code that didn't reject the invalid parameter from the customer input. It'd be good to surface
// and fix those.
// if error is vm size not available type of issue, we need to treat it as client error since sku api at front end couldn't catch some
// vmsize availability.
func treatInvalidParameterAsInternalError(category apierror.ErrorCategory, subCode ErrorSubCode, err error) (apierror.ErrorCategory, ErrorSubCode) {
	if category == apierror.ClientError && subCode == InvalidParameter {
		if strings.Contains(err.Error(), "The requested VM size") && strings.Contains(err.Error(), "is not available in the current region") {
			return category, subCode
		}
		return apierror.InternalError, subCode
	}
	return category, subCode
}

func setRetriableBasedOnCategorizedError(cerr *CategorizedError) *CategorizedError {
	retriable, nonRetriable := true, false
	if cerr.Category == apierror.ClientError {
		cerr.Retriable = &nonRetriable
		return cerr
	}

	if strings.EqualFold(string(cerr.SubCode), "InvalidParameter") {
		cerr.Retriable = &nonRetriable
		return cerr
	}

	cerr.Retriable = &retriable
	return cerr
}

// for vmss extension, we name it as "vmssCSE"
// for vmas extension, the naming convention is "cse-agent-<nodeIndex>"
func customerProvidedVMExtension(err error) bool {
	return strings.Contains(err.Error(), "VM has reported a failure when processing extension") && !strings.Contains(err.Error(), "vmssCSE") && !strings.Contains(err.Error(), "cse-agent-")
}

func allowedVMSSLinuxExtension(err error) bool {
	for _, extension := range consts.VMSSLinuxExtensionAllowList {
		if strings.Contains(err.Error(), extension.Type) {
			return true
		}
	}
	return false
}

func allowedVMSSWindowsExtension(err error) bool {
	for _, extension := range consts.VMSSWindowsExtensionAllowList {
		if strings.Contains(err.Error(), extension.Type) {
			return true
		}
	}
	return false
}
