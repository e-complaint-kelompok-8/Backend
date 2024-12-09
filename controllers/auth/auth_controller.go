package auth

import (
	"capstone/controllers/auth/request"
	"capstone/controllers/auth/response"
	"capstone/services/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService auth.AuthServiceInterface
}

func NewAuthController(authService auth.AuthServiceInterface) *AuthController {
	return &AuthController{AuthService: authService}
}

func (uc *AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	c.Bind(&userRegister)
	user, err := uc.AuthService.RegisterUser(userRegister.ToEntities())
	if err != nil {
		// Periksa error untuk memberikan pesan yang lebih spesifik
		if err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "email already exists",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "registration successful",
		"user":    response.RegisterFromEntities(user),
	})
}

func (uc *AuthController) VerifyOTPController(c echo.Context) error {
	type VerifyRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	req := VerifyRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	err := uc.AuthService.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "email verified successfully",
	})
}

func (uc *AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := uc.AuthService.LoginUser(userLogin.ToEntities())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user":    response.LoginFromEntities(user),
	})
}

func (ac *AuthController) GetProfile(c echo.Context) error {
	// Ambil user ID dari JWT
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authorized"})
	}

	// Panggil service untuk mendapatkan profil
	user, err := ac.AuthService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, response.UserProfileFromEntities(user))
}

func (ac *AuthController) UpdateName(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authorized"})
	}

	req := request.UpdateNameRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	user, err := ac.AuthService.UpdateName(userID, req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Name updated successfully",
		"user":    response.UserProfileFromEntities(user),
	})
}

func (ac *AuthController) UpdatePhoto(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authorized"})
	}

	req := request.UpdatePhotoRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Validasi URL photo
	if req.PhotoURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Photo URL cannot be empty"})
	}

	user, err := ac.AuthService.UpdatePhoto(userID, req.PhotoURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Photo updated successfully",
		"user":    response.UserProfileFromEntities(user),
	})
}

func (ac *AuthController) UpdatePassword(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authorized"})
	}

	req := request.UpdatePasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	err := ac.AuthService.UpdatePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
}
