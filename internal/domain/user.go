package domain

import (
	"context"
	"time"
)

// User represents user
type User struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UserRepository encapsulates the logic to access users from the data source.
type UserRepository interface {
	// Get returns the user with the specified user ID.
	Get(ctx context.Context, id uint) (User, error)
	// Count returns the number of users.
	Count(ctx context.Context) (int, error)
	// Query returns the list of users with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]User, error)
	// Create saves a new user in the storage.
	Create(ctx context.Context, user User) (User, error)
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, user User) error
	// Delete removes the user with given ID from the storage.
	Delete(ctx context.Context, id uint) error
}

// UserService encapsulates usecase logic for user.
type UserService interface {
	Get(ctx context.Context, id uint) (User, error)
	Query(ctx context.Context, offset, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input User) (User, error)
	Update(ctx context.Context, id uint, input User) (User, error)
	Delete(ctx context.Context, id uint) (User, error)
}
