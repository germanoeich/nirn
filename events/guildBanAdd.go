package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildBanAddPayload struct {
	GuildId discord.Snowflake `json:"guild_id"`
	User    *discord.User     `json:"user"`
}

func (g *GuildBanAddPayload) ToProto() *event.GuildBanAddEvent_PayloadData {
	return &event.GuildBanAddEvent_PayloadData{
		User:    g.User.ToProto(),
		GuildId: g.GuildId.ToProto(),
	}
}

type GuildBanAdd struct {
	Scope   *event.EventScope
	Payload *GuildBanAddPayload
}

func (m GuildBanAdd) ToProto() *event.GuildBanAddEvent {
	return &event.GuildBanAddEvent{
		Scope:   m.Scope,
		Payload: m.Payload.ToProto(),
	}
}
