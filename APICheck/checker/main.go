package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	apiclient "github.com/BoyerDamien/42APIClient"
)

type Checker struct {
	UserList []apiclient.User
}

func main() {
	//url := "https://api.intra.42.fr"
	//api := &apiclient.APIClient{Url: url, Uid: uid, Secret: secret}

	response, err := http.Get("http://localhost:3000/users")
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := apiclient.ReadHTTPResponse(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	var result []apiclient.User
	json.Unmarshal(body, &result)

	fmt.Println(result)
}
