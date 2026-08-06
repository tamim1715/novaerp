package leaverequest

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leavetype"
	"go.uber.org/zap"
)

type Service interface {
	CreateLeaveRequest(ctx context.Context, req CreateLeaveRequest) (*LeaveRequest, error)
	FindAllLeaveRequests(ctx context.Context, req pagination.PageRequest) ([]LeaveRequest, int64, error)
	FindLeaveRequestByID(ctx context.Context, id string) (*LeaveRequest, error)
	UpdateLeaveStatus(ctx context.Context, id string, req UpdateLeaveStatusRequest) (*LeaveRequest, error)
}

type service struct {
	repo          Repository
	leaveTypeRepo leavetype.Repository
	employeeRepo  employee.Repository
	logger        *zap.Logger
}

func NewService(repo Repository, leaveTypeRepo leavetype.Repository, employeeRepo employee.Repository, logger *zap.Logger) Service {
	return &service{
		repo:          repo,
		leaveTypeRepo: leaveTypeRepo,
		employeeRepo:  employeeRepo,
		logger:        logger,
	}
}

func (s *service) CreateLeaveRequest(ctx context.Context, req CreateLeaveRequest) (*LeaveRequest, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	ltID, err := uuid.Parse(req.LeaveTypeID)
	if err != nil {
		return nil, errors.New("invalid leave type ID format")
	}

	// Verify employee exists
	if _, err := s.employeeRepo.FindByID(ctx, req.EmployeeID); err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	// Verify leave type exists
	if _, err := s.leaveTypeRepo.FindByID(ctx, req.LeaveTypeID); err != nil {
		return nil, fmt.Errorf("leave type not found: %w", err)
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date format (use YYYY-MM-DD)")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, errors.New("invalid end date format (use YYYY-MM-DD)")
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	totalDays := int(endDate.Sub(startDate).Hours()/24) + 1

	lr := &LeaveRequest{
		EmployeeID:  empID,
		LeaveTypeID: ltID,
		StartDate:   startDate,
		EndDate:     endDate,
		TotalDays:   totalDays,
		Reason:      req.Reason,
		Status:      "PENDING",
	}

	if err := s.repo.Create(ctx, lr); err != nil {
		return nil, fmt.Errorf("failed to create leave request: %w", err)
	}

	return s.repo.FindByID(ctx, lr.ID.String())
}

func (s *service) FindAllLeaveRequests(ctx context.Context, req pagination.PageRequest) ([]LeaveRequest, int64, error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) FindLeaveRequestByID(ctx context.Context, id string) (*LeaveRequest, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateLeaveStatus(ctx context.Context, id string, req UpdateLeaveStatusRequest) (*LeaveRequest, error) {
	lr, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	lr.Status = req.Status
	lr.Remarks = req.Remarks

	if req.ApprovedBy != "" {
		apprID, err := uuid.Parse(req.ApprovedBy)
		if err == nil {
			lr.ApprovedBy = &apprID
		}
	}

	if err := s.repo.Update(ctx, lr); err != nil {
		return nil, err
	}

	return lr, nil
}
