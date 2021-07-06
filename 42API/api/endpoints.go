package main

import (
	"encoding/json"
	"fmt"
	apiclient "github.com/BoyerDamien/42APIClient"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"strings"
)

func endpointsLogin() {
	App.Post("/user/:login", func(c *fiber.Ctx) error {
		exist, err := createFileIfNotExist("./data/" + c.Params("login") + ".json")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		if exist {
			message := fmt.Sprintf("User %s already exists", c.Params("login"))
			return c.Status(fiber.ErrBadRequest.Code).SendString(message)
		}
		user, err := Client.GetUser(c.Params("login"))
		if err != nil {
			return c.Status(404).SendString(err.Error())
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
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
		files, err := ioutil.ReadDir("./data/")
		if err != nil {
			logError(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		for _, f := range files {
			err = os.Remove("./data/" + f.Name())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}
		return c.SendString("Users have been deleted")
	})

	App.Get("/users", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir("./data/")
		if err != nil {
			logError(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		response := make([]apiclient.User, 0)
		for _, f := range files {
			var user = apiclient.User{}

			fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/%s", f.Name()))
			if !strings.HasSuffix(f.Name(), ".json") {
				continue
			}
			if err != nil {
				fmt.Println(f.Name())
				fmt.Println(string(fileData))
				logError(err)
				continue
			}
			err = json.Unmarshal(fileData, &user)
			if err != nil {
				fmt.Println(f.Name())
				fmt.Println(string(fileData))
				logError(err)
				continue
			}
			response = append(response, user)
		}
		return c.JSON(response)
	})
}
