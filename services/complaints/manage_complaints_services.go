package complaints

import (
	"capstone/entities"
	"capstone/utils"
	"errors"
	"fmt"
)

func (cs *ComplaintService) GetComplaintsByStatusAndCategory(status string, categoryID, page, limit int) ([]entities.Complaint, int64, error) {
	// Validasi input status hanya jika tidak kosong
	if status != "" {
		validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
		if !utils.StringInSlice(status, validStatuses) {
			return nil, 0, errors.New(utils.CapitalizeErrorMessage(errors.New("statusnya tidak valid")))
		}
	}

	// Ambil data dari repository
	complaints, total, err := cs.complaintRepo.AdminGetComplaintsByStatusAndCategory(status, categoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return complaints, total, nil
}

func (cs *ComplaintService) GetComplaintDetailByID(complaintID int) (entities.Complaint, error) {
	// Ambil detail complaint dari repository
	complaint, err := cs.complaintRepo.AdminGetComplaintDetailByID(complaintID)
	if err != nil {
		return entities.Complaint{}, err
	}

	return complaint, nil
}

func (cs *ComplaintService) GetComplaintByID(complaintID int) (entities.Complaint, error) {
	return cs.complaintRepo.AdminGetComplaintByID(complaintID)
}

func (cs *ComplaintService) UpdateComplaintByAdmin(complaintID int, updateData entities.Complaint) error {
	// Validasi status (jika diperlukan)
	if updateData.Status != "" {
		validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
		if !utils.StringInSlice(updateData.Status, validStatuses) {
			return errors.New(utils.CapitalizeErrorMessage(errors.New("statusnya tidak valid")))
		}
	}

	// Validasi apakah kategori ada (opsional)
	if updateData.CategoryID > 0 {
		exists, err := cs.complaintRepo.CheckCategoryExists(updateData.CategoryID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New(utils.CapitalizeErrorMessage(errors.New("ID kategori tidak valid")))
		}
	}

	// Perbarui data di repository
	err := cs.complaintRepo.AdminUpdateComplaint(complaintID, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ComplaintService) DeleteComplaintsByAdmin(complaintIDs []int) error {
	// Validasi complaint IDs
	existingIDs, err := cs.complaintRepo.ValidateComplaintIDs(complaintIDs)
	if err != nil {
		return err
	}

	// Jika ada ID yang tidak ditemukan, return error
	if len(existingIDs) != len(complaintIDs) {
		return fmt.Errorf("some complaint IDs do not exist")
	}

	// Hapus complaints
	err = cs.complaintRepo.DeleteComplaints(existingIDs)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ComplaintService) ImportComplaintsFromCSV(filePath string) error {
	err := cs.complaintRepo.ImportComplaintsFromCSV(filePath)
	if err != nil {
		return fmt.Errorf("failed to import complaints: %w", err)
	}

	return nil
}