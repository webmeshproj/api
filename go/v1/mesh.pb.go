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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/mesh.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetNodeRequest is a request to get a node.
type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the ID of the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{0}
}

func (x *GetNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// MeshEdge is an edge between two nodes.
type MeshEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source is the source node.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Target is the target node.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// Weight is the weight of the edge.
	Weight int32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	// Attributes is a list of attributes for the edge.
	Attributes map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MeshEdge) Reset() {
	*x = MeshEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshEdge) ProtoMessage() {}

func (x *MeshEdge) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshEdge.ProtoReflect.Descriptor instead.
func (*MeshEdge) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{1}
}

func (x *MeshEdge) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MeshEdge) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *MeshEdge) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *MeshEdge) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// MeshEdges is a list of edges.
type MeshEdges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Items is the list of edges.
	Items []*MeshEdge `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *MeshEdges) Reset() {
	*x = MeshEdges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshEdges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshEdges) ProtoMessage() {}

func (x *MeshEdges) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshEdges.ProtoReflect.Descriptor instead.
func (*MeshEdges) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{2}
}

func (x *MeshEdges) GetItems() []*MeshEdge {
	if x != nil {
		return x.Items
	}
	return nil
}

// MeshGraph is a graph of nodes.
type MeshGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nodes is the list of nodes.
	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// Edges is the list of edges.
	Edges []*MeshEdge `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`
	// DOT is the DOT representation of the graph.
	Dot string `protobuf:"bytes,3,opt,name=dot,proto3" json:"dot,omitempty"`
}

func (x *MeshGraph) Reset() {
	*x = MeshGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mesh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshGraph) ProtoMessage() {}

func (x *MeshGraph) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mesh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshGraph.ProtoReflect.Descriptor instead.
func (*MeshGraph) Descriptor() ([]byte, []int) {
	return file_v1_mesh_proto_rawDescGZIP(), []int{3}
}

func (x *MeshGraph) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *MeshGraph) GetEdges() []*MeshEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

func (x *MeshGraph) GetDot() string {
	if x != nil {
		return x.Dot
	}
	return ""
}

var File_v1_mesh_proto protoreflect.FileDescriptor

var file_v1_mesh_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x57, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x68, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x74, 0x32, 0xa3,
	0x01, 0x0a, 0x04, 0x4d, 0x65, 0x73, 0x68, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x22, 0x00, 0x42, 0x65, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x70,
	0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31,
	0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_mesh_proto_rawDescOnce sync.Once
	file_v1_mesh_proto_rawDescData = file_v1_mesh_proto_rawDesc
)

func file_v1_mesh_proto_rawDescGZIP() []byte {
	file_v1_mesh_proto_rawDescOnce.Do(func() {
		file_v1_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_mesh_proto_rawDescData)
	})
	return file_v1_mesh_proto_rawDescData
}

var file_v1_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_mesh_proto_goTypes = []interface{}{
	(*GetNodeRequest)(nil), // 0: v1.GetNodeRequest
	(*MeshEdge)(nil),       // 1: v1.MeshEdge
	(*MeshEdges)(nil),      // 2: v1.MeshEdges
	(*MeshGraph)(nil),      // 3: v1.MeshGraph
	nil,                    // 4: v1.MeshEdge.AttributesEntry
	(*emptypb.Empty)(nil),  // 5: google.protobuf.Empty
	(*MeshNode)(nil),       // 6: v1.MeshNode
	(*NodeList)(nil),       // 7: v1.NodeList
}
var file_v1_mesh_proto_depIdxs = []int32{
	4, // 0: v1.MeshEdge.attributes:type_name -> v1.MeshEdge.AttributesEntry
	1, // 1: v1.MeshEdges.items:type_name -> v1.MeshEdge
	1, // 2: v1.MeshGraph.edges:type_name -> v1.MeshEdge
	0, // 3: v1.Mesh.GetNode:input_type -> v1.GetNodeRequest
	5, // 4: v1.Mesh.ListNodes:input_type -> google.protobuf.Empty
	5, // 5: v1.Mesh.GetMeshGraph:input_type -> google.protobuf.Empty
	6, // 6: v1.Mesh.GetNode:output_type -> v1.MeshNode
	7, // 7: v1.Mesh.ListNodes:output_type -> v1.NodeList
	3, // 8: v1.Mesh.GetMeshGraph:output_type -> v1.MeshGraph
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_mesh_proto_init() }
func file_v1_mesh_proto_init() {
	if File_v1_mesh_proto != nil {
		return
	}
	file_v1_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_mesh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_v1_mesh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshEdge); i {
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
		file_v1_mesh_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshEdges); i {
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
		file_v1_mesh_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshGraph); i {
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
			RawDescriptor: file_v1_mesh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_mesh_proto_goTypes,
		DependencyIndexes: file_v1_mesh_proto_depIdxs,
		MessageInfos:      file_v1_mesh_proto_msgTypes,
	}.Build()
	File_v1_mesh_proto = out.File
	file_v1_mesh_proto_rawDesc = nil
	file_v1_mesh_proto_goTypes = nil
	file_v1_mesh_proto_depIdxs = nil
}
