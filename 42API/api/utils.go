package main

import (
	"log"

	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

// Exec runs endpoint functions
func Exec(c *fiber.Ctx, f func(mw.Database, *fiber.Ctx) error) error {
	if err := Wrapper.Init(MongoURL); err != nil {
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
