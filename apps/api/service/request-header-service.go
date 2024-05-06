package service

import (
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type MonitorHeaderService struct {
	MonitorHeaderRepository repository.RequestHeaderRepository
	MonitorRepository       repository.MonitorRepository
}

func (service *MonitorHeaderService) FindHeaders(monitorId uint) ([]model.RequestHeader, error) {
	headers, err := service.MonitorHeaderRepository.FindHeaders(monitorId)
	return headers, err
}

func (service *MonitorHeaderService) FindHeaderById(headerId uint) (*model.RequestHeader, error) {
	header, err := service.MonitorHeaderRepository.FindHeaderById(headerId)
	return header, err
}

func (service *MonitorHeaderService) CreateHeader(header *model.RequestHeader) (*model.RequestHeader, error) {
	createdHeader, err := service.MonitorHeaderRepository.CreateHeader(header)
	return createdHeader, err
}

func (service *MonitorHeaderService) UpdateHeader(header *model.RequestHeader) (*model.RequestHeader, error) {
	updatedHeader, err := service.MonitorHeaderRepository.UpdateHeader(header)
	return updatedHeader, err
}

func (service *MonitorHeaderService) DeleteHeader(headerId uint) error {
	err := service.MonitorHeaderRepository.DeleteHeader(headerId)
	return err
}
