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
// source: v1/mesh_messages.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClusterStatus is the status of the node in the cluster.
type ClusterStatus int32

const (
	// CLUSTER_STATUS_UNKNOWN is the default status.
	ClusterStatus_CLUSTER_STATUS_UNKNOWN ClusterStatus = 0
	// CLUSTER_LEADER is the status for the leader node.
	ClusterStatus_CLUSTER_LEADER ClusterStatus = 1
	// CLUSTER_VOTER is the status for a voter node.
	ClusterStatus_CLUSTER_VOTER ClusterStatus = 2
	// CLUSTER_NON_VOTER is the status for a non-voter node.
	ClusterStatus_CLUSTER_NON_VOTER ClusterStatus = 3
)

// Enum value maps for ClusterStatus.
var (
	ClusterStatus_name = map[int32]string{
		0: "CLUSTER_STATUS_UNKNOWN",
		1: "CLUSTER_LEADER",
		2: "CLUSTER_VOTER",
		3: "CLUSTER_NON_VOTER",
	}
	ClusterStatus_value = map[string]int32{
		"CLUSTER_STATUS_UNKNOWN": 0,
		"CLUSTER_LEADER":         1,
		"CLUSTER_VOTER":          2,
		"CLUSTER_NON_VOTER":      3,
	}
)

func (x ClusterStatus) Enum() *ClusterStatus {
	p := new(ClusterStatus)
	*p = x
	return p
}

func (x ClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_mesh_messages_proto_enumTypes[0].Descriptor()
}

func (ClusterStatus) Type() protoreflect.EnumType {
	return &file_v1_mesh_messages_proto_enumTypes[0]
}

func (x ClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterStatus.Descriptor instead.
func (ClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_mesh_messages_proto_rawDescGZIP(), []int{0}
}

// Feature is a list of features supported by a node.
type Feature int32

const (
	// FEATURE_NONE is the default feature set.
	Feature_FEATURE_NONE Feature = 0
	// NODES is the feature for nodes. This is always supported.
	Feature_NODES Feature = 1
	// LEADER_PROXY is the feature for leader proxying.
	Feature_LEADER_PROXY Feature = 2
	// MESH_API is the feature for the mesh API.
	Feature_MESH_API Feature = 3
	// PEER_DISCOVERY is the feature for peer discovery.
	Feature_PEER_DISCOVERY Feature = 4
	// METRICS_GRPC is the feature for gRPC metrics.
	Feature_METRICS_GRPC Feature = 5
	// METRICS_NODE is the feature for node metrics.
	Feature_METRICS_NODE Feature = 6
	// ICE_NEGOTIATION is the feature for ICE negotiation.
	Feature_ICE_NEGOTIATION Feature = 7
)

// Enum value maps for Feature.
var (
	Feature_name = map[int32]string{
		0: "FEATURE_NONE",
		1: "NODES",
		2: "LEADER_PROXY",
		3: "MESH_API",
		4: "PEER_DISCOVERY",
		5: "METRICS_GRPC",
		6: "METRICS_NODE",
		7: "ICE_NEGOTIATION",
	}
	Feature_value = map[string]int32{
		"FEATURE_NONE":    0,
		"NODES":           1,
		"LEADER_PROXY":    2,
		"MESH_API":        3,
		"PEER_DISCOVERY":  4,
		"METRICS_GRPC":    5,
		"METRICS_NODE":    6,
		"ICE_NEGOTIATION": 7,
	}
)

func (x Feature) Enum() *Feature {
	p := new(Feature)
	*p = x
	return p
}

func (x Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_mesh_messages_proto_enumTypes[1].Descriptor()
}

func (Feature) Type() protoreflect.EnumType {
	return &file_v1_mesh_messages_proto_enumTypes[1]
}

func (x Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Feature.Descriptor instead.
func (Feature) EnumDescriptor() ([]byte, []int) {
	return file_v1_mesh_messages_proto_rawDescGZIP(), []int{1}
}

// GetNodeRequest is a request to get a node.
type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_v1_mesh_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// MeshNode is a node that has been registered with the controller.
type MeshNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// endpoint is the endpoint of the node.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// raft_port is the Raft listen port of the node.
	RaftPort int32 `protobuf:"varint,3,opt,name=raft_port,json=raftPort,proto3" json:"raft_port,omitempty"`
	// grpc_port is the gRPC listen port of the node.
	GrpcPort int32 `protobuf:"varint,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// wireguard_port is the Wireguard listen port of the node.
	WireguardPort int32 `protobuf:"varint,5,opt,name=wireguard_port,json=wireguardPort,proto3" json:"wireguard_port,omitempty"`
	// public_key is the public key of the node.
	PublicKey string `protobuf:"bytes,6,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// private_ipv4 is the private IPv4 address of the node.
	PrivateIpv4 string `protobuf:"bytes,7,opt,name=private_ipv4,json=privateIpv4,proto3" json:"private_ipv4,omitempty"`
	// private_ipv6 is the private IPv6 address of the node.
	PrivateIpv6 string `protobuf:"bytes,8,opt,name=private_ipv6,json=privateIpv6,proto3" json:"private_ipv6,omitempty"`
	// updated_at is the last time the node joined the cluster.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// created_at is the creation time for the node.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// cluster_status is the status of the node in the cluster.
	ClusterStatus ClusterStatus `protobuf:"varint,11,opt,name=cluster_status,json=clusterStatus,proto3,enum=v1.ClusterStatus" json:"cluster_status,omitempty"`
}

func (x *MeshNode) Reset() {
	*x = MeshNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshNode) ProtoMessage() {}

func (x *MeshNode) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshNode.ProtoReflect.Descriptor instead.
func (*MeshNode) Descriptor() ([]byte, []int) {
	return file_v1_mesh_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MeshNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MeshNode) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *MeshNode) GetRaftPort() int32 {
	if x != nil {
		return x.RaftPort
	}
	return 0
}

