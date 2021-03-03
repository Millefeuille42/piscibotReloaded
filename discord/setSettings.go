package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"strings"
)

func setAdmin(session *discordgo.Session, message *discordgo.MessageCreate) {
	path := fmt.Sprintf("./data/guilds/%s.json", message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(session, message) {
		return
	}

	args := strings.Split(message.Content, "-")
	if len(args) <= 1 {
		return
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}

	if !Find(settings.Admins, message.Author.ID) {
		_, _ = session.ChannelMessageSend(message.ChannelID, "You are not an admin")
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
	_ = writeGuildSettings(session, message, settings)
}

func setChan(session *discordgo.Session, message *discordgo.MessageCreate) {
	path := fmt.Sprintf("./data/guilds/%s.json", message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(session, message) {
		return
	}

	args := strings.Split(message.Content, "-")
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}

	if !Find(settings.Admins, message.Author.ID) {
		_, _ = session.ChannelMessageSend(message.ChannelID, "You are not an admin")
		return
	}

	for _, channel := range args {
		switch channel {
		case "command":
			settings.Settings.Channels.Commands = message.ChannelID
		case "leaderboard":
			settings.Settings.Channels.Leaderboard = message.ChannelID
		case "success":
			settings.Settings.Channels.Success = message.ChannelID
		case "started":
			settings.Settings.Channels.Started = message.ChannelID
		case "location":
			settings.Settings.Channels.Location = message.ChannelID
		}
	}
	_ = writeGuildSettings(session, message, settings)
}
