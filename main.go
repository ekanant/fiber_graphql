package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

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
