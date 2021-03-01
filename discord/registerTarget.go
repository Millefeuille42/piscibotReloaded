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

type TargetSettings struct {
	Login      string
	GuildUsers map[string]string
}

func makeApiReq(path, login string, session *discordgo.Session, message *discordgo.MessageCreate) error {
	uri := fmt.Sprintf("%s:%s/user/%s", os.Getenv("42API"), os.Getenv("42PORT"), login)
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	if res.StatusCode == 404 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "This login doesn't exist")
		_ = os.Remove(path)
		return os.ErrNotExist
	}
	return nil
}

func loadOrCreate(path, login string, settings *TargetSettings, message *discordgo.MessageCreate) error {
	exists, err := createFileIfNotExist(path)
	if err != nil {
		_ = os.Remove(path)
		return err
	}
	if exists {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			_ = os.Remove(path)
			return err
		}
		err = json.Unmarshal(data, settings)
		if err != nil {
			return err
		}
	} else {
		*settings = TargetSettings{
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

func registerTarget(session *discordgo.Session, message *discordgo.MessageCreate) {
	settings := TargetSettings{}
	args := strings.Split(message.Content, "-")
	if len(args) < 2 {
		return
	}

	if !userInitialCheck(session, message) {
		return
	}

	path := fmt.Sprintf("./data/targets/%s.json", args[1])
	err := loadOrCreate(path, args[1], &settings, message)
	if err != nil {
		if err == os.ErrExist {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Someone is already tracking this person"+
				" on this server!")
			return
		}
		logErrorToChan(session, message, err)
		return
	}
	/*
		if err := makeApiReq(path, args[1], session, message); err != nil {
			return
		}
	*/

	settingsBytes, err := json.MarshalIndent(settings, "", "\t")
	err = ioutil.WriteFile(path, settingsBytes, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}

	_, _ = session.ChannelMessageSend(message.ChannelID, "You are now tracking "+args[1])
}
