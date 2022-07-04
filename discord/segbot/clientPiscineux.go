package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Piscineux []struct {
	Login         string    `json:"login"`
	LastName      string    `json:"last_name"`
	UsualFullName string    `json:"usual_full_name"`
	Url           string    `json:"url"`
	Displayname   string    `json:"displayname"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (s *APIClient) GetPiscineux(id, month, year, page string) (Piscineux, error) {
	endpoint := fmt.Sprintf("%s/v2/cursus/%s/users?filter[primary_campus_id]=1&filter[pool_year]=%s&filer[pool_month]=%s&per_page=100&page=%s", s.Url, id, year, month, page)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Piscineux{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return Piscineux{}, err
	}
	if resp.StatusCode != 200 {
		return Piscineux{}, fmt.Errorf(resp.Status)
	}
	body, err := ReadHTTPResponse(resp)
	if err != nil {
		return Piscineux{}, err
	}
	var piscineux Piscineux
	_ = json.Unmarshal(body, &piscineux)
	return piscineux, nil
}
