package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// targetWriteFile Writes target data to file
func targetWriteFile(data TargetData, agent discordAgent) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/targets/%s.json", data.Login), dataBytes, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

// targetLoadFile Returns target data from file
func targetLoadFile(id string, agent discordAgent) (TargetData, error) {
	target := TargetData{}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/targets/%s.json", id))
	if err != nil {
		logErrorToChan(agent, err)
		return TargetData{}, err
	}

	err = json.Unmarshal(fileData, &target)
	if err != nil {
		logErrorToChan(agent, err)
		return TargetData{}, err
	}

	return target, nil
}
