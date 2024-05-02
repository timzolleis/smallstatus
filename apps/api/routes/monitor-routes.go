package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/controller"
)

func RegisterMonitorRoutes(baseGroup *echo.Group) {

	monitorController := controller.MonitorController{}

	monitorHeaderController := controller.MonitorHeaderController{}

	monitorGroup := baseGroup.Group("/monitors")
	monitorDetailsGroup := monitorGroup.Group("/:monitorId")

	//All monitor routes
	monitorGroup.GET("", monitorController.FindAll)
	monitorGroup.POST("", monitorController.Create)
	//Monitor specific routes
	monitorDetailsGroup.GET("", monitorController.FindById)
	monitorDetailsGroup.PUT("", monitorController.Update)
	monitorDetailsGroup.DELETE("", monitorController.Delete)

	//Monitor header routes
	monitorHeaderGroup := monitorDetailsGroup.Group("/headers")
	monitorHeaderGroup.GET("", monitorHeaderController.FindHeaders)
	monitorHeaderGroup.POST("", monitorHeaderController.CreateHeader)
	monitorHeaderGroup.GET("/:headerId", monitorHeaderController.FindHeader)
	monitorHeaderGroup.PUT("/:headerId", monitorHeaderController.UpdateHeader)
	monitorHeaderGroup.DELETE("/:headerId", monitorHeaderController.DeleteHeader)
}
