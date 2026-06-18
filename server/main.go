package main

import (
	"log"
	"project-management-be/config"
	"project-management-be/database/seeders"
	"project-management-be/models"
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
}