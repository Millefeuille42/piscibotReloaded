package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func userInitialCheck(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	_, err := os.Stat(fmt.Sprintf("./data/users/%s.json", message.Author.ID))
	if !os.IsNotExist(err) {
		return true
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "This user doesn't exists, "+
		"create it with !start")
	return false
}
