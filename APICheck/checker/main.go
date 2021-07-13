package main

import (
	"log"
	"os"
	"time"

	apiclient "github.com/BoyerDamien/42APIClient"
)

func checkEnvVariables() {
	envVariables := []string{"USER_API_URL", "UID", "SECRET", "SEGBOT_URL"}
	for _, val := range envVariables {
		if os.Getenv(val) == "" {
			log.Fatalf("Missing %s env variable", val)
		}
		log.Printf("%s env variable [OK]", val)
	}
}

func checkAuth(api *apiclient.APIClient) error {
	token := api.Token()
	timestamp := time.Since(token.LastUpdate).Seconds()
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
	checker := Checker{UserAPIURL: os.Getenv("USER_API_URL")}
	segbotURL := os.Getenv("SEGBOT_URL")

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

			for i := 0; i < checker.Length(); i++ {
				apiUser, err := api.GetUser(checker.UserList[i].Login)
				if err != nil {
					log.Fatalf("Error: Cannot fetch user %s data\n", checker.UserList[i].Login)
				}
				checker.UpdateDB(&apiUser)
				time.Sleep(time.Millisecond * 250)
				messages := checker.Check(&checker.UserList[i], &apiUser)
				if len(messages) > 0 {
					if err := checker.Send(segbotURL, messages); err != nil {
						log.Println(err.Error())
					} else {
						log.Printf("Messages sent for %s\n", apiUser.Login)
					}
				} else {
					log.Printf("No messages sent for %s\n", apiUser.Login)
				}
				time.Sleep(time.Second * 3)
			}
		} else {
			log.Println("No tracked users")
			time.Sleep(time.Second * 3)
		}
	}
}
