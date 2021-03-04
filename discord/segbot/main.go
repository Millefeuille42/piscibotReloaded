package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

var gBot *discordgo.Session

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	checkError(err)
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	checkError(err)
	fmt.Println("Discord bot created")

	setUpCloseHandler(discordBot)

	return discordBot
}

// prepFileSystem Create required directories
func prepFileSystem() error {
	err := createDirIfNotExist("./data")
	if err != nil {
		return err
	}
	err = createDirIfNotExist("./data/guilds")
	if err != nil {
		return err
	}
	err = createDirIfNotExist("./data/targets")
	if err != nil {
		return err
	}
	err = createDirIfNotExist("./data/users")
	return err
}

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "You must provide and env file")
		return
	}

	checkError(prepFileSystem())
	gBot = startBot()
	startServer()

	for {
		time.Sleep(time.Second * 3)
	}
}
