package news

import (
	"capstone/entities"
	"capstone/utils"
	"errors"
)

func (ns *NewsService) GetAllNewsWithComments(page, limit int) ([]entities.News, int64, error) {
	return ns.newsRepo.GetAllNewsWithComments(page, limit)
}

func (ns *NewsService) GetNewsByIDWithComments(id string) (entities.News, error) {
	return ns.newsRepo.GetNewsByIDWithComments(id)
}

func (ns *NewsService) AddNews(news entities.News) (entities.News, error) {
	// Validasi kategori
	isValid, err := ns.newsRepo.IsCategoryValid(news.CategoryID)
	if err != nil {
		return entities.News{}, err
	}
	if !isValid {
		return entities.News{}, errors.New(utils.CapitalizeErrorMessage(errors.New("ID kategori tidak valid")))
	}

	if news.AdminID == 0 {
		return entities.News{}, errors.New(utils.CapitalizeErrorMessage(errors.New("admin tidak ditemukan")))
	}

	// Simpan berita baru
	newsEntity, err := ns.newsRepo.CreateNews(news)
	if err != nil {
		return entities.News{}, err
	}

	return newsEntity, nil
}

func (ns *NewsService) UpdateNewsByID(id string, updatedNews entities.News) (entities.News, error) {
	// Validasi kategori
	isValid, err := ns.newsRepo.IsCategoryValid(updatedNews.CategoryID)
	if err != nil {
		return entities.News{}, err
	}
	if !isValid {
		return entities.News{}, errors.New(utils.CapitalizeErrorMessage(errors.New("ID kategori tidak benar")))
	}

	// Panggil repository untuk update berita
	news, err := ns.newsRepo.UpdateNewsByID(id, updatedNews)
	if err != nil {
		return entities.News{}, err
	}

	return news, nil
}

func (ns *NewsService) DeleteMultipleNews(ids []int) error {
	// Validasi input
	if len(ids) == 0 {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("tidak ada berita yang dipilih untuk dihapus")))
	}

	// Validasi apakah ID berita ada di database
	existingIDs, err := ns.newsRepo.ValidateNewsIDs(ids)
	if err != nil {
		return err
	}

	// Cek jika ada ID yang tidak ditemukan
	if len(existingIDs) == 0 {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("berita tidak ditemukan")))
	}
	if len(existingIDs) != len(ids) {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("beberapa ID berita tidak ditemukan")))
	}

	// Hapus berita yang valid
	err = ns.newsRepo.DeleteMultipleNews(existingIDs)
	if err != nil {
		return err
	}

	return nil
}
