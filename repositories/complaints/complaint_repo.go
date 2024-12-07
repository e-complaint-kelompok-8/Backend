package complaints

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

type ComplaintRepoInterface interface {
	CreateComplaint(c entities.Complaint) (entities.Complaint, error)
	AddComplaintPhotos(photos []entities.ComplaintPhoto) ([]entities.ComplaintPhoto, error)
	IsComplaintNumberUnique(complaintNumber string) (bool, error)
	GetComplaintsByUserID(userID int) ([]entities.Complaint, error)
	GetComplaintByIDAndUser(id int, userID int) (entities.Complaint, error)
	GetComplaintsByStatusAndUser(status string, userID int) ([]entities.Complaint, error)
	GetAllComplaintsByUser(userID int) ([]entities.Complaint, error)
	CheckCategoryExists(categoryID int) (bool, error)
	GetComplaintsByCategoryAndUser(categoryID int, userID int) ([]entities.Complaint, error)
	GetComplaintByID(complaintID int) (models.Complaint, error)
	UpdateComplaintStatus(complaintID int, status string, reason string) error
	AdminGetComplaintsByStatusAndCategory(status string, categoryID, page, limit int) ([]entities.Complaint, int64, error)
	AdminGetComplaintDetailByID(complaintID int) (entities.Complaint, error)
	AdminUpdateComplaintStatus(complaintID int, newStatus string, adminID int) error
	AdminGetComplaintByID(complaintID int) (entities.Complaint, error)
	AdminUpdateComplaint(complaintID int, updateData entities.Complaint) error
	DeleteComplaint(complaintID int) error
}

type ComplaintRepo struct {
	db *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) *ComplaintRepo {
	return &ComplaintRepo{db: db}
}

func (cr *ComplaintRepo) CreateComplaint(c entities.Complaint) (entities.Complaint, error) {
	complaint := models.FromEntitiesComplaint(c)
	if err := cr.db.Create(&complaint).Error; err != nil {
		return entities.Complaint{}, err
	}

	// Preload User dan Category setelah data disimpan
	err := cr.db.Preload("User").Preload("Category").First(&complaint, "id = ?", complaint.ID).Error
	if err != nil {
		return entities.Complaint{}, err
	}

	return complaint.ToEntities(), nil
}

func (cr *ComplaintRepo) AddComplaintPhotos(photos []entities.ComplaintPhoto) ([]entities.ComplaintPhoto, error) {
	var photoModels []models.ComplaintPhoto
	for _, photo := range photos {
		photoModels = append(photoModels, models.FromEntitiesComplaintPhoto(photo))
	}

	// Simpan ke database
	if err := cr.db.Create(&photoModels).Error; err != nil {
		return nil, err
	}

	// Konversi kembali ke slice entities.ComplaintPhoto
	var savedPhotos []entities.ComplaintPhoto
	for _, photoModel := range photoModels {
		savedPhotos = append(savedPhotos, photoModel.ToEntities())
	}

	return savedPhotos, nil
}

func (cr *ComplaintRepo) IsComplaintNumberUnique(complaintNumber string) (bool, error) {
	var count int64
	err := cr.db.Model(&models.Complaint{}).Where("complaint_number = ?", complaintNumber).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (cr *ComplaintRepo) GetComplaintsByUserID(userID int) ([]entities.Complaint, error) {
	var complaints []models.Complaint

	// Query database untuk mendapatkan keluhan berdasarkan user ID
	err := cr.db.Preload("User").Preload("Category").Preload("Photos").Where("user_id = ?", userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	// Konversi hasil dari models ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, nil
}

func (ar *ComplaintRepo) GetComplaintByIDAndUser(id int, userID int) (entities.Complaint, error) {
	var complaint models.Complaint
	err := ar.db.Preload("User").Preload("Category").
		Where("id = ? AND user_id = ?", id, userID).First(&complaint).Error
	if err != nil {
		return entities.Complaint{}, err
	}
	return complaint.ToEntities(), nil
}

func (cr *ComplaintRepo) GetComplaintsByStatusAndUser(status string, userID int) ([]entities.Complaint, error) {
	var complaints []models.Complaint

	// Query database untuk mendapatkan keluhan berdasarkan status dan user ID
	err := cr.db.Preload("User").Preload("Category").Preload("Photos").
		Where("status = ? AND user_id = ?", status, userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	// Konversi hasil dari models ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, nil
}

func (cr *ComplaintRepo) GetAllComplaintsByUser(userID int) ([]entities.Complaint, error) {
	var complaints []models.Complaint

	// Query database untuk mendapatkan semua complaints milik user
	err := cr.db.Preload("User").Preload("Category").Preload("Photos").
		Where("user_id = ?", userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	// Konversi hasil dari models ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, nil
}

func (cr *ComplaintRepo) CheckCategoryExists(categoryID int) (bool, error) {
	var count int64
	err := cr.db.Model(&models.Category{}).Where("id = ?", categoryID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (cr *ComplaintRepo) GetComplaintsByCategoryAndUser(categoryID int, userID int) ([]entities.Complaint, error) {
	var complaints []models.Complaint

	// Query database untuk mendapatkan keluhan berdasarkan kategori dan user ID
	err := cr.db.Preload("User").Preload("Category").Preload("Photos").
		Where("category_id = ? AND user_id = ?", categoryID, userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	// Konversi hasil dari models ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, nil
}

func (cr *ComplaintRepo) GetComplaintByID(complaintID int) (models.Complaint, error) {
	var complaint models.Complaint
	err := cr.db.Preload("User").
		Preload("Category").
		Preload("Photos").
		First(&complaint, "id = ?", complaintID).Error
	if err != nil {
		return models.Complaint{}, err
	}
	return complaint, nil
}

func (cr *ComplaintRepo) UpdateComplaintStatus(complaintID int, status string, reason string) error {
	// Perbarui status dan simpan alasan pembatalan
	return cr.db.Model(&models.Complaint{}).Where("id = ?", complaintID).Updates(map[string]interface{}{
		"status":     status,
		"reason":     reason,
		"updated_at": gorm.Expr("NOW()"),
	}).Error
}
