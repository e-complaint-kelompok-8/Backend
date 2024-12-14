package routes

import (
	adminai "capstone/controllers/admin_ai"
	"capstone/controllers/auth"
	"capstone/controllers/category"
	"capstone/controllers/comment"
	"capstone/controllers/complaints"
	customerservice "capstone/controllers/customer_service"
	feedback "capstone/controllers/feedbacks"
	manageuser "capstone/controllers/manage_user"
	"capstone/controllers/news"
	"capstone/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController            auth.AuthController
	ComplaintController       complaints.ComplaintController
	NewsController            news.NewsController
	CommentController         comment.CommentController
	FeedbackController        feedback.FeedbackController
	jwtUser                   middlewares.JwtUser
	CustomerServiceController customerservice.CustomerServiceController
	AuthAdminController       auth.AdminController
	jwtAdmin                  middlewares.JwtAdmin
	ManageUserController      manageuser.ManageUserController
	AdminAISuggestion         adminai.AdminAIController
	CategoryController        category.CategoryController
}

// RegisterRoutes mengatur semua rute untuk aplikasi
func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	// endpoint user
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/verify-otp", rc.AuthController.VerifyOTPController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	// endpoint user profile
	eUserProfile := eJwt.Group("/user")
	eUserProfile.Use(rc.jwtUser.GetUserID)
	eUserProfile.GET("/profile", rc.AuthController.GetProfile)
	eUserProfile.PUT("/profile/name", rc.AuthController.UpdateName)         // Endpoint untuk memperbarui nama
	eUserProfile.PUT("/profile/photo", rc.AuthController.UpdatePhoto)       // Endpoint untuk memperbarui foto
	eUserProfile.PUT("/profile/password", rc.AuthController.UpdatePassword) // Endpoint untuk memperbarui password

	// endpoint complaints
	eComplaint := eJwt.Group("/complaint")
	eComplaint.Use(rc.jwtUser.GetUserID)
	eComplaint.POST("", rc.ComplaintController.CreateComplaintController)
	eComplaint.GET("", rc.ComplaintController.GetUserComplaintsByStatusAndCategory)
	// eComplaint.GET("", rc.ComplaintController.GetAllComplaintsByUser)
	eComplaint.GET("/:id", rc.ComplaintController.GetComplaintById)
	eComplaint.GET("/user", rc.ComplaintController.GetComplaintByUser)
	eComplaint.GET("/status/:status", rc.ComplaintController.GetComplaintsByStatus)
	eComplaint.GET("/category/:category_id", rc.ComplaintController.GetComplaintsByCategory)
	eComplaint.PUT("/:id/cancel", rc.ComplaintController.CancelComplaint)

	// endpoint news
	eNews := eJwt.Group("/news")
	eNews.GET("", rc.NewsController.GetAllNews)
	eNews.GET("/:id", rc.NewsController.GetNewsByID)

	// endpoint comment
	eComment := eJwt.Group("/comment")
	eComment.Use(rc.jwtUser.GetUserID)
	eComment.POST("", rc.CommentController.AddComment)
	eComment.GET("/user", rc.CommentController.GetCommentsByUser)
	eComment.GET("", rc.CommentController.GetAllComments)
	eComment.GET("/:id", rc.CommentController.GetCommentByID)

	// grup feedback
	eFeedback := eJwt.Group("/feedback")
	eFeedback.Use(rc.jwtUser.GetUserID)
	eFeedback.GET("/complaint/:complaint_id", rc.FeedbackController.GetFeedbackByComplaint)
	eFeedback.GET("", rc.FeedbackController.GetFeedbacksByUser)
	eFeedback.POST("/:id/response", rc.FeedbackController.AddResponseToFeedback)

	// grup customer service (ai)
	eCustomerService := eJwt.Group("/chatbot")
	eCustomerService.Use(rc.jwtUser.GetUserID)
	eCustomerService.POST("", rc.CustomerServiceController.ChatbotQueryController)
	eCustomerService.GET("/user-responses", rc.CustomerServiceController.GetUserResponses)

	// Grup Admin
	eAdmin := e.Group("/admin")
	eAdmin.POST("/register", rc.AuthAdminController.RegisterAdminHandler)
	eAdmin.POST("/login", rc.AuthAdminController.LoginAdminHandler)

	// Rute Admin yang dilindungi JWTAdminMiddleware
	eAdminJwt := eAdmin.Group("")
	eAdminJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	eAdminJwt.Use(rc.jwtAdmin.JWTAdminMiddleware)
	eAdminJwt.GET("", rc.AuthAdminController.GetAllAdminsHandler)
	eAdminJwt.GET("/:id", rc.AuthAdminController.GetAdminByIDHandler)
	eAdminJwt.PUT("/:id", rc.AuthAdminController.UpdateAdminHandler)
	eAdminJwt.DELETE("/:id", rc.AuthAdminController.DeleteAdminHandler)
	eAdminJwt.GET("/profile", rc.AuthAdminController.GetAdminProfile)
	eAdminJwt.PUT("/profile", rc.AuthAdminController.UpdateAdminProfile)

	// Rute Admin untuk Kelola Complaints
	eAdminComplaints := eAdminJwt.Group("/complaint")
	eAdminComplaints.GET("/filter", rc.ComplaintController.GetComplaintsByStatusAndCategory)
	eAdminComplaints.GET("/:id", rc.ComplaintController.GetComplaintDetailByAdmin)
	eAdminComplaints.POST("/feedback", rc.FeedbackController.ProvideFeedback)
	eAdminComplaints.PUT("/:id", rc.ComplaintController.UpdateComplaintByAdmin)
	eAdminComplaints.DELETE("/:id", rc.ComplaintController.DeleteComplaintByAdmin)

	// Rute Admin untuk Kelola news
	eAdminNews := eAdminJwt.Group("/news")
	eAdminNews.GET("", rc.NewsController.GetAllNewsWithComments)
	eAdminNews.GET("/:id", rc.NewsController.GetNewsDetailByAdmin)
	eAdminNews.POST("", rc.NewsController.AddNews)
	eAdminNews.PUT("/:id", rc.NewsController.UpdateNewsByAdmin)
	eAdminNews.GET("/:news_id/comment", rc.CommentController.GetCommentsByNewsID)
	eAdminNews.DELETE("/bulk-delete", rc.NewsController.DeleteMultipleNewsByAdmin)

	// Rute Admin untuk Kelola Feedback
	eAdminFeedback := eAdminJwt.Group("/feedback")
	eAdminFeedback.PUT("/:id", rc.FeedbackController.UpdateFeedback)

	// Rute Admin untuk Kelola Category
	eAdminCategory := eAdminJwt.Group("/category")
	eAdminCategory.GET("", rc.CategoryController.GetAllCategories)
	eAdminCategory.GET("/:id", rc.CategoryController.GetCategoryByID)
	eAdminCategory.POST("", rc.CategoryController.CreateCategory)
	eAdminCategory.PUT("/:id", rc.CategoryController.UpdateCategory)
	eAdminCategory.DELETE("/:id", rc.CategoryController.DeleteCategory)

	// Rute Admin untuk kelola comment
	eAdminComment := eAdminJwt.Group("/comment")
	eAdminComment.DELETE("", rc.CommentController.DeleteComments)
	eAdminComment.GET("", rc.CommentController.GetAllComments)
	eAdminComment.GET("/:id", rc.CommentController.GetCommentByID)
	eAdminComment.GET("/user/:user_id", rc.CommentController.GetCommentsByUserID)

	// Rute Admin untuk kelola user
	eAdminManageUser := eAdminJwt.Group("/users")
	eAdminManageUser.GET("", rc.ManageUserController.GetAllUsers)
	eAdminManageUser.GET("/:id", rc.ManageUserController.GetUserDetail)

	// Rute Admin untuk kelola chatbot
	eAdminAI := eAdminJwt.Group("/ai-suggestions")
	eAdminAI.POST("", rc.AdminAISuggestion.GetAISuggestion)
	eAdminAI.POST("/:id/follow-up", rc.AdminAISuggestion.FollowUpAISuggestion)
	eAdminAI.GET("", rc.AdminAISuggestion.GetAllAISuggestions)
}
