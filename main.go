package main

import (
	"capstone/config"
	"capstone/controllers/auth"
	"capstone/controllers/comment"
	complaintsController "capstone/controllers/complaints"
	"capstone/controllers/news"
	"capstone/middlewares"
	AuthRepositories "capstone/repositories/auth"
	commentRepositories "capstone/repositories/comment"
	complaintsRepo "capstone/repositories/complaints"
	newsRepositories "capstone/repositories/news"
	"capstone/routes"
	AuthServices "capstone/services/auth"
	commentService "capstone/services/comment"
	complaintsService "capstone/services/complaints"
	newsService "capstone/services/news"
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

	// Inisialisasi dependency untuk create complaint
	complaintRepo := complaintsRepo.NewComplaintRepo(db)
	complaintService := complaintsService.NewComplaintService(complaintRepo)
	complaintController := complaintsController.NewComplaintController(complaintService)

	// Inisialisasi dependency untuk Auth Admin
	jwtAdmin := middlewares.JwtAdmin{}
	repoAdmin := AuthRepositories.NewAdminRepository(db)
	serviceAdmin := AuthServices.NewAdminService(repoAdmin, jwtAdmin)
	authControllerAdmin := auth.NewAdminController(serviceAdmin)

	newsRepo := newsRepositories.NewNewsRepository(db)
	newsService := newsService.NewNewsService(newsRepo)
	newsController := news.NewNewsController(newsService)

	commentRepo := commentRepositories.NewCommentRepository(db)
	commentService := commentService.NewCommentService(commentRepo)
	commentController := comment.NewCommentController(commentService)

	// Mendaftarkan routes
	routeController := routes.RouteController{
		AuthController:      *authController,
		ComplaintController: *complaintController,
		NewsController:      *newsController,
		CommentController:   *commentController,
		AuthAdminController: *authControllerAdmin,
	}
	routeController.RegisterRoutes(e)

	// Menjalankan server pada port 8080
	log.Println("Server starting on port 8000...")
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
