package response

import "capstone/entities"

type RegisterAdminResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func RegisterAdminFromEntities(admin entities.Admin) RegisterAdminResponse {
	return RegisterAdminResponse{
		ID:    admin.ID,
		Email: admin.Email,
		Role:  admin.Role,
	}
}