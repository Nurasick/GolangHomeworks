package service

import (
	"university/model"
	"university/pkg/repository"
)

// UserService handles user-related operations
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// GetCurrentUser retrieves the current user's information by their ID.
func (r *UserService) GetCurrentUser(id int) (*model.MeResponse, error) {
	user, err := r.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &model.MeResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (r *UserService) UpdateCurrentUser(user *model.User) error {
	return r.Repo.UpdateUser(user)
}

func (r *UserService) DeactivateCurrentUser(id int) error {
	return r.Repo.DeactivateUser(id)
}
