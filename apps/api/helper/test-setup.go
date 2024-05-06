package helper

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	database.DB = db
	database.Migrate()
}

func SetupSession(c echo.Context) {
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))
}
