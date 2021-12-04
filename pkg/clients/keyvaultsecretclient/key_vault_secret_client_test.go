package keyvaultsecretclient_test

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/Azure/aks-deployer/pkg/apierror"
	cgerror "github.com/Azure/aks-deployer/pkg/categorizederror"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/aks-deployer/pkg/clients/keyvaultsecretclient"
	"github.com/Azure/aks-deployer/pkg/log"
)

var _ = Describe("KeyVaultSecretClient", func() {
	Describe("Real key vault test", func() {
		BeforeEach(func() {
			if tokenString == "" {
				Skip("Token is not set. The integration test is skipped.")
			}
		})

		It("Should not return error when getting nonexistent secret", func() {
			result, rErr := keyVaultSecretClient.GetSecret(ctx, keyVaultBaseURL, secretName, "")
			Expect(rErr).To(BeNil())
			Expect(result).To(BeNil())
		})

		It("Should successfully set secret", func() {
			parameter := keyvault.SecretSetParameters{
				Value: &secretValue,
				Tags: map[string]*string{
					secretTagName: &secretTagValue,
				},
			}
			result, rErr := keyVaultSecretClient.SetSecret(ctx, keyVaultBaseURL, secretName, parameter)
			Expect(rErr).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Value).ToNot(BeNil())
			Expect(*result.Value).To(Equal(secretValue))
			Expect(result.Tags).ToNot(BeNil())
			Expect(result.Tags[secretTagName]).ToNot(BeNil())
			Expect(*result.Tags[secretTagName]).To(Equal(secretTagValue))
		})

		It("Should successfully get secret after setting", func() {
			result, rErr := keyVaultSecretClient.GetSecret(ctx, keyVaultBaseURL, secretName, "")
			Expect(rErr).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Value).ToNot(BeNil())
			Expect(*result.Value).To(Equal(secretValue))
			Expect(result.Tags).ToNot(BeNil())
			Expect(result.Tags[secretTagName]).ToNot(BeNil())
			Expect(*result.Tags[secretTagName]).To(Equal(secretTagValue))
		})

		It("Should successfully get secrets after setting", func() {
			result, rErr := keyVaultSecretClient.GetSecrets(ctx, keyVaultBaseURL, &maxResults)
			Expect(rErr).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Value).ToNot(BeNil())
			var found bool
			for _, item := range *result.Value {
				if item.Tags[secretTagName] != nil && *item.Tags[secretTagName] == secretTagValue {
					found = true
				}
			}
			Expect(found).To(BeTrue())
		})

		It("Should successfully delete the secret", func() {
			result, rErr := keyVaultSecretClient.DeleteSecret(ctx, keyVaultBaseURL, secretName)
			Expect(rErr).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
	})

	Describe("Validation test", func() {
		fakeAuthorizer := autorest.NewAPIKeyAuthorizer(nil, nil)
		fakeKVClient := keyvaultsecretclient.New(fakeAuthorizer, "centralus")

		It("Should report error for invalid secret name", func() {
			parameter := keyvault.SecretSetParameters{
				Value: &secretValue,
			}
			result, rErr := fakeKVClient.SetSecret(ctx, keyVaultBaseURL, "invalid_*^_^*_secret_0_0_name", parameter)
			Expect(rErr).ToNot(BeNil())
			var cerr *cgerror.CategorizedError
			Expect(errors.As(rErr.Error, &cerr)).To(BeTrue())
			Expect(cerr.Category).To(Equal(apierror.InternalError))
			Expect(cerr.SubCode).To(Equal(cgerror.FailedValidateSecretName))
			Expect(cerr.Dependency).To(Equal(cgerror.KeyVaultSecret))
			Expect(*cerr.Retriable).To(BeFalse())
			Expect(rErr.Error.Error()).To(ContainSubstring("validation#ValidateSecretName"))
			Expect(result).To(BeNil())
		})

		It("Should report error for invalid secret parameter", func() {
			inValidParameter := keyvault.SecretSetParameters{
				Value: nil,
			}
			result, rErr := fakeKVClient.SetSecret(ctx, keyVaultBaseURL, "secretName", inValidParameter)
			Expect(rErr).ToNot(BeNil())
			var cerr *cgerror.CategorizedError
			Expect(errors.As(rErr.Error, &cerr)).To(BeTrue())
			Expect(cerr.Category).To(Equal(apierror.InternalError))
			Expect(cerr.SubCode).To(Equal(cgerror.FailedValidateSecretSetParameters))
			Expect(cerr.Dependency).To(Equal(cgerror.KeyVaultSecret))
			Expect(*cerr.Retriable).To(BeFalse())
			Expect(rErr.Error.Error()).To(ContainSubstring("validation#ValidateSecretSetParameters"))
			Expect(result).To(BeNil())
		})
	})

	Describe("e2e with stubbed key vault server", func() {
		It("Should successfully get a mock secret", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.String()).To(Equal("/secrets/fooSecretName/barSecretVersion?api-version=2016-10-01"))
				rw.Write([]byte(`
				{
					"value": "fakeSecretVaule",
					"id": "https://unknownbaseURL/secrets/fooSecretName/barSecretVersion",
					"attributes": {
					  "enabled": true,
					  "created": 1493938410,
					  "updated": 1493938410,
					  "recoveryLevel": "Recoverable+Purgeable"
					}
				  }`))
				rw.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			secretBundle, rErr := kvSecretClient.GetSecret(ctx, server.URL, "fooSecretName", "barSecretVersion")
			Expect(rErr).To(BeNil())
			Expect(secretBundle).NotTo(BeNil())
			Expect(*secretBundle.ID).NotTo(BeNil())
			Expect(*secretBundle.ID).To(Equal("https://unknownbaseURL/secrets/fooSecretName/barSecretVersion"))
			Expect(*secretBundle.Value).NotTo(BeNil())
			Expect(*secretBundle.Value).To(Equal("fakeSecretVaule"))
		})

		It("Should successfully get mock secrets", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.String()).To(Equal("/secrets?api-version=2016-10-01&maxresults=25"))
				rw.Write([]byte(`
				{
					"value": [
					  {
						"contentType": "plainText",
						"id": "https://testvault1021.vault.azure.net/secrets/listsecrettest0",
						"attributes": {
						  "enabled": true,
						  "created": 1482189047,
						  "updated": 1482189047
						}
					  }
					],
					"nextLink": "https://testvault1021.vault.azure.net:443/secrets?api-version=7.0&$skiptoken=eyJOZXh0TWFya2VyIjoiMiE4OCFNREF3TURJeUlYTmxZM0psZEM5TVNWTlVVMFZEVWtWVVZFVlRWREVoTURBd01ESTRJVEl3TVRZdE1USXRNVGxVTWpNNk1UQTZORFV1T0RneE9ERXhNRm9oIiwiVGFyZ2V0TG9jYXRpb24iOjB9&maxresults=1"
				  }`))
				rw.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			var maxResults int32 = 25
			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			listResult, rErr := kvSecretClient.GetSecrets(ctx, server.URL, &maxResults)
			Expect(rErr).To(BeNil())
			Expect(listResult).NotTo(BeNil())
			Expect(listResult.Value).NotTo(BeNil())
			Expect(len(*listResult.Value)).To(Equal(1))
			Expect((*listResult.Value)[0].ID).NotTo(BeNil())
			Expect(*((*listResult.Value)[0].ID)).To(Equal("https://testvault1021.vault.azure.net/secrets/listsecrettest0"))
		})

		It("Should successfully delete a mock secret", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.String()).To(Equal("/secrets/fooSecretName?api-version=2016-10-01"))
				rw.Write([]byte(`
				{
					"recoveryId": "https://unknownBaseURL/deletedsecrets/fooSecretName",
					"deletedDate": 1493938433,
					"scheduledPurgeDate": 1501714433,
					"id": "https://unknownbaseURL/secrets/fooSecretName/barSecretVersion",
					"attributes": {
					  "enabled": true,
					  "created": 1493938433,
					  "updated": 1493938433,
					  "recoveryLevel": "Recoverable+Purgeable"
					}
				  }`))
				rw.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			secretBundle, rErr := kvSecretClient.DeleteSecret(ctx, server.URL, "fooSecretName")
			Expect(rErr).To(BeNil())
			Expect(secretBundle).NotTo(BeNil())
			Expect(*secretBundle.ID).To(Equal("https://unknownbaseURL/secrets/fooSecretName/barSecretVersion"))
			Expect(*secretBundle.RecoveryID).To(Equal("https://unknownBaseURL/deletedsecrets/fooSecretName"))
		})

		It("Should successsfully send a PUT request", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.String()).To(Equal("/secrets/fooSecretName?api-version=2016-10-01"))
				requestBody, err := ioutil.ReadAll(req.Body)
				Expect(err).To(BeNil())
				var parameter keyvault.SecretSetParameters
				Expect(json.Unmarshal(requestBody, &parameter)).To(BeNil())
				Expect(*parameter.Value).To(Equal("fakeSecretValue"))
				rw.Write([]byte(`
				{
					"value": "fakeSecretValue",
					"id": "https://unknownbaseURL/secrets/fooSecretName/barSecretVersion",
					"attributes": {
					  "enabled": true,
					  "created": 1493938459,
					  "updated": 1493938459,
					  "recoveryLevel": "Recoverable+Purgeable"
					}
				}`))
				rw.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			secretValue := "fakeSecretValue"
			setParameter := keyvault.SecretSetParameters{
				Value: &secretValue,
			}
			secretBundle, rErr := kvSecretClient.SetSecret(ctx, server.URL, "fooSecretName", setParameter)
			Expect(rErr).To(BeNil())
			Expect(secretBundle).NotTo(BeNil())
			Expect(*secretBundle.Value).To(Equal(secretValue))
		})

		It("Should not return error when trying to get a secret and server returns 404", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.String()).To(Equal("/secrets/fooSecretName/barSecretVersion?api-version=2016-10-01"))
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte(`
				{
					"code": "SecretNotFound",
					"message": "Secret not found: fooSecretName"
				}`))
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			secretBundle, rErr := kvSecretClient.GetSecret(ctx, server.URL, "fooSecretName", "barSecretVersion")
			Expect(rErr).To(BeNil())
			Expect(secretBundle).To(BeNil())
		})

		It("Should successfully purge a deleted secret", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.String()).To(Equal("/deletedsecrets/fooSecretName?api-version=2016-10-01"))
				rw.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			response, rErr := kvSecretClient.PurgeDeletedSecret(ctx, server.URL, "fooSecretName")
			Expect(rErr).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.Response.StatusCode).To(Equal(http.StatusNoContent))
		})

		It("Should successfully get mock deleted secrets", func() {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.String()).To(Equal("/deletedsecrets?api-version=2016-10-01&maxresults=25"))
				rw.Write([]byte(`
				{
					"value": [
						{
							"recoveryId": "https://myvault.vault.azure.net/deletedsecrets/listdeletedsecrettest0",
							"deletedDate": 1493937855,
							"scheduledPurgeDate": 1501713855,
							"contentType": "plainText",
							"id": "https://myvault.vault.azure.net/secrets/listdeletedsecrettest0",
							"attributes": {
							  "enabled": true,
							  "created": 1493937855,
							  "updated": 1493937855,
							  "recoveryLevel": "Recoverable+Purgeable"
							}
						  }
						],
					"nextLink": "https://testvault1021.vault.azure.net:443/deletedsecrets?api-version=7.0&$skiptoken=eyJOZXh0TWFya2VyIjoiMiE4OCFNREF3TURJeUlYTmxZM0psZEM5TVNWTlVVMFZEVWtWVVZFVlRWREVoTURBd01ESTRJVEl3TVRZdE1USXRNVGxVTWpNNk1UQTZORFV1T0RneE9ERXhNRm9oIiwiVGFyZ2V0TG9jYXRpb24iOjB9&maxresults=1"
				  }`))
				rw.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			logger := log.InitializeTestLogger()
			ctx := log.WithLogger(context.Background(), logger)
			apiTracking := log.NewAPITrackingFromParametersMap(nil)
			ctx = log.WithAPITracking(ctx, apiTracking)

			var maxResults int32 = 25
			kvSecretClient := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), "eastus")
			listResult, rErr := kvSecretClient.GetDeletedSecrets(ctx, server.URL, &maxResults)
			Expect(rErr).To(BeNil())
			Expect(listResult).NotTo(BeNil())
			Expect(listResult.Value).NotTo(BeNil())
			Expect(len(*listResult.Value)).To(Equal(1))
			Expect((*listResult.Value)[0].ID).NotTo(BeNil())
			Expect(*((*listResult.Value)[0].ID)).To(Equal("https://myvault.vault.azure.net/secrets/listdeletedsecrettest0"))
		})
	})

})
