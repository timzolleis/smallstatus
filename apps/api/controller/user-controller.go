package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"status/dto"
	"status/helper"
	"status/model"
	_ "status/repository"
	"status/service"
	"strconv"
)

type UserController struct {
	Service service.UserService
}

func (controller *UserController) FindAll(c echo.Context) error {
	users, err := controller.Service.FindAll()
	if err != nil {
		return helper.HandleError(err, c)
	}
	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = mapToDTO(user)
	}

	return c.JSON(http.StatusOK, userDTOs)
}

func (controller *UserController) FindById(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := controller.Service.FindUserById(id)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusOK, mapToDTO(*user))
}

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (controller *UserController) Create(c echo.Context) error {
	var body CreateUserBody
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	user, err := controller.Service.CreateUser(body.Name, body.Email, body.Password)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.JSON(http.StatusCreated, mapToDTO(*user))
}

func (controller *UserController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := controller.Service.Delete(id)
	if err != nil {
		return helper.HandleError(err, c)
	}
	return c.NoContent(http.StatusNoContent)
}

func mapToDTO(user model.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
