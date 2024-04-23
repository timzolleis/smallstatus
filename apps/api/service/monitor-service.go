package service

import (
	"status/model"
	"status/repository"
)

type MonitorService struct {
	repository repository.MonitorRepository
}

func (service *MonitorService) CreateMonitor(name string, url string, interval int, monitorType string) (*model.Monitor, error) {
	monitor := model.Monitor{
		Name:     name,
		Url:      url,
		Interval: interval,
		Type:     monitorType,
	}
	createdMonitor, err := service.repository.Create(&monitor)
	if err != nil {
		return nil, err
	}
	return createdMonitor, nil
}

func (service *MonitorService) FindMonitorById(id int) (*model.Monitor, error) {
	monitor, err := service.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return monitor, nil
}

func (service *MonitorService) FindAll(workspace int) ([]model.Monitor, error) {
	monitors, err := service.repository.FindAll(workspace)
	if err != nil {
		return nil, err
	}
	return monitors, nil
}
