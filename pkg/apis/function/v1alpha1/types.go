/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Function is a specification for a Function resource
type Function struct {
	metav1.TypeMeta       `json:",inline"`
	// +optional
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec FunctionSpec     `json:"spec"`
	// Most recently observed status of the Deployment.
	// +optional
	Status FunctionStatus `json:"status"`
}

type FunctionSpec struct {
	// The code of this function served for.
	Code string    `json:"code"`  // TODO: configmapName as code.
	// The runtime for this function.
	Runtime string `json:"runtime"`
}

type FunctionStatus struct {
	// Total number of non-terminated pods targeted by this function (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas"`
}

type FunctionConditionType string

const (
	// Progressing means the function is progressing. Progress for a function is
	// considered when a new deployment set is created or adopted
	FunctionProgressing FunctionConditionType = "Progressing"

	// Ready means the function is ready for requests.
	FunctionReady FunctionConditionType = "Ready"
)

type FunctionCondition struct {
	// Type of function condition.
	Type FunctionConditionType
	// Status of the condition, one of True, False, Unknown.
	Status api.ConditionStatus
	// The last time this condition was updated.
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	Reason string
	// A human readable message indicating details about the transition.
	Message string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FunctionList defines multiple functions.
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is the list of functions
	Items []Function `json:"items"`
}
