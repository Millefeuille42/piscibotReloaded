package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"piscibotReloaded/discord/segbot/utils"
	"strings"
)

type MessageList []struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
	Cursus  string `json:"cursus"`
}

type Message struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
	Cursus  string `json:"cursus"`
}

func sendMessageToOld(agent discordAgent, message Message) error {
	files, err := ioutil.ReadDir("./data/users")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}

	for _, f := range files {
		var user = UserData{}
		var channel string
		var param string

		fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s", f.Name()))
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}
		if err != nil {
			fmt.Println(f.Name())
			fmt.Println(string(fileData))
			utils.LogError(err)
			continue
		}
		err = json.Unmarshal(fileData, &user)
		if err != nil {
			fmt.Println(f.Name())
			fmt.Println(string(fileData))
			utils.LogError(err)
			continue
		}
		for guild, target := range user.ExGuildTargets {
			if target == message.Login {
				guildData, err := guildLoadFile(agent, true, guild)
				if err != nil {
					logErrorToChan(agent, err)
					continue
				}
				param = user.Settings.Success
				channel = guildData.Settings.Channels.Commands
			}
			sendMessageToUser(message.Message, channel, user.UserID, param, agent)
		}
	}
	return err
}

// sendMessage Internal, sends the message to concerned user
func sendMessage(message Message) error {
	agent := discordAgent{
		session: gBot,
		channel: os.Getenv("BOT_DEV_CHANNEL"),
	}

	if message.Channel == "student" {
		return sendMessageToOld(agent, message)
	}

	target, err := targetLoadFile(message.Login, agent)
	if err != nil {
		return err
	}

	for guild, user := range target.GuildUsers {
		var channel string
		var param string

		userData, err := userLoadFile(user, agent)
		if err != nil {
			logErrorToChan(agent, err)
			continue
		}
		guildData, err := guildLoadFile(agent, true, guild)
		if err != nil {
			logErrorToChan(agent, err)
			continue
		}

		switch message.Channel {
		case "success":
			param = userData.Settings.Success
			channel = guildData.Settings.Channels.Success
			err = sendMessageWrapper(agent.session, guildData.Settings.Channels.Leaderboard,
				"```"+createLeaderboard(agent, message.Cursus, guild)+"```")
		case "started":
			param = userData.Settings.Started
			channel = guildData.Settings.Channels.Started
		case "location":
			param = userData.Settings.Location
			channel = guildData.Settings.Channels.Location
		}
		sendMessageToUser(message.Message, channel, userData.UserID, param, agent)
	}
	return err
}

// sendHandler Internal, the handler for the endpoint
func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w = writeErrorToResponse(w, 500, err.Error())
			return
		}
		defer r.Body.Close()

		messages := parseMessage(data)
		if len(messages) <= 0 {
			w = writeErrorToResponse(w, 400, "No messages")
			return
		}
		for _, message := range messages {
			if message.Message == "" {
				w = writeErrorToResponse(w, 500, "Error parsing the message")
				return
			}

			if err = sendMessage(message); err != nil {
				w = writeErrorToResponse(w, 500, "Error sending the message\n"+err.Error())
				return
			}
		}
	}
	w.WriteHeader(200)
}

// startServer Starts the http endpoint for sending messages
func startServer() {
	http.HandleFunc("/discord", sendHandler)
	http.HandleFunc("/auth", authHandler)
	fmt.Println("Starting server")
	if err := http.ListenAndServe(":"+os.Getenv("SEGBOT_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
