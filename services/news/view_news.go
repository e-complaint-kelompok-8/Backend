package news

import (
	"capstone/entities"
	"capstone/repositories/news"
)

type NewsServiceInterface interface {
	GetAllNews(page int, limit int) ([]entities.News, int64, error)
	GetNewsByID(id string) (entities.News, error)
	GetAllNewsWithComments(page, limit int) ([]entities.News, int64, error)
	GetNewsByIDWithComments(id string) (entities.News, error)
	AddNews(news entities.News) (entities.News, error)
	UpdateNewsByID(id string, updatedNews entities.News) (entities.News, error)
	DeleteMultipleNews(ids []int) error
}

type NewsService struct {
	newsRepo news.NewsRepositoryInterface
}

func NewNewsService(repo news.NewsRepositoryInterface) *NewsService {
	return &NewsService{newsRepo: repo}
}

func (ns *NewsService) GetAllNews(page int, limit int) ([]entities.News, int64, error) {
	return ns.newsRepo.GetAllNews(page, limit)
}

func (ns *NewsService) GetNewsByID(id string) (entities.News, error) {
	// Panggil repository untuk mendapatkan detail berita berdasarkan ID
	return ns.newsRepo.GetNewsByID(id)
}
