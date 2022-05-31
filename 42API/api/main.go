package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	CheckEnvVariables()
	err := createDirIfNotExist("./data")
	if err != nil {
		return
	}

	if err := Client.Auth(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("42 API authentication [OK]")

	// Logger middleware
	App.Use(logger.New())

	// Auth token refresh middleware
	App.Use(func(c *fiber.Ctx) error {
		token := Client.Token()
		timestamp := time.Since(token.LastUpdate).Seconds()
		if timestamp > float64(token.ExpiresIn) {
			err := Client.Auth()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}
		return c.Next()
	})

	endpointsLogin()
	endpointsUsers()

	log.Fatal(App.Listen(fmt.Sprintf(":%s", AppPort)))
}
