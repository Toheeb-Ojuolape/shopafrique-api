package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func UserController(c *fiber.Ctx) error {
	user_id, ok := c.Locals("user_id").(string)
	if !ok {
		return handleErrors.HandleBadRequest(c, "Invalid user ID")
	}

	var user models.User
	err := initializers.DB.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "User details fetched successfully",
		Data: map[string]interface{}{
			"firstName":    user.FirstName,
			"lastName":     user.LastName,
			"email":        user.Email,
			"balance":      user.Balance,
			"country":      user.Country,
			"businessName": user.BusinessName,
			"businessType": user.BusinessType,
		},
	})
}
