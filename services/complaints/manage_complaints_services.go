package complaints

import (
	"capstone/entities"
	"capstone/utils"
	"errors"
)

func (cs *ComplaintService) GetComplaintsByStatusAndCategory(status string, categoryID int) ([]entities.Complaint, error) {
	// Validasi input status hanya jika tidak kosong
	if status != "" {
		validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
		if !utils.StringInSlice(status, validStatuses) {
			return nil, errors.New("status tidak valid")
		}
	}

	// Ambil data dari repository
	complaints, err := cs.complaintRepo.AdminGetComplaintsByStatusAndCategory(status, categoryID)
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func (cs *ComplaintService) GetComplaintDetailByID(complaintID int) (entities.Complaint, error) {
	// Ambil detail complaint dari repository
	complaint, err := cs.complaintRepo.AdminGetComplaintDetailByID(complaintID)
	if err != nil {
		return entities.Complaint{}, err
	}

	return complaint, nil
}

func (cs *ComplaintService) UpdateComplaintStatus(complaintID int, adminID int, newStatus string) error {
	// Validasi status baru
	validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
	if !utils.StringInSlice(newStatus, validStatuses) {
		return errors.New("status tidak valid")
	}

	// Ambil data complaint dari repository
	complaint, err := cs.complaintRepo.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("pengaduan tidak ditemukan")
	}

	// Pastikan admin berhak mengubah status
	if complaint.AdminID != nil && *complaint.AdminID != adminID {
		return errors.New("anda tidak memiliki akses untuk mengubah status pengaduan ini")
	}

	// Perbarui status pengaduan
	err = cs.complaintRepo.AdminUpdateComplaintStatus(complaintID, newStatus, adminID)
	if err != nil {
		return errors.New("gagal memperbarui status pengaduan")
	}

	return nil
}

func (cs *ComplaintService) GetComplaintByID(complaintID int) (entities.Complaint, error) {
	return cs.complaintRepo.AdminGetComplaintByID(complaintID)
}

func (cs *ComplaintService) UpdateComplaintByAdmin(complaintID int, updateData entities.Complaint) error {
	// Validasi status (jika diperlukan)
	if updateData.Status != "" {
		validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
		if !utils.StringInSlice(updateData.Status, validStatuses) {
			return errors.New("invalid status")
		}
	}

	// Validasi apakah kategori ada (opsional)
	if updateData.CategoryID > 0 {
		exists, err := cs.complaintRepo.CheckCategoryExists(updateData.CategoryID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("invalid category ID")
		}
	}

	// Perbarui data di repository
	err := cs.complaintRepo.AdminUpdateComplaint(complaintID, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ComplaintService) DeleteComplaintByAdmin(complaintID int) error {
	// Periksa apakah complaint ada di database
	complaint, err := cs.complaintRepo.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("complaint not found")
	}

	// Validasi jika diperlukan (opsional)
	if complaint.Status == "selesai" {
		return errors.New("completed complaints cannot be deleted")
	}

	// Hapus complaint di repository
	err = cs.complaintRepo.DeleteComplaint(complaintID)
	if err != nil {
		return errors.New("failed to delete complaint")
	}

	return nil
}
