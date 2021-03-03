package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var gBot *discordgo.Session

// Load env based on provided file
func loadEnv() {
	err := godotenv.Load(os.Args[1])
	checkError(err)
	fmt.Println("Loaded Env")
}

// Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOTTOKEN"))
	checkError(err)
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	checkError(err)
	fmt.Println("Discord bot created")

	setupCloseHandler(discordBot)

	return discordBot
}

// Create required directories
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
	loadEnv()
	gBot = startBot()
	startServer()

	for {
		time.Sleep(time.Second * 3)
	}
}
