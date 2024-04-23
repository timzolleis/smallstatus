package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
)

func RegisterMonitorRoutes(baseGroup *echo.Group) {

	monitorController := controller.MonitorController{}

	monitorHeaderController := controller.MonitorHeaderController{}

	//Monitor routes
	monitorGroup := baseGroup.Group("/monitors")
	monitorGroup.GET("", monitorController.FindAll)
	monitorGroup.POST("", monitorController.Create)
	monitorGroup.GET("/:id", monitorController.FindById)
	monitorGroup.PUT("/:id", monitorController.Update)
	monitorGroup.DELETE("/:id", monitorController.Delete)

	//Monitor header routes
	monitorHeaderGroup := monitorGroup.Group("/:id/headers")
	monitorHeaderGroup.GET("", monitorHeaderController.FindHeaders)
	monitorHeaderGroup.POST("", monitorHeaderController.CreateHeader)
	monitorHeaderGroup.GET("/:headerId", monitorHeaderController.FindHeader)
	monitorHeaderGroup.PUT("/:headerId", monitorHeaderController.UpdateHeader)
	monitorHeaderGroup.DELETE("/:headerId", monitorHeaderController.DeleteHeader)
}
