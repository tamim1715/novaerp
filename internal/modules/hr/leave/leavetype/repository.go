package leavetype

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, lt *LeaveType) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]LeaveType, int64, error)
	FindByID(ctx context.Context, id string) (*LeaveType, error)
	Update(ctx context.Context, lt *LeaveType) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, lt *LeaveType) error {
	return r.db.WithContext(ctx).Create(lt).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]LeaveType, int64, error) {
	var list []LeaveType
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&LeaveType{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ?", search, search)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&list).Error

	return list, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*LeaveType, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var lt LeaveType
	if err := r.db.WithContext(ctx).First(&lt, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &lt, nil
}

func (r *repository) Update(ctx context.Context, lt *LeaveType) error {
	return r.db.WithContext(ctx).Save(lt).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&LeaveType{}, "id = ?", uid).Error
}
