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
// source: v1/webrtc.proto

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
	WebRTC_StartDataChannel_FullMethodName = "/v1.WebRTC/StartDataChannel"
)

// WebRTCClient is the client API for WebRTC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebRTCClient interface {
	// StartDataChannel requests a new WebRTC connection to a node.
	// The client speaks first with the request containing the node ID
	// and where forwarded packets should be sent. The server responds
	// with an offer and STUN servers to be used to establish a WebRTC connection.
	// The client should then respond with an answer to the offer that matches the
	// spec of the DataChannel.CHANNELS enum value.
	//
	// After the offer is accepted, the stream can be used to exchange ICE candidates
	// to speed up the connection process.
	StartDataChannel(ctx context.Context, opts ...grpc.CallOption) (WebRTC_StartDataChannelClient, error)
}

type webRTCClient struct {
	cc grpc.ClientConnInterface
}

func NewWebRTCClient(cc grpc.ClientConnInterface) WebRTCClient {
	return &webRTCClient{cc}
}

func (c *webRTCClient) StartDataChannel(ctx context.Context, opts ...grpc.CallOption) (WebRTC_StartDataChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &WebRTC_ServiceDesc.Streams[0], WebRTC_StartDataChannel_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &webRTCStartDataChannelClient{stream}
	return x, nil
}

type WebRTC_StartDataChannelClient interface {
	Send(*StartDataChannelRequest) error
	Recv() (*DataChannelOffer, error)
	grpc.ClientStream
}

type webRTCStartDataChannelClient struct {
	grpc.ClientStream
}

func (x *webRTCStartDataChannelClient) Send(m *StartDataChannelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *webRTCStartDataChannelClient) Recv() (*DataChannelOffer, error) {
	m := new(DataChannelOffer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebRTCServer is the server API for WebRTC service.
// All implementations must embed UnimplementedWebRTCServer
// for forward compatibility
type WebRTCServer interface {
	// StartDataChannel requests a new WebRTC connection to a node.
	// The client speaks first with the request containing the node ID
	// and where forwarded packets should be sent. The server responds
	// with an offer and STUN servers to be used to establish a WebRTC connection.
	// The client should then respond with an answer to the offer that matches the
	// spec of the DataChannel.CHANNELS enum value.
	//
	// After the offer is accepted, the stream can be used to exchange ICE candidates
	// to speed up the connection process.
	StartDataChannel(WebRTC_StartDataChannelServer) error
	mustEmbedUnimplementedWebRTCServer()
}

// UnimplementedWebRTCServer must be embedded to have forward compatible implementations.
type UnimplementedWebRTCServer struct {
}

func (UnimplementedWebRTCServer) StartDataChannel(WebRTC_StartDataChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method StartDataChannel not implemented")
}
func (UnimplementedWebRTCServer) mustEmbedUnimplementedWebRTCServer() {}

// UnsafeWebRTCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebRTCServer will
// result in compilation errors.
type UnsafeWebRTCServer interface {
	mustEmbedUnimplementedWebRTCServer()
}

func RegisterWebRTCServer(s grpc.ServiceRegistrar, srv WebRTCServer) {
	s.RegisterService(&WebRTC_ServiceDesc, srv)
}

func _WebRTC_StartDataChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WebRTCServer).StartDataChannel(&webRTCStartDataChannelServer{stream})
}

type WebRTC_StartDataChannelServer interface {
	Send(*DataChannelOffer) error
	Recv() (*StartDataChannelRequest, error)
	grpc.ServerStream
}

type webRTCStartDataChannelServer struct {
	grpc.ServerStream
}

func (x *webRTCStartDataChannelServer) Send(m *DataChannelOffer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *webRTCStartDataChannelServer) Recv() (*StartDataChannelRequest, error) {
	m := new(StartDataChannelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebRTC_ServiceDesc is the grpc.ServiceDesc for WebRTC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebRTC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WebRTC",
	HandlerType: (*WebRTCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartDataChannel",
			Handler:       _WebRTC_StartDataChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/webrtc.proto",
}
