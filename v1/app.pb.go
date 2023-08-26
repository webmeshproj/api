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
// source: v1/app.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusResponse_ConnectionStatus int32

const (
	// DISCONNECTED indicates that the node is not connected to a mesh.
	StatusResponse_DISCONNECTED StatusResponse_ConnectionStatus = 0
	// CONNECTING indicates that the node is in the process of connecting to a mesh.
	StatusResponse_CONNECTING StatusResponse_ConnectionStatus = 1
	// CONNECTED indicates that the node is connected to a mesh.
	StatusResponse_CONNECTED StatusResponse_ConnectionStatus = 2
)

// Enum value maps for StatusResponse_ConnectionStatus.
var (
	StatusResponse_ConnectionStatus_name = map[int32]string{
		0: "DISCONNECTED",
		1: "CONNECTING",
		2: "CONNECTED",
	}
	StatusResponse_ConnectionStatus_value = map[string]int32{
		"DISCONNECTED": 0,
		"CONNECTING":   1,
		"CONNECTED":    2,
	}
)

func (x StatusResponse_ConnectionStatus) Enum() *StatusResponse_ConnectionStatus {
	p := new(StatusResponse_ConnectionStatus)
	*p = x
	return p
}

func (x StatusResponse_ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusResponse_ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_app_proto_enumTypes[0].Descriptor()
}

func (StatusResponse_ConnectionStatus) Type() protoreflect.EnumType {
	return &file_v1_app_proto_enumTypes[0]
}

func (x StatusResponse_ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusResponse_ConnectionStatus.Descriptor instead.
func (StatusResponse_ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{7, 0}
}

// ConnectRequest is sent by the application to the node to establish a
// connection to a mesh. This message will eventually contain unique
// identifiers to allow creating connections to multiple meshes.
type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Config is used to override any defaults configured on the node.
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// Disable bootstrap tells a node that is otherwise configured to
	// bootstrap to not bootstrap for this connection.
	DisableBootstrap bool `protobuf:"varint,2,opt,name=disable_bootstrap,json=disableBootstrap,proto3" json:"disable_bootstrap,omitempty"`
	// Join PSK is the pre-shared key to use for joining the mesh.
	JoinPsk string `protobuf:"bytes,3,opt,name=join_psk,json=joinPsk,proto3" json:"join_psk,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ConnectRequest) GetDisableBootstrap() bool {
	if x != nil {
		return x.DisableBootstrap
	}
	return false
}

func (x *ConnectRequest) GetJoinPsk() string {
	if x != nil {
		return x.JoinPsk
	}
	return ""
}

// ConnectResponse is returned by the Connect RPC.
type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// node id is the unique identifier of the node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// mesh domain is the domain of the mesh.
	MeshDomain string `protobuf:"bytes,2,opt,name=mesh_domain,json=meshDomain,proto3" json:"mesh_domain,omitempty"`
	// ipv4 is the IPv4 address of the node.
	Ipv4 string `protobuf:"bytes,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// ipv6 is the IPv6 address of the node.
	Ipv6 string `protobuf:"bytes,4,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ConnectResponse) GetMeshDomain() string {
	if x != nil {
		return x.MeshDomain
	}
	return ""
}

func (x *ConnectResponse) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *ConnectResponse) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

// DisconnectRequest is sent by the application to the node to disconnect
// from a mesh. This message will eventually contain unique identifiers
// for allowing the application to disconnect from a specific mesh.
type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{2}
}

// DisconnectResponse is returned by the Disconnect RPC.
type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{3}
}

// MetricsRequest is sent by the application to the node to retrieve interface
// metrics. It is intentionally empty for now, but can eventually be used to
// query specific interfaces/metrics.
type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{4}
}

// MetricsResponse is a message containing interface metrics.
type MetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// interfaces is a map of interface names to metrics.
	Interfaces map[string]*InterfaceMetrics `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetricsResponse) Reset() {
	*x = MetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsResponse) ProtoMessage() {}

func (x *MetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsResponse.ProtoReflect.Descriptor instead.
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{5}
}

func (x *MetricsResponse) GetInterfaces() map[string]*InterfaceMetrics {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

// StatusRequest is sent by the application to the node to retrieve the status
// of the node.
type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{6}
}

