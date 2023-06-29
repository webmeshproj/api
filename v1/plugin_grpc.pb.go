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
// source: v1/plugin.proto

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
	Plugin_GetInfo_FullMethodName   = "/v1.Plugin/GetInfo"
	Plugin_Configure_FullMethodName = "/v1.Plugin/Configure"
)

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	// GetInfo returns the information for the plugin.
	GetInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
	// Configure configures the plugin.
	Configure(ctx context.Context, in *PluginConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) GetInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Configure(ctx context.Context, in *PluginConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Plugin_Configure_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	// GetInfo returns the information for the plugin.
	GetInfo(context.Context, *emptypb.Empty) (*PluginInfo, error)
	// Configure configures the plugin.
	Configure(context.Context, *PluginConfiguration) (*emptypb.Empty, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) GetInfo(context.Context, *emptypb.Empty) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedPluginServer) Configure(context.Context, *PluginConfiguration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Configure(ctx, req.(*PluginConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Plugin_GetInfo_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Plugin_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/plugin.proto",
}

const (
	StoragePlugin_Store_FullMethodName           = "/v1.StoragePlugin/Store"
	StoragePlugin_RestoreSnapshot_FullMethodName = "/v1.StoragePlugin/RestoreSnapshot"
)

// StoragePluginClient is the client API for StoragePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoragePluginClient interface {
	// Store dispatches a Raft log entry for storage.
	Store(ctx context.Context, in *StoreLogRequest, opts ...grpc.CallOption) (*RaftApplyResponse, error)
	// RestoreSnapshot should drop any existing state and restore from the snapshot.
	RestoreSnapshot(ctx context.Context, in *DataSnapshot, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type storagePluginClient struct {
	cc grpc.ClientConnInterface
}

func NewStoragePluginClient(cc grpc.ClientConnInterface) StoragePluginClient {
	return &storagePluginClient{cc}
}

func (c *storagePluginClient) Store(ctx context.Context, in *StoreLogRequest, opts ...grpc.CallOption) (*RaftApplyResponse, error) {
	out := new(RaftApplyResponse)
	err := c.cc.Invoke(ctx, StoragePlugin_Store_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storagePluginClient) RestoreSnapshot(ctx context.Context, in *DataSnapshot, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, StoragePlugin_RestoreSnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoragePluginServer is the server API for StoragePlugin service.
// All implementations must embed UnimplementedStoragePluginServer
// for forward compatibility
type StoragePluginServer interface {
	// Store dispatches a Raft log entry for storage.
	Store(context.Context, *StoreLogRequest) (*RaftApplyResponse, error)
	// RestoreSnapshot should drop any existing state and restore from the snapshot.
	RestoreSnapshot(context.Context, *DataSnapshot) (*emptypb.Empty, error)
	mustEmbedUnimplementedStoragePluginServer()
}

// UnimplementedStoragePluginServer must be embedded to have forward compatible implementations.
type UnimplementedStoragePluginServer struct {
}

func (UnimplementedStoragePluginServer) Store(context.Context, *StoreLogRequest) (*RaftApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedStoragePluginServer) RestoreSnapshot(context.Context, *DataSnapshot) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreSnapshot not implemented")
}
func (UnimplementedStoragePluginServer) mustEmbedUnimplementedStoragePluginServer() {}

// UnsafeStoragePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoragePluginServer will
// result in compilation errors.
type UnsafeStoragePluginServer interface {
	mustEmbedUnimplementedStoragePluginServer()
}

func RegisterStoragePluginServer(s grpc.ServiceRegistrar, srv StoragePluginServer) {
	s.RegisterService(&StoragePlugin_ServiceDesc, srv)
}

func _StoragePlugin_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoragePluginServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoragePlugin_Store_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoragePluginServer).Store(ctx, req.(*StoreLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoragePlugin_RestoreSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSnapshot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoragePluginServer).RestoreSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoragePlugin_RestoreSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoragePluginServer).RestoreSnapshot(ctx, req.(*DataSnapshot))
	}
	return interceptor(ctx, in, info, handler)
}

// StoragePlugin_ServiceDesc is the grpc.ServiceDesc for StoragePlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoragePlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.StoragePlugin",
	HandlerType: (*StoragePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _StoragePlugin_Store_Handler,
		},
		{
			MethodName: "RestoreSnapshot",
			Handler:    _StoragePlugin_RestoreSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/plugin.proto",
}

const (
	AuthPlugin_Authenticate_FullMethodName = "/v1.AuthPlugin/Authenticate"
)

// AuthPluginClient is the client API for AuthPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthPluginClient interface {
	// Authenticate authenticates a request.
	Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type authPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthPluginClient(cc grpc.ClientConnInterface) AuthPluginClient {
	return &authPluginClient{cc}
}

func (c *authPluginClient) Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, AuthPlugin_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthPluginServer is the server API for AuthPlugin service.
// All implementations must embed UnimplementedAuthPluginServer
// for forward compatibility
type AuthPluginServer interface {
	// Authenticate authenticates a request.
	Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	mustEmbedUnimplementedAuthPluginServer()
}

// UnimplementedAuthPluginServer must be embedded to have forward compatible implementations.
type UnimplementedAuthPluginServer struct {
}

func (UnimplementedAuthPluginServer) Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAuthPluginServer) mustEmbedUnimplementedAuthPluginServer() {}

// UnsafeAuthPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthPluginServer will
// result in compilation errors.
type UnsafeAuthPluginServer interface {
	mustEmbedUnimplementedAuthPluginServer()
}

