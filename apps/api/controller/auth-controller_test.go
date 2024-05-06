package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timzolleis/smallstatus/controller"
	"github.com/timzolleis/smallstatus/service"
)

func TestAuthController_SignUp(t *testing.T) { //Setup
	helper.SetupDb()

	// Arrange
	userService := service.UserService{}
	authController := controller.AuthController{UserService: userService}
	e := echo.New()
	e.POST("/api/auth/signup", authController.SignUp)

	createUserDTO := dto.CreateUserDTO{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	requestBody, _ := json.Marshal(createUserDTO)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	authController.SignUp(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	//Check if user was created
	user, _ := userService.FindByEmail("test@example.com")
	assert.Equal(t, "Test User", user.Name)
}

func TestAuthController_Login(t *testing.T) {
	// Setup
	helper.SetupDb()

	userService := service.UserService{}
	authController := controller.AuthController{UserService: userService}
	e := echo.New()
	e.POST("/api/auth/login", authController.Login)

	hash, _ := helper.HashPassword("password")
	user := model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: hash,
	}
	userService.Create(&user)

	tests := []struct {
		name     string
		password string
		wantCode int
	}{
		{"Correct password", "password", http.StatusOK},
		{"Wrong password", "wrongpassword", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginDTO := dto.LoginDTO{
				Email:    "test@example.com",
				Password: tt.password,
			}
			body, _ := json.Marshal(loginDTO)
			req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))
			authController.Login(c)
			assert.Equal(t, tt.wantCode, rec.Code)
		})
	}
}
