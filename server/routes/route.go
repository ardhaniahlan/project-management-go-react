package routes

import (
	"project-management-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, uc *controllers.UserController) {
	app.Post("/api/auth/register", uc.Register)
	app.Post("/api/auth/login", uc.Login)
}