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
	complaints, err := cs.complaintRepo.GetComplaintsByStatusAndCategory(status, categoryID)
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func (cs *ComplaintService) GetComplaintDetailByID(complaintID int) (entities.Complaint, error) {
	// Ambil detail complaint dari repository
	complaint, err := cs.complaintRepo.GetComplaintDetailByID(complaintID)
	if err != nil {
		return entities.Complaint{}, err
	}

	return complaint, nil
}
