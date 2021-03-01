package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

func userCheckTarget(session *discordgo.Session, message *discordgo.MessageCreate) error {
	if !userInitialCheck(session, message) {
		return os.ErrNotExist
	}
	user, err := userLoadFile("", session, message)
	if err != nil {
		return err
	}
	if _, isExist := user.GuildTargets[message.GuildID]; isExist {
		_, _ = session.ChannelMessageSend(message.ChannelID, "You are already tracking someone on this"+
			" server!")
		return os.ErrExist
	}
	return nil
}

func userWriteFile(data UserData, session *discordgo.Session, message *discordgo.MessageCreate) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/users/%s.json", message.Author.ID), dataBytes, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	return nil
}

func userLoadFile(id string, session *discordgo.Session, message *discordgo.MessageCreate) (UserData, error) {
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

	return user, nil
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
