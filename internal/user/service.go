package user

import (
	"context"

	"github.com/hazmihaz/gostart/internal/domain"
	"github.com/hazmihaz/gostart/pkg/log"
)

type service struct {
	repo   domain.UserRepository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo domain.UserRepository, logger log.Logger) domain.UserService {
	return &service{repo, logger}
}

// Get returns the user with the specified the user ID.
func (s service) Get(c context.Context, id uint) (domain.User, error) {
	user, err := s.repo.Get(c, id)
	return user, err
}

// Create creates a new user.
func (s service) Create(c context.Context, req domain.User) (domain.User, error) {
	user, err := s.repo.Create(c, req)
	return user, err
}

// Update updates the user with the specified ID.
func (s service) Update(c context.Context, req domain.User) error {
	return s.repo.Update(c, req)
}

// Delete deletes the user with the specified ID.
func (s service) Delete(c context.Context, id uint) error {
	return s.repo.Delete(c, id)
}

// Count returns the number of users.
func (s service) Count(c context.Context) (int, error) {
	return s.repo.Count(c)
}

// Query returns the users with the specified offset and limit.
func (s service) Query(c context.Context, offset, limit int) ([]domain.User, error) {
	items, err := s.repo.Query(c, offset, limit)
	return items, err
}
