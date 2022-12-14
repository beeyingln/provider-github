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

type GroupObservation struct {
}

type GroupParameters struct {

	// +kubebuilder:validation:Required
	GroupDescription *string `json:"groupDescription" tf:"group_description,omitempty"`

	// +kubebuilder:validation:Required
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`
}

type SyncGroupMappingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SyncGroupMappingParameters struct {

	// +kubebuilder:validation:Optional
	Group []GroupParameters `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Required
	TeamSlug *string `json:"teamSlug" tf:"team_slug,omitempty"`
}

// SyncGroupMappingSpec defines the desired state of SyncGroupMapping
type SyncGroupMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SyncGroupMappingParameters `json:"forProvider"`
}

// SyncGroupMappingStatus defines the observed state of SyncGroupMapping.
type SyncGroupMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SyncGroupMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SyncGroupMapping is the Schema for the SyncGroupMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type SyncGroupMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyncGroupMappingSpec   `json:"spec"`
	Status            SyncGroupMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SyncGroupMappingList contains a list of SyncGroupMappings
type SyncGroupMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SyncGroupMapping `json:"items"`
}

// Repository type metadata.
var (
	SyncGroupMapping_Kind             = "SyncGroupMapping"
	SyncGroupMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SyncGroupMapping_Kind}.String()
	SyncGroupMapping_KindAPIVersion   = SyncGroupMapping_Kind + "." + CRDGroupVersion.String()
	SyncGroupMapping_GroupVersionKind = CRDGroupVersion.WithKind(SyncGroupMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&SyncGroupMapping{}, &SyncGroupMappingList{})
}
