package attendance

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leaverequest"
	"github.com/tamim1715/novaerp/internal/modules/hr/overtime"
	"github.com/tamim1715/novaerp/internal/modules/hr/shift"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CheckIn(ctx context.Context, req CheckInRequest) (*Attendance, error)
	CheckOut(ctx context.Context, req CheckOutRequest) (*Attendance, error)
	CreateAttendance(ctx context.Context, req CreateAttendanceRequest) (*Attendance, error)
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Attendance, int64, error)
	FindByID(ctx context.Context, id string) (*Attendance, error)
	ProcessDailyAttendance(ctx context.Context, req ProcessDailyAttendanceRequest) (*DailyAttendanceProcessResult, error)
	GetMonthlySummary(ctx context.Context, employeeID uuid.UUID, startDate, endDate time.Time) (*MonthlyAttendanceSummary, error)
}

type service struct {
	repo             Repository
	employeeRepo     employee.Repository
	shiftService     shift.Service
	overtimeService  overtime.Service
	leaveRequestRepo leaverequest.Repository
	logger           *zap.Logger
}

func NewService(
	repo Repository,
	employeeRepo employee.Repository,
	shiftService shift.Service,
	overtimeService overtime.Service,
	leaveRequestRepo leaverequest.Repository,
	logger *zap.Logger,
) Service {
	return &service{
		repo:             repo,
		employeeRepo:     employeeRepo,
		shiftService:     shiftService,
		overtimeService:  overtimeService,
		leaveRequestRepo: leaveRequestRepo,
		logger:           logger,
	}
}

func (s *service) CheckIn(ctx context.Context, req CheckInRequest) (*Attendance, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	emp, err := s.employeeRepo.FindByID(ctx, req.EmployeeID)
	if err != nil {
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

	// Determine Shift and Late Status
	var shiftID *uuid.UUID
	status := StatusPresent
	lateMinutes := 0

	empShift, err := s.shiftService.FindActiveEmployeeShift(ctx, empID, today)
	if err == nil && empShift != nil && empShift.Shift.ID != uuid.Nil {
		sh := empShift.Shift
		shiftID = &sh.ID
		parts := strings.Split(sh.StartTime, ":")
		if len(parts) == 2 {
			shHour, _ := strconv.Atoi(parts[0])
			shMin, _ := strconv.Atoi(parts[1])
			shiftStartToday := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), shHour, shMin, 0, 0, localNow.Location())
			graceEnd := shiftStartToday.Add(time.Duration(sh.GracePeriodMinutes) * time.Minute)

			if localNow.After(graceEnd) {
				status = StatusLate
				diff := localNow.Sub(shiftStartToday)
				if diff > 0 {
					lateMinutes = int(diff.Minutes())
				}
			}
		}
	} else {
		// Fallback default: 09:30 AM
		if localNow.Hour() > 9 || (localNow.Hour() == 9 && localNow.Minute() > 30) {
			status = StatusLate
			lateMinutes = (localNow.Hour()-9)*60 + (localNow.Minute() - 30)
		}
	}

	att = &Attendance{
		EmployeeID:  empID,
		ShiftID:     shiftID,
		Date:        today,
		CheckIn:     &nowUTC,
		Status:      status,
		LateMinutes: lateMinutes,
		Notes:       req.Notes,
	}
	att.Employee = *emp

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

	duration := nowUTC.Sub(att.CheckIn.UTC()).Hours()
	if duration < 0 {
		duration = 0
	}
	duration = math.Round(duration*100) / 100
	att.WorkHours = duration

	// Work hours standard from assigned Shift
	standardHours := 8.0
	if att.Shift != nil && att.Shift.WorkHours > 0 {
		standardHours = att.Shift.WorkHours
	}

	// OT Policy calculation
	otPolicy, _ := s.overtimeService.GetDefaultPolicy(ctx)
	minOTMinutes := 30
	maxOTHours := 4.0
	if otPolicy != nil {
		if otPolicy.MinOvertimeMinutes > 0 {
			minOTMinutes = otPolicy.MinOvertimeMinutes
		}
		if otPolicy.MaxOvertimeHoursPerDay > 0 {
			maxOTHours = otPolicy.MaxOvertimeHoursPerDay
		}
	}

	if duration > standardHours {
		excessHours := duration - standardHours
		if excessHours*60 >= float64(minOTMinutes) {
			ot := excessHours
			if ot > maxOTHours {
				ot = maxOTHours
			}
			att.OvertimeHours = math.Round(ot*100) / 100
		} else {
			att.OvertimeHours = 0
		}
	} else {
		att.OvertimeHours = 0
	}

	// Check if Half Day
	halfDayThreshold := 4.0
	if att.Shift != nil && att.Shift.HalfDayHours > 0 {
		halfDayThreshold = att.Shift.HalfDayHours
	}
	if duration > 0 && duration < halfDayThreshold && att.Status != StatusLate {
		att.Status = StatusHalfDay
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

	var shiftID *uuid.UUID
	if req.ShiftID != "" {
		sid, err := uuid.Parse(req.ShiftID)
		if err == nil {
			shiftID = &sid
		}
	}

	att := &Attendance{
		EmployeeID: empID,
		ShiftID:    shiftID,
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

func (s *service) ProcessDailyAttendance(ctx context.Context, req ProcessDailyAttendanceRequest) (*DailyAttendanceProcessResult, error) {
	targetDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format (use YYYY-MM-DD)")
	}
	targetDate = time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), 0, 0, 0, 0, time.UTC)

	// Fetch active employees
	pageReq := pagination.PageRequest{Page: 1, Size: 1000}
	employees, total, err := s.employeeRepo.FindAll(ctx, pageReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employees: %w", err)
	}

	res := &DailyAttendanceProcessResult{
		Date:           req.Date,
		TotalEmployees: int(total),
	}

	isWeekend := (targetDate.Weekday() == time.Saturday || targetDate.Weekday() == time.Sunday)

	for _, emp := range employees {
		if req.DepartmentID != "" && emp.DepartmentID != req.DepartmentID {
			continue
		}

		att, err := s.repo.FindByEmployeeAndDate(ctx, emp.ID, targetDate)
		if err == nil && att != nil {
			// Record already exists
			res.ProcessedCount++
			switch att.Status {
			case StatusPresent:
				res.PresentCount++
			case StatusLate:
				res.LateCount++
			case StatusAbsent:
				res.AbsentCount++
			case StatusOnLeave:
				res.LeaveCount++
			}
			continue
		}

		// Check if employee has an approved leave request on targetDate
		leaveReq, err := s.leaveRequestRepo.FindApprovedByEmployeeAndDate(ctx, emp.ID, targetDate)
		if err == nil && leaveReq != nil {
			leaveAtt := &Attendance{
				EmployeeID: emp.ID,
				Date:       targetDate,
				Status:     StatusOnLeave,
				Notes:      fmt.Sprintf("Approved Leave (%s)", leaveReq.LeaveType.Name),
			}
			_ = s.repo.Create(ctx, leaveAtt)
			res.ProcessedCount++
			res.LeaveCount++
			continue
		}

		// If weekend, mark WEEKEND
		if isWeekend {
			weekendAtt := &Attendance{
				EmployeeID: emp.ID,
				Date:       targetDate,
				Status:     StatusWeekend,
				Notes:      "Weekly Weekend",
			}
			_ = s.repo.Create(ctx, weekendAtt)
			res.ProcessedCount++
			continue
		}

		// Otherwise, employee missed clock-in without approved leave -> mark ABSENT
		absentAtt := &Attendance{
			EmployeeID: emp.ID,
			Date:       targetDate,
			Status:     StatusAbsent,
			Notes:      "Marked absent by daily attendance process",
		}
		_ = s.repo.Create(ctx, absentAtt)
		res.ProcessedCount++
		res.AbsentCount++
	}

	return res, nil
}

