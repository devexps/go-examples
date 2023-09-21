// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: common/conf/logger.proto

package conf

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

type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Zap    *Logger_Zap    `protobuf:"bytes,2,opt,name=zap,proto3" json:"zap,omitempty"`
	Logrus *Logger_Logrus `protobuf:"bytes,3,opt,name=logrus,proto3" json:"logrus,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Logger) GetZap() *Logger_Zap {
	if x != nil {
		return x.Zap
	}
	return nil
}

func (x *Logger) GetLogrus() *Logger_Logrus {
	if x != nil {
		return x.Logrus
	}
	return nil
}

// Zap
type Logger_Zap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Level      string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	MaxSize    int32  `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	MaxAge     int32  `protobuf:"varint,4,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	MaxBackups int32  `protobuf:"varint,5,opt,name=max_backups,json=maxBackups,proto3" json:"max_backups,omitempty"`
}

func (x *Logger_Zap) Reset() {
	*x = Logger_Zap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Zap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Zap) ProtoMessage() {}

func (x *Logger_Zap) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Zap.ProtoReflect.Descriptor instead.
func (*Logger_Zap) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Logger_Zap) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Logger_Zap) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Zap) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Logger_Zap) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Logger_Zap) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

// Logrus
type Logger_Logrus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level            string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`                                                // log level
	Formatter        string `protobuf:"bytes,2,opt,name=formatter,proto3" json:"formatter,omitempty"`                                        // output format：text, json.
	TimestampFormat  string `protobuf:"bytes,3,opt,name=timestamp_format,json=timestampFormat,proto3" json:"timestamp_format,omitempty"`     // define the timestamp format, e.g. "2006-01-02 15:04:05"
	DisableColors    bool   `protobuf:"varint,4,opt,name=disable_colors,json=disableColors,proto3" json:"disable_colors,omitempty"`          // no need for colored logs
	DisableTimestamp bool   `protobuf:"varint,5,opt,name=disable_timestamp,json=disableTimestamp,proto3" json:"disable_timestamp,omitempty"` // no timestamp required
}

func (x *Logger_Logrus) Reset() {
	*x = Logger_Logrus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_logger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Logrus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Logrus) ProtoMessage() {}

func (x *Logger_Logrus) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Logrus.ProtoReflect.Descriptor instead.
func (*Logger_Logrus) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Logger_Logrus) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Logrus) GetFormatter() string {
	if x != nil {
		return x.Formatter
	}
	return ""
}

func (x *Logger_Logrus) GetTimestampFormat() string {
	if x != nil {
		return x.TimestampFormat
	}
	return ""
}

func (x *Logger_Logrus) GetDisableColors() bool {
	if x != nil {
		return x.DisableColors
	}
	return false
}

func (x *Logger_Logrus) GetDisableTimestamp() bool {
	if x != nil {
		return x.DisableTimestamp
	}
	return false
}

var File_common_conf_logger_proto protoreflect.FileDescriptor

var file_common_conf_logger_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0xc8, 0x03, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x7a, 0x61, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x5a, 0x61, 0x70, 0x52, 0x03, 0x7a, 0x61,
	0x70, 0x12, 0x32, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x52, 0x06, 0x6c,
	0x6f, 0x67, 0x72, 0x75, 0x73, 0x1a, 0x8c, 0x01, 0x0a, 0x03, 0x5a, 0x61, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78,
	0x41, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x1a, 0xbb, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x65, 0x76, 0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_logger_proto_rawDescOnce sync.Once
	file_common_conf_logger_proto_rawDescData = file_common_conf_logger_proto_rawDesc
)

func file_common_conf_logger_proto_rawDescGZIP() []byte {
	file_common_conf_logger_proto_rawDescOnce.Do(func() {
		file_common_conf_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_logger_proto_rawDescData)
	})
	return file_common_conf_logger_proto_rawDescData
}

var file_common_conf_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_conf_logger_proto_goTypes = []interface{}{
	(*Logger)(nil),        // 0: common.conf.Logger
	(*Logger_Zap)(nil),    // 1: common.conf.Logger.Zap
	(*Logger_Logrus)(nil), // 2: common.conf.Logger.Logrus
}
var file_common_conf_logger_proto_depIdxs = []int32{
	1, // 0: common.conf.Logger.zap:type_name -> common.conf.Logger.Zap
	2, // 1: common.conf.Logger.logrus:type_name -> common.conf.Logger.Logrus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_conf_logger_proto_init() }
func file_common_conf_logger_proto_init() {
	if File_common_conf_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_conf_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
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
		file_common_conf_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Zap); i {
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
		file_common_conf_logger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Logrus); i {
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
			RawDescriptor: file_common_conf_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_logger_proto_goTypes,
		DependencyIndexes: file_common_conf_logger_proto_depIdxs,
		MessageInfos:      file_common_conf_logger_proto_msgTypes,
	}.Build()
	File_common_conf_logger_proto = out.File
	file_common_conf_logger_proto_rawDesc = nil
	file_common_conf_logger_proto_goTypes = nil
	file_common_conf_logger_proto_depIdxs = nil
}
