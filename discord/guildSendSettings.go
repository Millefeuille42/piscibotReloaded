package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func sendAdminSettings(agent discordAgent) {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(agent) {
		return
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(agent, err)
		return
	}

	mess := fmt.Sprintf(
		"```\n"+
			"Channels:\n"+
			"    Commands:     #%s\n"+
			"    Leaderboards: #%s\n"+
			"    Success:      #%s\n"+
			"    Started:      #%s\n"+
			"    Location:     #%s\n\n"+
			"Admins:           ",
		getChannelName(agent.session, settings.Settings.Channels.Commands),
		getChannelName(agent.session, settings.Settings.Channels.Leaderboard),
		getChannelName(agent.session, settings.Settings.Channels.Success),
		getChannelName(agent.session, settings.Settings.Channels.Started),
		getChannelName(agent.session, settings.Settings.Channels.Location),
	)

	for i, admin := range settings.Admins {
		if i == len(settings.Admins)-1 {
			mess = fmt.Sprintf("%s@%s\n```", mess, getUser(agent.session, admin))
			break
		}
		mess = fmt.Sprintf("%s@%s, ", mess, getUser(agent.session, admin))
	}
	_, _ = agent.session.ChannelMessageSend(agent.channel, mess)
}
