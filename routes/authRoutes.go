package routes

import (
	authControllers "github.com/Toheeb-Ojuolape/shopafrique-api/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(auth fiber.Router) {
	auth.Post("/signup", authControllers.Signup)
	auth.Post("/verify-email", authControllers.VerifyEmail)
	auth.Post("/verify-otp", authControllers.VerifyOtp)
	auth.Post("/forgot-password", authControllers.ForgotPassword)
	auth.Post("/reset-password", authControllers.ResetPassword)
}
