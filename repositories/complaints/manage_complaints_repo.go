package complaints

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

func (cr *ComplaintRepo) AdminGetComplaintsByStatusAndCategory(status string, categoryID int) ([]entities.Complaint, error) {
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

func (cr *ComplaintRepo) AdminGetComplaintDetailByID(complaintID int) (entities.Complaint, error) {
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

func (cr *ComplaintRepo) AdminUpdateComplaintStatus(complaintID int, newStatus string, adminID int) error {
	// Perbarui status dan admin ID di tabel complaints
	return cr.db.Model(&models.Complaint{}).Where("id = ?", complaintID).Updates(map[string]interface{}{
		"status":     newStatus,
		"admin_id":   adminID,
		"updated_at": gorm.Expr("NOW()"),
	}).Error
}

func (cr *ComplaintRepo) AdminGetComplaintByID(complaintID int) (entities.Complaint, error) {
	var complaint models.Complaint
	err := cr.db.Preload("User").Preload("Category").Preload("Admin").Preload("Photos").First(&complaint, "id = ?", complaintID).Error
	if err != nil {
		return entities.Complaint{}, err
	}
	return complaint.ToEntities(), nil
}

func (cr *ComplaintRepo) AdminUpdateComplaint(complaintID int, updateData entities.Complaint) error {
	// Bangun map untuk kolom yang akan diperbarui
	updateFields := map[string]interface{}{
		"category_id":      updateData.CategoryID,
		"title":            updateData.Title,
		"location":         updateData.Location,
		"status":           updateData.Status,
		"description":      updateData.Description,
		"complaint_number": updateData.ComplaintNumber,
		"admin_id":         updateData.AdminID, // Tambahkan AdminID
		"updated_at":       gorm.Expr("NOW()"),
	}

	// Perbarui data di database
	err := cr.db.Model(&models.Complaint{}).Where("id = ?", complaintID).Updates(updateFields).Error
	if err != nil {
		return err
	}

	return nil
}
