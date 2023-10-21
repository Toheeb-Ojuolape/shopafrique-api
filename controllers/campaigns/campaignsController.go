package controllers

import (
	"strconv"

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

	var user models.User
	var req types.Campaign
	c.BodyParser(&req)

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while creating campaign")
	}

	//debit wallet of funding dedicated to budget

	//fetch user
	err = initializers.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Error querying user")
	}

	budget, err := strconv.ParseFloat(req.Budget, 64)

	if err != nil {
		return handleErrors.HandleBadRequest(c, "Oops!. Budget cannot be 0 or negative.")
	}

	if user.Balance < budget {
		return handleErrors.HandleBadRequest(c, "Insufficient funds, kindly top-up your wallet")
	}

	// Update the user's balance
	user.Balance -= budget

	// Save the updated user with the new balance
	if err := initializers.DB.Save(&user).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while processing transaction")
	}

	//throw error if funds insufficient

	campaign := models.Campaign{
		ID:          newUUID,
		Title:       req.Title,
		Description: req.Description,
		Media:       req.Media,
		MediaType:   req.MediaType,
		Budget:      budget,
		Status:      "active",
		Views:       0,
		Clicks:      0,
		Impressions: 0,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Objective:   req.Objective,
		Audience:    req.Audience,
		UserId:      id,
	}

	transaction := models.Transaction{
		UserID:        id,
		Amount:        budget,
		ID:            newUUID,
		Type:          "create-campaign",
		Status:        "completed",
		CustomerEmail: user.Email,
		CustomerName:  user.FirstName,
		PaymentMethod: "vyouz-wallet",
	}

	// create a campaign record of the funding based on the info from request
	if err := initializers.DB.Create(&campaign).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while creating campaign")
	}

	// create a transaction record of the funding based on the info from request
	if err := initializers.DB.Create(&transaction).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while processing transaction")
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
