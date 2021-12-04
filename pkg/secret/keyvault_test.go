package secret

import (
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"

	"github.com/Azure/aks-deployer/pkg/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	aadEndpoint  = "aadEndpoint"
	tenantID     = "tenantID"
	clientID     = "clientID"
	clientSecret = "clientSecret"
)

var resource = "resource"
var logger = log.InitializeTestLogger()

var _ = Describe("Secret Provider", func() {
	Context("Client", func() {
		It("invalid resource should return error", func() {
			_, err := NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, clientID, clientSecret, "")
			Expect(err).Should(HaveOccurred())

			_, err = NewMsiKeyvaultSecretProvider("")
			Expect(err).Should(HaveOccurred())
		})

		It("empty clientID or secret should return error", func() {
			_, err := NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, "", clientSecret, resource)
			Expect(err).Should(HaveOccurred())

			_, err = NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, clientID, "", resource)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Get Secret", func() {
		It("invalid secret uri should return InvalidSecretURI", func() {
			kv, err := NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, clientID, clientSecret, resource)
			Expect(err).Should(BeNil())

			_, err = kv.Get(logger, "randomstring")
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(ErrInvalidSecretURI))

			_, err = kv.Get(logger, "https://myvault.vault.azure.net/secrets/foo/version/bar")
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(ErrInvalidSecretURI))
		})
	})

	Context("Set Secret", func() {
		It("invalid secret uri should return InvalidSecretURI", func() {
			kv, err := NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, clientID, clientSecret, resource)
			Expect(err).Should(BeNil())

			_, err = kv.Get(logger, "randomstring")
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(ErrInvalidSecretURI))

			_, err = kv.Set(logger, "https://myvault.vault.azure.net/secrets/foo/version", keyvault.SecretSetParameters{})
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(ErrInvalidSecretURI))
		})

		It("empty value should return ErrEmptyValue", func() {
			kv, err := NewClientSecretKeyvaultSecretProvider(aadEndpoint, tenantID, clientID, clientSecret, resource)
			Expect(err).Should(BeNil())

			_, err = kv.Set(logger, "https://myvault.vault.azure.net/secrets/foo", keyvault.SecretSetParameters{})
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(ErrEmptyValue))
		})
	})
})
