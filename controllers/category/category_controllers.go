package category

import (
	"capstone/controllers/category/request"
	"capstone/controllers/category/response"
	"capstone/services/category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CategoryController defines the handlers for category-related operations
type CategoryController struct {
	service category.CategoryService
}

// NewCategoryController creates a new instance of CategoryController
func NewCategoryController(service category.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// CreateCategory handles the creation of a new category
func (c *CategoryController) CreateCategory(ctx echo.Context) error {
	var req request.CategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request data"})
	}

	createdCategory, err := c.service.CreateCategory(ctx.Request().Context(), req.ToEntity())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, response.FromEntity(createdCategory))
}

// GetCategoryByID handles retrieving a category by its ID
func (c *CategoryController) GetCategoryByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
	}

	category, err := c.service.GetCategoryByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.FromEntity(category))
}

// GetAllCategories handles retrieving all categories
func (c *CategoryController) GetAllCategories(ctx echo.Context) error {
	categories, err := c.service.GetAllCategories(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.FromEntities(categories))
}

// UpdateCategory handles updating an existing category
func (c *CategoryController) UpdateCategory(ctx echo.Context) error {
	// Parse and validate the ID from the URL parameter
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
	}

	// Check if the category exists
	existingCategory, err := c.service.GetCategoryByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Category not found"})
	}

	// Bind the request data
	var req request.UpdateCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request data"})
	}

	// Ensure the ID in the request matches the ID parameter
	req.ID = id

	// Convert request to entity using existing category
	updatedEntity := req.ToEntity(existingCategory)

	// Perform the update
	updatedCategory, err := c.service.UpdateCategory(ctx.Request().Context(), updatedEntity)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return the updated category in the response
	return ctx.JSON(http.StatusOK, response.FromEntity(updatedCategory))
}

// DeleteCategory handles deleting a category by its ID
func (c *CategoryController) DeleteCategory(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
	}

	err = c.service.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}