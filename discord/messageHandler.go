package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Router for admin
func adminRouter(session *discordgo.Session, message *discordgo.MessageCreate) {
	switch {
	case strings.HasPrefix(message.Content, "!init"):
		initGuild(session, message)
	case strings.HasPrefix(message.Content, "!chan"):
		setChan(session, message)
	case strings.HasPrefix(message.Content, "!admin"):
		setAdmin(session, message)
	case message.Content == "!params":
		sendAdminSettings(session, message)
	}
}

// Router for user commands, returns true if a command was found
func userRouter(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	switch {
	case strings.HasPrefix(message.Content, "!start"):
		initUser(session, message)
		return true
	case strings.HasPrefix(message.Content, "!register"):
		registerTarget(session, message)
		return true
	case strings.HasPrefix(message.Content, "!ping"):
		// editPings
		fmt.Println("PING")
		return true
	case message.Content == "!settings":
		// getUserSettings
		fmt.Println("USER SETTINGS")
		return true
	}
	return false
}

// Discord bot message handler
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	botID, _ := session.User("@me")
	if message.Author.ID == botID.ID || !strings.HasPrefix(message.Content, "!") {
		return
	}
	if !userRouter(session, message) {
		adminRouter(session, message)
	}
}
