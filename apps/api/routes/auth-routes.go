package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
)

func RegisterAuthRoutes(group *echo.Group) {
	authController := controller.AuthController{}
	group.POST("/auth/signup", authController.SignUp)
	group.POST("/auth/login", authController.Login)
}
