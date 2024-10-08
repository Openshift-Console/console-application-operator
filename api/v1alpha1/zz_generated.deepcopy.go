//go:build !ignore_autogenerated

/*
Copyright 2024.

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
func (in *BuildConfiguration) DeepCopyInto(out *BuildConfiguration) {
	*out = *in
	out.BuilderImage = in.BuilderImage
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfiguration.
func (in *BuildConfiguration) DeepCopy() *BuildConfiguration {
	if in == nil {
		return nil
	}
	out := new(BuildConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderImage) DeepCopyInto(out *BuilderImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderImage.
func (in *BuilderImage) DeepCopy() *BuilderImage {
	if in == nil {
		return nil
	}
	out := new(BuilderImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplication) DeepCopyInto(out *ConsoleApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplication.
func (in *ConsoleApplication) DeepCopy() *ConsoleApplication {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationList) DeepCopyInto(out *ConsoleApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationList.
func (in *ConsoleApplicationList) DeepCopy() *ConsoleApplicationList {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationSpec) DeepCopyInto(out *ConsoleApplicationSpec) {
	*out = *in
	out.Git = in.Git
	in.BuildConfiguration.DeepCopyInto(&out.BuildConfiguration)
	in.DeploymentConfiguration.DeepCopyInto(&out.DeploymentConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationSpec.
func (in *ConsoleApplicationSpec) DeepCopy() *ConsoleApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationStatus) DeepCopyInto(out *ConsoleApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationStatus.
func (in *ConsoleApplicationStatus) DeepCopy() *ConsoleApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfiguration) DeepCopyInto(out *DeploymentConfiguration) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	in.Expose.DeepCopyInto(&out.Expose)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfiguration.
func (in *DeploymentConfiguration) DeepCopy() *DeploymentConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfiguration)
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
func (in *Expose) DeepCopyInto(out *Expose) {
	*out = *in
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Expose.
func (in *Expose) DeepCopy() *Expose {
	if in == nil {
		return nil
	}
	out := new(Expose)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Git) DeepCopyInto(out *Git) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Git.
func (in *Git) DeepCopy() *Git {
	if in == nil {
		return nil
	}
	out := new(Git)
	in.DeepCopyInto(out)
	return out
}
