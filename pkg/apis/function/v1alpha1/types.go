package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type Function struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec FunctionSpec

	// Most recently observed status of the Deployment.
	// +optional
	Status FunctionStatus
}

type FunctionSpec struct {
	// The code of this function served for.
	Code string

	// The runtime for this function.
	Runtime string
	// Template describes the pods that will be created, if not set, use crab by default.
	// +optional
	Template api.PodTemplateSpec
}

type FunctionStatus struct {
	// Total number of non-terminated pods targeted by this function (their labels match the selector).
	// +optional
	Replicas int32

	// Represents the latest available observations of a deployment's current state.
	Conditions []FunctionCondition
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
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	// Items is the list of deployments.
	Items []Function
}
