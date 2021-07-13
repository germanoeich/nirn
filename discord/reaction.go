package discord

import (
	"github.com/germanoeich/nirn/proto/discord/v1/model"
)

// Reaction ...
// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count uint          `json:"count"`
	Me    bool          `json:"me"`
	Emoji *PartialEmoji `json:"Emoji"`
}

type ReactionArr []*Reaction

func (u *Reaction) ToProto() *model.MessageData_MessageReactionData {
	return &model.MessageData_MessageReactionData{
		Count: uint32(u.Count),
		Me:    u.Me,
		Emoji: u.Emoji.ToMessageProto(),
	}
}

func (a ReactionArr) ToProto() []*model.MessageData_MessageReactionData {
	var ret []*model.MessageData_MessageReactionData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}

