package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
)

type MonitorHeaderRepository struct {
}

func (repository *MonitorHeaderRepository) FindHeaders(monitorId uint, workspace uint) ([]model.MonitorHeader, error) {
	var headers []model.MonitorHeader
	err := database.DB.Model(&model.Monitor{}).
		Where("id = ?", monitorId).
		Where("workspace_id = ?", workspace).
		Association("Headers").
		Find(&headers)
	return headers, err
}

func (repository *MonitorHeaderRepository) FindHeaderById(monitorId uint, workspace uint, headerId uint) (*model.MonitorHeader, error) {
	var header model.MonitorHeader
	err := database.DB.Model(&model.Monitor{}).
		Where("id = ?", monitorId).
		Where("workspace_id = ?", workspace).
		Association("Headers").
		Find(&header, headerId)
	return &header, err
}

func (repository *MonitorHeaderRepository) CreateHeader(header *model.MonitorHeader, monitorId uint, workspaceId uint) (*model.MonitorHeader, error) {
	err := database.DB.Model(&model.Monitor{}).Where("id = ?", monitorId).Where("workspace_id = ?", workspaceId).Association("Headers").Append(header)
	return header, err
}

func (repository *MonitorHeaderRepository) Update(header *model.MonitorHeader, monitorId uint, workspaceId uint) (*model.MonitorHeader, error) {
	err := database.DB.Model(&model.Monitor{}).Where("id = ?", monitorId).Where("workspace_id = ?", workspaceId).Association("Headers").Replace(header)
	return header, err
}
