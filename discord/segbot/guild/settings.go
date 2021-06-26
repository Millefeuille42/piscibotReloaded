package guild

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/utils"
	"strings"
)

// AdminSendSettings Send a guild settings, admin rights are not required for this
func AdminSendSettings(agent discord.Agent) {
	if !InitialCheck(agent) {
		return
	}
	settings, err := Load(agent, false, "")
	if err != nil {
		return
	}

	message := fmt.Sprintf(
		"```\n"+
			"Channels:\n"+
			"    Commands:     #%s\n"+
			"    Leaderboards: #%s\n"+
			"    Success:      #%s\n"+
			"    Started:      #%s\n"+
			"    Location:     #%s\n\n"+
			"Admins:           ",
		discord.GetChannelName(agent.Session, settings.Settings.Channels.Commands),
		discord.GetChannelName(agent.Session, settings.Settings.Channels.Leaderboard),
		discord.GetChannelName(agent.Session, settings.Settings.Channels.Success),
		discord.GetChannelName(agent.Session, settings.Settings.Channels.Started),
		discord.GetChannelName(agent.Session, settings.Settings.Channels.Location),
	)

	for i, admin := range settings.Admins {
		if i == len(settings.Admins)-1 {
			message = fmt.Sprintf("%s@%s\n```", message, discord.GetUser(agent.Session, admin))
			break
		}
		message = fmt.Sprintf("%s@%s, ", message, discord.GetUser(agent.Session, admin))
	}
	discord.SendMessageWithMention(message, "", agent)
}

// AdminSet Add provided admins to the guild
func AdminSet(agent discord.Agent) {
	if !InitialCheck(agent) {
		return
	}
	args := strings.Split(agent.Message.Content, " ")
	if len(args) <= 1 {
		return
	}
	data, err := Load(agent, false, "")
	if err != nil {
		return
	}

	if !utils.Find(data.Admins, agent.Message.Author.ID) {
		_, _ = agent.Session.ChannelMessageSend(agent.Channel, "You are not an admin")
		return
	}

	for _, user := range args[1:] {
		if !strings.Contains(user, "!") {
			continue
		}
		user = strings.TrimSpace(user)
		user = user[3 : len(user)-1]
		if !utils.Find(data.Admins, user) {
			data.Admins = append(data.Admins, user)
			discord.RoleSet(data, user, "admin", agent)
		}
	}
	if Write(agent, data) == nil {
		discord.SendMessageWithMention("Successfully added user(s) as admin", "", agent)
	}
}

// AdminSetChan Set provided channels to the originating channel
func AdminSetChan(agent discord.Agent) {

	if !InitialCheck(agent) {
		return
	}

	args := strings.Split(agent.Message.Content, " ")
	settings, err := Load(agent, false, "")
	if err != nil {
		return
	}

	if !utils.Find(settings.Admins, agent.Message.Author.ID) {
		_, _ = agent.Session.ChannelMessageSend(agent.Channel, "You are not an admin")
		return
	}

	for _, channel := range args {
		switch channel {
		case "command":
			settings.Settings.Channels.Commands = agent.Message.ChannelID
		case "leaderboard":
			settings.Settings.Channels.Leaderboard = agent.Message.ChannelID
		case "success":
			settings.Settings.Channels.Success = agent.Message.ChannelID
		case "started":
			settings.Settings.Channels.Started = agent.Message.ChannelID
		case "location":
			settings.Settings.Channels.Location = agent.Message.ChannelID
		}
	}
	if Write(agent, settings) == nil {
		discord.SendMessageWithMention("Successfully changed channels", "", agent)
	}
}
