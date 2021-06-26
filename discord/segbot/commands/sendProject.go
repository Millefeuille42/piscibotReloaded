package commands

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"strings"
)

func getProjectState(project map[string]interface{}) string {
	if project["status"].(string) == "finished" {
		if project["validated?"] == nil {
			return "gave_up"
		}
		token := '✘'
		if project["validated?"].(bool) {
			token = '✔'
		}
		return fmt.Sprintf("%s - %3d %c",
			project["status"].(string),
			int(project["final_mark"].(float64)), token)
	}
	return project["status"].(string)
}

func SendProject(agent discord.Agent) {
	if !discordUser.InitialCheck(agent) {
		return
	}
	targets, err := discord.GetTargetsOfGuild(agent, agent.Message.GuildID)
	if err != nil {
		return
	}
	args := strings.Split(agent.Message.Content, " ")
	if len(args) < 2 {
		discord.SendMessageWithMention("You must provide project(s) to check", "", agent)
		return
	}
	message := "```"
	for _, arg := range args[1:] {
		found := false
		for _, target := range targets {
			data, err := targetGetData(agent, target)
			if err != nil {
				return
			}
			for _, project := range data.ProjectsUsers {
				slug := project["project"].(map[string]interface{})["slug"]
				if slug == arg {
					if !found {
						message += "\n" + arg + "\n"
						found = true
					}
					message += fmt.Sprintf("\t%-9s- %s\n", data.Login, getProjectState(project))
					break
				}
			}
		}
	}
	if message == "```" {
		discord.SendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	discord.SendMessageWithMention(message+"```", "", agent)
}
