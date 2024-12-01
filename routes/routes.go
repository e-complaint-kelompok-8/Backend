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
	e.POST("/register", rc.AuthAdminController.RegisterAdminHandler)      
	e.POST("/login", rc.AuthAdminController.LoginAdminHandler)

	// Grup API
	api := e.Group("/admin")
	api.GET("", rc.AuthAdminController.GetAllAdminsHandler)       
	api.GET("/:id", rc.AuthAdminController.GetAdminByIDHandler) 
	api.PUT("", rc.AuthAdminController.UpdateAdminHandler)        
	api.DELETE("/:id", rc.AuthAdminController.DeleteAdminHandler)  

	// JWT Authentication Middleware for protected routes
	api.Use(middlewares.JWTMiddleware())
}
