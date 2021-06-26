package discordUser

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"piscibotReloaded/discord/segbot/authenticator"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/guild"
	"piscibotReloaded/discord/segbot/utils"
)

func SetSpectator(agent discord.Agent) {
	if !InitialCheck(agent) {
		return
	}
	user, err := Load("", agent)
	if err != nil {
		return
	}
	if _, isExist := user.GuildTargets[agent.Message.GuildID]; isExist {
		discord.SendMessageWithMention("You can't be a spectator if you are tracking someone!", "", agent)
		return
	}
	_ = discord.RoleSetLoad("", "spectator", agent)
	discord.SendMessageWithMention("You are now spectating", "", agent)
}

// Init Initializes user
func Init(agent discord.Agent) {
	path := fmt.Sprintf("./data/users/%s.json", agent.Message.Author.ID)

	if !guild.InitialCheck(agent) {
		return
	}

	exists, err := utils.CreateFileIfNotExist(path)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return
	}
	if exists {
		discord.SendMessageWithMention("You are already registered!", "", agent)
		return
	}

	link, state := authenticator.LinkCreator(agent)
	if link == "" {
		discord.SendMessageWithMention("Could not generate OAuth link", "", agent)
		return
	}

	data := User{
		UserID:       agent.Message.Author.ID,
		State:        state,
		GuildTargets: make(map[string]string),
		Settings: settings{
			Success:  "none",
			Started:  "none",
			Location: "none",
		},
		Verified: false,
	}
	if Write(data, agent, "") != nil {
		return
	}

	discord.SendMessageWithMention("", "", discord.Agent{
		Session: agent.Session, Message: agent.Message, Channel: agent.Message.ChannelID})

	_, err = agent.Session.ChannelMessageSendEmbed(agent.Message.ChannelID, &discordgo.MessageEmbed{
		URL:   link,
		Type:  "link",
		Title: "Verification Link",
		Description: "You are now registered, validate your profile with the link provided.\n" +
			"You will not be able to perform actions until you validate your profile through 42",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8d/42_Logo.svg/1200px-42_Logo.svg.png",
		},
	})
}
