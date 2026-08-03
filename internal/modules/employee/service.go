package employee

import (
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
	Create(CreateEmployeeRequest) (*Employee, error)
	FindAll(request pagination.PageRequest) ([]Employee, int64, error)
	FindByID(string) (*Employee, error)
	Update(string, UpdateEmployeeRequest) (*Employee, error)
	Delete(string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(req CreateEmployeeRequest) (*Employee, error) {

	joinDate, err := time.Parse("2006-01-02", req.JoiningDate)
	if err != nil {
		return nil, err
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
		Status:       req.Status,
	}

	if err := s.repository.Create(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *service) FindAll(request pagination.PageRequest) ([]Employee, int64, error) {
	return s.repository.FindAll(request)
}

func (s *service) FindByID(id string) (*Employee, error) {
	return s.repository.FindByID(id)
}

func (s *service) Update(id string, req UpdateEmployeeRequest) (*Employee, error) {

	employee, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	joinDate, err := time.Parse("2006-01-02", req.JoiningDate)
	if err != nil {
		return nil, err
	}

	employee.DepartmentID = req.DepartmentID
	employee.FirstName = req.FirstName
	employee.LastName = req.LastName
	employee.Email = req.Email
	employee.Phone = req.Phone
	employee.Designation = req.Designation
	employee.JoiningDate = joinDate
	employee.Salary = req.Salary
	employee.Status = req.Status

	if err := s.repository.Update(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *service) Delete(id string) error {

	_, err := s.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repository.Delete(id)
}
