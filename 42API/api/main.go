package main

import (
	"fmt"
	"log"
	"time"

	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	CheckEnvVariables()

	if err := Client.Auth(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("42 API authentication [OK]")
	log.Printf("MongoUrl: %s\n", MongoURL)
	log.Printf("Database name: %s\n", DatabaseName)

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

	App.Post("/user/:login", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			user, err := Client.GetUser(c.Params("login"))
			if err != nil {
				return c.Status(404).SendString(err.Error())
			}
			if IsExists(db, &user) {
				message := fmt.Sprintf("User %s already exists", c.Params("login"))
				return c.Status(fiber.ErrBadRequest.Code).SendString(message)
			}
			_, err = db.InsertOne(DatabaseName, user)
			if err != nil {
				return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
			}
			return c.SendString(fmt.Sprintf("User %s successfully created", c.Params("login")))
		})
	})

	App.Put("/user/:login", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			var user apiclient.User
			if err := c.BodyParser(&user); err != nil {
				_ = c.Status(fiber.StatusBadRequest).SendString(err.Error())
			}
			result, err := db.UpdateOne(DatabaseName, bson.M{"login": user.Login}, bson.M{"$set": user})
			if err != nil || result.ModifiedCount != 1 {
				return c.SendStatus(fiber.ErrBadRequest.Code)
			}
			return c.JSON(user)
		})
	})

	App.Delete("/user/:login", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			result, err := db.DeleteOne(DatabaseName, bson.M{"login": c.Params("login")})
			if err != nil || result.DeletedCount != 1 {
				return c.Status(404).SendString(fmt.Sprintf("User %s not found", c.Params("login")))
			}
			return c.SendString(fmt.Sprintf("User %s successfully deleted", c.Params("login")))
		})
	})

	App.Delete("/users", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			result, err := db.DeleteMany(DatabaseName, bson.M{})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			if result.DeletedCount == 0 {
				return c.SendStatus(404)
			}
			return c.SendString(fmt.Sprintf("%d users have been deleted", result.DeletedCount))
		})
	})

	App.Get("/user/:login", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			user, err := db.FindOne(DatabaseName, bson.M{"login": c.Params("login")})
			if err != nil || user.Err() != nil {
				return c.Status(404).SendString(fmt.Sprintf("User %s not found", c.Params("login")))
			}
			var response apiclient.User
			if err := user.Decode(&response); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			return c.JSON(response)
		})
	})

	App.Get("/users", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			users, err := db.FindMany(DatabaseName, bson.M{})
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			var response []apiclient.User
			if err := users.All(db.GetContext(), &response); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			return c.JSON(response)
		})
	})

	log.Fatal(App.Listen(fmt.Sprintf(":%s", AppPort)))
}
