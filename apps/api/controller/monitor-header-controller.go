package controller

import (
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

func (controller *MonitorHeaderController) FindHeaders(c echo.Context) error {
	monitorId := helper.StringToUint(c.Param("id"))
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
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := helper.StringToUint(c.Param("id"))
	header, err := controller.monitorHeaderService.CreateHeader(&body, monitorId)
	if err != nil {
		return helper.HandleError(err, c)

	}
	return c.JSON(http.StatusCreated, mapMonitorHeaderDTO(header))
}

func (controller *MonitorHeaderController) FindHeader(c echo.Context) error {
	monitorId := helper.StringToUint(c.Param("id"))
	headerId := helper.StringToUint(c.Param("headerId"))
	header, err := controller.monitorHeaderService.FindHeaderById(monitorId, headerId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(header))
}

func (controller *MonitorHeaderController) UpdateHeader(c echo.Context) error {
	body := dto.MonitorHeaderDTO{}
	err := c.Bind(&body)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	monitorId := helper.StringToUint(c.Param("id"))
	headerId := helper.StringToUint(c.Param("headerId"))
	header := mapMonitorHeaderDTOToModel(&body, monitorId)
	header.ID = headerId
	updatedHeader, err := controller.monitorHeaderService.Update(&header, monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapMonitorHeaderDTO(updatedHeader))
}

func (controller *MonitorHeaderController) DeleteHeader(c echo.Context) error {
	monitorId := helper.StringToUint(c.Param("id"))
	headerId := helper.StringToUint(c.Param("headerId"))
	err := controller.monitorHeaderService.Delete(monitorId, headerId)
	if err != nil {
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
