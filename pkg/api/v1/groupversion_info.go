// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// Package v1 contains API Schema definitions for the deployer v1 API group
//+kubebuilder:object:generate=true
//+groupName=deployer.aks
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "deployer.aks", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
