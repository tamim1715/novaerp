package employee

import (
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Employee) error

	FindAll(
		pagination.PageRequest,
	) ([]Employee, int64, error)

	FindByID(string) (*Employee, error)

	Update(*Employee) error

	Delete(string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(employee *Employee) error {
	return r.db.Create(employee).Error
}

func (r *repository) FindAll(
	request pagination.PageRequest,
) ([]Employee, int64, error) {

	request.Normalize()

	var employees []Employee
	var total int64

	query := r.db.Model(&Employee{})

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
		Find(&employees).Error

	return employees, total, err
}

func (r *repository) FindByID(id string) (*Employee, error) {

	var employee Employee

	err := r.db.First(&employee, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *repository) Update(employee *Employee) error {
	return r.db.Save(employee).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Employee{}, "id = ?", id).Error
}
