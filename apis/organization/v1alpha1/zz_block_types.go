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

type BlockObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BlockParameters struct {

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// BlockSpec defines the desired state of Block
type BlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlockParameters `json:"forProvider"`
}

// BlockStatus defines the observed state of Block.
type BlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Block is the Schema for the Blocks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type Block struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlockSpec   `json:"spec"`
	Status            BlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlockList contains a list of Blocks
type BlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Block `json:"items"`
}

// Repository type metadata.
var (
	Block_Kind             = "Block"
	Block_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Block_Kind}.String()
	Block_KindAPIVersion   = Block_Kind + "." + CRDGroupVersion.String()
	Block_GroupVersionKind = CRDGroupVersion.WithKind(Block_Kind)
)

func init() {
	SchemeBuilder.Register(&Block{}, &BlockList{})
}
