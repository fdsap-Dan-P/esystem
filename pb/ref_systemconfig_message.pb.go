// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: ref_systemconfig_message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OfficeId     int64                  `protobuf:"varint,2,opt,name=officeId,proto3" json:"officeId,omitempty"`
	GlDate       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=glDate,proto3" json:"glDate,omitempty"`
	LastAccruals *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lastAccruals,proto3" json:"lastAccruals,omitempty"`
	LastMonthEnd *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=lastMonthEnd,proto3" json:"lastMonthEnd,omitempty"`
	NextMonthEnd *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=nextMonthEnd,proto3" json:"nextMonthEnd,omitempty"`
	SystemDate   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=systemDate,proto3" json:"systemDate,omitempty"`
	RunState     int32                  `protobuf:"varint,8,opt,name=runState,proto3" json:"runState,omitempty"`
	OtherInfo    *NullString            `protobuf:"bytes,9,opt,name=otherInfo,proto3" json:"otherInfo,omitempty"`
}

func (x *SystemConfig) Reset() {
	*x = SystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ref_systemconfig_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemConfig) ProtoMessage() {}

func (x *SystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ref_systemconfig_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemConfig.ProtoReflect.Descriptor instead.
func (*SystemConfig) Descriptor() ([]byte, []int) {
	return file_ref_systemconfig_message_proto_rawDescGZIP(), []int{0}
}

func (x *SystemConfig) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SystemConfig) GetOfficeId() int64 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *SystemConfig) GetGlDate() *timestamppb.Timestamp {
	if x != nil {
		return x.GlDate
	}
	return nil
}

func (x *SystemConfig) GetLastAccruals() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAccruals
	}
	return nil
}

func (x *SystemConfig) GetLastMonthEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.LastMonthEnd
	}
	return nil
}

func (x *SystemConfig) GetNextMonthEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.NextMonthEnd
	}
	return nil
}

func (x *SystemConfig) GetSystemDate() *timestamppb.Timestamp {
	if x != nil {
		return x.SystemDate
	}
	return nil
}

func (x *SystemConfig) GetRunState() int32 {
	if x != nil {
		return x.RunState
	}
	return 0
}

func (x *SystemConfig) GetOtherInfo() *NullString {
	if x != nil {
		return x.OtherInfo
	}
	return nil
}

var File_ref_systemconfig_message_proto protoreflect.FileDescriptor

var file_ref_systemconfig_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x66, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x67, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x67, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x72, 0x75, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x72, 0x75, 0x61, 0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x4e,
	0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x20, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x64, 0x73, 0x61,
	0x70, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x62, 0x50,
	0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ref_systemconfig_message_proto_rawDescOnce sync.Once
	file_ref_systemconfig_message_proto_rawDescData = file_ref_systemconfig_message_proto_rawDesc
)

func file_ref_systemconfig_message_proto_rawDescGZIP() []byte {
	file_ref_systemconfig_message_proto_rawDescOnce.Do(func() {
		file_ref_systemconfig_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_ref_systemconfig_message_proto_rawDescData)
	})
	return file_ref_systemconfig_message_proto_rawDescData
}

var file_ref_systemconfig_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ref_systemconfig_message_proto_goTypes = []interface{}{
	(*SystemConfig)(nil),          // 0: simplebank.SystemConfig
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*NullString)(nil),            // 2: simplebank.NullString
}
var file_ref_systemconfig_message_proto_depIdxs = []int32{
	1, // 0: simplebank.SystemConfig.glDate:type_name -> google.protobuf.Timestamp
	1, // 1: simplebank.SystemConfig.lastAccruals:type_name -> google.protobuf.Timestamp
	1, // 2: simplebank.SystemConfig.lastMonthEnd:type_name -> google.protobuf.Timestamp
	1, // 3: simplebank.SystemConfig.nextMonthEnd:type_name -> google.protobuf.Timestamp
	1, // 4: simplebank.SystemConfig.systemDate:type_name -> google.protobuf.Timestamp
	2, // 5: simplebank.SystemConfig.otherInfo:type_name -> simplebank.NullString
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ref_systemconfig_message_proto_init() }
func file_ref_systemconfig_message_proto_init() {
	if File_ref_systemconfig_message_proto != nil {
		return
	}
	file_type_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ref_systemconfig_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemConfig); i {
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
			RawDescriptor: file_ref_systemconfig_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ref_systemconfig_message_proto_goTypes,
		DependencyIndexes: file_ref_systemconfig_message_proto_depIdxs,
		MessageInfos:      file_ref_systemconfig_message_proto_msgTypes,
	}.Build()
	File_ref_systemconfig_message_proto = out.File
	file_ref_systemconfig_message_proto_rawDesc = nil
	file_ref_systemconfig_message_proto_goTypes = nil
	file_ref_systemconfig_message_proto_depIdxs = nil
}
