package adminai

import (
	"capstone/entities"
	"capstone/repositories/models"
	"fmt"

	"gorm.io/gorm"
)

type AISuggestionRepositoryInterface interface {
	Create(aiSuggestion models.AISuggestion) (entities.AISuggestion, error)
	GetByID(id string) (entities.AISuggestion, error)
	GetAllByAdminID(adminID int) ([]entities.AISuggestion, error)
}

func NewCustomerServiceseRepo(db *gorm.DB) *AISuggestionRepository {
	return &AISuggestionRepository{db: db}
}

type AISuggestionRepository struct {
	db *gorm.DB
}

func (repo *AISuggestionRepository) Create(aiSuggestion models.AISuggestion) (entities.AISuggestion, error) {
	// Simpan data ke database
	if err := repo.db.Create(&aiSuggestion).Error; err != nil {
		return entities.AISuggestion{}, fmt.Errorf("failed to save AI suggestion: %w", err)
	}

	// Ambil kembali data lengkap dengan relasi
	if err := repo.db.Preload("Complaint").Preload("Complaint.Category").
		Preload("Complaint.Photos").Preload("Complaint.User").Preload("Admin").
		First(&aiSuggestion, "id = ?", aiSuggestion.ID).Error; err != nil {
		return entities.AISuggestion{}, err
	}

	// Konversi ke entity sebelum dikembalikan
	return aiSuggestion.ToEntities(), nil
}

func (repo *AISuggestionRepository) GetByID(id string) (entities.AISuggestion, error) {
	var aiSuggestion models.AISuggestion
	if err := repo.db.Preload("Complaint").Preload("Admin").Preload("Complaint.Category").Preload("Complaint.Photos").First(&aiSuggestion, id).Error; err != nil {
		return entities.AISuggestion{}, fmt.Errorf("AI suggestion not found: %w", err)
	}
	return aiSuggestion.ToEntities(), nil
}

func (repo *AISuggestionRepository) GetAllByAdminID(adminID int) ([]entities.AISuggestion, error) {
	var aiSuggestions []models.AISuggestion
	if err := repo.db.Preload("Complaint").Preload("Complaint.User").Preload("Complaint.Category").Preload("Complaint.Photos").Preload("Admin").Where("admin_id = ?", adminID).Order("created_at DESC").Find(&aiSuggestions).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve AI suggestions: %w", err)
	}

	// Map models to entities
	result := []entities.AISuggestion{}
	for _, aiSuggestion := range aiSuggestions {
		result = append(result, aiSuggestion.ToEntities())
	}

	return result, nil
}