func (x *MeshNode) GetGrpcPort() int32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

func (x *MeshNode) GetWireguardPort() int32 {
	if x != nil {
		return x.WireguardPort
	}
	return 0
}

func (x *MeshNode) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *MeshNode) GetPrivateIpv4() string {
	if x != nil {
		return x.PrivateIpv4
	}
	return ""
}

func (x *MeshNode) GetPrivateIpv6() string {
	if x != nil {
		return x.PrivateIpv6
	}
	return ""
}

func (x *MeshNode) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MeshNode) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MeshNode) GetClusterStatus() ClusterStatus {
	if x != nil {
		return x.ClusterStatus
	}
	return ClusterStatus_CLUSTER_STATUS_UNKNOWN
}

// NodeList is a list of nodes.
type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nodes is the list of nodes.
	Nodes []*MeshNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_v1_mesh_messages_proto_rawDescGZIP(), []int{2}
}

func (x *NodeList) GetNodes() []*MeshNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_v1_mesh_messages_proto protoreflect.FileDescriptor

var file_v1_mesh_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x03, 0x0a,
	0x08, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x76, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x76, 0x36, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0x69, 0x0a, 0x0d, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x55, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x4f, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x56,
	0x4f, 0x54, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x93, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x44, 0x45, 0x53, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x53, 0x48, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52,
	0x59, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x47,
	0x52, 0x50, 0x43, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53,
	0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x43, 0x45, 0x5f, 0x4e,
	0x45, 0x47, 0x4f, 0x54, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x42, 0x69, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_mesh_messages_proto_rawDescOnce sync.Once
	file_v1_mesh_messages_proto_rawDescData = file_v1_mesh_messages_proto_rawDesc
)

func file_v1_mesh_messages_proto_rawDescGZIP() []byte {
	file_v1_mesh_messages_proto_rawDescOnce.Do(func() {
		file_v1_mesh_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_mesh_messages_proto_rawDescData)
	})
	return file_v1_mesh_messages_proto_rawDescData
}

var file_v1_mesh_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_mesh_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_mesh_messages_proto_goTypes = []interface{}{
	(ClusterStatus)(0),            // 0: v1.ClusterStatus
	(Feature)(0),                  // 1: v1.Feature
	(*GetNodeRequest)(nil),        // 2: v1.GetNodeRequest
	(*MeshNode)(nil),              // 3: v1.MeshNode
	(*NodeList)(nil),              // 4: v1.NodeList
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_v1_mesh_messages_proto_depIdxs = []int32{
	5, // 0: v1.MeshNode.updated_at:type_name -> google.protobuf.Timestamp
	5, // 1: v1.MeshNode.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: v1.MeshNode.cluster_status:type_name -> v1.ClusterStatus
	3, // 3: v1.NodeList.nodes:type_name -> v1.MeshNode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_mesh_messages_proto_init() }
func file_v1_mesh_messages_proto_init() {
	if File_v1_mesh_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_mesh_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_v1_mesh_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshNode); i {
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
		file_v1_mesh_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeList); i {
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
			RawDescriptor: file_v1_mesh_messages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_mesh_messages_proto_goTypes,
		DependencyIndexes: file_v1_mesh_messages_proto_depIdxs,
		EnumInfos:         file_v1_mesh_messages_proto_enumTypes,
		MessageInfos:      file_v1_mesh_messages_proto_msgTypes,
	}.Build()
	File_v1_mesh_messages_proto = out.File
	file_v1_mesh_messages_proto_rawDesc = nil
	file_v1_mesh_messages_proto_goTypes = nil
	file_v1_mesh_messages_proto_depIdxs = nil
}
