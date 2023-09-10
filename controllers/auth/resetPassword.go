package authControllers

import (
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordRequest struct {
	ProcessId string
	Password  string
}

func ResetPassword(c *fiber.Ctx) error {
	var req ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		// Handle error
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed")
	}

	var process models.Process
	if err := initializers.DB.Find(&process, "id = ?", req.ProcessId).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Invalid ProcessId")
	}

	if helpers.HasEmptyValues(process) {
		return handleErrors.HandleBadRequest(c, "Invalid ProcessId")
	}

	var user models.User
	userDb := initializers.DB.Where("email = ?", process.Email).First(&user)

	if userDb.Error != nil {
		return handleErrors.HandleBadRequest(c, "Error fetching user by email data")
	}

	if process.Email != user.Email {
		return handleErrors.HandleBadRequest(c, "Suspected maliciousness, the email passed is different from the email in user db")
	}

	if len(req.Password) < 9 {
		return handleErrors.HandleBadRequest(c, "You have entered a weak password, kindly use something stronger")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err == nil {
		return handleErrors.HandleBadRequest(c, "Password already used. Kindly use another one")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return handleErrors.HandleBadRequest(c, err.Error())
	}

	//set the new password
	if err := initializers.DB.Model(&user).Where("id = ?", user.ID).Update("password", string(hash)).Error; err != nil {
		return handleErrors.HandleBadRequest(c, err.Error())

	}
	//delete the process
	if err := initializers.DB.Delete(&process).Error; err != nil {
		return handleErrors.HandleBadRequest(c, err.Error())
	}
	// Password update successful
	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Password updated successfully",
	})

}
