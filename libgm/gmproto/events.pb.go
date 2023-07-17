// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: events.proto

package gmproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

import _ "embed"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlertType int32

const (
	AlertType_ALERT_TYPE_UNKNOWN                     AlertType = 0
	AlertType_BROWSER_INACTIVE                       AlertType = 1  // Emitted whenever browser connection becomes inactive
	AlertType_BROWSER_ACTIVE                         AlertType = 2  // Emitted whenever a new browser session is created
	AlertType_MOBILE_DATA_CONNECTION                 AlertType = 3  // Emitted when the paired device connects to data
	AlertType_MOBILE_WIFI_CONNECTION                 AlertType = 4  // Emitted when the paired device connects to wifi
	AlertType_MOBILE_BATTERY_LOW                     AlertType = 5  // Emitted if the paired device reaches low battery
	AlertType_MOBILE_BATTERY_RESTORED                AlertType = 6  // Emitted if the paired device has restored battery enough to not be considered low
	AlertType_BROWSER_INACTIVE_FROM_TIMEOUT          AlertType = 7  // Emitted whenever browser connection becomes inactive due to timeout
	AlertType_BROWSER_INACTIVE_FROM_INACTIVITY       AlertType = 8  // Emitted whenever browser connection becomes inactive due to inactivity
	AlertType_RCS_CONNECTION                         AlertType = 9  // Emitted whenever RCS connection has been established successfully
	AlertType_OBSERVER_REGISTERED                    AlertType = 10 // Unknown
	AlertType_MOBILE_DATABASE_SYNCING                AlertType = 11 // Emitted whenever the paired device is attempting to sync db
	AlertType_MOBILE_DATABASE_SYNC_COMPLETE          AlertType = 12 // Emitted whenever the paired device has completed the db sync
	AlertType_MOBILE_DATABASE_SYNC_STARTED           AlertType = 13 // Emitted whenever the paired device has begun syncing the db
	AlertType_MOBILE_DATABASE_PARTIAL_SYNC_COMPLETED AlertType = 14 // Emitted whenever the paired device has successfully synced a chunk of the db
	AlertType_MOBILE_DATABASE_PARTIAL_SYNC_STARTED   AlertType = 15 // Emitted whenever the paired device has begun syncing a chunk of the db
	AlertType_CONTACTS_REFRESH_STARTED               AlertType = 16 // Emitted whenever the paired device has begun refreshing contacts
	AlertType_CONTACTS_REFRESH_COMPLETED             AlertType = 17 // Emitted whenever the paired device has successfully refreshed contacts
)

// Enum value maps for AlertType.
var (
	AlertType_name = map[int32]string{
		0:  "ALERT_TYPE_UNKNOWN",
		1:  "BROWSER_INACTIVE",
		2:  "BROWSER_ACTIVE",
		3:  "MOBILE_DATA_CONNECTION",
		4:  "MOBILE_WIFI_CONNECTION",
		5:  "MOBILE_BATTERY_LOW",
		6:  "MOBILE_BATTERY_RESTORED",
		7:  "BROWSER_INACTIVE_FROM_TIMEOUT",
		8:  "BROWSER_INACTIVE_FROM_INACTIVITY",
		9:  "RCS_CONNECTION",
		10: "OBSERVER_REGISTERED",
		11: "MOBILE_DATABASE_SYNCING",
		12: "MOBILE_DATABASE_SYNC_COMPLETE",
		13: "MOBILE_DATABASE_SYNC_STARTED",
		14: "MOBILE_DATABASE_PARTIAL_SYNC_COMPLETED",
		15: "MOBILE_DATABASE_PARTIAL_SYNC_STARTED",
		16: "CONTACTS_REFRESH_STARTED",
		17: "CONTACTS_REFRESH_COMPLETED",
	}
	AlertType_value = map[string]int32{
		"ALERT_TYPE_UNKNOWN":                     0,
		"BROWSER_INACTIVE":                       1,
		"BROWSER_ACTIVE":                         2,
		"MOBILE_DATA_CONNECTION":                 3,
		"MOBILE_WIFI_CONNECTION":                 4,
		"MOBILE_BATTERY_LOW":                     5,
		"MOBILE_BATTERY_RESTORED":                6,
		"BROWSER_INACTIVE_FROM_TIMEOUT":          7,
		"BROWSER_INACTIVE_FROM_INACTIVITY":       8,
		"RCS_CONNECTION":                         9,
		"OBSERVER_REGISTERED":                    10,
		"MOBILE_DATABASE_SYNCING":                11,
		"MOBILE_DATABASE_SYNC_COMPLETE":          12,
		"MOBILE_DATABASE_SYNC_STARTED":           13,
		"MOBILE_DATABASE_PARTIAL_SYNC_COMPLETED": 14,
		"MOBILE_DATABASE_PARTIAL_SYNC_STARTED":   15,
		"CONTACTS_REFRESH_STARTED":               16,
		"CONTACTS_REFRESH_COMPLETED":             17,
	}
)

