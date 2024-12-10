package auth

import (
	"capstone/entities"
	"capstone/middlewares"
	repositories "capstone/repositories/auth"
	"capstone/utils"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	adminRepo    *repositories.AdminRepository
	jwtInterface middlewares.JwtAdminInterface
}

// NewAdminService creates a new instance of AdminService
func NewAdminService(adminRepo *repositories.AdminRepository, jwtInterface middlewares.JwtAdminInterface) *AdminService {
	return &AdminService{adminRepo: adminRepo, jwtInterface: jwtInterface}
}

// RegisterAdmin handles the registration of a new admin
func (service *AdminService) RegisterAdmin(admin entities.Admin) (entities.Admin, error) {
	// Periksa apakah email sudah ada
	exists, err := service.adminRepo.CheckEmailAdminExists(admin.Email)
	if err != nil {
		return entities.Admin{}, err
	}
	if exists {
		return entities.Admin{}, errors.New(utils.CapitalizeErrorMessage(errors.New("email sudah ada")))
	}

	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.Admin{}, err
	}
	admin.Password = string(hashedPassword)

	// Save the admin to the database
	createdAdmin, err := service.adminRepo.CreateAdmin(admin)
	if err != nil {
		return entities.Admin{}, err
	}

	return createdAdmin, nil
}

// AuthenticateAdmin validates admin credentials and returns a JWT token if valid
func (service *AdminService) AuthenticateAdmin(email, password string) (string, entities.Admin, error) {
	// Get all admins
	admins, err := service.adminRepo.GetAllAdmin()
	if err != nil {
		return "", entities.Admin{}, err
	}

	// Find the admin by email
	var foundAdmin entities.Admin
	for _, admin := range admins {
		if admin.Email == email {
			foundAdmin = admin
			break
		}
	}

	// If admin not found
	if foundAdmin.ID == 0 {
		return "", entities.Admin{}, errors.New("admin not found")
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(foundAdmin.Password), []byte(password)); err != nil {
		return "", entities.Admin{}, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := service.jwtInterface.GenerateJWT(foundAdmin.ID, foundAdmin.Role)
	if err != nil {
		return "", entities.Admin{}, err
	}

	return token, foundAdmin, nil
}

// GetAllAdmins retrieves all admins
func (service *AdminService) GetAllAdmins() ([]entities.Admin, error) {
	return service.adminRepo.GetAllAdmin()
}

// GetAdminByID retrieves an admin by ID
func (service *AdminService) GetAdminByID(id int) (entities.Admin, error) {
	return service.adminRepo.GetAdminByID(id)
}

// UpdateAdmin handles the update of an admin's information
func (service *AdminService) UpdateAdmin(admin entities.Admin) (entities.Admin, error) {
	// Check if the admin exists
	existingAdmin, err := service.adminRepo.GetAdminByID(admin.ID)
	if err != nil {
		return entities.Admin{}, fmt.Errorf("admin not found")
	}

	// Update only provided fields
	if admin.Email == "" {
		admin.Email = existingAdmin.Email
	}
	if admin.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			return entities.Admin{}, err
		}
		admin.Password = string(hashedPassword)
	} else {
		admin.Password = existingAdmin.Password
	}
	if admin.Role == "" {
		admin.Role = existingAdmin.Role
	}

	return service.adminRepo.UpdateAdmin(admin)
}

// DeleteAdmin handles the deletion of an admin by ID
func (service *AdminService) DeleteAdmin(id int) error {
	return service.adminRepo.DeleteAdmin(id)
}

// ValidateAdminRole checks if an admin has the required role
func (service *AdminService) ValidateAdminRole(adminID int, requiredRole string) error {
	admin, err := service.adminRepo.GetAdminByID(adminID)
	if err != nil {
		return err
	}

	if admin.Role != requiredRole {
		return errors.New("access denied: insufficient role")
	}

	return nil
}

func (service *AdminService) UpdateAdminProfile(adminID int, email, password, photo string) (entities.Admin, error) {
	admin, err := service.adminRepo.GetAdminByID(adminID)
	if err != nil {
		return entities.Admin{}, err
	}

	// Perbarui data jika ada perubahan
	if email != "" {
		admin.Email = email
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return entities.Admin{}, errors.New("failed to hash password")
		}
		admin.Password = string(hashedPassword)
	}

	if photo != "" {
		admin.Photo = photo
	}

	// Simpan perubahan ke database
	if err := service.adminRepo.UpdateAdminProfile(admin); err != nil {
		return entities.Admin{}, err
	}

	return admin, nil
}
