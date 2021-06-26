package authenticator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"piscibotReloaded/discord/segbot/commands"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/guild"
	"piscibotReloaded/discord/segbot/target"
)

type MessageList []struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
}

type Message struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
}

// sendMessage Internal, sends the message to concerned user
func sendMessage(message Message) error {
	agent := discord.Agent{
		Session: gDiscordBot,
		Channel: os.Getenv("BOT_DEV_CHANNEL"),
	}

	pisci, err := target.Load(message.Login, agent)
	if err != nil {
		return err
	}

	for server, user := range pisci.GuildUsers {
		var channel string
		var param string

		userData, err := discordUser.Load(user, agent)
		if err != nil {
			discord.LogErrorToChan(agent, err)
			continue
		}
		guildData, err := guild.Load(agent, true, server)
		if err != nil {
			discord.LogErrorToChan(agent, err)
			continue
		}

		switch message.Channel {
		case "success":
			param = userData.Settings.Success
			channel = guildData.Settings.Channels.Success
			_, _ = agent.Session.ChannelMessageSend(guildData.Settings.Channels.Leaderboard,
				commands.CreateLeaderboard(agent, "c-piscine", server))
		case "started":
			param = userData.Settings.Started
			channel = guildData.Settings.Channels.Started
		case "location":
			param = userData.Settings.Location
			channel = guildData.Settings.Channels.Location
		}
		discord.SendMessageToUser(message.Message, channel, userData.UserID, param, agent)
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
func StartServer() {
	http.HandleFunc("/discord", sendHandler)
	http.HandleFunc("/auth", authHandler)
	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
