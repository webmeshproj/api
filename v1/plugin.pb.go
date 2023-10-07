//
//Copyright 2023 Avi Zimmerman <avi.zimmerman@gmail.com>
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
// source: v1/plugin.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// PluginCapability is the capabilities of a plugin.
type PluginInfo_PluginCapability int32

const (
	// UNKNOWN is the default value of PluginCapability.
	PluginInfo_UNKNOWN PluginInfo_PluginCapability = 0
	// STORAGE_PROVIDER indicates that the plugin can provide storage and underlying consistency.
	PluginInfo_STORAGE_PROVIDER PluginInfo_PluginCapability = 1
	// AUTH indicates that the plugin is an auth plugin.
	PluginInfo_AUTH PluginInfo_PluginCapability = 2
	// WATCH indicates that the plugin wants to receive watch events.
	PluginInfo_WATCH PluginInfo_PluginCapability = 3
	// IPAMV4 indicates that the plugin is an IPv4 IPAM plugin.
	PluginInfo_IPAMV4 PluginInfo_PluginCapability = 4
	// STORAGE_QUERIER indicates a plugin that wants to interact with storage.
	PluginInfo_STORAGE_QUERIER PluginInfo_PluginCapability = 5
)

// Enum value maps for PluginInfo_PluginCapability.
var (
	PluginInfo_PluginCapability_name = map[int32]string{
		0: "UNKNOWN",
		1: "STORAGE_PROVIDER",
		2: "AUTH",
		3: "WATCH",
		4: "IPAMV4",
		5: "STORAGE_QUERIER",
	}
	PluginInfo_PluginCapability_value = map[string]int32{
		"UNKNOWN":          0,
		"STORAGE_PROVIDER": 1,
		"AUTH":             2,
		"WATCH":            3,
		"IPAMV4":           4,
		"STORAGE_QUERIER":  5,
	}
)

func (x PluginInfo_PluginCapability) Enum() *PluginInfo_PluginCapability {
	p := new(PluginInfo_PluginCapability)
	*p = x
	return p
}

func (x PluginInfo_PluginCapability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginInfo_PluginCapability) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_plugin_proto_enumTypes[0].Descriptor()
}

func (PluginInfo_PluginCapability) Type() protoreflect.EnumType {
	return &file_v1_plugin_proto_enumTypes[0]
}

func (x PluginInfo_PluginCapability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginInfo_PluginCapability.Descriptor instead.
func (PluginInfo_PluginCapability) EnumDescriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{1, 0}
}

// WatchEvent is the type of a watch event.
type Event_WatchEvent int32

const (
	// UNKNOWN is the default value of WatchEvent.
	Event_UNKNOWN Event_WatchEvent = 0
	// NODE_JOIN indicates that a node has joined the cluster.
	Event_NODE_JOIN Event_WatchEvent = 1
	// NODE_LEAVE indicates that a node has left the cluster.
	Event_NODE_LEAVE Event_WatchEvent = 2
	// LEADER_CHANGE indicates that the leader of the cluster has changed.
	Event_LEADER_CHANGE Event_WatchEvent = 3
)

// Enum value maps for Event_WatchEvent.
var (
	Event_WatchEvent_name = map[int32]string{
		0: "UNKNOWN",
		1: "NODE_JOIN",
		2: "NODE_LEAVE",
		3: "LEADER_CHANGE",
	}
	Event_WatchEvent_value = map[string]int32{
		"UNKNOWN":       0,
		"NODE_JOIN":     1,
		"NODE_LEAVE":    2,
		"LEADER_CHANGE": 3,
	}
)

func (x Event_WatchEvent) Enum() *Event_WatchEvent {
	p := new(Event_WatchEvent)
	*p = x
	return p
}

func (x Event_WatchEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_WatchEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_plugin_proto_enumTypes[1].Descriptor()
}

func (Event_WatchEvent) Type() protoreflect.EnumType {
	return &file_v1_plugin_proto_enumTypes[1]
}

func (x Event_WatchEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_WatchEvent.Descriptor instead.
func (Event_WatchEvent) EnumDescriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{4, 0}
}

// PluginConfiguration is the message containing the configuration of a plugin.
type PluginConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Config is the configuration of the plugin.
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *PluginConfiguration) Reset() {
	*x = PluginConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfiguration) ProtoMessage() {}

func (x *PluginConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfiguration.ProtoReflect.Descriptor instead.
func (*PluginConfiguration) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginConfiguration) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

