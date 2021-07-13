package main

import (
	"fmt"
	"strings"
	"time"
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

func sendProject(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	args := strings.Split(agent.message.Content, " ")
	if len(args) < 2 {
		sendMessageWithMention("You must provide project(s) to check", "", agent)
		return
	}
	message := "```"
	gAPiMutex.Lock()
	for _, arg := range args[1:] {
		found := false
		for _, target := range targets {
			data, err := targetGetData(agent, target)
			if err != nil {
				gAPiMutex.Unlock()
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
			time.Sleep(time.Millisecond * 500)
		}
	}
	gAPiMutex.Unlock()
	if message == "```" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
