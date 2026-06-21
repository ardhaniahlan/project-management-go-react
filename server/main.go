package main

import (
	"log"
	"project-management-be/config"
	"project-management-be/controllers"
	"project-management-be/database/seeders"
	"project-management-be/models"
	"project-management-be/repositories"
	"project-management-be/routes"
	"project-management-be/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	log.Println("Menjalankan sinkronisasi database...")
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.BoardMember{},
		&models.List{},
		&models.ListPosition{},
		&models.Card{},
		&models.CardPosition{},
		&models.CardLabel{},
		&models.CardAssignees{},
		&models.CardAttachment{},
		&models.Label{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("Gagal melakukan migrasi tabel: ", err)
	}
	log.Println("Migrasi tabel berhasil! Database siap digunakan.")

	log.Println("Menjalankan seeder...")
	seeders.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	boardRepo := repositories.NewBoardRepository(config.DB)
	boardMemberRepo := repositories.NewBoardMemberRepository(config.DB)
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.SetupRoutes(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server berjalan di port ", port)
	app.Listen(":" + port)
}