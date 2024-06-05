//*
// API to manage members of an organization.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: priv/members/v1alpha1/members.proto

package membersv1alpha1

import (
	v1alpha1 "go.jetpack.io/pkg/api/gen/priv/organizations/v1alpha1"
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

// The member object
type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the object
	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Organization *v1alpha1.Organization `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_members_v1alpha1_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_priv_members_v1alpha1_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_priv_members_v1alpha1_members_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetOrganization() *v1alpha1.Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

type GetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the member to retrieve.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemberRequest) Reset() {
	*x = GetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_members_v1alpha1_members_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberRequest) ProtoMessage() {}

func (x *GetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_priv_members_v1alpha1_members_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberRequest.ProtoReflect.Descriptor instead.
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return file_priv_members_v1alpha1_members_proto_rawDescGZIP(), []int{1}
}

func (x *GetMemberRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requested member object.
	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *GetMemberResponse) Reset() {
	*x = GetMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_members_v1alpha1_members_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberResponse) ProtoMessage() {}

func (x *GetMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_priv_members_v1alpha1_members_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberResponse.ProtoReflect.Descriptor instead.
func (*GetMemberResponse) Descriptor() ([]byte, []int) {
	return file_priv_members_v1alpha1_members_proto_rawDescGZIP(), []int{2}
}

func (x *GetMemberResponse) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

var File_priv_members_v1alpha1_members_proto protoreflect.FileDescriptor

var file_priv_members_v1alpha1_members_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x75, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0xe0, 0x01,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x6f, 0x2e,
	0x6a, 0x65, 0x74, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x4d, 0x58, 0xaa, 0x02, 0x15, 0x50, 0x72, 0x69, 0x76, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x50, 0x72, 0x69,
	0x76, 0x5c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x21, 0x50, 0x72, 0x69, 0x76, 0x5c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x50, 0x72, 0x69, 0x76, 0x3a, 0x3a, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_priv_members_v1alpha1_members_proto_rawDescOnce sync.Once
	file_priv_members_v1alpha1_members_proto_rawDescData = file_priv_members_v1alpha1_members_proto_rawDesc
)

func file_priv_members_v1alpha1_members_proto_rawDescGZIP() []byte {
	file_priv_members_v1alpha1_members_proto_rawDescOnce.Do(func() {
		file_priv_members_v1alpha1_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_priv_members_v1alpha1_members_proto_rawDescData)
	})
	return file_priv_members_v1alpha1_members_proto_rawDescData
}

var file_priv_members_v1alpha1_members_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_priv_members_v1alpha1_members_proto_goTypes = []interface{}{
	(*Member)(nil),                // 0: priv.members.v1alpha1.Member
	(*GetMemberRequest)(nil),      // 1: priv.members.v1alpha1.GetMemberRequest
	(*GetMemberResponse)(nil),     // 2: priv.members.v1alpha1.GetMemberResponse
	(*v1alpha1.Organization)(nil), // 3: priv.organizations.v1alpha1.Organization
}
var file_priv_members_v1alpha1_members_proto_depIdxs = []int32{
	3, // 0: priv.members.v1alpha1.Member.organization:type_name -> priv.organizations.v1alpha1.Organization
	0, // 1: priv.members.v1alpha1.GetMemberResponse.member:type_name -> priv.members.v1alpha1.Member
	1, // 2: priv.members.v1alpha1.MembersService.GetMember:input_type -> priv.members.v1alpha1.GetMemberRequest
	2, // 3: priv.members.v1alpha1.MembersService.GetMember:output_type -> priv.members.v1alpha1.GetMemberResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_priv_members_v1alpha1_members_proto_init() }
func file_priv_members_v1alpha1_members_proto_init() {
	if File_priv_members_v1alpha1_members_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_priv_members_v1alpha1_members_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_priv_members_v1alpha1_members_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberRequest); i {
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
		file_priv_members_v1alpha1_members_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberResponse); i {
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
			RawDescriptor: file_priv_members_v1alpha1_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_priv_members_v1alpha1_members_proto_goTypes,
		DependencyIndexes: file_priv_members_v1alpha1_members_proto_depIdxs,
		MessageInfos:      file_priv_members_v1alpha1_members_proto_msgTypes,
	}.Build()
	File_priv_members_v1alpha1_members_proto = out.File
	file_priv_members_v1alpha1_members_proto_rawDesc = nil
	file_priv_members_v1alpha1_members_proto_goTypes = nil
	file_priv_members_v1alpha1_members_proto_depIdxs = nil
}
