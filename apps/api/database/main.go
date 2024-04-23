package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"status/model"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not open database connection: %s", err.Error())
	}
	DB = db
	return db
}

func Migrate() {
	models := []interface{}{&model.User{}, &model.ApiKey{}, model.Monitor{}, model.Workspace{}}
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Could not migrate database: %s", err.Error())
	}
}
