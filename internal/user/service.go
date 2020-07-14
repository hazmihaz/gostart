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

// Get returns the album with the specified the user ID.
func (s service) Get(ctx context.Context, id uint) (domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	return user, err
}

// Create creates a new album.
func (s service) Create(ctx context.Context, req domain.User) (domain.User, error) {
	user, err := s.repo.Create(ctx, req)
	s.logger.Info(err)
	return user, err
}

// Update updates the album with the specified ID.
func (s service) Update(ctx context.Context, id uint, req domain.User) (domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	return user, err
}

// Delete deletes the album with the specified ID.
func (s service) Delete(ctx context.Context, id uint) (domain.User, error) {
	user, err := s.Get(ctx, id)
	return user, err
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]domain.User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	return items, err
}
