package services

import (
	"errors"
	"capstone/repositories/auth"
	"golang.org/x/crypto/bcrypt"
	"capstone/entities"
)

type AuthService interface {
	Login(email, password string) (*entities.Admin, error)
}

type authService struct {
	AuthRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &authService{AuthRepository: authRepository}
}

func (s *authService) Login(email, password string) (*entities.Admin, error) {
	// Fetch admin data by email
	admin, err := s.AuthRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return admin, nil
}