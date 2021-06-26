package commands

import (
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/utils"
)

func SendProjectList(agent discord.Agent) {
	if !discordUser.InitialCheck(agent) {
		return
	}
	targets, err := discord.GetTargetsOfGuild(agent, agent.Message.GuildID)
	if err != nil {
		return
	}
	projectList := make([]string, 0)
	message := "```"

	for _, target := range targets {
		data, err := targetGetData(agent, target)
		if err != nil {
			return
		}
		for _, project := range data.ProjectsUsers {
			slug := project["project"].(map[string]interface{})["slug"]
			if !utils.Find(projectList, slug.(string)) {
				projectList = append(projectList, slug.(string))
				message += slug.(string) + "\n"
			}
		}
	}
	if message == "```" {
		discord.SendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	discord.SendMessageWithMention(message+"```", "", agent)
}
