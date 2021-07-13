package main

import (
	"fmt"
	"github.com/andersfylling/discordgateway"
	"github.com/andersfylling/discordgateway/event"
	"github.com/andersfylling/discordgateway/json"
	"github.com/germanoeich/nirn/cache"
	"github.com/germanoeich/nirn/discord"
	"github.com/germanoeich/nirn/events"
	event2 "github.com/germanoeich/nirn/proto/discord/v1/event"
	"github.com/golang/protobuf/proto"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type RawEvent struct {
	Type event.Type
	Data []byte
}

type InternalEventHandler = func(e event.Type, event interface{})

type EventDispatcher struct {
	sync.Mutex
	Shard      *discordgateway.Shard
	ShardCount uint
	ShardId    uint
	BotId      discord.Snowflake
	NatsConn   *nats.Conn
	Logger     *logrus.Logger
	Cache	   *Cache
	eventc     chan RawEvent
	handlers   map[event.Type][]InternalEventHandler
}

func (ed *EventDispatcher) Setup() {
	ed.handlers = make(map[event.Type][]InternalEventHandler)
	ed.eventc = make(chan RawEvent, 100)
	ed.Cache = &Cache{
		Guilds: cache.NewGuildCache(),
	}
	ed.Cache.Setup(ed)
}

func (ed *EventDispatcher) Start() discordgateway.Handler {
	go ed.eventLoop()
	return func(shardId discordgateway.ShardID, e event.Type, message discordgateway.RawMessage) {
		ed.eventc <- RawEvent{Type: e, Data: message}
	}
}

func (ed *EventDispatcher) Subscribe(e event.Type, handler InternalEventHandler) {
	ed.Lock()
	defer ed.Unlock()
	ed.handlers[e] = append(ed.handlers[e], handler)
}

func (ed *EventDispatcher) eventLoop() {
	for e := range ed.eventc {
		switch e.Type {
		case event.MessageCreate:
			msg := &discord.Message{}
			err := json.Unmarshal(e.Data, msg)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.MessageCreate{
				Scope:   &event2.EventScope{BotId: uint64(ed.BotId), GuildId: msg.GuildID.ToProto()},
				Message: msg,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.MessageUpdate:
			msg := &discord.Message{}
			err := json.Unmarshal(e.Data, msg)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.MessageUpdate{
				Scope:          &event2.EventScope{BotId: uint64(ed.BotId), GuildId: msg.GuildID.ToProto()},
				PartialMessage: msg,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildCreate:
			guild := &discord.Guild{}
			err := json.Unmarshal(e.Data, guild)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.GuildCreate{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: guild.ID.ToProto()},
				Guild: guild,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildUpdate:
			guild := &discord.Guild{}
			err := json.Unmarshal(e.Data, guild)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			prevGuild := ed.Cache.Guilds.Get(guild.ID)
			ev := events.GuildUpdate{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: guild.ID.ToProto()},
				Guild: guild,
				CachedGuild: &prevGuild,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildDelete:
			guild := &discord.GuildUnavailable{}
			err := json.Unmarshal(e.Data, guild)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.GuildDelete{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: guild.ID.ToProto()},
				Guild: guild,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildBanCreate:
			data := &events.GuildBanAddPayload{}
			err := json.Unmarshal(e.Data, data)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.GuildBanAdd{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: data.GuildId.ToProto()},
				Payload: data,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildBanDelete:
			data := &events.GuildBanRemovePayload{}
			err := json.Unmarshal(e.Data, data)
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.GuildBanRemove{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: data.GuildId.ToProto()},
				Payload: data,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		case event.GuildEmojisUpdate:
			data := &events.GuildEmojisUpdatePayload{}
			err := json.Unmarshal(e.Data, data)
			fmt.Println(data.ToProto())
			if err != nil {
				ed.Logger.Error(err)
				return
			}

			ev := events.GuildEmojisUpdate{
				Scope: &event2.EventScope{BotId: uint64(ed.BotId), GuildId: data.GuildId.ToProto()},
				Payload: data,
			}

			protoBytes, perr := proto.Marshal(ev.ToProto())
			if perr != nil {
				ed.Logger.Error(perr)
				return
			}
			ed.dispatch(e.Type, protoBytes, ev)
		}
	}
}

func (ed *EventDispatcher) dispatch(e event.Type, protoBytes []byte, parsedEvent interface{}) {
	handlers, ok := ed.handlers[e]
	if ok {
		for _, h := range handlers {
			go h(e, parsedEvent)
		}
	}
	err := ed.NatsConn.Publish("events."+string(e), protoBytes)
	if err != nil {
		ed.Logger.Error(err)
	}
}
