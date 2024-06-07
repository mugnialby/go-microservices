// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: protobuf/user.proto

package user

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

type FindByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *FindByUsernameRequest) Reset() {
	*x = FindByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUsernameRequest) ProtoMessage() {}

func (x *FindByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUsernameRequest.ProtoReflect.Descriptor instead.
func (*FindByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{0}
}

func (x *FindByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FindByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *FindByUsernameResponse) Reset() {
	*x = FindByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUsernameResponse) ProtoMessage() {}

func (x *FindByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUsernameResponse.ProtoReflect.Descriptor instead.
func (*FindByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindByUsernameResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FindByUsernameResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_protobuf_user_proto protoreflect.FileDescriptor

var file_protobuf_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x15, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4a, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x5a, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_user_proto_rawDescOnce sync.Once
	file_protobuf_user_proto_rawDescData = file_protobuf_user_proto_rawDesc
)

func file_protobuf_user_proto_rawDescGZIP() []byte {
	file_protobuf_user_proto_rawDescOnce.Do(func() {
		file_protobuf_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_user_proto_rawDescData)
	})
	return file_protobuf_user_proto_rawDescData
}

var file_protobuf_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_user_proto_goTypes = []interface{}{
	(*FindByUsernameRequest)(nil),  // 0: user.FindByUsernameRequest
	(*FindByUsernameResponse)(nil), // 1: user.FindByUsernameResponse
}
var file_protobuf_user_proto_depIdxs = []int32{
	0, // 0: user.UserService.FindByUsername:input_type -> user.FindByUsernameRequest
	1, // 1: user.UserService.FindByUsername:output_type -> user.FindByUsernameResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_user_proto_init() }
func file_protobuf_user_proto_init() {
	if File_protobuf_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUsernameRequest); i {
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
		file_protobuf_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUsernameResponse); i {
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
			RawDescriptor: file_protobuf_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_user_proto_goTypes,
		DependencyIndexes: file_protobuf_user_proto_depIdxs,
		MessageInfos:      file_protobuf_user_proto_msgTypes,
	}.Build()
	File_protobuf_user_proto = out.File
	file_protobuf_user_proto_rawDesc = nil
	file_protobuf_user_proto_goTypes = nil
	file_protobuf_user_proto_depIdxs = nil
}
