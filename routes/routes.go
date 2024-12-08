package routes

import (
	"capstone/controllers/auth"
	"capstone/controllers/comment"
	"capstone/controllers/complaints"
	feedback "capstone/controllers/feedbacks"
	"capstone/controllers/news"
	"capstone/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController      auth.AuthController
	ComplaintController complaints.ComplaintController
	NewsController      news.NewsController
	CommentController   comment.CommentController
	FeedbackController  feedback.FeedbackController
	jwtUser             middlewares.JwtUser
	AuthAdminController auth.AdminController
	jwtAdmin            middlewares.JwtAdmin
}

// RegisterRoutes mengatur semua rute untuk aplikasi
func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	// endpoint user
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/verify-otp", rc.AuthController.VerifyOTPController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	// endpoint complaints
	eComplaint := eJwt.Group("/complaint")
	eComplaint.Use(rc.jwtUser.GetUserID)
	eComplaint.POST("", rc.ComplaintController.CreateComplaintController)
	eComplaint.GET("", rc.ComplaintController.GetAllComplaintsByUser)
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

	eFeedback := eJwt.Group("/feedback")
	eFeedback.Use(rc.jwtUser.GetUserID)
	eFeedback.GET("/complaint/:complaint_id", rc.FeedbackController.GetFeedbackByComplaint)
	eFeedback.GET("", rc.FeedbackController.GetFeedbacksByUser)
	eFeedback.POST("/:id/response", rc.FeedbackController.AddResponseToFeedback)

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

	// Rute Admin untuk Kelola Complaints
	eAdminComplaints := eAdminJwt.Group("/complaint")
	eAdminComplaints.GET("/filter", rc.ComplaintController.GetComplaintsByStatusAndCategory)
	eAdminComplaints.GET("/:id", rc.ComplaintController.GetComplaintDetailByAdmin)
	eAdminComplaints.PUT("/status/:id", rc.ComplaintController.UpdateComplaintStatus)
	eAdminComplaints.POST("/feedback", rc.FeedbackController.ProvideFeedback)
	eAdminComplaints.PUT("/:id", rc.ComplaintController.UpdateComplaintByAdmin)
	eAdminComplaints.DELETE("/:id", rc.ComplaintController.DeleteComplaintByAdmin)

	// Rute Admin untuk Kelola news
	eAdminNews := eAdminJwt.Group("/news")
	eAdminNews.GET("", rc.NewsController.GetAllNewsWithComments)
	eAdminNews.GET("/:id", rc.NewsController.GetNewsDetailByAdmin)
	eAdminNews.POST("", rc.NewsController.AddNews)
	eAdminNews.PUT("/:id", rc.NewsController.UpdateNewsByAdmin)
	eAdminNews.DELETE("/bulk-delete", rc.NewsController.DeleteMultipleNewsByAdmin)

}
