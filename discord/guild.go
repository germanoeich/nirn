package discord

import (
	"fmt"
	"github.com/germanoeich/nirn/proto/discord/v1/model"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

// NOTE! Credit for defining the Permission consts in a clean way goes to DiscordGo.
// This is pretty much a copy from their project. I would have made it a dependency if
// the consts were in a isolated sub-pkg. Note that in respect to their license, Disgord
// has no affiliation with DiscordGo.
//
// Source code reference:
//  https://github.com/bwmarrin/discordgo/blob/8325a6bf6dd6c91ed4040a1617b07287b8fb0eba/structs.go#L854

// PermissionBit is used to define the permission bit(s) which are set.
type PermissionBit uint64

func (b *PermissionBit) MarshalJSON() ([]byte, error) {
	str := strconv.FormatUint(uint64(*b), 10)
	return []byte(strconv.Quote(str)), nil
}

func (b *PermissionBit) UnmarshalJSON(bytes []byte) error {
	sb := string(bytes)
	str, err := strconv.Unquote(sb)
	if err != nil {
		return fmt.Errorf("PermissionBit#UnmarshalJSON - unable to unquote bytes{%s}: %w", sb, err)
	}

	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fmt.Errorf("PermissionBit#UnmarshalJSON - parsing string to uint64 failed: %w", err)
	}

	*b = PermissionBit(v)
	return nil
}

// Contains is used to check if the permission bits contains the bits specified.
func (b PermissionBit) Contains(Bits PermissionBit) bool {
	return (b & Bits) == Bits
}

// Constants for the different bit offsets of text channel permissions
const (
	PermissionReadMessages PermissionBit = 1 << (iota + 10)
	PermissionSendMessages
	PermissionSendTTSMessages
	PermissionManageMessages
	PermissionEmbedLinks
	PermissionAttachFiles
	PermissionReadMessageHistory
	PermissionMentionEveryone
	PermissionUseExternalEmojis
	PermissionViewGuildInsights
)

// Constants for the different bit offsets of voice permissions
const (
	PermissionVoiceConnect PermissionBit = 1 << (iota + 20)
	PermissionVoiceSpeak
	PermissionVoiceMuteMembers
	PermissionVoiceDeafenMembers
	PermissionVoiceMoveMembers
	PermissionVoiceUseVAD
	PermissionVoicePrioritySpeaker PermissionBit = 1 << (iota + 2)
	PermissionStream
)

// Constants for general management.
const (
	PermissionChangeNickname PermissionBit = 1 << (iota + 26)
	PermissionManageNicknames
	PermissionManageRoles
	PermissionManageWebhooks
	PermissionManageEmojis
)

// Constants for the different bit offsets of general permissions
const (
	PermissionCreateInstantInvite PermissionBit = 1 << iota
	PermissionKickMembers
	PermissionBanMembers
	PermissionAdministrator
	PermissionManageChannels
	PermissionManageServer
	PermissionAddReactions
	PermissionViewAuditLogs

	PermissionTextAll = PermissionReadMessages |
		PermissionSendMessages |
		PermissionSendTTSMessages |
		PermissionManageMessages |
		PermissionEmbedLinks |
		PermissionAttachFiles |
		PermissionReadMessageHistory |
		PermissionMentionEveryone
	PermissionAllVoice = PermissionVoiceConnect |
		PermissionVoiceSpeak |
		PermissionVoiceMuteMembers |
		PermissionVoiceDeafenMembers |
		PermissionVoiceMoveMembers |
		PermissionVoiceUseVAD
	PermissionChannelAll = PermissionTextAll |
		PermissionAllVoice |
		PermissionCreateInstantInvite |
		PermissionManageRoles |
		PermissionManageChannels |
		PermissionAddReactions |
		PermissionViewAuditLogs
	PermissionAll = PermissionChannelAll |
		PermissionKickMembers |
		PermissionBanMembers |
		PermissionManageServer |
		PermissionAdministrator
)

// GuildUnavailable is a partial Guild object.
type GuildUnavailable struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"` // ?*|
}

func (g GuildUnavailable) ToProto() *model.GuildData {
	return &model.GuildData{
		Id: g.ID.ToProto(),
		Unavailable: g.Unavailable,
	}
}

