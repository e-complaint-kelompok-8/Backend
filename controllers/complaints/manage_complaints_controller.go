package complaints

import (
	"capstone/controllers/complaints/request"
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
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	// Konversi categoryID, page, dan limit jika ada
	var categoryID, page, limit int
	if categoryIDParam != "" {
		categoryID, err = strconv.Atoi(categoryIDParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
		}
	}
	page, _ = strconv.Atoi(pageParam)   // Default 0 jika kosong
	limit, _ = strconv.Atoi(limitParam) // Default 0 jika kosong

	// Ambil data dari service
	complaints, total, err := cc.complaintService.GetComplaintsByStatusAndCategory(status, categoryID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Konversi entities ke response
	responseData := response.ComplaintsFromEntities(complaints)

	// Kirim respons dengan data dan metadata pagination
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data": map[string]interface{}{
			"complaints": responseData,
			"total":      total,
			"page":       page,
			"limit":      limit,
		},
	})
}

func (cc *ComplaintController) GetComplaintDetailByAdmin(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil complaint ID dari parameter
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid complaint ID"})
	}

	// Ambil detail complaint dari service
	complaint, err := cc.complaintService.GetComplaintDetailByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Konversi ke response dengan feedback
	response := response.ComplaintFromEntitiesWithFeedback(complaint)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Complaint detail retrieved successfully",
		"complaint": response,
	})
}

func (cc *ComplaintController) UpdateComplaintByAdmin(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil AdminID dari token JWT
	adminID, ok := c.Get("admin_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid admin token"})
	}

	// Ambil ID pengaduan dari parameter
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid complaint ID"})
	}

	// Bind data dari request body
	request := request.RequestUpdateComplaint{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// Masukkan AdminID ke dalam data pembaruan
	updateData := request.ToEntity()

	// Pastikan AdminID diinisialisasi jika nil
	if updateData.AdminID == nil {
		updateData.AdminID = new(int) // Alokasikan memori untuk pointer AdminID
	}
	*updateData.AdminID = adminID

	// Update data pengaduan melalui service
	err = cc.complaintService.UpdateComplaintByAdmin(complaintID, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Ambil data complaint yang telah diperbarui
	complaint, err := cc.complaintService.GetComplaintByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch updated complaint"})
	}

	// Kirim respons
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Complaint updated successfully",
		"complaint": response.ComplaintFromEntitiesWithAdmin(complaint),
	})
}

func (cc *ComplaintController) DeleteComplaintByAdmin(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil ID pengaduan dari parameter
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid complaint ID"})
	}

	// Hapus complaint melalui service
	err = cc.complaintService.DeleteComplaintByAdmin(complaintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Kirim respons
	return c.JSON(http.StatusOK, map[string]string{"message": "Complaint deleted successfully"})
}
