package helper

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func HandleError(err error, c echo.Context) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code := http.StatusNotFound
		return c.JSON(code, NewErrorResponse(err.Error(), code))
	}
	code := http.StatusInternalServerError

	return c.JSON(code, NewErrorResponse(err.Error(), code))
}

func InvalidRequest(c echo.Context) error {
	code := http.StatusBadRequest
	return c.JSON(code, NewErrorResponse("Invalid request body", code))
}
