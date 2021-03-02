package main

import "fmt"

// userSendSettings Send user's ping related settings to the channel
func userSendSettings(agent discordAgent) {
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
