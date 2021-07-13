//go:generate go run generate/merge.go
package main

import (
	"github.com/andersfylling/discordgateway/event"
	"github.com/germanoeich/nirn/cache"
)

type Cache struct {
	Guilds *cache.GuildCache
}

func (c *Cache) Setup(dispatcher *EventDispatcher) {
	dispatcher.Subscribe(event.GuildCreate, c.Guilds.HandleCreate)
	dispatcher.Subscribe(event.GuildUpdate, c.Guilds.HandleUpdate)
	dispatcher.Subscribe(event.GuildDelete, c.Guilds.HandleDelete)
	dispatcher.Subscribe(event.GuildEmojisUpdate, c.Guilds.HandleEmojisUpdate)
}
