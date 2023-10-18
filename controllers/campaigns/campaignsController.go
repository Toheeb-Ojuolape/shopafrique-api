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

func CreateCampaign(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return handleErrors.HandleBadRequest(c, "Invalid user ID")
	}
	var req types.Campaign
	c.BodyParser(&req)

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while creating campaign")
	}

	campaign := models.Campaign{
		ID:          newUUID,
		Title:       req.Title,
		Description: req.Description,
		Media:       req.Media,
		MediaType:   req.MediaType,
		Budget:      req.Budget,
		Status:      req.Status,
		Views:       0,
		Clicks:      0,
		Impressions: 0,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Objective:   req.Objective,
		Audience:    req.Audience,
		UserId:      id,
	}

	// create a transaction record of the funding based on the info from request
	if err := initializers.DB.Create(&campaign).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while creating campaign")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Campaign created successfully",
		Data: map[string]interface{}{
			"title":       req.Title,
			"budget":      req.Budget,
			"description": req.Description,
		},
	})
}
