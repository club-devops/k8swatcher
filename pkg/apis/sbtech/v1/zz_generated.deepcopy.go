// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestSlack) DeepCopyInto(out *DestSlack) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestSlack.
func (in *DestSlack) DeepCopy() *DestSlack {
	if in == nil {
		return nil
	}
	out := new(DestSlack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SWatcherNotifier) DeepCopyInto(out *K8SWatcherNotifier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SWatcherNotifier.
func (in *K8SWatcherNotifier) DeepCopy() *K8SWatcherNotifier {
	if in == nil {
		return nil
	}
	out := new(K8SWatcherNotifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8SWatcherNotifier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SWatcherNotifierList) DeepCopyInto(out *K8SWatcherNotifierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K8SWatcherNotifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SWatcherNotifierList.
func (in *K8SWatcherNotifierList) DeepCopy() *K8SWatcherNotifierList {
	if in == nil {
		return nil
	}
	out := new(K8SWatcherNotifierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8SWatcherNotifierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SWatcherNotifierSpec) DeepCopyInto(out *K8SWatcherNotifierSpec) {
	*out = *in
	out.Slack = in.Slack
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SWatcherNotifierSpec.
func (in *K8SWatcherNotifierSpec) DeepCopy() *K8SWatcherNotifierSpec {
	if in == nil {
		return nil
	}
	out := new(K8SWatcherNotifierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SWatcherNotifierStatus) DeepCopyInto(out *K8SWatcherNotifierStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SWatcherNotifierStatus.
func (in *K8SWatcherNotifierStatus) DeepCopy() *K8SWatcherNotifierStatus {
	if in == nil {
		return nil
	}
	out := new(K8SWatcherNotifierStatus)
	in.DeepCopyInto(out)
	return out
}
