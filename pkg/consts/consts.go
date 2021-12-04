// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package consts

const (
	// DefaultPrivateEndpointName is the default private endpoint name created in the MC resource group
	DefaultPrivateEndpointName = "kube-apiserver"
	// Kilobyte is the number of bytes in a Kilobyte (technically, Kebibyte).
	// We use this value primarily for unit conversions between GB, MB, KB, and bytes.
	Kilobyte = 1024
	// Megabyte is the number of bytes in a megabyte. Technically a Mebibyte.
	Megabyte = Kilobyte * Kilobyte
	// Gigabyte is the number of bytes in a gigabyte. Technically a Gibibyte.
	Gigabyte = Megabyte * Kilobyte

	// HyperVGen2MinimumK8sVersion is the minimum Kubernetes version for which AKS
	// will default to deploying Generation 2 VMs, assuming the VM size supports it.
	HyperVGen2MinimumK8sVersion = "1.18.0"

	// GoalStateOnlyQuery is used as a query key in url
	// Example: <resource-id>?goalStateOnly=true
	GoalStateOnlyQuery = "goalStateOnly"

	// AKSPrefix is used for both system annotation and label keys
	AKSPrefixValue = "kubernetes.azure.com"
	AKSPrefix      = AKSPrefixValue + "/"

	// NetworkAssociationLinkName is the name for NetworkAssociationLink
	NetworkAssociationLinkName = "kube-apiserver"
	// AksServiceName - used in subnet delegation
	AksServiceName = "Microsoft.ContainerService/managedClusters"
	// AksDelegationName
	AKSDelegationName = "aks-delegation"

	// ccp kube-apiserver deployment name
	APIServerDeploymentName = "kube-apiserver"

	// E2EFixedAgentPoolNameSeed is the hard coded agent pool name seed used in E2E tests.
	E2EFixedAgentPoolNameSeed = int64(42)
)
