package discord

import "github.com/bwmarrin/discordgo"

// Agent Contains discord's session and message structs, and the guild's command channel
type Agent struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Channel string
}
