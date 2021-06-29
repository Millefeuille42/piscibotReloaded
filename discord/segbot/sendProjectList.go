package main

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

	for _, target := range targets {
		data, err := targetGetData(agent, target)
		if err != nil {
			return
		}
		for _, project := range data.ProjectsUsers {
			slug := project["project"].(map[string]interface{})["slug"]
			if !Find(projectList, slug.(string)) {
				projectList = append(projectList, slug.(string))
				message += slug.(string) + "\n"
			}
		}
	}
	if message == "```" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
