package controllers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/models"
	"github.com/koko-linoo/go-fiber-api/services"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(&user); err != nil {
		errs := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, e := range errs {
			messages = append(messages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}
		return c.Status(400).JSON(fiber.Map{"errors": messages})
	}

	newUser, err := services.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(newUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(&user); err != nil {
		errs := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, e := range errs {
			messages = append(messages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}
		return c.Status(400).JSON(fiber.Map{"errors": messages})
	}

	updatedUser, err := services.UpdateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(updatedUser)
}

func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")

	user, err := services.GetUser(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	userID := uint(idInt)

	err = services.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
