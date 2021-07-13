package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type MessageCreate struct {
	Scope *event.EventScope
	Message *discord.Message
}

func (m MessageCreate) ToProto() *event.MessageCreateEvent {
	return &event.MessageCreateEvent{
		Scope:       m.Scope,
		MessageData: m.Message.ToProtoForMessageCreateOrUpdate(),
	}
}