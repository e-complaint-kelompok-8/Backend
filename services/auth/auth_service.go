package auth

import (
	"capstone/entities"
	"capstone/middlewares"
	repositories "capstone/repositories/auth"
	"capstone/utils"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type AuthServiceInterface interface {
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

func (as *AuthService) RegisterUser(user entities.User) (entities.User, error) {
	// Validasi email
	if user.Email == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("email kosong")))
	}

	// Validasi password
	if user.Password == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("password kosong")))
	}

	// Validasi name
	if user.Name == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("nama kosong")))
	}

	// Validasi phone
	if user.Phone == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("nomor telepon kosong")))
	}

	// Periksa apakah email sudah ada
	exists, err := as.AuthRepository.CheckEmailExists(user.Email)
	if err != nil {
		return entities.User{}, err
	}
	if exists {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("email sudah ada")))
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
	// Informasi pengirim
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	// Subjek dan isi email
	subject := "Terima Kasih Telah Mendaftar di Laporin"
	body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>Selamat Datang di Laporin</title>
        </head>
        <body style="font-family: Arial, sans-serif; line-height: 1.6;">
            <h2>Halo, %s</h2>
            <p>Terima kasih telah mendaftar di aplikasi <strong>Laporin</strong>. Untuk menyelesaikan proses pendaftaran, gunakan kode OTP berikut:</p>
            <h1 style="text-align: center; color: #4CAF50;">%s</h1>
            <p>Kode ini hanya berlaku selama <strong>10 menit</strong>. Jika Anda tidak meminta kode ini, silakan abaikan email ini.</p>
            <p>Terima kasih,</p>
            <p><strong>Tim Laporin</strong></p>
        </body>
        </html>
    `, user.Name, user.OTP)

	// Konfigurasi SMTP Gmail
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Membuat pesan email
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Membuat koneksi ke server SMTP Gmail
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// Kirim email
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Printf("Failed to send email to %s: %s\n", user.Email, err.Error())
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("OTP %s sent to email %s\n", user.OTP, user.Email)
	return nil
}

func (as *AuthService) LoginUser(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("email kosong")))
	} else if user.Password == "" {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("password kosong")))
	}

	oldPassword := user.Password

	// Cari pengguna berdasarkan email
	user, err := as.AuthRepository.LoginUser(user)
	if err != nil {
		return entities.User{}, err
	}

	// Cek apakah email sudah diverifikasi
	if !user.Verified {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("email tidak terverifikasi")))
	}

	// Cek kecocokan password
	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return entities.User{}, errors.New(utils.CapitalizeErrorMessage(errors.New("password salah")))
	}

	// Generate token JWT
	token, err := as.jwtInterface.GenerateJWT(user.ID)
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

func (as *AuthService) VerifyOTP(email, otp string) error {
	user, err := as.AuthRepository.GetUserByEmail(email)
	if err != nil {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("pengguna tidak ditemukan")))
	}

	fmt.Printf("Verifying OTP %s for email %s. Stored OTP: %s\n", otp, email, user.OTP)

	// Periksa apakah OTP cocok
	if user.OTP != otp {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("OTP tidak valid")))
	}

	// Periksa apakah OTP sudah kedaluwarsa
	if time.Now().After(user.OTPExpiry) {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("OTP sudah habis masa berlakunya")))
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
		return errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memverifikasi email")))
	}

	return nil
}
