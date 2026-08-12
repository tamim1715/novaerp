package attendance

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CheckIn(ctx context.Context, req CheckInRequest) (*Attendance, error)
	CheckOut(ctx context.Context, req CheckOutRequest) (*Attendance, error)
	CreateAttendance(ctx context.Context, req CreateAttendanceRequest) (*Attendance, error)
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Attendance, int64, error)
	FindByID(ctx context.Context, id string) (*Attendance, error)
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

func (s *service) CheckIn(ctx context.Context, req CheckInRequest) (*Attendance, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	if _, err := s.employeeRepo.FindByID(ctx, req.EmployeeID); err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	localNow := time.Now()
	nowUTC := localNow.UTC()
	today := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), 0, 0, 0, 0, time.UTC)

	att, err := s.repo.FindByEmployeeAndDate(ctx, empID, today)
	if err == nil {
		if att.CheckIn != nil {
			return nil, errors.New("employee already checked in today")
		}
		att.CheckIn = &nowUTC
		if req.Notes != "" {
			att.Notes = req.Notes
		}
		if err := s.repo.Update(ctx, att); err != nil {
			return nil, err
		}
		return att, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	status := StatusPresent
	// If check in after 09:30 AM local time, mark as LATE
	if localNow.Hour() > 9 || (localNow.Hour() == 9 && localNow.Minute() > 30) {
		status = StatusLate
	}

	att = &Attendance{
		EmployeeID: empID,
		Date:       today,
		CheckIn:    &nowUTC,
		Status:     status,
		Notes:      req.Notes,
	}

	if err := s.repo.Create(ctx, att); err != nil {
		return nil, fmt.Errorf("failed to check in: %w", err)
	}

	return s.repo.FindByID(ctx, att.ID.String())
}

func (s *service) CheckOut(ctx context.Context, req CheckOutRequest) (*Attendance, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	localNow := time.Now()
	nowUTC := localNow.UTC()
	today := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), 0, 0, 0, 0, time.UTC)

	att, err := s.repo.FindByEmployeeAndDate(ctx, empID, today)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no check-in record found for today")
		}
		return nil, err
	}

	if att.CheckIn == nil {
		return nil, errors.New("cannot check out without prior check-in")
	}

	att.CheckOut = &nowUTC

	// Calculate duration between CheckOut (nowUTC) and CheckIn (att.CheckIn)
	duration := nowUTC.Sub(att.CheckIn.UTC()).Hours()
	if duration < 0 {
		duration = 0
	}
	duration = math.Round(duration*100) / 100

	att.WorkHours = duration

	// Standard work shift = 8 hours; any excess is overtime
	if duration > 8.0 {
		att.OvertimeHours = math.Round((duration-8.0)*100) / 100
	} else {
		att.OvertimeHours = 0
	}

	if req.Notes != "" {
		if att.Notes != "" {
			att.Notes = att.Notes + " | " + req.Notes
		} else {
			att.Notes = req.Notes
		}
	}

	if err := s.repo.Update(ctx, att); err != nil {
		return nil, err
	}

	return att, nil
}

func (s *service) CreateAttendance(ctx context.Context, req CreateAttendanceRequest) (*Attendance, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	dateVal, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format (use YYYY-MM-DD)")
	}

	status := Status(req.Status)
	if status == "" {
		status = StatusPresent
	}

	att := &Attendance{
		EmployeeID: empID,
		Date:       dateVal,
		Status:     status,
		Notes:      req.Notes,
	}

	if req.CheckIn != "" {
		ci, err := time.Parse(time.RFC3339, req.CheckIn)
		if err == nil {
			ciUTC := ci.UTC()
			att.CheckIn = &ciUTC
		}
	}

	if req.CheckOut != "" {
		co, err := time.Parse(time.RFC3339, req.CheckOut)
		if err == nil {
			coUTC := co.UTC()
			att.CheckOut = &coUTC
		}
	}

	if att.CheckIn != nil && att.CheckOut != nil {
		dur := att.CheckOut.UTC().Sub(att.CheckIn.UTC()).Hours()
		if dur < 0 {
			dur = 0
		}
		dur = math.Round(dur*100) / 100
		att.WorkHours = dur
		if dur > 8.0 {
			att.OvertimeHours = math.Round((dur-8.0)*100) / 100
		}
	}

	if err := s.repo.Create(ctx, att); err != nil {
		return nil, fmt.Errorf("failed to create attendance record: %w", err)
	}

	return s.repo.FindByID(ctx, att.ID.String())
}

func (s *service) FindAll(ctx context.Context, req pagination.PageRequest) ([]Attendance, int64, error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) FindByID(ctx context.Context, id string) (*Attendance, error) {
	return s.repo.FindByID(ctx, id)
}
