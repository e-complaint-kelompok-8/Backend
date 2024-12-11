package response

import (
	"capstone/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AIResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	User      User   `json:"user"`
	Request   string `json:"request"`
	Response  string `json:"response"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
	Photo string `json:"photo"`
}

type CustomerServiceResponse struct {
	User     User   `json:"user"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type SuccessResponseCS struct {
	Message string
	Data    CustomerServiceResponse
}

func FormatAIResponse(ai entities.AIResponse) AIResponse {
	return AIResponse{
		ID:        ai.ID,
		UserID:    ai.UserID,
		Request:   ai.Request,
		Response:  ai.Response,
		CreatedAt: ai.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func SuccessResponse(c echo.Context, user entities.User, req string, res string) error {
	return c.JSON(http.StatusOK, SuccessResponseCS{
		Message: "AI response retrieved successfully",
		Data: CustomerServiceResponse{
			User: User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Phone: user.Phone,
				Photo: user.PhotoURL,
			},
			Request:  req,
			Response: res,
		},
	})
}
