package main

import (
	"piscibotReloaded/discord/segbot/utils"
	"time"
)

func sendProjectList(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	projectList := make([]string, 0)
	message := "```\n"

	gAPiMutex.Lock()
	for _, target := range targets {
		data, err := targetGetData(agent, target)
		if err != nil {
			gAPiMutex.Unlock()
			return
		}
		for _, project := range data.ProjectsUsers {
			slug := project["project"].(map[string]interface{})["slug"]
			if !utils.Find(projectList, slug.(string)) {
				projectList = append(projectList, slug.(string))
				message += slug.(string) + "\n"
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
	gAPiMutex.Unlock()
	if message == "```" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
