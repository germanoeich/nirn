package cache

import (
	"github.com/andersfylling/discordgateway/event"
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/events"
	"sync"
)

type ChannelCache struct {
	sync.RWMutex
	channels map[discord.Snowflake]*discord.Channel
	channelsPerGuild map[discord.Snowflake][]*discord.Channel
}

func NewChannelCache() *ChannelCache {
	c := ChannelCache{
		channels: make(map[discord.Snowflake]*discord.Channel),
		channelsPerGuild: make(map[discord.Snowflake][]*discord.Channel),
	}
	return &c
}

func (m ChannelCache) HandleGuildCreate(e event.Type, event interface{}) {
	m.Lock()
	defer m.Unlock()
	ev := event.(events.GuildCreate)
	for _, ch := range ev.Guild.Channels {
		m.channels[ch.ID] = ch
	}
}

func (m ChannelCache) Get(id discord.Snowflake) *discord.Channel {
	m.RLock()
	defer m.RUnlock()
	g := m.channels[id]
	return g
}