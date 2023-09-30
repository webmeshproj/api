// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using RaftLogEntry within kubernetes types, where deepcopy-gen is used.
func (in *RaftLogEntry) DeepCopyInto(out *RaftLogEntry) {
	p := proto.Clone(in).(*RaftLogEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftLogEntry. Required by controller-gen.
func (in *RaftLogEntry) DeepCopy() *RaftLogEntry {
	if in == nil {
		return nil
	}
	out := new(RaftLogEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RaftLogEntry. Required by controller-gen.
func (in *RaftLogEntry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RaftApplyResponse within kubernetes types, where deepcopy-gen is used.
func (in *RaftApplyResponse) DeepCopyInto(out *RaftApplyResponse) {
	p := proto.Clone(in).(*RaftApplyResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftApplyResponse. Required by controller-gen.
func (in *RaftApplyResponse) DeepCopy() *RaftApplyResponse {
	if in == nil {
		return nil
	}
	out := new(RaftApplyResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RaftApplyResponse. Required by controller-gen.
func (in *RaftApplyResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RaftSnapshot within kubernetes types, where deepcopy-gen is used.
func (in *RaftSnapshot) DeepCopyInto(out *RaftSnapshot) {
	p := proto.Clone(in).(*RaftSnapshot)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftSnapshot. Required by controller-gen.
func (in *RaftSnapshot) DeepCopy() *RaftSnapshot {
	if in == nil {
		return nil
	}
	out := new(RaftSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RaftSnapshot. Required by controller-gen.
func (in *RaftSnapshot) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RaftDataItem within kubernetes types, where deepcopy-gen is used.
func (in *RaftDataItem) DeepCopyInto(out *RaftDataItem) {
	p := proto.Clone(in).(*RaftDataItem)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftDataItem. Required by controller-gen.
func (in *RaftDataItem) DeepCopy() *RaftDataItem {
	if in == nil {
		return nil
	}
	out := new(RaftDataItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RaftDataItem. Required by controller-gen.
func (in *RaftDataItem) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