func (x AlertType) Enum() *AlertType {
	p := new(AlertType)
	*p = x
	return p
}

func (x AlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[0].Descriptor()
}

func (AlertType) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[0]
}

func (x AlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertType.Descriptor instead.
func (AlertType) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

type TypingTypes int32

const (
	TypingTypes_STOPPED_TYPING TypingTypes = 0
	TypingTypes_STARTED_TYPING TypingTypes = 1
)

// Enum value maps for TypingTypes.
var (
	TypingTypes_name = map[int32]string{
		0: "STOPPED_TYPING",
		1: "STARTED_TYPING",
	}
	TypingTypes_value = map[string]int32{
		"STOPPED_TYPING": 0,
		"STARTED_TYPING": 1,
	}
)

func (x TypingTypes) Enum() *TypingTypes {
	p := new(TypingTypes)
	*p = x
	return p
}

func (x TypingTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypingTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[1].Descriptor()
}

func (TypingTypes) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[1]
}

func (x TypingTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypingTypes.Descriptor instead.
func (TypingTypes) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

type UpdateEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*UpdateEvents_ConversationEvent
	//	*UpdateEvents_MessageEvent
	//	*UpdateEvents_TypingEvent
	//	*UpdateEvents_SettingsEvent
	//	*UpdateEvents_UserAlertEvent
	Event isUpdateEvents_Event `protobuf_oneof:"event"`
}

func (x *UpdateEvents) Reset() {
	*x = UpdateEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEvents) ProtoMessage() {}

func (x *UpdateEvents) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEvents.ProtoReflect.Descriptor instead.
func (*UpdateEvents) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (m *UpdateEvents) GetEvent() isUpdateEvents_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *UpdateEvents) GetConversationEvent() *ConversationEvent {
	if x, ok := x.GetEvent().(*UpdateEvents_ConversationEvent); ok {
		return x.ConversationEvent
	}
	return nil
}

func (x *UpdateEvents) GetMessageEvent() *MessageEvent {
	if x, ok := x.GetEvent().(*UpdateEvents_MessageEvent); ok {
		return x.MessageEvent
	}
	return nil
}

func (x *UpdateEvents) GetTypingEvent() *TypingEvent {
	if x, ok := x.GetEvent().(*UpdateEvents_TypingEvent); ok {
		return x.TypingEvent
	}
	return nil
}

func (x *UpdateEvents) GetSettingsEvent() *Settings {
	if x, ok := x.GetEvent().(*UpdateEvents_SettingsEvent); ok {
		return x.SettingsEvent
	}
	return nil
}

func (x *UpdateEvents) GetUserAlertEvent() *UserAlertEvent {
	if x, ok := x.GetEvent().(*UpdateEvents_UserAlertEvent); ok {
		return x.UserAlertEvent
	}
	return nil
}

type isUpdateEvents_Event interface {
	isUpdateEvents_Event()
}

type UpdateEvents_ConversationEvent struct {
	ConversationEvent *ConversationEvent `protobuf:"bytes,2,opt,name=conversationEvent,proto3,oneof"`
}

