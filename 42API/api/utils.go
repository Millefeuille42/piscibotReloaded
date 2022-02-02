package main

import (
	"log"
	"os"

	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

// Exec runs endpoint functions
func Exec(c *fiber.Ctx, f func(mw.Database, *fiber.Ctx) error) error {
	if err := Wrapper.Init(MongoURL, DBCredentials); err != nil {
		log.Println(err.Error())
	}
	defer Wrapper.Close()
	db := Wrapper.GetDatabase(DatabaseName)
	db.AddCollections(DatabaseName)
	return f(db, c)
}

// IsExists checks if a user already exists in data base
func IsExists(db mw.Database, user *apiclient.User) bool {
	res, err := db.FindOne(DatabaseName, bson.M{"login": user.Login})
	if err != nil || res.Err() != nil {
		return false
	}
	return true
}

// CheckEnvVariables tests the existence of required env variables
func CheckEnvVariables() {
	envVariables := []string{"DB_NAME", "DB_URL", "UID", "SECRET", "PORT", "DB_USERNAME", "DB_PASSWORD"}
	log.Println("Checking:", envVariables)
	for _, val := range envVariables {
		if os.Getenv(val) == "" {
			log.Fatalf("Missing %s env variable", val)
		}
		log.Printf("%s env variable [OK]", val)
	}
}
