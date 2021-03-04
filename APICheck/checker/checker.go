package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	apiclient "github.com/BoyerDamien/42APIClient"
)

// Checker implements checker interface
type Checker struct {
	UserList   []apiclient.User
	UserAPIURL string
}

// FetchUsers retrieve 42 users stored in user database
func (s *Checker) FetchUsers() error {
	response, err := http.Get(fmt.Sprintf("%s/users", s.UserAPIURL))
	if err != nil {
		return err
	}
	body, err := apiclient.ReadHTTPResponse(response)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &s.UserList)
	return nil
}

// Length returns the number of fetched users
func (s *Checker) Length() int {
	return len(s.UserList)
}

// Check does all checks between old and new user data
func (s *Checker) Check(dbUser *apiclient.User, apiUser *apiclient.User) []string {
	var messages []string

	if err := CheckProjectSubscribed(dbUser, apiUser); err != nil {
		messages = append(messages, err.Error())
	}
	if err := CheckUserLocation(dbUser, apiUser); err != nil {
		messages = append(messages, err.Error())
	}

	dbUserProjectsLen := len(dbUser.ProjectsUsers)
	for i := 0; i < dbUserProjectsLen; i++ {
		p1 := BuildProject(dbUser.ProjectsUsers[i])
		if p1.Validated {
			return messages
		}
		p2Index := sort.Search(len(apiUser.ProjectsUsers), func(i int) bool {
			return BuildProject(apiUser.ProjectsUsers[i]).Name == p1.Name
		})
		if p2Index < len(apiUser.ProjectsUsers) {
			p2 := BuildProject(apiUser.ProjectsUsers[p2Index])
			if err := CheckProjectStatus(dbUser.Login, &p1, &p2); err != nil {
				messages = append(messages, err.Error())
			}
		}
	}
	return messages
}

// UpdateDB updates user data in database
func (s *Checker) UpdateDB(apiUser *apiclient.User) {
	body, err := json.Marshal(apiUser)
	if err != nil {
		log.Fatal(err.Error())
	}
	apiURL := fmt.Sprintf("%s/user/%s", s.UserAPIURL, apiUser.Login)
	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status code: %d\n", resp.StatusCode)
	}
}
