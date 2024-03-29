// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: token-srv/v1/token.proto

package v1

import (
	authn "github.com/devexps/go-examples/k8s/api/gen/go/common/authn"
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

type TokenPlatform int32

const (
	TokenPlatform_TOKEN_PLATFORM_UNSPECIFIED TokenPlatform = 0
	TokenPlatform_TOKEN_PLATFORM_WEB         TokenPlatform = 1
	TokenPlatform_TOKEN_PLATFORM_ANDROID     TokenPlatform = 2
	TokenPlatform_TOKEN_PLATFORM_IOS         TokenPlatform = 3
)

// Enum value maps for TokenPlatform.
var (
	TokenPlatform_name = map[int32]string{
		0: "TOKEN_PLATFORM_UNSPECIFIED",
		1: "TOKEN_PLATFORM_WEB",
		2: "TOKEN_PLATFORM_ANDROID",
		3: "TOKEN_PLATFORM_IOS",
	}
	TokenPlatform_value = map[string]int32{
		"TOKEN_PLATFORM_UNSPECIFIED": 0,
		"TOKEN_PLATFORM_WEB":         1,
		"TOKEN_PLATFORM_ANDROID":     2,
		"TOKEN_PLATFORM_IOS":         3,
	}
)

func (x TokenPlatform) Enum() *TokenPlatform {
	p := new(TokenPlatform)
	*p = x
	return p
}

func (x TokenPlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenPlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_token_srv_v1_token_proto_enumTypes[0].Descriptor()
}

func (TokenPlatform) Type() protoreflect.EnumType {
	return &file_token_srv_v1_token_proto_enumTypes[0]
}

func (x TokenPlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenPlatform.Descriptor instead.
func (TokenPlatform) EnumDescriptor() ([]byte, []int) {
	return file_token_srv_v1_token_proto_rawDescGZIP(), []int{0}
}

type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenInfo *authn.Claims `protobuf:"bytes,1,opt,name=tokenInfo,proto3" json:"tokenInfo,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_srv_v1_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_srv_v1_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_srv_v1_token_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetTokenInfo() *authn.Claims {
	if x != nil {
		return x.TokenInfo
	}
	return nil
}

type GenerateTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenerateTokenReply) Reset() {
	*x = GenerateTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_srv_v1_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenReply) ProtoMessage() {}

func (x *GenerateTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_token_srv_v1_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenReply.ProtoReflect.Descriptor instead.
func (*GenerateTokenReply) Descriptor() ([]byte, []int) {
	return file_token_srv_v1_token_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_token_srv_v1_token_proto protoreflect.FileDescriptor

var file_token_srv_v1_token_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2a, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x7b, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x1a,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57,
	0x45, 0x42, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x03, 0x32, 0x65, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x5f, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x76, 0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_srv_v1_token_proto_rawDescOnce sync.Once
	file_token_srv_v1_token_proto_rawDescData = file_token_srv_v1_token_proto_rawDesc
)

func file_token_srv_v1_token_proto_rawDescGZIP() []byte {
	file_token_srv_v1_token_proto_rawDescOnce.Do(func() {
		file_token_srv_v1_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_srv_v1_token_proto_rawDescData)
	})
	return file_token_srv_v1_token_proto_rawDescData
}

var file_token_srv_v1_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_token_srv_v1_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_token_srv_v1_token_proto_goTypes = []interface{}{
	(TokenPlatform)(0),           // 0: token_srv.v1.TokenPlatform
	(*GenerateTokenRequest)(nil), // 1: token_srv.v1.GenerateTokenRequest
	(*GenerateTokenReply)(nil),   // 2: token_srv.v1.GenerateTokenReply
	(*authn.Claims)(nil),         // 3: common.authn.Claims
}
var file_token_srv_v1_token_proto_depIdxs = []int32{
	3, // 0: token_srv.v1.GenerateTokenRequest.tokenInfo:type_name -> common.authn.Claims
	1, // 1: token_srv.v1.TokenService.GenerateToken:input_type -> token_srv.v1.GenerateTokenRequest
	2, // 2: token_srv.v1.TokenService.GenerateToken:output_type -> token_srv.v1.GenerateTokenReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_token_srv_v1_token_proto_init() }
func file_token_srv_v1_token_proto_init() {
	if File_token_srv_v1_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_srv_v1_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_token_srv_v1_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenReply); i {
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
			RawDescriptor: file_token_srv_v1_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_srv_v1_token_proto_goTypes,
		DependencyIndexes: file_token_srv_v1_token_proto_depIdxs,
		EnumInfos:         file_token_srv_v1_token_proto_enumTypes,
		MessageInfos:      file_token_srv_v1_token_proto_msgTypes,
	}.Build()
	File_token_srv_v1_token_proto = out.File
	file_token_srv_v1_token_proto_rawDesc = nil
	file_token_srv_v1_token_proto_goTypes = nil
	file_token_srv_v1_token_proto_depIdxs = nil
}
