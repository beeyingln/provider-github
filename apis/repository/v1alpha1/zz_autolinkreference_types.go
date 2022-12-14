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

type AutolinkReferenceObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AutolinkReferenceParameters struct {

	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
	// +kubebuilder:validation:Required
	KeyPrefix *string `json:"keyPrefix" tf:"key_prefix,omitempty"`

	// The repository name
	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	// +kubebuilder:validation:Required
	TargetURLTemplate *string `json:"targetUrlTemplate" tf:"target_url_template,omitempty"`
}

// AutolinkReferenceSpec defines the desired state of AutolinkReference
type AutolinkReferenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutolinkReferenceParameters `json:"forProvider"`
}

// AutolinkReferenceStatus defines the observed state of AutolinkReference.
type AutolinkReferenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutolinkReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutolinkReference is the Schema for the AutolinkReferences API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type AutolinkReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutolinkReferenceSpec   `json:"spec"`
	Status            AutolinkReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutolinkReferenceList contains a list of AutolinkReferences
type AutolinkReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutolinkReference `json:"items"`
}

// Repository type metadata.
var (
	AutolinkReference_Kind             = "AutolinkReference"
	AutolinkReference_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutolinkReference_Kind}.String()
	AutolinkReference_KindAPIVersion   = AutolinkReference_Kind + "." + CRDGroupVersion.String()
	AutolinkReference_GroupVersionKind = CRDGroupVersion.WithKind(AutolinkReference_Kind)
)

func init() {
	SchemeBuilder.Register(&AutolinkReference{}, &AutolinkReferenceList{})
}
