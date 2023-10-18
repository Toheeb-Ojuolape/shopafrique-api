package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func FetchCampaigns(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return handleErrors.HandleBadRequest(c, "Invalid user ID")
	}

	var campaigns []models.Campaign

	err := initializers.DB.Where("user_id = ?", id).Find(&campaigns).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Error fetching campaigns")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Campaigns fetched successfully",
		Data:    campaigns,
	})
}
