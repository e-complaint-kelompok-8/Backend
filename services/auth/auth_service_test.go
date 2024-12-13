package auth

import (
	"capstone/entities"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var authService AuthService

type AuthRepoDummy struct {
	GetUserByEmailFunc    func(email string) (entities.User, error)
	CheckEmailExistsDummy func(email string) (bool, error)
	RegisterUserFunc      func(user entities.User) (entities.User, error)
	LoginUserFunc         func(user entities.User) (entities.User, error)
	UpdateUserFunc        func(user entities.User) error
	GetUserByIDFunc       func(userID int) (entities.User, error)
	UpdateUserProfileFunc func(user entities.User) error
}
type JWTRepoDummy struct {
}

func (authRepoDummy AuthRepoDummy) CheckEmailExists(email string) (bool, error) {
	if authRepoDummy.CheckEmailExistsDummy != nil {
		return authRepoDummy.CheckEmailExistsDummy(email) // Panggil mock jika tersedia
	}
	return false, nil // Default: email tidak ditemukan
}

func (authRepoDummy AuthRepoDummy) RegisterUser(user entities.User) (entities.User, error) {
	if authRepoDummy.RegisterUserFunc != nil {
		return authRepoDummy.RegisterUserFunc(user)
	}
	return entities.User{ID: 1,
		Name:     "gilang",
		Email:    "filipi.ketaren@gmail.com",
		Password: "123",
		Phone:    "09321312"}, nil
}

func (authRepoDummy AuthRepoDummy) GetUserByEmail(email string) (entities.User, error) {
	if authRepoDummy.GetUserByEmailFunc != nil {
		return authRepoDummy.GetUserByEmailFunc(email) // Panggil fungsi dinamis jika tersedia
	}
	return entities.User{}, nil
}

func (authRepoDummy AuthRepoDummy) UpdateUser(user entities.User) error {
	if authRepoDummy.UpdateUserFunc != nil {
		return authRepoDummy.UpdateUserFunc(user)
	}
	return nil
}

func (authRepoDummy AuthRepoDummy) GetUserByID(userID int) (entities.User, error) {
	if authRepoDummy.GetUserByIDFunc != nil {
		return authRepoDummy.GetUserByIDFunc(userID)
	}
	return entities.User{}, nil
}

func (authRepoDummy AuthRepoDummy) LoginUser(user entities.User) (entities.User, error) {
	// Cek apakah LoginUserFunc tersedia
	if authRepoDummy.LoginUserFunc != nil {
		return authRepoDummy.LoginUserFunc(user) // Panggil fungsi mock
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost) // Hash password
	return entities.User{
		ID:       1,
		Name:     "filipi",
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Token:    "321",
		Verified: true,
	}, nil
}

func (authRepoDummy AuthRepoDummy) UpdateUserProfile(user entities.User) error {
	return nil
}

func (jwtRepo JWTRepoDummy) GenerateJWT(userID int) (string, error) {
	// Mengembalikan token yang valid
	return "ValidToken", nil
}

// var sendOTPEmailMock = func(user entities.User) error {
// 	return nil // Tidak melakukan apa-apa dalam pengujian
// }

func setup() {
	jwtRepo := JWTRepoDummy{}
	authRepoDummy := AuthRepoDummy{}
	authService = *NewAuthService(authRepoDummy, jwtRepo)
	// sendOTPEmailFunc = sendOTPEmailMock // Mock fungsi
}

func TestAuthService_RegisterUser(t *testing.T) {
	setup()
	t.Run("gagal register karena email kosong", func(t *testing.T) {
		user, err := authService.RegisterUser(entities.User{
			Name:     "test user",
			Password: "password123",
			Phone:    "08123456789",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Email Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal register karena password kosong", func(t *testing.T) {
		user, err := authService.RegisterUser(entities.User{
			Name:  "test user",
			Email: "test@gmail.com",
			Phone: "08123456789",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Password Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal register karena nama kosong", func(t *testing.T) {
		user, err := authService.RegisterUser(entities.User{
			Email:    "test@gmail.com",
			Password: "password123",
			Phone:    "08123456789",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Nama Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal register karena nomor telepon kosong", func(t *testing.T) {
		user, err := authService.RegisterUser(entities.User{
			Name:     "test user",
			Email:    "test@gmail.com",
			Password: "password123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Nomor Telepon Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal register karena email sudah ada", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			CheckEmailExistsDummy: func(email string) (bool, error) {
				return true, nil // Simulasikan email sudah ada
			},
		}

		user, err := authService.RegisterUser(entities.User{
			Name:     "test user",
			Email:    "test@gmail.com",
			Password: "password123",
			Phone:    "08123456789",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Email Sudah Ada", err.Error())
		assert.Equal(t, 0, user.ID)
	})
}

func TestAuthService_LoginUser(t *testing.T) {
	setup()

	t.Run("sukses login", func(t *testing.T) {
		user, err := authService.LoginUser(entities.User{
			Email:    "test@example.com",
			Password: "123",
		})
		assert.Nil(t, err)
		assert.Equal(t, "filipi", user.Name)
		assert.Equal(t, "ValidToken", user.Token) // Pastikan token sesuai
	})

	t.Run("gagal login karena email kosong", func(t *testing.T) {
		user, err := authService.LoginUser(entities.User{
			Password: "password123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Email Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena password kosong", func(t *testing.T) {
		user, err := authService.LoginUser(entities.User{
			Email: "test@example.com",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Password Kosong", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena email tidak ditemukan", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			LoginUserFunc: func(user entities.User) (entities.User, error) {
				return entities.User{}, errors.New("email tidak ditemukan")
			},
		}

		user, err := authService.LoginUser(entities.User{
			Email:    "notfound@example.com",
			Password: "password123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "email tidak ditemukan", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena email belum diverifikasi", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			LoginUserFunc: func(user entities.User) (entities.User, error) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				return entities.User{
					ID:       1,
					Name:     "test user",
					Email:    "test@example.com",
					Password: string(hashedPassword),
					Verified: false,
				}, nil
			},
		}

		user, err := authService.LoginUser(entities.User{
			Email:    "test@example.com",
			Password: "password123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Email Tidak Terverifikasi", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena password salah", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			LoginUserFunc: func(user entities.User) (entities.User, error) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				return entities.User{
					ID:       1,
					Name:     "test user",
					Email:    "test@example.com",
					Password: string(hashedPassword),
					Verified: true,
				}, nil
			},
		}

		user, err := authService.LoginUser(entities.User{
			Email:    "test@example.com",
			Password: "wrongpassword",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "Password Salah", err.Error())
		assert.Equal(t, 0, user.ID)
	})
}

func TestAuthService_VerifyOTP(t *testing.T) {
	setup()

	t.Run("sukses verify OTP", func(t *testing.T) {
		// Mock repository untuk GetUserByEmail
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: func(email string) (entities.User, error) {
				return entities.User{
					Email:     "test@example.com",
					OTP:       "123456",
					OTPExpiry: time.Now().Add(10 * time.Minute),
				}, nil
			},
			UpdateUserFunc: func(user entities.User) error {
				return nil // Simulasi sukses update user
			},
		}

		err := authService.VerifyOTP("test@example.com", "123456")
		assert.Nil(t, err)
	})

	t.Run("gagal verify OTP karena pengguna tidak ditemukan", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: func(email string) (entities.User, error) {
				return entities.User{}, errors.New("pengguna tidak ditemukan")
			},
		}

		err := authService.VerifyOTP("notfound@example.com", "123456")
		assert.NotNil(t, err)
		assert.Equal(t, "Pengguna Tidak Ditemukan", err.Error())
	})

	t.Run("gagal verify OTP karena OTP salah", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: func(email string) (entities.User, error) {
				return entities.User{
					Email:     "test@example.com",
					OTP:       "654321",
					OTPExpiry: time.Now().Add(10 * time.Minute),
				}, nil
			},
		}

		err := authService.VerifyOTP("test@example.com", "123456")
		assert.NotNil(t, err)
		assert.Equal(t, "OTP Tidak Valid", err.Error())
	})

	t.Run("gagal verify OTP karena OTP sudah kedaluwarsa", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: func(email string) (entities.User, error) {
				return entities.User{
					Email:     "test@example.com",
					OTP:       "123456",
					OTPExpiry: time.Now().Add(-10 * time.Minute),
				}, nil
			},
		}

		err := authService.VerifyOTP("test@example.com", "123456")
		assert.NotNil(t, err)
		assert.Equal(t, "OTP Sudah Habis Masa Berlakunya", err.Error())
	})
}

func TestPasswordHashing(t *testing.T) {
	password := "123"
	hashedPassword, err := HashPassword(password)
	assert.Nil(t, err)
	assert.NotEmpty(t, hashedPassword)

	match := CheckPasswordHash(password, hashedPassword)
	assert.True(t, match)
}

func TestGenerateOTP(t *testing.T) {
	t.Run("generate OTP menghasilkan 6 digit angka", func(t *testing.T) {
		otp := GenerateOTP()
		assert.NotEmpty(t, otp)
		assert.Len(t, otp, 6, "OTP harus memiliki panjang 6 karakter")
	})
}

func TestAuthService_GetUserByID(t *testing.T) {
	setup()

	t.Run("sukses mendapatkan user dengan ID yang valid", func(t *testing.T) {
		// Mock behavior untuk GetUserByID
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: nil,
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{
					ID:    1,
					Name:  "Filipi",
					Email: "test@example.com",
				}, nil
			},
		}

		user, err := authService.GetUserByID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "Filipi", user.Name)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("gagal mendapatkan user dengan ID yang tidak valid", func(t *testing.T) {
		// Mock behavior untuk GetUserByID dengan error
		authService.AuthRepository = AuthRepoDummy{
			GetUserByEmailFunc: nil,
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{}, errors.New("user not found")
			},
		}

		user, err := authService.GetUserByID(999)
		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Equal(t, 0, user.ID)
	})
}

func TestAuthService_UpdateName(t *testing.T) {
	setup()

	t.Run("sukses update nama", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{ID: userID, Name: "Old Name"}, nil
			},
			UpdateUserFunc: func(user entities.User) error {
				return nil
			},
		}

		user, err := authService.UpdateName(1, "New Name")
		assert.Nil(t, err)
		assert.Equal(t, "New Name", user.Name)
	})

	t.Run("gagal update nama karena user tidak ditemukan", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{}, errors.New("user not found")
			},
		}

		user, err := authService.UpdateName(1, "New Name")
		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Equal(t, "", user.Name)
	})
}

