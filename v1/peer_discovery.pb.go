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
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/peer_discovery.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListRaftPeersResponse is the response to ListPeers.
type ListRaftPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peers is the list of peers.
	Peers []*RaftPeer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *ListRaftPeersResponse) Reset() {
	*x = ListRaftPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_peer_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRaftPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRaftPeersResponse) ProtoMessage() {}

func (x *ListRaftPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_peer_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRaftPeersResponse.ProtoReflect.Descriptor instead.
func (*ListRaftPeersResponse) Descriptor() ([]byte, []int) {
	return file_v1_peer_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *ListRaftPeersResponse) GetPeers() []*RaftPeer {
	if x != nil {
		return x.Peers
	}
	return nil
}

// RaftPeer is a peer in the Raft cluster.
type RaftPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the ID of the peer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address is the public gRPC address of the peer.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Voter is whether the peer is a voter.
	Voter bool `protobuf:"varint,3,opt,name=voter,proto3" json:"voter,omitempty"`
	// Leader is whether the peer is the leader.
	Leader bool `protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *RaftPeer) Reset() {
	*x = RaftPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_peer_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftPeer) ProtoMessage() {}

func (x *RaftPeer) ProtoReflect() protoreflect.Message {
	mi := &file_v1_peer_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftPeer.ProtoReflect.Descriptor instead.
func (*RaftPeer) Descriptor() ([]byte, []int) {
	return file_v1_peer_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *RaftPeer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RaftPeer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RaftPeer) GetVoter() bool {
	if x != nil {
		return x.Voter
	}
	return false
}

func (x *RaftPeer) GetLeader() bool {
	if x != nil {
		return x.Leader
	}
	return false
}

var File_v1_peer_discovery_proto protoreflect.FileDescriptor

var file_v1_peer_discovery_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x61, 0x66, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x62, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x32, 0x51, 0x0a, 0x0d, 0x50,
	0x65, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x40, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x66, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6a,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x65, 0x65, 0x72, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca,
	0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_peer_discovery_proto_rawDescOnce sync.Once
	file_v1_peer_discovery_proto_rawDescData = file_v1_peer_discovery_proto_rawDesc
)

func file_v1_peer_discovery_proto_rawDescGZIP() []byte {
	file_v1_peer_discovery_proto_rawDescOnce.Do(func() {
		file_v1_peer_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_peer_discovery_proto_rawDescData)
	})
	return file_v1_peer_discovery_proto_rawDescData
}

var file_v1_peer_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_peer_discovery_proto_goTypes = []interface{}{
	(*ListRaftPeersResponse)(nil), // 0: v1.ListRaftPeersResponse
	(*RaftPeer)(nil),              // 1: v1.RaftPeer
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_v1_peer_discovery_proto_depIdxs = []int32{
	1, // 0: v1.ListRaftPeersResponse.peers:type_name -> v1.RaftPeer
	2, // 1: v1.PeerDiscovery.ListPeers:input_type -> google.protobuf.Empty
	0, // 2: v1.PeerDiscovery.ListPeers:output_type -> v1.ListRaftPeersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_peer_discovery_proto_init() }
func file_v1_peer_discovery_proto_init() {
	if File_v1_peer_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_peer_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRaftPeersResponse); i {
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
		file_v1_peer_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftPeer); i {
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
			RawDescriptor: file_v1_peer_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_peer_discovery_proto_goTypes,
		DependencyIndexes: file_v1_peer_discovery_proto_depIdxs,
		MessageInfos:      file_v1_peer_discovery_proto_msgTypes,
	}.Build()
	File_v1_peer_discovery_proto = out.File
	file_v1_peer_discovery_proto_rawDesc = nil
	file_v1_peer_discovery_proto_goTypes = nil
	file_v1_peer_discovery_proto_depIdxs = nil
}
