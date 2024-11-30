package request

import "capstone/entities"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (loginReguest LoginRequest) ToEntities() entities.User {
	return entities.User{
		Email:    loginReguest.Email,
		Password: loginReguest.Password,
	}
}
