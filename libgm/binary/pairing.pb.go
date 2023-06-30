// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pairing.proto

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

type BrowserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent string `protobuf:"bytes,1,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	SomeInt   int32  `protobuf:"varint,2,opt,name=someInt,proto3" json:"someInt,omitempty"`
	Os        string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	SomeBool  bool   `protobuf:"varint,6,opt,name=someBool,proto3" json:"someBool,omitempty"`
}

func (x *BrowserDetails) Reset() {
	*x = BrowserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserDetails) ProtoMessage() {}

func (x *BrowserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserDetails.ProtoReflect.Descriptor instead.
func (*BrowserDetails) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{0}
}

func (x *BrowserDetails) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *BrowserDetails) GetSomeInt() int32 {
	if x != nil {
		return x.SomeInt
	}
	return 0
}

func (x *BrowserDetails) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *BrowserDetails) GetSomeBool() bool {
	if x != nil {
		return x.SomeBool
	}
	return false
}

type PhoneRelayBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Bugle  string `protobuf:"bytes,3,opt,name=bugle,proto3" json:"bugle,omitempty"`
	RpcKey []byte `protobuf:"bytes,6,opt,name=rpcKey,proto3" json:"rpcKey,omitempty"`
	Date   *Date  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *PhoneRelayBody) Reset() {
	*x = PhoneRelayBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneRelayBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneRelayBody) ProtoMessage() {}

func (x *PhoneRelayBody) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneRelayBody.ProtoReflect.Descriptor instead.
func (*PhoneRelayBody) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{1}
}

func (x *PhoneRelayBody) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PhoneRelayBody) GetBugle() string {
	if x != nil {
		return x.Bugle
	}
	return ""
}

func (x *PhoneRelayBody) GetRpcKey() []byte {
	if x != nil {
		return x.RpcKey
	}
	return nil
}

func (x *PhoneRelayBody) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

type ECDSAKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoVersion  int64  `protobuf:"varint,1,opt,name=protoVersion,proto3" json:"protoVersion,omitempty"` // idk?
	EncryptedKeys []byte `protobuf:"bytes,2,opt,name=encryptedKeys,proto3" json:"encryptedKeys,omitempty"`
}

func (x *ECDSAKeys) Reset() {
	*x = ECDSAKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECDSAKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECDSAKeys) ProtoMessage() {}

func (x *ECDSAKeys) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECDSAKeys.ProtoReflect.Descriptor instead.
func (*ECDSAKeys) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{2}
}

func (x *ECDSAKeys) GetProtoVersion() int64 {
	if x != nil {
		return x.ProtoVersion
	}
	return 0
}

func (x *ECDSAKeys) GetEncryptedKeys() []byte {
	if x != nil {
		return x.EncryptedKeys
	}
	return nil
}

type PairDeviceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile         *Device     `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	EcdsaKeys      *ECDSAKeys  `protobuf:"bytes,6,opt,name=ecdsaKeys,proto3" json:"ecdsaKeys,omitempty"`
	WebAuthKeyData *WebAuthKey `protobuf:"bytes,2,opt,name=webAuthKeyData,proto3" json:"webAuthKeyData,omitempty"`
	Browser        *Device     `protobuf:"bytes,3,opt,name=browser,proto3" json:"browser,omitempty"`
}

func (x *PairDeviceData) Reset() {
	*x = PairDeviceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairDeviceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairDeviceData) ProtoMessage() {}

func (x *PairDeviceData) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairDeviceData.ProtoReflect.Descriptor instead.
func (*PairDeviceData) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{3}
}

func (x *PairDeviceData) GetMobile() *Device {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *PairDeviceData) GetEcdsaKeys() *ECDSAKeys {
	if x != nil {
		return x.EcdsaKeys
	}
	return nil
}

func (x *PairDeviceData) GetWebAuthKeyData() *WebAuthKey {
	if x != nil {
		return x.WebAuthKeyData
	}
	return nil
}

func (x *PairDeviceData) GetBrowser() *Device {
	if x != nil {
		return x.Browser
	}
	return nil
}

type WebAuthKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebAuthKey []byte `protobuf:"bytes,1,opt,name=webAuthKey,proto3" json:"webAuthKey,omitempty"`
	ValidFor   int64  `protobuf:"varint,2,opt,name=validFor,proto3" json:"validFor,omitempty"`
}

func (x *WebAuthKey) Reset() {
	*x = WebAuthKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebAuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthKey) ProtoMessage() {}

func (x *WebAuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthKey.ProtoReflect.Descriptor instead.
func (*WebAuthKey) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{4}
}

func (x *WebAuthKey) GetWebAuthKey() []byte {
	if x != nil {
		return x.WebAuthKey
	}
	return nil
}

