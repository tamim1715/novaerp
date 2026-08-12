package payroll

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreatePeriod(ctx context.Context, req CreatePayrollPeriodRequest) (*PayrollPeriod, error)
	ProcessPayroll(ctx context.Context, periodID string, req ProcessPayrollRequest) (*PayrollPeriod, error)
	FindAllPeriods(ctx context.Context, req pagination.PageRequest) ([]PayrollPeriod, int64, error)
	FindPeriodByID(ctx context.Context, id string) (*PayrollPeriod, error)
	GetPayslipsByPeriodID(ctx context.Context, periodID string) ([]Payslip, error)
	MarkPaid(ctx context.Context, periodID string) (*PayrollPeriod, error)
}

type service struct {
	repo         Repository
	employeeRepo employee.Repository
	logger       *zap.Logger
}

func NewService(repo Repository, employeeRepo employee.Repository, logger *zap.Logger) Service {
	return &service{
		repo:         repo,
		employeeRepo: employeeRepo,
		logger:       logger,
	}
}

func (s *service) CreatePeriod(ctx context.Context, req CreatePayrollPeriodRequest) (*PayrollPeriod, error) {
	existing, err := s.repo.FindPeriodByMonthYear(ctx, req.Month, req.Year)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("payroll period for %02d/%d already exists", req.Month, req.Year)
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	startDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1)

	period := &PayrollPeriod{
		Month:      req.Month,
		Year:       req.Year,
		StartDate:  startDate,
		EndDate:    endDate,
		Status:     StatusDraft,
		TotalGross: 0,
		TotalNet:   0,
	}

	if err := s.repo.CreatePeriod(ctx, period); err != nil {
		return nil, fmt.Errorf("failed to create payroll period: %w", err)
	}

	return period, nil
}

func (s *service) ProcessPayroll(ctx context.Context, periodID string, req ProcessPayrollRequest) (*PayrollPeriod, error) {
	period, err := s.repo.FindPeriodByID(ctx, periodID)
	if err != nil {
		return nil, err
	}

	if period.Status == StatusPaid {
		return nil, errors.New("cannot process an already paid payroll period")
	}

	// Fetch active employees
	pageReq := pagination.PageRequest{Page: 1, Size: 100}
	employees, _, err := s.employeeRepo.FindAll(ctx, pageReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employees for payroll: %w", err)
	}

	if len(employees) == 0 {
		return nil, errors.New("no active employees found to generate payslips")
	}

	var payslips []Payslip
	var totalGross float64
	var totalNet float64

	for _, emp := range employees {
		basic := emp.Salary
		allowances := req.AllowancesDefault
		deductions := req.DeductionsDefault
		unpaidDeduction := 0.0 // Reserved for unpaid leave integration

		gross := basic + allowances
		net := gross - deductions - unpaidDeduction
		if net < 0 {
			net = 0
		}

		payslips = append(payslips, Payslip{
			PayrollPeriodID:      period.ID,
			EmployeeID:           emp.ID,
			BasicSalary:          basic,
			Allowances:           allowances,
			Deductions:           deductions,
			UnpaidLeaveDeduction: unpaidDeduction,
			GrossSalary:          gross,
			NetSalary:            net,
			Status:               StatusDraft,
		})

		totalGross += gross
		totalNet += net
	}

	// Save payslips
	if err := s.repo.CreatePayslips(ctx, payslips); err != nil {
		return nil, fmt.Errorf("failed to generate payslips: %w", err)
	}

	now := time.Now()
	period.Status = StatusProcessing
	period.TotalGross = totalGross
	period.TotalNet = totalNet
	period.ProcessedAt = &now

	if err := s.repo.UpdatePeriod(ctx, period); err != nil {
		return nil, err
	}

	return s.repo.FindPeriodByID(ctx, periodID)
}

func (s *service) FindAllPeriods(ctx context.Context, req pagination.PageRequest) ([]PayrollPeriod, int64, error) {
	return s.repo.FindAllPeriods(ctx, req)
}

func (s *service) FindPeriodByID(ctx context.Context, id string) (*PayrollPeriod, error) {
	return s.repo.FindPeriodByID(ctx, id)
}

func (s *service) GetPayslipsByPeriodID(ctx context.Context, periodID string) ([]Payslip, error) {
	return s.repo.FindPayslipsByPeriodID(ctx, periodID)
}

func (s *service) MarkPaid(ctx context.Context, periodID string) (*PayrollPeriod, error) {
	period, err := s.repo.FindPeriodByID(ctx, periodID)
	if err != nil {
		return nil, err
	}

	if period.Status == StatusPaid {
		return period, nil
	}

	now := time.Now()
	period.Status = StatusPaid

	if err := s.repo.UpdatePeriod(ctx, period); err != nil {
		return nil, err
	}

	payslips, err := s.repo.FindPayslipsByPeriodID(ctx, periodID)
	if err == nil {
		for i := range payslips {
			payslips[i].Status = StatusPaid
			payslips[i].PaymentDate = &now
			_ = s.repo.UpdatePayslip(ctx, &payslips[i])
		}
	}

	return s.repo.FindPeriodByID(ctx, periodID)
}
