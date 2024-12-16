package customerservice

import (
	"capstone/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var customerService CustomerService

// Dummy repository untuk test
type DummyCustomerServiceRepo struct{}

func (repo DummyCustomerServiceRepo) SaveResponse(response entities.AIResponse) error {
	if response.Response == "" {
		return errors.New("content cannot be empty")
	}
	return nil
}

func (repo DummyCustomerServiceRepo) GetUserByID(userID int) (entities.User, error) {
	if userID == 1 {
		return entities.User{
			ID:   1,
			Name: "test_user",
		}, nil
	}
	return entities.User{}, errors.New("user not found")
}

func (repo DummyCustomerServiceRepo) GetUserResponses(userID int, offset int, limit int) ([]entities.AIResponse, int, error) {
	if userID != 1 {
		return nil, 0, errors.New("user not found")
	}
	responses := []entities.AIResponse{
		{ID: 1, Response: "Response 1"},
		{ID: 2, Response: "Response 2"},
	}
	return responses, len(responses), nil
}

// Setup untuk unit test
func setupTestService() {
	repo := DummyCustomerServiceRepo{}
	customerService = *NewCustomerService(repo)
}

func TestCustomerService_SaveAIResponse(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		response := entities.AIResponse{
			Request:  "ini adalah test request",
			Response: "ini adalah test response",
			UserID:   1,
		}
		err := customerService.SaveAIResponse(response.UserID, response.Request, response.Response)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
	})

	t.Run("gagal - konten kosong", func(t *testing.T) {
		// Data dummy dengan konten kosong
		response := entities.AIResponse{
			Response: "",
			UserID:   1,
		}

		// Panggil metode SaveResponse melalui service
		err := customerService.SaveAIResponse(response.UserID, response.Request, response.Response)

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "content cannot be empty", err.Error())
	})
}

func TestCustomerService_GetUserByID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		user, err := customerService.GetUserByID(1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "test_user", user.Name)
	})
}

func TestCustomerService_GetUserResponses(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		responses, total, err := customerService.GetUserResponses(1, 0, 2)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 2, total)
		assert.Equal(t, "Response 1", responses[0].Response)
		assert.Equal(t, "Response 2", responses[1].Response)
	})
}
