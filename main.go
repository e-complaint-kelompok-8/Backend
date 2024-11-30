package main

import (
	"capstone/config"
	auth "capstone/controllers"
	"capstone/middlewares"
	AuthRepositories "capstone/repositories/auth"
	"capstone/routes"
	AuthServices "capstone/services/auth"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	LoadEnv()
	// Connect to the database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Jalankan migrasi database
	config.RunMigrations(db)

	// Membuat instance Echo
	e := echo.New()
	e.Use(middleware.CORS())

	// Inisialisasi dependency untuk Auth
	jwtUser := middlewares.JwtUser{}
	repo := AuthRepositories.NewAuthRepository(db)
	service := AuthServices.NewAuthService(repo, jwtUser)
	authController := auth.NewAuthController(service)

	// Mendaftarkan routes
	routeController := routes.RouteController{
		AuthController: *authController,
	}
	routeController.RegisterRoutes(e)

	// Menjalankan server pada port 8080
	log.Println("Server starting on port 8080...")
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
