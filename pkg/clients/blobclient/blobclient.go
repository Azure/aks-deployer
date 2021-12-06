package blobclient

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	uuid "github.com/satori/go.uuid"

	"github.com/Azure/aks-deployer/pkg/apierror"
	cgerror "github.com/Azure/aks-deployer/pkg/categorizederror"
	"github.com/Azure/aks-deployer/pkg/clients/decorators"
)

var _ Interface = &BlobClient{}

type BlobClient struct {
	restClient            autorest.Client
	storageEndpointSuffix string
}

const (
	resource = "https://storage.azure.com/"
)

var statusCodesForRetry = []int{}

func init() {
	// dont retry on http.StatusTooManyRequests
	for _, code := range autorest.StatusCodesForRetry {
		if code != http.StatusTooManyRequests {
			statusCodesForRetry = append(statusCodesForRetry, code)
		}
	}
}

func New(authorizer autorest.Authorizer) *BlobClient {
	restClient := autorest.NewClientWithUserAgent("AKS-Blob-Client")
	restClient.RetryAttempts = 3
	restClient.RetryDuration = time.Second * 1
	restClient.Authorizer = authorizer

	return &BlobClient{
		restClient: restClient,
	}
}

func NewMsiBlobProviderWithClientID(clientID string) (*BlobClient, error) {
	if clientID == "" {
		return nil, fmt.Errorf("clientID is not set")
	}

	msiEndpoint, err := adal.GetMSIEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get msi endpoint")
	}

	spt, err := adal.NewServicePrincipalTokenFromMSIWithUserAssignedID(msiEndpoint, resource, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get msi token with clientID %q", clientID)
	}

	return New(autorest.NewBearerAuthorizer(spt)), nil
}

func NewMsiBlobProviderWithMiResourceID(miResID string) (*BlobClient, error) {
	if miResID == "" {
		return nil, fmt.Errorf("managed identity resourceID is not set")
	}

	msiEndpoint, err := adal.GetMSIEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get msi endpoint")
	}

	spt, err := adal.NewServicePrincipalTokenFromMSIWithIdentityResourceID(msiEndpoint, resource, miResID)
	if err != nil {
		return nil, fmt.Errorf("failed to get msi token with managed identity resourceID %q", miResID)
	}

	return New(autorest.NewBearerAuthorizer(spt)), nil
}

func NewMsiBlobProviderWithMiResourceIDWithCloudName(miResID string, cloudName string) (*BlobClient, error) {
	blobClient, err := NewMsiBlobProviderWithMiResourceID(miResID)
	if err != nil {
		return nil, err
	}
	cloudEnvironment, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		return nil, err
	}
	blobClient.storageEndpointSuffix = cloudEnvironment.StorageEndpointSuffix
	return blobClient, nil
}

// GetBlobData gets data from blob with URL
func (c *BlobClient) GetBlobData(ctx context.Context, url string) ([]byte, error) {
	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{autorest.AsGet()})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request, err: %s", err)
	}

	resp, err := c.sender(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to get response, err: %s", err)
	}

	bytes, err := c.getBlobDataResponder(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to handle response, err: %s", err)
	}
	return bytes, nil
}

// PutBlobData put data to blob with URL
func (c *BlobClient) PutBlobData(ctx context.Context, url string, data []byte) error {
	r, err := c.putPreparer(ctx, url, data, nil)
	if err != nil {
		return fmt.Errorf("failed to prepare request, err: %s", err)
	}

	resp, err := c.sender(ctx, r)

	if resp != nil {
		if resp.StatusCode == http.StatusCreated {
			return nil
		}
		return fmt.Errorf("Unexpected status code happened: %d when put blob: %s", resp.StatusCode, url)
	}
	return fmt.Errorf("failed to get response, err: %s", err)
}

