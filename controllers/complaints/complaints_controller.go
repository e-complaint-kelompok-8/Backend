package complaints

import (
	"capstone/controllers/complaints/request"
	"capstone/controllers/complaints/response"
	"capstone/services/complaints"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	complaintService complaints.ComplaintServiceInterface
}

func NewComplaintController(cs complaints.ComplaintServiceInterface) *ComplaintController {
	return &ComplaintController{complaintService: cs}
}

func (cc *ComplaintController) CreateComplaintController(c echo.Context) error {
	req := request.CreateComplaintRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	// Ambil UserID dari context (disimpan oleh middleware `GetUserID`)
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}
	req.UserID = userID

	// Ekstrak foto dari request
	photoURLs := req.PhotoURLs // Tambahkan di request JSON

	err := cc.complaintService.ValidateCategoryID(req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Kategori Tidak Valid",
		})
	}

	complaint, photos, err := cc.complaintService.CreateComplaint(req.ToEntity(), photoURLs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Complaint created successfully",
		"complaint": response.ComplaintFromEntitiesWithPhoto(complaint, photos),
	})
}

func (cc *ComplaintController) GetComplaintById(c echo.Context) error {
	// Ambil ID dari parameter URL
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Complaint ID",
		})
	}

	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil data keluhan dari service
	complaint, err := cc.complaintService.GetComplaintByIDAndUser(complaintID, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Pengaduan Tidak Ditemukan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Complaint status retrieved successfully",
		"complaint": response.ComplaintFromEntitiesWithPhoto(complaint, complaint.Photos),
	})
}

func (cc *ComplaintController) GetComplaintByUser(c echo.Context) error {
	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil keluhan berdasarkan user_id
	complaints, err := cc.complaintService.GetComplaintsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve complaints",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Complaints retrieved successfully",
		"complaints": response.ComplaintsFromEntities(complaints),
	})
}

func (cc *ComplaintController) GetComplaintsByStatus(c echo.Context) error {
	// Ambil status dari parameter URL
	status := c.Param("status")

	// Validasi status
	validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
	isValid := false
	for _, validStatus := range validStatuses {
		if status == validStatus {
			isValid = true
			break
		}
	}

	if !isValid {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid status value",
		})
	}

	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil data keluhan berdasarkan status dan user_id dari service
	complaints, err := cc.complaintService.GetComplaintsByStatusAndUser(status, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve complaints",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Complaints retrieved successfully",
		"complaints": response.ComplaintsFromEntities(complaints),
	})
}

func (cc *ComplaintController) GetAllComplaintsByUser(c echo.Context) error {
	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil semua data complaints milik user
	complaints, err := cc.complaintService.GetAllComplaintsByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve complaints",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Complaints retrieved successfully",
		"complaints": response.ComplaintsFromEntities(complaints),
	})
}

func (cc ComplaintController) GetComplaintsByCategory(c echo.Context) error {
	// Ambil ID kategori dari parameter URL
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Category ID",
		})
	}

	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	err = cc.complaintService.ValidateCategoryID(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Kategori Tidak Valid",
		})
	}

	// Ambil keluhan berdasarkan kategori dan user_id
	complaints, err := cc.complaintService.GetComplaintsByCategoryAndUser(categoryID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve complaints",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Complaints retrieved successfully",
		"complaints": response.ComplaintsFromEntities(complaints),
	})
}

func (cc *ComplaintController) CancelComplaint(c echo.Context) error {
	// Ambil User ID dari middleware
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User tidak memiliki otorisasi",
		})
	}

	// Ambil complaint ID dari parameter
	complaintID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID pengaduan tidak valid",
		})
	}

	// Ambil data alasan pembatalan dari body
	var request struct {
		Reason string `json:"reason" validate:"required"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Data pembatalan tidak valid",
		})
	}

	// Proses pembatalan melalui service
	updatedComplaint, err := cc.complaintService.CancelComplaint(complaintID, userID, request.Reason)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Kembalikan respons dengan data pengaduan yang telah diperbarui
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Pengaduan berhasil dibatalkan",
		"complaint": response.ComplaintFromEntitiesWithReason(updatedComplaint, updatedComplaint.Photos),
	})
}
