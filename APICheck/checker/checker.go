package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	apiclient "github.com/BoyerDamien/42APIClient"
)

// Checker interface
type Checker interface {
	UserList() []apiclient.User
	FetchUsers() error
	Length() int
}

// CheckerImpl implements checker interface
type CheckerImpl struct {
	userList   []apiclient.User
	UserAPIURL string
}

// FetchUsers retrieve 42 users stored in user database
func (s *CheckerImpl) FetchUsers() error {
	response, err := http.Get(fmt.Sprintf("%s/users", s.UserAPIURL))
	if err != nil {
		return err
	}
	body, err := apiclient.ReadHTTPResponse(response)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &s.userList)
	return nil
}

// UserList returns a list of fetched users
func (s *CheckerImpl) UserList() []apiclient.User {
	return s.userList
}

// Length returns the number of fetched users
func (s *CheckerImpl) Length() int {
	return len(s.UserList())
}
