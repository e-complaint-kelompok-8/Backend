package complaints

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

func (cr *ComplaintRepo) AdminGetComplaintsByStatusAndCategory(status string, categoryID, page, limit int) ([]entities.Complaint, int64, error) {
	var complaints []models.Complaint
	var total int64

	// Preload User, Category, dan Photos
	query := cr.db.Preload("User").Preload("Category").Preload("Photos").Preload("Feedbacks.Admin").
		Preload("Feedbacks")

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

	// Terapkan urutan descending berdasarkan waktu
	query = query.Order("created_at DESC") // Tambahkan ini

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
		Preload("Feedbacks.Admin"). // Tambahkan preload untuk feedback dan admin
		Preload("Feedbacks").
		First(&complaint, "id = ?", complaintID).Error
	if err != nil {
		return entities.Complaint{}, err
	}
	return complaint.ToEntitiesReason(), nil
}

func (cr *ComplaintRepo) AdminGetComplaintByID(complaintID int) (entities.Complaint, error) {
	var complaint models.Complaint
	err := cr.db.Preload("User").Preload("Category").Preload("Admin").Preload("Photos").Preload("Feedbacks").Preload("Feedbacks.Admin").First(&complaint, "id = ?", complaintID).Error
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

func (cr *ComplaintRepo) ValidateComplaintIDs(complaintIDs []int) ([]int, error) {
	var existingIDs []int

	// Query untuk mendapatkan complaint IDs yang valid di database
	if err := cr.db.Model(&models.Complaint{}).Where("id IN ?", complaintIDs).Pluck("id", &existingIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to validate complaint IDs: %w", err)
	}

	return existingIDs, nil
}

func (cr *ComplaintRepo) DeleteComplaints(complaintIDs []int) error {
	// Hapus data terkait complaints
	if err := cr.db.Where("complaint_id IN ?", complaintIDs).Delete(&models.ComplaintPhoto{}).Error; err != nil {
		return fmt.Errorf("failed to delete complaint photos: %w", err)
	}

	if err := cr.db.Where("complaint_id IN ?", complaintIDs).Delete(&models.Feedback{}).Error; err != nil {
		return fmt.Errorf("failed to delete feedbacks: %w", err)
	}

	if err := cr.db.Where("complaint_id IN ?", complaintIDs).Delete(&models.AISuggestion{}).Error; err != nil {
		return fmt.Errorf("failed to delete AI Suggestion: %w", err)
	}

	// Hapus complaints
	if err := cr.db.Where("id IN ?", complaintIDs).Delete(&models.Complaint{}).Error; err != nil {
		return fmt.Errorf("failed to delete complaints: %w", err)
	}

	return nil
}

func (cr *ComplaintRepo) ImportComplaintsFromCSV(filePath string) error {
	// Buka file CSV
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Baca data dari CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	// Validasi bahwa file tidak kosong dan memiliki header
	if len(records) < 2 {
		return errors.New("CSV file is empty or does not contain any data")
	}

	// Header CSV harus mencocokkan field yang akan dimasukkan
	header := records[0]
	if len(header) != 6 { // Contoh jumlah kolom: ID, CategoryID, Title, Description, Status, CreatedAt
		return errors.New("CSV header does not match the expected format")
	}

	// Parsing data keluhan
	var complaints []models.Complaint
	for _, record := range records[1:] {
		// Parsing setiap kolom
		categoryID, err := strconv.Atoi(record[1])
		if err != nil {
			return fmt.Errorf("invalid category ID: %w", err)
		}

		createdAt, err := time.Parse("2006-01-02", record[5]) // Format tanggal: YYYY-MM-DD
		if err != nil {
			return fmt.Errorf("invalid date format: %w", err)
		}

		// Tambahkan keluhan ke daftar
		complaints = append(complaints, models.Complaint{
			CategoryID:  categoryID,
			Title:       record[2],
			Description: record[3],
			Status:      record[4],
			CreatedAt:   createdAt,
		})
	}

	// Simpan data ke database menggunakan Bulk Insert
	if err := cr.db.Create(&complaints).Error; err != nil {
		return fmt.Errorf("failed to insert complaints into database: %w", err)
	}

	return nil
}