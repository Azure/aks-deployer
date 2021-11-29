// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AksAppSpec defines the desired state of AksApp
type AksAppSpec struct {
	Type      string            `json:"type"`
	Version   string            `json:"version"`
	Variables map[string]string `json:"variables"`
	// TODO: Remove 'Credentials' and leave only 'Secrets'
	// +optional
	// +nullable
	Credentials map[string]string `json:"credentials"`
	// +optional
	// +nullable
	Secrets map[string]string `json:"secrets"`
	// +optional
	// +nullable
	UnmanagedSecrets []string `json:"unmanagedSecrets"`
}

// RolloutStatus is the type for rollout status
type RolloutStatus string

// These are the valid rollout statuses.
const (
	RolloutCompleted  RolloutStatus = "Completed"
	RolloutFailed     RolloutStatus = "Failed"
	RolloutInProgress RolloutStatus = "InProgress"
)

// ReconciliationResult is the type for reconciliation result
type ReconciliationResult string

// These are the valid reconciliation results.
const (
	ReconciliationSucceeded ReconciliationResult = "Succeeded"
	ReconciliationFailed    ReconciliationResult = "Failed"
)

// Reconciliation is the type for reconciliation
type Reconciliation struct {
	// +nullable
	// +optional
	LastReconcileTime metav1.Time `json:"lastReconcileTime,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	OperationID string `json:"operationId,omitempty"`
	// +optional
	Result ReconciliationResult `json:"result,omitempty"`
}

// Rollout is the type for rollout
type Rollout struct {
	// +optional
	Name string `json:"name"`
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty"`
	// +optional
	Rollout RolloutStatus `json:"rollout"`
}

// AksAppStatus defines the observed state of AksApp
type AksAppStatus struct {
	// Number of total replicas per AksApp
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Number of total unavailable replicas per AksApp
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty"`
	// Rollout version of AksApp
	// +optional
	RolloutVersion string `json:"rolloutVersion"`
	// Rollout status of AksApp
	// +optional
	Rollout RolloutStatus `json:"rollout"`
	// Reconciliation result of AksApp
	// +optional
	Reconciliation Reconciliation `json:"reconciliation,omitempty"`
	// The list of all rollouts per AksApp
	// +optional
	Rollouts []Rollout `json:"rollouts,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="VERSION",type=string,JSONPath=`.spec.version`
// +kubebuilder:printcolumn:name="TYPE",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="AGE",type=date,JSONPath=`.metadata.creationTimestamp`

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
