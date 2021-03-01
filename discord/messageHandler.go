package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type discordAgent struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel string
}

// Router for admin
func adminRouter(agent discordAgent) {
	switch {
	case strings.HasPrefix(agent.message.Content, "!init"):
		initGuild(agent)
	case strings.HasPrefix(agent.message.Content, "!chan"):
		setChan(agent)
	case strings.HasPrefix(agent.message.Content, "!admin"):
		setAdmin(agent)
	case agent.message.Content == "!params":
		sendAdminSettings(agent)
	}
}

// Router for user commands, returns true if a command was found
func userRouter(agent discordAgent) bool {
	switch {
	case strings.HasPrefix(agent.message.Content, "!start"):
		userInit(agent)
		return true
	case strings.HasPrefix(agent.message.Content, "!track"):
		registerTarget(agent)
		return true
	case strings.HasPrefix(agent.message.Content, "!ping"):
		// editPings
		fmt.Println("PING")
		return true
	case agent.message.Content == "!settings":
		// getUserSettings
		fmt.Println("USER SETTINGS")
		return true
	case agent.message.Content == "!untrack":
		userUnTrack(agent)
		return true
	}
	return false
}

// Discord bot message handler
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	botID, _ := session.User("@me")
	agent := discordAgent{
		session: session,
		message: message,
	}
	agent.channel = guildGetChannel(agent)

	if message.Author.ID == botID.ID || !strings.HasPrefix(message.Content, "!") {
		return
	}
	if !userRouter(agent) {
		adminRouter(agent)
	}
}
