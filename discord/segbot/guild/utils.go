package guild

import (
	"fmt"
	"os"
	"piscibotReloaded/discord/segbot/discord"
)

// GetChannel Returns guild's command channel
func GetChannel(agent discord.Agent) string {
	guild, err := Load(agent, true, "")
	if err != nil {
		return agent.Message.ChannelID
	}
	return guild.Settings.Channels.Commands
}

// InitialCheck Required before guild related actions, checks if guild exists
func InitialCheck(agent discord.Agent) bool {
	_, err := os.Stat(fmt.Sprintf("./data/guilds/%s.json", agent.Message.GuildID))
	if !os.IsNotExist(err) {
		return true
	}
	discord.SendMessageWithMention("This guild doesn't exist, create it with !init", "", agent)
	return false
}
