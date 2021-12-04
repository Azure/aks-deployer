// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package keyvaultsecretclient

import (
	"context"
	"net/http"
	"time"

	cgerror "github.com/Azure/aks-deployer/pkg/categorizederror"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	uuid "github.com/satori/go.uuid"

	"github.com/Azure/aks-deployer/pkg/clients/decorators"
	"github.com/Azure/aks-deployer/pkg/clients/validation"
	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/Azure/aks-deployer/pkg/retry"
)

const (
	APIVersion     = "2016-10-01"
	headerClientID = "x-ms-client-request-id"
)

var _ Interface = &client{}

type client struct {
	// Here we don't use armclient like other clients because the
	// key vault base url is not fixed.
	restClient   autorest.Client
	clientRegion string
}

// New creates key vault secret client
func New(authorizer autorest.Authorizer, clientRegion string) Interface {
	restClient := autorest.NewClientWithUserAgent("KeyVault-Client")
	restClient.RetryAttempts = 3
	restClient.RetryDuration = time.Second * 1
	restClient.Authorizer = authorizer

	return &client{
		restClient:   restClient,
		clientRegion: clientRegion,
	}
}

func (c *client) SetSecret(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
	parameters keyvault.SecretSetParameters,
) (*keyvault.SecretBundle, *retry.Error) {
	if err := validation.ValidateSecretName(secretName); err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedValidateSecretName, cgerror.KeyVaultSecret, err).SetRetriable(false)
		traceKeyVaultSecretError(
			ctx,
			"error executing SetSecret when validating secret name, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(false, err)
	}

	if err := validation.ValidateSecretSetParameters(parameters); err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedValidateSecretSetParameters, cgerror.KeyVaultSecret, err).SetRetriable(false)
		traceKeyVaultSecretError(
			ctx,
			"error executing SetSecret when validating SetParameter, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(false, err)
	}

	request, err := c.setSecretPreparer(ctx, vaultBaseURL, secretName, parameters)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing SetSecret when constructing request, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing SetSecret when sending request, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.setSecretResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing SetSecret when extracting response, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			retriableErr.Error)
		return result, retriableErr
	}

	return result, nil
}

func (c *client) GetSecret(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
	secretVersion string,
) (*keyvault.SecretBundle, *retry.Error) {
	request, err := c.getSecretPreparer(ctx, vaultBaseURL, secretName, secretVersion)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecret when constructing request, vaultBaseURL %s, secretName %s, secretVersion %s, error %v",
			vaultBaseURL,
			secretName,
			secretVersion,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecret when sending request, vaultBaseURL %s, secretName %s, secretVersion %s, error %v",
			vaultBaseURL,
			secretName,
			secretVersion,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.getSecretResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecret when extracting response, vaultBaseURL %s, secretName %s, secretVersion %s, error %v",
			vaultBaseURL,
			secretName,
			secretVersion,
			retriableErr.Error)
		return result, retriableErr
	}

	log.GetLogger(ctx).Infof(ctx, "GetSecret succeeded with status code: %v", response.StatusCode)
	return result, nil
}

func (c *client) GetSecrets(
	ctx context.Context,
	vaultBaseURL string,
	maxResults *int32,
) (*keyvault.SecretListResult, *retry.Error) {
	if err := validation.ValidateSecretMaxResults(maxResults); err != nil {
		var maxResultsToLog interface{}
		if maxResults == nil {
			maxResultsToLog = "nil"
		} else {
			maxResultsToLog = *maxResults
		}
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedValidateSecretMaxResults, cgerror.KeyVaultSecret, err).SetRetriable(false)
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecrets when validating maxResults, vaultBaseURL %s, maxResults %v, error %v",
			vaultBaseURL,
			maxResultsToLog,
			err)
		return nil, retry.NewError(false, err)
	}

	request, err := c.getSecretsPreparer(ctx, vaultBaseURL, maxResults)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecrets when constructing request, vaultBaseURL %s, maxResults %v, error %v",
			vaultBaseURL,
			*maxResults,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecrets when sending request, vaultBaseURL %s, maxResults %v, error %v",
			vaultBaseURL,
			*maxResults,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.getSecretsResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetSecrets when extracting response, vaultBaseURL %s, maxResults %v, error %v",
			vaultBaseURL,
			maxResults,
			retriableErr.Error)
		return result, retriableErr
	}

	return result, nil
}

func (c *client) DeleteSecret(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
) (*keyvault.DeletedSecretBundle, *retry.Error) {
	request, err := c.deleteSecretPreparer(ctx, vaultBaseURL, secretName)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing DeleteSecret when constructing request, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing DeleteSecret when sending request, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.deleteSecretResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing DeleteSecret when extracting response, vaultBaseURL %s, secretName %s, error %v",
			vaultBaseURL,
			secretName,
			retriableErr.Error)
		return result, retriableErr
	}

	log.GetLogger(ctx).Infof(ctx, "Delete succeeded with status code: %v", result.Response.StatusCode)
	return result, nil
}

func (c *client) PurgeDeletedSecret(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
) (*autorest.Response, *retry.Error) {
	request, err := c.purgeDeletedSecretPreparer(ctx, vaultBaseURL, secretName)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing PurgeDeletedSecret when constructing request, vaultBaseURL %s, secretName %s, error %s",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing PurgeDeletedSecret when sending request, vaultBaseURL %s, secretName %s, error %s",
			vaultBaseURL,
			secretName,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.purgeDeletedSecretResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing PurgeDeletedSecret when extracting response, vaultBaseURL %s, secretName %s, error %s",
			vaultBaseURL,
			secretName,
			retriableErr.Error)
		return result, retriableErr
	}

	log.GetLogger(ctx).Infof(ctx, "PurgeDeletedSecret succeeded with status code: %d", result.Response.StatusCode)
	return result, nil
}

