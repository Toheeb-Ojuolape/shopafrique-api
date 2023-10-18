package authControllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string
	Password string
}

func LoginController(c *fiber.Ctx) error {
	//Get the email/password off the req body
	var req LoginRequest
	// this binds the req to the body struct
	if c.BodyParser(&req) != nil {
		return handleErrors.HandleBadRequest(c, "Invalid parameters passed")
	}

	//Look up the requested user
	var user models.User
	var count int64
	initializers.DB.Where("email = ?", req.Email).First(&user).Count(&count)

	if count == 0 {
		return handleErrors.HandleBadRequest(c, "You don't have an account, yet")
	}

	//Compare sent in password with the saved user's password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return handleErrors.HandleBadRequest(c, "The password you entered is wrong")
	}

	//Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return handleErrors.HandleBadRequest(c, "Failed to authenticate user")
	}

	//Send the token as a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(10 * time.Minute), // Expires in 10 mins
		SameSite: "Lax",
		Secure:   false, // Set to true if using HTTPS
		HTTPOnly: true,
	})
	//Send the token and other details to the user
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login Successful",
		"token":   tokenString,
		"data":    map[string]interface{}{"email": user.Email, "firstName": user.FirstName, "lastName": user.LastName, "country": user.Country},
	})
}
