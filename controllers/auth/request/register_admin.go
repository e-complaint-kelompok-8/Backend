package request

import "capstone/entities"

type RegisterAdminRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Photo    string `json:"photo"`
}

func (rar RegisterAdminRequest) ToEntities() entities.Admin {
	return entities.Admin{
		ID:       rar.ID,
		Email:    rar.Email,
		Password: rar.Password,
		Role:     rar.Role,
	}
}
