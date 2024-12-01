package routes

import (
	auth "capstone/controllers"
	"capstone/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController auth.AuthController
	// jwtUser        middlewares.JwtUser
	AuthAdminController auth.AdminController
	// jwtAdmin middlewares.JWTAdminClaims
}

// RegisterRoutes mengatur semua rute untuk aplikasi
func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/verify-otp", rc.AuthController.VerifyOTPController)

	// Auth Routes for Admin
	e.POST("/register", rc.AuthAdminController.RegisterAdminHandler)      // Admin registration
	e.POST("/login", rc.AuthAdminController.LoginAdminHandler)

	// Grup API
	api := e.Group("/api/v1")
	api.GET("/admins", rc.AuthAdminController.GetAllAdminsHandler)        // Get all admins
	api.GET("/admin/:id", rc.AuthAdminController.GetAdminByIDHandler)     // Get admin by ID
	api.PUT("/admin", rc.AuthAdminController.UpdateAdminHandler)         // Update admin details
	api.DELETE("/admin/:id", rc.AuthAdminController.DeleteAdminHandler)  

	// JWT Authentication Middleware for protected routes
	api.Use(middlewares.JWTMiddleware())
}
