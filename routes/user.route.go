package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/controllers"
)

func InitRoutes(app fiber.Router) {
	user := app.Group("/users")

	user.Post("/", controllers.CreateUser)
	user.Put("/:id", controllers.UpdateUser)
	user.Get("/:username", controllers.GetUser)
	user.Get("/", controllers.GetAllUsers)
	user.Delete("/:id", controllers.DeleteUser)

	log.Println("Registered /user route")
}
