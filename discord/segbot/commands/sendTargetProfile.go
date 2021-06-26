package commands

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/target"
	"strings"
)

func SendTargetProfile(agent discord.Agent) {
	message := ""

	if !discordUser.InitialCheck(agent) {
		return
	}
	userFile, err := discordUser.Load("", agent)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return
	}
	args := strings.Split(agent.Message.Content, " ")
	if len(args) < 2 {
		if _, ok := userFile.GuildTargets[agent.Message.GuildID]; !ok {
			discord.SendMessageWithMention("You must be tracking someone or provide logins", "", agent)
			return
		}
		args = append(args, userFile.GuildTargets[agent.Message.GuildID])
	}
	for _, login := range args[1:] {
		targetFile, err := target.Load(login, agent)
		if err != nil {
			discord.SendMessageWithMention(login+" not found!", "", agent)
			return
		}
		if _, ok := targetFile.GuildUsers[agent.Message.GuildID]; !ok {
			discord.SendMessageWithMention(login+" not found!", "", agent)
			return
		}
		data, err := targetGetData(agent, login)
		if err != nil {
			discord.LogErrorToChan(agent, err)
			return
		}
		if data.Location == nil {
			data.Location = "Offline"
		}
		userId := targetFile.GuildUsers[agent.Message.GuildID]
		userId = discord.GetUser(agent.Session, userId)
		message += fmt.Sprintf("```"+
			"Profile of %s (%s)\n"+
			"\tName:              %s\n"+
			"\tLocation:          %s\n"+
			"\tCorrection Points: %d\n"+
			"\tWallet:            %d₳\n",
			data.Login, userId, data.UsualFullName, data.Location, data.CorrectionPoint, data.Wallet)
		for _, cursus := range data.CursusUsers {
			message += fmt.Sprintf(""+
				"\t%s\n"+
				"\t\tLevel:         %.2f\n",
				cursus.Cursus.Slug, cursus.Level)
		}
		message += "```\n"
	}
	discord.SendMessageWithMention(message, "", agent)
}
