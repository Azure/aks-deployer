package keyvaultsecretclient

//go:generate sh -c "mockgen goms.io/aks/rp/core/clients/keyvaultsecretclient Interface >./mock_$GOPACKAGE/interface.go"

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/aks-deployer/pkg/retry"
)

type Interface interface {
	SetSecret(ctx context.Context, vaultBaseURL string, secretName string, parameters keyvault.SecretSetParameters) (*keyvault.SecretBundle, *retry.Error)
	GetSecret(ctx context.Context, vaultBaseURL string, secretName string, secretVersion string) (*keyvault.SecretBundle, *retry.Error)
	GetSecrets(ctx context.Context, vaultBaseURL string, maxResults *int32) (*keyvault.SecretListResult, *retry.Error)
	DeleteSecret(ctx context.Context, vaultBaseURL string, secretName string) (*keyvault.DeletedSecretBundle, *retry.Error)
	PurgeDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (*autorest.Response, *retry.Error)
	GetDeletedSecrets(ctx context.Context, vaultBaseURL string, maxResults *int32) (*keyvault.DeletedSecretListResult, *retry.Error)
}
