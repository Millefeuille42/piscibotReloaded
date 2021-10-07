package main

func adminLock(agent discordAgent) {
	guildInitialCheck(agent)
	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return
	}
	if !Find(data.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
		return
	}
	if data.Locked {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "The guild is already locked")
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
	if !Find(data.Admins, agent.message.Author.ID) {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "You are not an admin")
		return
	}
	if !data.Locked {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "The guild is already unlocked")
		return
	}
	data.Locked = false
	sendMessageWithMention("Tracking is now unlocked!", "", agent)
	_ = guildWriteFile(agent, data)
}
