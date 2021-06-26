package guild

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"piscibotReloaded/discord/segbot/discord"
)

// Write Writes guild data to file
func Write(agent discord.Agent, data Guild) error {
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/guilds/%s.json", data.GuildID), jsonGuild, 0677)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	return nil
}

// Load Returns guild data from file
func Load(agent discord.Agent, silent bool, id string) (Guild, error) {
	data := Guild{}

	if id == "" {
		id = agent.Message.GuildID
	}
	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/guilds/%s.json", id))
	if err != nil {
		if !silent {
			discord.LogErrorToChan(agent, err)
		}
		return Guild{}, err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		if !silent {
			discord.LogErrorToChan(agent, err)
		}
		return Guild{}, err
	}

	return data, nil
}
