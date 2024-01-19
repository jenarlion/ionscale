// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ionscale/v1/routes.proto

package ionscalev1

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

type GetMachineRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64 `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *GetMachineRoutesRequest) Reset() {
	*x = GetMachineRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineRoutesRequest) ProtoMessage() {}

func (x *GetMachineRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineRoutesRequest.ProtoReflect.Descriptor instead.
func (*GetMachineRoutesRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{0}
}

func (x *GetMachineRoutesRequest) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

type GetMachineRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64         `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    *MachineRoutes `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *GetMachineRoutesResponse) Reset() {
	*x = GetMachineRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineRoutesResponse) ProtoMessage() {}

func (x *GetMachineRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineRoutesResponse.ProtoReflect.Descriptor instead.
func (*GetMachineRoutesResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{1}
}

func (x *GetMachineRoutesResponse) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *GetMachineRoutesResponse) GetRoutes() *MachineRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type EnableMachineRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64   `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    []string `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	Replace   bool     `protobuf:"varint,3,opt,name=replace,proto3" json:"replace,omitempty"`
}

func (x *EnableMachineRoutesRequest) Reset() {
	*x = EnableMachineRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableMachineRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableMachineRoutesRequest) ProtoMessage() {}

func (x *EnableMachineRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableMachineRoutesRequest.ProtoReflect.Descriptor instead.
func (*EnableMachineRoutesRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{2}
}

func (x *EnableMachineRoutesRequest) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *EnableMachineRoutesRequest) GetRoutes() []string {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *EnableMachineRoutesRequest) GetReplace() bool {
	if x != nil {
		return x.Replace
	}
	return false
}

type EnableMachineRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64         `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    *MachineRoutes `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *EnableMachineRoutesResponse) Reset() {
	*x = EnableMachineRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableMachineRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableMachineRoutesResponse) ProtoMessage() {}

func (x *EnableMachineRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableMachineRoutesResponse.ProtoReflect.Descriptor instead.
func (*EnableMachineRoutesResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{3}
}

func (x *EnableMachineRoutesResponse) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *EnableMachineRoutesResponse) GetRoutes() *MachineRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type DisableMachineRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64   `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    []string `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *DisableMachineRoutesRequest) Reset() {
	*x = DisableMachineRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableMachineRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableMachineRoutesRequest) ProtoMessage() {}

func (x *DisableMachineRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableMachineRoutesRequest.ProtoReflect.Descriptor instead.
func (*DisableMachineRoutesRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{4}
}

func (x *DisableMachineRoutesRequest) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *DisableMachineRoutesRequest) GetRoutes() []string {
	if x != nil {
		return x.Routes
	}
	return nil
}

type DisableMachineRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64         `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    *MachineRoutes `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *DisableMachineRoutesResponse) Reset() {
	*x = DisableMachineRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableMachineRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableMachineRoutesResponse) ProtoMessage() {}

func (x *DisableMachineRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableMachineRoutesResponse.ProtoReflect.Descriptor instead.
func (*DisableMachineRoutesResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{5}
}

func (x *DisableMachineRoutesResponse) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *DisableMachineRoutesResponse) GetRoutes() *MachineRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type EnableExitNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64 `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *EnableExitNodeRequest) Reset() {
	*x = EnableExitNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableExitNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableExitNodeRequest) ProtoMessage() {}

func (x *EnableExitNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableExitNodeRequest.ProtoReflect.Descriptor instead.
func (*EnableExitNodeRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{6}
}

func (x *EnableExitNodeRequest) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

type EnableExitNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64         `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    *MachineRoutes `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *EnableExitNodeResponse) Reset() {
	*x = EnableExitNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableExitNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableExitNodeResponse) ProtoMessage() {}

