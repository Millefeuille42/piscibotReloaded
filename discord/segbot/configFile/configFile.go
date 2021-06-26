package configFile

import "piscibotReloaded/discord/segbot/discord"

type ConfigFile interface {
	Load(file ConfigFile, agent discord.Agent)
	Write(file ConfigFile, agent discord.Agent)
}
