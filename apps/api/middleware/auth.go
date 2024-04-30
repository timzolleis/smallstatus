package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["user_id"] == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please login to view this page")
		}
		userService := service.UserService{}
		userID := sess.Values["user_id"].(uint)
		user, err := userService.FindUserById(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user session")
		}
		c.Set("user", user)
		return next(c)
	}
}
