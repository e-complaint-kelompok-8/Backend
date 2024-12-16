package complaints

import (
	"capstone/controllers/complaints/request"
	"capstone/controllers/complaints/response"
	"capstone/middlewares"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

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

func (cc *ComplaintController) DeleteComplaintsByAdmin(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var complaintIDs []int
	if err := c.Bind(&complaintIDs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	// Validasi input
	if len(complaintIDs) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Complaint IDs are required",
		})
	}

	// Hapus complaints melalui service
	err = cc.complaintService.DeleteComplaintsByAdmin(complaintIDs)
	if err != nil {
		if err.Error() == "some complaint IDs do not exist" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Some complaint IDs do not exist",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete complaints",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Complaints deleted successfully",
	})
}

func (cc *ComplaintController) ImportComplaintsFromCSV(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil file dari form-data
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to retrieve file from request",
		})
	}

	// Cek ekstensi file
	if !strings.HasSuffix(file.Filename, ".csv") {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Only CSV files are allowed",
		})
	}

	// Buka file untuk membaca
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to open file",
		})
	}
	defer src.Close()

	// Pastikan direktori 'uploads/' ada
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create directory for file storage",
		})
	}

	// Simpan file sementara di server
	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to save file",
		})
	}
	defer dst.Close()

	// Salin konten file ke file sementara
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to copy file content",
		})
	}

	// Panggil service untuk memproses file CSV
	err = cc.complaintService.ImportComplaintsFromCSV(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to import complaints from CSV",
		})
	}

	// Kirim respon sukses
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Complaints imported successfully",
	})
}