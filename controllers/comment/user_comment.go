package comment

import (
	"capstone/controllers/comment/request"
	"capstone/controllers/comment/response"
	"capstone/services/comment"
	"net/http"

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

	// Ambil data komentar berdasarkan user_id dari service
	comments, err := cc.commentService.GetCommentsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve comments",
		})
	}

	// Konversi response
	commentResponses := response.FromEntityComments(comments)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Comments retrieved successfully",
		"comments": commentResponses,
	})
}
