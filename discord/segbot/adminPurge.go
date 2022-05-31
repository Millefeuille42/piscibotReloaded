package main

import (
	"fmt"
	"piscibotReloaded/discord/segbot/utils"
	"time"
)

func channelPurge(agent discordAgent, channel string) error {
	for {
		messages, err := agent.session.ChannelMessages(channel, 100, "", "", "")
		if err != nil {
			return err
		}
		for _, message := range messages {
			err = agent.session.ChannelMessageDelete(channel, message.ID)
			fmt.Println("Purging " + message.ID)
			if err != nil {
				return err
			}
			time.Sleep(time.Minute / 20)
		}
		if len(messages) < 100 {
			break
		}
	}
	sendMessageWithMention("Channel <#"+channel+"> purged!", "", agent)
	return nil
}

func adminPurge(agent discordAgent) {
	guildInitialCheck(agent)
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}
	if !utils.Find(data.Admins, agent.message.Author.ID) {
		err = sendMessageWrapper(agent.session, agent.channel, "You are not an admin")
		return
	}
	logErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Location))
	logErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Started))
	logErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Success))
	logErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Leaderboard))
	logErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Commands))
}
