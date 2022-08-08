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

type OrganizationSecretRepositoriesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrganizationSecretRepositoriesParameters struct {

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`

	// +kubebuilder:validation:Required
	SelectedRepositoryIds []*float64 `json:"selectedRepositoryIds" tf:"selected_repository_ids,omitempty"`
}

// OrganizationSecretRepositoriesSpec defines the desired state of OrganizationSecretRepositories
type OrganizationSecretRepositoriesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationSecretRepositoriesParameters `json:"forProvider"`
}

// OrganizationSecretRepositoriesStatus defines the observed state of OrganizationSecretRepositories.
type OrganizationSecretRepositoriesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationSecretRepositoriesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationSecretRepositories is the Schema for the OrganizationSecretRepositoriess API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type OrganizationSecretRepositories struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSecretRepositoriesSpec   `json:"spec"`
	Status            OrganizationSecretRepositoriesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationSecretRepositoriesList contains a list of OrganizationSecretRepositoriess
type OrganizationSecretRepositoriesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationSecretRepositories `json:"items"`
}

// Repository type metadata.
var (
	OrganizationSecretRepositories_Kind             = "OrganizationSecretRepositories"
	OrganizationSecretRepositories_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationSecretRepositories_Kind}.String()
	OrganizationSecretRepositories_KindAPIVersion   = OrganizationSecretRepositories_Kind + "." + CRDGroupVersion.String()
	OrganizationSecretRepositories_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationSecretRepositories_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationSecretRepositories{}, &OrganizationSecretRepositoriesList{})
}