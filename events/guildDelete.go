package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildDelete struct {
	Scope *event.EventScope
	Guild *discord.GuildUnavailable
}

func (m GuildDelete) ToProto() *event.GuildDeleteEvent {
	return &event.GuildDeleteEvent{
		Scope:   m.Scope,
		Payload: m.Guild.ToProto(),
	}
}
