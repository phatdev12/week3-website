package database

import (
	"log"
	"os"

	"github.com/phatdev12/week3-website/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DEBUG: %v", err)
	} else {
		log.Println("DEBUG: Connected to database")
	}

	DB.AutoMigrate(&models.Category{}, &models.Product{})
}
