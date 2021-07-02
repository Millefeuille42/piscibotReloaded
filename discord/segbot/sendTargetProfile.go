package main

import (
	"fmt"
	"strings"
)

func sendTargetProfile(agent discordAgent) {
	message := ""

	if !userInitialCheck(agent) {
		return
	}
	userFile, err := userLoadFile("", agent)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	args := strings.Split(agent.message.Content, " ")
	if len(args) < 2 {
		if _, ok := userFile.GuildTargets[agent.message.GuildID]; !ok {
			sendMessageWithMention("You must be tracking someone or provide login(s)", "", agent)
			return
		}
		args = append(args, userFile.GuildTargets[agent.message.GuildID])
	}
	for _, login := range args[1:] {
		targetFile, err := targetLoadFile(login, agent)
		if err != nil {
			sendMessageWithMention(login+" not found!", "", agent)
			return
		}
		if _, ok := targetFile.GuildUsers[agent.message.GuildID]; !ok {
			sendMessageWithMention(login+" not found!", "", agent)
			return
		}
		data, err := targetGetData(agent, login)
		if err != nil {
			logErrorToChan(agent, err)
			return
		}
		if data.Location == nil {
			data.Location = "Offline"
		}
		userId := targetFile.GuildUsers[agent.message.GuildID]
		userId = getUser(agent.session, userId)
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
	sendMessageWithMention(message, "", agent)
}
