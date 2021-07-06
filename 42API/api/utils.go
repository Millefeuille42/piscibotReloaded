package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

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
	for _, val := range envVariables {
		if os.Getenv(val) == "" {
			log.Fatalf("Missing %s env variable", val)
		}
		log.Printf("%s env variable [OK]", val)
	}
}

// Find Check if val exists in Slice, true if it exists
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// logError Prints error + StackTrace to stderr if error
func logError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err, string(debug.Stack()))
	}
}

// checkError Panic if error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// createDirIfNotExist Check if dir exists, if not create it
func createDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

// createFileIfNotExist Check if file exists, if not create it
func createFileIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			logError(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}
