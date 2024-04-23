package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
	"gorm.io/gorm"
)

type MonitorRepository struct {
}

func (repository *MonitorRepository) Create(monitor *model.Monitor) (*model.Monitor, error) {
	err := database.DB.Create(monitor).Error
	if err != nil {
		return nil, err
	}
	return monitor, nil
}

func (repository *MonitorRepository) FindById(id int, workspace int) (*model.Monitor, error) {
	var monitor model.Monitor
	err := database.DB.Where("id = ?", id).Where("workspace_id = ?", workspace).Find(&monitor).Error
	if err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (repository *MonitorRepository) FindAll(workspace int) ([]model.Monitor, error) {
	var monitors []model.Monitor
	err := database.DB.Where("workspace_id = ?", workspace).Find(&monitors).Error
	if err != nil {
		return nil, err
	}
	return monitors, nil
}

func (repository *MonitorRepository) Update(monitor *model.Monitor) (*model.Monitor, error) {
	err := database.DB.Save(monitor).Error
	if err != nil {
		return nil, err
	}
	return monitor, nil
}

func (repository *MonitorRepository) Delete(id int, workspace int) error {
	result := database.DB.Where("id = ?", id).Where("workspace_id = ?", workspace).Delete(&model.Monitor{}, id, workspace)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Monitor headers

func (repository *MonitorRepository) FindHeaders(id int, workspace int) ([]model.MonitorHeader, error) {
	var headers []model.MonitorHeader
	err := database.DB.Joins("JOIN monitors ON monitors.id = monitor_headers.monitor_id").
		Where("monitors.id = ?", id).
		Where("monitors.workspace_id = ?", workspace).
		Find(&headers).Error
	if err != nil {
		return nil, err
	}
	return headers, nil
}
