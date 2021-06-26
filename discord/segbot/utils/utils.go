package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

// CreateDirIfNotExist Check if dir exists, if not create it
func CreateDirIfNotExist(path string) error {
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

// CreateFileIfNotExist Check if file exists, if not create it
func CreateFileIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			LogError(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

// Find Check if val exists in Slice, true if it exists
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// LogError Prints error + StackTrace to stderr if error
func LogError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err, string(debug.Stack()))
	}
}

// CheckError Panic if error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// SetUpCloseHandler Set up a handler for Ctrl+C and closing the bot
func SetUpCloseHandler(session *discordgo.Session) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		_ = session.Close()
		os.Exit(0)
	}()
}
