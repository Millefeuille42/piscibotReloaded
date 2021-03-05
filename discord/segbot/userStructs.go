package main

// userSettings Internal, contains user's ping settings per channel
type userSettings struct {
	Success  string
	Started  string
	Location string
}

// UserData Contains UserID a GuildTargets map and the Settings
type UserData struct {
	UserID       string
	GuildTargets map[string]string
	Settings     userSettings
	Verified     bool
}
