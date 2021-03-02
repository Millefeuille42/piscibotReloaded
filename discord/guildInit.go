package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// createRoles Internal, Creates or get appropriate role
func getOrCreateRole(name string, roles *[]*discordgo.Role, agent discordAgent) (*discordgo.Role, error) {
	var role *discordgo.Role
	skip := false

	for _, rl := range *roles {
		if rl.Name == name {
			skip = true
			role = rl
			break
		}
	}
	if !skip {
		role, err := agent.session.GuildRoleCreate(agent.message.GuildID)
		if err != nil {
			return nil, err
		}
		role, err = agent.session.GuildRoleEdit(
			agent.message.GuildID, role.ID,
			name, role.Color, false, role.Permissions, true,
		)
		if err != nil {
			return nil, err
		}
	}

	if role == nil {
		return nil, os.ErrInvalid
	}
	return role, nil
}

// createRoles Internal, Creates or get appropriate roles, and associate them to data
func createRoles(agent discordAgent, data *GuildData) error {
	names := []string{
		"SegBot - Admin",
		"SegBot - Registered",
		"SegBot - Unregistered",
		"SegBot - Spectator",
	}
	roles, err := agent.session.GuildRoles(agent.message.GuildID) // Set roles list here so not queried every time
	checkRoles := err == nil

	for _, name := range names {
		var role *discordgo.Role

		if checkRoles {
			role, err = getOrCreateRole(name, &roles, agent) // Pass roles as pointer reason is, as above
			if err != nil {
				return err
			}
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

// createData Internal, creates and returns data file
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
	}
	return data
}

// writeData Internal, checks if guild registered, if not registers guild
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
	if guildWriteFile(agent, data) != nil {
		return err
	}
	return nil
}

// guildInit Create guild's data file
func guildInit(agent discordAgent) {
	data := createData(agent)
	if data.GuildID == "" {
		return
	}
	if writeData(agent, data) != nil {
		return
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, "Guild registered successfully!")
}
