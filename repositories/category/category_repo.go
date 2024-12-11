package category

import (
	"capstone/entities"
	"capstone/repositories/models"
	"context"

	"gorm.io/gorm"
)

// CategoryRepository defines the methods for accessing the Category data
type CategoryRepository interface {
	CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error)
	GetCategoryByID(ctx context.Context, id int) (entities.Category, error)
	GetAllCategories(ctx context.Context) ([]entities.Category, error)
	UpdateCategory(ctx context.Context, category entities.Category) (entities.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

// categoryRepository is the concrete implementation of CategoryRepository
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of categoryRepository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// CreateCategory inserts a new category into the database
func (r *categoryRepository) CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error) {
	model := models.FromEntitiesCategory(category)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entities.Category{}, err
	}
	return model.ToEntities(), nil
}

// GetCategoryByID retrieves a category by its ID
func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entities.Category, error) {
	var model models.Category
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return entities.Category{}, err
	}
	return model.ToEntities(), nil
}

// GetAllCategories retrieves all categories from the database
func (r *categoryRepository) GetAllCategories(ctx context.Context) ([]entities.Category, error) {
	var modelsData []models.Category
	if err := r.db.WithContext(ctx).Find(&modelsData).Error; err != nil {
		return nil, err
	}

	var entitiesData []entities.Category
	for _, model := range modelsData {
		entitiesData = append(entitiesData, model.ToEntities())
	}
	return entitiesData, nil
}

// UpdateCategory updates an existing category in the database
func (r *categoryRepository) UpdateCategory(ctx context.Context, category entities.Category) (entities.Category, error) {
	model := models.FromEntitiesCategory(category)
	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return entities.Category{}, err
	}
	return model.ToEntities(), nil
}

// DeleteCategory removes a category from the database
func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}