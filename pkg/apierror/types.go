// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package apierror

import "encoding/json"

// Error is the OData v4 format, used by the RPC and
// will go into the v2.2 Azure REST API guidelines
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Target  string    `json:"target,omitempty"`
	// Details is used to deliver structured error information from HCP API to RP.
	Details []Error `json:"details,omitempty" faker:"-"`

	Category      ErrorCategory `json:"-"`
	ExceptionType string        `json:"-"`
	// InternalDetail conveys information within RP layers.
	InternalDetail string `json:"-"`
	Subcode        string `json:"-"`
	// InnerMessage will deliver the same log information to QoSInfo when Operation generates an apiError with errorCode as
	// InternalOperationError and subCode as InternalServerError. This will help internal people to directly find out more
	// detailed information when they browse API QoS information without digging into RP log.
	InnerMessage string `json:"-"`
}

//Conform to the error interface
var _ error = &Error{}

// Error implements error interface to return error in json
func (e *Error) Error() string {
	output, err := json.MarshalIndent(e, " ", " ")
	if err != nil {
		return err.Error()
	}
	return string(output)
}

// Error is the OData v4 format, used by the RPC and
// will go into the v2.2 Azure REST API guidelines
type HttpErrorResponse struct {
	ErrorResponse  ErrorResponse
	HttpStatusCode int
}

// ErrorResponse  defines Resource Provider API 2.0 Error Response Content structure
type ErrorResponse struct {
	Body Error `json:"error"`
}

func (e *HttpErrorResponse) Unwrap() error {
	return &e.ErrorResponse
}

func (e *HttpErrorResponse) Error() string {
	return e.ErrorResponse.Error()
}

// Error implements error interface to return error in json
func (e *ErrorResponse) Error() string {
	return e.Body.Error()
}

func (e *ErrorResponse) Unwrap() error {
	return &e.Body
}
