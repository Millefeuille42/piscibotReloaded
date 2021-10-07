package main

import "fmt"

func sendTrackedList(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	message := "```\n"
	for _, target := range targets {
		targetFile, err := targetLoadFile(target, agent)
		if err != nil {
			sendMessageWithMention(target+" not found!", "", agent)
			return
		}
		userId := targetFile.GuildUsers[agent.message.GuildID]
		userId = getUser(agent.session, userId)
		message += fmt.Sprintf("%-9s (%s)\n", target, userId)
	}
	message += "```"
	sendMessageWithMention(message, "", agent)
}
