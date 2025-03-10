package comment

import (
	"capstone/controllers/comment/request"
	"capstone/controllers/comment/response"
	"capstone/services/comment"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	commentService comment.CommentServiceInterface
}

func NewCommentController(service comment.CommentServiceInterface) *CommentController {
	return &CommentController{commentService: service}
}

func (cc *CommentController) AddComment(c echo.Context) error {
	req := request.AddCommentRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	// Ambil UserID dari middleware
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Konversi request ke entitas
	commentEntity := req.ToEntity(userID)

	// Tambahkan komentar melalui service
	comment, err := cc.commentService.AddComment(commentEntity)
	if err != nil {
		// Menangkap error yang lebih spesifik dari service
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Comment added successfully",
		"comment": response.FromEntityComment(comment),
	})
}

func (cc CommentController) GetCommentsByUser(c echo.Context) error {
	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil query parameter page dan limit
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10 // Default limit
	}

	// Ambil data komentar berdasarkan user_id dari service
	comments, total, err := cc.commentService.GetCommentsByUserID(userID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve comments",
		})
	}

	// Hitung total halaman
	totalPages := (total + limit - 1) / limit

	// Konversi response
	commentResponses := response.FromEntityComments(comments)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Comments retrieved successfully",
		"user_id":    userID,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"comments":   commentResponses,
	})
}

func (cc *CommentController) GetAllComments(c echo.Context) error {
	// Ambil query parameter page dan limit
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10 // Default limit
	}

	// Ambil semua komentar dari service dengan pagination
	comments, total, err := cc.commentService.GetAllComments(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve all comments",
		})
	}

	// Hitung total halaman
	totalPages := (total + limit - 1) / limit

	// Konversi komentar ke response
	commentResponses := response.FromEntityComments(comments)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "All comments retrieved successfully",
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"comments":   commentResponses,
	})
}

func (cc *CommentController) GetCommentByID(c echo.Context) error {
	// Ambil ID dari parameter URL
	commentID := c.Param("id")

	// Ambil detail komentar dari service
	comment, err := cc.commentService.GetCommentByID(commentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Komentar Tidak Ditemukan",
		})
	}

	// Konversi komentar ke response
	commentResponse := response.FromEntityComment(comment)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Comment retrieved successfully",
		"comment": commentResponse,
	})
}
