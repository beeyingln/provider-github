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
func (in *Default) DeepCopyInto(out *Default) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Default.
func (in *Default) DeepCopy() *Default {
	if in == nil {
		return nil
	}
	out := new(Default)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Default) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultList) DeepCopyInto(out *DefaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Default, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultList.
func (in *DefaultList) DeepCopy() *DefaultList {
	if in == nil {
		return nil
	}
	out := new(DefaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultObservation) DeepCopyInto(out *DefaultObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultObservation.
func (in *DefaultObservation) DeepCopy() *DefaultObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultParameters) DeepCopyInto(out *DefaultParameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultParameters.
func (in *DefaultParameters) DeepCopy() *DefaultParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSpec) DeepCopyInto(out *DefaultSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSpec.
func (in *DefaultSpec) DeepCopy() *DefaultSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultStatus) DeepCopyInto(out *DefaultStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultStatus.
func (in *DefaultStatus) DeepCopy() *DefaultStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Protection) DeepCopyInto(out *Protection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Protection.
func (in *Protection) DeepCopy() *Protection {
	if in == nil {
		return nil
	}
	out := new(Protection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Protection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionList) DeepCopyInto(out *ProtectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Protection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionList.
func (in *ProtectionList) DeepCopy() *ProtectionList {
	if in == nil {
		return nil
	}
	out := new(ProtectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionObservation) DeepCopyInto(out *ProtectionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionObservation.
func (in *ProtectionObservation) DeepCopy() *ProtectionObservation {
	if in == nil {
		return nil
	}
	out := new(ProtectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionParameters) DeepCopyInto(out *ProtectionParameters) {
	*out = *in
	if in.AllowsDeletions != nil {
		in, out := &in.AllowsDeletions, &out.AllowsDeletions
		*out = new(bool)
		**out = **in
	}
	if in.AllowsForcePushes != nil {
		in, out := &in.AllowsForcePushes, &out.AllowsForcePushes
		*out = new(bool)
		**out = **in
	}
	if in.EnforceAdmins != nil {
		in, out := &in.EnforceAdmins, &out.EnforceAdmins
		*out = new(bool)
		**out = **in
	}
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = new(string)
		**out = **in
	}
	if in.PushRestrictions != nil {
		in, out := &in.PushRestrictions, &out.PushRestrictions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RepositoryID != nil {
		in, out := &in.RepositoryID, &out.RepositoryID
		*out = new(string)
		**out = **in
	}
	if in.RequireConversationResolution != nil {
		in, out := &in.RequireConversationResolution, &out.RequireConversationResolution
		*out = new(bool)
		**out = **in
	}
	if in.RequireSignedCommits != nil {
		in, out := &in.RequireSignedCommits, &out.RequireSignedCommits
		*out = new(bool)
		**out = **in
	}
	if in.RequiredLinearHistory != nil {
		in, out := &in.RequiredLinearHistory, &out.RequiredLinearHistory
		*out = new(bool)
		**out = **in
	}
	if in.RequiredPullRequestReviews != nil {
		in, out := &in.RequiredPullRequestReviews, &out.RequiredPullRequestReviews
		*out = make([]RequiredPullRequestReviewsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequiredStatusChecks != nil {
		in, out := &in.RequiredStatusChecks, &out.RequiredStatusChecks
		*out = make([]RequiredStatusChecksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionParameters.
func (in *ProtectionParameters) DeepCopy() *ProtectionParameters {
	if in == nil {
		return nil
	}
	out := new(ProtectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionSpec) DeepCopyInto(out *ProtectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionSpec.
func (in *ProtectionSpec) DeepCopy() *ProtectionSpec {
	if in == nil {
		return nil
	}
	out := new(ProtectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionStatus) DeepCopyInto(out *ProtectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionStatus.
func (in *ProtectionStatus) DeepCopy() *ProtectionStatus {
	if in == nil {
		return nil
	}
	out := new(ProtectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3) DeepCopyInto(out *ProtectionV3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3.
func (in *ProtectionV3) DeepCopy() *ProtectionV3 {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionV3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3List) DeepCopyInto(out *ProtectionV3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProtectionV3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3List.
func (in *ProtectionV3List) DeepCopy() *ProtectionV3List {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionV3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3Observation) DeepCopyInto(out *ProtectionV3Observation) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3Observation.
func (in *ProtectionV3Observation) DeepCopy() *ProtectionV3Observation {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3Parameters) DeepCopyInto(out *ProtectionV3Parameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.EnforceAdmins != nil {
		in, out := &in.EnforceAdmins, &out.EnforceAdmins
		*out = new(bool)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.RequireConversationResolution != nil {
		in, out := &in.RequireConversationResolution, &out.RequireConversationResolution
		*out = new(bool)
		**out = **in
	}
	if in.RequireSignedCommits != nil {
		in, out := &in.RequireSignedCommits, &out.RequireSignedCommits
		*out = new(bool)
		**out = **in
	}
	if in.RequiredPullRequestReviews != nil {
		in, out := &in.RequiredPullRequestReviews, &out.RequiredPullRequestReviews
		*out = make([]ProtectionV3RequiredPullRequestReviewsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequiredStatusChecks != nil {
		in, out := &in.RequiredStatusChecks, &out.RequiredStatusChecks
		*out = make([]ProtectionV3RequiredStatusChecksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = make([]RestrictionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3Parameters.
func (in *ProtectionV3Parameters) DeepCopy() *ProtectionV3Parameters {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3RequiredPullRequestReviewsObservation) DeepCopyInto(out *ProtectionV3RequiredPullRequestReviewsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3RequiredPullRequestReviewsObservation.
func (in *ProtectionV3RequiredPullRequestReviewsObservation) DeepCopy() *ProtectionV3RequiredPullRequestReviewsObservation {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3RequiredPullRequestReviewsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3RequiredPullRequestReviewsParameters) DeepCopyInto(out *ProtectionV3RequiredPullRequestReviewsParameters) {
	*out = *in
	if in.DismissStaleReviews != nil {
		in, out := &in.DismissStaleReviews, &out.DismissStaleReviews
		*out = new(bool)
		**out = **in
	}
	if in.DismissalTeams != nil {
		in, out := &in.DismissalTeams, &out.DismissalTeams
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DismissalUsers != nil {
		in, out := &in.DismissalUsers, &out.DismissalUsers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludeAdmins != nil {
		in, out := &in.IncludeAdmins, &out.IncludeAdmins
		*out = new(bool)
		**out = **in
	}
	if in.RequireCodeOwnerReviews != nil {
		in, out := &in.RequireCodeOwnerReviews, &out.RequireCodeOwnerReviews
		*out = new(bool)
		**out = **in
	}
	if in.RequiredApprovingReviewCount != nil {
		in, out := &in.RequiredApprovingReviewCount, &out.RequiredApprovingReviewCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3RequiredPullRequestReviewsParameters.
func (in *ProtectionV3RequiredPullRequestReviewsParameters) DeepCopy() *ProtectionV3RequiredPullRequestReviewsParameters {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3RequiredPullRequestReviewsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3RequiredStatusChecksObservation) DeepCopyInto(out *ProtectionV3RequiredStatusChecksObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3RequiredStatusChecksObservation.
func (in *ProtectionV3RequiredStatusChecksObservation) DeepCopy() *ProtectionV3RequiredStatusChecksObservation {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3RequiredStatusChecksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3RequiredStatusChecksParameters) DeepCopyInto(out *ProtectionV3RequiredStatusChecksParameters) {
	*out = *in
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludeAdmins != nil {
		in, out := &in.IncludeAdmins, &out.IncludeAdmins
		*out = new(bool)
		**out = **in
	}
	if in.Strict != nil {
		in, out := &in.Strict, &out.Strict
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3RequiredStatusChecksParameters.
func (in *ProtectionV3RequiredStatusChecksParameters) DeepCopy() *ProtectionV3RequiredStatusChecksParameters {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3RequiredStatusChecksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3Spec) DeepCopyInto(out *ProtectionV3Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3Spec.
func (in *ProtectionV3Spec) DeepCopy() *ProtectionV3Spec {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionV3Status) DeepCopyInto(out *ProtectionV3Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionV3Status.
func (in *ProtectionV3Status) DeepCopy() *ProtectionV3Status {
	if in == nil {
		return nil
	}
	out := new(ProtectionV3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredPullRequestReviewsObservation) DeepCopyInto(out *RequiredPullRequestReviewsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredPullRequestReviewsObservation.
func (in *RequiredPullRequestReviewsObservation) DeepCopy() *RequiredPullRequestReviewsObservation {
	if in == nil {
		return nil
	}
	out := new(RequiredPullRequestReviewsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredPullRequestReviewsParameters) DeepCopyInto(out *RequiredPullRequestReviewsParameters) {
	*out = *in
	if in.DismissStaleReviews != nil {
		in, out := &in.DismissStaleReviews, &out.DismissStaleReviews
		*out = new(bool)
		**out = **in
	}
	if in.DismissalRestrictions != nil {
		in, out := &in.DismissalRestrictions, &out.DismissalRestrictions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequireCodeOwnerReviews != nil {
		in, out := &in.RequireCodeOwnerReviews, &out.RequireCodeOwnerReviews
		*out = new(bool)
		**out = **in
	}
	if in.RequiredApprovingReviewCount != nil {
		in, out := &in.RequiredApprovingReviewCount, &out.RequiredApprovingReviewCount
		*out = new(float64)
		**out = **in
	}
	if in.RestrictDismissals != nil {
		in, out := &in.RestrictDismissals, &out.RestrictDismissals
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredPullRequestReviewsParameters.
func (in *RequiredPullRequestReviewsParameters) DeepCopy() *RequiredPullRequestReviewsParameters {
	if in == nil {
		return nil
	}
	out := new(RequiredPullRequestReviewsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredStatusChecksObservation) DeepCopyInto(out *RequiredStatusChecksObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredStatusChecksObservation.
func (in *RequiredStatusChecksObservation) DeepCopy() *RequiredStatusChecksObservation {
	if in == nil {
		return nil
	}
	out := new(RequiredStatusChecksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredStatusChecksParameters) DeepCopyInto(out *RequiredStatusChecksParameters) {
	*out = *in
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Strict != nil {
		in, out := &in.Strict, &out.Strict
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredStatusChecksParameters.
func (in *RequiredStatusChecksParameters) DeepCopy() *RequiredStatusChecksParameters {
	if in == nil {
		return nil
	}
	out := new(RequiredStatusChecksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionsObservation) DeepCopyInto(out *RestrictionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionsObservation.
func (in *RestrictionsObservation) DeepCopy() *RestrictionsObservation {
	if in == nil {
		return nil
	}
	out := new(RestrictionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionsParameters) DeepCopyInto(out *RestrictionsParameters) {
	*out = *in
	if in.Apps != nil {
		in, out := &in.Apps, &out.Apps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionsParameters.
func (in *RestrictionsParameters) DeepCopy() *RestrictionsParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictionsParameters)
	in.DeepCopyInto(out)
	return out
}
