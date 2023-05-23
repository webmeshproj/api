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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/node_raft.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RaftCommandType is the type of command being sent to the
// Raft log.
type RaftCommandType int32

const (
	// UNKNOWN is the unknown command type.
	RaftCommandType_UNKNOWN RaftCommandType = 0
	// QUERY is the query command type.
	RaftCommandType_QUERY RaftCommandType = 1
	// EXECUTE is the execute command type.
	RaftCommandType_EXECUTE RaftCommandType = 2
)

// Enum value maps for RaftCommandType.
var (
	RaftCommandType_name = map[int32]string{
		0: "UNKNOWN",
		1: "QUERY",
		2: "EXECUTE",
	}
	RaftCommandType_value = map[string]int32{
		"UNKNOWN": 0,
		"QUERY":   1,
		"EXECUTE": 2,
	}
)

func (x RaftCommandType) Enum() *RaftCommandType {
	p := new(RaftCommandType)
	*p = x
	return p
}

func (x RaftCommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaftCommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_node_raft_proto_enumTypes[0].Descriptor()
}

func (RaftCommandType) Type() protoreflect.EnumType {
	return &file_v1_node_raft_proto_enumTypes[0]
}

func (x RaftCommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RaftCommandType.Descriptor instead.
func (RaftCommandType) EnumDescriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{0}
}

// SQLParameter is a parameter of a SQL query.
type SQLParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the parameter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value is the value of the parameter.
	//
	// Types that are assignable to Value:
	//
	//	*SQLParameter_Int64
	//	*SQLParameter_Double
	//	*SQLParameter_Bool
	//	*SQLParameter_Bytes
	//	*SQLParameter_Str
	//	*SQLParameter_Time
	Value isSQLParameter_Value `protobuf_oneof:"value"`
}

func (x *SQLParameter) Reset() {
	*x = SQLParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLParameter) ProtoMessage() {}

func (x *SQLParameter) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLParameter.ProtoReflect.Descriptor instead.
func (*SQLParameter) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{0}
}

func (x *SQLParameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *SQLParameter) GetValue() isSQLParameter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SQLParameter) GetInt64() int64 {
	if x, ok := x.GetValue().(*SQLParameter_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *SQLParameter) GetDouble() float64 {
	if x, ok := x.GetValue().(*SQLParameter_Double); ok {
		return x.Double
	}
	return 0
}

func (x *SQLParameter) GetBool() bool {
	if x, ok := x.GetValue().(*SQLParameter_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *SQLParameter) GetBytes() []byte {
	if x, ok := x.GetValue().(*SQLParameter_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *SQLParameter) GetStr() string {
	if x, ok := x.GetValue().(*SQLParameter_Str); ok {
		return x.Str
	}
	return ""
}

func (x *SQLParameter) GetTime() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*SQLParameter_Time); ok {
		return x.Time
	}
	return nil
}

type isSQLParameter_Value interface {
	isSQLParameter_Value()
}

type SQLParameter_Int64 struct {
	Int64 int64 `protobuf:"zigzag64,2,opt,name=int64,proto3,oneof"`
}

type SQLParameter_Double struct {
	Double float64 `protobuf:"fixed64,3,opt,name=double,proto3,oneof"`
}

type SQLParameter_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=bool,proto3,oneof"`
}

type SQLParameter_Bytes struct {
	Bytes []byte `protobuf:"bytes,5,opt,name=bytes,proto3,oneof"`
}

type SQLParameter_Str struct {
	Str string `protobuf:"bytes,6,opt,name=str,proto3,oneof"`
}

type SQLParameter_Time struct {
	Time *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=time,proto3,oneof"`
}

func (*SQLParameter_Int64) isSQLParameter_Value() {}

func (*SQLParameter_Double) isSQLParameter_Value() {}

func (*SQLParameter_Bool) isSQLParameter_Value() {}

func (*SQLParameter_Bytes) isSQLParameter_Value() {}

func (*SQLParameter_Str) isSQLParameter_Value() {}

func (*SQLParameter_Time) isSQLParameter_Value() {}

// SQLValues is a list of values.
type SQLValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// values is the list of values.
	Values []*SQLParameter `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SQLValues) Reset() {
	*x = SQLValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLValues) ProtoMessage() {}

func (x *SQLValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLValues.ProtoReflect.Descriptor instead.
func (*SQLValues) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{1}
}

func (x *SQLValues) GetValues() []*SQLParameter {
	if x != nil {
		return x.Values
	}
	return nil
}

// SQLStatement is a SQL statement.
type SQLStatement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sql is the SQL statement.
	Sql string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	// parameters is the parameters of the SQL statement.
	Parameters []*SQLParameter `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *SQLStatement) Reset() {
	*x = SQLStatement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLStatement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLStatement) ProtoMessage() {}