// Guild Guilds in Discord represent an isolated collection of Users and Channels,
//  and are often referred to as "servers" in the UI.
// https://discord.com/developers/docs/resources/guild#guild-object
// Fields with `*` are only sent within the GUILD_CREATE event
// reviewed: 2018-08-25
// ++gen_merge
type Guild struct {
	ID                          Snowflake                     `json:"id"`
	ApplicationID               Snowflake                     `json:"application_id"` //   |?
	Name                        string                        `json:"name" merge:"true"`
	Icon                        string                        `json:"icon" merge:"true"`   //  |?, icon hash
	Splash                      string                        `json:"splash" merge:"true"` //  |?, image hash
	Owner                       bool                          `json:"owner,omitempty"`     // ?|
	OwnerID                     Snowflake                     `json:"owner_id" merge:"true"`
	Permissions                 PermissionBit                 `json:"permissions,omitempty"` // ?|, permission flags for connected user `/users/@me/guilds`
	Region                      string                        `json:"region" merge:"true"`
	AfkChannelID                Snowflake                     `json:"afk_channel_id" merge:"true"` // |?
	AfkTimeout                  uint                          `json:"afk_timeout" merge:"true"`
	VerificationLevel           VerificationLvl               `json:"verification_level" merge:"true"`
	DefaultMessageNotifications DefaultMessageNotificationLvl `json:"default_message_notifications" merge:"true"`
	ExplicitContentFilter       ExplicitContentFilterLvl      `json:"explicit_content_filter" merge:"true"`
	Roles                       []*Role                       `json:"roles"`
	Emojis                      []*Emoji                      `json:"emojis" merge:"true"`
	Features                    []string                      `json:"features" merge:"true"`
	MFALevel                    MFALvl                        `json:"mfa_level" merge:"true"`
	WidgetEnabled               bool                          `json:"widget_enabled,omit_empty" merge:"true"`    //   |
	WidgetChannelID             Snowflake                     `json:"widget_channel_id,omit_empty" merge:"true"` //   |?
	SystemChannelID             Snowflake                     `json:"system_channel_id,omitempty" merge:"true"`  //   |?
	DiscoverySplash             string                        `json:"discovery_splash,omitempty" merge:"true"`
	VanityUrl                   string                        `json:"vanity_url_code,omitempty" merge:"true"`
	Description                 string                        `json:"description,omitempty" merge:"true"`
	Banner                      string                        `json:"banner,omitempty" merge:"true"`
	PremiumTier                 PremiumTier                   `json:"premium_tier" merge:"true"`
	PremiumSubscriptionCount    uint                          `json:"premium_subscription_count,omitempty" merge:"true"`
	MaxPresences                uint                          `json:"max_presences,omitempty"  merge:"true"`
	MaxMembers                  uint                          `json:"max_members,omitempty" merge:"true"`
	PublicUpdatesChannel        Snowflake                     `json:"public_updates_channel_id,omitempty" merge:"true"`
	RulesChannel                Snowflake                     `json:"rules_channel_id,omitempty" merge:"true"`
	JoinedAt                    *Time                         `json:"joined_at,omitempty" merge:"true"`                 // ?*|
	Large                       bool                          `json:"large,omitempty" merge:"true"`        // ?*|
	Unavailable                 bool                          `json:"unavailable" merge:"true"`                         // ?*| omitempty?
	MemberCount                 uint                          `json:"member_count,omitempty"` // ?*|
	VoiceStates                 []*VoiceState                 `json:"voice_states,omitempty"`              // ?*|
	Members                     []*Member                     `json:"members,omitempty"`                   // ?*|
	Channels                    []*Channel                    `json:"channels,omitempty"`                  // ?*|
	Presences                   []*UserPresence               `json:"presences,omitempty"`                 // ?*|
}

