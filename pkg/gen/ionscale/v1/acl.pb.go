// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ionscale/v1/acl.proto

package ionscalev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetACLPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TailnetId uint64 `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
}

func (x *GetACLPolicyRequest) Reset() {
	*x = GetACLPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLPolicyRequest) ProtoMessage() {}

func (x *GetACLPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetACLPolicyRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{0}
}

func (x *GetACLPolicyRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

type GetACLPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *ACLPolicy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *GetACLPolicyResponse) Reset() {
	*x = GetACLPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLPolicyResponse) ProtoMessage() {}

func (x *GetACLPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetACLPolicyResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{1}
}

func (x *GetACLPolicyResponse) GetPolicy() *ACLPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type SetACLPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TailnetId uint64     `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	Policy    *ACLPolicy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *SetACLPolicyRequest) Reset() {
	*x = SetACLPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetACLPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetACLPolicyRequest) ProtoMessage() {}

func (x *SetACLPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetACLPolicyRequest.ProtoReflect.Descriptor instead.
func (*SetACLPolicyRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{2}
}

func (x *SetACLPolicyRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

func (x *SetACLPolicyRequest) GetPolicy() *ACLPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type SetACLPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetACLPolicyResponse) Reset() {
	*x = SetACLPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetACLPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetACLPolicyResponse) ProtoMessage() {}

func (x *SetACLPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetACLPolicyResponse.ProtoReflect.Descriptor instead.
func (*SetACLPolicyResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{3}
}

type ACLPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts         map[string]string              `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Groups        map[string]*structpb.ListValue `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Acls          []*ACL                         `protobuf:"bytes,3,rep,name=acls,proto3" json:"acls,omitempty"`
	Tagowners     map[string]*structpb.ListValue `protobuf:"bytes,4,rep,name=tagowners,proto3" json:"tagowners,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Autoapprovers *AutoApprovers                 `protobuf:"bytes,5,opt,name=autoapprovers,proto3,oneof" json:"autoapprovers,omitempty"`
	Ssh           []*SSHRule                     `protobuf:"bytes,6,rep,name=ssh,proto3" json:"ssh,omitempty"`
	Nodeattrs     []*NodeAttr                    `protobuf:"bytes,7,rep,name=nodeattrs,proto3" json:"nodeattrs,omitempty"`
}

func (x *ACLPolicy) Reset() {
	*x = ACLPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLPolicy) ProtoMessage() {}

func (x *ACLPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLPolicy.ProtoReflect.Descriptor instead.
func (*ACLPolicy) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{4}
}

func (x *ACLPolicy) GetHosts() map[string]string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *ACLPolicy) GetGroups() map[string]*structpb.ListValue {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ACLPolicy) GetAcls() []*ACL {
	if x != nil {
		return x.Acls
	}
	return nil
}

func (x *ACLPolicy) GetTagowners() map[string]*structpb.ListValue {
	if x != nil {
		return x.Tagowners
	}
	return nil
}

func (x *ACLPolicy) GetAutoapprovers() *AutoApprovers {
	if x != nil {
		return x.Autoapprovers
	}
	return nil
}

func (x *ACLPolicy) GetSsh() []*SSHRule {
	if x != nil {
		return x.Ssh
	}
	return nil
}

func (x *ACLPolicy) GetNodeattrs() []*NodeAttr {
	if x != nil {
		return x.Nodeattrs
	}
	return nil
}

type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Src    []string `protobuf:"bytes,2,rep,name=src,proto3" json:"src,omitempty"`
	Dst    []string `protobuf:"bytes,3,rep,name=dst,proto3" json:"dst,omitempty"`
	Proto  string   `protobuf:"bytes,4,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{5}
}

func (x *ACL) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ACL) GetSrc() []string {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *ACL) GetDst() []string {
	if x != nil {
		return x.Dst
	}
	return nil
}

func (x *ACL) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

