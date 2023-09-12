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
	"golang.org/x/crypto/bcrypt"
)

func ForgotPassword(c *fiber.Ctx) error {
	var req types.EmailRequest
	var user models.User
	var count int64

	c.BodyParser(&req)
	// check if email is defined
	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed. Request is missing "+missingProps)
	}

	initializers.DB.First(&user).Where("email = ?", req.Email).Count(&count)

	if count == 0 {
		return handleErrors.HandleBadRequest(c, "You don't have an account, yet")
	}

	sessionId := helpers.GenerateSessionId()
	otpNumber := helpers.GenerateOtp()
	expiry := helpers.GenerateExpiry()
	//hash the otpNumber stored in db
	hashedOtp, hashErr := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(otpNumber)), 10)

	if hashErr != nil {
		return handleErrors.HandleBadRequest(c, "Failed to hash password")
	}

	otp := models.Otp{ID: sessionId, Email: req.Email, Otp: string(hashedOtp), ExpiredAt: expiry}

	err := initializers.DB.Create(&otp)
	if err.Error != nil {
		return handleErrors.HandleBadRequest(c, "Otp not sent successfully")
	}

	emailErr := services.SendMail(
		"Verify Your Email",
		fmt.Sprintf("<h1>Hey %v.</h1> <p> Sorry you forgot your password. Kindly use this otp to verify your email: <strong>%v</strong></p>", user.FirstName, otpNumber),
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
