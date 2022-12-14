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

type PullRequestObservation struct {
	BaseSha *string `json:"baseSha,omitempty" tf:"base_sha,omitempty"`

	Draft *bool `json:"draft,omitempty" tf:"draft,omitempty"`

	HeadSha *string `json:"headSha,omitempty" tf:"head_sha,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`

	OpenedAt *float64 `json:"openedAt,omitempty" tf:"opened_at,omitempty"`

	OpenedBy *string `json:"openedBy,omitempty" tf:"opened_by,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type PullRequestParameters struct {

	// +kubebuilder:validation:Required
	BaseRef *string `json:"baseRef" tf:"base_ref,omitempty"`

	// +kubebuilder:validation:Required
	BaseRepository *string `json:"baseRepository" tf:"base_repository,omitempty"`

	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// +kubebuilder:validation:Required
	HeadRef *string `json:"headRef" tf:"head_ref,omitempty"`

	// +kubebuilder:validation:Optional
	MaintainerCanModify *bool `json:"maintainerCanModify,omitempty" tf:"maintainer_can_modify,omitempty"`

	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// PullRequestSpec defines the desired state of PullRequest
type PullRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PullRequestParameters `json:"forProvider"`
}

// PullRequestStatus defines the observed state of PullRequest.
type PullRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PullRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PullRequest is the Schema for the PullRequests API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type PullRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PullRequestSpec   `json:"spec"`
	Status            PullRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PullRequestList contains a list of PullRequests
type PullRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PullRequest `json:"items"`
}

// Repository type metadata.
var (
	PullRequest_Kind             = "PullRequest"
	PullRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PullRequest_Kind}.String()
	PullRequest_KindAPIVersion   = PullRequest_Kind + "." + CRDGroupVersion.String()
	PullRequest_GroupVersionKind = CRDGroupVersion.WithKind(PullRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&PullRequest{}, &PullRequestList{})
}
