package service

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type MonitorHeaderService struct {
	MonitorHeaderRepository repository.MonitorHeaderRepository
	MonitorRepository       repository.MonitorRepository
}

func (service *MonitorHeaderService) FindHeaders(monitorId uint) ([]model.MonitorHeader, error) {
	monitor, err := service.MonitorRepository.FindById(monitorId)
	if err != nil {
		return nil, err
	}
	headers, err := service.MonitorHeaderRepository.FindHeaders(monitor)
	return headers, err
}

func (service *MonitorHeaderService) FindHeaderById(monitorId uint, headerId uint) (*model.MonitorHeader, error) {
	monitor, err := service.MonitorRepository.FindById(monitorId)
	if err != nil {
		return nil, err
	}
	header, err := service.MonitorHeaderRepository.FindHeaderById(headerId, monitor)
	return header, err
}

func (service *MonitorHeaderService) CreateHeader(dto *dto.CreateMonitorHeaderDTO, monitorId uint) (*model.MonitorHeader, error) {
	header := &model.MonitorHeader{
		Key:   dto.Key,
		Value: dto.Value,
	}
	monitor, err := service.MonitorRepository.FindById(monitorId)
	if err != nil {
		return nil, err
	}
	createdHeader, err := service.MonitorHeaderRepository.CreateHeader(header, monitor)
	return createdHeader, err
}

func (service *MonitorHeaderService) Update(header *model.MonitorHeader, monitorId uint) (*model.MonitorHeader, error) {
	monitor, err := service.MonitorRepository.FindById(monitorId)
	if err != nil {
		return nil, err
	}
	updatedHeader, err := service.MonitorHeaderRepository.Update(header, monitor)
	return updatedHeader, err
}

func (service *MonitorHeaderService) Delete(monitorId uint, headerId uint) error {
	monitor, err := service.MonitorRepository.FindById(monitorId)
	if err != nil {
		return err
	}
	err = service.MonitorHeaderRepository.DeleteHeader(headerId, monitor)
	return err
}
