package shift

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	CreateShift(ctx context.Context, s *Shift) error
	FindAllShifts(ctx context.Context, req pagination.PageRequest) ([]Shift, int64, error)
	FindShiftByID(ctx context.Context, id string) (*Shift, error)
	FindShiftByCode(ctx context.Context, code string) (*Shift, error)
	UpdateShift(ctx context.Context, s *Shift) error
	DeleteShift(ctx context.Context, id string) error

	AssignEmployeeShift(ctx context.Context, es *EmployeeShift) error
	FindEmployeeShiftByID(ctx context.Context, id string) (*EmployeeShift, error)
	FindActiveEmployeeShift(ctx context.Context, employeeID uuid.UUID, date time.Time) (*EmployeeShift, error)
	FindEmployeeShifts(ctx context.Context, req pagination.PageRequest) ([]EmployeeShift, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateShift(ctx context.Context, s *Shift) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *repository) FindAllShifts(ctx context.Context, req pagination.PageRequest) ([]Shift, int64, error) {
	var list []Shift
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Shift{})

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

func (r *repository) FindShiftByID(ctx context.Context, id string) (*Shift, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var s Shift
	err = r.db.WithContext(ctx).First(&s, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *repository) FindShiftByCode(ctx context.Context, code string) (*Shift, error) {
	var s Shift
	err := r.db.WithContext(ctx).First(&s, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *repository) UpdateShift(ctx context.Context, s *Shift) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *repository) DeleteShift(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&Shift{}, "id = ?", uid).Error
}

func (r *repository) AssignEmployeeShift(ctx context.Context, es *EmployeeShift) error {
	return r.db.WithContext(ctx).Create(es).Error
}

func (r *repository) FindEmployeeShiftByID(ctx context.Context, id string) (*EmployeeShift, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var es EmployeeShift
	err = r.db.WithContext(ctx).Preload("Employee").Preload("Shift").First(&es, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &es, nil
}

func (r *repository) FindActiveEmployeeShift(ctx context.Context, employeeID uuid.UUID, date time.Time) (*EmployeeShift, error) {
	var es EmployeeShift
	targetDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	err := r.db.WithContext(ctx).Preload("Shift").
		Where("employee_id = ? AND is_active = true AND effective_from <= ? AND (effective_to IS NULL OR effective_to >= ?)", employeeID, targetDate, targetDate).
		Order("effective_from DESC").
		First(&es).Error

	if err != nil {
		return nil, err
	}
	return &es, nil
}

func (r *repository) FindEmployeeShifts(ctx context.Context, req pagination.PageRequest) ([]EmployeeShift, int64, error) {
	var list []EmployeeShift
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&EmployeeShift{}).Preload("Employee").Preload("Shift")

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&list).Error

	return list, total, err
}
