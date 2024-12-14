package manageuser

import (
	"capstone/entities"
	"capstone/repositories/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAllUsers(offset, limit int) ([]entities.User, int, error)
	GetUserByID(userID int) (entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetAllUsers(offset, limit int) ([]entities.User, int, error) {
	var users []models.User
	var total int64

	// Hitung total data
	err := repo.db.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Ambil data user dengan pagination
	err = repo.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi dari model ke entity dan tambahkan complaints
	var userEntities []entities.User
	for _, user := range users {
		userEntity := user.ToEntities()

		// Ambil complaints untuk user ini
		var complaints []models.Complaint
		err = repo.db.Preload("Category").
			Preload("Photos").
			Preload("Feedbacks").
			Where("user_id = ?", user.ID).Order("created_at DESC").Find(&complaints).Error
		if err != nil {
			return nil, 0, err
		}

		// Konversi complaints ke entities dan tambahkan ke user
		for _, complaint := range complaints {
			userEntity.Complaints = append(userEntity.Complaints, complaint.ToEntities())
		}

		userEntities = append(userEntities, userEntity)
	}

	return userEntities, int(total), nil
}

func (repo *UserRepository) GetUserByID(userID int) (entities.User, error) {
	var user models.User

	// Preload complaints dan data terkait lainnya
	err := repo.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}

	// Konversi dari model ke entity
	userEntity := user.ToEntities()

	// Ambil complaints untuk user ini
	var complaints []models.Complaint
	err = repo.db.Where("user_id = ?", userID).
		Preload("Category").
		Preload("Photos").
		Preload("Feedbacks").
		Order("created_at DESC").
		Find(&complaints).Error
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user complaints: %w", err)
	}

	// Konversi complaints ke entities dan tambahkan ke user
	for _, complaint := range complaints {
		userEntity.Complaints = append(userEntity.Complaints, complaint.ToEntities())
	}

	return userEntity, nil
}
