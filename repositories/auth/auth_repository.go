package repositories

import (
	"errors"
	"capstone/entities"
)

type AuthRepository interface {
	FindByEmail(email string) (*entities.Admin, error)
}

type authRepository struct {
	// Replace with actual DB connection or data source
	AdminData []entities.Admin
}

func NewAuthRepository() AuthRepository {
	return &authRepository{
		AdminData: []entities.Admin{
			// Example admin with hashed password for "password"
			{ID: 1, Email: "admin@example.com", Password: "$2a$10$ZJeNwuVrWfwR2q.cA2eMjuMdRTMxH4Uw0CgEvFj9lR6lQ3lJjAkdG"},
		},
	}
}

func (r *authRepository) FindByEmail(email string) (*entities.Admin, error) {
	for i := range r.AdminData { // Iterasi by reference
		if r.AdminData[i].Email == email {
			return &r.AdminData[i], nil
		}
	}
	return nil, errors.New("admin not found")
}