package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type JwtInterface interface {
	GenerateJWT(userID int, userRole string) (string, error)
}

var jwtKey = []byte("your_secret_key")

// JWTClaims defines the structure of JWT claims
type JWTClaims struct {
	AdminID int    `json:"admin_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for the admin
func (jwtUser JwtUser) GenerateToken(adminID int, email string) (string, error) {
	claims := &JWTClaims{
		AdminID: adminID,
		Email:   email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

type JwtUser struct {
}

type JwtCustomClaims struct {
	UserID   int    `json:"user_id"`
	UserRole string `json:"role"`
	jwt.RegisteredClaims
}

func (jwtUser JwtUser) GenerateJWT(userID int, userRole string) (string, error) {
	claims := &JwtCustomClaims{
		UserID:   userID,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plaintext one
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (jwtUser JwtUser) GetUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		// Gunakan MapClaims dan pastikan token valid
		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok || !userToken.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
		}

		// Ambil user_id dari claims dan simpan di context
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in token"})
		}
		c.Set("user_id", int(userID))

		return next(c)
	}
}
