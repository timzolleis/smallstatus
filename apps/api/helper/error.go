package helper

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func HandleError(err error, c echo.Context) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.String(http.StatusInternalServerError, err.Error())
}

func InvalidRequest(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Invalid request")
}
