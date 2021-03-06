package model

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	HashedPass  []byte    `db:"hashed_pass"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func CreateUser(ctx context.Context, name string, hashedPass []byte, description string) (*uuid.UUID, error) {
	var count int

	err := db.GetContext(ctx, &count, "SELECT COUNT(*) FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("user already exists")
	}
	userID := uuid.New()
	date := time.Now()
	_, err = db.Exec("INSERT INTO users (id, name, hashed_pass, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", userID, name, hashedPass, description, date, date)
	if err != nil {
		return nil, err
	}
	return &userID, nil
}

func GetUserByName(ctx context.Context, name string) (*User, error) {
	var user User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(ctx context.Context) ([]*UserResponse, error) {
	var users []*UserResponse
	err := db.SelectContext(ctx, &users, "SELECT id, name, description FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(ctx context.Context, ID uuid.UUID) (*UserResponse, error) {
	var user UserResponse
	err := db.GetContext(ctx, &user, "SELECT id, name, description FROM users WHERE id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
