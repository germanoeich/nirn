package discord

import (
	"github.com/germanoeich/nirn/proto/discord/v1/model"
	"unsafe"
)

// Emoji ...
type Emoji struct {
	ID            Snowflake      `json:"id"`
	Name          string         `json:"name"`
	Roles         SnowflakeArray `json:"roles,omitempty"`
	User          *User          `json:"user,omitempty"` // the user who created the emoji
	RequireColons bool           `json:"require_colons,omitempty"`
	Managed       bool           `json:"managed,omitempty"`
	Animated      bool           `json:"animated,omitempty"`
	Available     bool           `json:"available,omitempty"`
}

type PartialEmoji = Emoji
type EmojiArr []*Emoji

func (a EmojiArr) ToProto(guildId Snowflake) []*model.EmojiData {
	var ret []*model.EmojiData
	for _, el := range a {
		ret = append(ret, el.ToProto(guildId))
	}
	return ret
}


func (u *Emoji) ToProto(guildId Snowflake) *model.EmojiData {
	userId := uint64(0)
	if u.User != nil {
		userId = u.User.ID.ToProto()
	}
	return &model.EmojiData{
		Id:            uint64(u.ID),
		GuildId:       guildId.ToProto(),
		Name:          u.Name,
		Animated:      u.Animated,
		Roles:         *(*[]uint64)(unsafe.Pointer(&u.Roles)),
		Managed:       u.Managed,
		RequireColons: u.RequireColons,
		Available:     u.Available,
		UserId:        userId,
	}
}

func (u *Emoji) ToMessageProto() *model.MessageData_MessageReactionEmojiData {
	return &model.MessageData_MessageReactionEmojiData{
		Id:            uint64(u.ID),
		Name:          u.Name,
		Animated:      u.Animated,
	}
}

