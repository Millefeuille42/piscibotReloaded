package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	apiclient "github.com/BoyerDamien/42APIClient"
	"log"
	"net/http"
)

// Checker model
type Checker struct {
	UserList   []apiclient.User
	UserAPIURL string
}

// Message model for discord API
type Message struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
	Cursus  string `json:"cursus"`
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
	_ = json.Unmarshal(body, &s.UserList)
	return nil
}

// Length returns the number of fetched users
func (s *Checker) Length() int {
	return len(s.UserList)
}

// Check does all checks between old and new user data
func (_ *Checker) Check(dbUser, apiUser *apiclient.User) []Message {
	var messages []Message

	if err := CheckProjectSubscribed(dbUser, apiUser); err != nil {
		message := Message{Message: err.Error(), Channel: "started", Login: dbUser.Login}
		messages = append(messages, message)
	}
	if err := CheckUserLocation(dbUser, apiUser); err != nil {
		message := Message{Message: err.Error(), Channel: "location", Login: dbUser.Login}
		messages = append(messages, message)
	}

	if err := CheckUserStudent(dbUser, apiUser); err != nil {
		message := Message{Message: err.Error(), Channel: "student", Login: dbUser.Login}
		messages = append(messages, message)
	}

	dbUserProjectsLen := len(dbUser.ProjectsUsers)
	for i := 0; i < dbUserProjectsLen; i++ {
		p1 := BuildProject(dbUser.ProjectsUsers[i])
		//		if p1.Validated { 				// Projects are not sorted,
		//			return messages				// this conditon breaks the validated events
		//		}
		for _, val := range apiUser.ProjectsUsers {
			p2 := BuildProject(val)
			if p2.Slug == p1.Slug {
				if err := CheckProjectStatus(dbUser.Login, &p1, &p2); err != nil {
					message := Message{Message: err.Error(), Channel: "success", Login: dbUser.Login}
					if val["cursus_ids"] != nil && len(val["cursus_ids"].([]interface{})) > 0 { // To be done a better way, ex. by grabbing cursus slug with API
						if val["cursus_ids"].([]interface{})[0].(float64) == 21 {
							message.Cursus = "42cursus"
						}
					}
					messages = append(messages, message)
				}
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

// Send sends a list of messages to discord API
func (_ *Checker) Send(apiURL string, messages []Message) error {
	url := fmt.Sprintf("%s/discord", apiURL)
	body, err := json.Marshal(messages)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: Cannot send %v to %s", messages, apiURL)
	}
	return nil
}
