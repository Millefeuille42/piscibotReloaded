package main

import (
	"fmt"
	"time"
)

func sendLocationList(agent discordAgent) {
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
		data, err := targetGetData(agent, target)
		if err != nil {
			gAPiMutex.Unlock()
			return
		}
		if data.Location == nil {
			data.Location = "✘"
		}
		message = fmt.Sprintf("%s%-9s- %s\n", message, data.Login, data.Location)
		time.Sleep(time.Millisecond * 500)
	}
	gAPiMutex.Unlock()
	if message == "```\n" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
