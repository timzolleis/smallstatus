package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
)

type RequestHeaderRepository struct {
}

func (repository *RequestHeaderRepository) FindHeaders(monitorId uint) ([]model.RequestHeader, error) {
	var headers []model.RequestHeader
	err := database.DB.Where("monitor_id = ?", monitorId).Find(&headers).Error
	return headers, err
}

func (repository *RequestHeaderRepository) FindHeaderById(headerId uint) (*model.RequestHeader, error) {
	var header model.RequestHeader
	err := database.DB.Find(&header, headerId).Error
	return &header, err
}

func (repository *RequestHeaderRepository) CreateHeader(header *model.RequestHeader) (*model.RequestHeader, error) {
	err := database.DB.Create(header).Error
	return header, err
}

func (repository *RequestHeaderRepository) UpdateHeader(header *model.RequestHeader) (*model.RequestHeader, error) {
	err := database.DB.Save(header).Error
	return header, err
}

func (repository *RequestHeaderRepository) DeleteHeader(headerId uint) error {
	return database.DB.Delete(&model.RequestHeader{}, headerId).Error
}
