package department

import (
	"context"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, department *Department) error
	FindAll(ctx context.Context, request pagination.PageRequest) ([]Department, int64, error)
	FindByID(ctx context.Context, id string) (*Department, error)
	Update(ctx context.Context, department *Department) error
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

func (r *repository) Create(ctx context.Context, department *Department) error {
	return r.db.WithContext(ctx).Create(department).Error
}

func (r *repository) FindAll(
	ctx context.Context,
	request pagination.PageRequest,
) ([]Department, int64, error) {

	request.Normalize()

	var departments []Department
	var total int64

	query := r.db.WithContext(ctx).Model(&Department{})

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

func (r *repository) FindByID(ctx context.Context, id string) (*Department, error) {

	var department Department

	err := r.db.WithContext(ctx).First(&department, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) Update(ctx context.Context, department *Department) error {
	return r.db.WithContext(ctx).Save(department).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Department{}, "id = ?", id).Error
}
