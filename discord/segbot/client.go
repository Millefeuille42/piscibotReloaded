package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// AuthToken struct defines 42 api oAuth Token response
type AuthToken struct {
	AccessToken string    `json:"access_Token"`
	TokenType   string    `json:"Token_type"`
	ExpiresIn   int       `json:"expires_in"`
	LastUpdate  time.Time `json:"last_update"`
}

// API 42 interface
type API interface {
	Auth() error
	GetPiscineux(login string) (Piscineux, error)
	Token() AuthToken
}

// APIClient implements 42 API interface
type APIClient struct {
	Url    string
	Uid    string
	Secret string
	token  AuthToken
}

// Auth method implements 42 API oAuth authentication
func (s *APIClient) Auth() error {
	endpoint := BuildAuthURL(s.Url, s.Uid, s.Secret)
	response, err := http.Post(endpoint, "application/x-www-form-Urlencoded", nil)
	if err != nil {
		return err
	}
	body, err := ReadHTTPResponse(response)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(body, &s.token)
	s.token.LastUpdate = time.Now()
	return nil
}

// Token returns 42 api auth token
func (s *APIClient) Token() AuthToken {
	return s.token
}
