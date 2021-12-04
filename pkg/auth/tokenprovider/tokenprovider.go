// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package tokenprovider

import (
	"context"
	"errors"

	"github.com/Azure/aks-deployer/pkg/datastructs"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ProviderOptions struct {
	TenantID     string
	ClientID     string
	ClientSecret datastructs.RedactedString
	CloudName    string
}

type TokenProvider interface {
	GetToken(resource string) (*adal.ServicePrincipalToken, error)
}

type MsiTokenProvider struct {
	resourceID string
}

func (m *MsiTokenProvider) GetToken(resource string) (*adal.ServicePrincipalToken, error) {
	if resource == "" {
		return nil, errors.New("resource must be provided")
	}

	msiEndpoint, err := adal.GetMSIVMEndpoint()
	if err != nil {
		return nil, err
	}

	spt, err := adal.NewServicePrincipalTokenFromMSIWithIdentityResourceID(msiEndpoint, resource, m.resourceID)
	if err != nil {
		return nil, err
	}

	return spt, nil
}

func NewMsiTokenProvider(resourceID string) (*MsiTokenProvider, error) {
	if resourceID == "" {
		return nil, errors.New("resourceID must be provided")
	}
	return &MsiTokenProvider{resourceID: resourceID}, nil
}

type ServicePrincipalTokenProvider struct {
	tenantID     string
	clientID     string
	aadEndpoint  string
	clientSecret datastructs.RedactedString
}

func (s *ServicePrincipalTokenProvider) GetToken(resource string) (*adal.ServicePrincipalToken, error) {
	if resource == "" {
		return nil, errors.New("resource must be provided")
	}

	c, err := adal.NewOAuthConfig(s.aadEndpoint, s.tenantID)
	if err != nil {
		return nil, err
	}

	spt, err := adal.NewServicePrincipalToken(*c, s.clientID, string(s.clientSecret), resource)
	if err != nil {
		return nil, err
	}

	return spt, err
}

func NewServicePrincipalTokenProvider(tenantID string, clientID string,
	clientSecret datastructs.RedactedString, aadEndpoint string) (*ServicePrincipalTokenProvider, error) {
	if tenantID == "" {
		return nil, errors.New("tenantID must be provided")
	}
	if clientID == "" {
		return nil, errors.New("clientID must be provided")
	}
	if string(clientSecret) == "" {
		return nil, errors.New("clientSecret must be provided")
	}
	if aadEndpoint == "" {
		return nil, errors.New("aadEndpoint must be provided")
	}

	return &ServicePrincipalTokenProvider{tenantID: tenantID, clientID: clientID,
		clientSecret: clientSecret, aadEndpoint: aadEndpoint}, nil
}

func NewServicePrincipalTokenProviderWithOptions(o *ProviderOptions) (*ServicePrincipalTokenProvider, error) {
	if o.TenantID == "" {
		return nil, errors.New("tenantID must be provided")
	}
	if o.ClientID == "" {
		return nil, errors.New("clientID must be provided")
	}
	if string(o.ClientSecret) == "" {
		return nil, errors.New("clientSecret must be provided")
	}
	if o.CloudName == "" {
		return nil, errors.New("aadEndpoint must be provided")
	}

	env, err := azure.EnvironmentFromName(o.CloudName)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalTokenProvider{tenantID: o.TenantID, clientID: o.ClientID,
		clientSecret: o.ClientSecret, aadEndpoint: env.ActiveDirectoryEndpoint}, nil
}

func GetOAuthToken(ctx context.Context, spt *adal.ServicePrincipalToken) (string, error) {
	if spt == nil {
		return "", errors.New("must provide a valid service principal token")
	}

	err := spt.EnsureFreshWithContext(ctx)
	if err != nil {
		return "", err
	}
	return spt.OAuthToken(), nil
}
