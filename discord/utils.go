package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func getUser(session *discordgo.Session, id string) string {
	ret, err := session.User(id)
	if err != nil {
		return ""
	}
	return ret.Username
}

func getChannelName(session *discordgo.Session, id string) string {
	ret, _ := session.Channel(id)
	return ret.Name
}

func logErrorToChan(agent discordAgent, err error) {
	logError(err)
	_, _ = agent.session.ChannelMessageSend(agent.channel,
		fmt.Sprintf("An Error Occured, Please Try Again Later {%s}", err.Error()))
}

// Check if dir exists, if not create it
func createDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

// Check if file exists, if not create it
func createFileIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			logError(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Prints error + StackTrace to stderr if error
func logError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err, debug.Stack())
	}
}

// Panic if error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Setup an handler for Ctrl+C and closing the bot
func setupCloseHandler(session *discordgo.Session) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		_ = session.Close()
		os.Exit(0)
	}()
}
