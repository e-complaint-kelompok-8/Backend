package complaints

import (
	"capstone/entities"
	"capstone/repositories/complaints"
	"capstone/utils"
	"errors"
)

type ComplaintServiceInterface interface {
	CreateComplaint(c entities.Complaint, photoURLs []string) (entities.Complaint, []entities.ComplaintPhoto, error)
	GetUserComplaintsByStatusAndCategory(userID int, status string, categoryID, page, limit int) ([]entities.Complaint, int64, error)
	GetComplaintByIDAndUser(id int, userID int) (entities.Complaint, error)
	GetComplaintsByUserID(userID int) ([]entities.Complaint, error)
	GetComplaintsByStatusAndUser(status string, userID int) ([]entities.Complaint, error)
	GetAllComplaintsByUser(userID int) ([]entities.Complaint, error)
	ValidateCategoryID(categoryID int) error
	GetComplaintsByCategoryAndUser(categoryID int, userID int) ([]entities.Complaint, error)
	CancelComplaint(complaintID int, userID int, reason string) (entities.Complaint, error)
	GetComplaintsByStatusAndCategory(status string, categoryID, page, limit int) ([]entities.Complaint, int64, error)
	GetComplaintDetailByID(complaintID int) (entities.Complaint, error)
	// UpdateComplaintStatus(complaintID int, adminID int, newStatus string) error
	GetComplaintByID(complaintID int) (entities.Complaint, error)
	UpdateComplaintByAdmin(complaintID int, updateData entities.Complaint) error
	DeleteComplaintsByAdmin(complaintIDs []int) error
}

type ComplaintService struct {
	complaintRepo complaints.ComplaintRepoInterface
}

func NewComplaintService(cr complaints.ComplaintRepoInterface) *ComplaintService {
	return &ComplaintService{complaintRepo: cr}
}

func (cs *ComplaintService) CreateComplaint(c entities.Complaint, photoURLs []string) (entities.Complaint, []entities.ComplaintPhoto, error) {
	// Validasi data
	if c.Description == "" {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, errors.New(utils.CapitalizeErrorMessage(errors.New("deskripsi diperlukan")))
	}
	if c.CategoryID == 0 {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, errors.New(utils.CapitalizeErrorMessage(errors.New("kategori diperlukan")))
	}
	if c.ComplaintNumber == "" {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, errors.New(utils.CapitalizeErrorMessage(errors.New("nomor pengaduan diperlukan")))
	}
	if c.Title == "" {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, errors.New(utils.CapitalizeErrorMessage(errors.New("judul pengaduan diperlukan")))
	}
	if c.Location == "" {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, errors.New(utils.CapitalizeErrorMessage(errors.New("lokasi pengaduan diperlukan")))
	}

	// Cek keunikan nomor pengaduan
	isUnique, err := cs.complaintRepo.IsComplaintNumberUnique(c.ComplaintNumber)
	if err != nil {
		return entities.Complaint{}, nil, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memvalidasi nomor pengaduan")))
	}
	if !isUnique {
		return entities.Complaint{}, nil, errors.New(utils.CapitalizeErrorMessage(errors.New("nomor pengaduan harus unik")))
	}

	// Simpan complaint
	complaint, err := cs.complaintRepo.CreateComplaint(c)
	if err != nil {
		return entities.Complaint{}, []entities.ComplaintPhoto{}, err
	}

	// Simpan foto jika ada
	var photos []entities.ComplaintPhoto
	for _, url := range photoURLs {
		photos = append(photos, entities.ComplaintPhoto{
			ComplaintID: complaint.ID,
			PhotoURL:    url,
		})
	}

	// Simpan foto ke database
	if len(photos) > 0 {
		photos, err = cs.complaintRepo.AddComplaintPhotos(photos)
		if err != nil {
			return entities.Complaint{}, nil, err
		}
	}

	return complaint, photos, nil
}

func (cs *ComplaintService) GetUserComplaintsByStatusAndCategory(userID int, status string, categoryID, page, limit int) ([]entities.Complaint, int64, error) {
	// Validasi input status jika ada
	if status != "" {
		validStatuses := []string{"proses", "tanggapi", "batal", "selesai"}
		if !utils.StringInSlice(status, validStatuses) {
			return nil, 0, errors.New(utils.CapitalizeErrorMessage(errors.New("status tidak valid")))
		}
	}

	// Ambil data dari repository
	complaints, total, err := cs.complaintRepo.UserGetComplaintsByStatusAndCategory(userID, status, categoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return complaints, total, nil
}

func (cs *ComplaintService) GetComplaintByIDAndUser(id int, userID int) (entities.Complaint, error) {
	// Ambil data keluhan berdasarkan ID dan User ID
	return cs.complaintRepo.GetComplaintByIDAndUser(id, userID)
}

func (cs *ComplaintService) GetComplaintsByUserID(userID int) ([]entities.Complaint, error) {
	// Ambil data keluhan berdasarkan user ID
	complaints, err := cs.complaintRepo.GetComplaintsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func (cs *ComplaintService) GetComplaintsByStatusAndUser(status string, userID int) ([]entities.Complaint, error) {
	// Ambil data keluhan berdasarkan status dan user ID
	complaints, err := cs.complaintRepo.GetComplaintsByStatusAndUser(status, userID)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func (cs *ComplaintService) GetAllComplaintsByUser(userID int) ([]entities.Complaint, error) {
	// Ambil data complaints dari repository
	complaints, err := cs.complaintRepo.GetAllComplaintsByUser(userID)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func (cs *ComplaintService) ValidateCategoryID(categoryID int) error {
	exists, err := cs.complaintRepo.CheckCategoryExists(categoryID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("kategori tidak valid")))
	}
	return nil
}

func (cs ComplaintService) GetComplaintsByCategoryAndUser(categoryID int, userID int) ([]entities.Complaint, error) {
	// Ambil data keluhan berdasarkan kategori dan user ID
	complaints, err := cs.complaintRepo.GetComplaintsByCategoryAndUser(categoryID, userID)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func (cs *ComplaintService) CancelComplaint(complaintID int, userID int, reason string) (entities.Complaint, error) {
	// Ambil complaint untuk validasi
	complaint, err := cs.complaintRepo.GetComplaintByID(complaintID)
	if err != nil {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengaduan tidak ditemukan")))
	}

	// Pastikan complaint milik user
	if complaint.UserID != userID {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("anda tidak memiliki akses untuk membatalkan pengaduan ini")))
	}

	// Pastikan complaint memiliki status "proses"
	if complaint.Status != "proses" {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("hanya pengaduan dengan status 'proses' yang dapat dibatalkan")))
	}

	// Pastikan alasan belum ada
	if complaint.Reason != "" {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengaduan ini sudah memiliki alasan pembatalan")))
	}

	// Perbarui status complaint menjadi "batal" dan simpan alasan
	err = cs.complaintRepo.UpdateComplaintStatus(complaintID, "batal", reason)
	if err != nil {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal membatalkan pengaduan")))
	}

	// Ambil pengaduan yang telah diperbarui
	updatedComplaint, err := cs.complaintRepo.GetComplaintByID(complaintID)
	if err != nil {
		return entities.Complaint{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal mengambil pengaduan yang diperbarui")))
	}

	return updatedComplaint.ToEntitiesReason(), nil
}