// GetBlobData gets data from blob with URL
func (c *BlobClient) GetBlobDataV1(ctx context.Context, storageAccountName, containerName, blobName string) ([]byte, int, error) {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s", storageAccountName, c.storageEndpointSuffix, containerName, blobName)
	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{autorest.AsGet()})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to prepare request, err: %s", err)
	}

	resp, err := c.sender(ctx, r)
	if err != nil {
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		return nil, statusCode, fmt.Errorf("failed to get response, err: %s", err)
	}

	bytes, err := c.getBlobDataResponder(resp)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to handle response, err: %s", err)
	}
	return bytes, resp.StatusCode, nil
}

// BlobExists checks if a blob exist.
// ref: https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-properties
func (c *BlobClient) BlobExists(ctx context.Context, storageAccountName, containerName, blobName string) (bool, error) {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s", storageAccountName, c.storageEndpointSuffix, containerName, blobName)
	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{autorest.AsHead()})
	if err != nil {
		return false, cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to prepare request, err: %s", err))
	}

	resp, err := c.sender(ctx, r)
	if resp != nil {
		switch resp.StatusCode {
		case http.StatusOK:
			return true, nil
		case http.StatusNotFound:
			return false, nil
		default:
			err := fmt.Errorf("Unexpected status code happened: %d when check blob existence", resp.StatusCode)
			return false, cgerror.NewCategorizedError(
				ctx,
				apierror.InternalError,
				cgerror.UnexpectedStatusCodeCheckBlobExist,
				cgerror.BlobStorage,
				err)
		}
	}
	return false, cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

// PutBlobDataV1 gets data from blob with URL
func (c *BlobClient) PutBlobDataV1(ctx context.Context, storageAccountName, containerName, blobName string, data []byte, extraHeaders map[string]interface{}) error {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s", storageAccountName, c.storageEndpointSuffix, containerName, blobName)
	r, err := c.putPreparer(ctx, url, data, extraHeaders)
	if err != nil {
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to get response, err: %s", err))
	}

	resp, err := c.sender(ctx, r)

	if resp != nil {
		if resp.StatusCode == http.StatusCreated {
			return nil
		}
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.UnexpectedStatusCodePutBlob,
			cgerror.BlobStorage,
			fmt.Errorf("Unexpected status code happened: %d when put blob: %s", resp.StatusCode, url))
	}
	return cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

func (c *BlobClient) preparer(ctx context.Context, url string, decorators []autorest.PrepareDecorator) (*http.Request, error) {
	urlParam := map[string]interface{}{
		"url": url,
	}

	headers := map[string]interface{}{
		"x-ms-version":           "2017-11-09",
		"x-ms-client-request-id": uuid.NewV4().String(),
	}
	decorators = append(
		decorators,
		autorest.WithCustomBaseURL("{url}", urlParam),
		autorest.WithHeaders(headers))

	preparer := autorest.CreatePreparer(
		decorators...,
	)

	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (c *BlobClient) putPreparer(ctx context.Context, url string, data []byte, extraHeaders map[string]interface{}) (*http.Request, error) {
	urlParam := map[string]interface{}{
		"url": url,
	}

	headers := map[string]interface{}{
		"x-ms-version":           "2017-11-09",
		"x-ms-client-request-id": uuid.NewV4().String(),
		"x-ms-blob-type":         "BlockBlob",
	}

	if extraHeaders != nil {
		for k, v := range extraHeaders {
			headers[k] = v
		}
	}

	putPreparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{url}", urlParam),
		autorest.WithHeaders(headers),
		autorest.WithBytes(&data),
	)

	return putPreparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (c *BlobClient) sender(ctx context.Context, request *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(
		c.restClient,
		request,
		autorest.DoCloseIfError(),
		decorators.DoLogging(""),
		autorest.DoRetryForStatusCodes(c.restClient.RetryAttempts, c.restClient.RetryDuration, statusCodesForRetry...),
	)
}

func (c *BlobClient) getBlobDataResponder(resp *http.Response) ([]byte, error) {
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	err := autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body, err %s", err)
	}

	return result, nil
}

