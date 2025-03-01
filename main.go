package main

import (
	"github.com/gofiber/fiber/v2"
)

type Welcome struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New(fiber.Config{AppName: "fiber-demo"})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(Welcome{Message: "Hello World"})
	})

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
