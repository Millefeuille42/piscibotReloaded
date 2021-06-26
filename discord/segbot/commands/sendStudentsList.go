package commands

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/target"
)

func SendStudentsList(agent discord.Agent) {
	if !discordUser.InitialCheck(agent) {
		return
	}
	piscis, err := discord.GetTargetsOfGuild(agent, agent.Message.GuildID)
	if err != nil {
		return
	}
	message := "```\n"
	for _, pisci := range piscis {
		isStudent := false
		targetFile, err := target.Load(pisci, agent)
		if err != nil {
			discord.SendMessageWithMention(pisci+" not found!", "", agent)
			return
		}
		data, err := targetGetData(agent, pisci)
		if err != nil {
			return
		}
		for _, cursus := range data.CursusUsers {
			if cursus.Cursus.Slug != "c-piscine" {
				isStudent = true
			}
		}
		token := '✘'
		if isStudent {
			token = '✔'
		}
		userId := targetFile.GuildUsers[agent.Message.GuildID]
		userId = discord.GetUser(agent.Session, userId)
		message += fmt.Sprintf("%-9s%c (%s)\n", pisci, token, userId)
	}
	message += "```"
	discord.SendMessageWithMention(message, "", agent)
}
