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

func (ac *AuthController) Login(c echo.Context) error {
	var loginRequest request.LoginRequest

	// Bind request data
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid input"})
	}

	// Validasi menggunakan Validator Echo
	if err := c.Validate(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Validation failed"})
	}

	// Validate email and password
	admin, err := ac.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.LoginResponse{Message: "Login successful", Admin: *admin})
}

func (uc AuthController) RegisterController(c echo.Context) error {
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

func (uc AuthController) VerifyOTPController(c echo.Context) error {
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

func (uc AuthController) LoginController(c echo.Context) error {
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
