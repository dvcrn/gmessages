// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: util.proto

package gmproto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmptyArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyArr) Reset() {
	*x = EmptyArr{}
	mi := &file_util_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyArr) ProtoMessage() {}

func (x *EmptyArr) ProtoReflect() protoreflect.Message {
	mi := &file_util_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyArr.ProtoReflect.Descriptor instead.
func (*EmptyArr) Descriptor() ([]byte, []int) {
	return file_util_proto_rawDescGZIP(), []int{0}
}

var File_util_proto protoreflect.FileDescriptor

//go:embed util.pb.raw
var file_util_proto_rawDesc []byte

var (
	file_util_proto_rawDescOnce sync.Once
	file_util_proto_rawDescData = file_util_proto_rawDesc
)

func file_util_proto_rawDescGZIP() []byte {
	file_util_proto_rawDescOnce.Do(func() {
		file_util_proto_rawDescData = protoimpl.X.CompressGZIP(file_util_proto_rawDescData)
	})
	return file_util_proto_rawDescData
}

var file_util_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_util_proto_goTypes = []any{
	(*EmptyArr)(nil), // 0: util.EmptyArr
}
var file_util_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_util_proto_init() }
func file_util_proto_init() {
	if File_util_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_util_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_util_proto_goTypes,
		DependencyIndexes: file_util_proto_depIdxs,
		MessageInfos:      file_util_proto_msgTypes,
	}.Build()
	File_util_proto = out.File
	file_util_proto_rawDesc = nil
	file_util_proto_goTypes = nil
	file_util_proto_depIdxs = nil
}
