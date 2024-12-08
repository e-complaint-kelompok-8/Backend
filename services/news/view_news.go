package news

import (
	"capstone/entities"
	"capstone/repositories/news"
)

type NewsServiceInterface interface {
	GetAllNews() ([]entities.News, error)
	GetNewsByID(id string) (entities.News, error)
	GetAllNewsWithComments() ([]entities.News, error)
	GetNewsByIDWithComments(id string) (entities.News, error)
}

type NewsService struct {
	newsRepo news.NewsRepositoryInterface
}

func NewNewsService(repo news.NewsRepositoryInterface) *NewsService {
	return &NewsService{newsRepo: repo}
}

func (ns *NewsService) GetAllNews() ([]entities.News, error) {
	return ns.newsRepo.GetAllNews()
}

func (ns *NewsService) GetNewsByID(id string) (entities.News, error) {
	// Panggil repository untuk mendapatkan detail berita berdasarkan ID
	return ns.newsRepo.GetNewsByID(id)
}
