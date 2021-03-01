package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

type userSettings struct {
	Leaderboard string
	Success     string
	Started     string
	Location    string
}

type UserData struct {
	UserID   string
	settings userSettings
}

func initUser(session *discordgo.Session, message *discordgo.MessageCreate) {
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
		UserID: message.Author.ID,
		settings: userSettings{
			Leaderboard: "none",
			Success:     "none",
			Started:     "none",
			Location:    "none",
		},
	}

	jsonUser, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	err = ioutil.WriteFile(path, jsonUser, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "You are now registered")
}
