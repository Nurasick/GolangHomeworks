package repository

import (
	"context"
	"errors"
	"university/model"

	"github.com/jackc/pgx/v5"
)

// UserRepository handles user-related database operations
type UserRepository struct {
	Conn *pgx.Conn
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{Conn: conn}
}

// creating a new user in the datavase
func (r *UserRepository) CreateUser(email, passwordHash string) (*model.User, error) {
	query := `INSERT INTO users (email,password_hash)
	VALUES ($1,$2) RETURNING id, created_at;`
	//query to insert the user
	var user model.User
	err := r.Conn.QueryRow(context.Background(), query, email, passwordHash).Scan(&user.ID, &user.CreatedAt) //scanning the returned id and created_at
	if err != nil {
		return nil, errors.New("failed to create the user: %w" + err.Error()) //returning error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `Select id, email, password_hash, created_at FROM users Where email = $1;` //SQL
	var user model.User
	err := r.Conn.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt) //Only getting this three and not getting password hash
	if err != nil {
		return nil, errors.New("Failed to get user by email: %w" + err.Error())
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	query := `Select id, email, created_at FROM users Where id = $1;`
	var user model.User
	err := r.Conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Email, &user.CreatedAt) //getting password hash to compare with given password
	if err != nil {
		return nil, errors.New("Failed to get user by ID: %w" + err.Error())
	}
	return &user, nil
}
