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
}

type ComplaintRepo struct {
	db *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) *ComplaintRepo {
	return &ComplaintRepo{db: db}
}

func (cr ComplaintRepo) CreateComplaint(c entities.Complaint) (entities.Complaint, error) {
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

func (cr ComplaintRepo) AddComplaintPhotos(photos []entities.ComplaintPhoto) ([]entities.ComplaintPhoto, error) {
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

func (cr ComplaintRepo) IsComplaintNumberUnique(complaintNumber string) (bool, error) {
	var count int64
	err := cr.db.Model(&models.Complaint{}).Where("complaint_number = ?", complaintNumber).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
