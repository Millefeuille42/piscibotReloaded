package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"piscibotReloaded/discord/segbot/utils"
	"sync"
	"time"
)

var gBot *discordgo.Session
var ownerID = "268431730967314435" //Please change this when using my bot
var gAPiMutex = sync.Mutex{}
var gPrefix = os.Getenv("SEGBOT_PREFIX")
var commandMap = make(map[string]commandHandler)

func setupFunctionsMap() {
	//AdminCommands no args
	commandMap["init"] = guildInit
	commandMap["params"] = adminSendSettings
	commandMap["purge"] = adminPurge
	commandMap["lock"] = adminLock
	commandMap["unlock"] = adminUnlock
	commandMap["force-untrack"] = adminForceUntrack
	//ARGS
	commandMap["chan"] = adminSetChan
	commandMap["admin"] = adminSet

	//UserCommands no args
	commandMap["start"] = userInit
	commandMap["settings"] = userSendSettings
	commandMap["untrack"] = targetUntrack
	commandMap["spectate"] = userSetSpectator
	commandMap["help"] = sendHelp
	//ARGS
	commandMap["track"] = targetTrack
	commandMap["ping"] = userSetPings

	//Commands
	commandMap["profile"] = sendTargetProfile
	commandMap["list-students"] = sendStudentsList
	commandMap["list-tracked"] = sendTrackedList
	commandMap["list-projects"] = sendProjectList
	commandMap["list-location"] = sendLocationList
	commandMap["leaderboard"] = sendLeaderboard
	commandMap["project"] = sendProject
	commandMap["user-project"] = sendUserProject
}

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	utils.CheckError(err)
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	utils.CheckError(err)
	fmt.Println("Discord bot created")
	if os.Getenv("SEGBOT_IN_PROD") == "" {
		channel, err := discordBot.UserChannelCreate(ownerID)
		if err != nil {
			return nil
		}
		hostname, _ := os.Hostname()
		_, _ = discordBot.ChannelMessageSend(channel.ID, "Bot up - "+
			time.Now().Format(time.Stamp)+" - "+hostname)
	}
	if gPrefix == "" {
		gPrefix = "!"
	}
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
	setupFunctionsMap()
	gBot = startBot()
	startServer()

	for {
		time.Sleep(time.Second * 3)
	}
}
