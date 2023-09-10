package authControllers

import (
	"net/http"

	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Email        string
	Password     string
	Country      string
	FirstName    string
	LastName     string
	BusinessType string
	BusinessName string
	ProcessId    string
}

func Signup(c *fiber.Ctx) error {

	var request SignupRequest

	c.BodyParser(&request)
	missingProps := helpers.ValidateRequest(request)

	if missingProps != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid parameters passed",
			"error":   "Request is missing " + missingProps,
		})
	}

	// hash the password

	// validate processId

	//create user in database

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Signup is ready",
		"data":    request,
	})
}
