package handler

import (
	"net/http"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	AuthService *service.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Register handles user registration requests
func (h *AuthHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil { //binding the request body to struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	user, err := h.AuthService.Register(req.Email, req.Password) //calling the Register method of AuthService
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

func (h *AuthHandler) Login(c echo.Context) error { // Login handles user login requests
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil { //binding the request body to struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	token, err := h.AuthService.Login(req.Email, req.Password) //calling the Login method of AuthService
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
