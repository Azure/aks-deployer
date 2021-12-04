// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package tokenprovider

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	aadEndpoint  = "aadEndpoint"
	tenantID     = "tenantID"
	clientID     = "clientID"
	resourceID   = "resourceID"
	clientSecret = "clientSecret"
)

var _ = Describe("Token Provider", func() {
	Context("init", func() {
		It("empty endpoint should return error", func() {
			_, err := NewServicePrincipalTokenProvider(tenantID, clientID, clientSecret, "")
			Expect(err).Should(HaveOccurred())
		})

		It("empty msiID should return error", func() {
			_, err := NewMsiTokenProvider("")
			Expect(err).Should(HaveOccurred())
		})

		It("should create a token provider", func() {
			_, err := NewServicePrincipalTokenProvider(tenantID, clientID, clientSecret, aadEndpoint)
			Expect(err).Should(BeNil())

			_, err = NewMsiTokenProvider(resourceID)
			Expect(err).Should(BeNil())
		})

	})
	Context("Get", func() {
		It("should require resource", func() {
			sp, err := NewServicePrincipalTokenProvider(tenantID, clientID, clientSecret, aadEndpoint)
			Expect(err).Should(BeNil())
			_, err = sp.GetToken("")
			Expect(err).Should(HaveOccurred())

			msi, err := NewMsiTokenProvider(resourceID)
			Expect(err).Should(BeNil())
			_, err = msi.GetToken("")
			Expect(err).Should(HaveOccurred())
		})

	})
})
