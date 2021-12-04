package keyvaultsecretclient_test

import (
	"context"
	"os"

	"github.com/Azure/go-autorest/autorest"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/Azure/aks-deployer/pkg/clients/keyvaultsecretclient"
	"github.com/Azure/aks-deployer/pkg/log"

	"testing"
)

func TestKeyVaultSecretClient(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "KeyVaultSecretClient Test Suite", []Reporter{junitReporter})
}

var (
	keyVaultSecretClient keyvaultsecretclient.Interface
	logger               *log.Logger
	ctx                  context.Context
	location                   = "centralus"
	keyVaultBaseURL            = "https://akscitkeyvault.vault.azure.net/"
	secretName                 = uuid.NewV4().String()
	secretValue                = "secretValue"
	secretTagName              = uuid.NewV4().String()
	secretTagValue             = "cit-tag-value"
	maxResults           int32 = 25
)

type tokenProvider struct{}

// The token to access key vault is different from the token to access ARM
// So here we use a new envrionment variable
var tokenString = os.Getenv("CIT_KV_TOKEN")

func (tp *tokenProvider) OAuthToken() string {
	return tokenString
}

var _ = BeforeSuite(func() {
	logger = log.InitializeTestLogger()
	ctx = log.WithLogger(context.Background(), logger)
	apiTracking := log.NewAPITrackingFromParametersMap(nil)
	ctx = log.WithAPITracking(ctx, apiTracking)
	keyVaultSecretClient = keyvaultsecretclient.New(autorest.NewBearerAuthorizer(&tokenProvider{}), location)
})

var _ = AfterSuite(func() {
})
