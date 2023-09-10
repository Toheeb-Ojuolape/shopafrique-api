package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func UserController(c *fiber.Ctx) error {
	id := c.Get("id")

	var user models.User
	initializers.DB.First(&user, id)

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
