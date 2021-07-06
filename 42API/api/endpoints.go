package main

import (
	"fmt"
	apiclient "github.com/BoyerDamien/42APIClient"
	mw "github.com/BoyerDamien/mongodbWrapper"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func endpointsLogin() {
	App.Post("/user/:login", func(c *fiber.Ctx) error {
		user, err := Client.GetUser(c.Params("login"))
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		exist, err := createFileIfNotExist("./data/" + user.Login + ".json")
		if err != nil {
			return err
		}
		if exist {
			message := fmt.Sprintf("User %s already exists", c.Params("login"))
			return c.Status(fiber.ErrBadRequest.Code).SendString(message)
		}
		err = writeUserData(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("User %s successfully created", c.Params("login")))
	})

	App.Put("/user/:login", func(c *fiber.Ctx) error {
		var user apiclient.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		_, err := os.Stat(user.Login)
		if os.IsNotExist(err) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		err = writeUserData(user)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(user)
	})

	App.Delete("/user/:login", func(c *fiber.Ctx) error {
		_, err := os.Stat(c.Params("login"))
		if os.IsNotExist(err) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		err = os.Remove("./data/" + c.Params("login") + ".json")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("User %s successfully deleted", c.Params("login")))
	})

	App.Get("/user/:login", func(c *fiber.Ctx) error {
		_, err := os.Stat(c.Params("login"))
		if os.IsNotExist(err) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		resp, err := readUserData(c.Params("login"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(resp)
	})
}

func endpointsUsers() {
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

	App.Get("/users", func(c *fiber.Ctx) error {
		return Exec(c, func(db mw.Database, c *fiber.Ctx) error {
			users, err := db.FindMany(DatabaseName, bson.M{})
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			var response []apiclient.User
			if err = users.All(db.GetContext(), &response); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			return c.JSON(response)
		})
	})
}
