package response

import "capstone/entities"

type LoginAdminResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token"`
	Message   string `json:"message"`
}

func LoginAdminFromEntities(admin entities.Admin, token string) LoginAdminResponse {
	return LoginAdminResponse{
		ID:      admin.ID,
		Email:   admin.Email,
		Role:    admin.Role,
		Token:   token,
		Message: "Login successful",
	}
}