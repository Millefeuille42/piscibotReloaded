package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type userSettings struct {
	Leaderboard string
	Success     string
	Started     string
	Location    string
}

type UserData struct {
	UserID       string
	GuildTargets map[string]string
	Settings     userSettings
}

func userInit(session *discordgo.Session, message *discordgo.MessageCreate) {
	path := fmt.Sprintf("./data/users/%s.json", message.Author.ID)

	if !guildInitialCheck(session, message) {
		return
	}

	exists, err := createFileIfNotExist(path)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	if exists {
		_, _ = session.ChannelMessageSend(message.ChannelID, "You are already registered!")
		return
	}

	data := UserData{
		UserID:       message.Author.ID,
		GuildTargets: make(map[string]string),
		Settings: userSettings{
			Leaderboard: "none",
			Success:     "none",
			Started:     "none",
			Location:    "none",
		},
	}

	if userWriteFile(data, session, message) != nil {
		return
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "You are now registered")
}
