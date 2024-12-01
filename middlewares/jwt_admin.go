package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtInterfaceAdmin interface {
	GenerateJWT(adminID int, role string) (string, error)
}

type JwtAdmin struct {
}

type JWTAdminClaims struct {
	AdminID int    `json:"admin_id"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

var jwtSecretKey = []byte("your_secret_key") // Ganti dengan secret key yang lebih aman

// GenerateJWT generates a JWT token for admin
func (jwtAdmin JwtAdmin) GenerateJWT(adminID int, role string) (string, error) {
	claims := &JWTAdminClaims{
		AdminID: adminID,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token berlaku selama 24 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// JWTMiddleware configures the middleware for validating JWT tokens
func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecretKey,
		Claims:     &JWTAdminClaims{},
	})
}

// ExtractAdminID extracts the admin ID from the JWT token
func ExtractAdminID(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTAdminClaims)
	return claims.AdminID, nil
}

// ExtractAdminRole extracts the admin role from the JWT token
func ExtractAdminRole(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTAdminClaims)
	return claims.Role, nil
}