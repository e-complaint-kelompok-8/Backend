package adminai

import (
	adminai "capstone/services/admin_ai"
	"capstone/services/complaints"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminAIController struct {
	aiSuggestionService adminai.AISuggestionServiceInterface
	complaintService    complaints.ComplaintServiceInterface
}

func NewCustomerServiceController(ais adminai.AISuggestionServiceInterface, cs complaints.ComplaintServiceInterface) *AdminAIController {
	return &AdminAIController{aiSuggestionService: ais, complaintService: cs}
}

func (controller *AdminAIController) GetAISuggestion(c echo.Context) error {
	adminID := c.Get("admin_id").(int)
	var request struct {
		ComplaintID int    `json:"complaint_id"`
		Request     string `json:"request"`
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
	err = controller.aiSuggestionService.SaveAISuggestion(adminID, request.ComplaintID, request.Request, aiResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save AI suggestion"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"response": aiResponse,
	})
}
