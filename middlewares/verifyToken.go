package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Toheeb-Ojuolape/shopafrique-api/handleErrors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(c *fiber.Ctx) error {
	BearerToken := c.Get("Authorization")

	if BearerToken == "" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"message": "Access Denied",
		})
	}

	tokenString := strings.Split(BearerToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return handleErrors.HandleBadRequest(c, "Something went wrong")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Access Unauthorized",
			})
		}

		// Attach to request context
		c.Locals("user_id", claims["sub"])

		// Continue to the next middleware/handler
		return c.Next()
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"message": "Invalid Token",
	})
}
