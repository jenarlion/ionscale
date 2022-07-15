// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ionscale/v1/dns.proto

package ionscalev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type GetDNSConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TailnetId uint64 `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
}

func (x *GetDNSConfigRequest) Reset() {
	*x = GetDNSConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSConfigRequest) ProtoMessage() {}

func (x *GetDNSConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSConfigRequest.ProtoReflect.Descriptor instead.
func (*GetDNSConfigRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{0}
}

func (x *GetDNSConfigRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

type GetDNSConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *DNSConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetDNSConfigResponse) Reset() {
	*x = GetDNSConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSConfigResponse) ProtoMessage() {}

func (x *GetDNSConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSConfigResponse.ProtoReflect.Descriptor instead.
func (*GetDNSConfigResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{1}
}

func (x *GetDNSConfigResponse) GetConfig() *DNSConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type SetDNSConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TailnetId uint64     `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	Config    *DNSConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SetDNSConfigRequest) Reset() {
	*x = SetDNSConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSConfigRequest) ProtoMessage() {}

func (x *SetDNSConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSConfigRequest.ProtoReflect.Descriptor instead.
func (*SetDNSConfigRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{2}
}

func (x *SetDNSConfigRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

func (x *SetDNSConfigRequest) GetConfig() *DNSConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type SetDNSConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *DNSConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SetDNSConfigResponse) Reset() {
	*x = SetDNSConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSConfigResponse) ProtoMessage() {}

func (x *SetDNSConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSConfigResponse.ProtoReflect.Descriptor instead.
func (*SetDNSConfigResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{3}
}

func (x *SetDNSConfigResponse) GetConfig() *DNSConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type DNSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MagicDns         bool               `protobuf:"varint,1,opt,name=magic_dns,json=magicDns,proto3" json:"magic_dns,omitempty"`
	OverrideLocalDns bool               `protobuf:"varint,2,opt,name=override_local_dns,json=overrideLocalDns,proto3" json:"override_local_dns,omitempty"`
	Nameservers      []string           `protobuf:"bytes,3,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
	Routes           map[string]*Routes `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DNSConfig) Reset() {
	*x = DNSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSConfig) ProtoMessage() {}

func (x *DNSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSConfig.ProtoReflect.Descriptor instead.
func (*DNSConfig) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{4}
}

func (x *DNSConfig) GetMagicDns() bool {
	if x != nil {
		return x.MagicDns
	}
	return false
}

func (x *DNSConfig) GetOverrideLocalDns() bool {
	if x != nil {
		return x.OverrideLocalDns
	}
	return false
}

func (x *DNSConfig) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

func (x *DNSConfig) GetRoutes() map[string]*Routes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Routes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []string `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *Routes) Reset() {
	*x = Routes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routes) ProtoMessage() {}

func (x *Routes) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routes.ProtoReflect.Descriptor instead.
func (*Routes) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_dns_proto_rawDescGZIP(), []int{5}
}

func (x *Routes) GetRoutes() []string {
	if x != nil {
		return x.Routes
	}
	return nil
}

var File_ionscale_v1_dns_proto protoreflect.FileDescriptor

var file_ionscale_v1_dns_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x64, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x46, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x84, 0x02, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x5f, 0x64, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x44, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e,
	0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ionscale_v1_dns_proto_rawDescOnce sync.Once
	file_ionscale_v1_dns_proto_rawDescData = file_ionscale_v1_dns_proto_rawDesc
)

func file_ionscale_v1_dns_proto_rawDescGZIP() []byte {
	file_ionscale_v1_dns_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_ionscale_v1_dns_proto_rawDescData)
	})
	return file_ionscale_v1_dns_proto_rawDescData
}

var file_ionscale_v1_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ionscale_v1_dns_proto_goTypes = []interface{}{
	(*GetDNSConfigRequest)(nil),  // 0: ionscale.v1.GetDNSConfigRequest
	(*GetDNSConfigResponse)(nil), // 1: ionscale.v1.GetDNSConfigResponse
	(*SetDNSConfigRequest)(nil),  // 2: ionscale.v1.SetDNSConfigRequest
	(*SetDNSConfigResponse)(nil), // 3: ionscale.v1.SetDNSConfigResponse
	(*DNSConfig)(nil),            // 4: ionscale.v1.DNSConfig
	(*Routes)(nil),               // 5: ionscale.v1.Routes
	nil,                          // 6: ionscale.v1.DNSConfig.RoutesEntry
}
var file_ionscale_v1_dns_proto_depIdxs = []int32{
	4, // 0: ionscale.v1.GetDNSConfigResponse.config:type_name -> ionscale.v1.DNSConfig
	4, // 1: ionscale.v1.SetDNSConfigRequest.config:type_name -> ionscale.v1.DNSConfig
	4, // 2: ionscale.v1.SetDNSConfigResponse.config:type_name -> ionscale.v1.DNSConfig
	6, // 3: ionscale.v1.DNSConfig.routes:type_name -> ionscale.v1.DNSConfig.RoutesEntry
	5, // 4: ionscale.v1.DNSConfig.RoutesEntry.value:type_name -> ionscale.v1.Routes
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ionscale_v1_dns_proto_init() }
func file_ionscale_v1_dns_proto_init() {
	if File_ionscale_v1_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ionscale_v1_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSConfigRequest); i {
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
		file_ionscale_v1_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSConfigResponse); i {
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
		file_ionscale_v1_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSConfigRequest); i {
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
		file_ionscale_v1_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSConfigResponse); i {
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
		file_ionscale_v1_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSConfig); i {
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
		file_ionscale_v1_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routes); i {
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
			RawDescriptor: file_ionscale_v1_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_dns_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_dns_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_dns_proto_msgTypes,
	}.Build()
	File_ionscale_v1_dns_proto = out.File
	file_ionscale_v1_dns_proto_rawDesc = nil
	file_ionscale_v1_dns_proto_goTypes = nil
	file_ionscale_v1_dns_proto_depIdxs = nil
}