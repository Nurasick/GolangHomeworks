package handler

import (
	"net/http"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Me retrieves the current user's information
func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Get("userID").(int) //getting user ID from context
	user, err := h.UserService.GetCurrentUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}
	return c.JSON(http.StatusOK, user)
}
