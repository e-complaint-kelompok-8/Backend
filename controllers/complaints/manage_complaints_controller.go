package complaints

import (
	"capstone/controllers/complaints/response"
	"capstone/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (cc *ComplaintController) GetComplaintsByStatusAndCategory(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil parameter query
	status := c.QueryParam("status")
	categoryIDParam := c.QueryParam("category_id")

	// Konversi categoryID jika ada, atau gunakan default 0
	var categoryID int
	if categoryIDParam != "" {
		categoryID, err = strconv.Atoi(categoryIDParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
		}
	}

	// Ambil data dari service
	complaints, err := cc.complaintService.GetComplaintsByStatusAndCategory(status, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Konversi entities ke response
	response := response.ComplaintsFromEntities(complaints)

	// Kirim respons
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": response,
	})
}
