package manageuser

import (
	"capstone/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userService UserService

type UserRepoDummy struct {
	ShouldFail bool
}

func (repo UserRepoDummy) GetAllUsers(offset, limit int) ([]entities.User, int, error) {
	if repo.ShouldFail {
		return nil, 0, errors.New("failed to retrieve users")
	}

	responses := []entities.User{
		{ID: 1,
			Name:  "test",
			Email: "test@gmail.com",
			Phone: "123",
			Complaints: []entities.Complaint{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "123",
						Phone: "123",
					},
					Category: entities.Category{
						ID:          1,
						Name:        "test",
						Description: "ini cuman test",
					},
					ComplaintNumber: "13213",
					Title:           "ini cuman test",
					Status:          "proses",
					Description:     "ini contoh",
					Reason:          "",
					Location:        "test lokasi",
				},
			}},
		{ID: 2,
			Name:  "test 2",
			Email: "test2@gmail.com",
			Phone: "123",
			Complaints: []entities.Complaint{
				{
					ID: 2,
					User: entities.User{
						ID:       1,
						Name:     "test",
						Email:    "123",
						Password: "321",
						Phone:    "123",
					},
					Category: entities.Category{
						ID:          1,
						Name:        "test",
						Description: "ini cuman test",
					},
					ComplaintNumber: "13213",
					Title:           "ini cuman test",
					Status:          "proses",
					Description:     "ini contoh",
					Reason:          "",
					Location:        "test lokasi",
				},
			}},
	}

	return responses, len(responses), nil
}
func (repo UserRepoDummy) GetUserByID(userID int) (entities.User, error) {
	if userID == 1 {
		return entities.User{ID: 1,
			Name:  "test",
			Email: "test@gmail.com",
			Phone: "123",
			Complaints: []entities.Complaint{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "123",
						Phone: "123",
					},
					Category: entities.Category{
						ID:          1,
						Name:        "test",
						Description: "ini cuman test",
					},
					ComplaintNumber: "13213",
					Title:           "ini cuman test",
					Status:          "proses",
					Description:     "ini contoh",
					Reason:          "",
					Location:        "test lokasi",
				},
			}}, nil
	}

	return entities.User{}, errors.New("user not found")
}

func setupTestService() {
	repo := UserRepoDummy{}
	userService = *NewUserService(repo)
}

func TestCustomerService_GetAllUsers(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		user, total, err := userService.GetAllUsers(0, 2)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 2, total)
		assert.Equal(t, "test", user[0].Name)
		assert.Equal(t, 1, user[0].Complaints[0].ID)
		assert.Equal(t, "test 2", user[1].Name)
		assert.Equal(t, 2, user[1].Complaints[0].ID)
	})

	// Kondisi gagal
	t.Run("gagal", func(t *testing.T) {
		// Setup dengan repository yang mensimulasikan error
		repo := UserRepoDummy{ShouldFail: true}
		service := NewUserService(repo)

		// Data dummy untuk pengujian
		users, total, err := service.GetAllUsers(1, 10)

		// Periksa apakah error terjadi
		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, total)
		assert.Equal(t, "failed to retrieve users", err.Error())
	})
}

func TestCustomerService_GetUserByID(t *testing.T) {
	setupTestService()
	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		user, err := userService.GetUserDetail(1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, "test", user.Name)
		assert.Equal(t, 1, user.Complaints[0].ID)
	})

	// Kondisi gagal
	t.Run("gagal - user tidak ditemukan", func(t *testing.T) {
		// Setup dengan repository tanpa error
		repo := UserRepoDummy{ShouldFail: false}
		service := NewUserService(repo)

		_, err := service.GetUserDetail(999)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}