func (c *client) GetDeletedSecrets(
	ctx context.Context,
	vaultBaseURL string,
	maxResults *int32,
) (*keyvault.DeletedSecretListResult, *retry.Error) {
	var defaultMaxResults int32 = 25
	if maxResults == nil {
		maxResults = &defaultMaxResults
	}
	if err := validation.ValidateSecretMaxResults(maxResults); err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedValidateSecretMaxResults, cgerror.KeyVaultSecret, err).SetRetriable(false)
		traceKeyVaultSecretError(
			ctx,
			"error executing GetDeletedSecrets when validating maxResults, vaultBaseURL %s, maxResults %d, error %s",
			vaultBaseURL,
			*maxResults,
			err)
		return nil, retry.NewError(false, err)
	}

	request, err := c.getDeletedSecretsPreparer(ctx, vaultBaseURL, maxResults)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetDeletedSecrets when constructing request, vaultBaseURL %s, maxResults %d, error %s",
			vaultBaseURL,
			*maxResults,
			err)
		return nil, retry.NewError(false, err)
	}

	response, err := c.sender(ctx, request)
	if err != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetDeletedSecrets when sending request, vaultBaseURL %s, maxResults %d, error %s",
			vaultBaseURL,
			*maxResults,
			err)
		return nil, retry.NewError(true, err)
	}

	result, retriableErr := c.getDeletedSecretsResponder(ctx, response)
	if retriableErr != nil {
		traceKeyVaultSecretError(
			ctx,
			"error executing GetDeletedSecrets when extracting response, vaultBaseURL %s, maxResults %d, error %s",
			vaultBaseURL,
			maxResults,
			retriableErr.Error)
		return result, retriableErr
	}

	return result, nil
}

func (c *client) setSecretPreparer(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
	parameters keyvault.SecretSetParameters,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"secret-name": autorest.Encode("path", secretName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/secrets/{secret-name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)
	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) sender(ctx context.Context, request *http.Request) (*http.Response, error) {
	response, err := autorest.SendWithSender(
		c.restClient,
		request,
		autorest.DoCloseIfError(),
		autorest.DoRetryForStatusCodes(c.restClient.RetryAttempts, c.restClient.RetryDuration, autorest.StatusCodesForRetry...),
		decorators.DoLogging(c.clientRegion),
	)
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, response, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return response, err
}

func (c *client) setSecretResponder(ctx context.Context, resp *http.Response) (*keyvault.SecretBundle, *retry.Error) {
	var result keyvault.SecretBundle
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func (c *client) getSecretPreparer(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
	secretVersion string,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"secret-name":    autorest.Encode("path", secretName),
		"secret-version": autorest.Encode("path", secretVersion),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/secrets/{secret-name}/{secret-version}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)

	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) getSecretResponder(ctx context.Context, resp *http.Response) (*keyvault.SecretBundle, *retry.Error) {
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	var result keyvault.SecretBundle
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func (c *client) getSecretsPreparer(
	ctx context.Context,
	vaultBaseURL string,
	maxResults *int32,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/secrets"),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)

	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) getSecretsResponder(ctx context.Context, resp *http.Response) (*keyvault.SecretListResult, *retry.Error) {
	var result keyvault.SecretListResult
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func (c *client) deleteSecretPreparer(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"secret-name": autorest.Encode("path", secretName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/secrets/{secret-name}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)
	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) deleteSecretResponder(ctx context.Context, resp *http.Response) (*keyvault.DeletedSecretBundle, *retry.Error) {
	var result keyvault.DeletedSecretBundle
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func (c *client) purgeDeletedSecretPreparer(
	ctx context.Context,
	vaultBaseURL string,
	secretName string,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"secret-name": autorest.Encode("path", secretName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/deletedsecrets/{secret-name}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)
	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) purgeDeletedSecretResponder(ctx context.Context, resp *http.Response) (*autorest.Response, *retry.Error) {
	var result autorest.Response
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func (c *client) getDeletedSecretsPreparer(
	ctx context.Context,
	vaultBaseURL string,
	maxResults *int32,
) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/deletedsecrets"),
		autorest.WithQueryParameters(queryParameters),
		withClientID(),
	)
	request, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = cgerror.NewCategorizedError(ctx, apierror.InternalError, cgerror.FailedPrepareRequest, cgerror.KeyVaultSecret, err).SetRetriable(false)
	}
	return request, err
}

func (c *client) getDeletedSecretsResponder(ctx context.Context, resp *http.Response) (*keyvault.DeletedSecretListResult, *retry.Error) {
	var result keyvault.DeletedSecretListResult
	err := autorest.Respond(
		resp,
		c.restClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	if err != nil {
		err = cgerror.HandleErrorToCategorizedError(ctx, resp, err).SetDependency(cgerror.KeyVaultSecret)
	}
	return &result, retry.GetError(resp, err)
}

func traceKeyVaultSecretError(ctx context.Context, fmt string, args ...interface{}) {
	log.GetLogger(ctx).Errorf(ctx, fmt, args...)
}

func withClientID() autorest.PrepareDecorator {
	uuidString := uuid.NewV4().String()
	return autorest.WithHeader(headerClientID, uuidString)
}
