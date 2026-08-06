package payroll

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePeriod(ctx context.Context, period *PayrollPeriod) error
	FindAllPeriods(ctx context.Context, req pagination.PageRequest) ([]PayrollPeriod, int64, error)
	FindPeriodByID(ctx context.Context, id string) (*PayrollPeriod, error)
	FindPeriodByMonthYear(ctx context.Context, month, year int) (*PayrollPeriod, error)
	UpdatePeriod(ctx context.Context, period *PayrollPeriod) error

	CreatePayslips(ctx context.Context, payslips []Payslip) error
	FindPayslipsByPeriodID(ctx context.Context, periodID string) ([]Payslip, error)
	FindPayslipByID(ctx context.Context, id string) (*Payslip, error)
	UpdatePayslip(ctx context.Context, payslip *Payslip) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreatePeriod(ctx context.Context, period *PayrollPeriod) error {
	return r.db.WithContext(ctx).Create(period).Error
}

func (r *repository) FindAllPeriods(ctx context.Context, req pagination.PageRequest) ([]PayrollPeriod, int64, error) {
	var list []PayrollPeriod
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&PayrollPeriod{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&list).Error

	return list, total, err
}

func (r *repository) FindPeriodByID(ctx context.Context, id string) (*PayrollPeriod, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var period PayrollPeriod
	err = r.db.WithContext(ctx).Preload("Payslips.Employee").First(&period, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (r *repository) FindPeriodByMonthYear(ctx context.Context, month, year int) (*PayrollPeriod, error) {
	var period PayrollPeriod
	err := r.db.WithContext(ctx).First(&period, "month = ? AND year = ?", month, year).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (r *repository) UpdatePeriod(ctx context.Context, period *PayrollPeriod) error {
	return r.db.WithContext(ctx).Save(period).Error
}

func (r *repository) CreatePayslips(ctx context.Context, payslips []Payslip) error {
	if len(payslips) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&payslips).Error
}

func (r *repository) FindPayslipsByPeriodID(ctx context.Context, periodID string) ([]Payslip, error) {
	uid, err := uuid.Parse(periodID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var payslips []Payslip
	err = r.db.WithContext(ctx).Preload("Employee").Where("payroll_period_id = ?", uid).Find(&payslips).Error
	return payslips, err
}

func (r *repository) FindPayslipByID(ctx context.Context, id string) (*Payslip, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var ps Payslip
	err = r.db.WithContext(ctx).Preload("Employee").Preload("PayrollPeriod").First(&ps, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

func (r *repository) UpdatePayslip(ctx context.Context, payslip *Payslip) error {
	return r.db.WithContext(ctx).Save(payslip).Error
}
