package comment

import (
	"capstone/controllers/comment/response"
	"capstone/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (cc *CommentController) DeleteComments(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var ids []int
	if err := c.Bind(&ids); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	// Validasi input
	if len(ids) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Comment IDs are required",
		})
	}

	// Hapus komentar melalui service
	err = cc.commentService.DeleteComments(ids)
	if err != nil {
		if err.Error() == "some comment IDs do not exist" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Beberapa ID komentar tidak ada",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete comments",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Comments deleted successfully",
	})
}

func (cc *CommentController) GetCommentsByUserID(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil user_id dari parameter URL
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID format. It should be an integer.",
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

	// Ambil data komentar berdasarkan user_id dari service dengan pagination
	comments, total, err := cc.commentService.GetCommentsByUserID(userID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve comments",
		})
	}

	// Hitung total halaman
	totalPages := (total + limit - 1) / limit

	// Konversi komentar ke response
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
