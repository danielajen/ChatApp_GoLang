package main

import (
	"log"

	"github.com/daniel/go-rest-api/database"
	"github.com/daniel/go-rest-api/models"
	"github.com/daniel/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Connected to PostgreSQL!",
		})
	})
	routes.Setup(app)
	app.Listen(":3000")
}
