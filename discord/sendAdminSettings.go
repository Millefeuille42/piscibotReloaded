package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

func sendAdminSettings(session *discordgo.Session, message *discordgo.MessageCreate) {
	path := fmt.Sprintf("./data/guilds/%s.json", message.GuildID)
	settings := GuildData{}

	if !guildInitialCheck(session, message) {
		return
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		logErrorToChan(session, message, err)
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
		getChannelName(session, settings.Settings.Channels.Commands),
		getChannelName(session, settings.Settings.Channels.Leaderboard),
		getChannelName(session, settings.Settings.Channels.Success),
		getChannelName(session, settings.Settings.Channels.Started),
		getChannelName(session, settings.Settings.Channels.Location),
	)

	for i, admin := range settings.Admins {
		if i == len(settings.Admins)-1 {
			mess = fmt.Sprintf("%s@%s\n```", mess, getUser(session, admin))
			break
		}
		mess = fmt.Sprintf("%s@%s, ", mess, getUser(session, admin))
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, mess)
}
