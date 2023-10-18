package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func FetchTransactions(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return handleErrors.HandleBadRequest(c, "Invalid user ID")
	}

	var transactions []models.Transaction

	err := initializers.DB.Where("user_id = ?", id).Find(&transactions).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Error fetching transactions")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Transactions fetched successfully",
		Data:    transactions,
	})
}
