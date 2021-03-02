package main

import (
	"fmt"
	"strings"
)

// userSetPings Sets user ping
func userSetPings(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}
	args := strings.Split(agent.message.Content, "-")

	for _, channel := range args {
		subArgs := strings.Split(channel, ":")
		if len(subArgs) <= 1 || !Find([]string{"all", "none", "dm", "channel"}, subArgs[1]) {
			continue
		}
		switch subArgs[0] {
		case "leaderboard":
			user.Settings.Leaderboard = subArgs[1]
		case "success":
			user.Settings.Success = subArgs[1]
		case "started":
			user.Settings.Started = subArgs[1]
		case "location":
			user.Settings.Location = subArgs[1]
		}
	}
	if userWriteFile(user, agent) == nil {
		sendMessageWithMention("Ping settings updated", "", agent)
	}
}

// userSendSettings Send user's ping related settings to the channel
func userSendSettings(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}

	message := fmt.Sprintf("```\n"+
		"Leaderboard: %s\n"+
		"Success:     %s\n"+
		"Started:     %s\n"+
		"Location     %s\n"+
		"```", user.Settings.Leaderboard, user.Settings.Success, user.Settings.Started, user.Settings.Location)
	sendMessageWithMention(message, "", agent)
}
