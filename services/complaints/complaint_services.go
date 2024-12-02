package complaints

import (
	"capstone/entities"
	"capstone/repositories/complaints"
	"capstone/utils"
	"errors"
)

type ComplaintServiceInterface interface {
	CreateComplaint(c entities.Complaint, photoURLs []string) (entities.Complaint, []entities.ComplaintPhoto, error)
}

type ComplaintService struct {
	complaintRepo complaints.ComplaintRepoInterface
}

func NewComplaintService(cr complaints.ComplaintRepoInterface) *ComplaintService {
	return &ComplaintService{complaintRepo: cr}
}

func (cs ComplaintService) CreateComplaint(c entities.Complaint, photoURLs []string) (entities.Complaint, []entities.ComplaintPhoto, error) {
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
