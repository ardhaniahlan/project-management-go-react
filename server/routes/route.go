package routes

import (
	"project-management-be/controllers"
	"project-management-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, uc *controllers.UserController) {
	auth := app.Group("/api/auth")
	auth.Post("/register", uc.Register)
	auth.Post("/login", uc.Login)

	user := app.Group("/api/users", middleware.Protected())
	user.Get("/page", uc.GetUsersPaginate)
	user.Get("/:id", uc.GetUser)
}