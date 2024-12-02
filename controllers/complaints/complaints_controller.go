package complaints

import (
	"capstone/controllers/complaints/request"
	"capstone/controllers/complaints/response"
	"capstone/services/complaints"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	complaintService complaints.ComplaintServiceInterface
}

func NewComplaintController(cs complaints.ComplaintServiceInterface) *ComplaintController {
	return &ComplaintController{complaintService: cs}
}

func (cc ComplaintController) CreateComplaintController(c echo.Context) error {
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
