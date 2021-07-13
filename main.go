package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/andersfylling/discordgateway"
	"github.com/andersfylling/discordgateway/event"
	"github.com/andersfylling/discordgateway/log"
	event2 "github.com/germanoeich/nirn/proto/discord/v1/event"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"os"
)

var logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

var natsClient *nats.Conn

func newShard(shardId uint, shardCount uint, ctx context.Context) {
	ed := EventDispatcher{
		NatsConn: natsClient,
		Logger:   logger,
		ShardId:  shardId,
		ShardCount: shardCount,
	}
	ed.Setup()

	shard, err := discordgateway.NewShard(ed.Start(), &discordgateway.ShardConfig{
		BotToken: os.Getenv("TOKEN"),
		GuildEvents: []event.Type{
			event.MessageCreate,
			event.MessageUpdate,
			event.GuildCreate,
			event.GuildUpdate,
			event.GuildBanCreate,
			event.GuildBanDelete,
			event.GuildEmojisUpdate,
			//event.GuildDelete,
			//event.GuildMemberCreate,
			//event.GuildMemberUpdate,
			//event.GuildMemberDelete,
			//event.GuildMembersChunk,
		},
		DMEvents:            nil,
		TotalNumberOfShards: shardCount,
		ShardID:             shardId,
		IdentifyProperties: discordgateway.GatewayIdentifyProperties{
			OS:      "linux",
			Browser: "github.com/andersfylling/discordgateway v0",
			Device:  "tester",
		},
	})
	if err != nil {
		logger.Fatal(err)
	}

reconnect:
	conn, err := shard.Dial(ctx, "wss://gateway.discord.gg/?v=9&encoding=json")
	if err != nil {
		logger.Error("failed to open websocket connection.", err)
	}

	if op, err := shard.EventLoop(ctx, conn); err != nil {
		var discordErr *discordgateway.CloseError
		if errors.As(err, &discordErr) {
			switch discordErr.Code {
			case 1001, 4000: // will initiate a resume
				fallthrough
			case 4007, 4009: // will do a fresh identify
				goto reconnect
			case 4001, 4002, 4003, 4004, 4005, 4008, 4010, 4011, 4012, 4013, 4014:
			default:
				logger.Errorf("unhandled close error, with discord op code(%d): %d", op, discordErr.Code)
			}
		}
		var errClosed *discordgateway.ErrClosed
		if errors.As(err, &errClosed) || errors.Is(err, net.ErrClosed) || errors.Is(err, io.ErrClosedPipe) {
			logger.Debug("connection closed/lost .. will try to reconnect")
			goto reconnect
		}
		logger.Error(err)
	} else {
		goto reconnect
	}
}

func main() {
	log.LogInstance = logger
	logger.Info("Connecting to NATS")
	n, err := nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Connected to NATS")
	natsClient = n
	for i := uint(0); i < 1; i++ {
		logger.Info("Starting shard", i)

		go newShard(i, 1, context.Background())

		logger.Info("Started shard", i)
		if err != nil {
			fmt.Println(err)
		}
	}
	//natsClient.Subscribe("events.MESSAGE_CREATE", func(m *nats.Msg) {
	//	msg := &event2.MessageCreateEvent{}
	//	proto.Unmarshal(m.Data, msg)
	//	logger.Debug("MESSAGE_CREATE:\n", msg)
	//})
	//
	//natsClient.Subscribe("events.MESSAGE_UPDATE", func(m *nats.Msg) {
	//	msg := &event2.MessageUpdateEvent{}
	//	proto.Unmarshal(m.Data, msg)
	//	logger.Debug("MESSAGE_UPDATE:\n", msg)
	//})
	//
	//natsClient.Subscribe("events.GUILD_CREATE", func(m *nats.Msg) {
	//	msg := &event2.GuildCreateEvent{}
	//	proto.Unmarshal(m.Data, msg)
	//	logger.Debug("GUILD_CREATE:\n", msg)
	//})

	natsClient.Subscribe("events.GUILD_UPDATE", func(m *nats.Msg) {
		msg := &event2.GuildUpdateEvent{}
		proto.Unmarshal(m.Data, msg)
		logger.Debug("GUILD_UPDATE:\n", msg)
	})

	natsClient.Subscribe("events.GUILD_BAN_ADD", func(m *nats.Msg) {
		msg := &event2.GuildBanAddEvent{}
		proto.Unmarshal(m.Data, msg)
		logger.Debug("GUILD_BAN_ADD:\n", msg)
	})

	natsClient.Subscribe("events.GUILD_BAN_REMOVE", func(m *nats.Msg) {
		msg := &event2.GuildBanRemoveEvent{}
		proto.Unmarshal(m.Data, msg)
		logger.Debug("GUILD_BAN_REMOVE:\n", msg)
	})

	natsClient.Subscribe("events.GUILD_EMOJIS_UPDATE", func(m *nats.Msg) {
		msg := &event2.GuildEmojisUpdateEvent{}
		proto.Unmarshal(m.Data, msg)
		logger.Debug("GUILD_EMOJIS_UPDATE:\n", msg)
	})
	select {}
}
