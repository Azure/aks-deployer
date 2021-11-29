// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AksAppSpec defines the desired state of AksApp
type AksAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AksApp. Edit aksapp_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// AksAppStatus defines the observed state of AksApp
type AksAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AksApp is the Schema for the aksapps API
type AksApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AksAppSpec   `json:"spec,omitempty"`
	Status AksAppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AksAppList contains a list of AksApp
type AksAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AksApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AksApp{}, &AksAppList{})
}
