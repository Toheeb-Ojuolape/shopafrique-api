package routes

import (
	userControllers "github.com/Toheeb-Ojuolape/shopafrique-api/controllers/user"
	middleware "github.com/Toheeb-Ojuolape/shopafrique-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(user fiber.Router) {
	user.Get("/", middleware.VerifyToken, userControllers.UserController)
}
