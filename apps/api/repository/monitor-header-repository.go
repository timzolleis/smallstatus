package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
	"gorm.io/gorm"
)

type MonitorHeaderRepository struct {
}

func (repository *MonitorHeaderRepository) FindHeaders(monitor *model.Monitor) ([]model.MonitorHeader, error) {
	var headers []model.MonitorHeader
	err := database.DB.Model(monitor).
		Association("Headers").
		Find(&headers)
	return headers, err
}

func (repository *MonitorHeaderRepository) FindHeaderById(headerId uint, monitor *model.Monitor) (*model.MonitorHeader, error) {
	var header model.MonitorHeader
	err := database.DB.Model(monitor).Association("Headers").Find(&header, headerId)
	if err != nil {
		return nil, err
	}
	if header.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &header, err
}

func (repository *MonitorHeaderRepository) CreateHeader(header *model.MonitorHeader, monitor *model.Monitor) (*model.MonitorHeader, error) {
	err := database.DB.Model(monitor).Association("Headers").Append(header)
	return header, err
}

func (repository *MonitorHeaderRepository) Update(header *model.MonitorHeader, monitor *model.Monitor) (*model.MonitorHeader, error) {
	err := database.DB.Model(monitor).Association("Headers").Replace(header)
	return header, err
}

func (repository *MonitorHeaderRepository) DeleteHeader(headerId uint, monitor *model.Monitor) error {
	return database.DB.Model(monitor).Association("Headers").Delete(headerId)
}
