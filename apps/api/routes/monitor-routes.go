package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
	"github.com/timzolleis/smallstatus/repository"
	"github.com/timzolleis/smallstatus/service"
)

func RegisterMonitorRoutes(baseGroup *echo.Group) {
	monitorController := controller.MonitorController{Service: service.MonitorService{Repository: repository.MonitorRepository{}}}
	group := baseGroup.Group("/monitors")

	//Monitor routes
	group.GET("", monitorController.FindAll)
	group.GET("/:id", monitorController.FindById)
	group.POST("", monitorController.Create)
	group.DELETE("/:id", monitorController.Delete)
	//Monitor header routes
	group.GET("/:id/headers", monitorController.FindHeaders)

}
