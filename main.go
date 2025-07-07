package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/config"
	"github.com/koko-linoo/go-fiber-api/routes"
)

func init() {
	config.InitDatabase()
}

func main() {
	app := fiber.New()

	routes.InitRoutes(app)

	app.Listen(":3000")

	log.Println("Server started on port :3000")
}
