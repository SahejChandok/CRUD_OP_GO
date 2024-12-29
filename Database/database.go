package database

import (
	"CRUD_GO/models"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err!=nil{
		log.Println("Error loading the env files")
	}
	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Println("Error connecting to db")
	}

	DB.AutoMigrate(&models.Item{})
	log.Println("Database connected successfully")
}


