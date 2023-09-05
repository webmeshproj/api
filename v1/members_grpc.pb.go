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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/members.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Membership_Join_FullMethodName                 = "/v1.Membership/Join"
	Membership_Update_FullMethodName               = "/v1.Membership/Update"
	Membership_Leave_FullMethodName                = "/v1.Membership/Leave"
	Membership_SubscribePeers_FullMethodName       = "/v1.Membership/SubscribePeers"
	Membership_Apply_FullMethodName                = "/v1.Membership/Apply"
	Membership_GetRaftConfiguration_FullMethodName = "/v1.Membership/GetRaftConfiguration"
)

// MembershipClient is the client API for Membership service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MembershipClient interface {
	// Join is used to join a node to the mesh.
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	// Update is used by a node to update its state in the mesh. The node will be updated
	// in the mesh and will be able to query the mesh state or vote in elections. Only
	// non-empty fields will be updated. It is almost semantically equivalent to a join request
	// with the same ID, but redefined to avoid confusion and to allow for expansion.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Leave is used to remove a node from the mesh. The node will be removed from the mesh
	// and will no longer be able to query the mesh state or vote in elections.
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
	// SubscribePeers subscribes to the peer configuration for the given node. The node
	// will receive updates to the peer configuration as it changes.
	SubscribePeers(ctx context.Context, in *SubscribePeersRequest, opts ...grpc.CallOption) (Membership_SubscribePeersClient, error)
	// Apply is used by voting nodes to request a log entry be applied to the state machine.
	// This is only available on the leader, and can only be called by nodes that are allowed
	// to vote.
	Apply(ctx context.Context, in *RaftLogEntry, opts ...grpc.CallOption) (*RaftApplyResponse, error)
	// GetRaftConfiguration returns the current Raft configuration.
	GetRaftConfiguration(ctx context.Context, in *RaftConfigurationRequest, opts ...grpc.CallOption) (*RaftConfigurationResponse, error)
}

type membershipClient struct {
	cc grpc.ClientConnInterface
}

func NewMembershipClient(cc grpc.ClientConnInterface) MembershipClient {
	return &membershipClient{cc}
}

func (c *membershipClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, Membership_Join_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, Membership_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, Membership_Leave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipClient) SubscribePeers(ctx context.Context, in *SubscribePeersRequest, opts ...grpc.CallOption) (Membership_SubscribePeersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Membership_ServiceDesc.Streams[0], Membership_SubscribePeers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &membershipSubscribePeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Membership_SubscribePeersClient interface {
	Recv() (*PeerConfigurations, error)
	grpc.ClientStream
}

type membershipSubscribePeersClient struct {
	grpc.ClientStream
}

func (x *membershipSubscribePeersClient) Recv() (*PeerConfigurations, error) {
	m := new(PeerConfigurations)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *membershipClient) Apply(ctx context.Context, in *RaftLogEntry, opts ...grpc.CallOption) (*RaftApplyResponse, error) {
	out := new(RaftApplyResponse)
	err := c.cc.Invoke(ctx, Membership_Apply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipClient) GetRaftConfiguration(ctx context.Context, in *RaftConfigurationRequest, opts ...grpc.CallOption) (*RaftConfigurationResponse, error) {
	out := new(RaftConfigurationResponse)
	err := c.cc.Invoke(ctx, Membership_GetRaftConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembershipServer is the server API for Membership service.
// All implementations must embed UnimplementedMembershipServer
// for forward compatibility
type MembershipServer interface {
	// Join is used to join a node to the mesh.
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	// Update is used by a node to update its state in the mesh. The node will be updated
	// in the mesh and will be able to query the mesh state or vote in elections. Only
	// non-empty fields will be updated. It is almost semantically equivalent to a join request
	// with the same ID, but redefined to avoid confusion and to allow for expansion.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Leave is used to remove a node from the mesh. The node will be removed from the mesh
	// and will no longer be able to query the mesh state or vote in elections.
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
	// SubscribePeers subscribes to the peer configuration for the given node. The node
	// will receive updates to the peer configuration as it changes.
	SubscribePeers(*SubscribePeersRequest, Membership_SubscribePeersServer) error
	// Apply is used by voting nodes to request a log entry be applied to the state machine.
	// This is only available on the leader, and can only be called by nodes that are allowed
	// to vote.
	Apply(context.Context, *RaftLogEntry) (*RaftApplyResponse, error)
	// GetRaftConfiguration returns the current Raft configuration.
	GetRaftConfiguration(context.Context, *RaftConfigurationRequest) (*RaftConfigurationResponse, error)
	mustEmbedUnimplementedMembershipServer()
}

// UnimplementedMembershipServer must be embedded to have forward compatible implementations.
type UnimplementedMembershipServer struct {
}

func (UnimplementedMembershipServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedMembershipServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMembershipServer) Leave(context.Context, *LeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedMembershipServer) SubscribePeers(*SubscribePeersRequest, Membership_SubscribePeersServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribePeers not implemented")
}
func (UnimplementedMembershipServer) Apply(context.Context, *RaftLogEntry) (*RaftApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedMembershipServer) GetRaftConfiguration(context.Context, *RaftConfigurationRequest) (*RaftConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRaftConfiguration not implemented")
}
func (UnimplementedMembershipServer) mustEmbedUnimplementedMembershipServer() {}

// UnsafeMembershipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MembershipServer will
// result in compilation errors.
type UnsafeMembershipServer interface {
	mustEmbedUnimplementedMembershipServer()
}

func RegisterMembershipServer(s grpc.ServiceRegistrar, srv MembershipServer) {
	s.RegisterService(&Membership_ServiceDesc, srv)
}

func _Membership_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Membership_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Membership_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Membership_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Membership_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Membership_Leave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Membership_SubscribePeers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribePeersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MembershipServer).SubscribePeers(m, &membershipSubscribePeersServer{stream})
}

type Membership_SubscribePeersServer interface {
	Send(*PeerConfigurations) error
	grpc.ServerStream
}

type membershipSubscribePeersServer struct {
	grpc.ServerStream
}

func (x *membershipSubscribePeersServer) Send(m *PeerConfigurations) error {
	return x.ServerStream.SendMsg(m)
}

func _Membership_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftLogEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Membership_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Apply(ctx, req.(*RaftLogEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Membership_GetRaftConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).GetRaftConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Membership_GetRaftConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).GetRaftConfiguration(ctx, req.(*RaftConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Membership_ServiceDesc is the grpc.ServiceDesc for Membership service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Membership_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Membership",
	HandlerType: (*MembershipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Membership_Join_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Membership_Update_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Membership_Leave_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Membership_Apply_Handler,
		},
		{
			MethodName: "GetRaftConfiguration",
			Handler:    _Membership_GetRaftConfiguration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePeers",
			Handler:       _Membership_SubscribePeers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/members.proto",
}