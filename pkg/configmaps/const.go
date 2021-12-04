// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package configmaps

const (
	aksAppConfigNameFormat = "%s-%s"
	// ClusterConfigMapName is cluster configuration ConfigMap name
	ClusterConfigMapName    = "cluster-config"
	clusterConfigNameFormat = "%s/%s"
	// ConfigDataKey is configuration ConfigMap key
	ConfigDataKey = "config"
	// ConfigAnnotationKey is configuration blob URL key
	ConfigAnnotationKey = "blob-url"
	// DefaultDeployerNamespace is deployer's default namespace
	DefaultDeployerNamespace = "deployer"

	configDataDelimiter = "---\n"
)
