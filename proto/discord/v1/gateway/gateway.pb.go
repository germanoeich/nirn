// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: discord/v1/gateway/gateway.proto

package gateway

import (
	model "github.com/germanoeich/nirn/proto/discord/v1/model"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type UpdateVoiceStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId   uint64 `protobuf:"fixed64,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId uint64 `protobuf:"fixed64,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	SelfMute  bool   `protobuf:"varint,3,opt,name=self_mute,json=selfMute,proto3" json:"self_mute,omitempty"`
	SelfDeaf  bool   `protobuf:"varint,4,opt,name=self_deaf,json=selfDeaf,proto3" json:"self_deaf,omitempty"`
}

func (x *UpdateVoiceStateRequest) Reset() {
	*x = UpdateVoiceStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoiceStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoiceStateRequest) ProtoMessage() {}

func (x *UpdateVoiceStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoiceStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateVoiceStateRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateVoiceStateRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *UpdateVoiceStateRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateVoiceStateRequest) GetSelfMute() bool {
	if x != nil {
		return x.SelfMute
	}
	return false
}

func (x *UpdateVoiceStateRequest) GetSelfDeaf() bool {
	if x != nil {
		return x.SelfDeaf
	}
	return false
}

type UpdateVoiceStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateVoiceStateResponse) Reset() {
	*x = UpdateVoiceStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoiceStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoiceStateResponse) ProtoMessage() {}

func (x *UpdateVoiceStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoiceStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateVoiceStateResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{1}
}

type UpdateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId    *wrappers.UInt32Value               `protobuf:"bytes,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Since      *timestamp.Timestamp                `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Activities []*UpdateStatusRequest_ActivityData `protobuf:"bytes,3,rep,name=activities,proto3" json:"activities,omitempty"`
	Status     model.PresenceData_OnlineStatus     `protobuf:"varint,4,opt,name=status,proto3,enum=pylon.discord.v1.model.PresenceData_OnlineStatus" json:"status,omitempty"`
	Afk        bool                                `protobuf:"varint,5,opt,name=afk,proto3" json:"afk,omitempty"`
}

func (x *UpdateStatusRequest) Reset() {
	*x = UpdateStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest) ProtoMessage() {}

func (x *UpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStatusRequest) GetShardId() *wrappers.UInt32Value {
	if x != nil {
		return x.ShardId
	}
	return nil
}

func (x *UpdateStatusRequest) GetSince() *timestamp.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *UpdateStatusRequest) GetActivities() []*UpdateStatusRequest_ActivityData {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *UpdateStatusRequest) GetStatus() model.PresenceData_OnlineStatus {
	if x != nil {
		return x.Status
	}
	return model.PresenceData_UNKNOWN
}

func (x *UpdateStatusRequest) GetAfk() bool {
	if x != nil {
		return x.Afk
	}
	return false
}

type UpdateStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStatusResponse) Reset() {
	*x = UpdateStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusResponse) ProtoMessage() {}

func (x *UpdateStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateStatusResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{3}
}

type FindUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"fixed64,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FindUserRequest) Reset() {
	*x = FindUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserRequest) ProtoMessage() {}

func (x *FindUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserRequest.ProtoReflect.Descriptor instead.
func (*FindUserRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *FindUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *model.UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *FindUserResponse) GetUser() *model.UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserMutualGuildsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"fixed64,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserMutualGuildsRequest) Reset() {
	*x = GetUserMutualGuildsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMutualGuildsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMutualGuildsRequest) ProtoMessage() {}

