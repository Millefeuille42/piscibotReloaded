package target

import (
	"fmt"
	"os"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"strings"
)

// Track Registers target for user and guild
func Track(agent discord.Agent) {
	settings := Target{}
	args := strings.Split(agent.Message.Content, " ")
	if discordUser.CheckHasTarget(agent) != nil {
		return
	}
	if len(args) < 2 {
		discord.SendMessageWithMention("I need more arguments!", "", agent)
		return
	}

	path := fmt.Sprintf("./data/targets/%s.json", args[1])
	err := loadOrCreate(path, args[1], &settings, agent.Message)
	if err != nil {
		if err == os.ErrExist {
			discord.SendMessageWithMention("Someone is already tracking this person"+
				" on this server!", "", agent)
			return
		}
		discord.LogErrorToChan(agent, err)
		return
	}

	if makeApiReq(path, args[1], agent) != nil {
		return
	}

	user, err := discordUser.Load("", agent)
	if err != nil {
		return
	}
	user.GuildTargets[agent.Message.GuildID] = settings.Login

	if Write(settings, agent) == nil && discordUser.Write(user, agent, "") == nil {
		discord.SendMessageWithMention("You are now tracking "+args[1], "", agent)
		_ = discord.RoleSetLoad("", "registered", agent)
	}
}

// Untrack Un-tracks target for user on guild
func Untrack(agent discord.Agent) {
	if !discordUser.IsTrackingCheck(agent) {
		return
	}

	user, err := discordUser.Load("", agent)
	if err != nil {
		return
	}
	targetName := user.GuildTargets[agent.Message.GuildID]
	delete(user.GuildTargets, agent.Message.GuildID)
	err = discordUser.Write(user, agent, "")
	if err != nil {
		return
	}

	target, err := Load(targetName, agent)
	if err != nil {
		return
	}
	delete(target.GuildUsers, agent.Message.GuildID)
	err = Write(target, agent)
	if err != nil {
		return
	}
	_ = discord.RoleSetLoad("", "spectator", agent)
	discord.SendMessageWithMention("You are not tracking someone on this server anymore!", "", agent)
}
