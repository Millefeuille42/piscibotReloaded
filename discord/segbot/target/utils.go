package target

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"os"
	"piscibotReloaded/discord/segbot/discord"
	"piscibotReloaded/discord/segbot/utils"
)

// makeApiReq Internal, Make calls to the 42API module to start data collecting and check if user exists
func makeApiReq(path, login string, agent discord.Agent) error {
	uri := fmt.Sprintf("http://%s:%s/user/%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"), login)
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		discord.LogErrorToChan(agent, err)
		return err
	}
	if res.StatusCode == 404 {
		_, _ = agent.Session.ChannelMessageSend(agent.Channel, "This login doesn't exist")
		_ = os.Remove(path)
		return os.ErrNotExist
	}
	return nil
}

// loadOrCreate Internal, Loads or creates Target file
func loadOrCreate(path, login string, settings *Target, message *discordgo.MessageCreate) error {
	exists, err := utils.CreateFileIfNotExist(path)
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
			_ = os.Remove(path)
			return loadOrCreate(path, login, settings, message)
		}
	} else {
		*settings = Target{
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
