//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Red Hat Inc..

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConfig) DeepCopyInto(out *AppConfig) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppConfigItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConfig.
func (in *AppConfig) DeepCopy() *AppConfig {
	if in == nil {
		return nil
	}
	out := new(AppConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConfigItem) DeepCopyInto(out *AppConfigItem) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ObjectRef)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConfigItem.
func (in *AppConfigItem) DeepCopy() *AppConfigItem {
	if in == nil {
		return nil
	}
	out := new(AppConfigItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	if in.BackendAuthSecretRef != nil {
		in, out := &in.BackendAuthSecretRef, &out.BackendAuthSecretRef
		*out = new(BackendAuthSecretRef)
		**out = **in
	}
	if in.AppConfig != nil {
		in, out := &in.AppConfig, &out.AppConfig
		*out = new(AppConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DynamicPluginsConfig != nil {
		in, out := &in.DynamicPluginsConfig, &out.DynamicPluginsConfig
		*out = new(DynamicPluginsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = new(ExtraConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]EnvFrom, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendAuthSecretRef) DeepCopyInto(out *BackendAuthSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendAuthSecretRef.
func (in *BackendAuthSecretRef) DeepCopy() *BackendAuthSecretRef {
	if in == nil {
		return nil
	}
	out := new(BackendAuthSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backstage) DeepCopyInto(out *Backstage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backstage.
func (in *Backstage) DeepCopy() *Backstage {
	if in == nil {
		return nil
	}
	out := new(Backstage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backstage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageList) DeepCopyInto(out *BackstageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backstage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageList.
func (in *BackstageList) DeepCopy() *BackstageList {
	if in == nil {
		return nil
	}
	out := new(BackstageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackstageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageSpec) DeepCopyInto(out *BackstageSpec) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(Application)
		(*in).DeepCopyInto(*out)
	}
	out.RawRuntimeConfig = in.RawRuntimeConfig
	if in.SkipLocalDb != nil {
		in, out := &in.SkipLocalDb, &out.SkipLocalDb
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageSpec.
func (in *BackstageSpec) DeepCopy() *BackstageSpec {
	if in == nil {
		return nil
	}
	out := new(BackstageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackstageStatus) DeepCopyInto(out *BackstageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackstageStatus.
func (in *BackstageStatus) DeepCopy() *BackstageStatus {
	if in == nil {
		return nil
	}
	out := new(BackstageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicPluginsConfig) DeepCopyInto(out *DynamicPluginsConfig) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ObjectRef)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicPluginsConfig.
func (in *DynamicPluginsConfig) DeepCopy() *DynamicPluginsConfig {
	if in == nil {
		return nil
	}
	out := new(DynamicPluginsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvFrom) DeepCopyInto(out *EnvFrom) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ObjectRef)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvFrom.
func (in *EnvFrom) DeepCopy() *EnvFrom {
	if in == nil {
		return nil
	}
	out := new(EnvFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraConfig) DeepCopyInto(out *ExtraConfig) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExtraConfigItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraConfig.
func (in *ExtraConfig) DeepCopy() *ExtraConfig {
	if in == nil {
		return nil
	}
	out := new(ExtraConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraConfigItem) DeepCopyInto(out *ExtraConfigItem) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ObjectRef)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraConfigItem.
func (in *ExtraConfigItem) DeepCopy() *ExtraConfigItem {
	if in == nil {
		return nil
	}
	out := new(ExtraConfigItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postgresql) DeepCopyInto(out *Postgresql) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postgresql.
func (in *Postgresql) DeepCopy() *Postgresql {
	if in == nil {
		return nil
	}
	out := new(Postgresql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeConfig) DeepCopyInto(out *RuntimeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeConfig.
func (in *RuntimeConfig) DeepCopy() *RuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(RuntimeConfig)
	in.DeepCopyInto(out)
	return out
}
