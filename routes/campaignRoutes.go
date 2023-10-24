package routes

import (
	campaignControllers "github.com/Toheeb-Ojuolape/shopafrique-api/controllers/campaigns"
	middleware "github.com/Toheeb-Ojuolape/shopafrique-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func CampaignRoutes(campaign fiber.Router) {
	campaign.Post("/", middleware.VerifyToken, campaignControllers.CreateCampaign)
	campaign.Get("/", middleware.VerifyToken, campaignControllers.FetchCampaigns)
	campaign.Get("/:id", campaignControllers.FetchSingleCampaign)
	campaign.Patch("/:id", campaignControllers.UpdateSingleCampaign)
}
