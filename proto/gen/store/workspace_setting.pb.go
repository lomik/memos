// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/workspace_setting.proto

package store

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

type WorkspaceSettingKey int32

const (
	WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED WorkspaceSettingKey = 0
	// WORKSPACE_SETTING_GENERAL is the key for general settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_GENERAL WorkspaceSettingKey = 1
	// WORKSPACE_SETTING_STORAGE is the key for storage settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_STORAGE WorkspaceSettingKey = 2
	// WORKSPACE_SETTING_MEMO_RELATED is the key for memo related settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_MEMO_RELATED WorkspaceSettingKey = 3
	// WORKSPACE_SETTING_TELEGRAM_INTEGRATION is the key for telegram integration settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_TELEGRAM_INTEGRATION WorkspaceSettingKey = 4
)

// Enum value maps for WorkspaceSettingKey.
var (
	WorkspaceSettingKey_name = map[int32]string{
		0: "WORKSPACE_SETTING_KEY_UNSPECIFIED",
		1: "WORKSPACE_SETTING_GENERAL",
		2: "WORKSPACE_SETTING_STORAGE",
		3: "WORKSPACE_SETTING_MEMO_RELATED",
		4: "WORKSPACE_SETTING_TELEGRAM_INTEGRATION",
	}
	WorkspaceSettingKey_value = map[string]int32{
		"WORKSPACE_SETTING_KEY_UNSPECIFIED":      0,
		"WORKSPACE_SETTING_GENERAL":              1,
		"WORKSPACE_SETTING_STORAGE":              2,
		"WORKSPACE_SETTING_MEMO_RELATED":         3,
		"WORKSPACE_SETTING_TELEGRAM_INTEGRATION": 4,
	}
)

func (x WorkspaceSettingKey) Enum() *WorkspaceSettingKey {
	p := new(WorkspaceSettingKey)
	*p = x
	return p
}

func (x WorkspaceSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[0].Descriptor()
}

func (WorkspaceSettingKey) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[0]
}

func (x WorkspaceSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceSettingKey.Descriptor instead.
func (WorkspaceSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

type WorkspaceStorageSetting_StorageType int32

const (
	WorkspaceStorageSetting_STORAGE_TYPE_UNSPECIFIED WorkspaceStorageSetting_StorageType = 0
	// STORAGE_TYPE_DATABASE is the database storage type.
	WorkspaceStorageSetting_STORAGE_TYPE_DATABASE WorkspaceStorageSetting_StorageType = 1
	// STORAGE_TYPE_LOCAL is the local storage type.
	WorkspaceStorageSetting_STORAGE_TYPE_LOCAL WorkspaceStorageSetting_StorageType = 2
	// STORAGE_TYPE_EXTERNAL is the external storage type.
	WorkspaceStorageSetting_STORAGE_TYPE_EXTERNAL WorkspaceStorageSetting_StorageType = 3
)

// Enum value maps for WorkspaceStorageSetting_StorageType.
var (
	WorkspaceStorageSetting_StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_UNSPECIFIED",
		1: "STORAGE_TYPE_DATABASE",
		2: "STORAGE_TYPE_LOCAL",
		3: "STORAGE_TYPE_EXTERNAL",
	}
	WorkspaceStorageSetting_StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNSPECIFIED": 0,
		"STORAGE_TYPE_DATABASE":    1,
		"STORAGE_TYPE_LOCAL":       2,
		"STORAGE_TYPE_EXTERNAL":    3,
	}
)

func (x WorkspaceStorageSetting_StorageType) Enum() *WorkspaceStorageSetting_StorageType {
	p := new(WorkspaceStorageSetting_StorageType)
	*p = x
	return p
}

func (x WorkspaceStorageSetting_StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceStorageSetting_StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[1].Descriptor()
}

func (WorkspaceStorageSetting_StorageType) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[1]
}

func (x WorkspaceStorageSetting_StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceStorageSetting_StorageType.Descriptor instead.
func (WorkspaceStorageSetting_StorageType) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{2, 0}
}

type WorkspaceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key WorkspaceSettingKey `protobuf:"varint,1,opt,name=key,proto3,enum=memos.store.WorkspaceSettingKey" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*WorkspaceSetting_GeneralSetting
	//	*WorkspaceSetting_StorageSetting
	//	*WorkspaceSetting_MemoRelatedSetting
	//	*WorkspaceSetting_TelegramIntegrationSetting
	Value isWorkspaceSetting_Value `protobuf_oneof:"value"`
}

func (x *WorkspaceSetting) Reset() {
	*x = WorkspaceSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting) ProtoMessage() {}

func (x *WorkspaceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceSetting) GetKey() WorkspaceSettingKey {
	if x != nil {
		return x.Key
	}
	return WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED
}

func (m *WorkspaceSetting) GetValue() isWorkspaceSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *WorkspaceSetting) GetGeneralSetting() *WorkspaceGeneralSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_GeneralSetting); ok {
		return x.GeneralSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetStorageSetting() *WorkspaceStorageSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_StorageSetting); ok {
		return x.StorageSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetMemoRelatedSetting() *WorkspaceMemoRelatedSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_MemoRelatedSetting); ok {
		return x.MemoRelatedSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetTelegramIntegrationSetting() *WorkspaceTelegramIntegrationSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_TelegramIntegrationSetting); ok {
		return x.TelegramIntegrationSetting
	}
	return nil
}

type isWorkspaceSetting_Value interface {
	isWorkspaceSetting_Value()
}

type WorkspaceSetting_GeneralSetting struct {
	GeneralSetting *WorkspaceGeneralSetting `protobuf:"bytes,2,opt,name=general_setting,json=generalSetting,proto3,oneof"`
}

type WorkspaceSetting_StorageSetting struct {
	StorageSetting *WorkspaceStorageSetting `protobuf:"bytes,3,opt,name=storage_setting,json=storageSetting,proto3,oneof"`
}

type WorkspaceSetting_MemoRelatedSetting struct {
	MemoRelatedSetting *WorkspaceMemoRelatedSetting `protobuf:"bytes,4,opt,name=memo_related_setting,json=memoRelatedSetting,proto3,oneof"`
}

type WorkspaceSetting_TelegramIntegrationSetting struct {
	TelegramIntegrationSetting *WorkspaceTelegramIntegrationSetting `protobuf:"bytes,5,opt,name=telegram_integration_setting,json=telegramIntegrationSetting,proto3,oneof"`
}

func (*WorkspaceSetting_GeneralSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_StorageSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_MemoRelatedSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_TelegramIntegrationSetting) isWorkspaceSetting_Value() {}

type WorkspaceGeneralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// instance_url is the instance URL.
	InstanceUrl string `protobuf:"bytes,1,opt,name=instance_url,json=instanceUrl,proto3" json:"instance_url,omitempty"`
	// disallow_signup is the flag to disallow signup.
	DisallowSignup bool `protobuf:"varint,2,opt,name=disallow_signup,json=disallowSignup,proto3" json:"disallow_signup,omitempty"`
	// disallow_password_login is the flag to disallow password login.
	DisallowPasswordLogin bool `protobuf:"varint,3,opt,name=disallow_password_login,json=disallowPasswordLogin,proto3" json:"disallow_password_login,omitempty"`
	// additional_script is the additional script.
	AdditionalScript string `protobuf:"bytes,5,opt,name=additional_script,json=additionalScript,proto3" json:"additional_script,omitempty"`
	// additional_style is the additional style.
	AdditionalStyle string `protobuf:"bytes,6,opt,name=additional_style,json=additionalStyle,proto3" json:"additional_style,omitempty"`
}

func (x *WorkspaceGeneralSetting) Reset() {
	*x = WorkspaceGeneralSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceGeneralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceGeneralSetting) ProtoMessage() {}

func (x *WorkspaceGeneralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceGeneralSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceGeneralSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceGeneralSetting) GetInstanceUrl() string {
	if x != nil {
		return x.InstanceUrl
	}
	return ""
}

func (x *WorkspaceGeneralSetting) GetDisallowSignup() bool {
	if x != nil {
		return x.DisallowSignup
	}
	return false
}

func (x *WorkspaceGeneralSetting) GetDisallowPasswordLogin() bool {
	if x != nil {
		return x.DisallowPasswordLogin
	}
	return false
}

func (x *WorkspaceGeneralSetting) GetAdditionalScript() string {
	if x != nil {
		return x.AdditionalScript
	}
	return ""
}

func (x *WorkspaceGeneralSetting) GetAdditionalStyle() string {
	if x != nil {
		return x.AdditionalStyle
	}
	return ""
}

type WorkspaceStorageSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// storage_type is the storage type.
	StorageType WorkspaceStorageSetting_StorageType `protobuf:"varint,1,opt,name=storage_type,json=storageType,proto3,enum=memos.store.WorkspaceStorageSetting_StorageType" json:"storage_type,omitempty"`
	// The local storage path for STORAGE_TYPE_LOCAL.
	// e.g. assets/{timestamp}_{filename}
	LocalStoragePath string `protobuf:"bytes,2,opt,name=local_storage_path,json=localStoragePath,proto3" json:"local_storage_path,omitempty"`
	// The max upload size in megabytes.
	UploadSizeLimitMb int64 `protobuf:"varint,3,opt,name=upload_size_limit_mb,json=uploadSizeLimitMb,proto3" json:"upload_size_limit_mb,omitempty"`
}

func (x *WorkspaceStorageSetting) Reset() {
	*x = WorkspaceStorageSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceStorageSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceStorageSetting) ProtoMessage() {}

func (x *WorkspaceStorageSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceStorageSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceStorageSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{2}
}

func (x *WorkspaceStorageSetting) GetStorageType() WorkspaceStorageSetting_StorageType {
	if x != nil {
		return x.StorageType
	}
	return WorkspaceStorageSetting_STORAGE_TYPE_UNSPECIFIED
}

func (x *WorkspaceStorageSetting) GetLocalStoragePath() string {
	if x != nil {
		return x.LocalStoragePath
	}
	return ""
}

func (x *WorkspaceStorageSetting) GetUploadSizeLimitMb() int64 {
	if x != nil {
		return x.UploadSizeLimitMb
	}
	return 0
}

type WorkspaceMemoRelatedSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// disallow_public_share disallows set memo as public visible.
	DisallowPublicVisible bool `protobuf:"varint,1,opt,name=disallow_public_visible,json=disallowPublicVisible,proto3" json:"disallow_public_visible,omitempty"`
	// display_with_update_time orders and displays memo with update time.
	DisplayWithUpdateTime bool `protobuf:"varint,2,opt,name=display_with_update_time,json=displayWithUpdateTime,proto3" json:"display_with_update_time,omitempty"`
}

func (x *WorkspaceMemoRelatedSetting) Reset() {
	*x = WorkspaceMemoRelatedSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceMemoRelatedSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceMemoRelatedSetting) ProtoMessage() {}

func (x *WorkspaceMemoRelatedSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceMemoRelatedSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceMemoRelatedSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{3}
}

func (x *WorkspaceMemoRelatedSetting) GetDisallowPublicVisible() bool {
	if x != nil {
		return x.DisallowPublicVisible
	}
	return false
}

func (x *WorkspaceMemoRelatedSetting) GetDisplayWithUpdateTime() bool {
	if x != nil {
		return x.DisplayWithUpdateTime
	}
	return false
}

type WorkspaceTelegramIntegrationSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bot_token is the telegram bot token.
	BotToken string `protobuf:"bytes,1,opt,name=bot_token,json=botToken,proto3" json:"bot_token,omitempty"`
}

func (x *WorkspaceTelegramIntegrationSetting) Reset() {
	*x = WorkspaceTelegramIntegrationSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceTelegramIntegrationSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceTelegramIntegrationSetting) ProtoMessage() {}

func (x *WorkspaceTelegramIntegrationSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceTelegramIntegrationSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceTelegramIntegrationSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{4}
}

func (x *WorkspaceTelegramIntegrationSetting) GetBotToken() string {
	if x != nil {
		return x.BotToken
	}
	return ""
}

var File_store_workspace_setting_proto protoreflect.FileDescriptor

var file_store_workspace_setting_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xc5, 0x03, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4f, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4f, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x12, 0x6d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x74, 0x0a, 0x1c, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x1a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x69,
	0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x36, 0x0a, 0x17,
	0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64,
	0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x22, 0xc8, 0x02, 0x0a,
	0x17, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x14, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x6d, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x62, 0x22, 0x79, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x4f,
	0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41,
	0x53, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x22, 0x8e, 0x01, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x37, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x23, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0xca, 0x01, 0x0a,
	0x13, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x4f,
	0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x57, 0x4f, 0x52,
	0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4d,
	0x45, 0x4d, 0x4f, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x2a, 0x0a,
	0x26, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x54, 0x45, 0x4c, 0x45, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42, 0xa0, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x15, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_workspace_setting_proto_rawDescOnce sync.Once
	file_store_workspace_setting_proto_rawDescData = file_store_workspace_setting_proto_rawDesc
)

func file_store_workspace_setting_proto_rawDescGZIP() []byte {
	file_store_workspace_setting_proto_rawDescOnce.Do(func() {
		file_store_workspace_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_workspace_setting_proto_rawDescData)
	})
	return file_store_workspace_setting_proto_rawDescData
}

var file_store_workspace_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_workspace_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_workspace_setting_proto_goTypes = []interface{}{
	(WorkspaceSettingKey)(0),                    // 0: memos.store.WorkspaceSettingKey
	(WorkspaceStorageSetting_StorageType)(0),    // 1: memos.store.WorkspaceStorageSetting.StorageType
	(*WorkspaceSetting)(nil),                    // 2: memos.store.WorkspaceSetting
	(*WorkspaceGeneralSetting)(nil),             // 3: memos.store.WorkspaceGeneralSetting
	(*WorkspaceStorageSetting)(nil),             // 4: memos.store.WorkspaceStorageSetting
	(*WorkspaceMemoRelatedSetting)(nil),         // 5: memos.store.WorkspaceMemoRelatedSetting
	(*WorkspaceTelegramIntegrationSetting)(nil), // 6: memos.store.WorkspaceTelegramIntegrationSetting
}
var file_store_workspace_setting_proto_depIdxs = []int32{
	0, // 0: memos.store.WorkspaceSetting.key:type_name -> memos.store.WorkspaceSettingKey
	3, // 1: memos.store.WorkspaceSetting.general_setting:type_name -> memos.store.WorkspaceGeneralSetting
	4, // 2: memos.store.WorkspaceSetting.storage_setting:type_name -> memos.store.WorkspaceStorageSetting
	5, // 3: memos.store.WorkspaceSetting.memo_related_setting:type_name -> memos.store.WorkspaceMemoRelatedSetting
	6, // 4: memos.store.WorkspaceSetting.telegram_integration_setting:type_name -> memos.store.WorkspaceTelegramIntegrationSetting
	1, // 5: memos.store.WorkspaceStorageSetting.storage_type:type_name -> memos.store.WorkspaceStorageSetting.StorageType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_store_workspace_setting_proto_init() }
func file_store_workspace_setting_proto_init() {
	if File_store_workspace_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_workspace_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceGeneralSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceStorageSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceMemoRelatedSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceTelegramIntegrationSetting); i {
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
	file_store_workspace_setting_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WorkspaceSetting_GeneralSetting)(nil),
		(*WorkspaceSetting_StorageSetting)(nil),
		(*WorkspaceSetting_MemoRelatedSetting)(nil),
		(*WorkspaceSetting_TelegramIntegrationSetting)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_workspace_setting_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_workspace_setting_proto_goTypes,
		DependencyIndexes: file_store_workspace_setting_proto_depIdxs,
		EnumInfos:         file_store_workspace_setting_proto_enumTypes,
		MessageInfos:      file_store_workspace_setting_proto_msgTypes,
	}.Build()
	File_store_workspace_setting_proto = out.File
	file_store_workspace_setting_proto_rawDesc = nil
	file_store_workspace_setting_proto_goTypes = nil
	file_store_workspace_setting_proto_depIdxs = nil
}
