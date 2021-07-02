package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func userSetSpectator(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}
	if _, isExist := user.GuildTargets[agent.message.GuildID]; isExist {
		sendMessageWithMention("You can't be a spectator if you are tracking someone!", "", agent)
		return
	}
	_ = discordRoleSetLoad("", "spectator", agent)
	sendMessageWithMention("You are now spectating", "", agent)
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

	link, state := authLinkCreator(agent)
	if link == "" {
		sendMessageWithMention("Could not generate OAuth link", "", agent)
		return
	}

	data := UserData{
		UserID:       agent.message.Author.ID,
		State:        state,
		GuildTargets: make(map[string]string),
		Settings: userSettings{
			Success:  "none",
			Started:  "none",
			Location: "none",
		},
		Verified: false,
	}
	if userWriteFile(data, agent, "") != nil {
		return
	}

	sendMessageWithMention("", "", discordAgent{
		session: agent.session, message: agent.message, channel: agent.message.ChannelID})

	_, err = agent.session.ChannelMessageSendEmbed(agent.message.ChannelID, &discordgo.MessageEmbed{
		URL:   link,
		Type:  "link",
		Title: "Verification Link",
		Description: "You are now registered, validate your profile with the link provided.\n" +
			"You will not be able to perform actions until you validate your profile through 42",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8d/42_Logo.svg/1200px-42_Logo.svg.png",
		},
	})
}
