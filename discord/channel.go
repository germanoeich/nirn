package discord

import "github.com/germanoeich/nirn/proto/discord/v1/model"

// Channel types
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType uint

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

// Deprecated: use PermissionOverwrite* instead (note the Type keyword is removed)
// PermissionOverwriteTypeMember => PermissionOverwriteMember
const (
	PermissionOverwriteTypeRole uint8 = iota
	PermissionOverwriteTypeMember
)

// Attachment https://discord.com/developers/docs/resources/channel#attachment-object
type Attachment struct {
	ID       Snowflake `json:"id"`
	Filename string    `json:"filename"`
	Size     uint      `json:"size"`
	URL      string    `json:"url"`
	ProxyURL string    `json:"proxy_url"`
	Height   uint      `json:"height"`
	Width    uint      `json:"width"`

	SpoilerTag bool `json:"-"`
}


func (a Attachment) ToProto() *model.MessageData_MessageAttachmentData {
	return &model.MessageData_MessageAttachmentData{
		Id:       a.ID.ToProto(),
		Filename: a.Filename,
		Size:     uint64(a.Size),
		Url:      a.URL,
		ProxyUrl: a.ProxyURL,
		Width:    uint64(a.Width),
		Height:   uint64(a.Height),
	}
}

type AttachmentArray []*Attachment

func (a AttachmentArray) ToProto() []*model.MessageData_MessageAttachmentData {
	var ret []*model.MessageData_MessageAttachmentData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}

type PermissionOverwriteType uint8

const (
	PermissionOverwriteRole PermissionOverwriteType = iota
	PermissionOverwriteMember
)

// PermissionOverwrite https://discord.com/developers/docs/resources/channel#overwrite-object
//
// WARNING! Discord is bugged, and the Type field needs to be a string to read Permission Overwrites from audit log
type PermissionOverwrite struct {
	ID    Snowflake               `json:"id"` // role or user id
	Type  PermissionOverwriteType `json:"type"`
	Allow PermissionBit           `json:"allow"`
	Deny  PermissionBit           `json:"deny"`
}

// type ChannelDeleter interface { DeleteChannel(id Snowflake) (err error) }
// type ChannelUpdater interface {}

// PartialChannel ...
// example of partial channel
// // "channel": {
// //   "id": "165176875973476352",
// //   "name": "illuminati",
// //   "type": 0
// // }
type PartialChannel struct {
	ID   Snowflake   `json:"id"`
	Name string      `json:"name"`
	Type ChannelType `json:"type"`
}

// Channel ...
type Channel struct {
	ID                   Snowflake             `json:"id"`
	Type                 ChannelType           `json:"type"`
	GuildID              Snowflake             `json:"guild_id,omitempty"`
	Position             int                   `json:"position,omitempty"` // can be less than 0
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Topic                string                `json:"topic,omitempty"`
	NSFW                 bool                  `json:"nsfw,omitempty"`
	LastMessageID        Snowflake             `json:"last_message_id,omitempty"`
	Bitrate              uint                  `json:"bitrate,omitempty"`
	UserLimit            uint                  `json:"user_limit,omitempty"`
	RateLimitPerUser     uint                  `json:"rate_limit_per_user,omitempty"`
	Recipients           []*User               `json:"recipients,omitempty"` // empty if not DM/GroupDM
	Icon                 string                `json:"icon,omitempty"`
	OwnerID              Snowflake             `json:"owner_id,omitempty"`
	ApplicationID        Snowflake             `json:"application_id,omitempty"`
	ParentID             Snowflake             `json:"parent_id,omitempty"`
	LastPinTimestamp     Time                  `json:"last_pin_timestamp,omitempty"`
}


// GetPermissions is used to get a members permissions in a channel.
//func (c *Channel) GetPermissions(ctx context.Context, s GuildQueryBuilderCaller, member *Member, flags ...Flag) (permissions PermissionBit, err error) {
//	// Get the guild permissions.
//	permissions, err = member.GetPermissions(ctx, s, flags...)
//	if err != nil {
//		return 0, err
//	}
//
//	// Handle permission overwrites.
//	apply := func(o PermissionOverwrite) {
//		permissions |= o.Allow
//		permissions &= (-o.Deny) - 1
//	}
//	for _, overwrite := range c.PermissionOverwrites {
//		if overwrite.Type == PermissionOverwriteMember {
//			// This is a member. Is it me?
//			if overwrite.ID == member.UserID {
//				// It is! Time to apply the overwrites.
//				apply(overwrite)
//			}
//			continue
//		}
//
//		for _, role := range member.Roles {
//			if role == overwrite.ID {
//				apply(overwrite)
//				break
//			}
//		}
//	}
//
//	// Return the result.
//	return
//}
