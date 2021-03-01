package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

func targetWriteFile(data TargetData, session *discordgo.Session, message *discordgo.MessageCreate) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/targets/%s.json", data.Login), dataBytes, 0677)
	if err != nil {
		logErrorToChan(session, message, err)
		return err
	}
	return nil
}

func targetLoadFile(id string, session *discordgo.Session, message *discordgo.MessageCreate) (TargetData, error) {
	target := TargetData{}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s.json", id))
	if err != nil {
		logErrorToChan(session, message, err)
		return TargetData{}, err
	}

	err = json.Unmarshal(fileData, &target)
	if err != nil {
		logErrorToChan(session, message, err)
		return TargetData{}, err
	}

	return target, nil
}
