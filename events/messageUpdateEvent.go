package events

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/proto/discord/v1/event"
)

type MessageUpdate struct {
	Scope *event.EventScope
	PartialMessage *discord.Message
	//CachedMessage *discord.Message
}

func (m MessageUpdate) ToProto() *event.MessageUpdateEvent {
	return &event.MessageUpdateEvent{
		Scope:            m.Scope,
		Payload:          &event.MessageUpdateEvent_Raw{Raw: m.PartialMessage.ToPartialProto()},
		PreviouslyCached: nil,
	}
}