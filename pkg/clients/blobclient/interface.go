package blobclient

//go:generate sh -c "mockgen goms.io/aks/rp/core/clients/blobclient Interface >./mock_$GOPACKAGE/interface.go"

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// BlobPublicAccessLevel specifies the storage account blob public access level.
// ref: https://docs.microsoft.com/en-us/rest/api/storageservices/create-container
type BlobPublicAccessLevel string

const (
	// BlobPublicAccessLevelNotSet - default ACL, private to the account owner
	BlobPublicAccessLevelNotSet BlobPublicAccessLevel = ""
	// BlobPublicAccessLevelContainer - specifies full public read access for container and blob data.
	BlobPublicAccessLevelContainer BlobPublicAccessLevel = "container"
	// BlobPublicAccessLevelBlob - specifies public read access for blobs.
	BlobPublicAccessLevelBlob BlobPublicAccessLevel = "blob"
)

// HeaderLeaseID - specify the header name for operation lease id.
const HeaderLeaseID = "x-ms-lease-id"

// Interface blob client APIs
type Interface interface {
	GetBlobData(ctx context.Context, url string) ([]byte, error)
	GetBlobDataV1(ctx context.Context, storageAccountName, containerName, blobName string) ([]byte, int, error)
	BlobExists(ctx context.Context, storageAccountName, containerName, blobName string) (bool, error)
	PutBlobData(ctx context.Context, url string, data []byte) error
	PutBlobDataV1(ctx context.Context, storageAccountName, containerName, blobName string, data []byte, extraHeaders map[string]interface{}) error
	ListBlobsWithPrefix(ctx context.Context, url, prefix string) (storage.BlobListResponse, error)
	ContainerExists(ctx context.Context, storageAccountName, containerName string) (bool, error)
	DeleteContainer(ctx context.Context, storageAccountName, containerName string) error
	CreateContainer(ctx context.Context, storageAccountName, containerName string) error
	CreateContainerWithACL(ctx context.Context, storageAccountName, containerName string, acl BlobPublicAccessLevel) error
	AcquireLease(ctx context.Context, storageAccountName, containerName, blobName string, durationInSeconds uint) (string, error)
	ReleaseLease(ctx context.Context, storageAccountName, containerName, blobName, leaseID string) error
	BreakLease(ctx context.Context, storageAccountName, containerName, blobName, leaseID string) error
}
