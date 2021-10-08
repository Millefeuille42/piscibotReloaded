package main

import (
	"piscibotReloaded/discord/segbot/utils"
)

func adminForceUntrack(agent discordAgent) {
	guildInitialCheck(agent)
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}
	if !utils.Find(data.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
		return
	}
	targets, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return
	}
	for _, target := range targets {
		targetFile, err := targetLoadFile(target, agent)
		if err != nil {
			sendMessageWithMention(target+" not found!", "", agent)
			return
		}
		delete(targetFile.GuildUsers, agent.message.GuildID)
		err = targetWriteFile(targetFile, agent)
		if err != nil {
			return
		}

		userId := targetFile.GuildUsers[agent.message.GuildID]
		user, err := userLoadFile(userId, agent)
		if err != nil {
			sendMessageWithMention("User "+userId+" not found!", "", agent)
			return
		}
		delete(user.GuildTargets, agent.message.GuildID)
		err = userWriteFile(user, agent, "")
		if err != nil {
			return
		}
		_ = discordRoleSetLoad(userId, "spectator", agent)
		sendMessageWithMention("You are not tracking someone on this server anymore!", userId, agent)
	}
}

//func targetUntrack(agent discordAgent) {
//	data, err := guildLoadFile(agent, false, "")
//	if err != nil {
//		return
//	}
//	if data.Locked {
//		sendMessageWithMention("This server is locked, you can't change your tracking settings", "", agent)
//		return
//	}
//	if !userIsTrackingCheck(agent) {
//		return
//	}
//
//	user, err := userLoadFile("", agent)
//	if err != nil {
//		return
//	}
//	targetName := user.GuildTargets[agent.message.GuildID]
//	delete(user.GuildTargets, agent.message.GuildID)
//	err = userWriteFile(user, agent, "")
//	if err != nil {
//		return
//	}
//
//	target, err := targetLoadFile(targetName, agent)
//	if err != nil {
//		return
//	}
//	delete(target.GuildUsers, agent.message.GuildID)
//	err = targetWriteFile(target, agent)
//	if err != nil {
//		return
//	}
//	_ = discordRoleSetLoad("", "spectator", agent)
//	sendMessageWithMention("You are not tracking someone on this server anymore!", "", agent)
//}