func (x *EnableExitNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableExitNodeResponse.ProtoReflect.Descriptor instead.
func (*EnableExitNodeResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{7}
}

func (x *EnableExitNodeResponse) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *EnableExitNodeResponse) GetRoutes() *MachineRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type DisableExitNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64 `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *DisableExitNodeRequest) Reset() {
	*x = DisableExitNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableExitNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableExitNodeRequest) ProtoMessage() {}

func (x *DisableExitNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableExitNodeRequest.ProtoReflect.Descriptor instead.
func (*DisableExitNodeRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{8}
}

func (x *DisableExitNodeRequest) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

type DisableExitNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint64         `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Routes    *MachineRoutes `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *DisableExitNodeResponse) Reset() {
	*x = DisableExitNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableExitNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableExitNodeResponse) ProtoMessage() {}

func (x *DisableExitNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableExitNodeResponse.ProtoReflect.Descriptor instead.
func (*DisableExitNodeResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{9}
}

func (x *DisableExitNodeResponse) GetMachineId() uint64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *DisableExitNodeResponse) GetRoutes() *MachineRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type MachineRoutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdvertisedRoutes   []string `protobuf:"bytes,1,rep,name=advertised_routes,json=advertisedRoutes,proto3" json:"advertised_routes,omitempty"`
	EnabledRoutes      []string `protobuf:"bytes,2,rep,name=enabled_routes,json=enabledRoutes,proto3" json:"enabled_routes,omitempty"`
	AdvertisedExitNode bool     `protobuf:"varint,3,opt,name=advertised_exit_node,json=advertisedExitNode,proto3" json:"advertised_exit_node,omitempty"`
	EnabledExitNode    bool     `protobuf:"varint,4,opt,name=enabled_exit_node,json=enabledExitNode,proto3" json:"enabled_exit_node,omitempty"`
}

func (x *MachineRoutes) Reset() {
	*x = MachineRoutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_routes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineRoutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineRoutes) ProtoMessage() {}

func (x *MachineRoutes) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_routes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineRoutes.ProtoReflect.Descriptor instead.
func (*MachineRoutes) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_routes_proto_rawDescGZIP(), []int{10}
}

func (x *MachineRoutes) GetAdvertisedRoutes() []string {
	if x != nil {
		return x.AdvertisedRoutes
	}
	return nil
}

func (x *MachineRoutes) GetEnabledRoutes() []string {
	if x != nil {
		return x.EnabledRoutes
	}
	return nil
}

func (x *MachineRoutes) GetAdvertisedExitNode() bool {
	if x != nil {
		return x.AdvertisedExitNode
	}
	return false
}

func (x *MachineRoutes) GetEnabledExitNode() bool {
	if x != nil {
		return x.EnabledExitNode
	}
	return false
}

var File_ionscale_v1_routes_proto protoreflect.FileDescriptor

var file_ionscale_v1_routes_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x22, 0x6d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x6d, 0x0a, 0x1a, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22,
	0x70, 0x0a, 0x1b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x54, 0x0a, 0x1b, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x1c, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x22, 0x6b, 0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22,
	0x37, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x45, 0x78, 0x69, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x45, 0x78, 0x69, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e,
	0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ionscale_v1_routes_proto_rawDescOnce sync.Once
	file_ionscale_v1_routes_proto_rawDescData = file_ionscale_v1_routes_proto_rawDesc
)

func file_ionscale_v1_routes_proto_rawDescGZIP() []byte {
	file_ionscale_v1_routes_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_routes_proto_rawDescData = protoimpl.X.CompressGZIP(file_ionscale_v1_routes_proto_rawDescData)
	})
	return file_ionscale_v1_routes_proto_rawDescData
}

var file_ionscale_v1_routes_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ionscale_v1_routes_proto_goTypes = []interface{}{
	(*GetMachineRoutesRequest)(nil),      // 0: ionscale.v1.GetMachineRoutesRequest
	(*GetMachineRoutesResponse)(nil),     // 1: ionscale.v1.GetMachineRoutesResponse
	(*EnableMachineRoutesRequest)(nil),   // 2: ionscale.v1.EnableMachineRoutesRequest
	(*EnableMachineRoutesResponse)(nil),  // 3: ionscale.v1.EnableMachineRoutesResponse
	(*DisableMachineRoutesRequest)(nil),  // 4: ionscale.v1.DisableMachineRoutesRequest
	(*DisableMachineRoutesResponse)(nil), // 5: ionscale.v1.DisableMachineRoutesResponse
	(*EnableExitNodeRequest)(nil),        // 6: ionscale.v1.EnableExitNodeRequest
	(*EnableExitNodeResponse)(nil),       // 7: ionscale.v1.EnableExitNodeResponse
	(*DisableExitNodeRequest)(nil),       // 8: ionscale.v1.DisableExitNodeRequest
	(*DisableExitNodeResponse)(nil),      // 9: ionscale.v1.DisableExitNodeResponse
	(*MachineRoutes)(nil),                // 10: ionscale.v1.MachineRoutes
}
var file_ionscale_v1_routes_proto_depIdxs = []int32{
	10, // 0: ionscale.v1.GetMachineRoutesResponse.routes:type_name -> ionscale.v1.MachineRoutes
	10, // 1: ionscale.v1.EnableMachineRoutesResponse.routes:type_name -> ionscale.v1.MachineRoutes
	10, // 2: ionscale.v1.DisableMachineRoutesResponse.routes:type_name -> ionscale.v1.MachineRoutes
	10, // 3: ionscale.v1.EnableExitNodeResponse.routes:type_name -> ionscale.v1.MachineRoutes
	10, // 4: ionscale.v1.DisableExitNodeResponse.routes:type_name -> ionscale.v1.MachineRoutes
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ionscale_v1_routes_proto_init() }
func file_ionscale_v1_routes_proto_init() {
	if File_ionscale_v1_routes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ionscale_v1_routes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineRoutesRequest); i {
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
		file_ionscale_v1_routes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineRoutesResponse); i {
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
		file_ionscale_v1_routes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableMachineRoutesRequest); i {
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
		file_ionscale_v1_routes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableMachineRoutesResponse); i {
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
		file_ionscale_v1_routes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableMachineRoutesRequest); i {
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
		file_ionscale_v1_routes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableMachineRoutesResponse); i {
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
		file_ionscale_v1_routes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableExitNodeRequest); i {
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
		file_ionscale_v1_routes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableExitNodeResponse); i {
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
		file_ionscale_v1_routes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableExitNodeRequest); i {
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
		file_ionscale_v1_routes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableExitNodeResponse); i {
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
		file_ionscale_v1_routes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineRoutes); i {
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
			RawDescriptor: file_ionscale_v1_routes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_routes_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_routes_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_routes_proto_msgTypes,
	}.Build()
	File_ionscale_v1_routes_proto = out.File
	file_ionscale_v1_routes_proto_rawDesc = nil
	file_ionscale_v1_routes_proto_goTypes = nil
	file_ionscale_v1_routes_proto_depIdxs = nil
}
