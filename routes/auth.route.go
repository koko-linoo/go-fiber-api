package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/controllers"
)

func InitAuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/login", controllers.Login)

	log.Println("Registered /auth route")
}
