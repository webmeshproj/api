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
// source: v1/node_messages.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// DataChannel are the data channels used when communicating over ICE
// with a node.
type DataChannel int32

const (
	// CHANNELS is the data channel used for negotiating new channels.
	// This is the first channel that is opened. The ID of the channel
	// should be 0.
	DataChannel_CHANNELS DataChannel = 0
	// CONNECTIONS is the data channel used for negotiating new connections.
	// This is a channel that is opened for each incoming connection from a
	// client. The ID should start at 0 and be incremented for each new connection.
	DataChannel_CONNECTIONS DataChannel = 1
)

// Enum value maps for DataChannel.
var (
	DataChannel_name = map[int32]string{
		0: "CHANNELS",
		1: "CONNECTIONS",
	}
	DataChannel_value = map[string]int32{
		"CHANNELS":    0,
		"CONNECTIONS": 1,
	}
)

func (x DataChannel) Enum() *DataChannel {
	p := new(DataChannel)
	*p = x
	return p
}

func (x DataChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_node_messages_proto_enumTypes[0].Descriptor()
}

func (DataChannel) Type() protoreflect.EnumType {
	return &file_v1_node_messages_proto_enumTypes[0]
}

func (x DataChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataChannel.Descriptor instead.
func (DataChannel) EnumDescriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{0}
}

// JoinRequest is a request to join the cluster.
type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// public_key is the public wireguard key of the node to broadcast to peers.
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// raft_port is the Raft listen port of the node.
	RaftPort int32 `protobuf:"varint,3,opt,name=raft_port,json=raftPort,proto3" json:"raft_port,omitempty"`
	// grpc_port is the gRPC listen port of the node.
	GrpcPort int32 `protobuf:"varint,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// primary_endpoint is a routable address for the node. If left unset,
	// the node is assumed to be behind a NAT and not directly accessible.
	PrimaryEndpoint string `protobuf:"bytes,6,opt,name=primary_endpoint,json=primaryEndpoint,proto3" json:"primary_endpoint,omitempty"`
	// wireguard_endpoints is a list of WireGuard endpoints for the node.
	WireguardEndpoints []string `protobuf:"bytes,7,rep,name=wireguard_endpoints,json=wireguardEndpoints,proto3" json:"wireguard_endpoints,omitempty"`
	// zone_awareness_id is the zone awareness ID of the node.
	ZoneAwarenessId string `protobuf:"bytes,8,opt,name=zone_awareness_id,json=zoneAwarenessId,proto3" json:"zone_awareness_id,omitempty"`
	// assign_ipv4 is whether an IPv4 address should be assigned to the node.
	AssignIpv4 bool `protobuf:"varint,9,opt,name=assign_ipv4,json=assignIpv4,proto3" json:"assign_ipv4,omitempty"`
	// prefer_raft_ipv6 is whether IPv6 should be preferred over IPv4 for raft communication.
	// This is only used if assign_ipv4 is true.
	PreferRaftIpv6 bool `protobuf:"varint,10,opt,name=prefer_raft_ipv6,json=preferRaftIpv6,proto3" json:"prefer_raft_ipv6,omitempty"`
	// as_voter is whether the node should receive a vote in elections.
	AsVoter bool `protobuf:"varint,11,opt,name=as_voter,json=asVoter,proto3" json:"as_voter,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JoinRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *JoinRequest) GetRaftPort() int32 {
	if x != nil {
		return x.RaftPort
	}
	return 0
}

func (x *JoinRequest) GetGrpcPort() int32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

func (x *JoinRequest) GetPrimaryEndpoint() string {
	if x != nil {
		return x.PrimaryEndpoint
	}
	return ""
}

func (x *JoinRequest) GetWireguardEndpoints() []string {
	if x != nil {
		return x.WireguardEndpoints
	}
	return nil
}

func (x *JoinRequest) GetZoneAwarenessId() string {
	if x != nil {
		return x.ZoneAwarenessId
	}
	return ""
}