// PluginInfo is the information of a plugin.
type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version is the version of the plugin.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Description is the description of the plugin.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Capabilities is the capabilities of the plugin.
	Capabilities []PluginInfo_PluginCapability `protobuf:"varint,5,rep,packed,name=capabilities,proto3,enum=v1.PluginInfo_PluginCapability" json:"capabilities,omitempty"`
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PluginInfo) GetCapabilities() []PluginInfo_PluginCapability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

// AuthenticationRequest is the message containing an authentication request.
type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// headers are the headers of the request.
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// certificates are the DER encoded certificates of the request.
	Certificates [][]byte `protobuf:"bytes,2,rep,name=certificates,proto3" json:"certificates,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *AuthenticationRequest) GetCertificates() [][]byte {
	if x != nil {
		return x.Certificates
	}
	return nil
}

// AuthenticationResponse is the message containing an authentication response.
type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the id of the authenticated user.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Event is the message containing a watch event.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the type of the watch event.
	Type Event_WatchEvent `protobuf:"varint,1,opt,name=type,proto3,enum=v1.Event_WatchEvent" json:"type,omitempty"`
	// event is the data of the watch event.
	//
	// Types that are assignable to Event:
	//
	//	*Event_Node
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetType() Event_WatchEvent {
	if x != nil {
		return x.Type
	}
	return Event_UNKNOWN
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetNode() *MeshNode {
	if x, ok := x.GetEvent().(*Event_Node); ok {
		return x.Node
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Node struct {
	// node is the node that the event is about.
	Node *MeshNode `protobuf:"bytes,2,opt,name=node,proto3,oneof"`
}

func (*Event_Node) isEvent_Event() {}

// AllocateIPRequest is the message containing an IP allocation request.
type AllocateIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// node_id is the node that the IP should be allocated for.
	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	// subnet is the subnet that the IP should be allocated from.
	Subnet string `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *AllocateIPRequest) Reset() {
	*x = AllocateIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateIPRequest) ProtoMessage() {}

func (x *AllocateIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateIPRequest.ProtoReflect.Descriptor instead.
func (*AllocateIPRequest) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *AllocateIPRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *AllocateIPRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

// AllocatedIP is the message containing an allocated IP.
type AllocatedIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ip is the allocated IP. It should be returned in CIDR notation.
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *AllocatedIP) Reset() {
	*x = AllocatedIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocatedIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocatedIP) ProtoMessage() {}

func (x *AllocatedIP) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocatedIP.ProtoReflect.Descriptor instead.
func (*AllocatedIP) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *AllocatedIP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// ReleaseIPRequest is the message containing an IP release request.
type ReleaseIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// node_id is the node that the IP should be released for.
	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	// ip is the IP that should be released.
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ReleaseIPRequest) Reset() {
	*x = ReleaseIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseIPRequest) ProtoMessage() {}

func (x *ReleaseIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseIPRequest.ProtoReflect.Descriptor instead.
func (*ReleaseIPRequest) Descriptor() ([]byte, []int) {
	return file_v1_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *ReleaseIPRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *ReleaseIPRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_v1_plugin_proto protoreflect.FileDescriptor

var file_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x8e,
	0x02, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x6b, 0x0a, 0x10, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54,
	0x48, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x49, 0x50, 0x41, 0x4d, 0x56, 0x34, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54,
	0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x49, 0x45, 0x52, 0x10, 0x05, 0x22,
	0xb9, 0x01, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x16, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x68, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a,
	0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x44, 0x45,
	0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x44, 0x45, 0x5f,
	0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x41, 0x44, 0x45,
	0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x03, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x32, 0xb8, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x33,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0x55,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x47, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x3a, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x45, 0x6d, 0x69, 0x74, 0x12, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x32, 0x7d, 0x0a, 0x0a, 0x49, 0x50, 0x41, 0x4d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x34, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x50, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x50, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x32, 0x52, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x0d, 0x49, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x67, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73,
	0x68, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02,
	0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_plugin_proto_rawDescOnce sync.Once
	file_v1_plugin_proto_rawDescData = file_v1_plugin_proto_rawDesc
)

func file_v1_plugin_proto_rawDescGZIP() []byte {
	file_v1_plugin_proto_rawDescOnce.Do(func() {
		file_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_plugin_proto_rawDescData)
	})
	return file_v1_plugin_proto_rawDescData
}

