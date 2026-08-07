package period

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateFiscalYear(ctx context.Context, fy *FiscalYear, autoCreatePeriods bool) error
	FindAllFiscalYears(ctx context.Context) ([]FiscalYear, error)
	FindFiscalYearByID(ctx context.Context, id string) (*FiscalYear, error)
	CloseFiscalYear(ctx context.Context, id string) error
	FindPeriodByID(ctx context.Context, id string) (*AccountingPeriod, error)
	FindPeriodByDate(ctx context.Context, date time.Time) (*AccountingPeriod, error)
	ClosePeriod(ctx context.Context, id string) error
	ReopenPeriod(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateFiscalYear(ctx context.Context, fy *FiscalYear, autoCreatePeriods bool) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(fy).Error; err != nil {
			return err
		}

		if autoCreatePeriods {
			start := fy.StartDate
			for m := 1; m <= 12; m++ {
				// Calculate start and end of this month
				periodStart := time.Date(start.Year(), start.Month()+time.Month(m-1), 1, 0, 0, 0, 0, time.UTC)
				periodEnd := periodStart.AddDate(0, 1, -1)
				if periodEnd.After(fy.EndDate) {
					periodEnd = fy.EndDate
				}

				period := AccountingPeriod{
					FiscalYearID: fy.ID,
					PeriodNumber: m,
					Name:         periodStart.Format("2006-01 (Jan)"),
					StartDate:    periodStart,
					EndDate:      periodEnd,
					Status:       StatusOpen,
				}
				if err := tx.Create(&period).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (r *repository) FindAllFiscalYears(ctx context.Context) ([]FiscalYear, error) {
	var fys []FiscalYear
	err := r.db.WithContext(ctx).Preload("Periods", func(db *gorm.DB) *gorm.DB {
		return db.Order("period_number ASC")
	}).Order("start_date DESC").Find(&fys).Error
	return fys, err
}

func (r *repository) FindFiscalYearByID(ctx context.Context, id string) (*FiscalYear, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var fy FiscalYear
	err = r.db.WithContext(ctx).Preload("Periods", func(db *gorm.DB) *gorm.DB {
		return db.Order("period_number ASC")
	}).First(&fy, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &fy, nil
}

func (r *repository) CloseFiscalYear(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Close all periods
		if err := tx.Model(&AccountingPeriod{}).Where("fiscal_year_id = ?", uid).Update("status", StatusClosed).Error; err != nil {
			return err
		}
		return tx.Model(&FiscalYear{}).Where("id = ?", uid).Update("is_closed", true).Error
	})
}

func (r *repository) FindPeriodByID(ctx context.Context, id string) (*AccountingPeriod, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var p AccountingPeriod
	err = r.db.WithContext(ctx).Preload("FiscalYear").First(&p, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) FindPeriodByDate(ctx context.Context, date time.Time) (*AccountingPeriod, error) {
	var p AccountingPeriod
	err := r.db.WithContext(ctx).
		Where("? BETWEEN start_date AND end_date", date.Format("2006-01-02")).
		First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) ClosePeriod(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Model(&AccountingPeriod{}).Where("id = ?", uid).Update("status", StatusClosed).Error
}

func (r *repository) ReopenPeriod(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	// Check if parent fiscal year is closed
	var p AccountingPeriod
	if err := r.db.WithContext(ctx).Preload("FiscalYear").First(&p, "id = ?", uid).Error; err != nil {
		return err
	}
	if p.FiscalYear != nil && p.FiscalYear.IsClosed {
		return errors.New("cannot reopen period of a closed fiscal year")
	}
	return r.db.WithContext(ctx).Model(&AccountingPeriod{}).Where("id = ?", uid).Update("status", StatusOpen).Error
}
