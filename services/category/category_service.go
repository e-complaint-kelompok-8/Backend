package category

import (
	"capstone/entities"
	"capstone/repositories/category"
	"capstone/utils"
	"context"
	"errors"
)

// CategoryService defines the methods for handling category business logic
type CategoryService interface {
	CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error)
	GetCategoryByID(ctx context.Context, id int) (entities.Category, error)
	GetAllCategories(ctx context.Context) ([]entities.Category, error)
	UpdateCategory(ctx context.Context, category entities.Category) (entities.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

// categoryService is the concrete implementation of CategoryService
type categoryService struct {
	repo category.CategoryRepository
}

// NewCategoryService creates a new instance of categoryService
func NewCategoryService(repo category.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// CreateCategory creates a new category after applying business logic
func (s *categoryService) CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error) {
	createdCategory, err := s.repo.CreateCategory(ctx, category)
	if err != nil {
		// Capitalize the error message using utils
		return entities.Category{}, errors.New(utils.CapitalizeErrorMessage(err))
	}
	return createdCategory, nil
}

// GetCategoryByID retrieves a category by its ID
func (s *categoryService) GetCategoryByID(ctx context.Context, id int) (entities.Category, error) {
	category, err := s.repo.GetCategoryByID(ctx, id)
	if err != nil {
		return entities.Category{}, errors.New(utils.CapitalizeErrorMessage(err))
	}
	return category, nil
}

// GetAllCategories retrieves all categories
func (s *categoryService) GetAllCategories(ctx context.Context) ([]entities.Category, error) {
	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, errors.New(utils.CapitalizeErrorMessage(err))
	}
	return categories, nil
}

// UpdateCategory updates a category after applying business logic
func (s *categoryService) UpdateCategory(ctx context.Context, category entities.Category) (entities.Category, error) {
	updatedCategory, err := s.repo.UpdateCategory(ctx, category)
	if err != nil {
		return entities.Category{}, errors.New(utils.CapitalizeErrorMessage(err))
	}
	return updatedCategory, nil
}

// DeleteCategory deletes a category by its ID
func (s *categoryService) DeleteCategory(ctx context.Context, id int) error {
	err := s.repo.DeleteCategory(ctx, id)
	if err != nil {
		return errors.New(utils.CapitalizeErrorMessage(err))
	}
	return nil
}