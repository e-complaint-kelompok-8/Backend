package routes

import (
	auth "capstone/controllers"
	"capstone/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController auth.AuthController
	jwtUser        middlewares.JwtUser
}

// RegisterRoutes mengatur semua rute untuk aplikasi
func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/verify-otp", rc.AuthController.VerifyOTPController)
	// Grup API
	api := e.Group("/api/v1")

	// Route Auth
	api.POST("/login", rc.AuthController.Login)
}
