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
// source: v1/storage_query.proto

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

// QueryCommand is the type of the query.
type QueryRequest_QueryCommand int32

const (
	// GET is the command to get a value.
	QueryRequest_GET QueryRequest_QueryCommand = 0
	// LIST is the command to list keys with an optional prefix.
	QueryRequest_LIST QueryRequest_QueryCommand = 1
)

// Enum value maps for QueryRequest_QueryCommand.
var (
	QueryRequest_QueryCommand_name = map[int32]string{
		0: "GET",
		1: "LIST",
	}
	QueryRequest_QueryCommand_value = map[string]int32{
		"GET":  0,
		"LIST": 1,
	}
)

func (x QueryRequest_QueryCommand) Enum() *QueryRequest_QueryCommand {
	p := new(QueryRequest_QueryCommand)
	*p = x
	return p
}

func (x QueryRequest_QueryCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryRequest_QueryCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_storage_query_proto_enumTypes[0].Descriptor()
}

func (QueryRequest_QueryCommand) Type() protoreflect.EnumType {
	return &file_v1_storage_query_proto_enumTypes[0]
}

func (x QueryRequest_QueryCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryRequest_QueryCommand.Descriptor instead.
func (QueryRequest_QueryCommand) EnumDescriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{1, 0}
}

// QueryType is the type of object being queried.
type QueryRequest_QueryType int32

const (
	// VALUE represents a raw value query at a supplied key.
	QueryRequest_VALUE QueryRequest_QueryType = 0
	// KEYS is the type for querying keys.
	QueryRequest_KEYS QueryRequest_QueryType = 1
	// PEERS is the type for querying peers.
	QueryRequest_PEERS QueryRequest_QueryType = 2
	// EDGES is the type for querying edges.
	QueryRequest_EDGES QueryRequest_QueryType = 3
	// ROUTES is the type for querying routes.
	QueryRequest_ROUTES QueryRequest_QueryType = 4
	// ACLS is the type for querying ACLs.
	QueryRequest_ACLS QueryRequest_QueryType = 5
	// ROLES is the type for querying roles.
	QueryRequest_ROLES QueryRequest_QueryType = 6
	// ROLEBINDINGS is the type for querying role bindings.
	QueryRequest_ROLEBINDINGS QueryRequest_QueryType = 7
	// GROUPS is the type for querying groups.
	QueryRequest_GROUPS QueryRequest_QueryType = 8
	// NETWORK_STATE is the type for querying network configuration.
	QueryRequest_NETWORK_STATE QueryRequest_QueryType = 9
	// RBAC_STATE is the type for querying RBAC configuration.
	// This will return a single item of true or false.
	QueryRequest_RBAC_STATE QueryRequest_QueryType = 10
)

// Enum value maps for QueryRequest_QueryType.
var (
	QueryRequest_QueryType_name = map[int32]string{
		0:  "VALUE",
		1:  "KEYS",
		2:  "PEERS",
		3:  "EDGES",
		4:  "ROUTES",
		5:  "ACLS",
		6:  "ROLES",
		7:  "ROLEBINDINGS",
		8:  "GROUPS",
		9:  "NETWORK_STATE",
		10: "RBAC_STATE",
	}
	QueryRequest_QueryType_value = map[string]int32{
		"VALUE":         0,
		"KEYS":          1,
		"PEERS":         2,
		"EDGES":         3,
		"ROUTES":        4,
		"ACLS":          5,
		"ROLES":         6,
		"ROLEBINDINGS":  7,
		"GROUPS":        8,
		"NETWORK_STATE": 9,
		"RBAC_STATE":    10,
	}
)

func (x QueryRequest_QueryType) Enum() *QueryRequest_QueryType {
	p := new(QueryRequest_QueryType)
	*p = x
	return p
}

func (x QueryRequest_QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryRequest_QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_storage_query_proto_enumTypes[1].Descriptor()
}

func (QueryRequest_QueryType) Type() protoreflect.EnumType {
	return &file_v1_storage_query_proto_enumTypes[1]
}

func (x QueryRequest_QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryRequest_QueryType.Descriptor instead.
func (QueryRequest_QueryType) EnumDescriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{1, 1}
}

// NetworkState represents the full network state as returned by
// a network state query.
type NetworkState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkV4 string `protobuf:"bytes,1,opt,name=networkV4,proto3" json:"networkV4,omitempty"`
	NetworkV6 string `protobuf:"bytes,2,opt,name=networkV6,proto3" json:"networkV6,omitempty"`
	Domain    string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *NetworkState) Reset() {
	*x = NetworkState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkState) ProtoMessage() {}

func (x *NetworkState) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkState.ProtoReflect.Descriptor instead.
func (*NetworkState) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkState) GetNetworkV4() string {
	if x != nil {
		return x.NetworkV4
	}
	return ""
}

func (x *NetworkState) GetNetworkV6() string {
	if x != nil {
		return x.NetworkV6
	}
	return ""
}

func (x *NetworkState) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// QueryRequest is sent by the application to the node to query the mesh for
// information.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// command is the command of the query.
	Command QueryRequest_QueryCommand `protobuf:"varint,1,opt,name=command,proto3,enum=v1.QueryRequest_QueryCommand" json:"command,omitempty"`
	// type is the type of the query.
	Type QueryRequest_QueryType `protobuf:"varint,2,opt,name=type,proto3,enum=v1.QueryRequest_QueryType" json:"type,omitempty"`
	// query is the string of the query. This follows the format of a label
	// selector and is only applicable for certain queries. For get queries
	// this will usually be an ID. For list queries this will usually be one
	// or more filters.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRequest) GetCommand() QueryRequest_QueryCommand {
	if x != nil {
		return x.Command
	}
	return QueryRequest_GET
}

