package manageuser

import (
	"capstone/entities"
	manageuser "capstone/repositories/manage_user"
)

type UserServiceInterface interface {
	GetAllUsers(page, limit int) ([]entities.User, int, error)
	GetUserDetail(userID int) (entities.User, error)
}

type UserService struct {
	userRepo manageuser.UserRepositoryInterface
}

func NewUserService(repo manageuser.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: repo}
}

func (service *UserService) GetAllUsers(page, limit int) ([]entities.User, int, error) {
	// Hitung offset
	offset := (page - 1) * limit

	// Ambil semua user dari repository dengan pagination
	users, total, err := service.userRepo.GetAllUsers(offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (service *UserService) GetUserDetail(userID int) (entities.User, error) {
	// Ambil detail user dari repository
	user, err := service.userRepo.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
