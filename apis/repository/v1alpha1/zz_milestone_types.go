/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MilestoneObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`
}

type MilestoneParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// in yyyy-mm-dd format
	// +kubebuilder:validation:Optional
	DueDate *string `json:"dueDate,omitempty" tf:"due_date,omitempty"`

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// MilestoneSpec defines the desired state of Milestone
type MilestoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MilestoneParameters `json:"forProvider"`
}

// MilestoneStatus defines the observed state of Milestone.
type MilestoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MilestoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Milestone is the Schema for the Milestones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type Milestone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MilestoneSpec   `json:"spec"`
	Status            MilestoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MilestoneList contains a list of Milestones
type MilestoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Milestone `json:"items"`
}

// Repository type metadata.
var (
	Milestone_Kind             = "Milestone"
	Milestone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Milestone_Kind}.String()
	Milestone_KindAPIVersion   = Milestone_Kind + "." + CRDGroupVersion.String()
	Milestone_GroupVersionKind = CRDGroupVersion.WithKind(Milestone_Kind)
)

func init() {
	SchemeBuilder.Register(&Milestone{}, &MilestoneList{})
}