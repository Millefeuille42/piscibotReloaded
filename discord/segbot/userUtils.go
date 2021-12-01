package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// userWriteFile Writes user data to file
func userWriteFile(data UserData, agent discordAgent, id string) error {
	if id == "" {
		id = agent.message.Author.ID
	}
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/users/%s.json", id), dataBytes, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

// userLoadFile Returns user data from file
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

// userCheckHasTarget Check if user has already a target on guild
func userCheckHasTarget(agent discordAgent) error {
	if !userInitialCheck(agent) {
		return os.ErrNotExist
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return err
	}
	if _, isExist := user.GuildTargets[agent.message.GuildID]; isExist {
		sendMessageWithMention("You are already tracking someone on this server!", "", agent)
		return os.ErrExist
	}
	return nil
}

// userIsTrackingCheck Checks if user is tracking someone on guild
func userIsTrackingCheck(agent discordAgent) bool {
	if !userInitialCheck(agent) {
		return false
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return false
	}
	if _, isExist := user.GuildTargets[agent.message.GuildID]; isExist {
		return true
	}
	sendMessageWithMention("You are not tracking anyone on this server!", "", agent)
	return false
}

// userInitialCheck Checks if user is registered
func userInitialCheck(agent discordAgent) bool {
	if !guildInitialCheck(agent) {
		return false
	}
	_, err := os.Stat(fmt.Sprintf("./data/users/%s.json", agent.message.Author.ID))
	if !os.IsNotExist(err) {
		user, err := userLoadFile("", agent)
		if err != nil {
			return false
		}
		if !user.Verified {
			sendMessageWithMention("You are registered,"+
				" but your account is not verified!", "", agent)
			return false
		}
		return true
	}
	sendMessageWithMention("You are not registered, register with "+gPrefix+"start", "", agent)
	return false
}
