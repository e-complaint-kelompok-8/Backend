package complaints

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

func (cr *ComplaintRepo) AdminGetComplaintsByStatusAndCategory(status string, categoryID, page, limit int) ([]entities.Complaint, int64, error) {
	var complaints []models.Complaint
	var total int64

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

	// Hitung total data sebelum pagination
	if err := query.Model(&models.Complaint{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Terapkan pagination jika limit > 0
	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	// Eksekusi query
	err := query.Find(&complaints).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi ke entities
	var result []entities.Complaint
	for _, complaint := range complaints {
		result = append(result, complaint.ToEntities())
	}

	return result, total, nil
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

func (cr *ComplaintRepo) DeleteComplaint(complaintID int) error {
	// Hapus data foto terkait di tabel complaint_photos
	if err := cr.db.Where("complaint_id = ?", complaintID).Delete(&models.ComplaintPhoto{}).Error; err != nil {
		return err
	}

	// Hapus data complaint
	if err := cr.db.Where("id = ?", complaintID).Delete(&models.Complaint{}).Error; err != nil {
		return err
	}

	return nil
}
