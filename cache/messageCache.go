package cache

import (
	"github.com/andersfylling/discordgateway/event"
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/events"
)

type MessageCache struct {
	messages map[discord.Snowflake]*discord.Message
}

func (m MessageCache) Setup() {
	m.messages = make(map[discord.Snowflake]*discord.Message)
	//dispatcher.Subscribe(event.MessageCreate, m.handleCreate)
}

func (m MessageCache) handleCreate(e event.Type, event interface{}) {
	ev := event.(events.MessageCreate)
	m.messages[ev.Message.ID] = ev.Message
}