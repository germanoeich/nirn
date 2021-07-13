package discord

import (
	event2 "github.com/germanoeich/nirn/proto/discord/v1/event"
	"github.com/germanoeich/nirn/proto/discord/v1/model"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// different message activity types
const (
	_ = iota
	MessageActivityTypeJoin
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

// MessageFlag https://discord.com/developers/docs/resources/channel#message-object-message-flags
type MessageFlag uint

const (
	MessageFlagCrossposted MessageFlag = 1 << iota
	MessageFlagIsCrosspost
	MessageFlagSupressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
)

// The different message types usually generated by Discord. eg. "a new user joined"
type MessageType uint // TODO: once auto generated, un-export this.

const (
	MessageTypeUnknown MessageType = iota
	MessageTypeDefault
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	_
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	_
	_
	_
	MessageTypeReply
	MessageTypeApplicationCommand
)

const (
	AttachmentSpoilerPrefix = "SPOILER_"
)

// MessageActivity https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	Type    int    `json:"type"`
	PartyID string `json:"party_id"`
}

func (m MessageActivity) ToProto() *model.MessageData_MessageActivityData {
	return &model.MessageData_MessageActivityData{
		Type:    uint32(m.Type),
		PartyId: m.PartyID,
	}
}

type MentionChannel struct {
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

type MentionChannelArr []*MentionChannel

func (m MentionChannel) ToProto() *model.MessageData_MessageMentionChannelData {
	return &model.MessageData_MessageMentionChannelData{
		Id:      uint64(m.ID),
		GuildId: uint64(m.GuildID),
		Type:    model.ChannelData_ChannelType(m.Type + 1),
		Name:    m.Name,
	}
}

func (a MentionChannelArr) ToProto() []*model.MessageData_MessageMentionChannelData {
	var ret []*model.MessageData_MessageMentionChannelData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}

type MessageReference struct {
	MessageID Snowflake `json:"message_id"`
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id"`
}

func (m *MessageReference) ToProto() *model.MessageData_MessageReferenceData {
	if m == nil {
		return nil
	}
	return &model.MessageData_MessageReferenceData{
		MessageId: uint64(m.MessageID),
		ChannelId: uint64(m.ChannelID),
		GuildId:   m.GuildID.ToProto(),
	}
}

// MessageApplication https://discord.com/developers/docs/resources/channel#message-object-message-application-structure
type MessageApplication struct {
	ID          Snowflake `json:"id"`
	CoverImage  string    `json:"cover_image"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Name        string    `json:"name"`
}

func (m MessageApplication) ToProto() *model.MessageData_MessageApplicationData {
	return &model.MessageData_MessageApplicationData{
		Id:          m.ID.ToProto(),
		CoverImage:  m.CoverImage,
		Description: m.Description,
		Icon:        m.Icon,
		Name:        m.Name,
	}
}

type MessageStickerFormatType int

const (
	MessageStickerFormatUnknown MessageStickerFormatType = iota
	MessageStickerFormatPNG
	MessageStickerFormatAPNG
	MessageStickerFormatLOTTIE
)

type MessageSticker struct {
	ID           Snowflake                `json:"id"`
	PackID       Snowflake                `json:"pack_id"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
	Tags         string                   `json:"tags"`
	Asset        string                   `json:"asset"`
	PreviewAsset string                   `json:"preview_asset"`
	FormatType   MessageStickerFormatType `json:"format_type"`
}

type MessageStickerArr []*MessageSticker

func (m MessageSticker) ToProto() *model.MessageData_MessageStickerData {
	return &model.MessageData_MessageStickerData{
		Id:          m.ID.ToProto(),
		PackId:      uint64(m.PackID),
		Name:        m.Name,
		Description: m.Description,
		Tags:        m.Tags,
		Asset:       m.Asset,
		FormatType:  model.MessageData_MessageStickerData_MessageStickerFormatType(m.FormatType),
	}
}

func (a MessageStickerArr) ToProto() []*model.MessageData_MessageStickerData {
	var ret []*model.MessageData_MessageStickerData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}

// Message https://discord.com/developers/docs/resources/channel#message-object-message-structure
type Message struct {
	ID                Snowflake          `json:"id"`
	ChannelID         Snowflake          `json:"channel_id"`
	GuildID           Snowflake          `json:"guild_id"`
	Author            *User              `json:"author"`
	Member            *Member            `json:"member"`
	Content           string             `json:"content"`
	Timestamp         Time               `json:"timestamp"`
	EditedTimestamp   Time               `json:"edited_timestamp"` // ?
	Tts               bool               `json:"tts"`
	MentionEveryone   bool               `json:"mention_everyone"`
	Mentions          UserArr            `json:"mentions"`
	MentionRoles      SnowflakeArray     `json:"mention_roles"`
	// TODO: confirm this is being parsed correctly
	MentionChannels   MentionChannelArr  `json:"mention_channels"`
	Attachments       AttachmentArray    `json:"attachments"`
	Embeds            EmbedArray         `json:"embeds"`
	Reactions         ReactionArr        `json:"reactions"` // ?
	Nonce             interface{}        `json:"nonce"`     // NOT A SNOWFLAKE! DONT TOUCH!
	Pinned            bool               `json:"pinned"`
	WebhookID         Snowflake          `json:"webhook_id"` // ?
	Type              MessageType        `json:"type"`
	Activity          MessageActivity    `json:"activity"`
	Application       MessageApplication `json:"application"`
	MessageReference  *MessageReference  `json:"message_reference"`
	ReferencedMessage *Message           `json:"referenced_message"`
	Flags             MessageFlag        `json:"flags"`
	Stickers          MessageStickerArr  `json:"stickers"`

	// SpoilerTagContent is only true if the entire message text is tagged as a spoiler (aka completely wrapped in ||)
	SpoilerTagContent        bool `json:"-"`
	SpoilerTagAllAttachments bool `json:"-"`
	HasSpoilerImage          bool `json:"-"`
}

func (m Message) ToProto() *model.MessageData {
	return m.toProto(false)
}

func (m Message) ToProtoForMessageCreateOrUpdate() *model.MessageData {
	return m.toProto(true)
}

func (m Message) toProto(forMessageCreateOrUpdate bool) *model.MessageData {
	var member *model.MemberData
	if forMessageCreateOrUpdate {
		member = m.Member.ToProtoForMessageCreateOrUpdate(uint64(m.Author.ID))
	} else {
		member = m.Member.ToProto()
	}
	return &model.MessageData{
		Id:               m.ID.ToProto(),
		ChannelId:        m.ChannelID.ToProto(),
		GuildId:          m.GuildID.ToOptionalProto(),
		Content:          m.Content,
		Timestamp:        m.Timestamp.ToProto(),
		EditedTimestamp:  m.EditedTimestamp.ToProto(),
		MentionRoles:     m.MentionRoles.ToOptionalProto(),
		Tts:              m.Tts,
		MentionEveryone:  m.MentionEveryone,
		Attachments:      m.Attachments.ToProto(),
		Embeds:           m.Embeds.ToProto(),
		Mentions:         m.Mentions.ToMentionProto(),
		Reactions:        m.Reactions.ToProto(),
		Pinned:           m.Pinned,
		Type:             model.MessageData_MessageType(m.Type + 1),
		MentionChannels:  m.MentionChannels.ToProto(),
		Flags:            uint32(m.Flags),
		Activity:         m.Activity.ToProto(),
		Application:      m.Application.ToProto(),
		MessageReference: m.MessageReference.ToProto(),
		Author:           m.Author.ToProto(),
		Member:           member,
		WebhookId:        m.WebhookID.ToOptionalProto(),
		Stickers:         m.Stickers.ToProto(),
		// TODO: Map these
		Interaction: nil,
		Thread:      nil,
		Components:  nil,
	}
}

func (m Message) ToPartialProto() *event2.MessageUpdateEvent_PayloadData {
	ret := &event2.MessageUpdateEvent_PayloadData{
		Id:               m.ID.ToProto(),
		ChannelId:        m.ChannelID.ToProto(),
		GuildId:          m.GuildID.ToOptionalProto(),
		Content:          &wrappers.StringValue{Value: m.Content},
		EditedTimestamp:  m.EditedTimestamp.ToProto(),
		MentionRoles:     m.MentionRoles.ToOptionalProto(),
		Tts:              &wrappers.BoolValue{Value: m.Tts},
		MentionEveryone:  &wrappers.BoolValue{Value: m.MentionEveryone},
		Attachments:      &event2.MessageUpdateEvent_PayloadData_MessageAttachmentListValue{Values: m.Attachments.ToProto()},
		Embeds:           &event2.MessageUpdateEvent_PayloadData_MessageEmbedListValue{Values: m.Embeds.ToProto()},
		Mentions:         &event2.MessageUpdateEvent_PayloadData_MessageMentionListValue{Values: m.Mentions.ToMentionProto()},
		Reactions:        &event2.MessageUpdateEvent_PayloadData_MessageReactionListValue{Values: m.Reactions.ToProto()},
		Pinned:           &wrappers.BoolValue{Value: m.Pinned},
		Type:             &event2.MessageUpdateEvent_PayloadData_MessageTypeValue{Value: model.MessageData_MessageType(m.Type + 1) },
		MentionChannels:  &event2.MessageUpdateEvent_PayloadData_MessageMentionChannelListValue{Values: m.MentionChannels.ToProto()},
		Flags:            &wrappers.UInt32Value{ Value: uint32(m.Flags) },
		Activity:         &event2.MessageUpdateEvent_PayloadData_MessageActivityValue{Value: m.Activity.ToProto() },
		Application:      &event2.MessageUpdateEvent_PayloadData_MessageApplicationValue{Value: m.Application.ToProto()},
		MessageReference: &event2.MessageUpdateEvent_PayloadData_MessageReferenceValue{Value: m.MessageReference.ToProto()},
		Author:           &event2.MessageUpdateEvent_PayloadData_MessageAuthorValue{Value: m.Author.ToProto()},
		WebhookId:        m.WebhookID.ToOptionalProto(),
	}

	if m.Member != nil && m.Author != nil {
		ret.Member = &event2.MessageUpdateEvent_PayloadData_MessageMemberValue{Value: m.Member.ToProtoForMessageCreateOrUpdate(uint64(m.Author.ID))}
	}

	return ret
}
