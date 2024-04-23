package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
	"strconv"
)

type MonitorController struct {
	monitorService service.MonitorService
}

func (controller *MonitorController) FindAll(c echo.Context) error {
	workspaceIdString := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(workspaceIdString)
	monitors, err := controller.monitorService.FindAll(workspaceId)
	monitorDTOs := make([]dto.MonitorDTO, len(monitors))
	for i, monitor := range monitors {
		monitorDTOs[i] = mapMonitorToDTO(&monitor)
	}
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, monitorDTOs)
}

func (controller *MonitorController) FindById(c echo.Context) error {
	id := helper.StringToUint(c.Param("id"))
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	monitor, err := controller.monitorService.FindMonitorById(id, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorToDTO(monitor))
}

func (controller *MonitorController) Create(c echo.Context) error {
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	var body dto.CreateMonitorDTO
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitor, err := controller.monitorService.CreateMonitor(body, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusCreated, mapMonitorToDTO(monitor))
}

func (controller *MonitorController) Update(c echo.Context) error {
	var body dto.MonitorDTO
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitor := mapMonitorDTOToModel(&body, helper.StringToUint(c.Param("workspaceId")))
	monitor.ID = helper.StringToUint(c.Param("id"))
	updatedMonitor, err := controller.monitorService.Update(&monitor)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorToDTO(updatedMonitor))

}

func (controller *MonitorController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	workspace, _ := strconv.Atoi(c.Param("workspaceId"))
	err := controller.monitorService.Delete(id, workspace)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.NoContent(http.StatusNoContent)

}

func mapMonitorToDTO(monitor *model.Monitor) dto.MonitorDTO {
	return dto.MonitorDTO{
		ID:       monitor.ID,
		Name:     monitor.Name,
		Url:      monitor.Url,
		Interval: monitor.Interval,
		Type:     monitor.Type,
	}
}

func mapMonitorDTOToModel(dto *dto.MonitorDTO, workspace uint) model.Monitor {
	return model.Monitor{
		Name:        dto.Name,
		Url:         dto.Url,
		Interval:    dto.Interval,
		Type:        dto.Type,
		WorkspaceID: workspace,
	}
}
