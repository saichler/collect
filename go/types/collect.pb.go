// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: collect.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	Operation_Invalid_Operation Operation = 0
	Operation__Get              Operation = 1
	Operation__Map              Operation = 2
	Operation__Table            Operation = 3
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "Invalid_Operation",
		1: "_Get",
		2: "_Map",
		3: "_Table",
	}
	Operation_value = map[string]int32{
		"Invalid_Operation": 0,
		"_Get":              1,
		"_Map":              2,
		"_Table":            3,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_collect_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_collect_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{0}
}

type PollConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Vendor    string         `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Series    string         `protobuf:"bytes,3,opt,name=series,proto3" json:"series,omitempty"`
	Family    string         `protobuf:"bytes,4,opt,name=family,proto3" json:"family,omitempty"`
	Software  string         `protobuf:"bytes,5,opt,name=software,proto3" json:"software,omitempty"`
	Hardware  string         `protobuf:"bytes,6,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Version   string         `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Groups    []string       `protobuf:"bytes,8,rep,name=groups,proto3" json:"groups,omitempty"`
	What      string         `protobuf:"bytes,9,opt,name=what,proto3" json:"what,omitempty"`
	Operation Operation      `protobuf:"varint,10,opt,name=operation,proto3,enum=types.Operation" json:"operation,omitempty"`
	Protocol  Protocol       `protobuf:"varint,11,opt,name=protocol,proto3,enum=types.Protocol" json:"protocol,omitempty"`
	Parsing   *ParsingConfig `protobuf:"bytes,12,opt,name=parsing,proto3" json:"parsing,omitempty"`
	Cadence   int64          `protobuf:"varint,13,opt,name=cadence,proto3" json:"cadence,omitempty"`
	Timeout   int64          `protobuf:"varint,14,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *PollConfig) Reset() {
	*x = PollConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollConfig) ProtoMessage() {}

func (x *PollConfig) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollConfig.ProtoReflect.Descriptor instead.
func (*PollConfig) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{0}
}

func (x *PollConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PollConfig) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *PollConfig) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *PollConfig) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *PollConfig) GetSoftware() string {
	if x != nil {
		return x.Software
	}
	return ""
}

func (x *PollConfig) GetHardware() string {
	if x != nil {
		return x.Hardware
	}
	return ""
}

func (x *PollConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PollConfig) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *PollConfig) GetWhat() string {
	if x != nil {
		return x.What
	}
	return ""
}

func (x *PollConfig) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_Invalid_Operation
}

func (x *PollConfig) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_Invalid_Protocol
}

func (x *PollConfig) GetParsing() *ParsingConfig {
	if x != nil {
		return x.Parsing
	}
	return nil
}

func (x *PollConfig) GetCadence() int64 {
	if x != nil {
		return x.Cadence
	}
	return 0
}

func (x *PollConfig) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string             `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Result   []byte             `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Started  int64              `protobuf:"varint,3,opt,name=started,proto3" json:"started,omitempty"`
	Ended    int64              `protobuf:"varint,4,opt,name=ended,proto3" json:"ended,omitempty"`
	Cadence  int64              `protobuf:"varint,5,opt,name=cadence,proto3" json:"cadence,omitempty"`
	Timeout  int64              `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DeviceId string             `protobuf:"bytes,7,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	HostId   string             `protobuf:"bytes,8,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	PollName string             `protobuf:"bytes,9,opt,name=poll_name,json=pollName,proto3" json:"poll_name,omitempty"`
	IService *DeviceServiceInfo `protobuf:"bytes,10,opt,name=iService,proto3" json:"iService,omitempty"`
	PService *DeviceServiceInfo `protobuf:"bytes,11,opt,name=pService,proto3" json:"pService,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{1}
}

func (x *Job) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Job) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Job) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *Job) GetEnded() int64 {
	if x != nil {
		return x.Ended
	}
	return 0
}

func (x *Job) GetCadence() int64 {
	if x != nil {
		return x.Cadence
	}
	return 0
}

func (x *Job) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Job) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Job) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Job) GetPollName() string {
	if x != nil {
		return x.PollName
	}
	return ""
}

func (x *Job) GetIService() *DeviceServiceInfo {
	if x != nil {
		return x.IService
	}
	return nil
}

func (x *Job) GetPService() *DeviceServiceInfo {
	if x != nil {
		return x.PService
	}
	return nil
}

type CMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CMap) Reset() {
	*x = CMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMap) ProtoMessage() {}

func (x *CMap) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMap.ProtoReflect.Descriptor instead.
func (*CMap) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{2}
}

func (x *CMap) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns map[int32]string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Rows    map[int32]*CRow  `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CTable) Reset() {
	*x = CTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CTable) ProtoMessage() {}

