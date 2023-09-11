package authControllers

import (
	"fmt"
	"time"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/Toheeb-Ojuolape/shopafrique-api/services"
	"github.com/Toheeb-Ojuolape/shopafrique-api/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// sends an email otp to the user
func VerifyEmail(c *fiber.Ctx) error {
	var req types.EmailRequest
	var user models.User
	c.BodyParser(&req)

	// check if email is defined and is a valid email address
	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed. Request is missing "+missingProps)
	}

	// check if this user has an account already
	if err := initializers.DB.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return handleErrors.HandleBadRequest(c, "User already has an account")
	}

	sessionId := helpers.GenerateSessionId()
	otpNumber := helpers.GenerateOtp()
	expiry := time.Now().Add(time.Minute * 15)

	//hash the otpNumber stored in db
	hashedOtp, hashErr := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(otpNumber)), 10)

	if hashErr != nil {
		return handleErrors.HandleBadRequest(c, "Failed to hash password")
	}

	otp := models.Otp{ID: sessionId, Email: req.Email, Otp: string(hashedOtp), ExpiredAt: expiry}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&otp).Error; err != nil {
			return err
		}

		emailErr := services.SendMail(
			"Verify Your Email",
			fmt.Sprintf("<h1>Hey %v </h1> <p>Kindly use this otp to verify your email: <strong>%v</strong></p>", req.Email, otpNumber),
			string(req.Email),
		)

		if emailErr != nil {
			return emailErr
		}

		return nil
	})

	if err != nil {
		return handleErrors.HandleBadRequest(c, "Otp not sent successfully")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "OTP sent successfully",
		Data:    map[string]interface{}{"sessionId": sessionId},
	})
}
