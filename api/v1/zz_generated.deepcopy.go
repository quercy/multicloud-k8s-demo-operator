// +build !ignore_autogenerated

/*


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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Demo) DeepCopyInto(out *Demo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Demo.
func (in *Demo) DeepCopy() *Demo {
	if in == nil {
		return nil
	}
	out := new(Demo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Demo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DemoList) DeepCopyInto(out *DemoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Demo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DemoList.
func (in *DemoList) DeepCopy() *DemoList {
	if in == nil {
		return nil
	}
	out := new(DemoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DemoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DemoSpec) DeepCopyInto(out *DemoSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DemoSpec.
func (in *DemoSpec) DeepCopy() *DemoSpec {
	if in == nil {
		return nil
	}
	out := new(DemoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DemoStatus) DeepCopyInto(out *DemoStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DemoStatus.
func (in *DemoStatus) DeepCopy() *DemoStatus {
	if in == nil {
		return nil
	}
	out := new(DemoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G1) DeepCopyInto(out *G1) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G1.
func (in *G1) DeepCopy() *G1 {
	if in == nil {
		return nil
	}
	out := new(G1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcMethod) DeepCopyInto(out *GcMethod) {
	*out = *in
	out.G1 = in.G1
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcMethod.
func (in *GcMethod) DeepCopy() *GcMethod {
	if in == nil {
		return nil
	}
	out := new(GcMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.SecurityContext = in.SecurityContext
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmConfig) DeepCopyInto(out *JvmConfig) {
	*out = *in
	out.GcMethod = in.GcMethod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmConfig.
func (in *JvmConfig) DeepCopy() *JvmConfig {
	if in == nil {
		return nil
	}
	out := new(JvmConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Presto) DeepCopyInto(out *Presto) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Presto.
func (in *Presto) DeepCopy() *Presto {
	if in == nil {
		return nil
	}
	out := new(Presto)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Presto) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoConfig) DeepCopyInto(out *PrestoConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoConfig.
func (in *PrestoConfig) DeepCopy() *PrestoConfig {
	if in == nil {
		return nil
	}
	out := new(PrestoConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoList) DeepCopyInto(out *PrestoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Presto, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoList.
func (in *PrestoList) DeepCopy() *PrestoList {
	if in == nil {
		return nil
	}
	out := new(PrestoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrestoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoNode) DeepCopyInto(out *PrestoNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoNode.
func (in *PrestoNode) DeepCopy() *PrestoNode {
	if in == nil {
		return nil
	}
	out := new(PrestoNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoSpec) DeepCopyInto(out *PrestoSpec) {
	*out = *in
	out.Node = in.Node
	out.Config = in.Config
	out.JvmConfig = in.JvmConfig
	out.Image = in.Image
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoSpec.
func (in *PrestoSpec) DeepCopy() *PrestoSpec {
	if in == nil {
		return nil
	}
	out := new(PrestoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoStatus) DeepCopyInto(out *PrestoStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoStatus.
func (in *PrestoStatus) DeepCopy() *PrestoStatus {
	if in == nil {
		return nil
	}
	out := new(PrestoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContext) DeepCopyInto(out *SecurityContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContext.
func (in *SecurityContext) DeepCopy() *SecurityContext {
	if in == nil {
		return nil
	}
	out := new(SecurityContext)
	in.DeepCopyInto(out)
	return out
}