func TestAuthService_UpdatePhoto(t *testing.T) {
	setup()

	t.Run("sukses update foto", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{ID: userID, PhotoURL: "old_photo.jpg"}, nil
			},
			UpdateUserProfileFunc: func(user entities.User) error {
				return nil
			},
		}

		user, err := authService.UpdatePhoto(1, "new_photo.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "new_photo.jpg", user.PhotoURL)
	})

	t.Run("gagal update foto karena user tidak ditemukan", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{}, errors.New("user not found")
			},
		}

		user, err := authService.UpdatePhoto(1, "new_photo.jpg")
		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Equal(t, "", user.PhotoURL)
	})

	t.Run("gagal update foto karena URL kosong", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{ID: userID, PhotoURL: "old_photo.jpg"}, nil
			},
		}

		_, err := authService.UpdatePhoto(1, "")
		assert.NotNil(t, err)
		assert.Equal(t, "photo URL cannot be empty", err.Error())
	})
}

func TestAuthService_UpdatePassword(t *testing.T) {
	setup()

	t.Run("sukses update password", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				hashedPassword, _ := HashPassword("old_password")
				return entities.User{ID: userID, Password: hashedPassword}, nil
			},
			UpdateUserFunc: func(user entities.User) error {
				return nil
			},
		}

		err := authService.UpdatePassword(1, "old_password", "new_password")
		assert.Nil(t, err)
	})

	t.Run("gagal update password karena user tidak ditemukan", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				return entities.User{}, errors.New("user not found")
			},
		}

		err := authService.UpdatePassword(1, "old_password", "new_password")
		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("gagal update password karena password lama salah", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				hashedPassword, _ := HashPassword("correct_password")
				return entities.User{ID: userID, Password: hashedPassword}, nil
			},
		}

		err := authService.UpdatePassword(1, "wrong_password", "new_password")
		assert.NotNil(t, err)
		assert.Equal(t, "old password is incorrect", err.Error())
	})

	t.Run("gagal update password karena repository error", func(t *testing.T) {
		authService.AuthRepository = AuthRepoDummy{
			GetUserByIDFunc: func(userID int) (entities.User, error) {
				hashedPassword, _ := HashPassword("old_password")
				return entities.User{ID: userID, Password: hashedPassword}, nil
			},
			UpdateUserFunc: func(user entities.User) error {
				return errors.New("failed to update password")
			},
		}

		err := authService.UpdatePassword(1, "old_password", "new_password")
		assert.NotNil(t, err)
		assert.Equal(t, "failed to update password", err.Error())
	})
}