func (x *SQLStatement) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLStatement.ProtoReflect.Descriptor instead.
func (*SQLStatement) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{2}
}

func (x *SQLStatement) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

func (x *SQLStatement) GetParameters() []*SQLParameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// SQLQuery is a SQL query.
type SQLQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction flags whether the query is a transaction.
	Transaction bool `protobuf:"varint,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// statement is the statement of the SQL query.
	Statement *SQLStatement `protobuf:"bytes,2,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (x *SQLQuery) Reset() {
	*x = SQLQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLQuery) ProtoMessage() {}

func (x *SQLQuery) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLQuery.ProtoReflect.Descriptor instead.
func (*SQLQuery) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{3}
}

func (x *SQLQuery) GetTransaction() bool {
	if x != nil {
		return x.Transaction
	}
	return false
}

func (x *SQLQuery) GetStatement() *SQLStatement {
	if x != nil {
		return x.Statement
	}
	return nil
}

// SQLExec is a SQL exec.
type SQLExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction flags whether the exec is a transaction.
	Transaction bool `protobuf:"varint,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// statement is the statement of the SQL exec.
	Statement *SQLStatement `protobuf:"bytes,2,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (x *SQLExec) Reset() {
	*x = SQLExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLExec) ProtoMessage() {}

func (x *SQLExec) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLExec.ProtoReflect.Descriptor instead.
func (*SQLExec) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{4}
}

func (x *SQLExec) GetTransaction() bool {
	if x != nil {
		return x.Transaction
	}
	return false
}

func (x *SQLExec) GetStatement() *SQLStatement {
	if x != nil {
		return x.Statement
	}
	return nil
}

// SQLQueryResult contains the rows of a SQL query.
type SQLQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// columns is the list of columns.
	Columns []string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	// types is the list of types.
	Types []string `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	// values is the list of values.
	Values []*SQLValues `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	// error is an error that occurred during the query.
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	// time is the time it took to execute the query.
	Time string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SQLQueryResult) Reset() {
	*x = SQLQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLQueryResult) ProtoMessage() {}

func (x *SQLQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLQueryResult.ProtoReflect.Descriptor instead.
func (*SQLQueryResult) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{5}
}

func (x *SQLQueryResult) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *SQLQueryResult) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *SQLQueryResult) GetValues() []*SQLValues {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *SQLQueryResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SQLQueryResult) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

// SQLExecResult is the result of a SQL exec.
type SQLExecResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// last_insert_id is the last insert ID.
	LastInsertId int64 `protobuf:"varint,1,opt,name=last_insert_id,json=lastInsertId,proto3" json:"last_insert_id,omitempty"`
	// rows_affected is the number of rows affected.
	RowsAffected int64 `protobuf:"varint,2,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	// error is an error that occurred during the exec.
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// time is the time it took to execute the exec.
	Time string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SQLExecResult) Reset() {
	*x = SQLExecResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLExecResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLExecResult) ProtoMessage() {}

func (x *SQLExecResult) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLExecResult.ProtoReflect.Descriptor instead.
func (*SQLExecResult) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{6}
}

func (x *SQLExecResult) GetLastInsertId() int64 {
	if x != nil {
		return x.LastInsertId
	}
	return 0
}

func (x *SQLExecResult) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *SQLExecResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SQLExecResult) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

// RaftLogEntry is the data of an entry in the Raft log.
type RaftLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the type of the log entry.
	Type RaftCommandType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.RaftCommandType" json:"type,omitempty"`
	// data is the data of the log entry.
	//
	// Types that are assignable to Data:
	//
	//	*RaftLogEntry_SqlQuery
	//	*RaftLogEntry_SqlExec
	Data isRaftLogEntry_Data `protobuf_oneof:"data"`
}

func (x *RaftLogEntry) Reset() {
	*x = RaftLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftLogEntry) ProtoMessage() {}

func (x *RaftLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftLogEntry.ProtoReflect.Descriptor instead.
func (*RaftLogEntry) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{7}
}

func (x *RaftLogEntry) GetType() RaftCommandType {
	if x != nil {
		return x.Type
	}
	return RaftCommandType_UNKNOWN
}

