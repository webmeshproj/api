// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using RegisterRequest within kubernetes types, where deepcopy-gen is used.
func (in *RegisterRequest) DeepCopyInto(out *RegisterRequest) {
	p := proto.Clone(in).(*RegisterRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterRequest. Required by controller-gen.
func (in *RegisterRequest) DeepCopy() *RegisterRequest {
	if in == nil {
		return nil
	}
	out := new(RegisterRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RegisterRequest. Required by controller-gen.
func (in *RegisterRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RegisterResponse within kubernetes types, where deepcopy-gen is used.
func (in *RegisterResponse) DeepCopyInto(out *RegisterResponse) {
	p := proto.Clone(in).(*RegisterResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterResponse. Required by controller-gen.
func (in *RegisterResponse) DeepCopy() *RegisterResponse {
	if in == nil {
		return nil
	}
	out := new(RegisterResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RegisterResponse. Required by controller-gen.
func (in *RegisterResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LookupRequest within kubernetes types, where deepcopy-gen is used.
func (in *LookupRequest) DeepCopyInto(out *LookupRequest) {
	p := proto.Clone(in).(*LookupRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LookupRequest. Required by controller-gen.
func (in *LookupRequest) DeepCopy() *LookupRequest {
	if in == nil {
		return nil
	}
	out := new(LookupRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LookupRequest. Required by controller-gen.
func (in *LookupRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LookupResponse within kubernetes types, where deepcopy-gen is used.
func (in *LookupResponse) DeepCopyInto(out *LookupResponse) {
	p := proto.Clone(in).(*LookupResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LookupResponse. Required by controller-gen.
func (in *LookupResponse) DeepCopy() *LookupResponse {
	if in == nil {
		return nil
	}
	out := new(LookupResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LookupResponse. Required by controller-gen.
func (in *LookupResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
