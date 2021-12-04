// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package datastructs

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RedactedString", func() {
	Context("String() func", func() {
		It("should return <REDACTED>", func() {
			var redactedString RedactedString = "hello"
			Expect(redactedString.String()).Should(Equal("<REDACTED>"))
		})

		It("should return <REDACTED> when formatted using %v by fmt", func() {
			var redactedString RedactedString = "hello"
			Expect(fmt.Sprintf("%v", redactedString)).Should(Equal("<REDACTED>"))
		})

		It("should return <REDACTED> when formatted using %+v by fmt", func() {
			type StructWithSecret struct {
				NotSecret string
				Secret    RedactedString
			}

			testStructValue := StructWithSecret{NotSecret: "not secret", Secret: "secret"}
			Expect(fmt.Sprintf("%+v", testStructValue)).Should(Equal("{NotSecret:not secret Secret:<REDACTED>}"))
		})
	})

	Context("RealString() func", func() {
		It("should return real string", func() {
			var redactedString RedactedString = "hello"
			Expect(redactedString.RealString()).Should(Equal("hello"))
		})
	})

	Context("json Marshal and Unmarshal", func() {
		It("should keep the original content when going through a json Marshal and Unmarshal round trip", func() {
			type StructWithSecret struct {
				NotSecret string
				Secret    RedactedString
			}

			originalStructValue := StructWithSecret{NotSecret: "not secret", Secret: "secret"}
			jsonBytes, error := json.Marshal(originalStructValue)
			Expect(error).To(BeNil())

			var unmarshaledStruct StructWithSecret
			error = json.Unmarshal(jsonBytes, &unmarshaledStruct)
			Expect(error).To(BeNil())
			Expect(unmarshaledStruct.NotSecret).Should(Equal(originalStructValue.NotSecret))
			Expect(unmarshaledStruct.Secret).Should(Equal(originalStructValue.Secret))
		})
	})
})
