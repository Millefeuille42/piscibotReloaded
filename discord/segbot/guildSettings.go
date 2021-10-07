package main

import (
	"fmt"
	"piscibotReloaded/discord/segbot/utils"
	"strings"
)

// adminSendSettings Send a guild settings, admin rights are not required for this
func adminSendSettings(agent discordAgent) {
	if !guildInitialCheck(agent) {
		return
	}
	settings, err := guildLoadFile(agent, false, "")
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
		getChannelName(agent.session, settings.Settings.Channels.Commands),
		getChannelName(agent.session, settings.Settings.Channels.Leaderboard),
		getChannelName(agent.session, settings.Settings.Channels.Success),
		getChannelName(agent.session, settings.Settings.Channels.Started),
		getChannelName(agent.session, settings.Settings.Channels.Location),
	)

	for i, admin := range settings.Admins {
		if i == len(settings.Admins)-1 {
			message = fmt.Sprintf("%s@%s\n```", message, getUser(agent.session, admin))
			break
		}
		message = fmt.Sprintf("%s@%s, ", message, getUser(agent.session, admin))
	}
	sendMessageWithMention(message, "", agent)
}

// adminSet Add provided admins to the guild
func adminSet(agent discordAgent) {
	if !guildInitialCheck(agent) {
		return
	}
	args := strings.Split(agent.message.Content, " ")
	if len(args) <= 1 {
		return
	}
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}

	if !utils.Find(data.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
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
			discordRoleSet(data, user, "admin", agent)
		}
	}
	if guildWriteFile(agent, data) == nil {
		sendMessageWithMention("Successfully added user(s) as admin", "", agent)
	}
}

// adminSetChan Set provided channels to the originating channel
func adminSetChan(agent discordAgent) {

	if !guildInitialCheck(agent) {
		return
	}

	args := strings.Split(agent.message.Content, " ")
	settings, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}

	if !utils.Find(settings.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
		return
	}

	for _, channel := range args {
		switch channel {
		case "command":
			settings.Settings.Channels.Commands = agent.message.ChannelID
		case "leaderboard":
			settings.Settings.Channels.Leaderboard = agent.message.ChannelID
		case "success":
			settings.Settings.Channels.Success = agent.message.ChannelID
		case "started":
			settings.Settings.Channels.Started = agent.message.ChannelID
		case "location":
			settings.Settings.Channels.Location = agent.message.ChannelID
		}
	}
	if guildWriteFile(agent, settings) == nil {
		sendMessageWithMention("Successfully changed channels", "", agent)
	}
}
