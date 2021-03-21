package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// discordAgent Contains discord's session and message structs, and the guild's command channel
type discordAgent struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel string
}

// adminRouter Router for admin
func adminRouter(agent discordAgent) {
	switch {
	case strings.HasPrefix(agent.message.Content, "!init"):
		guildInit(agent)
	case strings.HasPrefix(agent.message.Content, "!chan"):
		adminSetChan(agent)
	case strings.HasPrefix(agent.message.Content, "!admin"):
		adminSet(agent)
	case agent.message.Content == "!params":
		adminSendSettings(agent)
	}
}

func commandsRouter(agent discordAgent) bool {
	switch {
	//	case strings.HasPrefix(agent.message.Content, "!profile"):
	//		sendTargetProfile(agent)
	//		return true
	//	case strings.HasPrefix(agent.message.Content, "!list students"):
	//		sendStudentsList(agent)
	//		return true
	//	case strings.HasPrefix(agent.message.Content, "!list tracked"):
	//		sendTrackedList(agent)
	//		return true
	case strings.HasPrefix(agent.message.Content, "!leaderboard"):
		sendLeaderboard(agent)
		return true
	}
	return false
}

// userRouter Router for user commands, returns true if a command was found
func userRouter(agent discordAgent) bool {
	switch {
	case strings.HasPrefix(agent.message.Content, "!start"):
		userInit(agent)
		return true
	case strings.HasPrefix(agent.message.Content, "!track"):
		targetTrack(agent)
		return true
	case strings.HasPrefix(agent.message.Content, "!ping"):
		userSetPings(agent)
		return true
	case agent.message.Content == "!settings":
		userSendSettings(agent)
		return true
	case agent.message.Content == "!untrack":
		targetUntrack(agent)
		return true
	case agent.message.Content == "!spectate":
		userSetSpectator(agent)
		return true
	}
	return false
}

// messageHandler Discord bot message handler
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
	if !commandsRouter(agent) && !userRouter(agent) {
		adminRouter(agent)
	}
}
