package blobclient

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// how to do real test
// 1. create an vmss which has msi enabled
// 2. create an storage account "aksconfigtest", and vmss should have storage blob reader role assignment
// 3. under the storage account, create container "test" which has private access, and upload a file "hello" to container.
// 4. login to the vmss instance, and get msi token by the curl command: curl 'http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fstorage.azure.com%2F' -H Metadata:true
// 5. assign the token to tokenString
// 6. for test msi -> create a user assigned managed identity, assign it to the vmss, and then assign the clientID to variable showed below.
var (
	tokenString  = os.Getenv("CIT_TOKEN")
	clientID     = os.Getenv("CIT_MSI_CLIENTID")
	miResourceID = os.Getenv("CIT_MSI_RESOURCEID")
	url          = "https://aksconfigtest.blob.core.windows.net/test/hello"
)

type tokenProvider struct{}

func (tp *tokenProvider) OAuthToken() string {
	return tokenString
}

func TestBlobClientWithToken(t *testing.T) {
	if tokenString == "" {
		t.Skip("tokenString is not set. The integration test is skipped.")
	}

	blobClient := New(autorest.NewBearerAuthorizer(&tokenProvider{}))
	ctx := context.Background()
	data, err := blobClient.GetBlobData(ctx, "https://aksconfigtest.blob.core.windows.net/test/non-existent")
	require.NoError(t, err)
	assert.Nil(t, data)

	data, err = blobClient.GetBlobData(ctx, url)
	require.NoError(t, err)
	assert.Equal(t, string(data), "hello\n")
}

func TestBlobClientWithMSIClientID(t *testing.T) {
	if clientID == "" {
		t.Skip("clientID is not set. The integration test is skipped.")
	}

	blobClient, err := NewMsiBlobProviderWithClientID(clientID)
	require.NoError(t, err)
	ctx := context.Background()

	data, err := blobClient.GetBlobData(ctx, "https://aksconfigtest.blob.core.windows.net/test/non-existent")
	require.NoError(t, err)
	assert.Nil(t, data)

	data, err = blobClient.GetBlobData(ctx, url)
	require.NoError(t, err)
	assert.Equal(t, string(data), "hello\n")
}

func TestBlobClientWithMSIResourceID(t *testing.T) {
	if miResourceID == "" {
		t.Skip("managed identity resourceID is not set. The integration test is skipped.")
	}

	blobClient, err := NewMsiBlobProviderWithMiResourceID(miResourceID)
	require.NoError(t, err)
	ctx := context.Background()

	data, err := blobClient.GetBlobData(ctx, "https://aksconfigtest.blob.core.windows.net/test/non-existent")
	require.NoError(t, err)
	assert.Nil(t, data)

	data, err = blobClient.GetBlobData(ctx, url)
	require.NoError(t, err)
	assert.Equal(t, string(data), "hello\n")
}

func TestBlobClientListBlobsWithPrefixWithMSIResourceID(t *testing.T) {
	if miResourceID == "" {
		t.Skip("managed identity resourceID is not set. The integration test is skipped.")
	}

	blobClient, err := NewMsiBlobProviderWithMiResourceID(miResourceID)
	require.NoError(t, err)
	ctx := context.Background()

	data, err := blobClient.ListBlobsWithPrefix(ctx, "https://aksconfigtest.blob.core.windows.net/test", "non-existent")
	require.NoError(t, err)
	assert.Nil(t, data)

	data, err = blobClient.ListBlobsWithPrefix(ctx, "https://aksconfigtest.blob.core.windows.net/test", "hello")
	require.NoError(t, err)
	assert.Equal(t, string(data.Blobs[0].Name), "hello")
}
