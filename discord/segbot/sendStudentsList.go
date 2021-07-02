package main

import "fmt"

func sendStudentsList(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	message := "```\n"
	for _, target := range targets {
		isStudent := false
		targetFile, err := targetLoadFile(target, agent)
		if err != nil {
			sendMessageWithMention(target+" not found!", "", agent)
			return
		}
		data, err := targetGetData(agent, target)
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
		userId := targetFile.GuildUsers[agent.message.GuildID]
		userId = getUser(agent.session, userId)
		message += fmt.Sprintf("%-9s%c (%s)\n", target, token, userId)
	}
	message += "```"
	sendMessageWithMention(message, "", agent)
}
