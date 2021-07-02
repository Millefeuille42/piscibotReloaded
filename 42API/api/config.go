package main

import (
	"os"

	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DatabaseName name of database
	DatabaseName string = os.Getenv("DB_NAME")

	// DBCredentials db credentials
	DBCredentials options.Credential = options.Credential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	// MongoURL url of mongodb
	MongoURL string = os.Getenv("DB_URL")

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
