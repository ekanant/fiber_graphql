package main

import (
	"fiber_graphql/graph"
	"fiber_graphql/graph/generated"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
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

	h1 := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	app.Post("/graphql", adaptor.HTTPHandlerFunc(h1.ServeHTTP))

	h := playground.Handler("GraphQL", "/graphql")
	app.Get("/graphql/playground", adaptor.HTTPHandlerFunc(h.ServeHTTP))

	log.Fatalln(app.Listen(":3000"))
}
