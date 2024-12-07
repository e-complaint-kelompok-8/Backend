package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtAdminInterface interface {
	GenerateJWT(adminID int, role string) (string, error)
}

type JwtAdmin struct {
}

type JwtAdminClaims struct {
	AdminID int    `json:"admin_id"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

func (jwtAdmin JwtAdmin) GenerateJWT(adminID int, role string) (string, error) {
	claims := &JwtAdminClaims{
		AdminID: adminID,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token berlaku selama 24 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secretKey))
}

// Middleware untuk validasi JWT Admin
func (jwtAdmin JwtAdmin) JWTAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari context
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		// Validasi token claims
		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok || !userToken.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
		}

		// Ambil admin_id dan role
		adminID, ok := claims["admin_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Admin ID not found in token"})
		}
		role, ok := claims["role"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Role not found in token"})
		}

		// Simpan admin_id dan role ke context
		c.Set("admin_id", int(adminID))
		c.Set("role", role)

		return next(c)
	}
}

// ExtractAdminID extracts the admin ID from the context
func ExtractAdminID(c echo.Context) (int, error) {
	adminID, ok := c.Get("admin_id").(int)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid admin ID")
	}
	return adminID, nil
}

// ExtractAdminRole extracts the admin role from the context
func ExtractAdminRole(c echo.Context) (string, error) {
	role, ok := c.Get("role").(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid role")
	}
	return role, nil
}
