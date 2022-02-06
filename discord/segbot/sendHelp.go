package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func sendHelp(agent discordAgent) {
	link := "https://rb.gy/ej2itp" //"https://github.com/Millefeuille42/piscibotReloaded#commands"

	i := rand.Intn(2)
	fmt.Println(i)
	if i == 1 {
		link = "https://rb.gy/enaq3a" //"https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	}
	sendMessageWithMention("Here. And, please, stop asking for help. Look the message history, somebody probably already asked...", "", agent)
	_, _ = agent.session.ChannelMessageSendEmbed(agent.channel, &discordgo.MessageEmbed{
		URL:         link,
		Type:        "link",
		Title:       "Help",
		Description: "Everyone has been in a situation where they could use some help. Unfortunately, it can feel tough to ask for help. Maybe you feel embarrassed or scared that you’ll be turned down. Don’t worry! Once you figure out what you need, you can make a polite and organized request. Chances are someone will be happy to give you the help you need!",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn-icons-png.flaticon.com/512/682/682055.png",
		},
	})
}
