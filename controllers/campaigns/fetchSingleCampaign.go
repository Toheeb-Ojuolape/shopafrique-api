package controllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
)

func FetchSingleCampaign(c *fiber.Ctx) error {
	campaignID := c.Params("id")
	var campaign models.Campaign

	err := initializers.DB.Where("id = ?", campaignID).Find(&campaign).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Error fetching campaigns")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Single campaign fetched successfully",
		Data:    campaign,
	})

}
