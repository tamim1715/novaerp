package employee

import (
	"context"
	"errors"
	"time"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewService(
	repository Repository,
	logger *zap.Logger,
) Service {

	return &service{
		repository: repository,
		logger:     logger,
	}
}

type Service interface {
	Create(ctx context.Context, request CreateEmployeeRequest) (*Employee, error)
	FindAll(ctx context.Context, request pagination.PageRequest) ([]Employee, int64, error)
	FindByID(ctx context.Context, id string) (*Employee, error)
	Update(ctx context.Context, id string, request UpdateEmployeeRequest) (*Employee, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(ctx context.Context, req CreateEmployeeRequest) (*Employee, error) {

	var joinDate time.Time
	var err error

	if req.JoiningDate != "" {
		joinDate, err = time.Parse("2006-01-02", req.JoiningDate)
		if err != nil {
			return nil, err
		}
	} else {
		joinDate = time.Now()
	}

	status := Status(req.Status)
	if status == "" {
		status = StatusActive
	}

	employee := &Employee{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		DepartmentID: req.DepartmentID,
		Email:        req.Email,
		Phone:        req.Phone,
		Designation:  req.Designation,
		JoiningDate:  joinDate,
		Salary:       req.Salary,
		Status:       status,
	}

	if err := s.repository.Create(ctx, employee); err != nil {
		return nil, err
	}

	return s.repository.FindByID(ctx, employee.ID.String())
}

func (s *service) FindAll(ctx context.Context, request pagination.PageRequest) ([]Employee, int64, error) {
	return s.repository.FindAll(ctx, request)
}

func (s *service) FindByID(ctx context.Context, id string) (*Employee, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id string, req UpdateEmployeeRequest) (*Employee, error) {

	employee, err := s.repository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if req.JoiningDate != "" {
		joinDate, err := time.Parse("2006-01-02", req.JoiningDate)
		if err != nil {
			return nil, err
		}
		employee.JoiningDate = joinDate
	}

	if req.DepartmentID != "" {
		employee.DepartmentID = req.DepartmentID
	}
	if req.FirstName != "" {
		employee.FirstName = req.FirstName
	}
	if req.LastName != "" {
		employee.LastName = req.LastName
	}
	if req.Email != "" {
		employee.Email = req.Email
	}
	if req.Phone != "" {
		employee.Phone = req.Phone
	}
	if req.Designation != "" {
		employee.Designation = req.Designation
	}
	if req.Salary != 0 {
		employee.Salary = req.Salary
	}
	if req.Status != "" {
		employee.Status = Status(req.Status)
	}

	if err := s.repository.Update(ctx, employee); err != nil {
		return nil, err
	}

	return s.repository.FindByID(ctx, employee.ID.String())
}

func (s *service) Delete(ctx context.Context, id string) error {

	_, err := s.repository.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repository.Delete(ctx, id)
}
