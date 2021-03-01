package main

import (
	"fmt"
	"log"
	"os"
	"time"

	apiclient "github.com/BoyerDamien/42APIClient"
)

func checkEnvVariables() {
	envVariables := []string{"USER_API_URL", "UID", "SECRET"}
	for _, val := range envVariables {
		if os.Getenv(val) == "" {
			log.Fatalf("Missing %s env variable", val)
		}
		log.Printf("%s env variable [OK]", val)
	}
}

func checkAuth(api *apiclient.APIClient) error {
	token := api.Token()
	timestamp := time.Now().Sub(token.LastUpdate).Seconds()
	if timestamp > float64(token.ExpiresIn) {
		log.Println("Refreshing auth token...")
		if err := api.Auth(); err != nil {
			return err
		}
	}
	log.Println("Auth token [OK]")
	return nil
}

func main() {
	checkEnvVariables()

	url := "https://api.intra.42.fr"
	api := &apiclient.APIClient{Url: url, Uid: os.Getenv("UID"), Secret: os.Getenv("SECRET")}
	checker := &CheckerImpl{UserAPIURL: os.Getenv("USER_API_URL")}

	if err := api.Auth(); err != nil {
		log.Fatal(err.Error())
	}

	for {
		if err := checkAuth(api); err != nil {
			log.Fatal(err.Error())
		}

		if err := checker.FetchUsers(); err != nil {
			log.Fatal(err.Error())
		}

		if checker.Length() > 0 {
			log.Printf("%d users detected\nAnalysis begin...", checker.Length())
			for _, val := range checker.UserList() {
				fmt.Println(val)
				time.Sleep(time.Second * 3)
			}
		}
	}
}