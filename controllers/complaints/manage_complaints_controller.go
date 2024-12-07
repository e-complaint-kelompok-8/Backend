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

	// Konversi ke response
	response := response.ComplaintFromEntitiesWithReason(complaint, complaint.Photos)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Complaint detail retrieved successfully",
		"complaint": response,
	})
}

func (cc *ComplaintController) UpdateComplaintStatus(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "Access denied",
		})
	}

	// Ambil admin ID dari middleware
	adminID, ok := c.Get("admin_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Admin tidak memiliki otorisasi",
		})
	}

	// Ambil complaint ID dari parameter
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID pengaduan tidak valid",
		})
	}

	// Ambil data status baru dari body
	var request struct {
		Status string `json:"status" validate:"required"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Data status tidak valid",
		})
	}

	// Proses pembaruan status melalui service
	err = cc.complaintService.UpdateComplaintStatus(complaintID, adminID, request.Status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Ambil data complaint terkini
	complaint, err := cc.complaintService.GetComplaintByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Gagal mengambil data pengaduan",
		})
	}

	// Kirim respons dengan data terkini
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Status pengaduan berhasil diperbarui",
		"complaint": response.ComplaintFromEntitiesWithAdmin(complaint),
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
	updateData.AdminID = adminID

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
