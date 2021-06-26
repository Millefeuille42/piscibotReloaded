package discord

import "piscibotReloaded/discord/segbot/guild"

func RoleSetLoad(id, role string, agent Agent) error {
	roleId := "none"
	otherRoleId := "none"

	data, err := guild.Load(agent, false, "")
	if err != nil {
		return err
	}
	if id == "" {
		id = agent.Message.Author.ID
	}

	switch role {
	case "admin":
		roleId = data.Settings.Roles.Admin
	case "registered":
		roleId = data.Settings.Roles.Registered
		otherRoleId = data.Settings.Roles.Spectator
	case "spectator":
		roleId = data.Settings.Roles.Spectator
		otherRoleId = data.Settings.Roles.Registered
	}

	if roleId != "none" {
		_ = agent.Session.GuildMemberRoleAdd(agent.Message.GuildID, id, roleId)
	}
	if otherRoleId != "none" {
		_ = agent.Session.GuildMemberRoleRemove(agent.Message.GuildID, id, otherRoleId)
	}
	return nil
}

func RoleSet(data guild.Guild, id, role string, agent Agent) {
	roleId := "none"
	otherRoleId := "none"

	if id == "" {
		id = agent.Message.Author.ID
	}

	switch role {
	case "admin":
		roleId = data.Settings.Roles.Admin
	case "registered":
		roleId = data.Settings.Roles.Registered
		otherRoleId = data.Settings.Roles.Spectator
	case "spectator":
		roleId = data.Settings.Roles.Spectator
		otherRoleId = data.Settings.Roles.Registered
	}

	if roleId != "none" {
		_ = agent.Session.GuildMemberRoleAdd(agent.Message.GuildID, id, roleId)
	}
	if otherRoleId != "none" {
		_ = agent.Session.GuildMemberRoleRemove(agent.Message.GuildID, id, roleId)
	}
}
