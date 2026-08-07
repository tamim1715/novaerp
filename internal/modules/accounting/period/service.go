package period

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

type Service interface {
	CreateFiscalYear(ctx context.Context, req CreateFiscalYearRequest) (*FiscalYear, error)
	FindAllFiscalYears(ctx context.Context) ([]FiscalYear, error)
	FindFiscalYearByID(ctx context.Context, id string) (*FiscalYear, error)
	CloseFiscalYear(ctx context.Context, id string) error
	FindPeriodByID(ctx context.Context, id string) (*AccountingPeriod, error)
	FindPeriodByDate(ctx context.Context, date time.Time) (*AccountingPeriod, error)
	ClosePeriod(ctx context.Context, id string) error
	ReopenPeriod(ctx context.Context, id string) error
}

type service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateFiscalYear(ctx context.Context, req CreateFiscalYearRequest) (*FiscalYear, error) {
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.New("fiscal year end date must be after start date")
	}

	fy := &FiscalYear{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		IsClosed:  false,
	}

	if err := s.repo.CreateFiscalYear(ctx, fy, req.AutoCreatePeriods); err != nil {
		return nil, err
	}

	return s.repo.FindFiscalYearByID(ctx, fy.ID.String())
}

func (s *service) FindAllFiscalYears(ctx context.Context) ([]FiscalYear, error) {
	return s.repo.FindAllFiscalYears(ctx)
}

func (s *service) FindFiscalYearByID(ctx context.Context, id string) (*FiscalYear, error) {
	return s.repo.FindFiscalYearByID(ctx, id)
}

func (s *service) CloseFiscalYear(ctx context.Context, id string) error {
	return s.repo.CloseFiscalYear(ctx, id)
}

func (s *service) FindPeriodByID(ctx context.Context, id string) (*AccountingPeriod, error) {
	return s.repo.FindPeriodByID(ctx, id)
}

func (s *service) FindPeriodByDate(ctx context.Context, date time.Time) (*AccountingPeriod, error) {
	return s.repo.FindPeriodByDate(ctx, date)
}

func (s *service) ClosePeriod(ctx context.Context, id string) error {
	return s.repo.ClosePeriod(ctx, id)
}

func (s *service) ReopenPeriod(ctx context.Context, id string) error {
	return s.repo.ReopenPeriod(ctx, id)
}