// ListBlobsWithPrefix lists blobs with URL and prefix
func (c *BlobClient) ListBlobsWithPrefix(ctx context.Context, url, prefix string) (storage.BlobListResponse, error) {
	r, err := c.preparer(ctx, fmt.Sprintf("%s?restype=container&comp=list&prefix=%s",
		url, prefix), []autorest.PrepareDecorator{autorest.AsGet()})
	if err != nil {
		return storage.BlobListResponse{}, fmt.Errorf("failed to prepare request, err: %s", err)
	}

	resp, err := c.sender(ctx, r)
	if err != nil {
		return storage.BlobListResponse{}, fmt.Errorf("failed to get response, err: %s", err)
	}

	blobs, err := c.listBlobsWithPrefixResponder(resp)
	if err != nil {
		return storage.BlobListResponse{}, fmt.Errorf("failed to handle response, err: %s", err)
	}
	return blobs, nil
}

func (c *BlobClient) listBlobsWithPrefixResponder(resp *http.Response) (storage.BlobListResponse, error) {
	if resp.StatusCode == http.StatusNotFound {
		return storage.BlobListResponse{}, nil
	}

	err := autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK))
	if err != nil {
		return storage.BlobListResponse{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return storage.BlobListResponse{}, fmt.Errorf("failed to read response body, err %s", err)
	}

	var out storage.BlobListResponse
	err = xml.Unmarshal(data, &out)
	if err != nil {
		return storage.BlobListResponse{}, fmt.Errorf("failed to unmarshal response data, err %s", err)
	}

	return out, nil
}

func (c *BlobClient) ContainerExists(ctx context.Context, storageAccountName, containerName string) (bool, error) {
	url := fmt.Sprintf("https://%s.blob.%s/%s?restype=container&comp=metadata", storageAccountName, c.storageEndpointSuffix, containerName)
	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{autorest.AsGet()})
	if err != nil {
		return false, cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to prepare request, err: %s", err))
	}
	resp, err := c.sender(ctx, r)
	if resp != nil {
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound {
			return resp.StatusCode == http.StatusOK, nil
		}
		return false, cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.UnexpectedStatusCodeCheckContainerExist,
			cgerror.BlobStorage,
			fmt.Errorf("Unexpected status code happened: %d when check container exist: %s", resp.StatusCode, url))
	}
	return false, cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

func (c *BlobClient) DeleteContainer(ctx context.Context, storageAccountName, containerName string) error {
	url := fmt.Sprintf("https://%s.blob.%s/%s?restype=container", storageAccountName, c.storageEndpointSuffix, containerName)
	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{autorest.AsDelete()})
	if err != nil {
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to prepare request, err: %s", err))
	}
	resp, err := c.sender(ctx, r)

	if resp != nil {
		if resp.StatusCode == http.StatusAccepted {
			return nil
		}
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.UnexpectedStatusCodeDeleteContainer,
			cgerror.BlobStorage,
			fmt.Errorf("Unexpected status code happened: %d when delete container: %s", resp.StatusCode, url))
	}
	return cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

const headerNameBlobPublicAccess = "x-ms-blob-public-access"

func (c *BlobClient) CreateContainerWithACL(ctx context.Context, storageAccountName, containerName string, acl BlobPublicAccessLevel) error {
	url := fmt.Sprintf("https://%s.blob.%s/%s?restype=container", storageAccountName, c.storageEndpointSuffix, containerName)

	preparers := []autorest.PrepareDecorator{autorest.AsPut()}
	switch acl {
	case BlobPublicAccessLevelNotSet:
	// don't set header, use API default
	default:
		// specify as x-ms-blob-public-access header
		preparers = append(preparers, autorest.WithHeader(headerNameBlobPublicAccess, string(acl)))
	}
	r, err := c.preparer(ctx, url, preparers)
	if err != nil {
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to prepare request, err: %s", err))
	}
	resp, err := c.sender(ctx, r)
	if resp != nil {
		if resp.StatusCode == http.StatusAccepted ||
			resp.StatusCode == http.StatusCreated {
			return nil
		}
		return cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.UnexpectedStatusCodeCreateContainer,
			cgerror.BlobStorage,
			fmt.Errorf("Unexpected status code happened: %d when delete container: %s", resp.StatusCode, url))
	}
	return cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

