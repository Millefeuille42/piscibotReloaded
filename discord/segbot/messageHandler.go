package main

import (
	"github.com/bwmarrin/discordgo"
	"piscibotReloaded/discord/segbot/commands"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/guild"
	"piscibotReloaded/discord/segbot/target"
	"strings"
)

// adminRouter Router for admin
func adminRouter(agent discord.Agent) {
	switch {
	case strings.HasPrefix(agent.Message.Content, "!init"):
		guild.GuildInit(agent)
	case strings.HasPrefix(agent.Message.Content, "!chan"):
		guild.AdminSetChan(agent)
	case strings.HasPrefix(agent.Message.Content, "!admin"):
		guild.AdminSet(agent)
	case agent.Message.Content == "!params":
		guild.AdminSendSettings(agent)
	case agent.Message.Content == "!purge":
		commands.AdminPurge(agent)
	}
}

func commandsRouter(agent discord.Agent) bool {
	switch {
	case strings.HasPrefix(agent.Message.Content, "!profile"):
		commands.SendTargetProfile(agent)
		return true
	case strings.HasPrefix(agent.Message.Content, "!list"):
		switch {
		case strings.HasPrefix(agent.Message.Content[6:], "students"):
			commands.SendStudentsList(agent)
			return true
		case strings.HasPrefix(agent.Message.Content[6:], "tracked"):
			commands.SendTrackedList(agent)
			return true
		case strings.HasPrefix(agent.Message.Content[6:], "projects"):
			commands.SendProjectList(agent)
			return true
		}
	case strings.HasPrefix(agent.Message.Content, "!leaderboard"):
		commands.SendLeaderboard(agent)
		return true
	case strings.HasPrefix(agent.Message.Content, "!project"):
		commands.SendProject(agent)
		return true
	}
	return false
}

// userRouter Router for user commands, returns true if a command was found
func userRouter(agent discord.Agent) bool {
	switch {
	case strings.HasPrefix(agent.Message.Content, "!start"):
		discordUser.Init(agent)
		return true
	case strings.HasPrefix(agent.Message.Content, "!track"):
		target.Track(agent)
		return true
	case strings.HasPrefix(agent.Message.Content, "!ping"):
		discordUser.SetPings(agent)
		return true
	case agent.Message.Content == "!settings":
		discordUser.SendSettings(agent)
		return true
	case agent.Message.Content == "!untrack":
		target.Untrack(agent)
		return true
	case agent.Message.Content == "!spectate":
		discordUser.SetSpectator(agent)
		return true
	case agent.Message.Content == "!help":
		commands.SendHelp(agent)
	}
	return false
}

// messageHandler Discord bot message handler
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	botID, _ := session.User("@me")
	agent := discord.Agent{
		Session: session,
		Message: message,
	}
	agent.Channel = guild.GetChannel(agent)

	if message.Author.ID == botID.ID || !strings.HasPrefix(message.Content, "!") {
		return
	}
	if !commandsRouter(agent) && !userRouter(agent) {
		adminRouter(agent)
	}
}
