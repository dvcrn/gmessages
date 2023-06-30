// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: media.proto

package binary

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

type StartMediaUploadPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageType int64        `protobuf:"varint,1,opt,name=imageType,proto3" json:"imageType,omitempty"`
	AuthData  *AuthMessage `protobuf:"bytes,2,opt,name=authData,proto3" json:"authData,omitempty"`
	Mobile    *Device      `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *StartMediaUploadPayload) Reset() {
	*x = StartMediaUploadPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMediaUploadPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMediaUploadPayload) ProtoMessage() {}

func (x *StartMediaUploadPayload) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMediaUploadPayload.ProtoReflect.Descriptor instead.
func (*StartMediaUploadPayload) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *StartMediaUploadPayload) GetImageType() int64 {
	if x != nil {
		return x.ImageType
	}
	return 0
}

func (x *StartMediaUploadPayload) GetAuthData() *AuthMessage {
	if x != nil {
		return x.AuthData
	}
	return nil
}

func (x *StartMediaUploadPayload) GetMobile() *Device {
	if x != nil {
		return x.Mobile
	}
	return nil
}

type UploadMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media   *Media `protobuf:"bytes,1,opt,name=media,proto3" json:"media,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UploadMediaResponse) Reset() {
	*x = UploadMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMediaResponse) ProtoMessage() {}

func (x *UploadMediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMediaResponse.ProtoReflect.Descriptor instead.
func (*UploadMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *UploadMediaResponse) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *UploadMediaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaID     string `protobuf:"bytes,1,opt,name=mediaID,proto3" json:"mediaID,omitempty"`
	MediaNumber int64  `protobuf:"varint,2,opt,name=mediaNumber,proto3" json:"mediaNumber,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *Media) GetMediaID() string {
	if x != nil {
		return x.MediaID
	}
	return ""
}

func (x *Media) GetMediaNumber() int64 {
	if x != nil {
		return x.MediaNumber
	}
	return 0
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x43, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData = file_media_proto_rawDesc
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_proto_rawDescData)
	})
	return file_media_proto_rawDescData
}

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_media_proto_goTypes = []interface{}{
	(*StartMediaUploadPayload)(nil), // 0: media.StartMediaUploadPayload
	(*UploadMediaResponse)(nil),     // 1: media.UploadMediaResponse
	(*Media)(nil),                   // 2: media.Media
	(*AuthMessage)(nil),             // 3: messages.AuthMessage
	(*Device)(nil),                  // 4: messages.Device
}
var file_media_proto_depIdxs = []int32{
	3, // 0: media.StartMediaUploadPayload.authData:type_name -> messages.AuthMessage
	4, // 1: media.StartMediaUploadPayload.mobile:type_name -> messages.Device
	2, // 2: media.UploadMediaResponse.media:type_name -> media.Media
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	file_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMediaUploadPayload); i {
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
		file_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadMediaResponse); i {
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
		file_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
			RawDescriptor: file_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_rawDesc = nil
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}
