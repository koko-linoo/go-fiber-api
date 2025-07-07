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
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		},
	})

	api := app.Group("/api")

	routes.InitRoutes(api)

	app.Listen(":3000")

	log.Println("Server started on port :3000")
}