func (x *JoinRequest) GetAssignIpv4() bool {
	if x != nil {
		return x.AssignIpv4
	}
	return false
}

func (x *JoinRequest) GetPreferRaftIpv6() bool {
	if x != nil {
		return x.PreferRaftIpv6
	}
	return false
}

func (x *JoinRequest) GetAsVoter() bool {
	if x != nil {
		return x.AsVoter
	}
	return false
}

// JoinResponse is a response to a join request.
type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address_ipv4 is the private IPv4 wireguard address of the node
	// in CIDR format representing the network. This is only set if
	// assign_ipv4 was set in the request or no network_ipv6 was provided.
	// The bits are set to the network bits of the Mesh IPv4 network.
	AddressIpv4 string `protobuf:"bytes,1,opt,name=address_ipv4,json=addressIpv4,proto3" json:"address_ipv4,omitempty"`
	// address_ipv6 is the IPv6 network assigned to the node.
	AddressIpv6 string `protobuf:"bytes,2,opt,name=address_ipv6,json=addressIpv6,proto3" json:"address_ipv6,omitempty"`
	// network_ipv6 is the IPv6 network of the Mesh.
	NetworkIpv6 string `protobuf:"bytes,3,opt,name=network_ipv6,json=networkIpv6,proto3" json:"network_ipv6,omitempty"`
	// peers is a list of wireguard peers to connect to.
	Peers []*WireGuardPeer `protobuf:"bytes,4,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponse) GetAddressIpv4() string {
	if x != nil {
		return x.AddressIpv4
	}
	return ""
}

func (x *JoinResponse) GetAddressIpv6() string {
	if x != nil {
		return x.AddressIpv6
	}
	return ""
}

func (x *JoinResponse) GetNetworkIpv6() string {
	if x != nil {
		return x.NetworkIpv6
	}
	return ""
}

func (x *JoinResponse) GetPeers() []*WireGuardPeer {
	if x != nil {
		return x.Peers
	}
	return nil
}

// WireGuardPeer is a peer in the Wireguard network.
type WireGuardPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the peer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// public_key is the public key of the peer.
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// primary_endpoint is the primary endpoint of the peer.
	PrimaryEndpoint string `protobuf:"bytes,3,opt,name=primary_endpoint,json=primaryEndpoint,proto3" json:"primary_endpoint,omitempty"`
	// wireguard_endpoints are the WireGuard endpoints for the peer, if applicable.
	WireguardEndpoints []string `protobuf:"bytes,4,rep,name=wireguard_endpoints,json=wireguardEndpoints,proto3" json:"wireguard_endpoints,omitempty"`
	// zone_awareness_id is the zone awareness ID of the peer.
	ZoneAwarenessId string `protobuf:"bytes,5,opt,name=zone_awareness_id,json=zoneAwarenessId,proto3" json:"zone_awareness_id,omitempty"`
	// address_ipv4 is the private IPv4 wireguard address of the peer.
	AddressIpv4 string `protobuf:"bytes,6,opt,name=address_ipv4,json=addressIpv4,proto3" json:"address_ipv4,omitempty"`
	// address_ipv6 is the private IPv6 wireguard address of the peer.
	AddressIpv6 string `protobuf:"bytes,7,opt,name=address_ipv6,json=addressIpv6,proto3" json:"address_ipv6,omitempty"`
	// allowed_ips is the list of allowed IPs for the peer.
	AllowedIps []string `protobuf:"bytes,8,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (x *WireGuardPeer) Reset() {
	*x = WireGuardPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireGuardPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireGuardPeer) ProtoMessage() {}

func (x *WireGuardPeer) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireGuardPeer.ProtoReflect.Descriptor instead.
func (*WireGuardPeer) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{2}
}

func (x *WireGuardPeer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WireGuardPeer) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *WireGuardPeer) GetPrimaryEndpoint() string {
	if x != nil {
		return x.PrimaryEndpoint
	}
	return ""
}

