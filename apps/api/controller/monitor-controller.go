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
	MonitorService       service.MonitorService
	MonitorHeaderService service.MonitorHeaderService
}

func (controller *MonitorController) FindAll(c echo.Context) error {
	workspaceIdString := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(workspaceIdString)
	monitors, err := controller.MonitorService.FindAll(workspaceId)
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
	id, _ := strconv.Atoi(c.Param("id"))
	workspace, _ := strconv.Atoi(c.Param("workspaceId"))
	monitor, err := controller.MonitorService.FindMonitorById(id, workspace)
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
	monitor, err := controller.MonitorService.CreateMonitor(body, workspaceId)
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
	updatedMonitor, err := controller.MonitorService.Update(&monitor)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorToDTO(updatedMonitor))

}

func (controller *MonitorController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	workspace, _ := strconv.Atoi(c.Param("workspaceId"))
	err := controller.MonitorService.Delete(id, workspace)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.NoContent(http.StatusNoContent)

}

func (controller *MonitorController) FindHeaders(c echo.Context) error {
	monitorId := helper.StringToUint(c.Param("id"))
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	headers, err := controller.MonitorHeaderService.FindHeaders(monitorId, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	headerDTOs := make([]dto.MonitorHeaderDTO, len(headers))
	for i, header := range headers {
		headerDTOs[i] = mapMonitorHeaderDTO(&header)
	}
	return c.JSON(http.StatusOK, headerDTOs)
}

func (controller *MonitorController) CreateHeader(c echo.Context) error {
	var body = dto.CreateMonitorHeaderDTO{}
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := helper.StringToUint(c.Param("id"))
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	header, err := controller.MonitorHeaderService.CreateHeader(&body, monitorId, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)

	}
	return c.JSON(http.StatusCreated, mapMonitorHeaderDTO(header))
}

func (controller *MonitorController) FindHeaderById(c echo.Context) error {
	monitorId := helper.StringToUint(c.Param("id"))
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	headerId := helper.StringToUint(c.Param("headerId"))
	header, err := controller.MonitorHeaderService.FindHeaderById(monitorId, workspaceId, headerId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(header))
}

func (controller *MonitorController) UpdateHeader(c echo.Context) error {
	body := dto.MonitorHeaderDTO{}
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := helper.StringToUint(c.Param("id"))
	workspaceId := helper.StringToUint(c.Param("workspaceId"))
	headerId := helper.StringToUint(c.Param("headerId"))
	header := mapMonitorHeaderDTOTOModel(&body, monitorId)
	header.ID = headerId
	updatedHeader, err := controller.MonitorHeaderService.Update(&header, monitorId, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(updatedHeader))

}

func mapMonitorHeaderDTO(header *model.MonitorHeader) dto.MonitorHeaderDTO {
	return dto.MonitorHeaderDTO{
		ID:    header.ID,
		Key:   header.Key,
		Value: header.Value,
	}

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

func mapMonitorHeaderDTOTOModel(dto *dto.MonitorHeaderDTO, monitorId uint) model.MonitorHeader {
	return model.MonitorHeader{
		Key:       dto.Key,
		Value:     dto.Value,
		MonitorID: monitorId,
	}
}
