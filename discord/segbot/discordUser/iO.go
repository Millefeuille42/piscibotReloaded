package discordUser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"piscibotReloaded/discord/segbot/discord"
)

// Write Writes user data to file
func Write(data User, agent discord.Agent, id string) error {
	if id == "" {
		id = agent.Message.Author.ID
	}
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/users/%s.json", id), dataBytes, 0677)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	return nil
}

// Load Returns user data from file
func Load(id string, agent discord.Agent) (User, error) {
	user := User{}
	if id == "" {
		id = agent.Message.Author.ID
	}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/users/%s.json", id))
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return User{}, err
	}

	err = json.Unmarshal(fileData, &user)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return User{}, err
	}

	return user, nil
}
