package main

import "piscibotReloaded/discord/segbot/utils"

func adminLock(agent discordAgent) {
	guildInitialCheck(agent)
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}
	if !utils.Find(data.Admins, agent.message.Author.ID) {
		err = sendMessageWrapper(agent.session, agent.channel, "You are not an admin")
		return
	}
	if data.Locked {
		err = sendMessageWrapper(agent.session, agent.channel, "The guild is already locked")
		return
	}
	data.Locked = true
	sendMessageWithMention("Tracking is now locked!", "", agent)
	_ = guildWriteFile(agent, data)
}

func adminUnlock(agent discordAgent) {
	guildInitialCheck(agent)
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}
	if !utils.Find(data.Admins, agent.message.Author.ID) {
		err = sendMessageWrapper(agent.session, agent.channel, "You are not an admin")
		return
	}
	if !data.Locked {
		err = sendMessageWrapper(agent.session, agent.channel, "The guild is already unlocked")
		return
	}
	data.Locked = false
	sendMessageWithMention("Tracking is now unlocked!", "", agent)
	_ = guildWriteFile(agent, data)
}
