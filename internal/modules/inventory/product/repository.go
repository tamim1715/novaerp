package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, product *Product) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, product *Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error) {
	var products []Product
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Product{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("name ILIKE ? OR sku ILIKE ? OR category ILIKE ?", search, search, search)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := req.Offset()
	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(offset).
		Find(&products).Error

	return products, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*Product, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	var product Product
	if err := r.db.WithContext(ctx).First(&product, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Update(ctx context.Context, product *Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&Product{}, "id = ?", uid).Error
}
