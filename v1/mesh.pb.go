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
// source: v1/mesh.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// EdgeAttributes are pre-defined edge attributes. They should
// be used as their string values.
type EdgeAttributes int32

const (
	// EDGE_ATTRIBUTE_UNKNOWN is an unknown edge attribute.
	EdgeAttributes_EDGE_ATTRIBUTE_UNKNOWN EdgeAttributes = 0
	// EDGE_ATTRIBUTE_ICE is an ICE edge attribute.
	EdgeAttributes_EDGE_ATTRIBUTE_ICE EdgeAttributes = 1
)

// Enum value maps for EdgeAttributes.
var (
	EdgeAttributes_name = map[int32]string{
		0: "EDGE_ATTRIBUTE_UNKNOWN",
		1: "EDGE_ATTRIBUTE_ICE",
	}
	EdgeAttributes_value = map[string]int32{
		"EDGE_ATTRIBUTE_UNKNOWN": 0,
		"EDGE_ATTRIBUTE_ICE":     1,
	}
)

func (x EdgeAttributes) Enum() *EdgeAttributes {
	p := new(EdgeAttributes)
	*p = x
	return p
}

func (x EdgeAttributes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EdgeAttributes) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_mesh_proto_enumTypes[0].Descriptor()
}

func (EdgeAttributes) Type() protoreflect.EnumType {
	return &file_v1_mesh_proto_enumTypes[0]
}

func (x EdgeAttributes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EdgeAttributes.Descriptor instead.
func (EdgeAttributes) EnumDescriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{0}
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
		mi := &file_v1_mesh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[0]
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
	return file_v1_mesh_proto_rawDescGZIP(), []int{0}
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
	// primary_endpoint is the primary endpoint of the node.
	PrimaryEndpoint string `protobuf:"bytes,2,opt,name=primary_endpoint,json=primaryEndpoint,proto3" json:"primary_endpoint,omitempty"`
	// wireguard_endpoints is a list of WireGuard endpoints for the node.
	WireguardEndpoints []string `protobuf:"bytes,3,rep,name=wireguard_endpoints,json=wireguardEndpoints,proto3" json:"wireguard_endpoints,omitempty"`
	// zone_awareness_id is the zone awareness ID of the node.
	ZoneAwarenessId string `protobuf:"bytes,4,opt,name=zone_awareness_id,json=zoneAwarenessId,proto3" json:"zone_awareness_id,omitempty"`
	// raft_port is the Raft listen port of the node.
	RaftPort int32 `protobuf:"varint,5,opt,name=raft_port,json=raftPort,proto3" json:"raft_port,omitempty"`
	// grpc_port is the gRPC listen port of the node.
	GrpcPort int32 `protobuf:"varint,6,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// public_key is the public key of the node.
	PublicKey string `protobuf:"bytes,8,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// private_ipv4 is the private IPv4 address of the node.
	PrivateIpv4 string `protobuf:"bytes,9,opt,name=private_ipv4,json=privateIpv4,proto3" json:"private_ipv4,omitempty"`
	// private_ipv6 is the private IPv6 address of the node.
	PrivateIpv6 string `protobuf:"bytes,10,opt,name=private_ipv6,json=privateIpv6,proto3" json:"private_ipv6,omitempty"`
	// updated_at is the last time the node joined the cluster.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// created_at is the creation time for the node.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// cluster_status is the status of the node in the cluster.
	ClusterStatus ClusterStatus `protobuf:"varint,13,opt,name=cluster_status,json=clusterStatus,proto3,enum=v1.ClusterStatus" json:"cluster_status,omitempty"`
}

func (x *MeshNode) Reset() {
	*x = MeshNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshNode) ProtoMessage() {}

func (x *MeshNode) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[1]
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
	return file_v1_mesh_proto_rawDescGZIP(), []int{1}
}

func (x *MeshNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MeshNode) GetPrimaryEndpoint() string {
	if x != nil {
		return x.PrimaryEndpoint
	}
	return ""
}

func (x *MeshNode) GetWireguardEndpoints() []string {
	if x != nil {
		return x.WireguardEndpoints
	}
	return nil
}

func (x *MeshNode) GetZoneAwarenessId() string {
	if x != nil {
		return x.ZoneAwarenessId
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
		mi := &file_v1_mesh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[2]
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
	return file_v1_mesh_proto_rawDescGZIP(), []int{2}
}

func (x *NodeList) GetNodes() []*MeshNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// MeshEdge is an edge between two nodes.
type MeshEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source is the source node.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// target is the target node.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// weight is the weight of the edge.
	Weight int32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	// attributes is a list of attributes for the edge.
	Attributes map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MeshEdge) Reset() {
	*x = MeshEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshEdge) ProtoMessage() {}

func (x *MeshEdge) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshEdge.ProtoReflect.Descriptor instead.
func (*MeshEdge) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{3}
}

