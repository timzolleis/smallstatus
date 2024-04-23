package repository

import (
	"status/database"
	"status/model"
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

func (repository *MonitorRepository) FindById(id int) (*model.Monitor, error) {
	var monitor model.Monitor
	err := database.DB.First(&monitor, id).Error
	if err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (repository *MonitorRepository) FindAll(workspace int) ([]model.Monitor, error) {
	var monitors []model.Monitor
	err := database.DB.Find(&monitors, "workspace_id = ?", workspace).Error
	if err != nil {
		return nil, err
	}
	return monitors, nil
}
