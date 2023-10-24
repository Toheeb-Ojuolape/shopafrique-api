package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateSingleCampaign(c *fiber.Ctx) error {
	campaignID := c.Params("id")
	var req models.Campaign
	var campaign models.Campaign

	c.BodyParser(&req)
	if err := initializers.DB.Where("id = ?", campaignID).Find(&campaign).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Error fetching campaign")
	}

	campaign.Views += req.Views
	campaign.Clicks += req.Clicks

	if err := initializers.DB.Save(&campaign).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Error updating campaign")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Campaign views and clicks updated successfully",
	})

}
