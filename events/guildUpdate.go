package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type GuildUpdate struct {
	Scope *event.EventScope
	Guild *discord.Guild
	CachedGuild *discord.Guild
}

func (m GuildUpdate) ToProto() *event.GuildUpdateEvent {
	return &event.GuildUpdateEvent{
		Scope:   m.Scope,
		Payload: m.Guild.ToProto(),
		PreviouslyCached: m.CachedGuild.ToProto(),
	}
}
