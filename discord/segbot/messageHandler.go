package main

import (
	"github.com/bwmarrin/discordgo"
	"piscibotReloaded/discord/segbot/utils"
	"strings"
)

// discordAgent Contains discord's session and message structs, and the guild's command channel
type discordAgent struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel string

	content string
	command string
	args    []string
}
type commandHandler func(agent discordAgent)

func commandRouter(agent discordAgent) {
	agent.content = strings.Replace(agent.message.Content, gPrefix, "", 1)
	splitBuffer := utils.CleanSplit(agent.content, ' ')
	if len(splitBuffer) < 1 {
		return
	}
	agent.command = splitBuffer[0]
	agent.args = splitBuffer
	if fc, ok := commandMap[agent.command]; ok {
		fc(agent)
	}
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
	commandRouter(agent)
}
