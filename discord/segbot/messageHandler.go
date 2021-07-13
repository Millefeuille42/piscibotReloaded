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
	case strings.HasPrefix(agent.message.Content, gPrefix+"init"):
		guildInit(agent)
	case strings.HasPrefix(agent.message.Content, gPrefix+"chan"):
		adminSetChan(agent)
	case strings.HasPrefix(agent.message.Content, gPrefix+"admin"):
		adminSet(agent)
	case agent.message.Content == gPrefix+"params":
		adminSendSettings(agent)
	case agent.message.Content == gPrefix+"purge":
		adminPurge(agent)
	case agent.message.Content == gPrefix+"lock":
		adminLock(agent)
	case agent.message.Content == gPrefix+"unlock":
		adminUnlock(agent)
	}
}

func commandsRouter(agent discordAgent) bool {
	switch {
	case strings.HasPrefix(agent.message.Content, gPrefix+"profile"):
		sendTargetProfile(agent)
		return true
	case strings.HasPrefix(agent.message.Content, gPrefix+"list"):
		switch {
		case strings.HasPrefix(agent.message.Content[6:], "students"):
			sendStudentsList(agent)
			return true
		case strings.HasPrefix(agent.message.Content[6:], "tracked"):
			sendTrackedList(agent)
			return true
		case strings.HasPrefix(agent.message.Content[6:], "projects"):
			sendProjectList(agent)
			return true
		case strings.HasPrefix(agent.message.Content[6:], "location"):
			sendLocationList(agent)
		}
	case strings.HasPrefix(agent.message.Content, gPrefix+"leaderboard"):
		sendLeaderboard(agent)
		return true
	case strings.HasPrefix(agent.message.Content, gPrefix+"project"):
		sendProject(agent)
		return true
	}
	return false
}

// userRouter Router for user commands, returns true if a command was found
func userRouter(agent discordAgent) bool {
	switch {
	case strings.HasPrefix(agent.message.Content, gPrefix+"start"):
		userInit(agent)
		return true
	case strings.HasPrefix(agent.message.Content, gPrefix+"track"):
		targetTrack(agent)
		return true
	case strings.HasPrefix(agent.message.Content, gPrefix+"ping"):
		userSetPings(agent)
		return true
	case agent.message.Content == gPrefix+"settings":
		userSendSettings(agent)
		return true
	case agent.message.Content == gPrefix+"untrack":
		targetUntrack(agent)
		return true
	case agent.message.Content == gPrefix+"spectate":
		userSetSpectator(agent)
		return true
	case agent.message.Content == gPrefix+"help":
		sendHelp(agent)
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

	if message.Author.ID == botID.ID || !strings.HasPrefix(message.Content, gPrefix) {
		return
	}
	if !commandsRouter(agent) && !userRouter(agent) {
		adminRouter(agent)
	}
}
