package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort         string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	JWTExpired      string
	JWTRefreshToken string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	AppConfig = &Config{
		AppPort:         getEnv("APP_PORT", "3030"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5433"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "password"),
		DBName:          getEnv("DB_NAME", "project_management"),
		JWTSecret:       getEnv("JWT_SECRET", "rahasia"),
		JWTExpired:      getEnv("JWT_EXPIRED", "2h"),
		JWTRefreshToken: getEnv("JWT_REFRESH_TOKEN", "24h"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectDB(){
	cfg := AppConfig

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get database connection", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	log.Println("Connected to database successfully")
}