func (m *RaftLogEntry) GetData() isRaftLogEntry_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *RaftLogEntry) GetSqlQuery() *SQLQuery {
	if x, ok := x.GetData().(*RaftLogEntry_SqlQuery); ok {
		return x.SqlQuery
	}
	return nil
}

func (x *RaftLogEntry) GetSqlExec() *SQLExec {
	if x, ok := x.GetData().(*RaftLogEntry_SqlExec); ok {
		return x.SqlExec
	}
	return nil
}

type isRaftLogEntry_Data interface {
	isRaftLogEntry_Data()
}

type RaftLogEntry_SqlQuery struct {
	// sql_query is the SQL query of the log entry.
	SqlQuery *SQLQuery `protobuf:"bytes,2,opt,name=sql_query,json=sqlQuery,proto3,oneof"`
}

type RaftLogEntry_SqlExec struct {
	// sql_exec is the SQL exec of the log entry.
	SqlExec *SQLExec `protobuf:"bytes,3,opt,name=sql_exec,json=sqlExec,proto3,oneof"`
}

func (*RaftLogEntry_SqlQuery) isRaftLogEntry_Data() {}

func (*RaftLogEntry_SqlExec) isRaftLogEntry_Data() {}

// RaftApplyResponse is the response to an apply request. It
// contains the result of applying the log entry.
type RaftApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time is the total time it took to apply the log entry.
	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// result is the result of the log entry.
	//
	// Types that are assignable to Result:
	//
	//	*RaftApplyResponse_Rows
	//	*RaftApplyResponse_Exec
	//	*RaftApplyResponse_Error
	Result isRaftApplyResponse_Result `protobuf_oneof:"result"`
}

func (x *RaftApplyResponse) Reset() {
	*x = RaftApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_node_raft_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftApplyResponse) ProtoMessage() {}