type UpdateEvents_MessageEvent struct {
	MessageEvent *MessageEvent `protobuf:"bytes,3,opt,name=messageEvent,proto3,oneof"`
}

type UpdateEvents_TypingEvent struct {
	TypingEvent *TypingEvent `protobuf:"bytes,4,opt,name=typingEvent,proto3,oneof"`
}

type UpdateEvents_SettingsEvent struct {
	SettingsEvent *Settings `protobuf:"bytes,5,opt,name=settingsEvent,proto3,oneof"`
}

type UpdateEvents_UserAlertEvent struct {
	UserAlertEvent *UserAlertEvent `protobuf:"bytes,6,opt,name=userAlertEvent,proto3,oneof"`
}

func (*UpdateEvents_ConversationEvent) isUpdateEvents_Event() {}

func (*UpdateEvents_MessageEvent) isUpdateEvents_Event() {}

func (*UpdateEvents_TypingEvent) isUpdateEvents_Event() {}

func (*UpdateEvents_SettingsEvent) isUpdateEvents_Event() {}

func (*UpdateEvents_UserAlertEvent) isUpdateEvents_Event() {}

type ConversationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Conversation `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConversationEvent) Reset() {
	*x = ConversationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationEvent) ProtoMessage() {}

func (x *ConversationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationEvent.ProtoReflect.Descriptor instead.
func (*ConversationEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *ConversationEvent) GetData() *Conversation {
	if x != nil {
		return x.Data
	}
	return nil
}

type TypingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *TypingData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TypingEvent) Reset() {
	*x = TypingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypingEvent) ProtoMessage() {}

func (x *TypingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypingEvent.ProtoReflect.Descriptor instead.
func (*TypingEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *TypingEvent) GetData() *TypingData {
	if x != nil {
		return x.Data
	}
	return nil
}

type MessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Message `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageEvent) Reset() {
	*x = MessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageEvent) ProtoMessage() {}

func (x *MessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageEvent.ProtoReflect.Descriptor instead.
func (*MessageEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *MessageEvent) GetData() *Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserAlertEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertType AlertType `protobuf:"varint,2,opt,name=alertType,proto3,enum=events.AlertType" json:"alertType,omitempty"`
}

func (x *UserAlertEvent) Reset() {
	*x = UserAlertEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAlertEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAlertEvent) ProtoMessage() {}

func (x *UserAlertEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAlertEvent.ProtoReflect.Descriptor instead.
func (*UserAlertEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *UserAlertEvent) GetAlertType() AlertType {
	if x != nil {
		return x.AlertType
	}
	return AlertType_ALERT_TYPE_UNKNOWN
}

type TypingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationID string      `protobuf:"bytes,1,opt,name=conversationID,proto3" json:"conversationID,omitempty"`
	User           *User       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Type           TypingTypes `protobuf:"varint,3,opt,name=type,proto3,enum=events.TypingTypes" json:"type,omitempty"`
}

func (x *TypingData) Reset() {
	*x = TypingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypingData) ProtoMessage() {}

func (x *TypingData) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypingData.ProtoReflect.Descriptor instead.
func (*TypingData) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{5}
}

func (x *TypingData) GetConversationID() string {
	if x != nil {
		return x.ConversationID
	}
	return ""
}

func (x *TypingData) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TypingData) GetType() TypingTypes {
	if x != nil {
		return x.Type
	}
	return TypingTypes_STOPPED_TYPING
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1 int64  `protobuf:"varint,1,opt,name=field1,proto3" json:"field1,omitempty"`
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetField1() int64 {
	if x != nil {
		return x.Field1
	}
	return 0
}

func (x *User) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type RPCPairData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*RPCPairData_Paired
	//	*RPCPairData_Revoked
	Event isRPCPairData_Event `protobuf_oneof:"event"`
}

func (x *RPCPairData) Reset() {
	*x = RPCPairData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCPairData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCPairData) ProtoMessage() {}

func (x *RPCPairData) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCPairData.ProtoReflect.Descriptor instead.
func (*RPCPairData) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{7}
}

