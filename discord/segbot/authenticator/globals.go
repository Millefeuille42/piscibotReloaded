package authenticator

import "github.com/bwmarrin/discordgo"

var gDiscordBot *discordgo.Session

func SetBot(bot *discordgo.Session) {
	gDiscordBot = bot
}
