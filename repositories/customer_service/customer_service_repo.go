package customerservice

import (
	"capstone/entities"
	"fmt"

	"gorm.io/gorm"
)

type AIResponseRepositoryInterface interface {
	SaveResponse(response entities.AIResponse) error
	GetUserByID(userID int) (entities.User, error)
}

func NewCustomerServiceseRepo(db *gorm.DB) *AIResponseRepository {
	return &AIResponseRepository{db: db}
}

type AIResponseRepository struct {
	db *gorm.DB
}

func (repo *AIResponseRepository) SaveResponse(response entities.AIResponse) error {
	if err := repo.db.Create(&response).Error; err != nil {
		return fmt.Errorf("failed to save AI response: %w", err)
	}
	return nil
}

func (repo *AIResponseRepository) GetUserByID(userID int) (entities.User, error) {
	var user entities.User
	if err := repo.db.First(&user, "id = ?", userID).Error; err != nil {
		return user, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}
