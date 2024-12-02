package routes

import (
	"capstone/controllers/auth"
	"capstone/controllers/complaints"
	"capstone/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController      auth.AuthController
	ComplaintController complaints.ComplaintController
	jwtUser             middlewares.JwtUser
	AuthAdminController auth.AdminController
	// jwtAdmin middlewares.JWTAdminClaims
}

// RegisterRoutes mengatur semua rute untuk aplikasi
func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	// end point user
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/verify-otp", rc.AuthController.VerifyOTPController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eComplaint := eJwt.Group("/complaint")
	eComplaint.Use(rc.jwtUser.GetUserID)
	eComplaint.POST("", rc.ComplaintController.CreateComplaintController)
	eComplaint.GET("/:id", rc.ComplaintController.GetComplaintById)
	eComplaint.GET("/user", rc.ComplaintController.GetComplaintByUser)
	eComplaint.GET("/status/:status", rc.ComplaintController.GetComplaintsByStatus)

	// Grup Admin
	api := e.Group("/admin")
	api.POST("/register", rc.AuthAdminController.RegisterAdminHandler)
	api.POST("/login", rc.AuthAdminController.LoginAdminHandler)
	api.GET("", rc.AuthAdminController.GetAllAdminsHandler)
	api.GET("/:id", rc.AuthAdminController.GetAdminByIDHandler)
	api.PUT("", rc.AuthAdminController.UpdateAdminHandler)
	api.DELETE("/:id", rc.AuthAdminController.DeleteAdminHandler)

	// JWT Authentication Middleware for protected routes
	api.Use(middlewares.JWTMiddleware())
}
