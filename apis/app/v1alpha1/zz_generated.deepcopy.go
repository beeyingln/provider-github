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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepository) DeepCopyInto(out *InstallationRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepository.
func (in *InstallationRepository) DeepCopy() *InstallationRepository {
	if in == nil {
		return nil
	}
	out := new(InstallationRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepositoryList) DeepCopyInto(out *InstallationRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallationRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepositoryList.
func (in *InstallationRepositoryList) DeepCopy() *InstallationRepositoryList {
	if in == nil {
		return nil
	}
	out := new(InstallationRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepositoryObservation) DeepCopyInto(out *InstallationRepositoryObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RepoID != nil {
		in, out := &in.RepoID, &out.RepoID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepositoryObservation.
func (in *InstallationRepositoryObservation) DeepCopy() *InstallationRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(InstallationRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepositoryParameters) DeepCopyInto(out *InstallationRepositoryParameters) {
	*out = *in
	if in.InstallationID != nil {
		in, out := &in.InstallationID, &out.InstallationID
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepositoryParameters.
func (in *InstallationRepositoryParameters) DeepCopy() *InstallationRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepositorySpec) DeepCopyInto(out *InstallationRepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepositorySpec.
func (in *InstallationRepositorySpec) DeepCopy() *InstallationRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(InstallationRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationRepositoryStatus) DeepCopyInto(out *InstallationRepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationRepositoryStatus.
func (in *InstallationRepositoryStatus) DeepCopy() *InstallationRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}