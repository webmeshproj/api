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
// source: v1/network_acls.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ACLAction is the action to take when a request matches an ACL.
type ACLAction int32

const (
	// ACTION_UNKNOWN is the default action for ACLs. It is synonymous with ACTION_DENY.
	ACLAction_ACTION_UNKNOWN ACLAction = 0
	// ACTION_ACCEPT allows the request to proceed.
	ACLAction_ACTION_ACCEPT ACLAction = 1
	// ACTION_DENY denies the request.
	ACLAction_ACTION_DENY ACLAction = 2
)

// Enum value maps for ACLAction.
var (
	ACLAction_name = map[int32]string{
		0: "ACTION_UNKNOWN",
		1: "ACTION_ACCEPT",
		2: "ACTION_DENY",
	}
	ACLAction_value = map[string]int32{
		"ACTION_UNKNOWN": 0,
		"ACTION_ACCEPT":  1,
		"ACTION_DENY":    2,
	}
)

func (x ACLAction) Enum() *ACLAction {
	p := new(ACLAction)
	*p = x
	return p
}

func (x ACLAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACLAction) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_network_acls_proto_enumTypes[0].Descriptor()
}

func (ACLAction) Type() protoreflect.EnumType {
	return &file_v1_network_acls_proto_enumTypes[0]
}

func (x ACLAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACLAction.Descriptor instead.
func (ACLAction) EnumDescriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{0}
}

// NetworkACL is a network ACL.
type NetworkACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the ACL.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// priority is the priority of the ACL. ACLs with higher priority are evaluated first.
	Priority int32 `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	// action is the action to take when a request matches the ACL.
	Action ACLAction `protobuf:"varint,3,opt,name=action,proto3,enum=v1.ACLAction" json:"action,omitempty"`
	// source_nodes is a list of source nodes to match against. If empty, all nodes are matched. Groups
	// can be specified with the prefix "group:".
	SourceNodes []string `protobuf:"bytes,4,rep,name=source_nodes,json=sourceNodes,proto3" json:"source_nodes,omitempty"`
	// destination_nodes is a list of destination nodes to match against. If empty, all nodes are matched.
	DestinationNodes []string `protobuf:"bytes,5,rep,name=destination_nodes,json=destinationNodes,proto3" json:"destination_nodes,omitempty"`
	// source_cidrs is a list of source CIDRs to match against. If empty, all CIDRs are matched.
	SourceCidrs []string `protobuf:"bytes,6,rep,name=source_cidrs,json=sourceCidrs,proto3" json:"source_cidrs,omitempty"`
	// destination_cidrs is a list of destination CIDRs to match against. If empty, all CIDRs are matched.
	DestinationCidrs []string `protobuf:"bytes,7,rep,name=destination_cidrs,json=destinationCidrs,proto3" json:"destination_cidrs,omitempty"`
	// protocols is a list of protocols to match against. If empty, all protocols are matched.
	// Protocols can be specified by name or number.
	Protocols []string `protobuf:"bytes,8,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// ports is a list of ports to match against. If empty, all ports are matched.
	Ports []uint32 `protobuf:"varint,9,rep,packed,name=ports,proto3" json:"ports,omitempty"`
}

func (x *NetworkACL) Reset() {
	*x = NetworkACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_network_acls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkACL) ProtoMessage() {}

func (x *NetworkACL) ProtoReflect() protoreflect.Message {
	mi := &file_v1_network_acls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkACL.ProtoReflect.Descriptor instead.
func (*NetworkACL) Descriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkACL) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkACL) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NetworkACL) GetAction() ACLAction {
	if x != nil {
		return x.Action
	}
	return ACLAction_ACTION_UNKNOWN
}

func (x *NetworkACL) GetSourceNodes() []string {
	if x != nil {
		return x.SourceNodes
	}
	return nil
}

func (x *NetworkACL) GetDestinationNodes() []string {
	if x != nil {
		return x.DestinationNodes
	}
	return nil
}

func (x *NetworkACL) GetSourceCidrs() []string {
	if x != nil {
		return x.SourceCidrs
	}
	return nil
}

func (x *NetworkACL) GetDestinationCidrs() []string {
	if x != nil {
		return x.DestinationCidrs
	}
	return nil
}

func (x *NetworkACL) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *NetworkACL) GetPorts() []uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

// NetworkACLs is a list of network ACLs.
type NetworkACLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items is the list of network ACLs.
	Items []*NetworkACL `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NetworkACLs) Reset() {
	*x = NetworkACLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_network_acls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkACLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkACLs) ProtoMessage() {}

func (x *NetworkACLs) ProtoReflect() protoreflect.Message {
	mi := &file_v1_network_acls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkACLs.ProtoReflect.Descriptor instead.
func (*NetworkACLs) Descriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkACLs) GetItems() []*NetworkACL {
	if x != nil {
		return x.Items
	}
	return nil
}

// Route is a route that is broadcasted by one or more nodes.
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the route.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// node is the node that broadcasts the route. A group can be specified with the prefix "group:".
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	// destination_cidrs are the destination CIDRs of the route.
	DestinationCidrs []string `protobuf:"bytes,3,rep,name=destination_cidrs,json=destinationCidrs,proto3" json:"destination_cidrs,omitempty"`
	// next_hop_node is an optional node that is used as the next hop for the route.
	NextHopNode string `protobuf:"bytes,4,opt,name=next_hop_node,json=nextHopNode,proto3" json:"next_hop_node,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_network_acls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_v1_network_acls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Route) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *Route) GetDestinationCidrs() []string {
	if x != nil {
		return x.DestinationCidrs
	}
	return nil
}