func (x *GetUserMutualGuildsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMutualGuildsRequest.ProtoReflect.Descriptor instead.
func (*GetUserMutualGuildsRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserMutualGuildsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserMutualGuildsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guilds []*model.GuildData `protobuf:"bytes,1,rep,name=guilds,proto3" json:"guilds,omitempty"`
}

func (x *GetUserMutualGuildsResponse) Reset() {
	*x = GetUserMutualGuildsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMutualGuildsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMutualGuildsResponse) ProtoMessage() {}

func (x *GetUserMutualGuildsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMutualGuildsResponse.ProtoReflect.Descriptor instead.
func (*GetUserMutualGuildsResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserMutualGuildsResponse) GetGuilds() []*model.GuildData {
	if x != nil {
		return x.Guilds
	}
	return nil
}

type FindEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmojiId uint64 `protobuf:"fixed64,1,opt,name=emoji_id,json=emojiId,proto3" json:"emoji_id,omitempty"`
}

func (x *FindEmojiRequest) Reset() {
	*x = FindEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmojiRequest) ProtoMessage() {}

func (x *FindEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmojiRequest.ProtoReflect.Descriptor instead.
func (*FindEmojiRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *FindEmojiRequest) GetEmojiId() uint64 {
	if x != nil {
		return x.EmojiId
	}
	return 0
}

type FindEmojiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emoji *model.EmojiData `protobuf:"bytes,1,opt,name=emoji,proto3" json:"emoji,omitempty"`
}

func (x *FindEmojiResponse) Reset() {
	*x = FindEmojiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmojiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmojiResponse) ProtoMessage() {}

func (x *FindEmojiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmojiResponse.ProtoReflect.Descriptor instead.
func (*FindEmojiResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *FindEmojiResponse) GetEmoji() *model.EmojiData {
	if x != nil {
		return x.Emoji
	}
	return nil
}

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{10}
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildCount        uint64 `protobuf:"varint,1,opt,name=guild_count,json=guildCount,proto3" json:"guild_count,omitempty"`
	UserCount         uint64 `protobuf:"varint,2,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	MemberCount       uint64 `protobuf:"varint,3,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	ConnectedChannels uint64 `protobuf:"varint,4,opt,name=connected_channels,json=connectedChannels,proto3" json:"connected_channels,omitempty"`
	ShardCount        uint32 `protobuf:"varint,5,opt,name=shard_count,json=shardCount,proto3" json:"shard_count,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{11}
}

func (x *GetStatsResponse) GetGuildCount() uint64 {
	if x != nil {
		return x.GuildCount
	}
	return 0
}

func (x *GetStatsResponse) GetUserCount() uint64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *GetStatsResponse) GetMemberCount() uint64 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *GetStatsResponse) GetConnectedChannels() uint64 {
	if x != nil {
		return x.ConnectedChannels
	}
	return 0
}

func (x *GetStatsResponse) GetShardCount() uint32 {
	if x != nil {
		return x.ShardCount
	}
	return 0
}

type UpdateStatusRequest_ActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type model.PresenceData_PresenceActivityData_ActivityType `protobuf:"varint,2,opt,name=type,proto3,enum=pylon.discord.v1.model.PresenceData_PresenceActivityData_ActivityType" json:"type,omitempty"`
}

func (x *UpdateStatusRequest_ActivityData) Reset() {
	*x = UpdateStatusRequest_ActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_v1_gateway_gateway_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusRequest_ActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest_ActivityData) ProtoMessage() {}

func (x *UpdateStatusRequest_ActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_discord_v1_gateway_gateway_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest_ActivityData.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest_ActivityData) Descriptor() ([]byte, []int) {
	return file_discord_v1_gateway_gateway_proto_rawDescGZIP(), []int{2, 0}
}

func (x *UpdateStatusRequest_ActivityData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStatusRequest_ActivityData) GetType() model.PresenceData_PresenceActivityData_ActivityType {
	if x != nil {
		return x.Type
	}
	return model.PresenceData_PresenceActivityData_UNKNOWN
}

var File_discord_v1_gateway_gateway_proto protoreflect.FileDescriptor

var file_discord_v1_gateway_gateway_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x66, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x4d, 0x75, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x64,
	0x65, 0x61, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x44,
	0x65, 0x61, 0x66, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xb9, 0x03, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x49,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x66, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x66, 0x6b, 0x1a, 0x7e, 0x0a, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x5a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e,
	0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x22, 0x31, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x6f,
	0x6a, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x79, 0x6c, 0x6f,
	0x6e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x67,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x0a, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0b, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x64,
	0x0a, 0x22, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x79, 0x6c, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x6f, 0x65, 0x69, 0x63, 0x68, 0x2f, 0x6e,
	0x69, 0x72, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discord_v1_gateway_gateway_proto_rawDescOnce sync.Once
	file_discord_v1_gateway_gateway_proto_rawDescData = file_discord_v1_gateway_gateway_proto_rawDesc
)

func file_discord_v1_gateway_gateway_proto_rawDescGZIP() []byte {
	file_discord_v1_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_discord_v1_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_discord_v1_gateway_gateway_proto_rawDescData)
	})
	return file_discord_v1_gateway_gateway_proto_rawDescData
}

var file_discord_v1_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_discord_v1_gateway_gateway_proto_goTypes = []interface{}{
	(*UpdateVoiceStateRequest)(nil),                           // 0: pylon.discord.v1.gateway.UpdateVoiceStateRequest
	(*UpdateVoiceStateResponse)(nil),                          // 1: pylon.discord.v1.gateway.UpdateVoiceStateResponse
	(*UpdateStatusRequest)(nil),                               // 2: pylon.discord.v1.gateway.UpdateStatusRequest
	(*UpdateStatusResponse)(nil),                              // 3: pylon.discord.v1.gateway.UpdateStatusResponse
	(*FindUserRequest)(nil),                                   // 4: pylon.discord.v1.gateway.FindUserRequest
	(*FindUserResponse)(nil),                                  // 5: pylon.discord.v1.gateway.FindUserResponse
	(*GetUserMutualGuildsRequest)(nil),                        // 6: pylon.discord.v1.gateway.GetUserMutualGuildsRequest
	(*GetUserMutualGuildsResponse)(nil),                       // 7: pylon.discord.v1.gateway.GetUserMutualGuildsResponse
	(*FindEmojiRequest)(nil),                                  // 8: pylon.discord.v1.gateway.FindEmojiRequest
	(*FindEmojiResponse)(nil),                                 // 9: pylon.discord.v1.gateway.FindEmojiResponse
	(*GetStatsRequest)(nil),                                   // 10: pylon.discord.v1.gateway.GetStatsRequest
	(*GetStatsResponse)(nil),                                  // 11: pylon.discord.v1.gateway.GetStatsResponse
	(*UpdateStatusRequest_ActivityData)(nil),                  // 12: pylon.discord.v1.gateway.UpdateStatusRequest.ActivityData
	(*wrappers.UInt32Value)(nil),                              // 13: google.protobuf.UInt32Value
	(*timestamp.Timestamp)(nil),                               // 14: google.protobuf.Timestamp
	(model.PresenceData_OnlineStatus)(0),                      // 15: pylon.discord.v1.model.PresenceData.OnlineStatus
	(*model.UserData)(nil),                                    // 16: pylon.discord.v1.model.UserData
	(*model.GuildData)(nil),                                   // 17: pylon.discord.v1.model.GuildData
	(*model.EmojiData)(nil),                                   // 18: pylon.discord.v1.model.EmojiData
	(model.PresenceData_PresenceActivityData_ActivityType)(0), // 19: pylon.discord.v1.model.PresenceData.PresenceActivityData.ActivityType
}
var file_discord_v1_gateway_gateway_proto_depIdxs = []int32{
	13, // 0: pylon.discord.v1.gateway.UpdateStatusRequest.shard_id:type_name -> google.protobuf.UInt32Value
	14, // 1: pylon.discord.v1.gateway.UpdateStatusRequest.since:type_name -> google.protobuf.Timestamp
	12, // 2: pylon.discord.v1.gateway.UpdateStatusRequest.activities:type_name -> pylon.discord.v1.gateway.UpdateStatusRequest.ActivityData
	15, // 3: pylon.discord.v1.gateway.UpdateStatusRequest.status:type_name -> pylon.discord.v1.model.PresenceData.OnlineStatus
	16, // 4: pylon.discord.v1.gateway.FindUserResponse.user:type_name -> pylon.discord.v1.model.UserData
	17, // 5: pylon.discord.v1.gateway.GetUserMutualGuildsResponse.guilds:type_name -> pylon.discord.v1.model.GuildData
	18, // 6: pylon.discord.v1.gateway.FindEmojiResponse.emoji:type_name -> pylon.discord.v1.model.EmojiData
	19, // 7: pylon.discord.v1.gateway.UpdateStatusRequest.ActivityData.type:type_name -> pylon.discord.v1.model.PresenceData.PresenceActivityData.ActivityType
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_discord_v1_gateway_gateway_proto_init() }
func file_discord_v1_gateway_gateway_proto_init() {
	if File_discord_v1_gateway_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discord_v1_gateway_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVoiceStateRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVoiceStateResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMutualGuildsRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMutualGuildsResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmojiRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmojiResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsRequest); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
		file_discord_v1_gateway_gateway_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusRequest_ActivityData); i {
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
			RawDescriptor: file_discord_v1_gateway_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discord_v1_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_discord_v1_gateway_gateway_proto_depIdxs,
		MessageInfos:      file_discord_v1_gateway_gateway_proto_msgTypes,
	}.Build()
	File_discord_v1_gateway_gateway_proto = out.File
	file_discord_v1_gateway_gateway_proto_rawDesc = nil
	file_discord_v1_gateway_gateway_proto_goTypes = nil
	file_discord_v1_gateway_gateway_proto_depIdxs = nil
}
