package routes

import (
	"github.com/labstack/echo/v4"
	"status/controller"
	"status/repository"
	"status/service"
)

func RegisterUserRoutes(e *echo.Echo) {
	userController := controller.UserController{Service: service.UserService{Repository: repository.UserRepository{}}}
	e.GET("/api/users", userController.FindAll)
	e.POST("/api/users", userController.Create)
	e.GET("/api/users/:id", userController.FindById)
	e.DELETE("/api/users/:id", userController.Delete)
}
