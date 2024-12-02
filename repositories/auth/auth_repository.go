package auth

import (
	"capstone/entities"
	"capstone/repositories/models"
	"time"

	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepositoryInterface interface {
	RegisterUser(entities.User) (entities.User, error)
	CheckEmailExists(email string) (bool, error)
	LoginUser(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	UpdateUser(user entities.User) error
}

type AuthRepository struct {
	db *gorm.DB
}

func (ar *AuthRepository) RegisterUser(user entities.User) (entities.User, error) {
	userDB := models.FromEntitiesUser(user)
	err := ar.db.Create(&userDB)
	if err.Error != nil {
		return entities.User{}, err.Error
	}
	return userDB.ToEntities(), nil
}

func (ar *AuthRepository) LoginUser(user entities.User) (entities.User, error) {
	userDB := models.FromEntitiesUser(user)
	err := ar.db.First(&userDB, "email = ?", userDB.Email)
	if err.Error != nil {
		return entities.User{}, err.Error
	}
	return userDB.ToEntities(), nil
}

func (ar *AuthRepository) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := ar.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *AuthRepository) GetUserByEmail(email string) (entities.User, error) {
	var user models.User
	err := ar.db.First(&user, "email = ?", email).Error
	if err != nil {
		return entities.User{}, err
	}
	return user.ToEntities(), nil
}

func (ar *AuthRepository) UpdateUser(user entities.User) error {
	userDB := models.FromEntitiesUser(user)

	// Siapkan data pembaruan
	updateData := map[string]interface{}{
		"name":       userDB.Name,
		"email":      userDB.Email,
		"password":   userDB.Password,
		"verified":   userDB.Verified,
		"otp":        userDB.OTP,
		"updated_at": time.Now(),
	}

	// Jika OTPExpiry kosong, atur ke NULL
	if userDB.OTPExpiry.IsZero() {
		updateData["otp_expiry"] = nil
	} else {
		updateData["otp_expiry"] = userDB.OTPExpiry
	}

	// Lakukan pembaruan
	return ar.db.Model(&models.User{}).Where("id = ?", userDB.ID).Updates(updateData).Error
}
