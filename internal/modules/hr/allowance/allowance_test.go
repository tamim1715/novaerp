package allowance

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

func TestToAllowanceTypeResponse(t *testing.T) {
	at := &AllowanceType{
		Name:            "Medical Allowance",
		Code:            "MED",
		CalculationType: CalculationFixed,
		DefaultAmount:   200.0,
		IsTaxable:       false,
		IsActive:        true,
	}
	at.ID = uuid.New()
	at.CreatedAt = time.Now()

	resp := ToAllowanceTypeResponse(at)
	if resp.Code != "MED" {
		t.Errorf("expected Code MED, got %s", resp.Code)
	}
	if resp.CalculationType != "FIXED" {
		t.Errorf("expected CalculationType FIXED, got %s", resp.CalculationType)
	}
}

func TestToEmployeeAllowanceResponse(t *testing.T) {
	empID := uuid.New()
	atID := uuid.New()

	ea := &EmployeeAllowance{
		EmployeeID:      empID,
		AllowanceTypeID: atID,
		Amount:          300.0,
		IsActive:        true,
		Employee:        employee.Employee{FirstName: "Alice", LastName: "Wonder"},
		AllowanceType:   AllowanceType{Name: "Housing", Code: "HRA"},
	}
	ea.ID = uuid.New()

	resp := ToEmployeeAllowanceResponse(ea)
	if resp.EmployeeName != "Alice Wonder" {
		t.Errorf("expected EmployeeName Alice Wonder, got %s", resp.EmployeeName)
	}
	if resp.AllowanceType == nil || resp.AllowanceType.Code != "HRA" {
		t.Errorf("expected AllowanceType Code HRA")
	}
}

// Mock repository for service calculation test
type mockAllowanceRepo struct {
	allowances []EmployeeAllowance
}

func (m *mockAllowanceRepo) CreateType(ctx context.Context, at *AllowanceType) error { return nil }
func (m *mockAllowanceRepo) FindAllTypes(ctx context.Context, req pagination.PageRequest) ([]AllowanceType, int64, error) {
	return nil, 0, nil
}
func (m *mockAllowanceRepo) FindTypeByID(ctx context.Context, id string) (*AllowanceType, error) {
	return nil, nil
}
func (m *mockAllowanceRepo) FindTypeByCode(ctx context.Context, code string) (*AllowanceType, error) {
	return nil, nil
}
func (m *mockAllowanceRepo) UpdateType(ctx context.Context, at *AllowanceType) error { return nil }
func (m *mockAllowanceRepo) DeleteType(ctx context.Context, id string) error         { return nil }
func (m *mockAllowanceRepo) AssignEmployeeAllowance(ctx context.Context, ea *EmployeeAllowance) error {
	return nil
}
func (m *mockAllowanceRepo) FindEmployeeAllowanceByID(ctx context.Context, id string) (*EmployeeAllowance, error) {
	return nil, nil
}
func (m *mockAllowanceRepo) FindActiveAllowancesByEmployeeID(ctx context.Context, employeeID uuid.UUID) ([]EmployeeAllowance, error) {
	return m.allowances, nil
}
func (m *mockAllowanceRepo) FindEmployeeAllowances(ctx context.Context, req pagination.PageRequest) ([]EmployeeAllowance, int64, error) {
	return nil, 0, nil
}
func (m *mockAllowanceRepo) DeleteEmployeeAllowance(ctx context.Context, id string) error { return nil }

func TestCalculateEmployeeAllowances(t *testing.T) {
	empID := uuid.New()
	mockRepo := &mockAllowanceRepo{
		allowances: []EmployeeAllowance{
			{
				EmployeeID: empID,
				Amount:     40.0, // 40%
				AllowanceType: AllowanceType{
					Code:            "HRA",
					Name:            "House Rent",
					CalculationType: CalculationPercentage,
				},
			},
			{
				EmployeeID: empID,
				Amount:     500.0, // Fixed 500
				AllowanceType: AllowanceType{
					Code:            "MED",
					Name:            "Medical",
					CalculationType: CalculationFixed,
				},
			},
		},
	}

	svc := NewService(mockRepo, nil, nil)
	basicSalary := 2000.00
	total, items, err := svc.CalculateEmployeeAllowances(context.Background(), empID, basicSalary)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 40% of 2000 = 800 + fixed 500 = 1300
	if total != 1300.00 {
		t.Errorf("expected total 1300.00, got %f", total)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 allowance items, got %d", len(items))
	}
	if items[0].CalculatedTotal != 800.00 {
		t.Errorf("expected HRA 800.00, got %f", items[0].CalculatedTotal)
	}
	if items[1].CalculatedTotal != 500.00 {
		t.Errorf("expected MED 500.00, got %f", items[1].CalculatedTotal)
	}
}
