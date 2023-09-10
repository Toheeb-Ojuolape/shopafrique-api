package authControllers

import (
	"fmt"
	"net/http"

	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/Toheeb-Ojuolape/shopafrique-api/services"
	"github.com/gofiber/fiber/v2"
)

type EmailRequest struct {
	Email string
}

// sends an email otp to the user
func VerifyEmail(c *fiber.Ctx) error {
	var req EmailRequest
	var user models.User
	c.BodyParser(&req)

	// check if email is defined
	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid parameters passed",
			"error":   "Request is missing " + missingProps,
		})
	}

	// check if this user has an account already
	result := initializers.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "User with this email already exist",
		})
	}

	sessionId := helpers.GenerateSessionId()
	otpNumber := helpers.GenerateOtp()
	expiry := helpers.GenerateExpiry()

	otp := models.Otp{ID: sessionId, Email: req.Email, Otp: fmt.Sprint(otpNumber), ExpiredAt: expiry}

	err := initializers.DB.Create(&otp)
	fmt.Println(err)
	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Otp not sent successfully",
			"error":   fmt.Sprint(err.Error),
		})
	}

	services.SendMail(
		"Verify Your Email",
		fmt.Sprintf("<h1>Hey %v </h1> <p>Kindly use this otp to verify your email: <strong>%v</strong></p>", req.Email, otpNumber),
		string(req.Email),
		fmt.Sprint(sessionId),
		c,
	)

	return nil
}
