package consts

const (
	// Below four consts are query parameter names msi-connector expect to receive
	// For other components in cx cluster which need to get token of an MSI-enabled cluster from
	// msi-connector, set below four query parameters correctly and send a GET request to
	// msi-connector(http://msi-connector.msi-connector.svc.cluster.local)
	QueryParameterSubscriptionID     = "subscription_id"
	QueryParameterResourceGroupName  = "resource_group_name"
	QueryParameterManagedClusterName = "managed_cluster_name"
	QueryParameterResource           = "resource"
	// HeaderRequest is an optional header send to msi-connector.
	// It is used to track an individual request
	HeaderRequestID = "x-ms-client-request-id"
	// Below two consts are expected values of `identity.type` property of resource definition
	TypeSystemAssigned = "systemassigned"
	TypeUserAssigned   = "userassigned"
	// This type will be automatically enabled after we enable user assigned identity,
	// we should disallow this case in validator.
	TypeSystemAssignedUserAssigned = "systemassigned, userassigned"
	TypeNone                       = "none"
	// Below three consts are header names in rp's http request.
	// These headers are auto-added by ARM when 'identity' property
	// is in the resource definition.
	IdentityURLHeader         = "x-ms-identity-url"
	IdentityPrincipalIDHeader = "x-ms-identity-principal-id"
	IdentityTenantIDHeader    = "x-ms-home-tenant-id"

	// Below three consts are tags of key vault secret "MSICredential"
	SubscriptionIDTag     = "subscriptionID"
	ResourceGroupNameTag  = "resourceGroupName"
	ManagedClusterNameTag = "managedClusterName"

	// KubeletIdentity is one of keys in managedCluster.IdentityProfile. Use this key to retrieve kubelet's identity in an MSI cluster
	KubeletIdentity = "kubeletidentity"

	// We need later two placeholder for some historical reason.
	// Without this placeholder, in generated ARM template, there will have syntax error.
	// see https://github.com/Azure/aks-engine/pull/2573. The easiest way to fix this is provide
	// a placeholder. I've tried but found it's really hard to fix in AKS-Engine to make it accpet
	// empty string as client id and secret.

	// ClientIDForMSICluster is ClientID property in servicePrincipalProfile
	// for MSI clusters.
	ClientIDForMSICluster = "msi"
	// SecretForMSICluster is Secret property in servicePrincipalProfile
	// for MSI clusters.
	SecretForMSICluster = "msi"
	// OmsAgentTokenSecretName is the secret name of omsagent addon's access token when using AAD auth
	/* #nosec */
	OmsAgentTokenSecretName string = "omsagent-aad-msi-token"
)

// IdentityChangeType is used to track the identity change
type IdentityChangeType int

const (
	// NewIdentity means the system identity is newly created in this operation.
	// In this case, asyncProcessor will create roleAssignment on this identity
	// and interact with MSI service to fetch the credential of this identity.
	NewIdentity IdentityChangeType = iota
	// OldIdentity means the identity already exists before this operation.
	// In this case, no action need to be performed.
	OldIdentity
	// UpdatingFromServicePrincipalToMSI belongs to NewIdentity, means the identity
	// is newly created and the cluster is updated from SPN cluster to MSI cluster.
	// UpdatingFromServicePrincipalToMSI will NOT be used in code, just write it
	// here for memorization.
	UpdatingFromServicePrincipalToMSI
	// RecoverFromMisDeletion belongs to NewIdentity, means the identity is newly created
	// because of recovery from mis-deletion.
	// UpdatingFromServicePrincipalToMSI will NOT be used in code, just write it
	// here for memorization.
	RecoverFromMisDeletion
)
