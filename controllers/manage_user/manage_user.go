package manageuser

import (
	"capstone/controllers/manage_user/response"
	"capstone/middlewares"
	manageuser "capstone/services/manage_user"
	"net/http"

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

	// Ambil data semua user dari service
	users, err := controller.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve users",
		})
	}

	// Konversi ke response
	userResponses := response.FromEntitiesUsers(users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users retrieved successfully",
		"data":    userResponses,
	})
}
