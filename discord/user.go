package discord

import (
	"github.com/germanoeich/nirn/proto/discord/v1/model"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
	Web     string `json:"web"`
}

// ActivityParty ...
type ActivityParty struct {
	ID   string `json:"id,omitempty"`   // the id of the party
	Size []int  `json:"size,omitempty"` // used to show the party's current and maximum size
}

// ActivityAssets ...
type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"` // the id for a large asset of the activity, usually a snowflake
	LargeText  string `json:"large_text,omitempty"`  //text displayed when hovering over the large image of the activity
	SmallImage string `json:"small_image,omitempty"` // the id for a small asset of the activity, usually a snowflake
	SmallText  string `json:"small_text,omitempty"`  //	text displayed when hovering over the small image of the activity
}

// ActivitySecrets ...
type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`     // the secret for joining a party
	Spectate string `json:"spectate,omitempty"` // the secret for spectating a game
	Match    string `json:"match,omitempty"`    // the secret for a specific instanced match
}

// ActivityEmoji ...
type ActivityEmoji struct {
	Name     string    `json:"name"`
	ID       Snowflake `json:"id,omitempty"`
	Animated bool      `json:"animated,omitempty"`
}

// ActivityTimestamp ...
type ActivityTimestamp struct {
	Start int `json:"start,omitempty"` // unix time (in milliseconds) of when the activity started
	End   int `json:"end,omitempty"`   // unix time (in milliseconds) of when the activity ends
}

// ######################
// ##
// ## Activity
// ##
// ######################

// ActivityType https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type ActivityType uint

const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	_
	ActivityTypeCustom
	ActivityTypeCompeting
)

// ActivityFlag https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
type ActivityFlag uint

// flags for the Activity object to signify the type of action taken place
const (
	ActivityFlagInstance ActivityFlag = 1 << iota
	ActivityFlagJoin
	ActivityFlagSpectate
	ActivityFlagJoinRequest
	ActivityFlagSync
	ActivityFlagPlay
)

// Activity https://discord.com/developers/docs/topics/gateway#activity-object-activity-structure
type Activity struct {
	Name          string             `json:"name"`
	Type          ActivityType       `json:"type"`
	URL           string             `json:"url,omitempty"`
	CreatedAt     int                `json:"created_at"`
	Timestamps    *ActivityTimestamp `json:"timestamps,omitempty"`
	ApplicationID Snowflake          `json:"application_id,omitempty"`
	Details       string             `json:"details,omitempty"`
	State         string             `json:"state,omitempty"`
	Emoji         *ActivityEmoji     `json:"emoji,omitempty"`
	Party         *ActivityParty     `json:"party,omitempty"`
	Assets        *ActivityAssets    `json:"assets,omitempty"`
	Secrets       *ActivitySecrets   `json:"secrets,omitempty"`
	Instance      bool               `json:"instance,omitempty"`
	Flags         ActivityFlag       `json:"flags,omitempty"`
}

// ---------

type UserFlag uint64

const (
	UserFlagNone            UserFlag = 0
	UserFlagDiscordEmployee UserFlag = 1 << iota
	UserFlagDiscordPartner
	UserFlagHypeSquadEvents
	UserFlagBugHunterLevel1
	_
	_
	UserFlagHouseBravery
	UserFlagHouseBrilliance
	UserFlagHouseBalance
	UserFlagEarlySupporter
	UserFlagTeamUser
	_
	UserFlagSystem
	_
	UserFlagBugHunterLevel2
	_
	UserFlagVerifiedBot
	UserFlagVerifiedBotDeveloper
)

type PremiumType int

func (p PremiumType) String() (t string) {
	switch p {
	case PremiumTypeNone:
		t = "None"
	case PremiumTypeNitroClassic:
		t = "Nitro Classic"
	case PremiumTypeNitro:
		t = "Nitro"
	default:
		t = ""
	}

	return t
}

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

// User the Discord user object which is reused in most other data structures.
type User struct {
	ID            Snowflake     `json:"id,omitempty"`
	Username      string        `json:"username,omitempty"`
	Discriminator Discriminator `json:"discriminator,omitempty"`
	Avatar        string        `json:"avatar"` // data:image/jpeg;base64,BASE64_ENCODED_JPEG_IMAGE_DATA
	Bot           bool          `json:"bot,omitempty"`
	System        bool          `json:"system,omitempty"`
	MFAEnabled    bool          `json:"mfa_enabled,omitempty"`
	Locale        string        `json:"locale,omitempty"`
	Verified      bool          `json:"verified,omitempty"`
	Email         string        `json:"email,omitempty"`
	Flags         UserFlag      `json:"flag,omitempty"`
	PremiumType   PremiumType   `json:"premium_type,omitempty"`
	PublicFlags   UserFlag      `json:"public_flag,omitempty"`
	PartialMember *Member       `json:"member"` // may be populated by Message
}

type UserArr []*User

func (u *User) ToProto() *model.UserData {
	if u == nil {
		return nil
	}
	return &model.UserData{
		Id:            u.ID.ToProto(),
		Username:      u.Username,
		Avatar:        &wrappers.StringValue{Value: u.Avatar},
		Discriminator: uint32(u.Discriminator),
		Bot:           u.Bot,
	}
}

func (a UserArr) ToMentionProto() []*model.MessageData_MessageMentionData {
	var ret []*model.MessageData_MessageMentionData
	for _, el := range a {
		ret = append(ret, &model.MessageData_MessageMentionData{
			Id:            uint64(el.ID),
			Username:      el.Username,
			Avatar:        el.Avatar,
			Discriminator: uint32(el.Discriminator),
		})
	}
	return ret
}

func (a UserArr) ToProto() []*model.UserData {
	var ret []*model.UserData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}

// -------

// UserPresence presence info for a guild member or friend/user in a DM
type UserPresence struct {
	User    *User       `json:"user"`
	Roles   []Snowflake `json:"roles"`
	Game    *Activity   `json:"activity"`
	GuildID Snowflake   `json:"guild_id"`
	Nick    string      `json:"nick"`
	Status  string      `json:"status"`
}

// UserConnection ...
type UserConnection struct {
	ID      string `json:"id"`      // id of the connection account
	Name    string `json:"name"`    // the username of the connection account
	Type    string `json:"type"`    // the service of the connection (twitch, youtube)
	Revoked bool   `json:"revoked"` // whether the connection is revoked
	//Integrations []*IntegrationAccount `json:"integrations"` // an array of partial server integrations
}
