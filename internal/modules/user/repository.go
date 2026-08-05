package user

import (
	"context"
	"errors"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

var ErrNilDB = errors.New("database connection is not initialized")

type Repository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, request pagination.PageRequest) ([]User, int64, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) Update(ctx context.Context, user *User) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Delete(&User{}, "id = ?", id).Error
}

func (r *repository) FindAll(ctx context.Context, request pagination.PageRequest) ([]User, int64, error) {
	if r == nil || r.db == nil {
		return nil, 0, ErrNilDB
	}
	request.Normalize()

	var users []User
	var total int64

	query := r.db.WithContext(ctx).Model(&User{})

	if request.Search != "" {
		search := "%" + request.Search + "%"
		query = query.Where(
			"username ILIKE ? OR email ILIKE ?",
			search,
			search,
		)
	}

	query.Count(&total)

	err := query.
		Order(request.SortBy + " " + request.Order).
		Limit(request.Size).
		Offset(request.Offset()).
		Find(&users).Error

	return users, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*User, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var user User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var user User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var user User
	if err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
