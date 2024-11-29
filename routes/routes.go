package routes

import (
	"capstone/controllers"
	authRepository "capstone/repositories/auth"
	authService "capstone/services/auth"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes mengatur semua rute untuk aplikasi
func RegisterRoutes(e *echo.Echo) {
	// Inisialisasi dependency untuk Auth
	repo := authRepository.NewAuthRepository()
	service := authService.NewAuthService(repo)
	authController := controllers.NewAuthController(service)

	// Grup API
	api := e.Group("/api/v1")

	// Route Auth
	api.POST("/login", authController.Login)
}