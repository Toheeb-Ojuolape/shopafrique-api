package authControllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/handleSuccess"
	"github.com/Toheeb-Ojuolape/shopafrique-api/helpers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Email        string
	Password     string
	Country      string
	FirstName    string
	LastName     string
	BusinessType string
	BusinessName string
	ProcessId    string
}

func Signup(c *fiber.Ctx) error {

	var req SignupRequest

	c.BodyParser(&req)
	missingProps := helpers.ValidateRequest(req)

	if missingProps != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid parameters passed",
			"error":   "Request is missing " + missingProps,
		})
	}

	// validate processId
	var process models.Process
	if err := initializers.DB.Find(&process, "id = ?", req.ProcessId).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Signup process is incomplete. Kindly try again")
	}

	if helpers.HasEmptyValues(process) {
		return handleErrors.HandleBadRequest(c, "Signup process is incomplete. Kindly restart signup")
	}

	if time.Now().Unix() > process.Expiry.Unix() {
		return handleErrors.HandleBadRequest(c, "Process has expired. Kindly try again from the beginning")
	}

	// check for fraud
	if req.Email != process.Email {
		return handleErrors.HandleBadRequest(c, "Suspected malicious signup. Verified email different from email at request")
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return handleErrors.HandleBadRequest(c, "Failed to hash password")
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong while creating user")
	}

	user := models.User{
		ID:               newUUID,
		Email:            process.Email,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Password:         string(hash),
		Country:          req.Country,
		BusinessName:     req.BusinessName,
		BusinessType:     req.BusinessType,
		Balance:          0.00,
		LightningAddress: "",
		Role:             "user",
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return handleErrors.HandleBadRequest(c, "User with this details already exits")
	}

	//delete the process
	if err := initializers.DB.Unscoped().Delete(&process).Error; err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong "+fmt.Sprint(err))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Failed to authenticate user")
	}

	return handleSuccess.HandleSuccessResponse(c, handleSuccess.SuccessResponse{
		Message: "Account created successfully",
		Data:    map[string]interface{}{"token": tokenString},
	})

}