func (x *MeshEdge) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MeshEdge) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *MeshEdge) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *MeshEdge) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// MeshEdges is a list of edges.
type MeshEdges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items is the list of edges.
	Items []*MeshEdge `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *MeshEdges) Reset() {
	*x = MeshEdges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshEdges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshEdges) ProtoMessage() {}

func (x *MeshEdges) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshEdges.ProtoReflect.Descriptor instead.
func (*MeshEdges) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{4}
}

func (x *MeshEdges) GetItems() []*MeshEdge {
	if x != nil {
		return x.Items
	}
	return nil
}

// MeshGraph is a graph of nodes.
type MeshGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nodes is the list of nodes.
	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// edges is the list of edges.
	Edges []*MeshEdge `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`
	// dot is the DOT representation of the graph.
	Dot string `protobuf:"bytes,3,opt,name=dot,proto3" json:"dot,omitempty"`
}

func (x *MeshGraph) Reset() {
	*x = MeshGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshGraph) ProtoMessage() {}

func (x *MeshGraph) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshGraph.ProtoReflect.Descriptor instead.
func (*MeshGraph) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{5}
}

func (x *MeshGraph) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *MeshGraph) GetEdges() []*MeshEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

func (x *MeshGraph) GetDot() string {
	if x != nil {
		return x.Dot
	}
	return ""
}

var File_v1_mesh_proto protoreflect.FileDescriptor

var file_v1_mesh_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xf1, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x77, 0x69,
	0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x7a,
	0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x77, 0x61, 0x72, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x7a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x61, 0x72,
	0x65, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x34,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49,
	0x70, 0x76, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x49, 0x70, 0x76, 0x36, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x0e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x68, 0x45,
	0x64, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x68,
	0x45, 0x64, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x57, 0x0a, 0x09, 0x4d, 0x65, 0x73,
	0x68, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05,
	0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x6f, 0x74, 0x2a, 0x44, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x54, 0x45, 0x5f, 0x49, 0x43, 0x45, 0x10, 0x01, 0x32, 0xa3, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x73,
	0x68, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x68,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22, 0x00, 0x42, 0x65,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2,
	0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_mesh_proto_rawDescOnce sync.Once
	file_v1_mesh_proto_rawDescData = file_v1_mesh_proto_rawDesc
)

func file_v1_mesh_proto_rawDescGZIP() []byte {
	file_v1_mesh_proto_rawDescOnce.Do(func() {
		file_v1_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_mesh_proto_rawDescData)
	})
	return file_v1_mesh_proto_rawDescData
}

var file_v1_mesh_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_mesh_proto_goTypes = []interface{}{
	(EdgeAttributes)(0),           // 0: v1.EdgeAttributes
	(*GetNodeRequest)(nil),        // 1: v1.GetNodeRequest
	(*MeshNode)(nil),              // 2: v1.MeshNode
	(*NodeList)(nil),              // 3: v1.NodeList
	(*MeshEdge)(nil),              // 4: v1.MeshEdge
	(*MeshEdges)(nil),             // 5: v1.MeshEdges
	(*MeshGraph)(nil),             // 6: v1.MeshGraph
	nil,                           // 7: v1.MeshEdge.AttributesEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(ClusterStatus)(0),            // 9: v1.ClusterStatus
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_v1_mesh_proto_depIdxs = []int32{
	8,  // 0: v1.MeshNode.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 1: v1.MeshNode.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: v1.MeshNode.cluster_status:type_name -> v1.ClusterStatus
	2,  // 3: v1.NodeList.nodes:type_name -> v1.MeshNode
	7,  // 4: v1.MeshEdge.attributes:type_name -> v1.MeshEdge.AttributesEntry
	4,  // 5: v1.MeshEdges.items:type_name -> v1.MeshEdge
	4,  // 6: v1.MeshGraph.edges:type_name -> v1.MeshEdge
	1,  // 7: v1.Mesh.GetNode:input_type -> v1.GetNodeRequest
	10, // 8: v1.Mesh.ListNodes:input_type -> google.protobuf.Empty
	10, // 9: v1.Mesh.GetMeshGraph:input_type -> google.protobuf.Empty
	2,  // 10: v1.Mesh.GetNode:output_type -> v1.MeshNode
	3,  // 11: v1.Mesh.ListNodes:output_type -> v1.NodeList
	6,  // 12: v1.Mesh.GetMeshGraph:output_type -> v1.MeshGraph
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_mesh_proto_init() }
func file_v1_mesh_proto_init() {
	if File_v1_mesh_proto != nil {
		return
	}
	file_v1_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_mesh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_mesh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_mesh_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_mesh_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshEdge); i {
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
		file_v1_mesh_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshEdges); i {
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
		file_v1_mesh_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshGraph); i {
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
			RawDescriptor: file_v1_mesh_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_mesh_proto_goTypes,
		DependencyIndexes: file_v1_mesh_proto_depIdxs,
		EnumInfos:         file_v1_mesh_proto_enumTypes,
		MessageInfos:      file_v1_mesh_proto_msgTypes,
	}.Build()
	File_v1_mesh_proto = out.File
	file_v1_mesh_proto_rawDesc = nil
	file_v1_mesh_proto_goTypes = nil
	file_v1_mesh_proto_depIdxs = nil
}
