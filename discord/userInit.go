package main

import (
	"fmt"
)

// targetUntrack Un-tracks target for user on guild
func targetUntrack(agent discordAgent) {
	if !userTrackCheck(agent) {
		return
	}

	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}
	targetName := user.GuildTargets[agent.message.GuildID]
	delete(user.GuildTargets, agent.message.GuildID)
	err = userWriteFile(user, agent)
	if err != nil {
		return
	}

	target, err := targetLoadFile(targetName, agent)
	if err != nil {
		return
	}
	delete(target.GuildUsers, agent.message.GuildID)
	err = targetWriteFile(target, agent)
	if err != nil {
		return
	}
}

// userInit Initializes user
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
