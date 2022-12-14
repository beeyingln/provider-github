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

type OrganizationSecretObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type OrganizationSecretParameters struct {

	// +kubebuilder:validation:Optional
	EncryptedValueSecretRef *v1.SecretKeySelector `json:"encryptedValueSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PlaintextValueSecretRef *v1.SecretKeySelector `json:"plaintextValueSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`

	// +kubebuilder:validation:Optional
	SelectedRepositoryIds []*float64 `json:"selectedRepositoryIds,omitempty" tf:"selected_repository_ids,omitempty"`

	// +kubebuilder:validation:Required
	Visibility *string `json:"visibility" tf:"visibility,omitempty"`
}

// OrganizationSecretSpec defines the desired state of OrganizationSecret
type OrganizationSecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationSecretParameters `json:"forProvider"`
}

// OrganizationSecretStatus defines the observed state of OrganizationSecret.
type OrganizationSecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationSecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationSecret is the Schema for the OrganizationSecrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type OrganizationSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSecretSpec   `json:"spec"`
	Status            OrganizationSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationSecretList contains a list of OrganizationSecrets
type OrganizationSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationSecret `json:"items"`
}

// Repository type metadata.
var (
	OrganizationSecret_Kind             = "OrganizationSecret"
	OrganizationSecret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationSecret_Kind}.String()
	OrganizationSecret_KindAPIVersion   = OrganizationSecret_Kind + "." + CRDGroupVersion.String()
	OrganizationSecret_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationSecret_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationSecret{}, &OrganizationSecretList{})
}
