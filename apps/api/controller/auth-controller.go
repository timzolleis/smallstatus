package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
)

type AuthController struct {
	UserService service.UserService
}

func (controller *AuthController) SignUp(c echo.Context) error {
	var createUserDTO = dto.CreateUserDTO{}
	err := c.Bind(&createUserDTO)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	_, err = controller.UserService.CreateUser(createUserDTO.Name, createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, helper.NewSuccessResponse("User created"))
}

func (controller *AuthController) Login(c echo.Context) error {
	var loginDTO = dto.LoginDTO{}
	err := c.Bind(&loginDTO)
	if err != nil {
		return helper.InvalidRequest(c)
	}
	user, err := controller.UserService.FindUserByEmail(loginDTO.Email)
	if err != nil {
		return helper.HandleError(err, c)
	}
	isValidPassword := helper.CheckPassword(loginDTO.Password, user.Password)
	if !isValidPassword {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse("invalid_credentials", http.StatusBadRequest))
	}
	return c.String(http.StatusOK, "Alright")
}
