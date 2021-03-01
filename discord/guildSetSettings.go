package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func setAdmin(agent discordAgent) {
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
	_ = writeGuildData(agent, settings)
}

func setChan(agent discordAgent) {
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
	_ = writeGuildData(agent, settings)
}
