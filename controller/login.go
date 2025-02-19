package controller

import (
	"github.com/dancankarani/delha-frontend/model"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminLogin(c *fiber.Ctx) error {
	var loginReq LoginRequest

	// Parse the JSON request body
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Check credentials against the hardcoded admin
	if loginReq.Username == model.HardcodedAdmin.Username && loginReq.Password == model.HardcodedAdmin.Password {
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"admin":   loginReq.Username,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid username or password",
	})
}
