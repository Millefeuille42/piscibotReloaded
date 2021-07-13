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
	args := strings.Split(agent.message.Content, " ")
	if len(args) <= 1 {
		sendMessageWithMention("I need more arguments", "", agent)
		return
	}

	didSomething := false
	for _, channel := range args {
		subArgs := strings.Split(channel, ":")
		if len(subArgs) <= 1 || !Find([]string{"all", "none", "dm", "channel", "mention"}, subArgs[1]) {
			continue
		}
		if subArgs[1] == "mention" {
			subArgs[1] = "channel"
		}
		switch subArgs[0] {
		case "success":
			user.Settings.Success = subArgs[1]
		case "started":
			user.Settings.Started = subArgs[1]
		case "location":
			user.Settings.Location = subArgs[1]
		}
		didSomething = true
		_, _ = agent.session.ChannelMessageSend(agent.channel, "Ping settings updated for "+subArgs[0])
	}
	if didSomething && userWriteFile(user, agent, "") == nil {
		sendMessageWithMention("Ping settings saved", "", agent)
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
		"Success:     %s\n"+
		"Started:     %s\n"+
		"Location     %s\n"+
		"```", user.Settings.Success, user.Settings.Started, user.Settings.Location)
	sendMessageWithMention(message, "", agent)
}
