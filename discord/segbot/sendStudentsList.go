package main

import (
	"fmt"
	"time"
)

func sendStudentsList(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	message := "```\n"
	gAPiMutex.Lock()
	for _, target := range targets {
		isStudent := false
		targetFile, err := targetLoadFile(target, agent)
		if err != nil {
			sendMessageWithMention(target+" not found!", "", agent)
			gAPiMutex.Unlock()
			return
		}
		data, err := targetGetData(agent, target)
		if err != nil {
			gAPiMutex.Unlock()
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
		time.Sleep(time.Millisecond * 500)
	}
	gAPiMutex.Unlock()
	message += "```"
	sendMessageWithMention(message, "", agent)
}
