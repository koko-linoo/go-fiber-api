package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app fiber.Router) {

	api := app.Group("/api")

	InitUserRoutes(api)
	InitAuthRoutes(api)

	app.All("/*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
			"status":  "error",
			"error":   "Request not found",
		})
	})

	log.Println("Initialized routes")
}
