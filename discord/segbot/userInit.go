package main

import (
	"fmt"
)

// targetUntrack Un-tracks target for user on guild
func targetUntrack(agent discordAgent) {
	if !userIsTrackingCheck(agent) {
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
		sendMessageWithMention("You are already registered!", "", agent)
		return
	}

	link := authLinkCreator(agent.message.Author.ID)
	if link == "" {
		sendMessageWithMention("Could not generate oauth link", "", agent)
		return
	}

	data := UserData{
		UserID:       agent.message.Author.ID,
		GuildTargets: make(map[string]string),
		Settings: userSettings{
			Success:  "none",
			Started:  "none",
			Location: "none",
		},
		Verified: false,
	}

	if userWriteFile(data, agent) != nil {
		return
	}
	sendMessageWithMention("You are now registered,"+
		" validate your profile with the link below", "", agent)
	_, _ = agent.session.ChannelMessageSend(agent.channel, link)
}
