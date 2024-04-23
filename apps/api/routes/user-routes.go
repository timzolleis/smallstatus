package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
)

func RegisterUserRoutes(baseGroup *echo.Group) {
	userController := controller.UserController{}
	group := baseGroup.Group("/users")
	group.GET("", userController.FindAll)
	group.POST("", userController.Create)
	group.GET("/:id", userController.FindById)
	group.DELETE("/:id", userController.Delete)
}
