package main

import (
	"fmt"
	"strings"
	"time"
)

func makeCursusString(data ApiData) string {
	message := ""
	for _, cursus := range data.CursusUsers {
		projectCount := 0
		for _, project := range data.ProjectsUsers {
			if project["cursus_ids"] == nil || len(project["cursus_ids"].([]interface{})) <= 0 {
				continue
			}
			for _, pCursus := range project["cursus_ids"].([]interface{}) {
				if int(pCursus.(float64)) == cursus.CursusID && isProjectValidated(project) {
					projectCount++
				}
			}
		}
		message += fmt.Sprintf(""+
			"\t%s\n"+
			"\t\tLevel:         %.2f\n"+
			"\t\tProjects OK:   %d\n",
			cursus.Cursus.Slug, cursus.Level, projectCount)
	}
	return message
}

func sendTargetProfile(agent discordAgent) {
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
	for _, login := range args[1:] {
		targetFile, err := targetLoadFile(login, agent)
		if err != nil {
			sendMessageWithMention(login+" not found!", "", agent)
			continue
		}
		if _, ok := targetFile.GuildUsers[agent.message.GuildID]; !ok {
			sendMessageWithMention(login+" not found!", "", agent)
			continue
		}
		data, err := targetGetData(agent, login)
		if err != nil {
			logErrorToChan(agent, err)
			continue
		}
		if data.Location == nil {
			data.Location = "Offline"
		}
		userId := targetFile.GuildUsers[agent.message.GuildID]
		userId = getUser(agent.session, userId)
		message += fmt.Sprintf("```"+
			"Profile of %s (%s)\n"+
			"\tName:              %s\n"+
			"\tLocation:          %s\n"+
			"\tCorrection Points: %d\n"+
			"\tWallet:            %d₳\n",
			data.Login, userId, data.UsualFullName, data.Location, data.CorrectionPoint, data.Wallet)
		message += makeCursusString(data)
		message += "```\n"
		time.Sleep(time.Millisecond * 500)
	}
	gAPiMutex.Unlock()
	sendMessageWithMention(message, "", agent)
}
