package manageuser

import (
	"capstone/controllers/manage_user/response"
	"capstone/middlewares"
	manageuser "capstone/services/manage_user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ManageUserController struct {
	userService manageuser.UserServiceInterface
}

func NewManageUserController(service manageuser.UserServiceInterface) *ManageUserController {
	return &ManageUserController{userService: service}
}

func (controller *ManageUserController) GetAllUsers(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
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

	// Ambil data user dari service dengan pagination
	users, total, err := controller.userService.GetAllUsers(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve users",
		})
	}

	// Hitung total halaman
	totalPages := (total + limit - 1) / limit

	// Konversi ke response
	userResponses := response.FromEntitiesUsers(users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Users retrieved successfully",
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"data":       userResponses,
	})
}

func (controller *ManageUserController) GetUserDetail(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil ID user dari parameter
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	// Ambil detail user dari service
	user, err := controller.userService.GetUserDetail(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Konversi ke response
	userResponse := response.FromEntityUser(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User retrieved successfully",
		"data":    userResponse,
	})
}