var file_v1_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_plugin_proto_goTypes = []interface{}{
	(PluginInfo_PluginCapability)(0), // 0: v1.PluginInfo.PluginCapability
	(Event_WatchEvent)(0),            // 1: v1.Event.WatchEvent
	(*PluginConfiguration)(nil),      // 2: v1.PluginConfiguration
	(*PluginInfo)(nil),               // 3: v1.PluginInfo
	(*AuthenticationRequest)(nil),    // 4: v1.AuthenticationRequest
	(*AuthenticationResponse)(nil),   // 5: v1.AuthenticationResponse
	(*Event)(nil),                    // 6: v1.Event
	(*AllocateIPRequest)(nil),        // 7: v1.AllocateIPRequest
	(*AllocatedIP)(nil),              // 8: v1.AllocatedIP
	(*ReleaseIPRequest)(nil),         // 9: v1.ReleaseIPRequest
	nil,                              // 10: v1.AuthenticationRequest.HeadersEntry
	(*structpb.Struct)(nil),          // 11: google.protobuf.Struct
	(*MeshNode)(nil),                 // 12: v1.MeshNode
	(*emptypb.Empty)(nil),            // 13: google.protobuf.Empty
	(*QueryRequest)(nil),             // 14: v1.QueryRequest
	(*QueryResponse)(nil),            // 15: v1.QueryResponse
}
var file_v1_plugin_proto_depIdxs = []int32{
	11, // 0: v1.PluginConfiguration.config:type_name -> google.protobuf.Struct
	0,  // 1: v1.PluginInfo.capabilities:type_name -> v1.PluginInfo.PluginCapability
	10, // 2: v1.AuthenticationRequest.headers:type_name -> v1.AuthenticationRequest.HeadersEntry
	1,  // 3: v1.Event.type:type_name -> v1.Event.WatchEvent
	12, // 4: v1.Event.node:type_name -> v1.MeshNode
	13, // 5: v1.Plugin.GetInfo:input_type -> google.protobuf.Empty
	2,  // 6: v1.Plugin.Configure:input_type -> v1.PluginConfiguration
	13, // 7: v1.Plugin.Close:input_type -> google.protobuf.Empty
	4,  // 8: v1.AuthPlugin.Authenticate:input_type -> v1.AuthenticationRequest
	6,  // 9: v1.WatchPlugin.Emit:input_type -> v1.Event
	7,  // 10: v1.IPAMPlugin.Allocate:input_type -> v1.AllocateIPRequest
	9,  // 11: v1.IPAMPlugin.Release:input_type -> v1.ReleaseIPRequest
	14, // 12: v1.StorageQuerierPlugin.InjectQuerier:input_type -> v1.QueryRequest
	3,  // 13: v1.Plugin.GetInfo:output_type -> v1.PluginInfo
	13, // 14: v1.Plugin.Configure:output_type -> google.protobuf.Empty
	13, // 15: v1.Plugin.Close:output_type -> google.protobuf.Empty
	5,  // 16: v1.AuthPlugin.Authenticate:output_type -> v1.AuthenticationResponse
	13, // 17: v1.WatchPlugin.Emit:output_type -> google.protobuf.Empty
	8,  // 18: v1.IPAMPlugin.Allocate:output_type -> v1.AllocatedIP
	13, // 19: v1.IPAMPlugin.Release:output_type -> google.protobuf.Empty
	15, // 20: v1.StorageQuerierPlugin.InjectQuerier:output_type -> v1.QueryResponse
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_plugin_proto_init() }
func file_v1_plugin_proto_init() {
	if File_v1_plugin_proto != nil {
		return
	}
	file_v1_node_proto_init()
	file_v1_raft_proto_init()
	file_v1_storage_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfiguration); i {
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
		file_v1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
		file_v1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_v1_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
		file_v1_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_v1_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateIPRequest); i {
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
		file_v1_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocatedIP); i {
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
		file_v1_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseIPRequest); i {
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
	file_v1_plugin_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Event_Node)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_plugin_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_v1_plugin_proto_goTypes,
		DependencyIndexes: file_v1_plugin_proto_depIdxs,
		EnumInfos:         file_v1_plugin_proto_enumTypes,
		MessageInfos:      file_v1_plugin_proto_msgTypes,
	}.Build()
	File_v1_plugin_proto = out.File
	file_v1_plugin_proto_rawDesc = nil
	file_v1_plugin_proto_goTypes = nil
	file_v1_plugin_proto_depIdxs = nil
}
