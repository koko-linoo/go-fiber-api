package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koko-linoo/go-fiber-api/helpers"
	"github.com/koko-linoo/go-fiber-api/models"
	"github.com/koko-linoo/go-fiber-api/services"
)

func CreateUser(c *fiber.Ctx) error {
	var createUser models.UserCreate
	if err := c.BodyParser(&createUser); err != nil {
		return fiber.ErrBadRequest
	}

	if err := helpers.Validate(createUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"errors":  err,
			"status":  "error",
		})
	}

	usr, err := services.CreateUser(createUser.ToUser())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "User Create Failed",
			"message": "Bad Request",
			"status":  "error",
		})
	}

	return c.JSON(usr)
}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return fiber.ErrBadRequest
	}

	var updateUser models.UserUpdate
	if err := c.BodyParser(&updateUser); err != nil {
		return fiber.ErrBadRequest
	}

	if err := helpers.Validate(updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"errors":  err,
			"status":  "error",
		})
	}

	usr, err := services.UpdateUser(idInt, updateUser.GetMap())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "User Update Failed",
			"message": "Bad Request",
			"status":  "error",
		})
	}

	return c.JSON(usr)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := services.GetUserByID(uint(idInt))
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
		return fiber.ErrBadRequest
	}

	err = services.DeleteUser(uint(idInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
