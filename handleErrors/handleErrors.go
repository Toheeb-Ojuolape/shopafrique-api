package handleErrors

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleBadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": msg,
	})
}