type AutoApprovers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes   map[string]*structpb.ListValue `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Exitnode []string                       `protobuf:"bytes,2,rep,name=exitnode,proto3" json:"exitnode,omitempty"`
}

func (x *AutoApprovers) Reset() {
	*x = AutoApprovers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoApprovers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoApprovers) ProtoMessage() {}

func (x *AutoApprovers) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoApprovers.ProtoReflect.Descriptor instead.
func (*AutoApprovers) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{6}
}

func (x *AutoApprovers) GetRoutes() map[string]*structpb.ListValue {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *AutoApprovers) GetExitnode() []string {
	if x != nil {
		return x.Exitnode
	}
	return nil
}

type SSHRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action      string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Src         []string `protobuf:"bytes,2,rep,name=src,proto3" json:"src,omitempty"`
	Dst         []string `protobuf:"bytes,3,rep,name=dst,proto3" json:"dst,omitempty"`
	Users       []string `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	Checkperiod string   `protobuf:"bytes,5,opt,name=checkperiod,proto3" json:"checkperiod,omitempty"`
}

func (x *SSHRule) Reset() {
	*x = SSHRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHRule) ProtoMessage() {}

func (x *SSHRule) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHRule.ProtoReflect.Descriptor instead.
func (*SSHRule) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{7}
}

func (x *SSHRule) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SSHRule) GetSrc() []string {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *SSHRule) GetDst() []string {
	if x != nil {
		return x.Dst
	}
	return nil
}

func (x *SSHRule) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SSHRule) GetCheckperiod() string {
	if x != nil {
		return x.Checkperiod
	}
	return ""
}

type NodeAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target []string `protobuf:"bytes,1,rep,name=target,proto3" json:"target,omitempty"`
	Attr   []string `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty"`
}

func (x *NodeAttr) Reset() {
	*x = NodeAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_acl_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAttr) ProtoMessage() {}

func (x *NodeAttr) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_acl_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAttr.ProtoReflect.Descriptor instead.
func (*NodeAttr) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_acl_proto_rawDescGZIP(), []int{8}
}

func (x *NodeAttr) GetTarget() []string {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *NodeAttr) GetAttr() []string {
	if x != nil {
		return x.Attr
	}
	return nil
}

var File_ionscale_v1_acl_proto protoreflect.FileDescriptor

var file_ionscale_v1_acl_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x69,
	0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x64, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x69, 0x6c, 0x6e,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x61, 0x69,
	0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8c,
	0x05, 0x0a, 0x09, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x37, 0x0a, 0x05,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43,
	0x4c, 0x52, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x54, 0x61, 0x67, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x74, 0x61, 0x67, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x0d,
	0x61, 0x75, 0x74, 0x6f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x48,
	0x00, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x03, 0x73, 0x73, 0x68, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x53, 0x48, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x03, 0x73, 0x73, 0x68, 0x12, 0x33, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x74, 0x74, 0x72, 0x73,
	0x1a, 0x38, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x0b, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x58, 0x0a, 0x0e, 0x54, 0x61, 0x67, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x22, 0x57, 0x0a,
	0x03, 0x41, 0x43, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x72, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x6f, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x55, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7d, 0x0a, 0x07, 0x53,
	0x53, 0x48, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x74,
	0x74, 0x72, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ionscale_v1_acl_proto_rawDescOnce sync.Once
	file_ionscale_v1_acl_proto_rawDescData = file_ionscale_v1_acl_proto_rawDesc
)

func file_ionscale_v1_acl_proto_rawDescGZIP() []byte {
	file_ionscale_v1_acl_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_ionscale_v1_acl_proto_rawDescData)
	})
	return file_ionscale_v1_acl_proto_rawDescData
}

