package service

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type MonitorHeaderService struct {
	Repository repository.MonitorHeaderRepository
}

func (service *MonitorHeaderService) FindHeaders(monitorId uint, workspace uint) ([]model.MonitorHeader, error) {
	headers, err := service.Repository.FindHeaders(monitorId, workspace)
	return headers, err
}

func (service *MonitorHeaderService) FindHeaderById(monitorId uint, workspace uint, headerId uint) (*model.MonitorHeader, error) {
	header, err := service.Repository.FindHeaderById(monitorId, workspace, headerId)
	return header, err
}

func (service *MonitorHeaderService) CreateHeader(dto *dto.CreateMonitorHeaderDTO, monitorId uint, workspaceId uint) (*model.MonitorHeader, error) {
	header := model.MonitorHeader{
		Key:   dto.Key,
		Value: dto.Value,
	}
	createdHeader, err := service.Repository.CreateHeader(&header, monitorId, workspaceId)
	return createdHeader, err
}

func (service *MonitorHeaderService) Update(header *model.MonitorHeader, monitorId uint, workspaceId uint) (*model.MonitorHeader, error) {
	updatedHeader, err := service.Repository.Update(header, monitorId, workspaceId)
	return updatedHeader, err
}
