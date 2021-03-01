package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

func writeGuildSettings(session *discordgo.Session, message *discordgo.MessageCreate, data GuildData) error {
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/guilds/%s.json", data.GuildID), jsonGuild, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	return nil
}

func guildInitialCheck(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	_, err := os.Stat(fmt.Sprintf("./data/guilds/%s.json", message.GuildID))
	if !os.IsNotExist(err) {
		return true
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "This guild doesn't exists, "+
		"create it with !init")
	return false
}
