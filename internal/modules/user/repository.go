package user

import (
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*User) error
	Update(*User) error
	Delete(string) error
	FindAll(request pagination.PageRequest) ([]User, int64, error)
	FindByID(string) (*User, error)
	FindByEmail(string) (*User, error)
	FindByUsername(string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func (r *repository) FindAll(request pagination.PageRequest) ([]User, int64, error) {
	request.Normalize()

	var users []User
	var total int64

	query := r.db.Model(&User{})

	if request.Search != "" {

		search := "%" + request.Search + "%"

		query = query.Where(
			"name ILIKE ? OR code ILIKE ?",
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

func (r *repository) FindByID(id string) (*User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
