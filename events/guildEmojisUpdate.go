package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildEmojisUpdatePayload struct {
	GuildId discord.Snowflake `json:"guild_id"`
	Emojis  discord.EmojiArr  `json:"emojis"`
}

func (g *GuildEmojisUpdatePayload) ToProto() *event.GuildEmojisUpdateEvent_PayloadData {
	return &event.GuildEmojisUpdateEvent_PayloadData{
		Emojis:  g.Emojis.ToProto(g.GuildId),
		GuildId: g.GuildId.ToProto(),
	}
}

type GuildEmojisUpdate struct {
	Scope   *event.EventScope
	Payload *GuildEmojisUpdatePayload
}

func (m *GuildEmojisUpdate) ToProto() *event.GuildEmojisUpdateEvent {
	return &event.GuildEmojisUpdateEvent{
		Scope:   m.Scope,
		Payload: m.Payload.ToProto(),
	}
}
