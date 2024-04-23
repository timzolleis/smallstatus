package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
	"github.com/timzolleis/smallstatus/repository"
	"github.com/timzolleis/smallstatus/service"
)

func RegisterMonitorRoutes(baseGroup *echo.Group) {
	monitorController := controller.MonitorController{MonitorService: service.MonitorService{Repository: repository.MonitorRepository{}}}
	group := baseGroup.Group("/monitors")

	//Monitor routes
	group.GET("", monitorController.FindAll)
	group.POST("", monitorController.Create)
	group.GET("/:id", monitorController.FindById)
	group.PUT("/:id", monitorController.Update)
	group.DELETE("/:id", monitorController.Delete)
	//Monitor header routes
	group.GET("/:id/headers", monitorController.FindHeaders)
	group.POST("/:id/headers", monitorController.CreateHeader)
	group.GET("/:id/headers/:headerId", monitorController.FindHeaderById)
	group.PUT("/:id/headers/:headerId", monitorController.UpdateHeader)
	group.DELETE("/:id/headers/:headerId", monitorController.DeleteHeader)
}
