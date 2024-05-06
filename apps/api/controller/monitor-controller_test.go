package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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
			expectedDtos:  []dto.MonitorDTO{convertMonitorModelToDto(monitor)},
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
				assert.Equal(t, convertMonitorModelToDto(tt.expectedMonitor), monitorDto)
			}
		})
	}
}

func Test_convertMonitorDtoToModel(t *testing.T) {
	//The DTO we want to convert
	monitorDTO := dto.MonitorDTO{
		ID:       1,
		Name:     "Test Monitor",
		Url:      "https://example.com",
		Interval: 60,
		Method:   "GET",
		Retries:  3,
		Timeout:  10,
	}
	//The expected model after conversion
	monitorModel := model.Monitor{
		Base:        model.Base{ID: 1},
		Name:        monitorDTO.Name,
		Url:         monitorDTO.Url,
		Interval:    monitorDTO.Interval,
		Retries:     monitorDTO.Retries,
		Timeout:     monitorDTO.Timeout,
		Method:      monitorDTO.Method,
		WorkspaceID: 0,
	}

	type args struct {
		dto       *dto.MonitorDTO
		workspace uint
	}
	tests := []struct {
		name string
		args args
		want model.Monitor
	}{
		{name: "Convert DTO to model", args: args{
			dto: &monitorDTO},
			want: monitorModel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertMonitorDtoToModel(tt.args.dto, tt.args.workspace), "convertMonitorDtoToModel(%v, %v)", tt.args.dto, tt.args.workspace)
		})
	}
}

func Test_convertMonitorModelToDto(t *testing.T) {
	//The model we want to convert
	monitorModel := model.Monitor{
		Base:        model.Base{ID: 1},
		Name:        "Test Monitor",
		Url:         "https://example.com",
		Interval:    60,
		Method:      "GET",
		Retries:     3,
		Timeout:     10,
		WorkspaceID: 0,
	}

	//The expected DTO after conversion
	monitorDTO := dto.MonitorDTO{
		ID:       1,
		Name:     "Test Monitor",
		Url:      "https://example.com",
		Interval: 60,
		Method:   "GET",
		Retries:  3,
		Timeout:  10,
	}

	type args struct {
		monitor *model.Monitor
	}
	tests := []struct {
		name string
		args args
		want dto.MonitorDTO
	}{
		{name: "Convert model to DTO", args: args{
			monitor: &monitorModel},
			want: monitorDTO,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertMonitorModelToDto(tt.args.monitor), "convertMonitorModelToDto(%v)", tt.args.monitor)
		})
	}
}
func TestValidateMonitorDTO(t *testing.T) {
	validDto := getCreateMonitorDTO()
	invalidDto := dto.CreateMonitorDTO{
		Name:     "Test Monitor",
		Url:      "example.com",
		Method:   "GETS",
		Timeout:  -1,
		Interval: 0,
		Retries:  -1,
	}

	testCases := []struct {
		name    string
		input   dto.CreateMonitorDTO
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Valid DTO",
			input:   *validDto,
			wantErr: assert.NoError,
		},
		{
			name:    "Invalid DTO",
			input:   invalidDto,
			wantErr: assert.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateMonitorDTO(tc.input)
			tc.wantErr(t, err, fmt.Sprintf("validateMonitorDTO(%v)", tc.input))

			if err != nil {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					// Collect validation error fields for checks
					seenFields := map[string]bool{}
					for _, vErr := range validationErrors {
						seenFields[vErr.Field()] = true
					}

					// Check all expected fields are included
					expectedFields := []string{"Url", "Method", "Timeout", "Interval", "Retries"}
					for _, field := range expectedFields {
						assert.Contains(t, seenFields, field, "Expected validation error for field %s", field)
					}
				} else {
					t.Fatalf("Expected validator.ValidationErrors, but got: %v", err)
				}
			}
		})
	}
}
