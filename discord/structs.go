package main

type guildSettingsChannels struct {
	Commands    string
	Leaderboard string
	Success     string
	Started     string
	Location    string
}

type guildSettingsRoles struct {
	Admin        string
	Registered   string
	Unregistered string
	Spectator    string
}

type guildSettings struct {
	Channels guildSettingsChannels
	Roles    guildSettingsRoles
}

// Guild Data
type GuildData struct {
	GuildID  string
	Admins   []string
	Settings guildSettings
}
