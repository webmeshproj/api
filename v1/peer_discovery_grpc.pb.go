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
// source: v1/peer_discovery.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PeerDiscovery_ListPeers_FullMethodName = "/v1.PeerDiscovery/ListPeers"
)

// PeerDiscoveryClient is the client API for PeerDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerDiscoveryClient interface {
	// ListPeers returns a list of peers currently known to the mesh.
	ListPeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRaftPeersResponse, error)
}

type peerDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerDiscoveryClient(cc grpc.ClientConnInterface) PeerDiscoveryClient {
	return &peerDiscoveryClient{cc}
}

func (c *peerDiscoveryClient) ListPeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRaftPeersResponse, error) {
	out := new(ListRaftPeersResponse)
	err := c.cc.Invoke(ctx, PeerDiscovery_ListPeers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerDiscoveryServer is the server API for PeerDiscovery service.
// All implementations must embed UnimplementedPeerDiscoveryServer
// for forward compatibility
type PeerDiscoveryServer interface {
	// ListPeers returns a list of peers currently known to the mesh.
	ListPeers(context.Context, *emptypb.Empty) (*ListRaftPeersResponse, error)
	mustEmbedUnimplementedPeerDiscoveryServer()
}

// UnimplementedPeerDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedPeerDiscoveryServer struct {
}

func (UnimplementedPeerDiscoveryServer) ListPeers(context.Context, *emptypb.Empty) (*ListRaftPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeers not implemented")
}
func (UnimplementedPeerDiscoveryServer) mustEmbedUnimplementedPeerDiscoveryServer() {}

// UnsafePeerDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerDiscoveryServer will
// result in compilation errors.
type UnsafePeerDiscoveryServer interface {
	mustEmbedUnimplementedPeerDiscoveryServer()
}

func RegisterPeerDiscoveryServer(s grpc.ServiceRegistrar, srv PeerDiscoveryServer) {
	s.RegisterService(&PeerDiscovery_ServiceDesc, srv)
}

func _PeerDiscovery_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerDiscoveryServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerDiscovery_ListPeers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerDiscoveryServer).ListPeers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerDiscovery_ServiceDesc is the grpc.ServiceDesc for PeerDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PeerDiscovery",
	HandlerType: (*PeerDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPeers",
			Handler:    _PeerDiscovery_ListPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/peer_discovery.proto",
}
