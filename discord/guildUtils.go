package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func guildGetChannel(agent discordAgent) string {
	guild, err := guildLoadFile(agent, true)
	if err != nil {
		return agent.message.ChannelID
	}
	return guild.Settings.Channels.Commands
}

func guildLoadFile(agent discordAgent, silent bool) (GuildData, error) {
	data := GuildData{}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/guild/%s.json", agent.message.GuildID))
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

func writeGuildData(agent discordAgent, data GuildData) error {
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

func guildInitialCheck(agent discordAgent) bool {
	_, err := os.Stat(fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID))
	if !os.IsNotExist(err) {
		return true
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "This guild doesn't exists, "+
		"create it with !init")
	return false
}
