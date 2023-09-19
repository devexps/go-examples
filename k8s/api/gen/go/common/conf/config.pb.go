// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: common/conf/config.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RemoteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Consul *RemoteConfig_Consul `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`
	Etcd   *RemoteConfig_Etcd   `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *RemoteConfig) Reset() {
	*x = RemoteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig) ProtoMessage() {}

func (x *RemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig.ProtoReflect.Descriptor instead.
func (*RemoteConfig) Descriptor() ([]byte, []int) {
	return file_common_conf_config_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RemoteConfig) GetConsul() *RemoteConfig_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *RemoteConfig) GetEtcd() *RemoteConfig_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type RemoteConfig_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme  string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RemoteConfig_Consul) Reset() {
	*x = RemoteConfig_Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Consul) ProtoMessage() {}

func (x *RemoteConfig_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Consul.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Consul) Descriptor() ([]byte, []int) {
	return file_common_conf_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RemoteConfig_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *RemoteConfig_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RemoteConfig_Consul) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RemoteConfig_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Timeout   *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key       string               `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RemoteConfig_Etcd) Reset() {
	*x = RemoteConfig_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Etcd) ProtoMessage() {}

func (x *RemoteConfig_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Etcd.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Etcd) Descriptor() ([]byte, []int) {
	return file_common_conf_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RemoteConfig_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *RemoteConfig_Etcd) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RemoteConfig_Etcd) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_common_conf_config_proto protoreflect.FileDescriptor

var file_common_conf_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x1a, 0x4c, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0x6b, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x33,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_config_proto_rawDescOnce sync.Once
	file_common_conf_config_proto_rawDescData = file_common_conf_config_proto_rawDesc
)

func file_common_conf_config_proto_rawDescGZIP() []byte {
	file_common_conf_config_proto_rawDescOnce.Do(func() {
		file_common_conf_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_config_proto_rawDescData)
	})
	return file_common_conf_config_proto_rawDescData
}

var file_common_conf_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_conf_config_proto_goTypes = []interface{}{
	(*RemoteConfig)(nil),        // 0: common.conf.RemoteConfig
	(*RemoteConfig_Consul)(nil), // 1: common.conf.RemoteConfig.Consul
	(*RemoteConfig_Etcd)(nil),   // 2: common.conf.RemoteConfig.Etcd
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_common_conf_config_proto_depIdxs = []int32{
	1, // 0: common.conf.RemoteConfig.consul:type_name -> common.conf.RemoteConfig.Consul
	2, // 1: common.conf.RemoteConfig.etcd:type_name -> common.conf.RemoteConfig.Etcd
	3, // 2: common.conf.RemoteConfig.Etcd.timeout:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_conf_config_proto_init() }
func file_common_conf_config_proto_init() {
	if File_common_conf_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_conf_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig); i {
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
		file_common_conf_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Consul); i {
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
		file_common_conf_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Etcd); i {
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
			RawDescriptor: file_common_conf_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_config_proto_goTypes,
		DependencyIndexes: file_common_conf_config_proto_depIdxs,
		MessageInfos:      file_common_conf_config_proto_msgTypes,
	}.Build()
	File_common_conf_config_proto = out.File
	file_common_conf_config_proto_rawDesc = nil
	file_common_conf_config_proto_goTypes = nil
	file_common_conf_config_proto_depIdxs = nil
}
