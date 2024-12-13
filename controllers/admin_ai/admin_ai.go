package adminai

import (
	"capstone/controllers/admin_ai/response"
	"capstone/middlewares"
	adminai "capstone/services/admin_ai"
	"capstone/services/auth"
	"capstone/services/complaints"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminAIController struct {
	aiSuggestionService adminai.AISuggestionServiceInterface
	complaintService    complaints.ComplaintServiceInterface
	authAdminService    auth.AdminService
}

func NewCustomerServiceController(ais adminai.AISuggestionServiceInterface, cs complaints.ComplaintServiceInterface, ads auth.AdminService) *AdminAIController {
	return &AdminAIController{aiSuggestionService: ais, complaintService: cs, authAdminService: ads}
}

func (controller *AdminAIController) GetAISuggestion(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	adminID := c.Get("admin_id").(int)
	var request struct {
		ComplaintID int    `json:"complaint_id"`
		Request     string `json:"request"`
	}

	admin, err := controller.authAdminService.GetAdminByID(adminID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Admin not found"})
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request format"})
	}

	// Ambil data pengaduan
	complaint, err := controller.complaintService.GetComplaintByID(request.ComplaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Complaint not found"})
	}

	// Kirim ke AI untuk mendapatkan rekomendasi
	aiResponse, err := controller.aiSuggestionService.GetAISuggestion(request.Request, complaint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get AI response"})
	}

	// Simpan ke database
	aiSuggestion, err := controller.aiSuggestionService.SaveAISuggestion(adminID, request.ComplaintID, request.Request, aiResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save AI suggestion"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.AISuggestionFromEntities(aiSuggestion, admin),
	})
}

func (controller *AdminAIController) FollowUpAISuggestion(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	adminID := c.Get("admin_id").(int)
	aiSuggestionID := c.Param("id")

	var request struct {
		FollowUpQuery string `json:"follow_up_request"`
	}

	admin, err := controller.authAdminService.GetAdminByID(adminID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Admin not found"})
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request format"})
	}

	// Ambil data AISuggestion berdasarkan ID
	aiSuggestion, err := controller.aiSuggestionService.GetAISuggestionByID(aiSuggestionID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "AI Suggestion not found"})
	}

	// Kirim ke AI untuk mendapatkan jawaban lanjutan
	aiResponse, err := controller.aiSuggestionService.GetFollowUpAISuggestion(request.FollowUpQuery, aiSuggestion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get follow-up AI response"})
	}

	// Simpan jawaban lanjutan ke database
	savedAISuggestion, err := controller.aiSuggestionService.SaveAISuggestion(adminID, aiSuggestion.ComplaintID, request.FollowUpQuery, aiResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save follow-up AI suggestion"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.AISuggestionFromEntities(savedAISuggestion, admin),
	})
}