func (s *service) GetMonthlySummary(ctx context.Context, employeeID uuid.UUID, startDate, endDate time.Time) (*MonthlyAttendanceSummary, error) {
	records, err := s.repo.FindByEmployeeAndDateRange(ctx, employeeID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	summary := &MonthlyAttendanceSummary{
		EmployeeID: employeeID.String(),
		StartDate:  startDate.Format("2006-01-02"),
		EndDate:    endDate.Format("2006-01-02"),
	}

	// Calculate calendar working days (excluding weekends)
	curr := startDate
	workingDaysCount := 0
	for !curr.After(endDate) {
		if curr.Weekday() != time.Saturday && curr.Weekday() != time.Sunday {
			workingDaysCount++
		}
		curr = curr.AddDate(0, 0, 1)
	}
	summary.TotalWorkingDays = workingDaysCount

	for _, rec := range records {
		summary.TotalWorkHours += rec.WorkHours
		summary.TotalOvertimeHours += rec.OvertimeHours

		switch rec.Status {
		case StatusPresent:
			summary.PresentDays++
		case StatusLate:
			summary.LateDays++
			summary.PresentDays++ // Late employees are still present for work
		case StatusHalfDay:
			summary.HalfDays++
		case StatusAbsent:
			summary.AbsentDays++
		case StatusOnLeave:
			summary.LeaveDays++
		}
	}

	summary.TotalWorkHours = math.Round(summary.TotalWorkHours*100) / 100
	summary.TotalOvertimeHours = math.Round(summary.TotalOvertimeHours*100) / 100

	return summary, nil
}
