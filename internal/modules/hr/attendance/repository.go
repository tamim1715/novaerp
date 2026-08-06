package attendance

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, att *Attendance) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Attendance, int64, error)
	FindByID(ctx context.Context, id string) (*Attendance, error)
	FindByEmployeeAndDate(ctx context.Context, employeeID uuid.UUID, date time.Time) (*Attendance, error)
	Update(ctx context.Context, att *Attendance) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, att *Attendance) error {
	return r.db.WithContext(ctx).Create(att).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]Attendance, int64, error) {
	var list []Attendance
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Attendance{}).Preload("Employee")

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("status ILIKE ? OR notes ILIKE ?", search, search)
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

func (r *repository) FindByID(ctx context.Context, id string) (*Attendance, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var att Attendance
	err = r.db.WithContext(ctx).Preload("Employee").First(&att, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &att, nil
}

func (r *repository) FindByEmployeeAndDate(ctx context.Context, employeeID uuid.UUID, date time.Time) (*Attendance, error) {
	var att Attendance
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	err := r.db.WithContext(ctx).Preload("Employee").First(&att, "employee_id = ? AND date = ?", employeeID, dateOnly).Error
	if err != nil {
		return nil, err
	}
	return &att, nil
}

func (r *repository) Update(ctx context.Context, att *Attendance) error {
	return r.db.WithContext(ctx).Save(att).Error
}
