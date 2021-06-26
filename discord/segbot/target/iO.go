package target

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"piscibotReloaded/discord/segbot/discord"
)

// Write Writes target data to file
func Write(data Target, agent discord.Agent) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/targets/%s.json", data.Login), dataBytes, 0677)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	return nil
}

// Load Returns target data from file
func Load(id string, agent discord.Agent) (Target, error) {
	target := Target{}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/targets/%s.json", id))
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return Target{}, err
	}

	err = json.Unmarshal(fileData, &target)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return Target{}, err
	}

	return target, nil
}
