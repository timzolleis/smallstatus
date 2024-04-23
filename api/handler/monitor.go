package handler

import (
	"encoding/json"
	"net/http"
	"status/api/helper"
	"status/internal/apiError"
	"status/internal/database/models"
	"status/internal/database/repository"
	"strconv"
)

type CreateMonitorBody struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	Interval int    `json:"interval"`
}

func CreateMonitorHandler(writer http.ResponseWriter, request *http.Request) {
	var createMonitoryBody CreateMonitorBody
	err := json.NewDecoder(request.Body).Decode(&createMonitoryBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	monitor := models.Monitor{Name: createMonitoryBody.Name, Interval: createMonitoryBody.Interval, Url: createMonitoryBody.Url}
	repository.CreateMonitor(&monitor)
	writer.WriteHeader(http.StatusCreated)
}

func GetMonitorHandler(writer http.ResponseWriter, request *http.Request) {
	idString := request.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		apiError.ReturnApiError(writer, apiError.GetApiError(400, "invalid_id"))
		return
	}
	monitor, err := repository.FindMonitor(id)
	if err != nil {
		apiError.ReturnApiError(writer, apiError.GetApiError(404, "monitor_not_found"))
		return
	}
	helper.ReturnJsonResponse(writer, monitor)

}

func DeleteMonitorHandler(writer http.ResponseWriter, request *http.Request) {
	idString := request.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(writer, "invalid_id", http.StatusBadRequest)
	}
	apiErr := repository.DeleteMonitor(id)
	if apiErr != nil {
		apiError.ReturnApiError(writer, apiErr)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
