package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"status/internal/database/models"
)

var Database *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not open database connection: %s", err.Error())
	}
	Database = db
	return db
}

func Migrate() {
	err := Database.AutoMigrate(&models.Monitor{}, &models.User{}, &models.ApiKey{})
	if err != nil {
		log.Fatalf("Could not migrate database: %s", err.Error())
	}
}
