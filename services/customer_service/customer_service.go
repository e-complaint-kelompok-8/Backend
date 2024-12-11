package customerservice

import (
	"capstone/entities"
	customerservice "capstone/repositories/customer_service"
	"time"
)

func NewCustomerService(ai customerservice.AIResponseRepositoryInterface) *CustomerService {
	return &CustomerService{aiResponseRepo: ai}
}

type CustomerServiceInterface interface {
	SaveAIResponse(userID int, request string, response string) error
	GetUserByID(userID int) (entities.User, error)
}

type CustomerService struct {
	aiResponseRepo customerservice.AIResponseRepositoryInterface
}

func (service *CustomerService) SaveAIResponse(userID int, request string, response string) error {
	aiResponse := entities.AIResponse{
		UserID:    userID,
		Request:   request,
		Response:  response,
		CreatedAt: time.Now(),
	}

	// Simpan data ke repository
	return service.aiResponseRepo.SaveResponse(aiResponse)
}

func (service *CustomerService) GetUserByID(userID int) (entities.User, error) {
	return service.aiResponseRepo.GetUserByID(userID)
}
