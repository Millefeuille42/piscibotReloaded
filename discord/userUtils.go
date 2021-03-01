package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func userCheckTarget(agent discordAgent) error {
	if !userInitialCheck(agent) {
		return os.ErrNotExist
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return err
	}
	if _, isExist := user.GuildTargets[agent.message.GuildID]; isExist {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are already tracking someone on this"+
			" server!")
		return os.ErrExist
	}
	return nil
}

func userWriteFile(data UserData, agent discordAgent) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/users/%s.json", agent.message.Author.ID), dataBytes, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

func userLoadFile(id string, agent discordAgent) (UserData, error) {
	user := UserData{}
	if id == "" {
		id = agent.message.Author.ID
	}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s.json", id))
	if err != nil {
		logErrorToChan(agent, err)
		return UserData{}, err
	}

	err = json.Unmarshal(fileData, &user)
	if err != nil {
		logErrorToChan(agent, err)
		return UserData{}, err
	}

	return user, nil
}

func userTrackCheck(agent discordAgent) bool {
	user, err := userLoadFile("", agent)
	if err != nil {
		return false
	}
	if _, isExist := user.GuildTargets[agent.message.GuildID]; isExist {
		return true
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not tracking anyone on this server!")
	return false
}

func userInitialCheck(agent discordAgent) bool {
	_, err := os.Stat(fmt.Sprintf("./data/users/%s.json", agent.message.Author.ID))
	if !os.IsNotExist(err) {
		return true
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "This user doesn't exists, "+
		"create it with !start")
	return false
}