func (x *QueryRequest) GetType() QueryRequest_QueryType {
	if x != nil {
		return x.Type
	}
	return QueryRequest_VALUE
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// QueryResponse is the message containing a mesh query result.
type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items contain the results of the query. These will be protobuf
	// json-encoded objects of the given query type.
	Items [][]byte `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// error is an error that happened during the query. This will always
	// be populated on errors, but single-flight queries will return
	// a coded error instead.
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResponse) GetItems() [][]byte {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *QueryResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// SubscribeRequest is sent by the application to the node to subscribe to
// events. This currently only supports database events.
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// prefix is the prefix of the events to subscribe to.
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRequest) GetPrefix() []byte {
	if x != nil {
		return x.Prefix
	}
	return nil
}

// SubscriptionEvent is a message containing a subscription event.
type SubscriptionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the key of the event.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value of the event. This will be the raw value of the key.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SubscriptionEvent) Reset() {
	*x = SubscriptionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionEvent) ProtoMessage() {}

func (x *SubscriptionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionEvent.ProtoReflect.Descriptor instead.
func (*SubscriptionEvent) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{4}
}

func (x *SubscriptionEvent) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SubscriptionEvent) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// PublishRequest is sent by the application to the node to publish events.
// This currently only supports database events.
type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the key of the event.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value of the event. This will be the raw value of the key.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// ttl is the time for the event to live in the database.
	Ttl *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{5}
}

func (x *PublishRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PublishRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PublishRequest) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// PublishResponse is the response to a publish request. This is currently
// empty.
type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_storage_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_storage_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_v1_storage_query_proto_rawDescGZIP(), []int{6}
}

var File_v1_storage_query_proto protoreflect.FileDescriptor

var file_v1_storage_query_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0c,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x56, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x56, 0x34, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x56, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x56, 0x36, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0xcb, 0x02, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x21, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53,
	0x54, 0x10, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4b, 0x45, 0x59, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x45, 0x45, 0x52, 0x53, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x44, 0x47, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x43, 0x4c, 0x53,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x4f, 0x4c, 0x45, 0x53, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x52, 0x4f, 0x4c, 0x45, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x07, 0x12,
	0x0a, 0x0a, 0x06, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x42, 0x41, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x0a, 0x22, 0x3b,
	0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x3b, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b,
	0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x11, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb9,
	0x01, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x6d, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x72, 0x6f,
	0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02,
	0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_storage_query_proto_rawDescOnce sync.Once
	file_v1_storage_query_proto_rawDescData = file_v1_storage_query_proto_rawDesc
)

func file_v1_storage_query_proto_rawDescGZIP() []byte {
	file_v1_storage_query_proto_rawDescOnce.Do(func() {
		file_v1_storage_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_storage_query_proto_rawDescData)
	})
	return file_v1_storage_query_proto_rawDescData
}

var file_v1_storage_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_storage_query_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_storage_query_proto_goTypes = []interface{}{
	(QueryRequest_QueryCommand)(0), // 0: v1.QueryRequest.QueryCommand
	(QueryRequest_QueryType)(0),    // 1: v1.QueryRequest.QueryType
	(*NetworkState)(nil),           // 2: v1.NetworkState
	(*QueryRequest)(nil),           // 3: v1.QueryRequest
	(*QueryResponse)(nil),          // 4: v1.QueryResponse
	(*SubscribeRequest)(nil),       // 5: v1.SubscribeRequest
	(*SubscriptionEvent)(nil),      // 6: v1.SubscriptionEvent
	(*PublishRequest)(nil),         // 7: v1.PublishRequest
	(*PublishResponse)(nil),        // 8: v1.PublishResponse
	(*durationpb.Duration)(nil),    // 9: google.protobuf.Duration
}
var file_v1_storage_query_proto_depIdxs = []int32{
	0, // 0: v1.QueryRequest.command:type_name -> v1.QueryRequest.QueryCommand
	1, // 1: v1.QueryRequest.type:type_name -> v1.QueryRequest.QueryType
	9, // 2: v1.PublishRequest.ttl:type_name -> google.protobuf.Duration
	3, // 3: v1.StorageQueryService.Query:input_type -> v1.QueryRequest
	7, // 4: v1.StorageQueryService.Publish:input_type -> v1.PublishRequest
	5, // 5: v1.StorageQueryService.Subscribe:input_type -> v1.SubscribeRequest
	4, // 6: v1.StorageQueryService.Query:output_type -> v1.QueryResponse
	8, // 7: v1.StorageQueryService.Publish:output_type -> v1.PublishResponse
	6, // 8: v1.StorageQueryService.Subscribe:output_type -> v1.SubscriptionEvent
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_storage_query_proto_init() }
func file_v1_storage_query_proto_init() {
	if File_v1_storage_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_storage_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkState); i {
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
		file_v1_storage_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_v1_storage_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_v1_storage_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_v1_storage_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionEvent); i {
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
		file_v1_storage_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_v1_storage_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
			RawDescriptor: file_v1_storage_query_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_storage_query_proto_goTypes,
		DependencyIndexes: file_v1_storage_query_proto_depIdxs,
		EnumInfos:         file_v1_storage_query_proto_enumTypes,
		MessageInfos:      file_v1_storage_query_proto_msgTypes,
	}.Build()
	File_v1_storage_query_proto = out.File
	file_v1_storage_query_proto_rawDesc = nil
	file_v1_storage_query_proto_goTypes = nil
	file_v1_storage_query_proto_depIdxs = nil
}
