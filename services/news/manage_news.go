package news

import "capstone/entities"

func (ns *NewsService) GetAllNewsWithComments() ([]entities.News, error) {
	return ns.newsRepo.GetAllNewsWithComments()
}
