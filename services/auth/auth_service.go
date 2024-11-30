package auth

import (
	"capstone/entities"
	"capstone/middlewares"
	repositories "capstone/repositories/auth"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type AuthServiceInterface interface {
	Login(email, password string) (*entities.Admin, error)
	RegisterUser(user entities.User) (entities.User, error)
	LoginUser(user entities.User) (entities.User, error)
	VerifyOTP(email, otp string) error
}

type AuthService struct {
	AuthRepository repositories.AuthRepositoryInterface
	jwtInterface   middlewares.JwtInterface
}

func NewAuthService(ar repositories.AuthRepositoryInterface, jwtInterface middlewares.JwtInterface) *AuthService {
	return &AuthService{
		AuthRepository: ar,
		jwtInterface:   jwtInterface}
}

func (s *AuthService) Login(email, password string) (*entities.Admin, error) {
	// Fetch admin data by email
	admin, err := s.AuthRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return admin, nil
}

func (as AuthService) RegisterUser(user entities.User) (entities.User, error) {
	// Periksa apakah email sudah ada
	exists, err := as.AuthRepository.CheckEmailExists(user.Email)
	if err != nil {
		return entities.User{}, err
	}
	if exists {
		return entities.User{}, errors.New("email already exists")
	}

	// Validasi password
	if user.Password == "" {
		return entities.User{}, errors.New("password is empty")
	}

	// Hash password
	hash, _ := HashPassword(user.Password)
	user.Password = hash

	// Generate OTP dan set waktu kedaluwarsa
	user.OTP = GenerateOTP()
	user.OTPExpiry = time.Now().Add(10 * time.Minute) // OTP berlaku 10 menit

	// Debug log untuk memastikan OTP terisi
	fmt.Printf("Generated OTP: %s\n", user.OTP)

	// Buat user baru di database
	fmt.Printf("Before saving to DB, OTP: %s\n", user.OTP)
	user, err = as.AuthRepository.RegisterUser(user)
	if err != nil {
		return entities.User{}, err
	}
	fmt.Printf("After saving to DB, OTP: %s\n", user.OTP)

	// Kirim OTP ke email
	err = sendOTPEmail(user)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to send OTP email: %w", err)
	}

	return user, nil
}

func sendOTPEmail(user entities.User) error {
	body := fmt.Sprintf(`
        <p>Hello %s,</p>
        <p>Your OTP code is: <strong>%s</strong></p>
        <p>This code will expire in 10 minutes.</p>
    `, user.Name, user.OTP)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Your OTP Code")
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),     // Host
		2525,                       // Port
		os.Getenv("SMTP_EMAIL"),    // Email pengirim
		os.Getenv("SMTP_PASSWORD"), // Password API Key
	)
	// fmt.Println("SMTP_HOST:", os.Getenv("SMTP_HOST"))
	// fmt.Println("SMTP_EMAIL:", os.Getenv("SMTP_EMAIL"))
	// fmt.Println("SMTP_PORT:", os.Getenv("EMAIL_PORT"))

	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Printf("Failed to send email to %s: %s\n", user.Email, err.Error())
		return fmt.Errorf("failed to send email: %w", err)
	}
	fmt.Printf("Sending OTP %s to email %s\n", user.OTP, user.Email)

	return nil
}

func (as AuthService) LoginUser(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, errors.New("email is empty")
	} else if user.Password == "" {
		return entities.User{}, errors.New("password is empty")
	}

	oldPassword := user.Password

	// Cari pengguna berdasarkan email
	user, err := as.AuthRepository.LoginUser(user)
	if err != nil {
		return entities.User{}, err
	}

	// Cek apakah email sudah diverifikasi
	if !user.Verified {
		return entities.User{}, errors.New("email is not verified")
	}

	// Cek kecocokan password
	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return entities.User{}, errors.New("password is wrong")
	}

	// Generate token JWT
	token, err := as.jwtInterface.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return entities.User{}, err
	}

	user.Token = token
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // OTP 6 digit
}

func (as AuthService) VerifyOTP(email, otp string) error {
	user, err := as.AuthRepository.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	fmt.Printf("Verifying OTP %s for email %s. Stored OTP: %s\n", otp, email, user.OTP)

	// Periksa apakah OTP cocok
	if user.OTP != otp {
		return errors.New("invalid OTP")
	}

	// Periksa apakah OTP sudah kedaluwarsa
	if time.Now().After(user.OTPExpiry) {
		return errors.New("OTP has expired")
	}

	if user.OTPExpiry.IsZero() {
		user.OTPExpiry = time.Now() // Atur ke nilai default yang valid jika diperlukan
	}

	// Perbarui status verifikasi
	user.Verified = true
	user.OTP = ""                // Hapus OTP setelah verifikasi
	user.OTPExpiry = time.Time{} // Atur waktu kedaluwarsa ke nol (atau gunakan NULL)

	err = as.AuthRepository.UpdateUser(user)
	if err != nil {
		return errors.New("failed to verify email")
	}

	return nil
}
