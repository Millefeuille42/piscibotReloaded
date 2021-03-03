package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// guildGetChannel Returns guild's command channel
func guildGetChannel(agent discordAgent) string {
	guild, err := guildLoadFile(agent, true, "")
	if err != nil {
		return agent.message.ChannelID
	}
	return guild.Settings.Channels.Commands
}

// guildLoadFile Returns guild data from file
func guildLoadFile(agent discordAgent, silent bool, id string) (GuildData, error) {
	data := GuildData{}

	if id == "" {
		id = agent.message.GuildID
	}
	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/guilds/%s.json", id))
	if err != nil {
		if !silent {
			logErrorToChan(agent, err)
		}
		return GuildData{}, err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		if !silent {
			logErrorToChan(agent, err)
		}
		return GuildData{}, err
	}

	return data, nil
}

// guildWriteFile Writes guild data to file
func guildWriteFile(agent discordAgent, data GuildData) error {
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/guilds/%s.json", data.GuildID), jsonGuild, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

// guildInitialCheck Required before guild related actions, checks if guild exists
func guildInitialCheck(agent discordAgent) bool {
	_, err := os.Stat(fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID))
	if !os.IsNotExist(err) {
		return true
	}
	sendMessageWithMention("This guild doesn't exist, create it with !init", "", agent)
	return false
}
