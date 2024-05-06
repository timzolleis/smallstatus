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
	// Initialize the test database and the Echo instance
	helper.SetupDb()
	e := echo.New()

	// Initialize the service and controller
	monitorService := service.MonitorService{
		Repository: repository.MonitorRepository{},
	}
	monitorController := MonitorController{monitorService: monitorService}

	// Define the endpoint to be tested
	e.GET("/workspaces/:workspaceId/monitors", monitorController.FindAll)

	// Create a sample monitor for testing
	createMonitorDTO := getCreateMonitorDTO()
	monitor, err := monitorService.CreateMonitor(*createMonitorDTO, 1)
	require.NoError(t, err)

	// Test cases
	tests := []struct {
		name          string
		workspaceId   uint
		expectedDtos  []dto.MonitorDTO
		expectedCount int
	}{
		{
			name:          "One monitor in workspace 1",
			workspaceId:   1,
			expectedDtos:  []dto.MonitorDTO{mapMonitorToDTO(monitor)},
			expectedCount: 1,
		},
		{
			name:          "No monitors in workspace 2",
			workspaceId:   2,
			expectedDtos:  []dto.MonitorDTO{},
			expectedCount: 0,
		},
	}

	// Execute each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request and context
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/workspaces/:workspaceId/monitors")
			c.SetParamNames("workspaceId")
			c.SetParamValues(helper.UintToString(tt.workspaceId))

			// Setup the user session in the context
			helper.SetupSession(c)

			// Invoke the handler and verify the response
			err := monitorController.FindAll(c)
			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			// Unmarshal and validate the response JSON
			var monitors []dto.MonitorDTO
			err = json.Unmarshal(rec.Body.Bytes(), &monitors)
			require.NoError(t, err, "Unable to unmarshal response JSON")
			assert.Equal(t, tt.expectedCount, len(monitors))
			assert.Equal(t, tt.expectedDtos, monitors)
		})
	}
}

func TestMonitorController_FindById(t *testing.T) {
	helper.SetupDb() // Set up the test database
	e := echo.New()

	// Initialize service and controller
	monitorService := service.MonitorService{
		Repository: repository.MonitorRepository{},
	}
	monitorController := MonitorController{monitorService: monitorService}

	// Define the GET endpoint to be tested
	e.GET("/workspaces/:workspaceId/monitors/:monitorId", monitorController.FindById)

	// Create a sample monitor for testing purposes
	createMonitorDTO := getCreateMonitorDTO()
	monitor, err := monitorService.CreateMonitor(*createMonitorDTO, 1)
	require.NoError(t, err)

	// Test cases
	tests := []struct {
		name            string
		monitorId       uint
		workspaceId     string
		expectedCode    int
		expectedMonitor *model.Monitor
	}{
		{
			name:            "Monitor in workspace 1",
			monitorId:       monitor.ID,
			workspaceId:     "1",
			expectedCode:    http.StatusOK,
			expectedMonitor: monitor,
		},
		{
			name:            "Monitor in workspace 2",
			monitorId:       monitor.ID,
			workspaceId:     "2",
			expectedCode:    http.StatusForbidden,
			expectedMonitor: nil,
		},
		{
			name:            "Monitor not found",
			monitorId:       999,
			workspaceId:     "1",
			expectedCode:    http.StatusNotFound,
			expectedMonitor: nil,
		},
	}

	// Execute each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request and context
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/workspaces/:workspaceId/monitors/:monitorId")
			c.SetParamNames("workspaceId", "monitorId")
			c.SetParamValues(tt.workspaceId, helper.UintToString(tt.monitorId))

			// Setup a user session in the context
			helper.SetupSession(c)

			// Invoke the handler and verify the response
			err := monitorController.FindById(c)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedCode, rec.Code)

			if tt.expectedMonitor != nil {
				var monitorDto dto.MonitorDTO
				err := json.Unmarshal(rec.Body.Bytes(), &monitorDto)
				require.NoError(t, err)
				assert.Equal(t, mapMonitorToDTO(tt.expectedMonitor), monitorDto)
			}
		})
	}
}
