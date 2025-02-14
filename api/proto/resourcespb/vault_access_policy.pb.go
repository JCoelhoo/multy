// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/resourcespb/vault_access_policy.proto

package resourcespb

import (
	commonpb "github.com/multycloud/multy/api/proto/commonpb"
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

type VaultAccess_Enum int32

const (
	VaultAccess_UNKNOWN VaultAccess_Enum = 0
	VaultAccess_READ    VaultAccess_Enum = 1
	VaultAccess_WRITE   VaultAccess_Enum = 2
	VaultAccess_OWNER   VaultAccess_Enum = 3
)

// Enum value maps for VaultAccess_Enum.
var (
	VaultAccess_Enum_name = map[int32]string{
		0: "UNKNOWN",
		1: "READ",
		2: "WRITE",
		3: "OWNER",
	}
	VaultAccess_Enum_value = map[string]int32{
		"UNKNOWN": 0,
		"READ":    1,
		"WRITE":   2,
		"OWNER":   3,
	}
)

func (x VaultAccess_Enum) Enum() *VaultAccess_Enum {
	p := new(VaultAccess_Enum)
	*p = x
	return p
}

func (x VaultAccess_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VaultAccess_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_resourcespb_vault_access_policy_proto_enumTypes[0].Descriptor()
}

func (VaultAccess_Enum) Type() protoreflect.EnumType {
	return &file_api_proto_resourcespb_vault_access_policy_proto_enumTypes[0]
}

func (x VaultAccess_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VaultAccess_Enum.Descriptor instead.
func (VaultAccess_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{4, 0}
}

type CreateVaultAccessPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *VaultAccessPolicyArgs `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateVaultAccessPolicyRequest) Reset() {
	*x = CreateVaultAccessPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVaultAccessPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVaultAccessPolicyRequest) ProtoMessage() {}

func (x *CreateVaultAccessPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVaultAccessPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateVaultAccessPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVaultAccessPolicyRequest) GetResource() *VaultAccessPolicyArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ReadVaultAccessPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadVaultAccessPolicyRequest) Reset() {
	*x = ReadVaultAccessPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadVaultAccessPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadVaultAccessPolicyRequest) ProtoMessage() {}

func (x *ReadVaultAccessPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadVaultAccessPolicyRequest.ProtoReflect.Descriptor instead.
func (*ReadVaultAccessPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{1}
}

func (x *ReadVaultAccessPolicyRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateVaultAccessPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string                 `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resource   *VaultAccessPolicyArgs `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateVaultAccessPolicyRequest) Reset() {
	*x = UpdateVaultAccessPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVaultAccessPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVaultAccessPolicyRequest) ProtoMessage() {}

func (x *UpdateVaultAccessPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVaultAccessPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateVaultAccessPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVaultAccessPolicyRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateVaultAccessPolicyRequest) GetResource() *VaultAccessPolicyArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteVaultAccessPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteVaultAccessPolicyRequest) Reset() {
	*x = DeleteVaultAccessPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVaultAccessPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVaultAccessPolicyRequest) ProtoMessage() {}

func (x *DeleteVaultAccessPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVaultAccessPolicyRequest.ProtoReflect.Descriptor instead.
func (*DeleteVaultAccessPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteVaultAccessPolicyRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type VaultAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VaultAccess) Reset() {
	*x = VaultAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultAccess) ProtoMessage() {}

func (x *VaultAccess) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultAccess.ProtoReflect.Descriptor instead.
func (*VaultAccess) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{4}
}

type VaultAccessPolicyArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.ChildResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	VaultId          string                            `protobuf:"bytes,2,opt,name=vault_id,json=vaultId,proto3" json:"vault_id,omitempty"`
	Identity         string                            `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Access           VaultAccess_Enum                  `protobuf:"varint,4,opt,name=access,proto3,enum=dev.multy.resources.VaultAccess_Enum" json:"access,omitempty"`
}

func (x *VaultAccessPolicyArgs) Reset() {
	*x = VaultAccessPolicyArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultAccessPolicyArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultAccessPolicyArgs) ProtoMessage() {}

func (x *VaultAccessPolicyArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultAccessPolicyArgs.ProtoReflect.Descriptor instead.
func (*VaultAccessPolicyArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{5}
}

func (x *VaultAccessPolicyArgs) GetCommonParameters() *commonpb.ChildResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *VaultAccessPolicyArgs) GetVaultId() string {
	if x != nil {
		return x.VaultId
	}
	return ""
}

func (x *VaultAccessPolicyArgs) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *VaultAccessPolicyArgs) GetAccess() VaultAccess_Enum {
	if x != nil {
		return x.Access
	}
	return VaultAccess_UNKNOWN
}

type VaultAccessPolicyResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.CommonChildResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	VaultId          string                                  `protobuf:"bytes,2,opt,name=vault_id,json=vaultId,proto3" json:"vault_id,omitempty"`
	Identity         string                                  `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Access           VaultAccess_Enum                        `protobuf:"varint,4,opt,name=access,proto3,enum=dev.multy.resources.VaultAccess_Enum" json:"access,omitempty"`
}

func (x *VaultAccessPolicyResource) Reset() {
	*x = VaultAccessPolicyResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultAccessPolicyResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultAccessPolicyResource) ProtoMessage() {}

func (x *VaultAccessPolicyResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultAccessPolicyResource.ProtoReflect.Descriptor instead.
func (*VaultAccessPolicyResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP(), []int{6}
}

func (x *VaultAccessPolicyResource) GetCommonParameters() *commonpb.CommonChildResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *VaultAccessPolicyResource) GetVaultId() string {
	if x != nil {
		return x.VaultId
	}
	return ""
}

func (x *VaultAccessPolicyResource) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *VaultAccessPolicyResource) GetAccess() VaultAccess_Enum {
	if x != nil {
		return x.Access
	}
	return VaultAccess_UNKNOWN
}

var File_api_proto_resourcespb_vault_access_policy_proto protoreflect.FileDescriptor

var file_api_proto_resourcespb_vault_access_policy_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x41, 0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x3f, 0x0a, 0x1c, 0x52, 0x65, 0x61, 0x64, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x41,
	0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x42, 0x0a, 0x0b, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x33, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57,
	0x4e, 0x45, 0x52, 0x10, 0x03, 0x22, 0xe5, 0x01, 0x0a, 0x15, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3d,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xef, 0x01,
	0x0a, 0x19, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x5e, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resourcespb_vault_access_policy_proto_rawDescOnce sync.Once
	file_api_proto_resourcespb_vault_access_policy_proto_rawDescData = file_api_proto_resourcespb_vault_access_policy_proto_rawDesc
)

func file_api_proto_resourcespb_vault_access_policy_proto_rawDescGZIP() []byte {
	file_api_proto_resourcespb_vault_access_policy_proto_rawDescOnce.Do(func() {
		file_api_proto_resourcespb_vault_access_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resourcespb_vault_access_policy_proto_rawDescData)
	})
	return file_api_proto_resourcespb_vault_access_policy_proto_rawDescData
}

var file_api_proto_resourcespb_vault_access_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_resourcespb_vault_access_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_resourcespb_vault_access_policy_proto_goTypes = []interface{}{
	(VaultAccess_Enum)(0),                          // 0: dev.multy.resources.VaultAccess.Enum
	(*CreateVaultAccessPolicyRequest)(nil),         // 1: dev.multy.resources.CreateVaultAccessPolicyRequest
	(*ReadVaultAccessPolicyRequest)(nil),           // 2: dev.multy.resources.ReadVaultAccessPolicyRequest
	(*UpdateVaultAccessPolicyRequest)(nil),         // 3: dev.multy.resources.UpdateVaultAccessPolicyRequest
	(*DeleteVaultAccessPolicyRequest)(nil),         // 4: dev.multy.resources.DeleteVaultAccessPolicyRequest
	(*VaultAccess)(nil),                            // 5: dev.multy.resources.VaultAccess
	(*VaultAccessPolicyArgs)(nil),                  // 6: dev.multy.resources.VaultAccessPolicyArgs
	(*VaultAccessPolicyResource)(nil),              // 7: dev.multy.resources.VaultAccessPolicyResource
	(*commonpb.ChildResourceCommonArgs)(nil),       // 8: dev.multy.common.ChildResourceCommonArgs
	(*commonpb.CommonChildResourceParameters)(nil), // 9: dev.multy.common.CommonChildResourceParameters
}
var file_api_proto_resourcespb_vault_access_policy_proto_depIdxs = []int32{
	6, // 0: dev.multy.resources.CreateVaultAccessPolicyRequest.resource:type_name -> dev.multy.resources.VaultAccessPolicyArgs
	6, // 1: dev.multy.resources.UpdateVaultAccessPolicyRequest.resource:type_name -> dev.multy.resources.VaultAccessPolicyArgs
	8, // 2: dev.multy.resources.VaultAccessPolicyArgs.common_parameters:type_name -> dev.multy.common.ChildResourceCommonArgs
	0, // 3: dev.multy.resources.VaultAccessPolicyArgs.access:type_name -> dev.multy.resources.VaultAccess.Enum
	9, // 4: dev.multy.resources.VaultAccessPolicyResource.common_parameters:type_name -> dev.multy.common.CommonChildResourceParameters
	0, // 5: dev.multy.resources.VaultAccessPolicyResource.access:type_name -> dev.multy.resources.VaultAccess.Enum
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_resourcespb_vault_access_policy_proto_init() }
func file_api_proto_resourcespb_vault_access_policy_proto_init() {
	if File_api_proto_resourcespb_vault_access_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVaultAccessPolicyRequest); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadVaultAccessPolicyRequest); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVaultAccessPolicyRequest); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVaultAccessPolicyRequest); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultAccess); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultAccessPolicyArgs); i {
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
		file_api_proto_resourcespb_vault_access_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultAccessPolicyResource); i {
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
			RawDescriptor: file_api_proto_resourcespb_vault_access_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resourcespb_vault_access_policy_proto_goTypes,
		DependencyIndexes: file_api_proto_resourcespb_vault_access_policy_proto_depIdxs,
		EnumInfos:         file_api_proto_resourcespb_vault_access_policy_proto_enumTypes,
		MessageInfos:      file_api_proto_resourcespb_vault_access_policy_proto_msgTypes,
	}.Build()
	File_api_proto_resourcespb_vault_access_policy_proto = out.File
	file_api_proto_resourcespb_vault_access_policy_proto_rawDesc = nil
	file_api_proto_resourcespb_vault_access_policy_proto_goTypes = nil
	file_api_proto_resourcespb_vault_access_policy_proto_depIdxs = nil
}