func (x *WebAuthKey) GetValidFor() int64 {
	if x != nil {
		return x.ValidFor
	}
	return 0
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneRelay     *PhoneRelayBody `protobuf:"bytes,1,opt,name=PhoneRelay,proto3" json:"PhoneRelay,omitempty"`
	BrowserDetails *BrowserDetails `protobuf:"bytes,3,opt,name=browserDetails,proto3" json:"browserDetails,omitempty"`
	PairDeviceData *PairDeviceData `protobuf:"bytes,4,opt,name=pairDeviceData,proto3" json:"pairDeviceData,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{5}
}

func (x *Container) GetPhoneRelay() *PhoneRelayBody {
	if x != nil {
		return x.PhoneRelay
	}
	return nil
}

func (x *Container) GetBrowserDetails() *BrowserDetails {
	if x != nil {
		return x.BrowserDetails
	}
	return nil
}

func (x *Container) GetPairDeviceData() *PairDeviceData {
	if x != nil {
		return x.PairDeviceData
	}
	return nil
}

type UrlData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PairingKey   []byte `protobuf:"bytes,1,opt,name=pairingKey,proto3" json:"pairingKey,omitempty"`
	AESCTR256Key []byte `protobuf:"bytes,2,opt,name=AESCTR256Key,proto3" json:"AESCTR256Key,omitempty"`
	SHA256Key    []byte `protobuf:"bytes,3,opt,name=SHA256Key,proto3" json:"SHA256Key,omitempty"`
}

func (x *UrlData) Reset() {
	*x = UrlData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pairing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlData) ProtoMessage() {}

func (x *UrlData) ProtoReflect() protoreflect.Message {
	mi := &file_pairing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlData.ProtoReflect.Descriptor instead.
func (*UrlData) Descriptor() ([]byte, []int) {
	return file_pairing_proto_rawDescGZIP(), []int{6}
}

func (x *UrlData) GetPairingKey() []byte {
	if x != nil {
		return x.PairingKey
	}
	return nil
}

func (x *UrlData) GetAESCTR256Key() []byte {
	if x != nil {
		return x.AESCTR256Key
	}
	return nil
}

func (x *UrlData) GetSHA256Key() []byte {
	if x != nil {
		return x.SHA256Key
	}
	return nil
}

var File_pairing_proto protoreflect.FileDescriptor

var file_pairing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x6d, 0x65,
	0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x49,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x72,
	0x0a, 0x0e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x67, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x75, 0x67, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x22,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x55, 0x0a, 0x09, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x0e, 0x50, 0x61,
	0x69, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x63, 0x64, 0x73, 0x61, 0x4b,
	0x65, 0x79, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x09, 0x65,
	0x63, 0x64, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x77, 0x65, 0x62, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x77, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x22, 0x48, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x22, 0xc6, 0x01, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0a, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x12, 0x3f, 0x0a, 0x0e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x70, 0x61, 0x69, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x70, 0x61, 0x69, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x07, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x45, 0x53, 0x43, 0x54, 0x52, 0x32, 0x35, 0x36, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x41, 0x45, 0x53, 0x43, 0x54, 0x52, 0x32, 0x35, 0x36,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x4b, 0x65,
	0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pairing_proto_rawDescOnce sync.Once
	file_pairing_proto_rawDescData = file_pairing_proto_rawDesc
)

func file_pairing_proto_rawDescGZIP() []byte {
	file_pairing_proto_rawDescOnce.Do(func() {
		file_pairing_proto_rawDescData = protoimpl.X.CompressGZIP(file_pairing_proto_rawDescData)
	})
	return file_pairing_proto_rawDescData
}

var file_pairing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pairing_proto_goTypes = []interface{}{
	(*BrowserDetails)(nil), // 0: pairing.BrowserDetails
	(*PhoneRelayBody)(nil), // 1: pairing.PhoneRelayBody
	(*ECDSAKeys)(nil),      // 2: pairing.ECDSAKeys
	(*PairDeviceData)(nil), // 3: pairing.PairDeviceData
	(*WebAuthKey)(nil),     // 4: pairing.WebAuthKey
	(*Container)(nil),      // 5: pairing.Container
	(*UrlData)(nil),        // 6: pairing.UrlData
	(*Date)(nil),           // 7: messages.Date
	(*Device)(nil),         // 8: messages.Device
}
var file_pairing_proto_depIdxs = []int32{
	7, // 0: pairing.PhoneRelayBody.date:type_name -> messages.Date
	8, // 1: pairing.PairDeviceData.mobile:type_name -> messages.Device
	2, // 2: pairing.PairDeviceData.ecdsaKeys:type_name -> pairing.ECDSAKeys
	4, // 3: pairing.PairDeviceData.webAuthKeyData:type_name -> pairing.WebAuthKey
	8, // 4: pairing.PairDeviceData.browser:type_name -> messages.Device
	1, // 5: pairing.Container.PhoneRelay:type_name -> pairing.PhoneRelayBody
	0, // 6: pairing.Container.browserDetails:type_name -> pairing.BrowserDetails
	3, // 7: pairing.Container.pairDeviceData:type_name -> pairing.PairDeviceData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pairing_proto_init() }
func file_pairing_proto_init() {
	if File_pairing_proto != nil {
		return
	}
	file_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pairing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserDetails); i {
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
		file_pairing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneRelayBody); i {
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
		file_pairing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECDSAKeys); i {
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
		file_pairing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairDeviceData); i {
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
		file_pairing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebAuthKey); i {
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
		file_pairing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_pairing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlData); i {
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
			RawDescriptor: file_pairing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pairing_proto_goTypes,
		DependencyIndexes: file_pairing_proto_depIdxs,
		MessageInfos:      file_pairing_proto_msgTypes,
	}.Build()
	File_pairing_proto = out.File
	file_pairing_proto_rawDesc = nil
	file_pairing_proto_goTypes = nil
	file_pairing_proto_depIdxs = nil
}
