package main

import (
	"fmt"
	"os"

	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"github.com/gofiber/fiber/v2"
)

var (
	// DatabaseName name of database
	DatabaseName string = os.Getenv("DB_NAME")

	// MongoURL url of mongodb
	MongoURL string = fmt.Sprintf("mongodb://%s", os.Getenv("DB_URL"))

	// Wrapper wrapper object
	Wrapper mw.Wrapper = &mw.WrapperData{}

	url    string = "https://api.intra.42.fr"
	uid    string = os.Getenv("UID")
	secret string = os.Getenv("SECRET")

	// Client 42 API client
	Client apiclient.API = &apiclient.APIClient{Url: url, Uid: uid, Secret: secret}

	// App app
	App *fiber.App = fiber.New()

	// AppPort port
	AppPort string = os.Getenv("PORT")
)
