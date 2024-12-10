package auth

import (
	"capstone/entities"
	"capstone/repositories/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

// NewAdminRepository creates a new instance of AdminRepository
func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// CreateAdmin creates a new admin in the database
func (repo *AdminRepository) CreateAdmin(admin entities.Admin) (entities.Admin, error) {
	modelAdmin := models.FromEntitiesAdmin(admin)
	if err := repo.db.Create(&modelAdmin).Error; err != nil {
		return entities.Admin{}, err
	}
	return modelAdmin.ToEntities(), nil
}

// GetAllAdmin retrieves all admins from the database
func (repo *AdminRepository) GetAllAdmin() ([]entities.Admin, error) {
	var modelAdmins []models.Admin
	if err := repo.db.Find(&modelAdmins).Error; err != nil {
		return nil, err
	}

	var admins []entities.Admin
	for _, modelAdmin := range modelAdmins {
		admins = append(admins, modelAdmin.ToEntities())
	}
	return admins, nil
}

// GetAdminByID retrieves an admin by ID
func (repo *AdminRepository) GetAdminByID(id int) (entities.Admin, error) {
	var modelAdmin models.Admin
	if err := repo.db.First(&modelAdmin, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Admin{}, errors.New("admin not found")
		}
		return entities.Admin{}, err
	}
	return modelAdmin.ToEntities(), nil
}

// UpdateAdmin updates an existing admin
func (repo *AdminRepository) UpdateAdmin(admin entities.Admin) (entities.Admin, error) {
	modelAdmin := models.FromEntitiesAdmin(admin)
	if err := repo.db.Model(&modelAdmin).Updates(map[string]interface{}{
		"email":      modelAdmin.Email,
		"password":   modelAdmin.Password,
		"role":       modelAdmin.Role,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return entities.Admin{}, err
	}
	return modelAdmin.ToEntities(), nil
}

// DeleteAdmin deletes an admin by ID
func (repo *AdminRepository) DeleteAdmin(id int) error {
	if err := repo.db.Delete(&models.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) UpdateAdminProfile(admin entities.Admin) error {
	modelAdmin := models.FromEntitiesAdmin(admin)

	return repo.db.Model(&models.Admin{}).Where("id = ?", admin.ID).Updates(map[string]interface{}{
		"email":      modelAdmin.Email,
		"password":   modelAdmin.Password,
		"photo":      modelAdmin.Photo,
		"updated_at": time.Now(),
	}).Error
}

func (ar *AdminRepository) CheckEmailAdminExists(email string) (bool, error) {
	var count int64
	err := ar.db.Model(&models.Admin{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
