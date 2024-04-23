package repository

import (
	"status/internal/apiError"
	"status/internal/database"
	"status/internal/database/models"
)

func CreateMonitor(monitor *models.Monitor) *models.Monitor {
	database.Database.Create(monitor)
	return monitor
}

func FindMonitor(id int) (*models.Monitor, error) {
	var monitor models.Monitor
	err := database.Database.First(&monitor, id).Error
	if err != nil {
		return nil, err
	}
	return &monitor, nil
}

func UpdateMonitor(monitor *models.Monitor) *models.Monitor {
	database.Database.Updates(monitor)
	return monitor
}

func DeleteMonitor(id int) *apiError.ApiError {
	result := database.Database.Delete(&models.Monitor{}, id)
	if result.RowsAffected < 1 {
		return apiError.GetApiError(400, "Monitor not found")
	}
	return nil
}
