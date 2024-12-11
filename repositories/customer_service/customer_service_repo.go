package customerservice

import (
	"capstone/entities"
	"capstone/repositories/models"
	"fmt"

	"gorm.io/gorm"
)

type AIResponseRepositoryInterface interface {
	SaveResponse(response entities.AIResponse) error
	GetUserByID(userID int) (entities.User, error)
	GetUserResponses(userID int, offset int, limit int) ([]entities.AIResponse, int, error)
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

func (repo *AIResponseRepository) GetUserResponses(userID int, offset int, limit int) ([]entities.AIResponse, int, error) {
	var responses []models.AIResponse
	var total int64

	// Hitung total data untuk user tertentu
	if err := repo.db.Model(&models.AIResponse{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count responses: %w", err)
	}

	// Ambil data dengan batasan offset dan limit
	if err := repo.db.Preload("User").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&responses).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve responses: %w", err)
	}

	// Konversi models ke entities
	var result []entities.AIResponse
	for _, response := range responses {
		result = append(result, response.ToEntities())
	}

	return result, int(total), nil
}
