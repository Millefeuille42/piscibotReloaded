package target

import "piscibotReloaded/discord/segbot/configFile"

// Target Contains a target Login, and a GuildUsers map
type Target struct {
	configFile.ConfigFile
	Login      string
	GuildUsers map[string]string
}
