package news

import "capstone/entities"

func (ns *NewsService) GetAllNewsWithComments() ([]entities.News, error) {
	return ns.newsRepo.GetAllNewsWithComments()
}

func (ns *NewsService) GetNewsByIDWithComments(id string) (entities.News, error) {
	return ns.newsRepo.GetNewsByIDWithComments(id)
}
