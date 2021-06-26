package commands

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/target"
)

func SendTrackedList(agent discord.Agent) {
	if !discordUser.InitialCheck(agent) {
		return
	}
	piscis, err := discord.GetTargetsOfGuild(agent, agent.Message.GuildID)
	if err != nil {
		return
	}
	message := "```\n"
	for _, pisci := range piscis {
		targetFile, err := target.Load(pisci, agent)
		if err != nil {
			discord.SendMessageWithMention(pisci+" not found!", "", agent)
			return
		}
		userId := targetFile.GuildUsers[agent.Message.GuildID]
		userId = discord.GetUser(agent.Session, userId)
		message += fmt.Sprintf("%-9s (%s)\n", pisci, userId)
	}
	message += "```"
	discord.SendMessageWithMention(message, "", agent)
}
