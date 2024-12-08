package news

import (
	"capstone/entities"
	"capstone/utils"
	"errors"
)

func (ns *NewsService) GetAllNewsWithComments() ([]entities.News, error) {
	return ns.newsRepo.GetAllNewsWithComments()
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
		return entities.News{}, errors.New("invalid category ID")
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
