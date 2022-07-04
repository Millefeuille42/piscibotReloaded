package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReadHTTPResponse read and return an httpresponse body
func ReadHTTPResponse(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	return body, nil
}

// BuildAuthURL Build the 42 endpoints based on the base url, uid and secret
func BuildAuthURL(base string, uid string, secret string) string {
	return fmt.Sprintf("%s/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", base, uid, secret)
}
