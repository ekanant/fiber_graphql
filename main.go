package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "pong",
		})
	})

	log.Fatalln(app.Listen(":3000"))
}