func (x *WireGuardPeer) GetWireguardEndpoints() []string {
	if x != nil {
		return x.WireguardEndpoints
	}
	return nil
}

func (x *WireGuardPeer) GetZoneAwarenessId() string {
	if x != nil {
		return x.ZoneAwarenessId
	}
	return ""
}

func (x *WireGuardPeer) GetAddressIpv4() string {
	if x != nil {
		return x.AddressIpv4
	}
	return ""
}

func (x *WireGuardPeer) GetAddressIpv6() string {
	if x != nil {
		return x.AddressIpv6
	}
	return ""
}

func (x *WireGuardPeer) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

// LeaveRequest is a request to leave the cluster.
type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetStatusRequest is a request to get the status of a node.
type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node. If unset, the status of the
	// local node is returned.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Status represents the status of a node.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// version is the version of the node.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// commit is the commit of the node.
	Commit string `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	// build_date is the build date of the node.
	BuildDate string `protobuf:"bytes,4,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	// uptime is the uptime of the node.
	Uptime string `protobuf:"bytes,5,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// started_at is the time the node started.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// features is the list of features currently enabled.
	Features []Feature `protobuf:"varint,7,rep,packed,name=features,proto3,enum=v1.Feature" json:"features,omitempty"`
	// wireguard_peers is the number of wireguard peers connected.
	WireguardPeers uint32 `protobuf:"varint,8,opt,name=wireguard_peers,json=wireguardPeers,proto3" json:"wireguard_peers,omitempty"`
	// cluster_status is the status of the node in the cluster.
	ClusterStatus ClusterStatus `protobuf:"varint,9,opt,name=cluster_status,json=clusterStatus,proto3,enum=v1.ClusterStatus" json:"cluster_status,omitempty"`
	// current_leader is the current leader of the cluster.
	CurrentLeader string `protobuf:"bytes,10,opt,name=current_leader,json=currentLeader,proto3" json:"current_leader,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Status) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Status) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *Status) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

func (x *Status) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *Status) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Status) GetFeatures() []Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Status) GetWireguardPeers() uint32 {
	if x != nil {
		return x.WireguardPeers
	}
	return 0
}

func (x *Status) GetClusterStatus() ClusterStatus {
	if x != nil {
		return x.ClusterStatus
	}
	return ClusterStatus_CLUSTER_STATUS_UNKNOWN
}

func (x *Status) GetCurrentLeader() string {
	if x != nil {
		return x.CurrentLeader
	}
	return ""
}

// DataChannelNegotiation is the message for communicating data channels to nodes.
type DataChannelNegotiation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proto is the protocol of the traffic.
	Proto string `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	// src is the address of the client that initiated the request.
	Src string `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	// dst is the destination address of the traffic.
	Dst string `protobuf:"bytes,3,opt,name=dst,proto3" json:"dst,omitempty"`
	// port is the destination port of the traffic.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// offer is the offer for the node to use as its local description.
	Offer string `protobuf:"bytes,5,opt,name=offer,proto3" json:"offer,omitempty"`
	// answer is the answer for the node to use as its remote description.
	Answer string `protobuf:"bytes,6,opt,name=answer,proto3" json:"answer,omitempty"`
	// candidate is an ICE candidate.
	Candidate string `protobuf:"bytes,7,opt,name=candidate,proto3" json:"candidate,omitempty"`
	// stun_servers is the list of STUN servers to use.
	StunServers []string `protobuf:"bytes,8,rep,name=stun_servers,json=stunServers,proto3" json:"stun_servers,omitempty"`
}

func (x *DataChannelNegotiation) Reset() {
	*x = DataChannelNegotiation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataChannelNegotiation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataChannelNegotiation) ProtoMessage() {}

func (x *DataChannelNegotiation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataChannelNegotiation.ProtoReflect.Descriptor instead.
func (*DataChannelNegotiation) Descriptor() ([]byte, []int) {
	return file_v1_node_messages_proto_rawDescGZIP(), []int{6}
}

func (x *DataChannelNegotiation) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *DataChannelNegotiation) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *DataChannelNegotiation) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *DataChannelNegotiation) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *DataChannelNegotiation) GetOffer() string {
	if x != nil {
		return x.Offer
	}
	return ""
}

