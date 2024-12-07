package complaints

import (
	"capstone/entities"
	"capstone/repositories/models"
)

func (cr *ComplaintRepo) GetComplaintsByStatusAndCategory(status string, categoryID int) ([]entities.Complaint, error) {
	var complaints []models.Complaint

	// Preload User, Category, dan Photos
	query := cr.db.Preload("User").Preload("Category").Preload("Photos")

	// Tambahkan kondisi pencarian status jika ada
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Tambahkan kondisi pencarian category_id jika ada
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	// Eksekusi query
	err := query.Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	// Konversi ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, nil
}

func (cr *ComplaintRepo) GetComplaintDetailByID(complaintID int) (entities.Complaint, error) {
	var complaint models.Complaint
	err := cr.db.Preload("User").
		Preload("Category").
		Preload("Photos").
		First(&complaint, "id = ?", complaintID).Error
	if err != nil {
		return entities.Complaint{}, err
	}
	return complaint.ToEntitiesReason(), nil
}