func (g Guild) ToProto() *model.GuildData {
	var joinedAt *timestamppb.Timestamp
	if g.JoinedAt != nil {
		joinedAt = g.JoinedAt.ToProto()
	}
	return &model.GuildData{
		Id:                          g.ID.ToProto(),
		Name:                        g.Name,
		Icon:                        &wrappers.StringValue{Value: g.Icon},
		Region:                      g.Region,
		AfkChannelId:                g.AfkChannelID.ToOptionalProto(),
		OwnerId:                     g.OwnerID.ToProto(),
		JoinedAt:                    joinedAt,
		Splash:                      &wrappers.StringValue{Value: g.Splash},
		DiscoverySplash:             &wrappers.StringValue{Value: g.DiscoverySplash},
		AfkTimeout:                  uint32(g.AfkTimeout),
		MemberCount:                 uint32(g.MemberCount),
		VerificationLevel:           uint32(g.VerificationLevel),
		DefaultMessageNotifications: uint32(g.DefaultMessageNotifications),
		ExplicitContentFilter:       uint32(g.ExplicitContentFilter),
		Features:                    g.Features,
		MfaLevel:                    uint32(g.MFALevel),
		WidgetEnabled:               g.WidgetEnabled,
		WidgetChannelId:             g.WidgetChannelID.ToOptionalProto(),
		SystemChannelId:             g.SystemChannelID.ToOptionalProto(),
		VanityUrlCode:               &wrappers.StringValue{Value: g.VanityUrl},
		Description:                 &wrappers.StringValue{Value: g.Description},
		Banner:                      &wrappers.StringValue{Value: g.Banner},
		PremiumTier:                 uint32(g.PremiumTier),
		PremiumSubscriptionCount:    uint32(g.PremiumSubscriptionCount),
		Unavailable:                 g.Unavailable,
	}
}

// --------------

// PartialBan is used by audit logs
type PartialBan struct {
	Reason                 string
	BannedUserID           Snowflake
	ModeratorResponsibleID Snowflake
}

// Ban https://discord.com/developers/docs/resources/guild#ban-object
type Ban struct {
	Reason string `json:"reason"`
	User   *User  `json:"user"`
}

// ------------

// GuildEmbed https://discord.com/developers/docs/resources/guild#guild-embed-object
type GuildEmbed struct {
	Enabled   bool      `json:"enabled"`
	ChannelID Snowflake `json:"channel_id"`
}

// -------

// Integration https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	ID                Snowflake           `json:"id"`
	Name              string              `json:"name"`
	Type              string              `json:"type"`
	Enabled           bool                `json:"enabled"`
	Syncing           bool                `json:"syncing"`
	RoleID            Snowflake           `json:"role_id"`
	ExpireBehavior    int                 `json:"expire_behavior"`
	ExpireGracePeriod int                 `json:"expire_grace_period"`
	User              *User               `json:"user"`
	Account           *IntegrationAccount `json:"account"`
}

// IntegrationAccount https://discord.com/developers/docs/resources/guild#integration-account-object
type IntegrationAccount struct {
	ID   string `json:"id"`   // id of the account
	Name string `json:"name"` // name of the account
}

// -------

// Member https://discord.com/developers/docs/resources/guild#guild-member-object
type Member struct {
	GuildID      Snowflake      `json:"guild_id,omitempty"`
	User         *User          `json:"user"`
	Nick         string         `json:"nick,omitempty"`
	Roles        SnowflakeArray `json:"roles"`
	JoinedAt     Time           `json:"joined_at,omitempty"`
	PremiumSince Time           `json:"premium_since,omitempty"`
	Deaf         bool           `json:"deaf"`
	Mute         bool           `json:"mute"`
	Pending      bool           `json:"pending"`
}

func (m *Member) toProto(userId uint64) *model.MemberData {
	if m == nil {
		return nil
	}

	return &model.MemberData{
		Id:           userId,
		GuildId:      m.GuildID.ToProto(),
		User:         m.User.ToProto(),
		Nick:         &wrappers.StringValue{Value: m.Nick},
		Roles:        m.Roles.ToProto(),
		JoinedAt:     m.JoinedAt.ToProto(),
		PremiumSince: m.PremiumSince.ToProto(),
		// TODO: Calculate this
		Permissions: 0,
		Pending:     &wrappers.BoolValue{Value: m.Pending},
	}
}

func (m *Member) ToProtoForMessageCreateOrUpdate(userId uint64) *model.MemberData {
	return m.toProto(userId)
}

func (m *Member) ToProto() *model.MemberData {
	var userId uint64 = 0
	if m.User != nil {
		userId = m.User.ID.ToProto()
	}
	return m.toProto(userId)
}
