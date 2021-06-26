package guild

import "piscibotReloaded/discord/segbot/configFile"

// settingsChannels Internal, Contains Channels associations
type settingsChannels struct {
	Commands    string
	Leaderboard string
	Success     string
	Started     string
	Location    string
}

// settingsRoles Internal, Contains Roles associations
type settingsRoles struct {
	Admin      string
	Registered string
	Spectator  string
}

// settings Internal, Contains structs of the Channels and Roles
type settings struct {
	Channels settingsChannels
	Roles    settingsRoles
}

// Guild Contains the GuildID, a list of Admins and the guild Settings
type Guild struct {
	configFile.ConfigFile
	GuildID  string
	Admins   []string
	Settings settings
}
