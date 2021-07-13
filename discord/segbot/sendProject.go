package main

import (
	"fmt"
	"strings"
	"time"
)

func isProjectValidated(project map[string]interface{}) bool {
	if project["status"].(string) == "finished" {
		if project["validated?"] == nil {
			return false
		}
		if project["validated?"].(bool) {
			return true
		}
	}
	return false
}

func getProjectState(project map[string]interface{}) string {
	if project["status"].(string) == "finished" {
		if project["validated?"] == nil {
			return "gave_up"
		}
		token := '✘'
		if project["validated?"].(bool) {
			token = '✔'
		}
		return fmt.Sprintf("%-11.11s - %3d %c",
			project["status"].(string),
			int(project["final_mark"].(float64)), token)
	}
	return fmt.Sprintf("%11.11s", project["status"].(string))
}

func sendProject(agent discordAgent) {
	message := ""

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
	gAPiMutex.Lock()
	for _, arg := range args[1:] {
		newMessage := "```\n"
		found := false
		for _, target := range targets {
			data, err := targetGetData(agent, target)
			if err != nil {
				break
			}
			for _, project := range data.ProjectsUsers {
				slug := project["project"].(map[string]interface{})["slug"]
				if slug == arg {
					if !found {
						newMessage += arg + "\n"
						found = true
					}
					newMessage += fmt.Sprintf("\t%-9s- %s\n", data.Login, getProjectState(project))
					break
				}
			}
			time.Sleep(time.Millisecond * 500)
		}
		if newMessage == "```\n" {
			sendMessageWithMention("Nothing to see here...("+arg+")", "", agent)
			continue
		}
		message += newMessage + "```\n"
	}
	sendMessageWithMention(message, "", agent)
	gAPiMutex.Unlock()
}

func sendUserProject(agent discordAgent) {
	message := ""

	if !userInitialCheck(agent) {
		return
	}
	userFile, err := userLoadFile("", agent)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	args := strings.Split(agent.message.Content, " ")
	if len(args) < 2 {
		if _, ok := userFile.GuildTargets[agent.message.GuildID]; !ok {
			sendMessageWithMention("You must be tracking someone or provide login(s)", "", agent)
			return
		}
		args = append(args, userFile.GuildTargets[agent.message.GuildID])
	}
	gAPiMutex.Lock()
	for _, arg := range args[1:] {
		targetFile, err := targetLoadFile(arg, agent)
		if err != nil {
			sendMessageWithMention(arg+" not found!", "", agent)
			continue
		}
		if _, ok := targetFile.GuildUsers[agent.message.GuildID]; !ok {
			sendMessageWithMention(arg+" not found!", "", agent)
			continue
		}
		newMessage := "```\n" + arg + " project's\n"
		data, err := targetGetData(agent, arg)
		if err != nil {
			continue
		}
		for _, project := range data.ProjectsUsers {
			slug := project["project"].(map[string]interface{})["slug"]
			newMessage += fmt.Sprintf("\t%-20s-     %s\n", slug, getProjectState(project))
		}
		time.Sleep(time.Millisecond * 500)
		if newMessage == arg+" project's\n" {
			sendMessageWithMention("Nothing to see here... ("+arg+")", "", agent)
			gAPiMutex.Unlock()
			continue
		}
		message += newMessage + "```\n"
	}
	if message != "" {
		sendMessageWithMention(message, "", agent)
	}
	gAPiMutex.Unlock()
}