func (x *RaftApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_node_raft_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftApplyResponse.ProtoReflect.Descriptor instead.
func (*RaftApplyResponse) Descriptor() ([]byte, []int) {
	return file_v1_node_raft_proto_rawDescGZIP(), []int{8}
}

func (x *RaftApplyResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (m *RaftApplyResponse) GetResult() isRaftApplyResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *RaftApplyResponse) GetRows() *SQLQueryResult {
	if x, ok := x.GetResult().(*RaftApplyResponse_Rows); ok {
		return x.Rows
	}
	return nil
}

func (x *RaftApplyResponse) GetExec() *SQLExecResult {
	if x, ok := x.GetResult().(*RaftApplyResponse_Exec); ok {
		return x.Exec
	}
	return nil
}

func (x *RaftApplyResponse) GetError() string {
	if x, ok := x.GetResult().(*RaftApplyResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isRaftApplyResponse_Result interface {
	isRaftApplyResponse_Result()
}

type RaftApplyResponse_Rows struct {
	// query is the query result of the log entry.
	Rows *SQLQueryResult `protobuf:"bytes,2,opt,name=rows,proto3,oneof"`
}

type RaftApplyResponse_Exec struct {
	// exec is the exec result of the log entry.
	Exec *SQLExecResult `protobuf:"bytes,3,opt,name=exec,proto3,oneof"`
}

type RaftApplyResponse_Error struct {
	// error is an error that occurred during the apply.
	Error string `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*RaftApplyResponse_Rows) isRaftApplyResponse_Result() {}

func (*RaftApplyResponse_Exec) isRaftApplyResponse_Result() {}

func (*RaftApplyResponse_Error) isRaftApplyResponse_Result() {}

var File_v1_node_raft_proto protoreflect.FileDescriptor

var file_v1_node_raft_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0c, 0x53, 0x51,
	0x4c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a,
	0x09, 0x53, 0x51, 0x4c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x51, 0x4c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0c, 0x53, 0x51, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x51, 0x4c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x5c, 0x0a, 0x08, 0x53, 0x51, 0x4c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x51, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x07, 0x53, 0x51, 0x4c, 0x45, 0x78, 0x65,
	0x63, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x51, 0x4c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x53, 0x51, 0x4c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x51, 0x4c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x53, 0x51, 0x4c, 0x45,
	0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x96,
	0x01, 0x0a, 0x0c, 0x52, 0x61, 0x66, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x73, 0x71, 0x6c, 0x5f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x51, 0x4c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x73, 0x71, 0x6c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x71, 0x6c, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x51, 0x4c,
	0x45, 0x78, 0x65, 0x63, 0x48, 0x00, 0x52, 0x07, 0x73, 0x71, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9c, 0x01, 0x0a, 0x11, 0x52, 0x61, 0x66, 0x74,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x51, 0x4c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x65,
	0x78, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x51, 0x4c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x65, 0x78, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x36, 0x0a, 0x0f, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x02, 0x42, 0x65,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x61,
	0x66, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2,
	0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_node_raft_proto_rawDescOnce sync.Once
	file_v1_node_raft_proto_rawDescData = file_v1_node_raft_proto_rawDesc
)

func file_v1_node_raft_proto_rawDescGZIP() []byte {
	file_v1_node_raft_proto_rawDescOnce.Do(func() {
		file_v1_node_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_node_raft_proto_rawDescData)
	})
	return file_v1_node_raft_proto_rawDescData
}

var file_v1_node_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_node_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_node_raft_proto_goTypes = []interface{}{
	(RaftCommandType)(0),          // 0: v1.RaftCommandType
	(*SQLParameter)(nil),          // 1: v1.SQLParameter
	(*SQLValues)(nil),             // 2: v1.SQLValues
	(*SQLStatement)(nil),          // 3: v1.SQLStatement
	(*SQLQuery)(nil),              // 4: v1.SQLQuery
	(*SQLExec)(nil),               // 5: v1.SQLExec
	(*SQLQueryResult)(nil),        // 6: v1.SQLQueryResult
	(*SQLExecResult)(nil),         // 7: v1.SQLExecResult
	(*RaftLogEntry)(nil),          // 8: v1.RaftLogEntry
	(*RaftApplyResponse)(nil),     // 9: v1.RaftApplyResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_v1_node_raft_proto_depIdxs = []int32{
	10, // 0: v1.SQLParameter.time:type_name -> google.protobuf.Timestamp
	1,  // 1: v1.SQLValues.values:type_name -> v1.SQLParameter
	1,  // 2: v1.SQLStatement.parameters:type_name -> v1.SQLParameter
	3,  // 3: v1.SQLQuery.statement:type_name -> v1.SQLStatement
	3,  // 4: v1.SQLExec.statement:type_name -> v1.SQLStatement
	2,  // 5: v1.SQLQueryResult.values:type_name -> v1.SQLValues
	0,  // 6: v1.RaftLogEntry.type:type_name -> v1.RaftCommandType
	4,  // 7: v1.RaftLogEntry.sql_query:type_name -> v1.SQLQuery
	5,  // 8: v1.RaftLogEntry.sql_exec:type_name -> v1.SQLExec
	6,  // 9: v1.RaftApplyResponse.rows:type_name -> v1.SQLQueryResult
	7,  // 10: v1.RaftApplyResponse.exec:type_name -> v1.SQLExecResult
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v1_node_raft_proto_init() }
func file_v1_node_raft_proto_init() {
	if File_v1_node_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_node_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLParameter); i {
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
		file_v1_node_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLValues); i {
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
		file_v1_node_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLStatement); i {
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
		file_v1_node_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLQuery); i {
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
		file_v1_node_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLExec); i {
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
		file_v1_node_raft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLQueryResult); i {
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
		file_v1_node_raft_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLExecResult); i {
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
		file_v1_node_raft_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftLogEntry); i {
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
		file_v1_node_raft_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftApplyResponse); i {
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
	file_v1_node_raft_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SQLParameter_Int64)(nil),
		(*SQLParameter_Double)(nil),
		(*SQLParameter_Bool)(nil),
		(*SQLParameter_Bytes)(nil),
		(*SQLParameter_Str)(nil),
		(*SQLParameter_Time)(nil),
	}
	file_v1_node_raft_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*RaftLogEntry_SqlQuery)(nil),
		(*RaftLogEntry_SqlExec)(nil),
	}
	file_v1_node_raft_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*RaftApplyResponse_Rows)(nil),
		(*RaftApplyResponse_Exec)(nil),
		(*RaftApplyResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_node_raft_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_node_raft_proto_goTypes,
		DependencyIndexes: file_v1_node_raft_proto_depIdxs,
		EnumInfos:         file_v1_node_raft_proto_enumTypes,
		MessageInfos:      file_v1_node_raft_proto_msgTypes,
	}.Build()
	File_v1_node_raft_proto = out.File
	file_v1_node_raft_proto_rawDesc = nil
	file_v1_node_raft_proto_goTypes = nil
	file_v1_node_raft_proto_depIdxs = nil
}
