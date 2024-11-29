package controllers

import (
	"net/http"
	"capstone/controllers/auth/request"
	"capstone/controllers/auth/response"
	"capstone/services/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Login(c echo.Context) error {
	var loginRequest request.LoginRequest

	// Bind request data
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid input"})
	}

	// Validasi menggunakan Validator Echo
	if err := c.Validate(&loginRequest); err != nil {
    	return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Validation failed"})
	}

	// Validate email and password
	admin, err := ac.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.LoginResponse{Message: "Login successful", Admin: *admin})
}