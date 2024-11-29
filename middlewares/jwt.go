package middlewares

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt"
// 	"golang.org/x/crypto/bcrypt"
// )

// var jwtKey = []byte("your_secret_key")

// // JWTClaims defines the structure of JWT claims
// type JWTClaims struct {
// 	AdminID int    `json:"admin_id"`
// 	Email   string `json:"email"`
// 	jwt.StandardClaims
// }

// // GenerateToken creates a new JWT token for the admin
// func GenerateToken(adminID int, email string) (string, error) {
// 	claims := &JWTClaims{
// 		AdminID: adminID,
// 		Email:   email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtKey)
// }

// // HashPassword hashes a password using bcrypt
// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// // CheckPasswordHash compares a hashed password with a plaintext one
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }