package main

import (
	"os"

	apiclient "github.com/BoyerDamien/42APIClient"
	"github.com/gofiber/fiber/v2"
)

var (
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
