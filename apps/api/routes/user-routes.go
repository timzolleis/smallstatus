package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
	"github.com/timzolleis/smallstatus/repository"
	"github.com/timzolleis/smallstatus/service"
)

func RegisterUserRoutes(baseGroup *echo.Group) {
	userController := controller.UserController{Service: service.UserService{Repository: repository.UserRepository{}}}
	group := baseGroup.Group("/users")
	group.GET("", userController.FindAll)
	group.POST("", userController.Create)
	group.GET("/:id", userController.FindById)
	group.DELETE("/:id", userController.Delete)
}
