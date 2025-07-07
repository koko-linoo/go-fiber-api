package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/models"
	"github.com/koko-linoo/go-fiber-api/services"
)

func Login(c *fiber.Ctx) error {

	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := services.Login(loginRequest.Username, loginRequest.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(user)
}
