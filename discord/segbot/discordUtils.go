package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"strings"
)

// sendMessageWithMention Sends a discord message according to user params
func sendMessageToUser(message, channel, userID, chanParam string, agent discordAgent) {
	switch chanParam {
	case "none":
		_, _ = agent.session.ChannelMessageSend(channel, message)
	case "channel":
		_, _ = agent.session.ChannelMessageSend(channel, fmt.Sprintf("<@%s>\n%s", userID, message))
	case "dm":
		dmChan, err := agent.session.UserChannelCreate(userID)
		if err == nil {
			_, _ = agent.session.ChannelMessageSend(dmChan.ID, message)
		}
		_, _ = agent.session.ChannelMessageSend(channel, message)
	case "all":
		dmChan, err := agent.session.UserChannelCreate(userID)
		if err == nil {
			_, _ = agent.session.ChannelMessageSend(dmChan.ID, message)
		}
		_, _ = agent.session.ChannelMessageSend(channel, fmt.Sprintf("<@%s>\n%s", userID, message))
	}
}

func getUsersOfGuild(agent discordAgent, guild string) ([]UserData, error) {
	var userList = make([]UserData, 0)

	files, err := ioutil.ReadDir("./data/users")
	if err != nil {
		logErrorToChan(agent, err)
		return nil, err
	}
	for _, f := range files {
		var user = UserData{}

		fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s", f.Name()))
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}
		if err != nil {
			fmt.Println(f.Name())
			fmt.Println(string(fileData))
			logError(err)
			continue
		}
		err = json.Unmarshal(fileData, &user)
		if err != nil {
			fmt.Println(f.Name())
			fmt.Println(string(fileData))
			logError(err)
			continue
		}
		if _, ok := user.GuildTargets[guild]; ok {
			userList = append(userList, user)
		}
	}
	return userList, nil
}

func getTargetsOfGuild(agent discordAgent, guild string) ([]string, error) {
	var targetList = make([]string, 0)

	userList, err := getUsersOfGuild(agent, guild)
	if err != nil {
		return nil, err
	}
	for _, user := range userList {
		if target, ok := user.GuildTargets[guild]; ok {
			targetList = append(targetList, target)
		}
	}
	return targetList, nil
}

// sendMessageWithMention Sends a discord message prepending a mention + \n to the message, if id == "", id becomes the message author
func sendMessageWithMention(message, id string, agent discordAgent) {
	if id == "" {
		id = agent.message.Author.ID
	}

	_, err := agent.session.ChannelMessageSend(agent.channel, fmt.Sprintf("<@%s>\n%s", id, message))
	if err != nil {
		logError(err)
	}
}

// getUser Returns associated user of provided id
func getUser(session *discordgo.Session, id string) string {
	ret, err := session.User(id)
	if err != nil {
		return ""
	}
	return ret.Username
}

// getChannelName Returns associated channel name of provided id
func getChannelName(session *discordgo.Session, id string) string {
	ret, _ := session.Channel(id)
	return ret.Name
}

// logErrorToChan Sends plain error to command channel
func logErrorToChan(agent discordAgent, err error) {
	if err == nil {
		return
	}
	logError(err)
	_, _ = agent.session.ChannelMessageSend(agent.channel,
		fmt.Sprintf("An Error Occured, Please Try Again Later {%s}", err.Error()))
}
