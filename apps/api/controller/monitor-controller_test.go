package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func getCreateMonitorDTO() *dto.CreateMonitorDTO {
	return &dto.CreateMonitorDTO{
		Name:     "Test Monitor",
		Url:      "https://example.com",
		Method:   "GET",
		Timeout:  10,
		Interval: 60,
		Retries:  3,
	}
}

func TestMonitorController_FindAll(t *testing.T) {
	helper.SetupDb()
	e := echo.New()

	monitorService := service.MonitorService{Repository: repository.MonitorRepository{}}
	monitorController := MonitorController{monitorService: monitorService}
	e.GET("/workspaces/:workspaceId", monitorController.FindAll)
	createMonitorDTO := getCreateMonitorDTO()
	monitor, err := monitorService.CreateMonitor(*createMonitorDTO, 1)
	require.NoError(t, err)
	tests := []struct {
		name          string
		workspaceId   uint
		expectedDtos  []dto.MonitorDTO
		expectedCount int
	}{
		{name: "One monitor in workspace 1", workspaceId: 1, expectedDtos: []dto.MonitorDTO{mapMonitorToDTO(monitor)}, expectedCount: 1},
		{name: "No monitors in workspace 2", workspaceId: 2, expectedDtos: make([]dto.MonitorDTO, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/workspaces/:workspaceId/monitors")
			c.SetParamNames("workspaceId")
			c.SetParamValues(strconv.FormatUint(uint64(tt.workspaceId), 10))
			helper.SetupSession(c)
			monitorController.FindAll(c)
			assert.Equal(t, http.StatusOK, rec.Code)
			var monitors []dto.MonitorDTO
			err = json.Unmarshal(rec.Body.Bytes(), &monitors)
			require.NoError(t, err, "Unable to unmarshal response JSON")
			assert.Equal(t, tt.expectedCount, len(monitors))
			assert.Equal(t, tt.expectedDtos, monitors)
		})
	}
}

func TestMonitorController_FindById(t *testing.T) {
	helper.SetupDb()
	e := echo.New()
	monitorService := service.MonitorService{Repository: repository.MonitorRepository{}}
	monitorController := MonitorController{monitorService: monitorService}

	e.GET("/workspaces/:workspaceId/monitors/:monitorId", monitorController.FindById)
	createMonitorDTO := getCreateMonitorDTO()
	monitor, err := monitorService.CreateMonitor(*createMonitorDTO, 1)
	require.NoError(t, err)

	tests := []struct {
		name         string
		monitorId    uint
		workspaceId  string
		expectedCode int
		monitor      *model.Monitor
	}{
		{name: "Monitor in workspace 1", monitorId: monitor.ID, workspaceId: "1", expectedCode: http.StatusOK, monitor: monitor},
		{name: "Monitor in workspace 2", monitorId: monitor.ID, workspaceId: "2", expectedCode: http.StatusForbidden, monitor: nil},
		{name: "Monitor not found", monitorId: 999, workspaceId: "1", expectedCode: http.StatusNotFound, monitor: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/workspaces/:workspaceId/monitors/:monitorId")
			c.SetParamNames("workspaceId", "monitorId")
			c.SetParamValues(tt.workspaceId, strconv.FormatUint(uint64(tt.monitorId), 10))
			helper.SetupSession(c)
			monitorController.FindById(c)
			assert.Equal(t, tt.expectedCode, c.Response().Status)
			if tt.monitor != nil {
				var monitorDto dto.MonitorDTO
				err := json.Unmarshal(rec.Body.Bytes(), &monitorDto)
				require.NoError(t, err)
				assert.Equal(t, mapMonitorToDTO(tt.monitor), monitorDto)
			}
		})
	}
}