// StatusResponse is a message containing the status of the node.
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// connection status is the status of the connection.
	ConnectionStatus StatusResponse_ConnectionStatus `protobuf:"varint,1,opt,name=connection_status,json=connectionStatus,proto3,enum=v1.StatusResponse_ConnectionStatus" json:"connection_status,omitempty"`
	// node is the node status. This is only populated if the node is connected.
	Node *MeshNode `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{7}
}

func (x *StatusResponse) GetConnectionStatus() StatusResponse_ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return StatusResponse_DISCONNECTED
}

func (x *StatusResponse) GetNode() *MeshNode {
	if x != nil {
		return x.Node
	}
	return nil
}

// AnnounceDHTRequest is sent by the application to the node to announce the
// node's presence on the Kademlia DHT for other nodes to discover.
type AnnounceDHTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bootstrap servers are optional bootstrap servers to use for bootstrapping
	// the DHT. If not provided, the node will use the default bootstrap servers.
	BootstrapServers []string `protobuf:"bytes,1,rep,name=bootstrap_servers,json=bootstrapServers,proto3" json:"bootstrap_servers,omitempty"`
	// PSK is the pre-shared key to use for the DHT.
	Psk string `protobuf:"bytes,2,opt,name=psk,proto3" json:"psk,omitempty"`
}

func (x *AnnounceDHTRequest) Reset() {
	*x = AnnounceDHTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceDHTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceDHTRequest) ProtoMessage() {}

func (x *AnnounceDHTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceDHTRequest.ProtoReflect.Descriptor instead.
func (*AnnounceDHTRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{8}
}

func (x *AnnounceDHTRequest) GetBootstrapServers() []string {
	if x != nil {
		return x.BootstrapServers
	}
	return nil
}

func (x *AnnounceDHTRequest) GetPsk() string {
	if x != nil {
		return x.Psk
	}
	return ""
}

// AnnounceDHTResponse is returned by the AnnounceDHT RPC.
type AnnounceDHTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnnounceDHTResponse) Reset() {
	*x = AnnounceDHTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceDHTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceDHTResponse) ProtoMessage() {}

func (x *AnnounceDHTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceDHTResponse.ProtoReflect.Descriptor instead.
func (*AnnounceDHTResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{9}
}

// LeaveDHTRequest is sent by the application to the node to leave the Kademlia
// DHT.
type LeaveDHTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PSK is the pre-shared key that was used to join the DHT.
	Psk string `protobuf:"bytes,1,opt,name=psk,proto3" json:"psk,omitempty"`
}

func (x *LeaveDHTRequest) Reset() {
	*x = LeaveDHTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveDHTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveDHTRequest) ProtoMessage() {}

func (x *LeaveDHTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveDHTRequest.ProtoReflect.Descriptor instead.
func (*LeaveDHTRequest) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{10}
}

func (x *LeaveDHTRequest) GetPsk() string {
	if x != nil {
		return x.Psk
	}
	return ""
}

// LeaveDHTResponse is returned by the LeaveDHT RPC.
type LeaveDHTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveDHTResponse) Reset() {
	*x = LeaveDHTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_app_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveDHTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveDHTResponse) ProtoMessage() {}

func (x *LeaveDHTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_app_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveDHTResponse.ProtoReflect.Descriptor instead.
func (*LeaveDHTResponse) Descriptor() ([]byte, []int) {
	return file_v1_app_proto_rawDescGZIP(), []int{11}
}

var File_v1_app_proto protoreflect.FileDescriptor

var file_v1_app_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x70, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x50, 0x73, 0x6b, 0x22, 0x73, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x68,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70,
	0x76, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x22, 0x13,
	0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x0f,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x1a, 0x53, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x22, 0x43, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x53, 0x0a, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x73, 0x6b, 0x22, 0x15, 0x0a, 0x13, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x73, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x44, 0x48, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8a, 0x04, 0x0a, 0x09,
	0x41, 0x70, 0x70, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x34, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x48, 0x54, 0x12, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x44, 0x48, 0x54, 0x12, 0x13, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x44, 0x48, 0x54, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02,
	0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_app_proto_rawDescOnce sync.Once
	file_v1_app_proto_rawDescData = file_v1_app_proto_rawDesc
)

func file_v1_app_proto_rawDescGZIP() []byte {
	file_v1_app_proto_rawDescOnce.Do(func() {
		file_v1_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_app_proto_rawDescData)
	})
	return file_v1_app_proto_rawDescData
}

var file_v1_app_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_app_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_v1_app_proto_goTypes = []interface{}{
	(StatusResponse_ConnectionStatus)(0), // 0: v1.StatusResponse.ConnectionStatus
	(*ConnectRequest)(nil),               // 1: v1.ConnectRequest
	(*ConnectResponse)(nil),              // 2: v1.ConnectResponse
	(*DisconnectRequest)(nil),            // 3: v1.DisconnectRequest
	(*DisconnectResponse)(nil),           // 4: v1.DisconnectResponse
	(*MetricsRequest)(nil),               // 5: v1.MetricsRequest
	(*MetricsResponse)(nil),              // 6: v1.MetricsResponse
	(*StatusRequest)(nil),                // 7: v1.StatusRequest
	(*StatusResponse)(nil),               // 8: v1.StatusResponse
	(*AnnounceDHTRequest)(nil),           // 9: v1.AnnounceDHTRequest
	(*AnnounceDHTResponse)(nil),          // 10: v1.AnnounceDHTResponse
	(*LeaveDHTRequest)(nil),              // 11: v1.LeaveDHTRequest
	(*LeaveDHTResponse)(nil),             // 12: v1.LeaveDHTResponse
	nil,                                  // 13: v1.MetricsResponse.InterfacesEntry
	(*structpb.Struct)(nil),              // 14: google.protobuf.Struct
	(*MeshNode)(nil),                     // 15: v1.MeshNode
	(*InterfaceMetrics)(nil),             // 16: v1.InterfaceMetrics
	(*QueryRequest)(nil),                 // 17: v1.QueryRequest
	(*SubscribeRequest)(nil),             // 18: v1.SubscribeRequest
	(*PublishRequest)(nil),               // 19: v1.PublishRequest
	(*QueryResponse)(nil),                // 20: v1.QueryResponse
	(*SubscriptionEvent)(nil),            // 21: v1.SubscriptionEvent
	(*PublishResponse)(nil),              // 22: v1.PublishResponse
}
var file_v1_app_proto_depIdxs = []int32{
	14, // 0: v1.ConnectRequest.config:type_name -> google.protobuf.Struct
	13, // 1: v1.MetricsResponse.interfaces:type_name -> v1.MetricsResponse.InterfacesEntry
	0,  // 2: v1.StatusResponse.connection_status:type_name -> v1.StatusResponse.ConnectionStatus
	15, // 3: v1.StatusResponse.node:type_name -> v1.MeshNode
	16, // 4: v1.MetricsResponse.InterfacesEntry.value:type_name -> v1.InterfaceMetrics
	1,  // 5: v1.AppDaemon.Connect:input_type -> v1.ConnectRequest
	3,  // 6: v1.AppDaemon.Disconnect:input_type -> v1.DisconnectRequest
	17, // 7: v1.AppDaemon.Query:input_type -> v1.QueryRequest
	5,  // 8: v1.AppDaemon.Metrics:input_type -> v1.MetricsRequest
	7,  // 9: v1.AppDaemon.Status:input_type -> v1.StatusRequest
	18, // 10: v1.AppDaemon.Subscribe:input_type -> v1.SubscribeRequest
	19, // 11: v1.AppDaemon.Publish:input_type -> v1.PublishRequest
	9,  // 12: v1.AppDaemon.AnnounceDHT:input_type -> v1.AnnounceDHTRequest
	11, // 13: v1.AppDaemon.LeaveDHT:input_type -> v1.LeaveDHTRequest
	2,  // 14: v1.AppDaemon.Connect:output_type -> v1.ConnectResponse
	4,  // 15: v1.AppDaemon.Disconnect:output_type -> v1.DisconnectResponse
	20, // 16: v1.AppDaemon.Query:output_type -> v1.QueryResponse
	6,  // 17: v1.AppDaemon.Metrics:output_type -> v1.MetricsResponse
	8,  // 18: v1.AppDaemon.Status:output_type -> v1.StatusResponse
	21, // 19: v1.AppDaemon.Subscribe:output_type -> v1.SubscriptionEvent
	22, // 20: v1.AppDaemon.Publish:output_type -> v1.PublishResponse
	10, // 21: v1.AppDaemon.AnnounceDHT:output_type -> v1.AnnounceDHTResponse
	12, // 22: v1.AppDaemon.LeaveDHT:output_type -> v1.LeaveDHTResponse
	14, // [14:23] is the sub-list for method output_type
	5,  // [5:14] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_app_proto_init() }
func file_v1_app_proto_init() {
	if File_v1_app_proto != nil {
		return
	}
	file_v1_node_proto_init()
	file_v1_storage_proto_init()
	file_v1_mesh_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_v1_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_v1_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_v1_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
		file_v1_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_v1_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsResponse); i {
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
		file_v1_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_v1_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_v1_app_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceDHTRequest); i {
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
		file_v1_app_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceDHTResponse); i {
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
		file_v1_app_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveDHTRequest); i {
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
		file_v1_app_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveDHTResponse); i {
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
			RawDescriptor: file_v1_app_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_app_proto_goTypes,
		DependencyIndexes: file_v1_app_proto_depIdxs,
		EnumInfos:         file_v1_app_proto_enumTypes,
		MessageInfos:      file_v1_app_proto_msgTypes,
	}.Build()
	File_v1_app_proto = out.File
	file_v1_app_proto_rawDesc = nil
	file_v1_app_proto_goTypes = nil
	file_v1_app_proto_depIdxs = nil
}
