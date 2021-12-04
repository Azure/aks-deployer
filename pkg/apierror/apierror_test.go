// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package apierror

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Error", func() {
	It("should create invalid param error", func() {
		apiError := New(
			ClientError,
			InvalidParameter,
			"error test")

		Expect(apiError.Body.Code).Should(Equal(ErrorCode("InvalidParameter")))
	})

	It("should create scale down internal error", func() {
		apiError := New(
			ClientError,
			ScaleDownInternalError,
			"error test")

		Expect(apiError.Body.Code).Should(Equal(ErrorCode("ScaleDownInternalError")))
	})

	It("should create error with subcode", func() {
		apiErr := NewWithSubcode(
			InternalError,
			InternalOperationError,
			"foobarSubcode",
			"panic!",
		)

		Expect(apiErr.Body.Subcode).To(Equal("foobarSubcode"))
	})
})
