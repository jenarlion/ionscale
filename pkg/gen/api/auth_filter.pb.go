// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/auth_filter.proto

package api

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

type AuthFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthMethod *Ref   `protobuf:"bytes,2,opt,name=auth_method,json=authMethod,proto3" json:"auth_method,omitempty"`
	Tailnet    *Ref   `protobuf:"bytes,3,opt,name=tailnet,proto3" json:"tailnet,omitempty"`
	Expr       string `protobuf:"bytes,4,opt,name=expr,proto3" json:"expr,omitempty"`
}

func (x *AuthFilter) Reset() {
	*x = AuthFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthFilter) ProtoMessage() {}

func (x *AuthFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthFilter.ProtoReflect.Descriptor instead.
func (*AuthFilter) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{0}
}

func (x *AuthFilter) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthFilter) GetAuthMethod() *Ref {
	if x != nil {
		return x.AuthMethod
	}
	return nil
}

func (x *AuthFilter) GetTailnet() *Ref {
	if x != nil {
		return x.Tailnet
	}
	return nil
}

func (x *AuthFilter) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

type CreateAuthFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthMethodId uint64 `protobuf:"varint,1,opt,name=auth_method_id,json=authMethodId,proto3" json:"auth_method_id,omitempty"`
	TailnetId    uint64 `protobuf:"varint,2,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	Expr         string `protobuf:"bytes,3,opt,name=expr,proto3" json:"expr,omitempty"`
}

func (x *CreateAuthFilterRequest) Reset() {
	*x = CreateAuthFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthFilterRequest) ProtoMessage() {}

func (x *CreateAuthFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthFilterRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuthFilterRequest) GetAuthMethodId() uint64 {
	if x != nil {
		return x.AuthMethodId
	}
	return 0
}

func (x *CreateAuthFilterRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

func (x *CreateAuthFilterRequest) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

type CreateAuthFilterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthFilter *AuthFilter `protobuf:"bytes,1,opt,name=auth_filter,json=authFilter,proto3" json:"auth_filter,omitempty"`
}

func (x *CreateAuthFilterResponse) Reset() {
	*x = CreateAuthFilterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthFilterResponse) ProtoMessage() {}

func (x *CreateAuthFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthFilterResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthFilterResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthFilterResponse) GetAuthFilter() *AuthFilter {
	if x != nil {
		return x.AuthFilter
	}
	return nil
}

type ListAuthFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthMethodId *uint64 `protobuf:"varint,1,opt,name=auth_method_id,json=authMethodId,proto3,oneof" json:"auth_method_id,omitempty"`
}

func (x *ListAuthFiltersRequest) Reset() {
	*x = ListAuthFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthFiltersRequest) ProtoMessage() {}

func (x *ListAuthFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListAuthFiltersRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{3}
}

func (x *ListAuthFiltersRequest) GetAuthMethodId() uint64 {
	if x != nil && x.AuthMethodId != nil {
		return *x.AuthMethodId
	}
	return 0
}

type ListAuthFiltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthFilters []*AuthFilter `protobuf:"bytes,1,rep,name=auth_filters,json=authFilters,proto3" json:"auth_filters,omitempty"`
}

func (x *ListAuthFiltersResponse) Reset() {
	*x = ListAuthFiltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthFiltersResponse) ProtoMessage() {}

func (x *ListAuthFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthFiltersResponse.ProtoReflect.Descriptor instead.
func (*ListAuthFiltersResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{4}
}

func (x *ListAuthFiltersResponse) GetAuthFilters() []*AuthFilter {
	if x != nil {
		return x.AuthFilters
	}
	return nil
}

type DeleteAuthFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthFilterId uint64 `protobuf:"varint,1,opt,name=auth_filter_id,json=authFilterId,proto3" json:"auth_filter_id,omitempty"`
}

func (x *DeleteAuthFilterRequest) Reset() {
	*x = DeleteAuthFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthFilterRequest) ProtoMessage() {}

func (x *DeleteAuthFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthFilterRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAuthFilterRequest) GetAuthFilterId() uint64 {
	if x != nil {
		return x.AuthFilterId
	}
	return 0
}

type DeleteAuthFilterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAuthFilterResponse) Reset() {
	*x = DeleteAuthFilterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_filter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthFilterResponse) ProtoMessage() {}

func (x *DeleteAuthFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_filter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthFilterResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthFilterResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_filter_proto_rawDescGZIP(), []int{6}
}

var File_api_auth_filter_proto protoreflect.FileDescriptor

var file_api_auth_filter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0d, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0a, 0x41,
	0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x52,
	0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x22, 0x72, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72,
	0x22, 0x4c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x56,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_filter_proto_rawDescOnce sync.Once
	file_api_auth_filter_proto_rawDescData = file_api_auth_filter_proto_rawDesc
)

func file_api_auth_filter_proto_rawDescGZIP() []byte {
	file_api_auth_filter_proto_rawDescOnce.Do(func() {
		file_api_auth_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_filter_proto_rawDescData)
	})
	return file_api_auth_filter_proto_rawDescData
}

var file_api_auth_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_auth_filter_proto_goTypes = []interface{}{
	(*AuthFilter)(nil),               // 0: api.AuthFilter
	(*CreateAuthFilterRequest)(nil),  // 1: api.CreateAuthFilterRequest
	(*CreateAuthFilterResponse)(nil), // 2: api.CreateAuthFilterResponse
	(*ListAuthFiltersRequest)(nil),   // 3: api.ListAuthFiltersRequest
	(*ListAuthFiltersResponse)(nil),  // 4: api.ListAuthFiltersResponse
	(*DeleteAuthFilterRequest)(nil),  // 5: api.DeleteAuthFilterRequest
	(*DeleteAuthFilterResponse)(nil), // 6: api.DeleteAuthFilterResponse
	(*Ref)(nil),                      // 7: api.Ref
}
var file_api_auth_filter_proto_depIdxs = []int32{
	7, // 0: api.AuthFilter.auth_method:type_name -> api.Ref
	7, // 1: api.AuthFilter.tailnet:type_name -> api.Ref
	0, // 2: api.CreateAuthFilterResponse.auth_filter:type_name -> api.AuthFilter
	0, // 3: api.ListAuthFiltersResponse.auth_filters:type_name -> api.AuthFilter
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_auth_filter_proto_init() }
func file_api_auth_filter_proto_init() {
	if File_api_auth_filter_proto != nil {
		return
	}
	file_api_ref_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_auth_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthFilter); i {
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
		file_api_auth_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthFilterRequest); i {
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
		file_api_auth_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthFilterResponse); i {
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
		file_api_auth_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthFiltersRequest); i {
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
		file_api_auth_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthFiltersResponse); i {
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
		file_api_auth_filter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthFilterRequest); i {
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
		file_api_auth_filter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthFilterResponse); i {
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
	file_api_auth_filter_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_auth_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_filter_proto_goTypes,
		DependencyIndexes: file_api_auth_filter_proto_depIdxs,
		MessageInfos:      file_api_auth_filter_proto_msgTypes,
	}.Build()
	File_api_auth_filter_proto = out.File
	file_api_auth_filter_proto_rawDesc = nil
	file_api_auth_filter_proto_goTypes = nil
	file_api_auth_filter_proto_depIdxs = nil
}
