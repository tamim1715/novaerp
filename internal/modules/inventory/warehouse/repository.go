package warehouse

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, warehouse *Warehouse) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error)
	FindByID(ctx context.Context, id string) (*Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, warehouse *Warehouse) error {
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error) {
	var warehouses []Warehouse
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Warehouse{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ? OR location ILIKE ?", search, search, search)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := req.Offset()
	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(offset).
		Find(&warehouses).Error

	return warehouses, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*Warehouse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	var warehouse Warehouse
	if err := r.db.WithContext(ctx).First(&warehouse, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *repository) Update(ctx context.Context, warehouse *Warehouse) error {
	return r.db.WithContext(ctx).Save(warehouse).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&Warehouse{}, "id = ?", uid).Error
}
