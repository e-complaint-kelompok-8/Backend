package comment

import (
	"capstone/middlewares"
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete comments",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Comments deleted successfully",
	})
}
