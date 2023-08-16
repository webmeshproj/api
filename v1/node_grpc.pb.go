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
// source: v1/node.proto

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
	Node_GetStatus_FullMethodName            = "/v1.Node/GetStatus"
	Node_Query_FullMethodName                = "/v1.Node/Query"
	Node_Publish_FullMethodName              = "/v1.Node/Publish"
	Node_Subscribe_FullMethodName            = "/v1.Node/Subscribe"
	Node_NegotiateDataChannel_FullMethodName = "/v1.Node/NegotiateDataChannel"
)

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	// GetStatus gets the status of a node in the cluster.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*Status, error)
	// Query is used to query the mesh for information.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Node_QueryClient, error)
	// Publish is used to publish events to the mesh database. A restricted set
	// of keys are allowed to be published to.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Subscribe is used by non-raft nodes to receive updates to the mesh state. This is only
	// available on nodes that are members of the raft cluster.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Node_SubscribeClient, error)
	// NegotiateDataChannel is used to negotiate a WebRTC connection between a webmesh client
	// and a node in the cluster. The handling server will send the target node the source address,
	// the destination for traffic, and STUN/TURN servers to use for the negotiation. The node
	// responds with an offer to be forwarded to the client. When the handler receives an answer
	// from the client, it forwards it to the node. Once the node receives the answer, the stream
	// can optionally be used to exchange ICE candidates.
	NegotiateDataChannel(ctx context.Context, opts ...grpc.CallOption) (Node_NegotiateDataChannelClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, Node_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Node_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], Node_Query_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_QueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type nodeQueryClient struct {
	grpc.ClientStream
}

func (x *nodeQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Node_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Node_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[1], Node_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_SubscribeClient interface {
	Recv() (*SubscriptionEvent, error)
	grpc.ClientStream
}

type nodeSubscribeClient struct {
	grpc.ClientStream
}

func (x *nodeSubscribeClient) Recv() (*SubscriptionEvent, error) {
	m := new(SubscriptionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) NegotiateDataChannel(ctx context.Context, opts ...grpc.CallOption) (Node_NegotiateDataChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[2], Node_NegotiateDataChannel_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeNegotiateDataChannelClient{stream}
	return x, nil
}

type Node_NegotiateDataChannelClient interface {
	Send(*DataChannelNegotiation) error
	Recv() (*DataChannelNegotiation, error)
	grpc.ClientStream
}

type nodeNegotiateDataChannelClient struct {
	grpc.ClientStream
}

func (x *nodeNegotiateDataChannelClient) Send(m *DataChannelNegotiation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeNegotiateDataChannelClient) Recv() (*DataChannelNegotiation, error) {
	m := new(DataChannelNegotiation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	// GetStatus gets the status of a node in the cluster.
	GetStatus(context.Context, *GetStatusRequest) (*Status, error)
	// Query is used to query the mesh for information.
	Query(*QueryRequest, Node_QueryServer) error
	// Publish is used to publish events to the mesh database. A restricted set
	// of keys are allowed to be published to.
	Publish(context.Context, *PublishRequest) (*emptypb.Empty, error)
	// Subscribe is used by non-raft nodes to receive updates to the mesh state. This is only
	// available on nodes that are members of the raft cluster.
	Subscribe(*SubscribeRequest, Node_SubscribeServer) error
	// NegotiateDataChannel is used to negotiate a WebRTC connection between a webmesh client
	// and a node in the cluster. The handling server will send the target node the source address,
	// the destination for traffic, and STUN/TURN servers to use for the negotiation. The node
	// responds with an offer to be forwarded to the client. When the handler receives an answer
	// from the client, it forwards it to the node. Once the node receives the answer, the stream
	// can optionally be used to exchange ICE candidates.
	NegotiateDataChannel(Node_NegotiateDataChannelServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) GetStatus(context.Context, *GetStatusRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedNodeServer) Query(*QueryRequest, Node_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedNodeServer) Publish(context.Context, *PublishRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedNodeServer) Subscribe(*SubscribeRequest, Node_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedNodeServer) NegotiateDataChannel(Node_NegotiateDataChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method NegotiateDataChannel not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).Query(m, &nodeQueryServer{stream})
}

type Node_QueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type nodeQueryServer struct {
	grpc.ServerStream
}

func (x *nodeQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).Subscribe(m, &nodeSubscribeServer{stream})
}

type Node_SubscribeServer interface {
	Send(*SubscriptionEvent) error
	grpc.ServerStream
}

type nodeSubscribeServer struct {
	grpc.ServerStream
}

func (x *nodeSubscribeServer) Send(m *SubscriptionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_NegotiateDataChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).NegotiateDataChannel(&nodeNegotiateDataChannelServer{stream})
}

type Node_NegotiateDataChannelServer interface {
	Send(*DataChannelNegotiation) error
	Recv() (*DataChannelNegotiation, error)
	grpc.ServerStream
}

type nodeNegotiateDataChannelServer struct {
	grpc.ServerStream
}

func (x *nodeNegotiateDataChannelServer) Send(m *DataChannelNegotiation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeNegotiateDataChannelServer) Recv() (*DataChannelNegotiation, error) {
	m := new(DataChannelNegotiation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Node_GetStatus_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Node_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _Node_Query_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Node_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NegotiateDataChannel",
			Handler:       _Node_NegotiateDataChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/node.proto",
}
