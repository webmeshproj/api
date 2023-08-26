//
//Copyright 2023.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/raft.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RaftCommandType is the type of command being sent to the
// Raft log.
type RaftCommandType int32

const (
	// UNKNOWN is the unknown command type.
	RaftCommandType_UNKNOWN RaftCommandType = 0
	// PUT is the command for putting a key/value pair.
	RaftCommandType_PUT RaftCommandType = 1
	// DELETE is the command for deleting a key/value pair.
	RaftCommandType_DELETE RaftCommandType = 2
)

// Enum value maps for RaftCommandType.
var (
	RaftCommandType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PUT",
		2: "DELETE",
	}
	RaftCommandType_value = map[string]int32{
		"UNKNOWN": 0,
		"PUT":     1,
		"DELETE":  2,
	}
)

func (x RaftCommandType) Enum() *RaftCommandType {
	p := new(RaftCommandType)
	*p = x
	return p
}

func (x RaftCommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaftCommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_raft_proto_enumTypes[0].Descriptor()
}

func (RaftCommandType) Type() protoreflect.EnumType {
	return &file_v1_raft_proto_enumTypes[0]
}

func (x RaftCommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RaftCommandType.Descriptor instead.
func (RaftCommandType) EnumDescriptor() ([]byte, []int) {
	return file_v1_raft_proto_rawDescGZIP(), []int{0}
}

// RaftLogEntry is the data of an entry in the Raft log.
type RaftLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the type of the log entry.
	Type RaftCommandType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.RaftCommandType" json:"type,omitempty"`
	// key is the key of the log entry.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value of the log entry.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// ttl is the time to live of the log entry.
	Ttl *durationpb.Duration `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *RaftLogEntry) Reset() {
	*x = RaftLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftLogEntry) ProtoMessage() {}

func (x *RaftLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftLogEntry.ProtoReflect.Descriptor instead.
func (*RaftLogEntry) Descriptor() ([]byte, []int) {
	return file_v1_raft_proto_rawDescGZIP(), []int{0}
}

func (x *RaftLogEntry) GetType() RaftCommandType {
	if x != nil {
		return x.Type
	}
	return RaftCommandType_UNKNOWN
}

func (x *RaftLogEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RaftLogEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RaftLogEntry) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// RaftApplyResponse is the response to an apply request. It
// contains the result of applying the log entry.
type RaftApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time is the total time it took to apply the log entry.
	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// error is an error that occurred during the apply.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RaftApplyResponse) Reset() {
	*x = RaftApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftApplyResponse) ProtoMessage() {}

func (x *RaftApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftApplyResponse.ProtoReflect.Descriptor instead.
func (*RaftApplyResponse) Descriptor() ([]byte, []int) {
	return file_v1_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RaftApplyResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *RaftApplyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// RaftSnapshot is the data of a snapshot.
type RaftSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kv map[string]string `protobuf:"bytes,1,rep,name=kv,proto3" json:"kv,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RaftSnapshot) Reset() {
	*x = RaftSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftSnapshot) ProtoMessage() {}

func (x *RaftSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_v1_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftSnapshot.ProtoReflect.Descriptor instead.
func (*RaftSnapshot) Descriptor() ([]byte, []int) {
	return file_v1_raft_proto_rawDescGZIP(), []int{2}
}

func (x *RaftSnapshot) GetKv() map[string]string {
	if x != nil {
		return x.Kv
	}
	return nil
}

var File_v1_raft_proto protoreflect.FileDescriptor

var file_v1_raft_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x52, 0x61, 0x66, 0x74, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74,
	0x74, 0x6c, 0x22, 0x3d, 0x0a, 0x11, 0x52, 0x61, 0x66, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x6f, 0x0a, 0x0c, 0x52, 0x61, 0x66, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x28, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x4b, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6b, 0x76, 0x1a, 0x35, 0x0a, 0x07, 0x4b,
	0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x33, 0x0a, 0x0f, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x65, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x52, 0x61, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02,
	0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_raft_proto_rawDescOnce sync.Once
	file_v1_raft_proto_rawDescData = file_v1_raft_proto_rawDesc
)

func file_v1_raft_proto_rawDescGZIP() []byte {
	file_v1_raft_proto_rawDescOnce.Do(func() {
		file_v1_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_raft_proto_rawDescData)
	})
	return file_v1_raft_proto_rawDescData
}

var file_v1_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_raft_proto_goTypes = []interface{}{
	(RaftCommandType)(0),        // 0: v1.RaftCommandType
	(*RaftLogEntry)(nil),        // 1: v1.RaftLogEntry
	(*RaftApplyResponse)(nil),   // 2: v1.RaftApplyResponse
	(*RaftSnapshot)(nil),        // 3: v1.RaftSnapshot
	nil,                         // 4: v1.RaftSnapshot.KvEntry
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
}
var file_v1_raft_proto_depIdxs = []int32{
	0, // 0: v1.RaftLogEntry.type:type_name -> v1.RaftCommandType
	5, // 1: v1.RaftLogEntry.ttl:type_name -> google.protobuf.Duration
	4, // 2: v1.RaftSnapshot.kv:type_name -> v1.RaftSnapshot.KvEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_raft_proto_init() }
func file_v1_raft_proto_init() {
	if File_v1_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftLogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftApplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftSnapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_raft_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_raft_proto_goTypes,
		DependencyIndexes: file_v1_raft_proto_depIdxs,
		EnumInfos:         file_v1_raft_proto_enumTypes,
		MessageInfos:      file_v1_raft_proto_msgTypes,
	}.Build()
	File_v1_raft_proto = out.File
	file_v1_raft_proto_rawDesc = nil
	file_v1_raft_proto_goTypes = nil
	file_v1_raft_proto_depIdxs = nil
}