func RegisterAuthPluginServer(s grpc.ServiceRegistrar, srv AuthPluginServer) {
	s.RegisterService(&AuthPlugin_ServiceDesc, srv)
}

func _AuthPlugin_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthPluginServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthPlugin_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthPluginServer).Authenticate(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthPlugin_ServiceDesc is the grpc.ServiceDesc for AuthPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AuthPlugin",
	HandlerType: (*AuthPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthPlugin_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/plugin.proto",
}

const (
	WatchPlugin_Emit_FullMethodName = "/v1.WatchPlugin/Emit"
)

// WatchPluginClient is the client API for WatchPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchPluginClient interface {
	// Emit handles a watch event.
	Emit(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type watchPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchPluginClient(cc grpc.ClientConnInterface) WatchPluginClient {
	return &watchPluginClient{cc}
}

func (c *watchPluginClient) Emit(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WatchPlugin_Emit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchPluginServer is the server API for WatchPlugin service.
// All implementations must embed UnimplementedWatchPluginServer
// for forward compatibility
type WatchPluginServer interface {
	// Emit handles a watch event.
	Emit(context.Context, *Event) (*emptypb.Empty, error)
	mustEmbedUnimplementedWatchPluginServer()
}

// UnimplementedWatchPluginServer must be embedded to have forward compatible implementations.
type UnimplementedWatchPluginServer struct {
}

func (UnimplementedWatchPluginServer) Emit(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Emit not implemented")
}
func (UnimplementedWatchPluginServer) mustEmbedUnimplementedWatchPluginServer() {}

// UnsafeWatchPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchPluginServer will
// result in compilation errors.
type UnsafeWatchPluginServer interface {
	mustEmbedUnimplementedWatchPluginServer()
}

func RegisterWatchPluginServer(s grpc.ServiceRegistrar, srv WatchPluginServer) {
	s.RegisterService(&WatchPlugin_ServiceDesc, srv)
}

func _WatchPlugin_Emit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchPluginServer).Emit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WatchPlugin_Emit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchPluginServer).Emit(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

// WatchPlugin_ServiceDesc is the grpc.ServiceDesc for WatchPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WatchPlugin",
	HandlerType: (*WatchPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emit",
			Handler:    _WatchPlugin_Emit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/plugin.proto",
}

const (
	IPAMPlugin_AllocateIP_FullMethodName = "/v1.IPAMPlugin/AllocateIP"
	IPAMPlugin_LookupIP_FullMethodName   = "/v1.IPAMPlugin/LookupIP"
)

// IPAMPluginClient is the client API for IPAMPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPAMPluginClient interface {
	// AllocateIP allocates an IP for a node.
	AllocateIP(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocatedIP, error)
	// LookupIP looks up an IP for a node.
	LookupIP(ctx context.Context, in *LookupIPRequest, opts ...grpc.CallOption) (*AllocatedIP, error)
}

type iPAMPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewIPAMPluginClient(cc grpc.ClientConnInterface) IPAMPluginClient {
	return &iPAMPluginClient{cc}
}

func (c *iPAMPluginClient) AllocateIP(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocatedIP, error) {
	out := new(AllocatedIP)
	err := c.cc.Invoke(ctx, IPAMPlugin_AllocateIP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAMPluginClient) LookupIP(ctx context.Context, in *LookupIPRequest, opts ...grpc.CallOption) (*AllocatedIP, error) {
	out := new(AllocatedIP)
	err := c.cc.Invoke(ctx, IPAMPlugin_LookupIP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPAMPluginServer is the server API for IPAMPlugin service.
// All implementations must embed UnimplementedIPAMPluginServer
// for forward compatibility
type IPAMPluginServer interface {
	// AllocateIP allocates an IP for a node.
	AllocateIP(context.Context, *AllocateIPRequest) (*AllocatedIP, error)
	// LookupIP looks up an IP for a node.
	LookupIP(context.Context, *LookupIPRequest) (*AllocatedIP, error)
	mustEmbedUnimplementedIPAMPluginServer()
}

// UnimplementedIPAMPluginServer must be embedded to have forward compatible implementations.
type UnimplementedIPAMPluginServer struct {
}

func (UnimplementedIPAMPluginServer) AllocateIP(context.Context, *AllocateIPRequest) (*AllocatedIP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateIP not implemented")
}
func (UnimplementedIPAMPluginServer) LookupIP(context.Context, *LookupIPRequest) (*AllocatedIP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupIP not implemented")
}
func (UnimplementedIPAMPluginServer) mustEmbedUnimplementedIPAMPluginServer() {}

// UnsafeIPAMPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPAMPluginServer will
// result in compilation errors.
type UnsafeIPAMPluginServer interface {
	mustEmbedUnimplementedIPAMPluginServer()
}

func RegisterIPAMPluginServer(s grpc.ServiceRegistrar, srv IPAMPluginServer) {
	s.RegisterService(&IPAMPlugin_ServiceDesc, srv)
}

func _IPAMPlugin_AllocateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMPluginServer).AllocateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPAMPlugin_AllocateIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMPluginServer).AllocateIP(ctx, req.(*AllocateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAMPlugin_LookupIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMPluginServer).LookupIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPAMPlugin_LookupIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMPluginServer).LookupIP(ctx, req.(*LookupIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPAMPlugin_ServiceDesc is the grpc.ServiceDesc for IPAMPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPAMPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.IPAMPlugin",
	HandlerType: (*IPAMPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateIP",
			Handler:    _IPAMPlugin_AllocateIP_Handler,
		},
		{
			MethodName: "LookupIP",
			Handler:    _IPAMPlugin_LookupIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/plugin.proto",
}
