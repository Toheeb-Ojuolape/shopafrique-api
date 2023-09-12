package authControllers

import (
	"fmt"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/Toheeb-Ojuolape/shopafrique-api/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SessionRequest struct {
	SessionId string
}

func ResendOtp(c *fiber.Ctx) error {

	var req SessionRequest
	var session models.Otp

	if err := c.BodyParser(&req); err != nil {
		// Handle error
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed")
	}

	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed. Request is missing "+missingProps)
	}

	if err := initializers.DB.Find(&session, "id = ?", req.SessionId).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Invalid Session")
	}

	otpNumber := helpers.GenerateOtp()
	expiry := helpers.GenerateExpiry()

	//hash the otpNumber stored in db
	hashedOtp, hashErr := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(otpNumber)), 10)

	if hashErr != nil {
		return handleErrors.HandleBadRequest(c, "Failed to hash password")
	}

	otp := models.Otp{ID: req.SessionId, Email: session.Email, Otp: string(hashedOtp), ExpiredAt: expiry}

	err := initializers.DB.Where("id = ?", req.SessionId).Updates(&otp)
	if err.Error != nil {
		return handleErrors.HandleBadRequest(c, "Otp not sent successfully")
	}

	emailErr := services.SendMail(
		"Verify Your Email",
		fmt.Sprintf("<h1>Hey %v </h1> <p>Kindly use this otp to verify your email: <strong>%v</strong></p>", session.Email, otpNumber),
		string(session.Email),
	)

	if emailErr != nil {
		return handleErrors.HandleBadRequest(c, fmt.Sprint(err))
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "OTP sent successfully",
		Data:    map[string]interface{}{"sessionId": req.SessionId},
	})
}
