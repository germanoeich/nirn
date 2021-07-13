package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildBanRemovePayload struct {
	GuildId discord.Snowflake `json:"guild_id"`
	User    *discord.User     `json:"user"`
}

func (g *GuildBanRemovePayload) ToProto() *event.GuildBanRemoveEvent_PayloadData {
	return &event.GuildBanRemoveEvent_PayloadData{
		User:    g.User.ToProto(),
		GuildId: g.GuildId.ToProto(),
	}
}

type GuildBanRemove struct {
	Scope   *event.EventScope
	Payload *GuildBanRemovePayload
}

func (m GuildBanRemove) ToProto() *event.GuildBanRemoveEvent {
	return &event.GuildBanRemoveEvent{
		Scope:   m.Scope,
		Payload: m.Payload.ToProto(),
	}
}
