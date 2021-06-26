package discordUser

import (
	"fmt"
	"os"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/guild"
)

// CheckHasTarget Check if user has already a target on guild
func CheckHasTarget(agent discord.Agent) error {
	if !InitialCheck(agent) {
		return os.ErrNotExist
	}
	user, err := Load("", agent)
	if err != nil {
		return err
	}
	if _, isExist := user.GuildTargets[agent.Message.GuildID]; isExist {
		discord.SendMessageWithMention("You are already tracking someone on this server!", "", agent)
		return os.ErrExist
	}
	return nil
}

// IsTrackingCheck Checks if user is tracking someone on guild
func IsTrackingCheck(agent discord.Agent) bool {
	if !InitialCheck(agent) {
		return false
	}
	user, err := Load("", agent)
	if err != nil {
		return false
	}
	if _, isExist := user.GuildTargets[agent.Message.GuildID]; isExist {
		return true
	}
	discord.SendMessageWithMention("You are not tracking anyone on this server!", "", agent)
	return false
}

// InitialCheck Checks if user is registered
func InitialCheck(agent discord.Agent) bool {
	if !guild.InitialCheck(agent) {
		return false
	}
	_, err := os.Stat(fmt.Sprintf("./data/users/%s.json", agent.Message.Author.ID))
	if !os.IsNotExist(err) {
		user, err := Load("", agent)
		if err != nil {
			return false
		}
		if !user.Verified {
			discord.SendMessageWithMention("You are registered,"+
				" but your account is not verified!", "", agent)
			return false
		}
		return true
	}
	discord.SendMessageWithMention("You are not registered, register with !start", "", agent)
	return false
}
