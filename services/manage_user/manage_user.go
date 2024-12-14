package manageuser

import (
	"capstone/entities"
	manageuser "capstone/repositories/manage_user"
)

type UserServiceInterface interface {
	GetAllUsers() ([]entities.User, error)
}

type UserService struct {
	userRepo manageuser.UserRepositoryInterface
}

func NewUserService(repo manageuser.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: repo}
}

func (service *UserService) GetAllUsers() ([]entities.User, error) {
	// Ambil semua user dari repository
	users, err := service.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
