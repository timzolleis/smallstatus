package service

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type MonitorService struct {
	Repository repository.MonitorRepository
}

type CreateMonitorBody struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
	Type     string `json:"type"`
}

func (service *MonitorService) CreateMonitor(dto dto.CreateMonitorDTO, workspace uint) (*model.Monitor, error) {
	headers := make([]model.MonitorHeader, len(dto.Headers))
	for i, header := range dto.Headers {
		headers[i] = model.MonitorHeader{
			Key:   header.Key,
			Value: header.Value,
		}
	}
	monitor := model.Monitor{
		Name:        dto.Name,
		Url:         dto.Url,
		Interval:    dto.Interval,
		Type:        dto.Type,
		WorkspaceID: workspace,
		Headers:     headers,
	}
	createdMonitor, err := service.Repository.Create(&monitor)
	if err != nil {
		return nil, err
	}
	return createdMonitor, nil
}

func (service *MonitorService) FindMonitorById(id uint, workspace uint) (*model.Monitor, error) {
	monitor, err := service.Repository.FindById(id, workspace)
	if err != nil {
		return nil, err
	}
	return monitor, nil
}

func (service *MonitorService) FindAll(workspace int) ([]model.Monitor, error) {
	monitors, err := service.Repository.FindAll(workspace)
	if err != nil {
		return nil, err
	}
	return monitors, nil
}

func (service *MonitorService) Update(monitor *model.Monitor) (*model.Monitor, error) {
	updatedMonitor, err := service.Repository.Update(monitor)
	if err != nil {
		return nil, err
	}
	return updatedMonitor, nil
}

func (service *MonitorService) Delete(id int, workspace int) error {
	err := service.Repository.Delete(id, workspace)
	if err != nil {
		return err
	}
	return nil
}

// Monitor headers

func (service *MonitorService) FindHeaders(id int, workspaceId int) ([]model.MonitorHeader, error) {
	headers, err := service.Repository.FindHeaders(id, workspaceId)
	if err != nil {
		return nil, err
	}
	return headers, nil
}
