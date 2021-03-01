package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

func loadUserFile(id string, session *discordgo.Session, message *discordgo.MessageCreate) (UserData, error) {
	user := UserData{}
	if id == "" {
		id = message.Author.ID
	}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s.json", id))
	if err != nil {
		logErrorToChan(session, message, err)
		return UserData{}, err
	}

	err = json.Unmarshal(fileData, &user)
	if err != nil {
		logErrorToChan(session, message, err)
		return UserData{}, err
	}

	return user, err
}

func userInitialCheck(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	_, err := os.Stat(fmt.Sprintf("./data/users/%s.json", message.Author.ID))
	if !os.IsNotExist(err) {
		return true
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "This user doesn't exists, "+
		"create it with !start")
	return false
}
