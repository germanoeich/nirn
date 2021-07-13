package cache

import (
	"fmt"
	"github.com/andersfylling/discordgateway/event"
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/events"
	"sync"
)

type GuildCache struct {
	sync.RWMutex
	guilds map[discord.Snowflake]*discord.Guild
}

func NewGuildCache() *GuildCache {
	c := GuildCache{
		guilds: make(map[discord.Snowflake]*discord.Guild),
	}
	return &c
}

func (m *GuildCache) HandleCreate(e event.Type, event interface{}) {
	m.Lock()
	defer m.Unlock()
	ev := event.(events.GuildCreate)
	m.guilds[ev.Guild.ID] = ev.Guild
}

func (m *GuildCache) HandleUpdate(e event.Type, event interface{}) {
	m.Lock()
	defer m.Unlock()
	ev := event.(events.GuildUpdate)
	m.guilds[ev.Guild.ID].Merge(ev.Guild)
	fmt.Println(m.guilds[ev.Guild.ID].ToProto())
}

func (m *GuildCache) HandleDelete(e event.Type, event interface{}) {
	m.Lock()
	defer m.Unlock()
	ev := event.(events.GuildDelete)
	// Bot was removed from guild
	if ev.Guild.Unavailable == false {
		m.guilds[ev.Guild.ID] = nil
	} else {
		m.guilds[ev.Guild.ID].Unavailable = true
	}
}

func (m *GuildCache) HandleEmojisUpdate(e event.Type, event interface{}) {
	m.Lock()
	defer m.Unlock()
	ev := event.(events.GuildEmojisUpdate)
	m.guilds[ev.Payload.GuildId].Emojis = ev.Payload.Emojis
}

func (m *GuildCache) Get(id discord.Snowflake) discord.Guild {
	m.RLock()
	defer m.RUnlock()
	g := *m.guilds[id]
	return g
}