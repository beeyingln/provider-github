//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedActionsConfigObservation) DeepCopyInto(out *AllowedActionsConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedActionsConfigObservation.
func (in *AllowedActionsConfigObservation) DeepCopy() *AllowedActionsConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedActionsConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedActionsConfigParameters) DeepCopyInto(out *AllowedActionsConfigParameters) {
	*out = *in
	if in.GithubOwnedAllowed != nil {
		in, out := &in.GithubOwnedAllowed, &out.GithubOwnedAllowed
		*out = new(bool)
		**out = **in
	}
	if in.PatternsAllowed != nil {
		in, out := &in.PatternsAllowed, &out.PatternsAllowed
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VerifiedAllowed != nil {
		in, out := &in.VerifiedAllowed, &out.VerifiedAllowed
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedActionsConfigParameters.
func (in *AllowedActionsConfigParameters) DeepCopy() *AllowedActionsConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedActionsConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnabledRepositoriesConfigObservation) DeepCopyInto(out *EnabledRepositoriesConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnabledRepositoriesConfigObservation.
func (in *EnabledRepositoriesConfigObservation) DeepCopy() *EnabledRepositoriesConfigObservation {
	if in == nil {
		return nil
	}
	out := new(EnabledRepositoriesConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnabledRepositoriesConfigParameters) DeepCopyInto(out *EnabledRepositoriesConfigParameters) {
	*out = *in
	if in.RepositoryIds != nil {
		in, out := &in.RepositoryIds, &out.RepositoryIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnabledRepositoriesConfigParameters.
func (in *EnabledRepositoriesConfigParameters) DeepCopy() *EnabledRepositoriesConfigParameters {
	if in == nil {
		return nil
	}
	out := new(EnabledRepositoriesConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecret) DeepCopyInto(out *EnvironmentSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecret.
func (in *EnvironmentSecret) DeepCopy() *EnvironmentSecret {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecretList) DeepCopyInto(out *EnvironmentSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecretList.
func (in *EnvironmentSecretList) DeepCopy() *EnvironmentSecretList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecretObservation) DeepCopyInto(out *EnvironmentSecretObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecretObservation.
func (in *EnvironmentSecretObservation) DeepCopy() *EnvironmentSecretObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecretParameters) DeepCopyInto(out *EnvironmentSecretParameters) {
	*out = *in
	if in.EncryptedValueSecretRef != nil {
		in, out := &in.EncryptedValueSecretRef, &out.EncryptedValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(string)
		**out = **in
	}
	if in.PlaintextValueSecretRef != nil {
		in, out := &in.PlaintextValueSecretRef, &out.PlaintextValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecretParameters.
func (in *EnvironmentSecretParameters) DeepCopy() *EnvironmentSecretParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecretSpec) DeepCopyInto(out *EnvironmentSecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecretSpec.
func (in *EnvironmentSecretSpec) DeepCopy() *EnvironmentSecretSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSecretStatus) DeepCopyInto(out *EnvironmentSecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSecretStatus.
func (in *EnvironmentSecretStatus) DeepCopy() *EnvironmentSecretStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissions) DeepCopyInto(out *OrganizationPermissions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissions.
func (in *OrganizationPermissions) DeepCopy() *OrganizationPermissions {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationPermissions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissionsList) DeepCopyInto(out *OrganizationPermissionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationPermissions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissionsList.
func (in *OrganizationPermissionsList) DeepCopy() *OrganizationPermissionsList {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationPermissionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissionsObservation) DeepCopyInto(out *OrganizationPermissionsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissionsObservation.
func (in *OrganizationPermissionsObservation) DeepCopy() *OrganizationPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissionsParameters) DeepCopyInto(out *OrganizationPermissionsParameters) {
	*out = *in
	if in.AllowedActions != nil {
		in, out := &in.AllowedActions, &out.AllowedActions
		*out = new(string)
		**out = **in
	}
	if in.AllowedActionsConfig != nil {
		in, out := &in.AllowedActionsConfig, &out.AllowedActionsConfig
		*out = make([]AllowedActionsConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnabledRepositories != nil {
		in, out := &in.EnabledRepositories, &out.EnabledRepositories
		*out = new(string)
		**out = **in
	}
	if in.EnabledRepositoriesConfig != nil {
		in, out := &in.EnabledRepositoriesConfig, &out.EnabledRepositoriesConfig
		*out = make([]EnabledRepositoriesConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissionsParameters.
func (in *OrganizationPermissionsParameters) DeepCopy() *OrganizationPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissionsSpec) DeepCopyInto(out *OrganizationPermissionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissionsSpec.
func (in *OrganizationPermissionsSpec) DeepCopy() *OrganizationPermissionsSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPermissionsStatus) DeepCopyInto(out *OrganizationPermissionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPermissionsStatus.
func (in *OrganizationPermissionsStatus) DeepCopy() *OrganizationPermissionsStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationPermissionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecret) DeepCopyInto(out *OrganizationSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecret.
func (in *OrganizationSecret) DeepCopy() *OrganizationSecret {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretList) DeepCopyInto(out *OrganizationSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretList.
func (in *OrganizationSecretList) DeepCopy() *OrganizationSecretList {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretObservation) DeepCopyInto(out *OrganizationSecretObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretObservation.
func (in *OrganizationSecretObservation) DeepCopy() *OrganizationSecretObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretParameters) DeepCopyInto(out *OrganizationSecretParameters) {
	*out = *in
	if in.EncryptedValueSecretRef != nil {
		in, out := &in.EncryptedValueSecretRef, &out.EncryptedValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PlaintextValueSecretRef != nil {
		in, out := &in.PlaintextValueSecretRef, &out.PlaintextValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	if in.SelectedRepositoryIds != nil {
		in, out := &in.SelectedRepositoryIds, &out.SelectedRepositoryIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretParameters.
func (in *OrganizationSecretParameters) DeepCopy() *OrganizationSecretParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositories) DeepCopyInto(out *OrganizationSecretRepositories) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositories.
func (in *OrganizationSecretRepositories) DeepCopy() *OrganizationSecretRepositories {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositories)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationSecretRepositories) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositoriesList) DeepCopyInto(out *OrganizationSecretRepositoriesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationSecretRepositories, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositoriesList.
func (in *OrganizationSecretRepositoriesList) DeepCopy() *OrganizationSecretRepositoriesList {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositoriesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationSecretRepositoriesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositoriesObservation) DeepCopyInto(out *OrganizationSecretRepositoriesObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositoriesObservation.
func (in *OrganizationSecretRepositoriesObservation) DeepCopy() *OrganizationSecretRepositoriesObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositoriesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositoriesParameters) DeepCopyInto(out *OrganizationSecretRepositoriesParameters) {
	*out = *in
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	if in.SelectedRepositoryIds != nil {
		in, out := &in.SelectedRepositoryIds, &out.SelectedRepositoryIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositoriesParameters.
func (in *OrganizationSecretRepositoriesParameters) DeepCopy() *OrganizationSecretRepositoriesParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositoriesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositoriesSpec) DeepCopyInto(out *OrganizationSecretRepositoriesSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositoriesSpec.
func (in *OrganizationSecretRepositoriesSpec) DeepCopy() *OrganizationSecretRepositoriesSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositoriesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretRepositoriesStatus) DeepCopyInto(out *OrganizationSecretRepositoriesStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretRepositoriesStatus.
func (in *OrganizationSecretRepositoriesStatus) DeepCopy() *OrganizationSecretRepositoriesStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretRepositoriesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretSpec) DeepCopyInto(out *OrganizationSecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretSpec.
func (in *OrganizationSecretSpec) DeepCopy() *OrganizationSecretSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSecretStatus) DeepCopyInto(out *OrganizationSecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSecretStatus.
func (in *OrganizationSecretStatus) DeepCopy() *OrganizationSecretStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroup) DeepCopyInto(out *RunnerGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroup.
func (in *RunnerGroup) DeepCopy() *RunnerGroup {
	if in == nil {
		return nil
	}
	out := new(RunnerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroupList) DeepCopyInto(out *RunnerGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RunnerGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroupList.
func (in *RunnerGroupList) DeepCopy() *RunnerGroupList {
	if in == nil {
		return nil
	}
	out := new(RunnerGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroupObservation) DeepCopyInto(out *RunnerGroupObservation) {
	*out = *in
	if in.AllowsPublicRepositories != nil {
		in, out := &in.AllowsPublicRepositories, &out.AllowsPublicRepositories
		*out = new(bool)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Inherited != nil {
		in, out := &in.Inherited, &out.Inherited
		*out = new(bool)
		**out = **in
	}
	if in.RunnersURL != nil {
		in, out := &in.RunnersURL, &out.RunnersURL
		*out = new(string)
		**out = **in
	}
	if in.SelectedRepositoriesURL != nil {
		in, out := &in.SelectedRepositoriesURL, &out.SelectedRepositoriesURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroupObservation.
func (in *RunnerGroupObservation) DeepCopy() *RunnerGroupObservation {
	if in == nil {
		return nil
	}
	out := new(RunnerGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroupParameters) DeepCopyInto(out *RunnerGroupParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SelectedRepositoryIds != nil {
		in, out := &in.SelectedRepositoryIds, &out.SelectedRepositoryIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroupParameters.
func (in *RunnerGroupParameters) DeepCopy() *RunnerGroupParameters {
	if in == nil {
		return nil
	}
	out := new(RunnerGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroupSpec) DeepCopyInto(out *RunnerGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroupSpec.
func (in *RunnerGroupSpec) DeepCopy() *RunnerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerGroupStatus) DeepCopyInto(out *RunnerGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerGroupStatus.
func (in *RunnerGroupStatus) DeepCopy() *RunnerGroupStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Secret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretList) DeepCopyInto(out *SecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretList.
func (in *SecretList) DeepCopy() *SecretList {
	if in == nil {
		return nil
	}
	out := new(SecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObservation) DeepCopyInto(out *SecretObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObservation.
func (in *SecretObservation) DeepCopy() *SecretObservation {
	if in == nil {
		return nil
	}
	out := new(SecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretParameters) DeepCopyInto(out *SecretParameters) {
	*out = *in
	if in.EncryptedValueSecretRef != nil {
		in, out := &in.EncryptedValueSecretRef, &out.EncryptedValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PlaintextValueSecretRef != nil {
		in, out := &in.PlaintextValueSecretRef, &out.PlaintextValueSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretParameters.
func (in *SecretParameters) DeepCopy() *SecretParameters {
	if in == nil {
		return nil
	}
	out := new(SecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}
