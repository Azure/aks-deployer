// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package apierror

// New creates an ErrorResponse
func New(errorCategory ErrorCategory, errorCode ErrorCode, message string) *ErrorResponse {
	return NewWithSubcode(errorCategory, errorCode, "", message)
}

// NewWithSubcode returns error response with Subcode
func NewWithSubcode(errorCategory ErrorCategory, errorCode ErrorCode, subcode, message string) *ErrorResponse {
	return &ErrorResponse{
		Body: Error{
			Code:     errorCode,
			Message:  message,
			Category: errorCategory,
			Subcode:  subcode,
		},
	}
}

// NewWithInnerMessage returns error response with Subcode and InnerMessage
func NewWithInnerMessage(errorCategory ErrorCategory, errorCode ErrorCode, subcode, message, innerMessage string) *ErrorResponse {
	return &ErrorResponse{
		Body: Error{
			Code:         errorCode,
			Message:      message,
			Category:     errorCategory,
			Subcode:      subcode,
			InnerMessage: innerMessage,
		},
	}
}

// NewWithSubcodeAndTarget returns error response with Subcode and target
func NewWithSubcodeAndTarget(errorCategory ErrorCategory, errorCode ErrorCode, subcode, target, message string) *ErrorResponse {
	return &ErrorResponse{
		Body: Error{
			Category: errorCategory,
			Code:     errorCode,
			Subcode:  subcode,
			Target:   target,
			Message:  message,
		},
	}
}

func WithHttpStatusCode(httpStatusCode int, err *ErrorResponse) *HttpErrorResponse {
	return &HttpErrorResponse{
		ErrorResponse:  *err,
		HttpStatusCode: httpStatusCode,
	}
}

// NewError creates Error object
func NewError(errorCategory ErrorCategory, errorCode ErrorCode, message string, subCode string, innerMsg string) *Error {
	return &Error{
		Code:         errorCode,
		Message:      message,
		Category:     errorCategory,
		Subcode:      subCode,
		InnerMessage: innerMsg,
	}
}
