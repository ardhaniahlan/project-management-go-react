package routes

import (
	"project-management-be/controllers"
	"project-management-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, uc *controllers.UserController, bc *controllers.BoardController) {
	auth := app.Group("/api/auth")
	auth.Post("/register", uc.Register)
	auth.Post("/login", uc.Login)

	user := app.Group("/api/users", middleware.Protected())
	user.Get("/page", uc.GetUsersPaginate)
	user.Get("/:id", uc.GetUser)
	user.Put("/:id", uc.UpdateUser)
	user.Delete("/:id", uc.DeleteUser)

	board := app.Group("/api/boards", middleware.Protected())
	board.Post("/", bc.CreateBoard)
	board.Put("/:id", bc.UpdateBoard)
	board.Post("/:id/members", bc.AddBoardMembers)
	board.Delete("/:id/members", bc.RemoveBoardMembers)
}