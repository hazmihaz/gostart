package user

import (
	"context"

	"github.com/hazmihaz/gostart/pkg/log"
	"github.com/jinzhu/gorm"

	"github.com/hazmihaz/gostart/internal/domain"
)

// repository persists user in database
type repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *gorm.DB, logger log.Logger) domain.UserRepository {
	return repository{db, logger}
}

// Get reads the user with the specified ID from the database.
func (r repository) Get(c context.Context, id uint) (domain.User, error) {
	user := domain.User{}
	db := r.db.First(&user, id)
	return user, db.Error
}

// Create saves a new user record in the database.
// It returns the ID of the newly inserted user record.
func (r repository) Create(c context.Context, user domain.User) (domain.User, error) {
	db := r.db.Create(&user)

	return user, db.Error
}

// Update saves the changes to a user in the database.
func (r repository) Update(c context.Context, user domain.User) error {
	res := r.db.Save(&user)

	return res.Error
}

// Delete deletes a user with the specified ID from the database.
func (r repository) Delete(c context.Context, id uint) error {
	res := r.db.Where("id = ?", id).Delete(&domain.User{})
	return res.Error
}

// Count returns the number of the user records in the database.
func (r repository) Count(c context.Context) (int, error) {
	count := 0
	res := r.db.Model(&domain.User{}).Count(&count)

	return count, res.Error
}

// Query retrieves the user records with the specified offset and limit from the database.
func (r repository) Query(c context.Context, offset, limit int) ([]domain.User, error) {
	users := []domain.User{}

	res := r.db.
		Offset(offset).
		Limit(limit).
		Find(&users)

	return users, res.Error
}
