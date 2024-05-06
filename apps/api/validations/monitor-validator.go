package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/timzolleis/smallstatus/service"
)

func ValidateHttpMethod(input validator.FieldLevel) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE"}
	for _, method := range validMethods {
		if input.Field().String() == method {
			return true
		}
	}
	return false
}

func IsMonitorPartOfWorkspace(monitorId uint, workspaceId uint) bool {
	monitorService := service.MonitorService{}
	monitor, err := monitorService.FindMonitorById(monitorId)
	if err != nil {
		return false
	}
	return monitor.WorkspaceID == workspaceId
}
