package routes

import (
	"github.com/daniel/go-rest-api/controllers"
	"github.com/daniel/go-rest-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controllers.RegisterUser)
	api.Post("/login", controllers.LoginUser)
	api.Get("/profile", middlewares.RequireAuth, controllers.Profile)
	api.Put("/profile", middlewares.RequireAuth, controllers.UpdateProfile)
	api.Patch("/profile", middlewares.RequireAuth, controllers.UpdateProfile)
}
