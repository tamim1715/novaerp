package employee

import (
	"context"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, employee *Employee) error
	FindAll(ctx context.Context, request pagination.PageRequest) ([]Employee, int64, error)
	FindByID(ctx context.Context, id string) (*Employee, error)
	Update(ctx context.Context, employee *Employee) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

func (r *repository) FindAll(
	ctx context.Context,
	request pagination.PageRequest,
) ([]Employee, int64, error) {

	request.Normalize()

	var employees []Employee
	var total int64

	query := r.db.WithContext(ctx).Model(&Employee{})

	if request.Search != "" {
		search := "%" + request.Search + "%"
		query = query.Where(
			"first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ? OR designation ILIKE ?",
			search,
			search,
			search,
			search,
		)
	}

	query.Count(&total)

	err := query.
		Preload("Department").
		Order(request.SortBy + " " + request.Order).
		Limit(request.Size).
		Offset(request.Offset()).
		Find(&employees).Error

	return employees, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*Employee, error) {

	var employee Employee

	err := r.db.WithContext(ctx).Preload("Department").First(&employee, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *repository) Update(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Save(employee).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Employee{}, "id = ?", id).Error
}
