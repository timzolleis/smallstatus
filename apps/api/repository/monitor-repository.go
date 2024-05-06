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
	return monitor, err
}

func (repository *MonitorRepository) FindById(id uint) (*model.Monitor, error) {
	var monitor model.Monitor
	err := database.DB.First(&monitor, id).Error
	return &monitor, err
}

func (repository *MonitorRepository) FindAllByWorkspace(workspace uint) ([]model.Monitor, error) {
	var monitors []model.Monitor
	err := database.DB.Where("workspace_id = ?", workspace).Find(&monitors).Error
	if err != nil {
		return nil, err
	}
	return monitors, nil
}

func (repository *MonitorRepository) Update(monitor *model.Monitor) (*model.Monitor, error) {
	result := database.DB.Save(monitor)
	err := result.Error
	if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return monitor, err
}

func (repository *MonitorRepository) Delete(id uint) error {
	result := database.DB.Delete(&model.Monitor{}, id)
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
