package guild

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/utils"
)

// createRoles Internal, Creates or get appropriate role
func getOrCreateRole(name string, roles *[]*discordgo.Role, agent discord.Agent) (*discordgo.Role, error) {
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
		role, err := agent.Session.GuildRoleCreate(agent.Message.GuildID)
		if err != nil {
			return nil, err
		}
		role, err = agent.Session.GuildRoleEdit(
			agent.Message.GuildID, role.ID,
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
func createRoles(agent discord.Agent, data *Guild) error {
	names := []string{
		"SegBot - Admin",
		"SegBot - Registered",
		"SegBot - Unregistered",
		"SegBot - Spectator",
	}
	roles, err := agent.Session.GuildRoles(agent.Message.GuildID) // Set roles list here so not queried every time
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
		case "SegBot - Spectator":
			data.Settings.Roles.Spectator = role.ID
		}
	}
	return nil
}

// createData Internal, creates and returns data file
func createData(agent discord.Agent) Guild {
	data := Guild{
		GuildID: agent.Message.GuildID,
		Admins:  append(make([]string, 0), agent.Message.Author.ID),
		Settings: settings{
			Channels: settingsChannels{
				Commands:    agent.Message.ChannelID,
				Leaderboard: agent.Message.ChannelID,
				Success:     agent.Message.ChannelID,
				Started:     agent.Message.ChannelID,
				Location:    agent.Message.ChannelID,
			},
			Roles: settingsRoles{
				Admin:      "none",
				Registered: "none",
				Spectator:  "none",
			},
		},
	}
	if createRoles(agent, &data) != nil {
		_, _ = agent.Session.ChannelMessageSend(agent.Channel,
			"Failed to create roles, you'll have to create and configure the missing ones")
	}
	return data
}

// writeData Internal, checks if guild registered, if not registers guild
func writeData(agent discord.Agent, data Guild) error {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.Message.GuildID)

	exists, err := utils.CreateFileIfNotExist(path)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	if exists {
		discord.SendMessageWithMention("This Guild is already registered!", "", agent)
		return os.ErrExist
	}
	if Write(agent, data) != nil {
		return err
	}
	return nil
}

// GuildInit Create guild's data file
func GuildInit(agent discord.Agent) {
	data := createData(agent)
	if data.GuildID == "" {
		return
	}
	if writeData(agent, data) != nil {
		return
	}
	discord.RoleSet(data, "", "admin", agent)
	discord.SendMessageWithMention("Guild registered successfully!", "", agent)
}
