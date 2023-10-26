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
// source: v1/registrar.proto

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
	Registrar_Register_FullMethodName = "/v1.Registrar/Register"
	Registrar_Lookup_FullMethodName   = "/v1.Registrar/Lookup"
)

// RegistrarClient is the client API for Registrar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrarClient interface {
	// Register a public key with the registrar. An alias can be provided to make it easier
	// to lookup the public key later. If the alias is already in use, the request will fail.
	// This method can be used to change the alias of a public key by providing the same
	// public key with a different alias.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Lookup a public key by ID or alias. If the ID is not found, the
	// request will fail.
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
}

type registrarClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrarClient(cc grpc.ClientConnInterface) RegistrarClient {
	return &registrarClient{cc}
}

func (c *registrarClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Registrar_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrarClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, Registrar_Lookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrarServer is the server API for Registrar service.
// All implementations must embed UnimplementedRegistrarServer
// for forward compatibility
type RegistrarServer interface {
	// Register a public key with the registrar. An alias can be provided to make it easier
	// to lookup the public key later. If the alias is already in use, the request will fail.
	// This method can be used to change the alias of a public key by providing the same
	// public key with a different alias.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Lookup a public key by ID or alias. If the ID is not found, the
	// request will fail.
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	mustEmbedUnimplementedRegistrarServer()
}

// UnimplementedRegistrarServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrarServer struct {
}

func (UnimplementedRegistrarServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistrarServer) Lookup(context.Context, *LookupRequest) (*LookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedRegistrarServer) mustEmbedUnimplementedRegistrarServer() {}

// UnsafeRegistrarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrarServer will
// result in compilation errors.
type UnsafeRegistrarServer interface {
	mustEmbedUnimplementedRegistrarServer()
}

func RegisterRegistrarServer(s grpc.ServiceRegistrar, srv RegistrarServer) {
	s.RegisterService(&Registrar_ServiceDesc, srv)
}

func _Registrar_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrarServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registrar_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrarServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registrar_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrarServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registrar_Lookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrarServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registrar_ServiceDesc is the grpc.ServiceDesc for Registrar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registrar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Registrar",
	HandlerType: (*RegistrarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registrar_Register_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Registrar_Lookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/registrar.proto",
}