package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildCreate struct {
	Scope *event.EventScope
	Guild *discord.Guild
}

func (m GuildCreate) ToProto() *event.GuildCreateEvent {
	return &event.GuildCreateEvent{
		Scope:   m.Scope,
		Payload: m.Guild.ToProto(),
	}
}
