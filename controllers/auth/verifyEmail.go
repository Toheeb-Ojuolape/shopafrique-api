package authControllers

import (
	"fmt"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/Toheeb-Ojuolape/shopafrique-api/services"
	"github.com/Toheeb-Ojuolape/shopafrique-api/types"
	"github.com/gofiber/fiber/v2"
)

// sends an email otp to the user
func VerifyEmail(c *fiber.Ctx) error {
	var req types.EmailRequest
	var user models.User
	c.BodyParser(&req)

	// check if email is defined
	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed. Request is missing "+missingProps)
	}

	// check if this user has an account already
	if err := initializers.DB.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return handleErrors.HandleBadRequest(c, err.Error())
	}

	sessionId := helpers.GenerateSessionId()
	otpNumber := helpers.GenerateOtp()
	expiry := helpers.GenerateExpiry()

	otp := models.Otp{ID: sessionId, Email: req.Email, Otp: fmt.Sprint(otpNumber), ExpiredAt: expiry}

	err := initializers.DB.Create(&otp)
	if err.Error != nil {
		return handleErrors.HandleBadRequest(c, "Otp not sent successfully")
	}

	emailErr := services.SendMail(
		"Verify Your Email",
		fmt.Sprintf("<h1>Hey %v </h1> <p>Kindly use this otp to verify your email: <strong>%v</strong></p>", req.Email, otpNumber),
		string(req.Email),
	)

	if emailErr != nil {
		return handleErrors.HandleBadRequest(c, fmt.Sprint(err))
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "OTP sent successfully",
		Data:    map[string]interface{}{"sessionId": sessionId},
	})
}
