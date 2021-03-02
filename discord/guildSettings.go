package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// adminSendSettings Send guild's settings, admin rights are not required for this
func adminSendSettings(agent discordAgent) {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(agent) {
		return
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}

	mess := fmt.Sprintf(
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
			mess = fmt.Sprintf("%s@%s\n```", mess, getUser(agent.session, admin))
			break
		}
		mess = fmt.Sprintf("%s@%s, ", mess, getUser(agent.session, admin))
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, mess)
}

// adminSet Add provided admins to the guild
func adminSet(agent discordAgent) {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(agent) {
		return
	}

	args := strings.Split(agent.message.Content, "-")
	if len(args) <= 1 {
		return
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}

	if !Find(settings.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
		return
	}

	for _, user := range args[1:] {
		if !strings.Contains(user, "!") {
			continue
		}
		user = strings.TrimSpace(user)
		user = user[3 : len(user)-1]
		if !Find(settings.Admins, user) {
			settings.Admins = append(settings.Admins, user)
		}
	}
	_ = guildWriteFile(agent, settings)
}

// adminSetChan Set provided channels to the originating channel
func adminSetChan(agent discordAgent) {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(agent) {
		return
	}

	args := strings.Split(agent.message.Content, "-")
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}

	if !Find(settings.Admins, agent.message.Author.ID) {
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
	_ = guildWriteFile(agent, settings)
}