func (x *CTable) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CTable.ProtoReflect.Descriptor instead.
func (*CTable) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{3}
}

func (x *CTable) GetColumns() map[int32]string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *CTable) GetRows() map[int32]*CRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

type CRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[int32][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CRow) Reset() {
	*x = CRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRow) ProtoMessage() {}

func (x *CRow) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRow.ProtoReflect.Descriptor instead.
func (*CRow) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{4}
}

func (x *CRow) GetData() map[int32][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_collect_proto protoreflect.FileDescriptor

var file_collect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x61, 0x72,
	0x73, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x68, 0x61, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x73, 0x69,
	0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xd6, 0x02, 0x0a, 0x03,
	0x4a, 0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x6a, 0x0a, 0x04, 0x43, 0x4d, 0x61, 0x70, 0x12, 0x29, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x43, 0x4d, 0x61, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xed, 0x01, 0x0a, 0x06, 0x43, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x12, 0x2b, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52,
	0x6f, 0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x1a, 0x3a,
	0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x09, 0x52, 0x6f,
	0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x52, 0x6f, 0x77, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x6a, 0x0a, 0x04, 0x43, 0x52, 0x6f, 0x77, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x52, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x42, 0x0a, 0x09,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x5f, 0x4d,
	0x61, 0x70, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x03,
	0x42, 0x25, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x07,
	0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_proto_rawDescOnce sync.Once
	file_collect_proto_rawDescData = file_collect_proto_rawDesc
)

func file_collect_proto_rawDescGZIP() []byte {
	file_collect_proto_rawDescOnce.Do(func() {
		file_collect_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_proto_rawDescData)
	})
	return file_collect_proto_rawDescData
}

var file_collect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_collect_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_collect_proto_goTypes = []interface{}{
	(Operation)(0),            // 0: types.Operation
	(*PollConfig)(nil),        // 1: types.PollConfig
	(*Job)(nil),               // 2: types.Job
	(*CMap)(nil),              // 3: types.CMap
	(*CTable)(nil),            // 4: types.CTable
	(*CRow)(nil),              // 5: types.CRow
	nil,                       // 6: types.CMap.DataEntry
	nil,                       // 7: types.CTable.ColumnsEntry
	nil,                       // 8: types.CTable.RowsEntry
	nil,                       // 9: types.CRow.DataEntry
	(Protocol)(0),             // 10: types.Protocol
	(*ParsingConfig)(nil),     // 11: types.ParsingConfig
	(*DeviceServiceInfo)(nil), // 12: types.DeviceServiceInfo
}
var file_collect_proto_depIdxs = []int32{
	0,  // 0: types.PollConfig.operation:type_name -> types.Operation
	10, // 1: types.PollConfig.protocol:type_name -> types.Protocol
	11, // 2: types.PollConfig.parsing:type_name -> types.ParsingConfig
	12, // 3: types.Job.iService:type_name -> types.DeviceServiceInfo
	12, // 4: types.Job.pService:type_name -> types.DeviceServiceInfo
	6,  // 5: types.CMap.data:type_name -> types.CMap.DataEntry
	7,  // 6: types.CTable.columns:type_name -> types.CTable.ColumnsEntry
	8,  // 7: types.CTable.rows:type_name -> types.CTable.RowsEntry
	9,  // 8: types.CRow.data:type_name -> types.CRow.DataEntry
	5,  // 9: types.CTable.RowsEntry.value:type_name -> types.CRow
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_collect_proto_init() }
func file_collect_proto_init() {
	if File_collect_proto != nil {
		return
	}
	file_device_config_proto_init()
	file_parsing_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_collect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollConfig); i {
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
		file_collect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_collect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMap); i {
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
		file_collect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CTable); i {
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
		file_collect_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CRow); i {
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
			RawDescriptor: file_collect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collect_proto_goTypes,
		DependencyIndexes: file_collect_proto_depIdxs,
		EnumInfos:         file_collect_proto_enumTypes,
		MessageInfos:      file_collect_proto_msgTypes,
	}.Build()
	File_collect_proto = out.File
	file_collect_proto_rawDesc = nil
	file_collect_proto_goTypes = nil
	file_collect_proto_depIdxs = nil
}