func (x *Route) GetNextHopNode() string {
	if x != nil {
		return x.NextHopNode
	}
	return ""
}

// Routes is a list of routes.
type Routes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items is the list of routes.
	Items []*Route `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Routes) Reset() {
	*x = Routes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_network_acls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routes) ProtoMessage() {}

func (x *Routes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_network_acls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routes.ProtoReflect.Descriptor instead.
func (*Routes) Descriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{3}
}

func (x *Routes) GetItems() []*Route {
	if x != nil {
		return x.Items
	}
	return nil
}

// NetworkAction is an action that can be performed on a network resource. It is used
// by implementations to evaluate network ACLs.
type NetworkAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// src_node is the source node of the action.
	SrcNode string `protobuf:"bytes,1,opt,name=src_node,json=srcNode,proto3" json:"src_node,omitempty"`
	// src_cidr is the source CIDR of the action.
	SrcCidr string `protobuf:"bytes,2,opt,name=src_cidr,json=srcCidr,proto3" json:"src_cidr,omitempty"`
	// dst_node is the destination node of the action.
	DstNode string `protobuf:"bytes,3,opt,name=dst_node,json=dstNode,proto3" json:"dst_node,omitempty"`
	// dst_cidr is the destination CIDR of the action.
	DstCidr string `protobuf:"bytes,4,opt,name=dst_cidr,json=dstCidr,proto3" json:"dst_cidr,omitempty"`
	// protocol is the protocol of the action.
	Protocol string `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// port is the port of the action.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *NetworkAction) Reset() {
	*x = NetworkAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_network_acls_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAction) ProtoMessage() {}

func (x *NetworkAction) ProtoReflect() protoreflect.Message {
	mi := &file_v1_network_acls_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAction.ProtoReflect.Descriptor instead.
func (*NetworkAction) Descriptor() ([]byte, []int) {
	return file_v1_network_acls_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkAction) GetSrcNode() string {
	if x != nil {
		return x.SrcNode
	}
	return ""
}

func (x *NetworkAction) GetSrcCidr() string {
	if x != nil {
		return x.SrcCidr
	}
	return ""
}

func (x *NetworkAction) GetDstNode() string {
	if x != nil {
		return x.DstNode
	}
	return ""
}

func (x *NetworkAction) GetDstCidr() string {
	if x != nil {
		return x.DstCidr
	}
	return ""
}

func (x *NetworkAction) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *NetworkAction) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_v1_network_acls_proto protoreflect.FileDescriptor

var file_v1_network_acls_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x61, 0x63, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xb7, 0x02, 0x0a, 0x0a,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x43, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x43, 0x4c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x69, 0x64, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x64, 0x72,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x43, 0x4c, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x43, 0x4c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x05, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x64, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x64, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x0a,
	0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72,
	0x63, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72,
	0x63, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x69, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x43, 0x69, 0x64, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64,
	0x73, 0x74, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x73, 0x74, 0x43, 0x69, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x2a, 0x43, 0x0a, 0x09, 0x41, 0x43, 0x4c, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02, 0x42, 0x6c, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63,
	0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x72, 0x6f,
	0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02,
	0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_network_acls_proto_rawDescOnce sync.Once
	file_v1_network_acls_proto_rawDescData = file_v1_network_acls_proto_rawDesc
)

func file_v1_network_acls_proto_rawDescGZIP() []byte {
	file_v1_network_acls_proto_rawDescOnce.Do(func() {
		file_v1_network_acls_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_network_acls_proto_rawDescData)
	})
	return file_v1_network_acls_proto_rawDescData
}

var file_v1_network_acls_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_network_acls_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_network_acls_proto_goTypes = []interface{}{
	(ACLAction)(0),        // 0: v1.ACLAction
	(*NetworkACL)(nil),    // 1: v1.NetworkACL
	(*NetworkACLs)(nil),   // 2: v1.NetworkACLs
	(*Route)(nil),         // 3: v1.Route
	(*Routes)(nil),        // 4: v1.Routes
	(*NetworkAction)(nil), // 5: v1.NetworkAction
}
var file_v1_network_acls_proto_depIdxs = []int32{
	0, // 0: v1.NetworkACL.action:type_name -> v1.ACLAction
	1, // 1: v1.NetworkACLs.items:type_name -> v1.NetworkACL
	3, // 2: v1.Routes.items:type_name -> v1.Route
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_network_acls_proto_init() }
func file_v1_network_acls_proto_init() {
	if File_v1_network_acls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_network_acls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkACL); i {
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
		file_v1_network_acls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkACLs); i {
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
		file_v1_network_acls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_v1_network_acls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routes); i {
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
		file_v1_network_acls_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAction); i {
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
			RawDescriptor: file_v1_network_acls_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_network_acls_proto_goTypes,
		DependencyIndexes: file_v1_network_acls_proto_depIdxs,
		EnumInfos:         file_v1_network_acls_proto_enumTypes,
		MessageInfos:      file_v1_network_acls_proto_msgTypes,
	}.Build()
	File_v1_network_acls_proto = out.File
	file_v1_network_acls_proto_rawDesc = nil
	file_v1_network_acls_proto_goTypes = nil
	file_v1_network_acls_proto_depIdxs = nil
}
