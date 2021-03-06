package main

// guildSettingsChannels Internal, Contains Channels associations
type guildSettingsChannels struct {
	Commands    string
	Leaderboard string
	Success     string
	Started     string
	Location    string
}

// guildSettingsRoles Internal, Contains Roles associations
type guildSettingsRoles struct {
	Admin      string
	Registered string
	Spectator  string
}

// guildSettings Internal, Contains structs of the Channels and Roles
type guildSettings struct {
	Channels guildSettingsChannels
	Roles    guildSettingsRoles
}

// GuildData Contains the GuildID, a list of Admins and the guild Settings
type GuildData struct {
	GuildID  string
	Admins   []string
	Settings guildSettings
}
