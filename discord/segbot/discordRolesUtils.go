package main

func discordRoleSetLoad(id, role string, agent discordAgent) error {
	roleId := "none"
	otherRoleId := "none"

	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return err
	}
	if id == "" {
		id = agent.message.Author.ID
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
		_ = agent.session.GuildMemberRoleAdd(agent.message.GuildID, id, roleId)
	}
	if otherRoleId != "none" {
		_ = agent.session.GuildMemberRoleRemove(agent.message.GuildID, id, otherRoleId)
	}
	return nil
}

func discordRoleSet(data GuildData, id, role string, agent discordAgent) {
	roleId := "none"
	otherRoleId := "none"

	if id == "" {
		id = agent.message.Author.ID
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
		_ = agent.session.GuildMemberRoleAdd(agent.message.GuildID, id, roleId)
	}
	if otherRoleId != "none" {
		_ = agent.session.GuildMemberRoleRemove(agent.message.GuildID, id, roleId)
	}
}
