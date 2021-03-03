package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

func createRoles(session *discordgo.Session, message *discordgo.MessageCreate, data *GuildData) error {
	names := []string{
		"SegBot - Admin",
		"SegBot - Registered",
		"SegBot - Unregistered",
		"SegBot - Spectator",
	}

	for _, name := range names {
		role, err := session.GuildRoleCreate(message.GuildID)
		if err != nil {
			return err
		}
		role, err = session.GuildRoleEdit(
			message.GuildID, role.ID,
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

func createData(session *discordgo.Session, message *discordgo.MessageCreate) GuildData {

	data := GuildData{
		GuildID: message.GuildID,
		Admins:  append(make([]string, 0), message.Author.ID),
		Settings: guildSettings{
			Channels: guildSettingsChannels{
				Commands:    message.ChannelID,
				Leaderboard: message.ChannelID,
				Success:     message.ChannelID,
				Started:     message.ChannelID,
				Location:    message.ChannelID,
			},
			Roles: guildSettingsRoles{
				Admin:        "none",
				Registered:   "none",
				Unregistered: "none",
				Spectator:    "none",
			},
		},
	}
	if createRoles(session, message, &data) != nil {
		_, _ = session.ChannelMessageSend(message.ChannelID,
			"Failed to create roles, you'll have to create and configure the missing ones")
		return GuildData{GuildID: ""}
	}
	return data
}

func writeData(session *discordgo.Session, message *discordgo.MessageCreate, data GuildData) error {
	path := fmt.Sprintf("./data/guilds/%s.json", message.GuildID)

	exists, err := createFileIfNotExist(path)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	if exists {
		_, _ = session.ChannelMessageSend(message.ChannelID, "This Guild is already registered!")
		return os.ErrExist
	}
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	err = ioutil.WriteFile(path, jsonGuild, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	return nil
}

// Create guild's data file
func initGuild(session *discordgo.Session, message *discordgo.MessageCreate) {
	data := createData(session, message)
	if data.GuildID == "" {
		return
	}
	if writeData(session, message, data) != nil {
		return
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "Guild registered successfully!")
}
