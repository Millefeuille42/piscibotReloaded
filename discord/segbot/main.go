package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"piscibotReloaded/discord/segbot/authenticator"
	"piscibotReloaded/discord/segbot/utils"
	"time"
)

var gBot *discordgo.Session
var ownerID string = "268431730967314435" //Please change this when using my bot

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	utils.CheckError(err)
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	utils.CheckError(err)
	fmt.Println("Discord bot created")
	channel, err := discordBot.UserChannelCreate(ownerID)
	if err != nil {
		return nil
	}
	hostname, _ := os.Hostname()
	_, _ = discordBot.ChannelMessageSend(channel.ID, "Bot up - "+
		time.Now().Format(time.Stamp)+hostname)

	utils.SetUpCloseHandler(discordBot)

	return discordBot
}

// prepFileSystem Create required directories
func prepFileSystem() error {
	err := utils.CreateDirIfNotExist("./data")
	if err != nil {
		return err
	}
	err = utils.CreateDirIfNotExist("./data/guilds")
	if err != nil {
		return err
	}
	err = utils.CreateDirIfNotExist("./data/targets")
	if err != nil {
		return err
	}
	err = utils.CreateDirIfNotExist("./data/users")
	return err
}

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "You must provide and env file")
		return
	}

	utils.CheckError(prepFileSystem())
	gBot = startBot()
	authenticator.SetBot(gBot)
	authenticator.StartServer()

	for {
		time.Sleep(time.Second * 3)
	}
}
