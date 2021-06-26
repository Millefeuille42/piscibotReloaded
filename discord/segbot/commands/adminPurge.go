package commands

import (
	"fmt"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/guild"
	"piscibotReloaded/discord/segbot/utils"
	"time"
)

func channelPurge(agent discord.Agent, channel string) error {
	for {
		messages, err := agent.Session.ChannelMessages(channel, 100, "", "", "")
		if err != nil {
			return err
		}
		for _, message := range messages {
			err = agent.Session.ChannelMessageDelete(channel, message.ID)
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
	discord.SendMessageWithMention("Channel <#"+channel+"> purged!", "", agent)
	return nil
}

func AdminPurge(agent discord.Agent) {
	guild.InitialCheck(agent)
	data, err := guild.Load(agent, false, "")
	if err != nil {
		return
	}
	if !utils.Find(data.Admins, agent.Message.Author.ID) {
		_, _ = agent.Session.ChannelMessageSend(agent.Channel, "You are not an admin")
		return
	}
	discord.LogErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Location))
	discord.LogErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Started))
	discord.LogErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Success))
	discord.LogErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Leaderboard))
	discord.LogErrorToChan(agent, channelPurge(agent, data.Settings.Channels.Commands))
}
