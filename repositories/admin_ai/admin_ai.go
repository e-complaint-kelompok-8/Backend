package adminai

import (
	"capstone/repositories/models"
	"fmt"

	"gorm.io/gorm"
)

type AISuggestionRepositoryInterface interface {
	Create(aiSuggestion models.AISuggestion) error
}

func NewCustomerServiceseRepo(db *gorm.DB) *AISuggestionRepository {
	return &AISuggestionRepository{db: db}
}

type AISuggestionRepository struct {
	db *gorm.DB
}

func (repo *AISuggestionRepository) Create(aiSuggestion models.AISuggestion) error {
	if err := repo.db.Create(&aiSuggestion).Error; err != nil {
		return fmt.Errorf("failed to save AI suggestion: %w", err)
	}
	return nil
}
