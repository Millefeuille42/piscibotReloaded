package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// TargetData Contains a target Login, and a GuildUsers map
type TargetData struct {
	Login      string
	GuildUsers map[string]string
}

// makeApiReq Internal, Make calls to the 42API module to start data collecting and check if user exists
func makeApiReq(path, login string, agent discordAgent) error {
	uri := fmt.Sprintf("http://%s:%s/user/%s", os.Getenv("42API"), os.Getenv("42PORT"), login)
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	if res.StatusCode == 404 {
		_, _ = agent.session.ChannelMessageSend(agent.channel, "This login doesn't exist")
		_ = os.Remove(path)
		return os.ErrNotExist
	}
	return nil
}

// loadOrCreate Internal, Loads or creates Target file
func loadOrCreate(path, login string, settings *TargetData, message *discordgo.MessageCreate) error {
	exists, err := createFileIfNotExist(path)
	if err != nil {
		return err
	}
	if exists {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, settings)
		if err != nil {
			return err
		}
	} else {
		*settings = TargetData{
			Login:      login,
			GuildUsers: make(map[string]string),
		}
	}
	if _, isExist := settings.GuildUsers[message.GuildID]; !isExist {
		settings.GuildUsers[message.GuildID] = message.Author.ID
	} else {
		return os.ErrExist
	}
	return nil
}

// targetRegister Registers target for user and guild
func targetRegister(agent discordAgent) {
	settings := TargetData{}
	args := strings.Split(agent.message.Content, "-")
	if userCheckHasTarget(agent) != nil {
		return
	}
	if len(args) < 2 {
		sendMessageWithMention("I need more arguments!", "", agent)
		return
	}

	path := fmt.Sprintf("./data/targets/%s.json", args[1])
	err := loadOrCreate(path, args[1], &settings, agent.message)
	if err != nil {
		if err == os.ErrExist {
			sendMessageWithMention("Someone is already tracking this person"+
				" on this server!", "", agent)
			return
		}
		logErrorToChan(agent, err)
		return
	}

	if makeApiReq(path, args[1], agent) != nil {
		return
	}

	user, err := userLoadFile("", agent)
	if err != nil {
		return
	}
	user.GuildTargets[agent.message.GuildID] = settings.Login

	if targetWriteFile(settings, agent) == nil && userWriteFile(user, agent) == nil {
		sendMessageWithMention("You are now tracking "+args[1], "", agent)
	}
}
