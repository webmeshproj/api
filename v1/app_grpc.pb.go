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
	AppDaemon_Connect_FullMethodName     = "/v1.AppDaemon/Connect"
	AppDaemon_Disconnect_FullMethodName  = "/v1.AppDaemon/Disconnect"
	AppDaemon_Query_FullMethodName       = "/v1.AppDaemon/Query"
	AppDaemon_Metrics_FullMethodName     = "/v1.AppDaemon/Metrics"
	AppDaemon_Status_FullMethodName      = "/v1.AppDaemon/Status"
	AppDaemon_Subscribe_FullMethodName   = "/v1.AppDaemon/Subscribe"
	AppDaemon_Publish_FullMethodName     = "/v1.AppDaemon/Publish"
	AppDaemon_AnnounceDHT_FullMethodName = "/v1.AppDaemon/AnnounceDHT"
	AppDaemon_LeaveDHT_FullMethodName    = "/v1.AppDaemon/LeaveDHT"
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
	// Query is used to query the mesh for information.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Metrics is used to retrieve interface metrics from the node.
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
	// Status is used to retrieve the status of the node.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Subscribe is used to subscribe to events in the mesh database.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (AppDaemon_SubscribeClient, error)
	// Publish is used to publish events to the mesh database. A restricted set
	// of keys are allowed to be published to.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// AnnounceDHT is used to announce the node's presence on the Kademlia DHT
	// for other nodes to discover.
	AnnounceDHT(ctx context.Context, in *AnnounceDHTRequest, opts ...grpc.CallOption) (*AnnounceDHTResponse, error)
	// LeaveDHT is used to leave the Kademlia DHT.
	LeaveDHT(ctx context.Context, in *LeaveDHTRequest, opts ...grpc.CallOption) (*LeaveDHTResponse, error)
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

func (c *appDaemonClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Metrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (AppDaemon_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppDaemon_ServiceDesc.Streams[0], AppDaemon_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &appDaemonSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppDaemon_SubscribeClient interface {
	Recv() (*SubscriptionEvent, error)
	grpc.ClientStream
}

type appDaemonSubscribeClient struct {
	grpc.ClientStream
}

func (x *appDaemonSubscribeClient) Recv() (*SubscriptionEvent, error) {
	m := new(SubscriptionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appDaemonClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, AppDaemon_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) AnnounceDHT(ctx context.Context, in *AnnounceDHTRequest, opts ...grpc.CallOption) (*AnnounceDHTResponse, error) {
	out := new(AnnounceDHTResponse)
	err := c.cc.Invoke(ctx, AppDaemon_AnnounceDHT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appDaemonClient) LeaveDHT(ctx context.Context, in *LeaveDHTRequest, opts ...grpc.CallOption) (*LeaveDHTResponse, error) {
	out := new(LeaveDHTResponse)
	err := c.cc.Invoke(ctx, AppDaemon_LeaveDHT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// Query is used to query the mesh for information.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Metrics is used to retrieve interface metrics from the node.
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
	// Status is used to retrieve the status of the node.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Subscribe is used to subscribe to events in the mesh database.
	Subscribe(*SubscribeRequest, AppDaemon_SubscribeServer) error
	// Publish is used to publish events to the mesh database. A restricted set
	// of keys are allowed to be published to.
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	// AnnounceDHT is used to announce the node's presence on the Kademlia DHT
	// for other nodes to discover.
	AnnounceDHT(context.Context, *AnnounceDHTRequest) (*AnnounceDHTResponse, error)
	// LeaveDHT is used to leave the Kademlia DHT.
	LeaveDHT(context.Context, *LeaveDHTRequest) (*LeaveDHTResponse, error)
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
func (UnimplementedAppDaemonServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedAppDaemonServer) Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (UnimplementedAppDaemonServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedAppDaemonServer) Subscribe(*SubscribeRequest, AppDaemon_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedAppDaemonServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedAppDaemonServer) AnnounceDHT(context.Context, *AnnounceDHTRequest) (*AnnounceDHTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceDHT not implemented")
}
func (UnimplementedAppDaemonServer) LeaveDHT(context.Context, *LeaveDHTRequest) (*LeaveDHTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveDHT not implemented")
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

func _AppDaemon_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Metrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppDaemonServer).Subscribe(m, &appDaemonSubscribeServer{stream})
}

type AppDaemon_SubscribeServer interface {
	Send(*SubscriptionEvent) error
	grpc.ServerStream
}

type appDaemonSubscribeServer struct {
	grpc.ServerStream
}

func (x *appDaemonSubscribeServer) Send(m *SubscriptionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _AppDaemon_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_AnnounceDHT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnounceDHTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).AnnounceDHT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_AnnounceDHT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).AnnounceDHT(ctx, req.(*AnnounceDHTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppDaemon_LeaveDHT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveDHTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDaemonServer).LeaveDHT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppDaemon_LeaveDHT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDaemonServer).LeaveDHT(ctx, req.(*LeaveDHTRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "Query",
			Handler:    _AppDaemon_Query_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _AppDaemon_Metrics_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _AppDaemon_Status_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _AppDaemon_Publish_Handler,
		},
		{
			MethodName: "AnnounceDHT",
			Handler:    _AppDaemon_AnnounceDHT_Handler,
		},
		{
			MethodName: "LeaveDHT",
			Handler:    _AppDaemon_LeaveDHT_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _AppDaemon_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/app.proto",
}
