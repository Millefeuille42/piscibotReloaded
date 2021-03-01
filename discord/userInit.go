package main

import (
	"fmt"
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

func userUnTrack(agent discordAgent) {
	if !userTrackCheck(agent) {
		return
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}
	delete(user.GuildTargets, agent.message.GuildID)
}

func userInit(agent discordAgent) {
	path := fmt.Sprintf("./data/users/%s.json", agent.message.Author.ID)

	if !guildInitialCheck(agent) {
		return
	}

	exists, err := createFileIfNotExist(path)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	if exists {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are already registered!")
		return
	}

	data := UserData{
		UserID:       agent.message.Author.ID,
		GuildTargets: make(map[string]string),
		Settings: userSettings{
			Leaderboard: "none",
			Success:     "none",
			Started:     "none",
			Location:    "none",
		},
	}

	if userWriteFile(data, agent) != nil {
		return
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "You are now registered")
}
