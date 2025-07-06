package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/config"
	"github.com/koko-linoo/go-fiber-api/routes"
)

func main() {
	app := fiber.New()

	config.InitDatabase()

	api := app.Group("/api")

	routes.InitRoutes(api)

	app.Listen(":3000")
}
