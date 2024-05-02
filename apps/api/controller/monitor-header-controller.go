package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
)

type MonitorHeaderController struct {
	monitorHeaderService service.MonitorHeaderService
}

func getHeaderId(c echo.Context) uint {
	return helper.StringToUint(c.Param("headerId"))
}

func (controller *MonitorHeaderController) FindHeaders(c echo.Context) error {
	monitorId := getMonitorId(c)
	headers, err := controller.monitorHeaderService.FindHeaders(monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	headerDTOs := make([]dto.MonitorHeaderDTO, len(headers))
	for i, header := range headers {
		headerDTOs[i] = mapMonitorHeaderDTO(&header)
	}
	return c.JSON(http.StatusOK, headerDTOs)
}

func (controller *MonitorHeaderController) CreateHeader(c echo.Context) error {
	var body = dto.CreateMonitorHeaderDTO{}
	if err := c.Bind(&body); err != nil {
		return helper.InvalidRequest(c)
	}
	if err := validateCreateHeaderDTO(body); err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := getMonitorId(c)
	header, err := controller.monitorHeaderService.CreateHeader(&body, monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}

	return c.JSON(http.StatusCreated, mapMonitorHeaderDTO(header))
}

func validateCreateHeaderDTO(createHeaderDTO dto.CreateMonitorHeaderDTO) error {
	validate := validator.New()
	return validate.Struct(createHeaderDTO)
}

func (controller *MonitorHeaderController) FindHeader(c echo.Context) error {
	monitorId := getMonitorId(c)
	headerId := getHeaderId(c)
	header, err := controller.monitorHeaderService.FindHeaderById(monitorId, headerId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(header))
}

func (controller *MonitorHeaderController) UpdateHeader(c echo.Context) error {
	body := dto.MonitorHeaderDTO{}
	if err := c.Bind(&body); err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := getMonitorId(c)
	header := mapMonitorHeaderDTOToModel(&body, monitorId)
	updatedHeader, err := controller.monitorHeaderService.Update(&header, monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(updatedHeader))
}

func (controller *MonitorHeaderController) DeleteHeader(c echo.Context) error {
	monitorId := getMonitorId(c)
	headerId := getHeaderId(c)
	if err := controller.monitorHeaderService.Delete(monitorId, headerId); err != nil {
		return helper.HandleError(err, c)
	}

	return c.NoContent(http.StatusNoContent)
}

func mapMonitorHeaderDTO(header *model.MonitorHeader) dto.MonitorHeaderDTO {
	return dto.MonitorHeaderDTO{
		ID:    header.ID,
		Key:   header.Key,
		Value: header.Value,
	}
}

func mapMonitorHeaderDTOToModel(dto *dto.MonitorHeaderDTO, monitorId uint) model.MonitorHeader {
	return model.MonitorHeader{
		Key:       dto.Key,
		Value:     dto.Value,
		MonitorID: monitorId,
	}
}
