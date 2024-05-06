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

type MonitorHeaderController struct {
	monitorHeaderService service.MonitorHeaderService
}

func getHeaderId(c echo.Context) uint {
	return helper.StringToUint(c.Param("headerId"))
}

func validateRequest(c echo.Context, service *service.MonitorHeaderService) error {
	workspaceId := getWorkspaceId(c)
	monitorId := getMonitorId(c)
	if !validations.IsMonitorPartOfWorkspace(monitorId, workspaceId) {
		return helper.NewForbiddenError(c)
	}
	headerId := getHeaderId(c)
	if headerId != 0 {
		header, err := service.FindHeaderById(headerId)
		if err != nil || header.MonitorID != monitorId {
			return helper.NewNotFoundError(c)
		}

	}
	return nil

}

func (controller *MonitorHeaderController) FindHeaders(c echo.Context) error {
	if err := validateRequest(c, &controller.monitorHeaderService); err != nil {
		return err
	}
	monitorId := getMonitorId(c)
	headers, err := controller.monitorHeaderService.FindHeaders(monitorId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	headerDTOs := make([]dto.MonitorHeaderDTO, len(headers))
	for i, header := range headers {
		headerDTOs[i] = *convertModelToHeaderDTO(&header)
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
	if err := validateRequest(c, &controller.monitorHeaderService); err != nil {
		return err
	}
	header := convertCreateHeaderDTOToModel(&body, monitorId)
	createdHeader, err := controller.monitorHeaderService.CreateHeader(header)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusCreated, convertModelToHeaderDTO(createdHeader))
}

func validateCreateHeaderDTO(createHeaderDTO dto.CreateMonitorHeaderDTO) error {
	validate := validator.New()
	return validate.Struct(createHeaderDTO)
}

func (controller *MonitorHeaderController) FindHeader(c echo.Context) error {
	headerId := getHeaderId(c)
	if err := validateRequest(c, &controller.monitorHeaderService); err != nil {
		return err
	}
	header, err := controller.monitorHeaderService.FindHeaderById(headerId)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, convertModelToHeaderDTO(header))
}

func (controller *MonitorHeaderController) UpdateHeader(c echo.Context) error {
	body := dto.MonitorHeaderDTO{}
	if err := c.Bind(&body); err != nil {
		return helper.InvalidRequest(c)
	}
	if err := validateRequest(c, &controller.monitorHeaderService); err != nil {
		return err
	}
	headerId := getHeaderId(c)
	header, err := controller.monitorHeaderService.FindHeaderById(headerId)
	headerToUpdate := convertHeaderDTOToModel(&body, header)
	updatedHeader, err := controller.monitorHeaderService.UpdateHeader(headerToUpdate)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, convertModelToHeaderDTO(updatedHeader))
}

func (controller *MonitorHeaderController) DeleteHeader(c echo.Context) error {

	if err := validateRequest(c, &controller.monitorHeaderService); err != nil {
		return err
	}
	headerId := getHeaderId(c)
	if err := controller.monitorHeaderService.DeleteHeader(headerId); err != nil {
		return helper.HandleError(err, c)
	}

	return c.NoContent(http.StatusNoContent)
}

func convertCreateHeaderDTOToModel(dto *dto.CreateMonitorHeaderDTO, monitorId uint) *model.RequestHeader {
	return &model.RequestHeader{
		Key:       dto.Key,
		Value:     dto.Value,
		MonitorID: monitorId,
	}
}

func convertHeaderDTOToModel(dto *dto.MonitorHeaderDTO, oldHeader *model.RequestHeader) *model.RequestHeader {
	oldHeader.Key = dto.Key
	oldHeader.Value = dto.Value
	return oldHeader
}

func convertModelToHeaderDTO(header *model.RequestHeader) *dto.MonitorHeaderDTO {
	return &dto.MonitorHeaderDTO{
		ID:    header.ID,
		Key:   header.Key,
		Value: header.Value,
	}
}
