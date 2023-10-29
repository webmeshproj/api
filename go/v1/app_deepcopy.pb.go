// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using ConnectRequest within kubernetes types, where deepcopy-gen is used.
func (in *ConnectRequest) DeepCopyInto(out *ConnectRequest) {
	p := proto.Clone(in).(*ConnectRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectRequest. Required by controller-gen.
func (in *ConnectRequest) DeepCopy() *ConnectRequest {
	if in == nil {
		return nil
	}
	out := new(ConnectRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectRequest. Required by controller-gen.
func (in *ConnectRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MeshConnNetworking within kubernetes types, where deepcopy-gen is used.
func (in *MeshConnNetworking) DeepCopyInto(out *MeshConnNetworking) {
	p := proto.Clone(in).(*MeshConnNetworking)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnNetworking. Required by controller-gen.
func (in *MeshConnNetworking) DeepCopy() *MeshConnNetworking {
	if in == nil {
		return nil
	}
	out := new(MeshConnNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnNetworking. Required by controller-gen.
func (in *MeshConnNetworking) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MeshConnServices within kubernetes types, where deepcopy-gen is used.
func (in *MeshConnServices) DeepCopyInto(out *MeshConnServices) {
	p := proto.Clone(in).(*MeshConnServices)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnServices. Required by controller-gen.
func (in *MeshConnServices) DeepCopy() *MeshConnServices {
	if in == nil {
		return nil
	}
	out := new(MeshConnServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnServices. Required by controller-gen.
func (in *MeshConnServices) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MeshConnBootstrap within kubernetes types, where deepcopy-gen is used.
func (in *MeshConnBootstrap) DeepCopyInto(out *MeshConnBootstrap) {
	p := proto.Clone(in).(*MeshConnBootstrap)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnBootstrap. Required by controller-gen.
func (in *MeshConnBootstrap) DeepCopy() *MeshConnBootstrap {
	if in == nil {
		return nil
	}
	out := new(MeshConnBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnBootstrap. Required by controller-gen.
func (in *MeshConnBootstrap) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MeshConnTLS within kubernetes types, where deepcopy-gen is used.
func (in *MeshConnTLS) DeepCopyInto(out *MeshConnTLS) {
	p := proto.Clone(in).(*MeshConnTLS)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnTLS. Required by controller-gen.
func (in *MeshConnTLS) DeepCopy() *MeshConnTLS {
	if in == nil {
		return nil
	}
	out := new(MeshConnTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshConnTLS. Required by controller-gen.
func (in *MeshConnTLS) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectResponse within kubernetes types, where deepcopy-gen is used.
func (in *ConnectResponse) DeepCopyInto(out *ConnectResponse) {
	p := proto.Clone(in).(*ConnectResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectResponse. Required by controller-gen.
func (in *ConnectResponse) DeepCopy() *ConnectResponse {
	if in == nil {
		return nil
	}
	out := new(ConnectResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectResponse. Required by controller-gen.
func (in *ConnectResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DisconnectRequest within kubernetes types, where deepcopy-gen is used.
func (in *DisconnectRequest) DeepCopyInto(out *DisconnectRequest) {
	p := proto.Clone(in).(*DisconnectRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisconnectRequest. Required by controller-gen.
func (in *DisconnectRequest) DeepCopy() *DisconnectRequest {
	if in == nil {
		return nil
	}
	out := new(DisconnectRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DisconnectRequest. Required by controller-gen.
func (in *DisconnectRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DisconnectResponse within kubernetes types, where deepcopy-gen is used.
func (in *DisconnectResponse) DeepCopyInto(out *DisconnectResponse) {
	p := proto.Clone(in).(*DisconnectResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisconnectResponse. Required by controller-gen.
func (in *DisconnectResponse) DeepCopy() *DisconnectResponse {
	if in == nil {
		return nil
	}
	out := new(DisconnectResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DisconnectResponse. Required by controller-gen.
func (in *DisconnectResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MetricsRequest within kubernetes types, where deepcopy-gen is used.
func (in *MetricsRequest) DeepCopyInto(out *MetricsRequest) {
	p := proto.Clone(in).(*MetricsRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsRequest. Required by controller-gen.
func (in *MetricsRequest) DeepCopy() *MetricsRequest {
	if in == nil {
		return nil
	}
	out := new(MetricsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MetricsRequest. Required by controller-gen.
func (in *MetricsRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MetricsResponse within kubernetes types, where deepcopy-gen is used.
func (in *MetricsResponse) DeepCopyInto(out *MetricsResponse) {
	p := proto.Clone(in).(*MetricsResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsResponse. Required by controller-gen.
func (in *MetricsResponse) DeepCopy() *MetricsResponse {
	if in == nil {
		return nil
	}
	out := new(MetricsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MetricsResponse. Required by controller-gen.
func (in *MetricsResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using StatusRequest within kubernetes types, where deepcopy-gen is used.
func (in *StatusRequest) DeepCopyInto(out *StatusRequest) {
	p := proto.Clone(in).(*StatusRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusRequest. Required by controller-gen.
func (in *StatusRequest) DeepCopy() *StatusRequest {
	if in == nil {
		return nil
	}
	out := new(StatusRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StatusRequest. Required by controller-gen.
func (in *StatusRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using StatusResponse within kubernetes types, where deepcopy-gen is used.
func (in *StatusResponse) DeepCopyInto(out *StatusResponse) {
	p := proto.Clone(in).(*StatusResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusResponse. Required by controller-gen.
func (in *StatusResponse) DeepCopy() *StatusResponse {
	if in == nil {
		return nil
	}
	out := new(StatusResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StatusResponse. Required by controller-gen.
func (in *StatusResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AppQueryRequest within kubernetes types, where deepcopy-gen is used.
func (in *AppQueryRequest) DeepCopyInto(out *AppQueryRequest) {
	p := proto.Clone(in).(*AppQueryRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppQueryRequest. Required by controller-gen.
func (in *AppQueryRequest) DeepCopy() *AppQueryRequest {
	if in == nil {
		return nil
	}
	out := new(AppQueryRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AppQueryRequest. Required by controller-gen.
func (in *AppQueryRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}