package discordUser

import "piscibotReloaded/discord/segbot/configFile"

// settings Internal, contains user's ping settings per channel
type settings struct {
	Success  string
	Started  string
	Location string
}

// User Contains UserID a GuildTargets map and the Settings
type User struct {
	configFile.ConfigFile
	UserID       string
	State        string
	GuildTargets map[string]string
	Settings     settings
	Verified     bool
}
