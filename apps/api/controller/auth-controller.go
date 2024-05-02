package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/middleware"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/service"
	"net/http"
)

type AuthController struct {
	UserService service.UserService
}

func (controller *AuthController) SignUp(c echo.Context) error {
	var createUserDTO = dto.CreateUserDTO{}
	if err := c.Bind(&createUserDTO); err != nil {
		return helper.InvalidRequest(c)
	}

	if err := validateCreateUserDTO(createUserDTO); err != nil {
		return helper.InvalidRequest(c)
	}

	hashedPassword, err := helper.HashPassword(createUserDTO.Password)
	if err != nil {
		return helper.HandleError(err, c)
	}

	user := &model.User{
		Name:     createUserDTO.Name,
		Email:    createUserDTO.Email,
		Password: hashedPassword,
	}

	if _, err := controller.UserService.Create(user); err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, helper.NewSuccessResponse("User created"))
}

func validateCreateUserDTO(createUserDTO dto.CreateUserDTO) error {
	validate := validator.New()
	err := validate.Struct(createUserDTO)
	if err != nil {
		return err
	}
	return nil
}

func (controller *AuthController) Login(c echo.Context) error {
	var loginDTO = dto.LoginDTO{}
	if err := c.Bind(&loginDTO); err != nil {
		return helper.InvalidRequest(c)
	}
	if err := validateLoginDTO(loginDTO); err != nil {
		return helper.InvalidRequest(c)
	}
	user, err := controller.UserService.FindByEmail(loginDTO.Email)
	if err != nil {
		return helper.HandleError(err, c)
	}
	if isValidPassword := helper.CheckPassword(loginDTO.Password, user.Password); !isValidPassword {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse("invalid_credentials", http.StatusBadRequest))
	}
	sess := createLoginSession(c, user)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.NewErrorResponse("session_error", http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, helper.NewSuccessResponse("Logged in"))
}

func validateLoginDTO(loginDTO dto.LoginDTO) error {
	validate := validator.New()
	err := validate.Struct(loginDTO)
	if err != nil {
		return err
	}
	return nil
}

func createLoginSession(c echo.Context, user *model.User) *sessions.Session {
	sess, _ := session.Get(middleware.SessionName, c)
	sess.Values["user_id"] = user.ID
	sess.Values["user_email"] = user.Email
	return sess
}
