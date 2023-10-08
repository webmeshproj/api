// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using NetworkACL within kubernetes types, where deepcopy-gen is used.
func (in *NetworkACL) DeepCopyInto(out *NetworkACL) {
	p := proto.Clone(in).(*NetworkACL)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACL. Required by controller-gen.
func (in *NetworkACL) DeepCopy() *NetworkACL {
	if in == nil {
		return nil
	}
	out := new(NetworkACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACL. Required by controller-gen.
func (in *NetworkACL) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using NetworkACLs within kubernetes types, where deepcopy-gen is used.
func (in *NetworkACLs) DeepCopyInto(out *NetworkACLs) {
	p := proto.Clone(in).(*NetworkACLs)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACLs. Required by controller-gen.
func (in *NetworkACLs) DeepCopy() *NetworkACLs {
	if in == nil {
		return nil
	}
	out := new(NetworkACLs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACLs. Required by controller-gen.
func (in *NetworkACLs) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Route within kubernetes types, where deepcopy-gen is used.
func (in *Route) DeepCopyInto(out *Route) {
	p := proto.Clone(in).(*Route)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route. Required by controller-gen.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Route. Required by controller-gen.
func (in *Route) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Routes within kubernetes types, where deepcopy-gen is used.
func (in *Routes) DeepCopyInto(out *Routes) {
	p := proto.Clone(in).(*Routes)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Routes. Required by controller-gen.
func (in *Routes) DeepCopy() *Routes {
	if in == nil {
		return nil
	}
	out := new(Routes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Routes. Required by controller-gen.
func (in *Routes) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using NetworkAction within kubernetes types, where deepcopy-gen is used.
func (in *NetworkAction) DeepCopyInto(out *NetworkAction) {
	p := proto.Clone(in).(*NetworkAction)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAction. Required by controller-gen.
func (in *NetworkAction) DeepCopy() *NetworkAction {
	if in == nil {
		return nil
	}
	out := new(NetworkAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAction. Required by controller-gen.
func (in *NetworkAction) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}