package department

import (
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Department) error
	FindAll(request pagination.PageRequest) ([]Department, int64, error)
	FindByID(string) (*Department, error)
	Update(*Department) error
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

func (r *repository) Create(department *Department) error {
	return r.db.Create(department).Error
}

func (r *repository) FindAll(
	request pagination.PageRequest,
) ([]Department, int64, error) {

	request.Normalize()

	var departments []Department
	var total int64

	query := r.db.Model(&Department{})

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
		Find(&departments).Error

	return departments, total, err
}

func (r *repository) FindByID(id string) (*Department, error) {

	var department Department

	err := r.db.First(&department, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) Update(department *Department) error {
	return r.db.Save(department).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Department{}, "id = ?", id).Error
}
