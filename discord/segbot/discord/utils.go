package discord

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"piscibotReloaded/discord/segbot/discordUser"
	"piscibotReloaded/discord/segbot/utils"
)

func SendMessageToUser(message, channel, userID, chanParam string, agent Agent) {
	switch chanParam {
	case "none":
		_, _ = agent.Session.ChannelMessageSend(channel, message)
	case "channel":
		_, _ = agent.Session.ChannelMessageSend(channel, fmt.Sprintf("<@%s>\n%s", userID, message))
	case "dm":
		dmChan, err := agent.Session.UserChannelCreate(userID)
		if err == nil {
			_, _ = agent.Session.ChannelMessageSend(dmChan.ID, message)
		}
		_, _ = agent.Session.ChannelMessageSend(channel, message)
	case "all":
		dmChan, err := agent.Session.UserChannelCreate(userID)
		if err == nil {
			_, _ = agent.Session.ChannelMessageSend(dmChan.ID, message)
		}
		_, _ = agent.Session.ChannelMessageSend(channel, fmt.Sprintf("<@%s>\n%s", userID, message))
	}
}

func GetUsersOfGuild(agent Agent, guild string) ([]discordUser.User, error) {
	var userList = make([]discordUser.User, 0)

	files, err := ioutil.ReadDir("./data/users")
	if err != nil {
		LogErrorToChan(agent, err)
		return nil, err
	}
	for _, f := range files {
		var user = discordUser.User{}

		fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s", f.Name()))
		if err != nil {
			utils.LogError(err)
			continue
		}
		err = json.Unmarshal(fileData, &user)
		if err != nil {
			utils.LogError(err)
			continue
		}
		if _, ok := user.GuildTargets[guild]; ok {
			userList = append(userList, user)
		}
	}
	return userList, nil
}

func GetTargetsOfGuild(agent Agent, guild string) ([]string, error) {
	var targetList = make([]string, 0)

	userList, err := GetUsersOfGuild(agent, guild)
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

// SendMessageWithMention Sends a discord message prepending a mention + \n to the message, if id == "", id becomes the message author
func SendMessageWithMention(message, id string, agent Agent) {
	if id == "" {
		id = agent.Message.Author.ID
	}

	_, _ = agent.Session.ChannelMessageSend(agent.Channel, fmt.Sprintf("<@%s>\n%s", id, message))
}

// GetUser Returns associated user of provided id
func GetUser(session *discordgo.Session, id string) string {
	ret, err := session.User(id)
	if err != nil {
		return ""
	}
	return ret.Username
}

// GetChannelName Returns associated channel name of provided id
func GetChannelName(session *discordgo.Session, id string) string {
	ret, _ := session.Channel(id)
	return ret.Name
}

// LogErrorToChan Sends plain error to command channel
func LogErrorToChan(agent Agent, err error) {
	if err == nil {
		return
	}
	utils.LogError(err)
	_, _ = agent.Session.ChannelMessageSend(agent.Channel,
		fmt.Sprintf("An Error Occured, Please Try Again Later {%s}", err.Error()))
}
