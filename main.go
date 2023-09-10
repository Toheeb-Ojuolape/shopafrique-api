package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Everything's ok, stop worrying",
		})
	})

	auth := app.Group("auth")
	routes.AuthRoutes(auth)

	user := app.Group("user")
	routes.UserRoutes(user)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
