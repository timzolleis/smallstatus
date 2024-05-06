package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/service"
	"github.com/timzolleis/smallstatus/validations"
	"net/http"
)

type MonitorController struct {
	monitorService service.MonitorService
}

func getWorkspaceId(c echo.Context) uint {
	return helper.StringToUint(c.Param("workspaceId"))
}

func getMonitorId(c echo.Context) uint {
	return helper.StringToUint(c.Param("monitorId"))
}

func (controller *MonitorController) isMonitorInWorkspace(monitorId uint, workspaceId uint) bool {
	monitor, err := controller.monitorService.FindMonitorById(monitorId)
	if err != nil {
		return false
	}
	return monitor.WorkspaceID == workspaceId

}

func (controller *MonitorController) FindAll(c echo.Context) error {
	workspaceId := getWorkspaceId(c)
	monitors, err := controller.monitorService.FindAll(workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	monitorDTOs := make([]dto.MonitorDTO, len(monitors))
	for i, monitor := range monitors {
		monitorDTOs[i] = mapMonitorToDTO(&monitor)
	}
	return c.JSON(http.StatusOK, monitorDTOs)
}

func (controller *MonitorController) FindById(c echo.Context) error {
	monitorId := getMonitorId(c)
	workspaceId := getWorkspaceId(c)
	monitor, err := controller.monitorService.FindMonitorById(monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	if monitor.WorkspaceID != workspaceId {
		return helper.NewForbiddenError(c)
	}
	return c.JSON(http.StatusOK, mapMonitorToDTO(monitor))
}

func (controller *MonitorController) Create(c echo.Context) error {
	workspaceId := getWorkspaceId(c)
	var body dto.CreateMonitorDTO

	if err := c.Bind(&body); err != nil {
		return helper.InvalidRequest(c)
	}

	if err := validateMonitorDTO(body); err != nil {
		return helper.InvalidRequest(c)
	}
	monitor, err := controller.monitorService.CreateMonitor(body, workspaceId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusCreated, mapMonitorToDTO(monitor))
}

func validateMonitorDTO(body dto.CreateMonitorDTO) error {
	validate := validator.New()
	validate.RegisterValidation("http_method", validations.ValidateHttpMethod)
	return validate.Struct(body)
}

func (controller *MonitorController) Update(c echo.Context) error {
	var body dto.MonitorDTO
	if err := c.Bind(&body); err != nil {
		return helper.InvalidRequest(c)
	}
	workspaceId := getWorkspaceId(c)
	monitorId := getMonitorId(c)
	if !controller.isMonitorInWorkspace(monitorId, workspaceId) {
		return helper.NewForbiddenError(c)
	}
	monitor := mapMonitorDTOToModel(&body, workspaceId)
	updatedMonitor, err := controller.monitorService.Update(&monitor)

	if err != nil {
		return helper.HandleError(err, c)
	}

	return c.JSON(http.StatusOK, mapMonitorToDTO(updatedMonitor))
}

func (controller *MonitorController) Delete(c echo.Context) error {
	monitorId := getMonitorId(c)
	workspaceId := getWorkspaceId(c)
	if !controller.isMonitorInWorkspace(monitorId, workspaceId) {
		return helper.NewForbiddenError(c)
	}

	if err := controller.monitorService.Delete(monitorId); err != nil {
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
		Method:   monitor.Method,
		Retries:  monitor.Retries,
		Timeout:  monitor.Timeout,
	}
}

func mapMonitorDTOToModel(dto *dto.MonitorDTO, workspace uint) model.Monitor {
	return model.Monitor{
		Name:        dto.Name,
		Url:         dto.Url,
		Interval:    dto.Interval,
		Retries:     dto.Retries,
		Timeout:     dto.Timeout,
		Method:      dto.Method,
		WorkspaceID: workspace,
	}
}
