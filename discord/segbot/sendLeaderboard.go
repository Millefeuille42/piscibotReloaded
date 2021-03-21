package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

// slugs: c-piscine

type targetLevelPair struct {
	name  string
	level float64
}

type ApiData struct {
	CursusUsers []struct {
		ID       int         `json:"id"`
		Grade    interface{} `json:"grade"`
		Level    float64     `json:"level"`
		CursusID int         `json:"cursus_id"`
		Cursus   struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"cursus"`
	} `json:"cursus_users"`
}

func targetGetData(agent discordAgent, target string) (ApiData, error) {
	apiData := ApiData{}

	uri := fmt.Sprintf("http://%s:%s/user/%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"), target)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		logErrorToChan(agent, err)
		return apiData, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logErrorToChan(agent, err)
		return apiData, err
	}
	if res.StatusCode == 404 {
		fmt.Println("Not Found")
		return apiData, os.ErrNotExist
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logErrorToChan(agent, err)
		return apiData, err
	}
	err = json.Unmarshal(data, &apiData)
	if err != nil {
		logErrorToChan(agent, err)
		return apiData, err
	}
	return apiData, nil
}

func createLevelPairList(agent discordAgent, slug string) []targetLevelPair {
	var pairList = make([]targetLevelPair, 0)

	targetList, err := getTargetsOfGuild(agent, agent.message.GuildID)
	if err != nil {
		return nil
	}
	for _, target := range targetList {
		apiData, err := targetGetData(agent, target)
		if err != nil {
			return nil
		}
		for _, cursus := range apiData.CursusUsers {
			if cursus.Cursus.Slug == slug {
				pairList = append(pairList, targetLevelPair{name: target, level: cursus.Level})
			}
		}
	}
	return pairList
}

func createLeaderboard(agent discordAgent, slug string) string {
	leaderBoard := "\t--- c-piscine ---\n"
	pairList := createLevelPairList(agent, slug)
	if pairList == nil {
		return ""
	}

	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].level > pairList[j].level
	})

	for i, pair := range pairList {
		leaderBoard = fmt.Sprintf("%s\n%d. %-15.15s :  %.2f", leaderBoard, i+1, pair.name, pair.level)
	}
	return leaderBoard
}

func sendLeaderboard(agent discordAgent) {
	if !userInitialCheck(agent) {
		return
	}
	args := strings.Split(agent.message.Content, " ")
	if len(args) != 2 {
		sendMessageWithMention("Invalid Number of Arguments", "", agent)
		return
	}
	leaderboard := createLeaderboard(agent, args[1])
	if leaderboard != "" {
		sendMessageWithMention("```"+leaderboard+"```", "", agent)
	}
}
