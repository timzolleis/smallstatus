package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"status/helper"
	"status/service"
	"strconv"
)

type MonitorController struct {
	service service.MonitorService
}

func (controller *MonitorController) FindAll(c echo.Context) error {
	workspaceIdString := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(workspaceIdString)
	monitors, err := controller.service.FindAll(workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, monitors)
}