func (m *RPCPairData) GetEvent() isRPCPairData_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *RPCPairData) GetPaired() *PairedData {
	if x, ok := x.GetEvent().(*RPCPairData_Paired); ok {
		return x.Paired
	}
	return nil
}

func (x *RPCPairData) GetRevoked() *RevokePairData {
	if x, ok := x.GetEvent().(*RPCPairData_Revoked); ok {
		return x.Revoked
	}
	return nil
}

type isRPCPairData_Event interface {
	isRPCPairData_Event()
}

type RPCPairData_Paired struct {
	Paired *PairedData `protobuf:"bytes,4,opt,name=paired,proto3,oneof"`
}

type RPCPairData_Revoked struct {
	Revoked *RevokePairData `protobuf:"bytes,5,opt,name=revoked,proto3,oneof"`
}

func (*RPCPairData_Paired) isRPCPairData_Event() {}

func (*RPCPairData_Revoked) isRPCPairData_Event() {}

var File_events_proto protoreflect.FileDescriptor

//go:embed events.pb.raw
var file_events_proto_rawDesc []byte

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_events_proto_goTypes = []interface{}{
	(AlertType)(0),            // 0: events.AlertType
	(TypingTypes)(0),          // 1: events.TypingTypes
	(*UpdateEvents)(nil),      // 2: events.UpdateEvents
	(*ConversationEvent)(nil), // 3: events.ConversationEvent
	(*TypingEvent)(nil),       // 4: events.TypingEvent
	(*MessageEvent)(nil),      // 5: events.MessageEvent
	(*UserAlertEvent)(nil),    // 6: events.UserAlertEvent
	(*TypingData)(nil),        // 7: events.TypingData
	(*User)(nil),              // 8: events.User
	(*RPCPairData)(nil),       // 9: events.RPCPairData
	(*Settings)(nil),          // 10: settings.Settings
	(*Conversation)(nil),      // 11: conversations.Conversation
	(*Message)(nil),           // 12: conversations.Message
	(*PairedData)(nil),        // 13: authentication.PairedData
	(*RevokePairData)(nil),    // 14: authentication.RevokePairData
}
var file_events_proto_depIdxs = []int32{
	3,  // 0: events.UpdateEvents.conversationEvent:type_name -> events.ConversationEvent
	5,  // 1: events.UpdateEvents.messageEvent:type_name -> events.MessageEvent
	4,  // 2: events.UpdateEvents.typingEvent:type_name -> events.TypingEvent
	10, // 3: events.UpdateEvents.settingsEvent:type_name -> settings.Settings
	6,  // 4: events.UpdateEvents.userAlertEvent:type_name -> events.UserAlertEvent
	11, // 5: events.ConversationEvent.data:type_name -> conversations.Conversation
	7,  // 6: events.TypingEvent.data:type_name -> events.TypingData
	12, // 7: events.MessageEvent.data:type_name -> conversations.Message
	0,  // 8: events.UserAlertEvent.alertType:type_name -> events.AlertType
	8,  // 9: events.TypingData.user:type_name -> events.User
	1,  // 10: events.TypingData.type:type_name -> events.TypingTypes
	13, // 11: events.RPCPairData.paired:type_name -> authentication.PairedData
	14, // 12: events.RPCPairData.revoked:type_name -> authentication.RevokePairData
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	file_conversations_proto_init()
	file_authentication_proto_init()
	file_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEvents); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationEvent); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypingEvent); i {
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
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageEvent); i {
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
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAlertEvent); i {
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
		file_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypingData); i {
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
		file_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCPairData); i {
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
	file_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UpdateEvents_ConversationEvent)(nil),
		(*UpdateEvents_MessageEvent)(nil),
		(*UpdateEvents_TypingEvent)(nil),
		(*UpdateEvents_SettingsEvent)(nil),
		(*UpdateEvents_UserAlertEvent)(nil),
	}
	file_events_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*RPCPairData_Paired)(nil),
		(*RPCPairData_Revoked)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		EnumInfos:         file_events_proto_enumTypes,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}
