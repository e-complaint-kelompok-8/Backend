package auth

import (
	"capstone/controllers/auth/request"
	"capstone/controllers/auth/response"
	"capstone/entities"
	"capstone/middlewares"
	"capstone/services/auth"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminService *auth.AdminService
}

// NewAdminController creates a new instance of AdminController
func NewAdminController(adminService *auth.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

// RegisterAdminHandler handles admin registration
func (controller *AdminController) RegisterAdminHandler(c echo.Context) error {
	var admin entities.Admin

	// Bind JSON request body to admin struct
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request body"})
	}

	// Register admin using the service
	createdAdmin, err := controller.adminService.RegisterAdmin(admin)
	if err != nil {
		// Periksa error untuk memberikan pesan yang lebih spesifik
		if err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "Email Already Exists",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.RegisterAdminFromEntities(createdAdmin))
}

func (controller *AdminController) LoginAdminHandler(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON request body to credentials struct
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// Authenticate admin
	token, admin, err := controller.adminService.AuthenticateAdmin(credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
	}

	// Return token and admin info
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"admin": map[string]interface{}{
			"id":    admin.ID,
			"email": admin.Email,
			"role":  admin.Role,
		},
	})
}

// GetAllAdminsHandler handles retrieving all admins
func (controller *AdminController) GetAllAdminsHandler(c echo.Context) error {
	admins, err := controller.adminService.GetAllAdmins()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, admins)
}

// GetAdminByIDHandler handles retrieving an admin by ID
func (controller *AdminController) GetAdminByIDHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid admin ID"})
	}

	admin, err := controller.adminService.GetAdminByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, admin)
}

// UpdateAdminHandler handles updating an admin
func (controller *AdminController) UpdateAdminHandler(c echo.Context) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid admin ID"})
	}

	var admin entities.Admin

	// Bind JSON request body to admin struct
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request body"})
	}

	// Ensure the ID in the request matches the parameter ID
	admin.ID = id

	// Update admin using the service
	updatedAdmin, err := controller.adminService.UpdateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedAdmin)
}

// DeleteAdminHandler handles deleting an admin by ID
func (controller *AdminController) DeleteAdminHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid admin ID"})
	}

	// Delete admin using the service
	if err := controller.adminService.DeleteAdmin(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "admin deleted successfully"})
}

func (ac *AdminController) SomeAdminEndpoint(c echo.Context) error {
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Lanjutkan dengan logika endpoint
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome, Admin!"})
}

func (controller *AdminController) GetAdminProfile(c echo.Context) error {
	adminID, err := middlewares.ExtractAdminID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	admin, err := controller.adminService.GetAdminByID(adminID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Admin profile fetched successfully",
		"admin":   response.AdminProfileFromEntities(admin),
	})
}

func (controller *AdminController) UpdateAdminProfile(c echo.Context) error {
	adminID, err := middlewares.ExtractAdminID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	updateRequest := request.UpdateAdminRequest{}

	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	admin, err := controller.adminService.UpdateAdminProfile(adminID, updateRequest.Email, updateRequest.Password, updateRequest.Photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Admin profile updated successfully",
		"admin":   response.AdminProfileFromEntities(admin),
	})
}
