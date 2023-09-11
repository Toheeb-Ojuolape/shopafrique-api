package authControllers

import (
	"fmt"
	"time"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type OtpRequest struct {
	Otp         string
	SessionId   string
	ProcessType string
}

func VerifyOtp(c *fiber.Ctx) error {
	var req OtpRequest
	c.BodyParser(&req)

	missingProps := helpers.ValidateRequest(req)
	if missingProps != "" {
		return handleErrors.HandleBadRequest(c, "Missing parameters. Kindly pass in "+missingProps)
	}

	var otp models.Otp
	if err := initializers.DB.Find(&otp, "id = ?", req.SessionId).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Invalid Session")
	}

	if helpers.HasEmptyValues(otp) {
		return handleErrors.HandleBadRequest(c, "Invalid Session")
	}

	if time.Now().Unix() > otp.ExpiredAt.Unix() {
		return handleErrors.HandleBadRequest(c, "OTP has expired")
	}

	// compare otp from request with hashed otp in db
	hashErr := bcrypt.CompareHashAndPassword([]byte(otp.Otp), []byte(req.Otp))
	if hashErr != nil {
		return handleErrors.HandleBadRequest(c, "OTP is invalid")
	}

	//if all pass, delete the session and return a processId
	if err := initializers.DB.Unscoped().Delete(&otp).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong "+fmt.Sprint(err))
	}
	//create processId
	processId := helpers.GenerateProcessId()
	expiry := helpers.GenerateExpiry()
	processRecord := models.Process{ID: processId, Email: otp.Email, Expiry: expiry, Process: req.ProcessType}

	err := initializers.DB.Create(&processRecord)

	if err.Error != nil {
		return handleErrors.HandleBadRequest(c, fmt.Sprint(err.Error))
	} else {
		return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
			Message: "OTP verified successfully",
			Data:    map[string]interface{}{"processId": processId},
		})
	}
}
