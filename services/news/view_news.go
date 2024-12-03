package news

import (
	"capstone/entities"
	"capstone/repositories/news"
)

type NewsServiceInterface interface {
	GetAllNews() ([]entities.News, error)
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