var file_ionscale_v1_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_ionscale_v1_acl_proto_goTypes = []interface{}{
	(*GetACLPolicyRequest)(nil),  // 0: ionscale.v1.GetACLPolicyRequest
	(*GetACLPolicyResponse)(nil), // 1: ionscale.v1.GetACLPolicyResponse
	(*SetACLPolicyRequest)(nil),  // 2: ionscale.v1.SetACLPolicyRequest
	(*SetACLPolicyResponse)(nil), // 3: ionscale.v1.SetACLPolicyResponse
	(*ACLPolicy)(nil),            // 4: ionscale.v1.ACLPolicy
	(*ACL)(nil),                  // 5: ionscale.v1.ACL
	(*AutoApprovers)(nil),        // 6: ionscale.v1.AutoApprovers
	(*SSHRule)(nil),              // 7: ionscale.v1.SSHRule
	(*NodeAttr)(nil),             // 8: ionscale.v1.NodeAttr
	nil,                          // 9: ionscale.v1.ACLPolicy.HostsEntry
	nil,                          // 10: ionscale.v1.ACLPolicy.GroupsEntry
	nil,                          // 11: ionscale.v1.ACLPolicy.TagownersEntry
	nil,                          // 12: ionscale.v1.AutoApprovers.RoutesEntry
	(*structpb.ListValue)(nil),   // 13: google.protobuf.ListValue
}
var file_ionscale_v1_acl_proto_depIdxs = []int32{
	4,  // 0: ionscale.v1.GetACLPolicyResponse.policy:type_name -> ionscale.v1.ACLPolicy
	4,  // 1: ionscale.v1.SetACLPolicyRequest.policy:type_name -> ionscale.v1.ACLPolicy
	9,  // 2: ionscale.v1.ACLPolicy.hosts:type_name -> ionscale.v1.ACLPolicy.HostsEntry
	10, // 3: ionscale.v1.ACLPolicy.groups:type_name -> ionscale.v1.ACLPolicy.GroupsEntry
	5,  // 4: ionscale.v1.ACLPolicy.acls:type_name -> ionscale.v1.ACL
	11, // 5: ionscale.v1.ACLPolicy.tagowners:type_name -> ionscale.v1.ACLPolicy.TagownersEntry
	6,  // 6: ionscale.v1.ACLPolicy.autoapprovers:type_name -> ionscale.v1.AutoApprovers
	7,  // 7: ionscale.v1.ACLPolicy.ssh:type_name -> ionscale.v1.SSHRule
	8,  // 8: ionscale.v1.ACLPolicy.nodeattrs:type_name -> ionscale.v1.NodeAttr
	12, // 9: ionscale.v1.AutoApprovers.routes:type_name -> ionscale.v1.AutoApprovers.RoutesEntry
	13, // 10: ionscale.v1.ACLPolicy.GroupsEntry.value:type_name -> google.protobuf.ListValue
	13, // 11: ionscale.v1.ACLPolicy.TagownersEntry.value:type_name -> google.protobuf.ListValue
	13, // 12: ionscale.v1.AutoApprovers.RoutesEntry.value:type_name -> google.protobuf.ListValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_ionscale_v1_acl_proto_init() }
func file_ionscale_v1_acl_proto_init() {
	if File_ionscale_v1_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ionscale_v1_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLPolicyRequest); i {
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
		file_ionscale_v1_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLPolicyResponse); i {
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
		file_ionscale_v1_acl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetACLPolicyRequest); i {
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
		file_ionscale_v1_acl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetACLPolicyResponse); i {
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
		file_ionscale_v1_acl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLPolicy); i {
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
		file_ionscale_v1_acl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACL); i {
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
		file_ionscale_v1_acl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoApprovers); i {
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
		file_ionscale_v1_acl_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHRule); i {
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
		file_ionscale_v1_acl_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAttr); i {
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
	file_ionscale_v1_acl_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ionscale_v1_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_acl_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_acl_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_acl_proto_msgTypes,
	}.Build()
	File_ionscale_v1_acl_proto = out.File
	file_ionscale_v1_acl_proto_rawDesc = nil
	file_ionscale_v1_acl_proto_goTypes = nil
	file_ionscale_v1_acl_proto_depIdxs = nil
}
