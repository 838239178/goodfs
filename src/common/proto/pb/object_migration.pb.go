// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: object_migration.proto

package pb

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

type ObjectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ObjectData) Reset() {
	*x = ObjectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_migration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectData) ProtoMessage() {}

func (x *ObjectData) ProtoReflect() protoreflect.Message {
	mi := &file_object_migration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectData.ProtoReflect.Descriptor instead.
func (*ObjectData) Descriptor() ([]byte, []int) {
	return file_object_migration_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectData) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ObjectData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ObjectData) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName     string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	OriginLocate string `protobuf:"bytes,2,opt,name=originLocate,proto3" json:"originLocate,omitempty"`
	Size         int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ObjectInfo) Reset() {
	*x = ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_migration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectInfo) ProtoMessage() {}

func (x *ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_object_migration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectInfo.ProtoReflect.Descriptor instead.
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return file_object_migration_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ObjectInfo) GetOriginLocate() string {
	if x != nil {
		return x.OriginLocate
	}
	return ""
}

func (x *ObjectInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RequiredInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AverageCap    int64  `protobuf:"varint,1,opt,name=averageCap,proto3" json:"averageCap,omitempty"`
	RequiredSize  int64  `protobuf:"varint,2,opt,name=requiredSize,proto3" json:"requiredSize,omitempty"`
	TargetAddress string `protobuf:"bytes,3,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
}

func (x *RequiredInfo) Reset() {
	*x = RequiredInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_migration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredInfo) ProtoMessage() {}

func (x *RequiredInfo) ProtoReflect() protoreflect.Message {
	mi := &file_object_migration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredInfo.ProtoReflect.Descriptor instead.
func (*RequiredInfo) Descriptor() ([]byte, []int) {
	return file_object_migration_proto_rawDescGZIP(), []int{2}
}

func (x *RequiredInfo) GetAverageCap() int64 {
	if x != nil {
		return x.AverageCap
	}
	return 0
}

func (x *RequiredInfo) GetRequiredSize() int64 {
	if x != nil {
		return x.RequiredSize
	}
	return 0
}

func (x *RequiredInfo) GetTargetAddress() string {
	if x != nil {
		return x.TargetAddress
	}
	return ""
}

var File_object_migration_proto protoreflect.FileDescriptor

var file_object_migration_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50,
	0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x60, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x78, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x61, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x93, 0x02, 0x0a,
	0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_object_migration_proto_rawDescOnce sync.Once
	file_object_migration_proto_rawDescData = file_object_migration_proto_rawDesc
)

func file_object_migration_proto_rawDescGZIP() []byte {
	file_object_migration_proto_rawDescOnce.Do(func() {
		file_object_migration_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_migration_proto_rawDescData)
	})
	return file_object_migration_proto_rawDescData
}

var file_object_migration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_object_migration_proto_goTypes = []interface{}{
	(*ObjectData)(nil),   // 0: proto.ObjectData
	(*ObjectInfo)(nil),   // 1: proto.ObjectInfo
	(*RequiredInfo)(nil), // 2: proto.RequiredInfo
	(*EmptyReq)(nil),     // 3: proto.EmptyReq
	(*Response)(nil),     // 4: proto.Response
}
var file_object_migration_proto_depIdxs = []int32{
	0, // 0: proto.ObjectMigration.ReceiveData:input_type -> proto.ObjectData
	1, // 1: proto.ObjectMigration.FinishReceive:input_type -> proto.ObjectInfo
	2, // 2: proto.ObjectMigration.RequireSend:input_type -> proto.RequiredInfo
	3, // 3: proto.ObjectMigration.LeaveCommand:input_type -> proto.EmptyReq
	3, // 4: proto.ObjectMigration.JoinCommand:input_type -> proto.EmptyReq
	4, // 5: proto.ObjectMigration.ReceiveData:output_type -> proto.Response
	4, // 6: proto.ObjectMigration.FinishReceive:output_type -> proto.Response
	4, // 7: proto.ObjectMigration.RequireSend:output_type -> proto.Response
	4, // 8: proto.ObjectMigration.LeaveCommand:output_type -> proto.Response
	4, // 9: proto.ObjectMigration.JoinCommand:output_type -> proto.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_object_migration_proto_init() }
func file_object_migration_proto_init() {
	if File_object_migration_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_object_migration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectData); i {
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
		file_object_migration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectInfo); i {
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
		file_object_migration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredInfo); i {
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
			RawDescriptor: file_object_migration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_object_migration_proto_goTypes,
		DependencyIndexes: file_object_migration_proto_depIdxs,
		MessageInfos:      file_object_migration_proto_msgTypes,
	}.Build()
	File_object_migration_proto = out.File
	file_object_migration_proto_rawDesc = nil
	file_object_migration_proto_goTypes = nil
	file_object_migration_proto_depIdxs = nil
}
