package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

// Find Check if val exists in Slice, true if it exists
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
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
