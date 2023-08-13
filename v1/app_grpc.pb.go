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
// source: v1/app.proto

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
	AppDaemon_Connect_FullMethodName       = "/v1.AppDaemon/Connect"
	AppDaemon_Disconnect_FullMethodName    = "/v1.AppDaemon/Disconnect"
	AppDaemon_StartCampfire_FullMethodName = "/v1.AppDaemon/StartCampfire"
	AppDaemon_Query_FullMethodName         = "/v1.AppDaemon/Query"
)

// AppDaemonClient is the client API for AppDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppDaemonClient interface {
	// Connect is used to establish a connection between the node and a
	// mesh. The provided struct is used to override any defaults configured
	// on the node.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// Disconnect is used to disconnect the node from the mesh.
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// StartCampfire is used to start a campfire.
	StartCampfire(ctx context.Context, in *StartCampfireRequest, opts ...grpc.CallOption) (*StartCampfireResponse, error)
	// Query is used to query the mesh for information.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (AppDaemon_QueryClient, error)
}

type appDaemonClient struct {
	cc grpc.ClientConnInterface
}

func NewAppDaemonClient(cc grpc.ClientConnInterface) AppDaemonClient {
	return &appDaemonClient{cc}
}

func (c *appDaemonClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Disconnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) StartCampfire(ctx context.Context, in *StartCampfireRequest, opts ...grpc.CallOption) (*StartCampfireResponse, error) {
	out := new(StartCampfireResponse)
	err := c.cc.Invoke(ctx, AppDaemon_StartCampfire_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (AppDaemon_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppDaemon_ServiceDesc.Streams[0], AppDaemon_Query_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &appDaemonQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppDaemon_QueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type appDaemonQueryClient struct {
	grpc.ClientStream
}

func (x *appDaemonQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppDaemonServer is the server API for AppDaemon service.
// All implementations must embed UnimplementedAppDaemonServer
// for forward compatibility
type AppDaemonServer interface {
	// Connect is used to establish a connection between the node and a
	// mesh. The provided struct is used to override any defaults configured
	// on the node.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Disconnect is used to disconnect the node from the mesh.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// StartCampfire is used to start a campfire.
	StartCampfire(context.Context, *StartCampfireRequest) (*StartCampfireResponse, error)
	// Query is used to query the mesh for information.
	Query(*QueryRequest, AppDaemon_QueryServer) error
	mustEmbedUnimplementedAppDaemonServer()
}

// UnimplementedAppDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedAppDaemonServer struct {
}

func (UnimplementedAppDaemonServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedAppDaemonServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedAppDaemonServer) StartCampfire(context.Context, *StartCampfireRequest) (*StartCampfireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCampfire not implemented")
}
func (UnimplementedAppDaemonServer) Query(*QueryRequest, AppDaemon_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedAppDaemonServer) mustEmbedUnimplementedAppDaemonServer() {}

// UnsafeAppDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppDaemonServer will
// result in compilation errors.
type UnsafeAppDaemonServer interface {
	mustEmbedUnimplementedAppDaemonServer()
}

func RegisterAppDaemonServer(s grpc.ServiceRegistrar, srv AppDaemonServer) {
	s.RegisterService(&AppDaemon_ServiceDesc, srv)
}

func _AppDaemon_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_StartCampfire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartCampfireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).StartCampfire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_StartCampfire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).StartCampfire(ctx, req.(*StartCampfireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppDaemonServer).Query(m, &appDaemonQueryServer{stream})
}

type AppDaemon_QueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type appDaemonQueryServer struct {
	grpc.ServerStream
}

func (x *appDaemonQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AppDaemon_ServiceDesc is the grpc.ServiceDesc for AppDaemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppDaemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AppDaemon",
	HandlerType: (*AppDaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _AppDaemon_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _AppDaemon_Disconnect_Handler,
		},
		{
			MethodName: "StartCampfire",
			Handler:    _AppDaemon_StartCampfire_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _AppDaemon_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/app.proto",
}
