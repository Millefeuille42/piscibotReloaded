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
			data.Location = "Offline"
		}
		message = fmt.Sprintf("%s%-9s- %s", message, data.Login, data.Login)
		time.Sleep(time.Millisecond * 500)
	}
	gAPiMutex.Unlock()
	if message == "```" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
