package discordUser

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/utils"
	"strings"
)

// SetPings Sets user ping
func SetPings(agent discord.Agent) {
	if !InitialCheck(agent) {
		return
	}
	user, err := Load("", agent)
	if err != nil {
		return
	}
	args := strings.Split(agent.Message.Content, " ")
	if len(args) <= 1 {
		discord.SendMessageWithMention("I need more arguments", "", agent)
		return
	}

	for _, channel := range args {
		subArgs := strings.Split(channel, ":")
		if len(subArgs) <= 1 || !utils.Find([]string{"all", "none", "dm", "channel"}, subArgs[1]) {
			continue
		}
		switch subArgs[0] {
		case "success":
			user.Settings.Success = subArgs[1]
		case "started":
			user.Settings.Started = subArgs[1]
		case "location":
			user.Settings.Location = subArgs[1]
		}
		_, _ = agent.Session.ChannelMessageSend(agent.Channel, "Ping settings updated for "+subArgs[0])
	}
	if Write(user, agent, "") == nil {
		discord.SendMessageWithMention("Ping settings saved", "", agent)
	}
}

// SendSettings Send user's ping related settings to the channel
func SendSettings(agent discord.Agent) {
	if !InitialCheck(agent) {
		return
	}
	user, err := Load("", agent)
	if err != nil {
		return
	}

	message := fmt.Sprintf("```\n"+
		"Success:     %s\n"+
		"Started:     %s\n"+
		"Location     %s\n"+
		"```", user.Settings.Success, user.Settings.Started, user.Settings.Location)
	discord.SendMessageWithMention(message, "", agent)
}
