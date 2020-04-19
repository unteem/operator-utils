// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	status "github.com/operator-framework/operator-sdk/pkg/status"
	apis "github.com/redhat-cop/operator-utils/pkg/util/apis"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcingCRD) DeepCopyInto(out *EnforcingCRD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcingCRD.
func (in *EnforcingCRD) DeepCopy() *EnforcingCRD {
	if in == nil {
		return nil
	}
	out := new(EnforcingCRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnforcingCRD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcingCRDList) DeepCopyInto(out *EnforcingCRDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnforcingCRD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcingCRDList.
func (in *EnforcingCRDList) DeepCopy() *EnforcingCRDList {
	if in == nil {
		return nil
	}
	out := new(EnforcingCRDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnforcingCRDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcingCRDSpec) DeepCopyInto(out *EnforcingCRDSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]apis.LockedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcingCRDSpec.
func (in *EnforcingCRDSpec) DeepCopy() *EnforcingCRDSpec {
	if in == nil {
		return nil
	}
	out := new(EnforcingCRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcingCRDStatus) DeepCopyInto(out *EnforcingCRDStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcingCRDStatus.
func (in *EnforcingCRDStatus) DeepCopy() *EnforcingCRDStatus {
	if in == nil {
		return nil
	}
	out := new(EnforcingCRDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyCRD) DeepCopyInto(out *MyCRD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyCRD.
func (in *MyCRD) DeepCopy() *MyCRD {
	if in == nil {
		return nil
	}
	out := new(MyCRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MyCRD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyCRDList) DeepCopyInto(out *MyCRDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MyCRD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyCRDList.
func (in *MyCRDList) DeepCopy() *MyCRDList {
	if in == nil {
		return nil
	}
	out := new(MyCRDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MyCRDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyCRDSpec) DeepCopyInto(out *MyCRDSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyCRDSpec.
func (in *MyCRDSpec) DeepCopy() *MyCRDSpec {
	if in == nil {
		return nil
	}
	out := new(MyCRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyCRDStatus) DeepCopyInto(out *MyCRDStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyCRDStatus.
func (in *MyCRDStatus) DeepCopy() *MyCRDStatus {
	if in == nil {
		return nil
	}
	out := new(MyCRDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplatedEnforcingCRD) DeepCopyInto(out *TemplatedEnforcingCRD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplatedEnforcingCRD.
func (in *TemplatedEnforcingCRD) DeepCopy() *TemplatedEnforcingCRD {
	if in == nil {
		return nil
	}
	out := new(TemplatedEnforcingCRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplatedEnforcingCRD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplatedEnforcingCRDList) DeepCopyInto(out *TemplatedEnforcingCRDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplatedEnforcingCRD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplatedEnforcingCRDList.
func (in *TemplatedEnforcingCRDList) DeepCopy() *TemplatedEnforcingCRDList {
	if in == nil {
		return nil
	}
	out := new(TemplatedEnforcingCRDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplatedEnforcingCRDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplatedEnforcingCRDSpec) DeepCopyInto(out *TemplatedEnforcingCRDSpec) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]apis.LockedResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplatedEnforcingCRDSpec.
func (in *TemplatedEnforcingCRDSpec) DeepCopy() *TemplatedEnforcingCRDSpec {
	if in == nil {
		return nil
	}
	out := new(TemplatedEnforcingCRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplatedEnforcingCRDStatus) DeepCopyInto(out *TemplatedEnforcingCRDStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplatedEnforcingCRDStatus.
func (in *TemplatedEnforcingCRDStatus) DeepCopy() *TemplatedEnforcingCRDStatus {
	if in == nil {
		return nil
	}
	out := new(TemplatedEnforcingCRDStatus)
	in.DeepCopyInto(out)
	return out
}
