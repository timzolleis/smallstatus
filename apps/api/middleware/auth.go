package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/service"
	"log"
	"net/http"
)

var SessionName = "smallstatus-authentication-session"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionName, c)
		if err != nil {
			log.Printf("Could not retrieve session: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Please login to view this page")
		}
		if sess.Values["user_id"] == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please login to view this page")
		}
		userService := service.UserService{}
		userID := sess.Values["user_id"].(uint)
		user, err := userService.FindById(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user session")
		}
		c.Set("user", user)
		return next(c)
	}
}
