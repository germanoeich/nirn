package discord

// Role https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID          Snowflake     `json:"id"`
	Name        string        `json:"name"`
	Color       uint          `json:"color"`
	Hoist       bool          `json:"hoist"`
	Position    int           `json:"position"` // can be -1
	Permissions PermissionBit `json:"permissions"`
	Managed     bool          `json:"managed"`
	Mentionable bool          `json:"mentionable"`
	guildID     Snowflake
}
