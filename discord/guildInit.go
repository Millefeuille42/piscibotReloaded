package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func createRoles(agent discordAgent, data *GuildData) error {
	names := []string{
		"SegBot - Admin",
		"SegBot - Registered",
		"SegBot - Unregistered",
		"SegBot - Spectator",
	}

	for _, name := range names {
		role, err := agent.session.GuildRoleCreate(agent.message.GuildID)
		if err != nil {
			return err
		}
		role, err = agent.session.GuildRoleEdit(
			agent.message.GuildID, role.ID,
			name, role.Color, false, role.Permissions, true,
		)
		if err != nil {
			return err
		}
		switch name {
		case "SegBot - Admin":
			data.Settings.Roles.Admin = role.ID
		case "SegBot - Registered":
			data.Settings.Roles.Registered = role.ID
		case "SegBot - Unregistered":
			data.Settings.Roles.Unregistered = role.ID
		case "SegBot - Spectator":
			data.Settings.Roles.Spectator = role.ID
		}
	}
	return nil
}

func createData(agent discordAgent) GuildData {

	data := GuildData{
		GuildID: agent.message.GuildID,
		Admins:  append(make([]string, 0), agent.message.Author.ID),
		Settings: guildSettings{
			Channels: guildSettingsChannels{
				Commands:    agent.message.ChannelID,
				Leaderboard: agent.message.ChannelID,
				Success:     agent.message.ChannelID,
				Started:     agent.message.ChannelID,
				Location:    agent.message.ChannelID,
			},
			Roles: guildSettingsRoles{
				Admin:        "none",
				Registered:   "none",
				Unregistered: "none",
				Spectator:    "none",
			},
		},
	}
	if createRoles(agent, &data) != nil {
		_, _ = agent.session.ChannelMessageSend(agent.channel,
			"Failed to create roles, you'll have to create and configure the missing ones")
		return GuildData{GuildID: ""}
	}
	return data
}

func writeData(agent discordAgent, data GuildData) error {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID)

	exists, err := createFileIfNotExist(path)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	if exists {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "This Guild is already registered!")
		return os.ErrExist
	}
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(path, jsonGuild, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

// Create guild's data file
func initGuild(agent discordAgent) {
	data := createData(agent)
	if data.GuildID == "" {
		return
	}
	if writeData(agent, data) != nil {
		return
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "Guild registered successfully!")
}