func (x *DataChannelNegotiation) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *DataChannelNegotiation) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

func (x *DataChannelNegotiation) GetStunServers() []string {
	if x != nil {
		return x.StunServers
	}
	return nil
}

var File_v1_node_messages_proto protoreflect.FileDescriptor

var file_v1_node_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x02, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x77, 0x69, 0x72,
	0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x61, 0x77, 0x61, 0x72, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x7a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x61, 0x72, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x49, 0x70, 0x76, 0x34, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x49, 0x70, 0x76,
	0x36, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x22, 0xa0, 0x01, 0x0a,
	0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x76, 0x34,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x76, 0x36,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x70, 0x76, 0x36, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x70, 0x76, 0x36, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22,
	0xad, 0x02, 0x0a, 0x0d, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x77,
	0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x77, 0x61, 0x72, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x7a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x61,
	0x72, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x76, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x76, 0x36, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x22,
	0x1e, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xef, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77,
	0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xd5, 0x01, 0x0a, 0x16, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x75, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2a, 0x2c, 0x0a,
	0x0b, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x42, 0x6d, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x72,
	0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca,
	0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_node_messages_proto_rawDescOnce sync.Once
	file_v1_node_messages_proto_rawDescData = file_v1_node_messages_proto_rawDesc
)

func file_v1_node_messages_proto_rawDescGZIP() []byte {
	file_v1_node_messages_proto_rawDescOnce.Do(func() {
		file_v1_node_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_node_messages_proto_rawDescData)
	})
	return file_v1_node_messages_proto_rawDescData
}

var file_v1_node_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_node_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_node_messages_proto_goTypes = []interface{}{
	(DataChannel)(0),               // 0: v1.DataChannel
	(*JoinRequest)(nil),            // 1: v1.JoinRequest
	(*JoinResponse)(nil),           // 2: v1.JoinResponse
	(*WireGuardPeer)(nil),          // 3: v1.WireGuardPeer
	(*LeaveRequest)(nil),           // 4: v1.LeaveRequest
	(*GetStatusRequest)(nil),       // 5: v1.GetStatusRequest
	(*Status)(nil),                 // 6: v1.Status
	(*DataChannelNegotiation)(nil), // 7: v1.DataChannelNegotiation
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(Feature)(0),                   // 9: v1.Feature
	(ClusterStatus)(0),             // 10: v1.ClusterStatus
}
var file_v1_node_messages_proto_depIdxs = []int32{
	3,  // 0: v1.JoinResponse.peers:type_name -> v1.WireGuardPeer
	8,  // 1: v1.Status.started_at:type_name -> google.protobuf.Timestamp
	9,  // 2: v1.Status.features:type_name -> v1.Feature
	10, // 3: v1.Status.cluster_status:type_name -> v1.ClusterStatus
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v1_node_messages_proto_init() }
func file_v1_node_messages_proto_init() {
	if File_v1_node_messages_proto != nil {
		return
	}
	file_v1_mesh_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_node_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_v1_node_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_v1_node_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireGuardPeer); i {
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
		file_v1_node_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_v1_node_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusRequest); i {
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
		file_v1_node_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_v1_node_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataChannelNegotiation); i {
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
			RawDescriptor: file_v1_node_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_node_messages_proto_goTypes,
		DependencyIndexes: file_v1_node_messages_proto_depIdxs,
		EnumInfos:         file_v1_node_messages_proto_enumTypes,
		MessageInfos:      file_v1_node_messages_proto_msgTypes,
	}.Build()
	File_v1_node_messages_proto = out.File
	file_v1_node_messages_proto_rawDesc = nil
	file_v1_node_messages_proto_goTypes = nil
	file_v1_node_messages_proto_depIdxs = nil
}
