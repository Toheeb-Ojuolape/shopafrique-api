package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/Toheeb-Ojuolape/shopafrique-api/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func FundWallet(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return handleErrors.HandleBadRequest(c, "Invalid user ID")
	}
	var user models.User
	var req types.Transaction
	c.BodyParser(&req)
	// Retrieve the user by ID
	err := initializers.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Error querying user")
	}

	if req.Amount <= 0 {
		return handleErrors.HandleBadRequest(c, "Amount cannot be 0 or negative. Suspected fraud!")
	}

	// Update the user's balance
	user.Balance += req.Amount

	// Save the updated user with the new balance
	if err := initializers.DB.Save(&user).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while processing transaction")
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while processing transaction")
	}

	transaction := models.Transaction{
		UserID:        id,
		Amount:        req.Amount,
		ID:            newUUID,
		Type:          req.Type,
		Status:        "completed",
		CustomerEmail: user.Email,
		CustomerName:  user.FirstName,
		PaymentMethod: req.PaymentMethod,
	}

	// create a transaction record of the funding based on the info from request
	if err := initializers.DB.Create(&transaction).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while processing transaction")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Wallet funded successfully",
		Data: map[string]interface{}{
			"method": req.PaymentMethod,
			"type":   req.Type,
			"amount": req.Amount,
		},
	})
}
