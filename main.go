package main

import (
	"capstone/config"
	adminai "capstone/controllers/admin_ai"
	"capstone/controllers/auth"
	"capstone/controllers/category"
	"capstone/controllers/comment"
	complaintsController "capstone/controllers/complaints"
	customerservice "capstone/controllers/customer_service"
	"capstone/controllers/feedbacks"
	"capstone/controllers/news"
	"capstone/middlewares"
	adminaiRepositories "capstone/repositories/admin_ai"
	AuthRepositories "capstone/repositories/auth"
	categoryRepositories "capstone/repositories/category"
	commentRepositories "capstone/repositories/comment"
	complaintsRepo "capstone/repositories/complaints"
	csRepositories "capstone/repositories/customer_service"
	feedbackRepositories "capstone/repositories/feedbacks"
	newsRepositories "capstone/repositories/news"
	"capstone/routes"
	adminaiService "capstone/services/admin_ai"
	AuthServices "capstone/services/auth"
	categoryService "capstone/services/category"
	commentService "capstone/services/comment"
	complaintsService "capstone/services/complaints"
	csService "capstone/services/customer_service"
	feedbackService "capstone/services/feedbacks"
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
	e.Use(middleware.Logger())
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	feedbackRepo := feedbackRepositories.NewFeedbackRepository(db)
	feedbackService := feedbackService.NewFeedbackService(feedbackRepo)
	feedbackController := feedbacks.NewFeedbackController(feedbackService)

	customerServiceRepo := csRepositories.NewCustomerServiceseRepo(db)
	customerService := csService.NewCustomerService(customerServiceRepo)
	customerServiceController := customerservice.NewCustomerServiceController(customerService)

	categoryRepo := categoryRepositories.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepo)
	categoryController := category.NewCategoryController(categoryService)

	adminAIRepo := adminaiRepositories.NewCustomerServiceseRepo(db)
	adminAIService := adminaiService.NewCustomerService(adminAIRepo)
	adminAIController := adminai.NewCustomerServiceController(adminAIService, complaintService, *serviceAdmin)

	// Mendaftarkan routes
	routeController := routes.RouteController{
		AuthController:            *authController,
		ComplaintController:       *complaintController,
		NewsController:            *newsController,
		CommentController:         *commentController,
		FeedbackController:        *feedbackController,
		CustomerServiceController: *customerServiceController,
		AuthAdminController:       *authControllerAdmin,
		AdminAISuggestion:         *adminAIController,
		CategoryController:        *categoryController,
	}
	routeController.RegisterRoutes(e)

	// Menjalankan server pada port 8000
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
