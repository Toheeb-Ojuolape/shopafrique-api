package handleSuccess

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Message string
	Data    interface{}
}

func HandleSuccessResponse(c *fiber.Ctx, response SuccessResponse) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": response.Message,
		"data":    response.Data,
	})
}
