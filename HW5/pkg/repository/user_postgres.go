package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
	"university/model"

	"github.com/jackc/pgx/v5"
)

/*
This repository layer is used for authentication and authorization,
admin creation flows,
linking teacher/students
It implements All CRUD operations related to User entity

*/

type UserRepositoryInterface interface {
	CreateUser(*model.User) (int, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	UpdateUser(*model.User) error
	DeactivateUser(id int) error
}

type UserRepository struct {
	Conn *pgx.Conn
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{Conn: conn}
}

// creating a new user in the database
func (r *UserRepository) CreateUser(user *model.User) (int, error) {
	query := `INSERT INTO users (email,password_hash,username, role_id, status)
	VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at, updated_at;`
	//query to insert the user

	var id int
	var createdAt time.Time
	var updatedAt time.Time

	err := r.Conn.QueryRow(context.Background(), query, user.Email, user.PasswordHash, user.Username, user.RoleID, user.Status).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		fmt.Errorf("failed to create the user: %w", err)
		return 0, errors.New("failed to create the user: " + err.Error()) //returning error
	}
	user.ID = id
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return id, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `Select id, email,username, role_id, password_hash, created_at, status FROM users Where email = $1;` //SQL
	var user model.User
	err := r.Conn.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.Username, &user.RoleID, &user.PasswordHash, &user.CreatedAt, &user.Status)
	if err != nil {
		fmt.Errorf("Failed to get user by email: %w", err)
		return nil, errors.New("Failed to get user by email: " + err.Error())
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	query := `Select id, email, username, role_id, created_at, updated_at FROM users Where id = $1;`
	var user model.User
	err := r.Conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Email, &user.Username, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Errorf("Failed to get user by ID: %w", err)
		return nil, errors.New("Failed to get user by ID: " + err.Error())
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	query := `UPDATE users SET email=$1, username=$2, role_id=$3, status=$4, updated_at=$5 WHERE id=$6;`
	_, err := r.Conn.Exec(context.Background(), query, user.Email, user.Username, user.RoleID, user.Status, time.Now(), user.ID)
	if err != nil {
		fmt.Errorf("Failed to update user: %w", err)
		return errors.New("Failed to update user: " + err.Error())
	}
	return nil
}

func (r *UserRepository) DeactivateUser(id int) error {
	query := `UPDATE users SET status=$1, updated_at=$2 WHERE id=$3;`
	_, err := r.Conn.Exec(context.Background(), query, model.InactiveStatus, time.Now(), id)
	if err != nil {
		fmt.Errorf("Failed to deactivate user: %w", err)
		return errors.New("Failed to deactivate user: " + err.Error())
	}
	return nil
}