func (c *BlobClient) CreateContainer(ctx context.Context, storageAccountName, containerName string) error {
	return c.CreateContainerWithACL(ctx, storageAccountName, containerName, BlobPublicAccessLevelNotSet)
}

func (c *BlobClient) AcquireLease(ctx context.Context,
	storageAccountName,
	containerName,
	blobName string,
	duration uint) (string, error) {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s?comp=lease", storageAccountName, c.storageEndpointSuffix, containerName, blobName)

	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{
		autorest.AsPut(),
		autorest.WithHeader("x-ms-lease-action", "acquire"),
		autorest.WithHeader("x-ms-lease-duration", fmt.Sprintf("%d", duration)),
	})
	if err != nil {
		return "", cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.FailedPrepareRequest,
			cgerror.BlobStorage,
			fmt.Errorf("failed to prepare request, err: %s", err))
	}
	resp, err := c.sender(ctx, r)
	if resp != nil {
		if resp.StatusCode == http.StatusAccepted ||
			resp.StatusCode == http.StatusCreated {
			leaseIDHeader := resp.Header.Get("x-ms-lease-id")
			return leaseIDHeader, nil
		}
		return "", cgerror.NewCategorizedError(
			ctx,
			apierror.InternalError,
			cgerror.UnexpectedStatusCodeAcquireLease,
			cgerror.BlobStorage,
			fmt.Errorf("Unexpected status code happened: %d when acquire lease for: %s", resp.StatusCode, url))
	}
	return "", cgerror.NewCategorizedError(
		ctx,
		apierror.InternalError,
		cgerror.FailedGetResponse,
		cgerror.BlobStorage,
		fmt.Errorf("failed to get response, err: %s", err))
}

func (c *BlobClient) ReleaseLease(ctx context.Context, storageAccountName, containerName, blobName, leaseID string) error {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s?comp=lease", storageAccountName, c.storageEndpointSuffix, containerName, blobName)

	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{
		autorest.AsPut(),
		autorest.WithHeader("x-ms-lease-action", "release"),
		autorest.WithHeader(HeaderLeaseID, leaseID),
	})
	if err != nil {
		return fmt.Errorf("failed to prepare request, err: %s", err)
	}
	resp, err := c.sender(ctx, r)
	if resp != nil {
		if resp.StatusCode == http.StatusOK {
			return nil
		}
		return fmt.Errorf("Unexpected status code happened: %d when release lease for: %s", resp.StatusCode, url)
	}
	return fmt.Errorf("failed to get response, err: %s", err)
}

func (c *BlobClient) BreakLease(ctx context.Context, storageAccountName, containerName, blobName, leaseID string) error {
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s?comp=lease", storageAccountName, c.storageEndpointSuffix, containerName, blobName)

	r, err := c.preparer(ctx, url, []autorest.PrepareDecorator{
		autorest.AsPut(),
		autorest.WithHeader("x-ms-lease-action", "break"),
		autorest.WithHeader(HeaderLeaseID, leaseID),
	})
	if err != nil {
		return fmt.Errorf("failed to prepare request, err: %s", err)
	}
	resp, err := c.sender(ctx, r)
	if resp != nil {
		if resp.StatusCode == http.StatusAccepted ||
			resp.StatusCode == http.StatusCreated {
			return nil
		}
		return fmt.Errorf("Unexpected status code happened: %d when break lease for: %s", resp.StatusCode, url)
	}
	return fmt.Errorf("failed to get response, err: %s", err)
}
