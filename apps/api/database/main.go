package database

import (
	"github.com/timzolleis/smallstatus/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate() {
	models := []interface{}{&model.User{}, &model.ApiKey{}, model.Monitor{}, model.Workspace{}, model.RequestHeader{}}
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Could not migrate database: %s", err.Error())
	}
}
