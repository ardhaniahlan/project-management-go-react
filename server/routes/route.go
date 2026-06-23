package routes

import (
	"project-management-be/controllers"
	"project-management-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, uc *controllers.UserController, bc *controllers.BoardController, lc *controllers.ListController) {
	auth := app.Group("/api/auth")
	auth.Post("/register", uc.Register)
	auth.Post("/login", uc.Login)

	user := app.Group("/api/users", middleware.Protected())
	user.Get("/page", uc.GetUsersPaginate)
	user.Get("/:id", uc.GetUser)
	user.Put("/:id", uc.UpdateUser)
	user.Delete("/:id", uc.DeleteUser)

	board := app.Group("/api/boards", middleware.Protected())
	board.Get("/my", bc.GetMyBoardPaginate)
	board.Post("/", bc.CreateBoard)
	board.Put("/:id", bc.UpdateBoard)
	board.Post("/:id/members", bc.AddBoardMembers)
	board.Delete("/:id/members", bc.RemoveBoardMembers)

	list := app.Group("/api/lists", middleware.Protected())
	list.Post("/", lc.CreateList)
	list.Put("/:id", lc.UpdateList)
}