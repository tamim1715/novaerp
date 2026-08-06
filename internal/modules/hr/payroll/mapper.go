package payroll

func ToPayrollPeriodResponse(p *PayrollPeriod) PayrollPeriodResponse {
	resp := PayrollPeriodResponse{
		ID:         p.ID.String(),
		Month:      p.Month,
		Year:       p.Year,
		StartDate:  p.StartDate.Format("2006-01-02"),
		EndDate:    p.EndDate.Format("2006-01-02"),
		Status:     p.Status,
		TotalGross: p.TotalGross,
		TotalNet:   p.TotalNet,
		CreatedAt:  p.CreatedAt.Unix(),
	}

	if p.ProcessedAt != nil {
		resp.ProcessedAt = p.ProcessedAt.Format("2006-01-02T15:04:05Z07:00")
	}

	return resp
}

func ToPayrollPeriodResponseList(list []PayrollPeriod) []PayrollPeriodResponse {
	resp := make([]PayrollPeriodResponse, len(list))
	for i, item := range list {
		resp[i] = ToPayrollPeriodResponse(&item)
	}
	return resp
}

func ToPayslipResponse(ps *Payslip) PayslipResponse {
	resp := PayslipResponse{
		ID:                  ps.ID.String(),
		PayrollPeriodID:     ps.PayrollPeriodID.String(),
		EmployeeID:          ps.EmployeeID.String(),
		BasicSalary:         ps.BasicSalary,
		Allowances:          ps.Allowances,
		Deductions:          ps.Deductions,
		UnpaidLeaveDeduction: ps.UnpaidLeaveDeduction,
		GrossSalary:         ps.GrossSalary,
		NetSalary:           ps.NetSalary,
		Status:              ps.Status,
		Notes:               ps.Notes,
		CreatedAt:           ps.CreatedAt.Unix(),
	}

	if ps.PaymentDate != nil {
		resp.PaymentDate = ps.PaymentDate.Format("2006-01-02T15:04:05Z07:00")
	}
	if ps.Employee.FirstName != "" {
		resp.EmployeeName = ps.Employee.FirstName + " " + ps.Employee.LastName
		resp.EmployeeDesignation = ps.Employee.Designation
	}

	return resp
}

func ToPayslipResponseList(list []Payslip) []PayslipResponse {
	resp := make([]PayslipResponse, len(list))
	for i, item := range list {
		resp[i] = ToPayslipResponse(&item)
	}
	return resp
}
