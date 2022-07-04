package main

import (
	"fmt"
	"log"
	"time"
)

func sendPiscineuxList(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}

	message := "```\n"
	page := 1
	piscineux, err := Client.GetPiscineux("9", agent.args[1], agent.args[2], fmt.Sprintf("%d", page))
	for len(piscineux) > 0 {
		page++
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, user := range piscineux {
			message += user.Login + "\n"
		}
		piscineux, err = Client.GetPiscineux("9", agent.args[1], agent.args[2], fmt.Sprintf("%d", page))
		time.Sleep(1)
	}

	if message == "```\n" {
		sendMessageWithMention("Nothing to see here...", "", agent)
		return
	}
	sendMessageWithMention(message+"```", "", agent)
}
