package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
)

func WorkspaceMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*model.User)
		workspaceIdString := c.Param("workspaceId")
		if workspaceIdString == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Workspace ID is required")
		}
		workspaceId := helper.StringToUint(workspaceIdString)
		workspaceService := service.WorkspaceService{}
		if !workspaceService.IsPartOfWorkspace(user.ID, workspaceId) {
			return echo.NewHTTPError(http.StatusUnauthorized, "You are not part of this workspace")
		}
		return next(c)
	}
}
