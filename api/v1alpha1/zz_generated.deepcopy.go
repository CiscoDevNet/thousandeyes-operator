// +build !ignore_autogenerated

/*
Copyright 2021.

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
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = make([]Agent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThousandEyesTest) DeepCopyInto(out *ThousandEyesTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThousandEyesTest.
func (in *ThousandEyesTest) DeepCopy() *ThousandEyesTest {
	if in == nil {
		return nil
	}
	out := new(ThousandEyesTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThousandEyesTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThousandEyesTestList) DeepCopyInto(out *ThousandEyesTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThousandEyesTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThousandEyesTestList.
func (in *ThousandEyesTestList) DeepCopy() *ThousandEyesTestList {
	if in == nil {
		return nil
	}
	out := new(ThousandEyesTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThousandEyesTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThousandEyesTestSpec) DeepCopyInto(out *ThousandEyesTestSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThousandEyesTestSpec.
func (in *ThousandEyesTestSpec) DeepCopy() *ThousandEyesTestSpec {
	if in == nil {
		return nil
	}
	out := new(ThousandEyesTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThousandEyesTestStatus) DeepCopyInto(out *ThousandEyesTestStatus) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThousandEyesTestStatus.
func (in *ThousandEyesTestStatus) DeepCopy() *ThousandEyesTestStatus {
	if in == nil {
		return nil
	}
	out := new(ThousandEyesTestStatus)
	in.DeepCopyInto(out)
	return out
}
