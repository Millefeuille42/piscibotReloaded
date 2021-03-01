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
	Username string
	Users    []string
}

func loadOrCreate(path, id, login string, settings *TargetSettings) error {
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
			_ = os.Remove(path)
			return err
		}
	} else {
		*settings = TargetSettings{
			Username: login,
			Users:    make([]string, 0),
		}
	}
	settings.Users = append(settings.Users, id)
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
	err := loadOrCreate(path, message.Author.ID, args[1], &settings)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}

	uri := fmt.Sprintf("%s:%s/user/%s", os.Getenv("42API"), os.Getenv("42PORT"), args[1])
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logErrorToChan(session, message, err)
		return
	}
	if res.StatusCode == 404 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "This login doesn't exist")
		_ = os.Remove(path)
		return
	}
	_, _ = session.ChannelMessageSend(message.ChannelID, "You are now tracking "+args[1])
}
