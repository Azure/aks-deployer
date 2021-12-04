// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package secret

//go:generate sh -c "mockgen -package mock_$GOPACKAGE -destination ./mock_$GOPACKAGE/$GOFILE github.com/aks-deployer/pkg/secret KeyvaultSecretProvider"

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/aks-deployer/pkg/auth/tokenprovider"
	"github.com/Azure/aks-deployer/pkg/clients/keyvaultsecretclient"
	"github.com/Azure/aks-deployer/pkg/datastructs"
	"github.com/Azure/aks-deployer/pkg/log"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

var (
	// ErrInvalidSecretURI indicates the secretURI does not match keyvault secret pattern
	ErrInvalidSecretURI = errors.New("invalid secret URI")
	// ErrInvalidCertificateURI indicates the certificateURI does not match keyvault certificate pattern
	ErrInvalidCertificateURI = errors.New("invalid certificate URI")
	// ErrEmptyResource  indicates resourceID is empty
	ErrEmptyResource = errors.New("resource should not be empty")
	// ErrEmptyValue indicates the value to set is empty
	ErrEmptyValue = errors.New("value to set should not be empty")
)

// pattern for https://myvault.vault.net/secerts/mysecret/version
// or https://myvault.vault.net/secerts/mysecret
const getPattern = `^(https://[^/]+)/secrets/([^/]+)(/[^/]+)?$`

const setPattern = `^(https://[^/]+)/secrets/([^/]+)$`

var (
	re                            = regexp.MustCompile(getPattern)
	setReg                        = regexp.MustCompile(setPattern)
	_      KeyvaultSecretProvider = (*keyvaultSecretProvider)(nil)
)

// KeyvaultSecretProvider is an interface to retrieve a Keyvault secret
type KeyvaultSecretProvider interface {
	Get(logger *log.Logger, secretURI string) (keyvault.SecretBundle, error)
	// Set set the value into the keyvault
	Set(logger *log.Logger, secretURI string, secretSetParameter keyvault.SecretSetParameters) (keyvault.SecretBundle, error)
}

// keyvaultSecretProvider is an implementation of KeyvaultSecretProvider
type keyvaultSecretProvider struct {
	client keyvaultsecretclient.Interface
}

// Get returns the SecretBundle for a given keyvault secret url.
// url must be a valid keyvault secret url such as "https://foo.vault.net/secerts/bar/version"
// or https://foo.vault.net/secerts/bar
func (p *keyvaultSecretProvider) Get(logger *log.Logger, secretURI string) (keyvault.SecretBundle, error) {
	parts := re.FindStringSubmatch(secretURI)
	if len(parts) != 4 {
		return keyvault.SecretBundle{}, ErrInvalidSecretURI
	}
	vaultBaseURL := parts[1]
	secretname := parts[2]
	secretVersion := strings.TrimPrefix(parts[3], "/")
	ctx := log.WithLogger(context.Background(), logger)
	secretBundle, retriableError := p.client.GetSecret(ctx, vaultBaseURL, secretname, secretVersion)
	if retriableError != nil {
		logger.Warningf(ctx, "keyvault get secret failure URI: %s", secretURI)
		return keyvault.SecretBundle{}, retriableError.Error
	}

	if secretBundle == nil {
		logger.Warningf(ctx, "keyvault doesn't get secret at URI: %s", secretURI)
		return keyvault.SecretBundle{}, nil
	}

	logger.Infof(ctx, "retrieved secret from %s", *secretBundle.ID)
	return *secretBundle, nil
}

// Get returns the SecretBundle for a given keyvault secret url.
// url must be a valid keyvault secret url such as "https://foo.vault.net/secerts/bar/version"
// or https://foo.vault.net/secerts/bar
func (p *keyvaultSecretProvider) Set(logger *log.Logger, secretURI string, secretSetParameter keyvault.SecretSetParameters) (keyvault.SecretBundle, error) {
	parts := setReg.FindStringSubmatch(secretURI)
	if len(parts) != 3 {
		return keyvault.SecretBundle{}, ErrInvalidSecretURI
	}
	vaultBaseURL := parts[1]
	secretname := parts[2]

	if secretSetParameter.Value == nil ||
		*secretSetParameter.Value == "" {
		return keyvault.SecretBundle{}, ErrEmptyValue
	}
	ctx := log.WithLogger(context.Background(), logger)
	secretBundle, retriableError := p.client.SetSecret(ctx, vaultBaseURL, secretname, secretSetParameter)
	if retriableError != nil {
		logger.Warningf(ctx, "keyvault get secret failure URI: %s", secretURI)
		return keyvault.SecretBundle{}, retriableError.Error
	}

	if secretBundle == nil {
		logger.Warningf(ctx, "keyvault doesn't get secret at URI: %s", secretURI)
		return keyvault.SecretBundle{}, nil
	}

	logger.Infof(ctx, "retrieved secret from %s", *secretBundle.ID)
	return *secretBundle, nil
}

// NewKeyvaultSecretProviderFromTokenProvider returns an instance of keyvaultSecretProvider using a token provider
func NewKeyvaultSecretProviderFromTokenProvider(resource string, provider tokenprovider.TokenProvider) (KeyvaultSecretProvider, error) {
	spt, err := provider.GetToken(resource)
	if err != nil {
		return nil, err
	}

	client := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(spt), "")
	return &keyvaultSecretProvider{
		client: client,
	}, nil
}

// NewMsiKeyvaultSecretProvider returns an instance of keyvaultSecretProvider using MSI auth
func NewMsiKeyvaultSecretProvider(resource string) (KeyvaultSecretProvider, error) {
	if resource == "" {
		return nil, ErrEmptyResource
	}

	msiEndpoint, err := adal.GetMSIVMEndpoint()
	if err != nil {
		return nil, err
	}

	spt, err := adal.NewServicePrincipalTokenFromMSI(msiEndpoint, resource)
	if err != nil {
		return nil, err
	}

	client := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(spt), "")
	return &keyvaultSecretProvider{
		client: client,
	}, nil
}

// NewClientSecretKeyvaultSecretProvider returns an instance of keyvaultSecretProvider
// using client secret auth
func NewClientSecretKeyvaultSecretProvider(
	aadEndpoint, tenantID, clientID string,
	clientSecret datastructs.RedactedString,
	resource string) (KeyvaultSecretProvider, error) {
	if resource == "" {
		return nil, ErrEmptyResource
	}

	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("clientID and clientSecret cannot be empty")
	}

	c, err := adal.NewOAuthConfig(aadEndpoint, tenantID)
	if err != nil {
		return nil, err
	}

	spt, err := adal.NewServicePrincipalToken(*c, clientID, string(clientSecret), resource)
	if err != nil {
		return nil, err
	}

	client := keyvaultsecretclient.New(autorest.NewBearerAuthorizer(spt), "")
	return &keyvaultSecretProvider{
		client: client,
	}, nil
}
