package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

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

func CompareLocation(apiUser apiclient.User, dbUser apiclient.User) (string, error) {
	if apiUser.Location != dbUser.Location {
		if apiUser.Location == nil {
			return fmt.Sprintf("%s s'est déconnecté", apiUser.Login), nil
		} else if apiUser.Location != nil {
			return fmt.Sprintf("%s s'est connecté en %s", apiUser.Login, apiUser.Location), nil
		}
	}
	return string(""), fmt.Errorf("No changes for user %s", apiUser.Login)
}

/*
func CompareProjects(apiUser apiclient.User, dbUser apiclient.User) (string, error){
	if apiUser.ProjectsUsers
}*/

func main() {
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatal(err.Error())
	}

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
			for _, val := range checker.UserList {
				apiUser, err := api.GetUser(val.Login)
				if err != nil {
					log.Printf("Error: Cannot fetch user %s data", val.Login)
				} else {
					msg, err := CompareLocation(apiUser, val)
					if err != nil {
						log.Println(err.Error())
					} else {
						log.Println(msg)
					}
				}
				time.Sleep(time.Second * 3)
			}
		} else {
			log.Println("No tracked users")
			time.Sleep(time.Second * 3)
		}
	}
}
