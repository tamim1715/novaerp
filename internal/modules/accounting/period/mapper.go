package period

func ToAccountingPeriodResponse(p *AccountingPeriod) AccountingPeriodResponse {
	if p == nil {
		return AccountingPeriodResponse{}
	}
	return AccountingPeriodResponse{
		ID:           p.ID,
		FiscalYearID: p.FiscalYearID,
		PeriodNumber: p.PeriodNumber,
		Name:         p.Name,
		StartDate:    p.StartDate,
		EndDate:      p.EndDate,
		Status:       p.Status.String(),
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func ToAccountingPeriodResponseList(periods []AccountingPeriod) []AccountingPeriodResponse {
	res := make([]AccountingPeriodResponse, len(periods))
	for i, p := range periods {
		res[i] = ToAccountingPeriodResponse(&p)
	}
	return res
}

func ToFiscalYearResponse(fy *FiscalYear) FiscalYearResponse {
	if fy == nil {
		return FiscalYearResponse{}
	}
	res := FiscalYearResponse{
		ID:        fy.ID,
		Name:      fy.Name,
		StartDate: fy.StartDate,
		EndDate:   fy.EndDate,
		IsClosed:  fy.IsClosed,
		CreatedAt: fy.CreatedAt,
		UpdatedAt: fy.UpdatedAt,
	}
	if len(fy.Periods) > 0 {
		res.Periods = ToAccountingPeriodResponseList(fy.Periods)
	}
	return res
}

func ToFiscalYearResponseList(fys []FiscalYear) []FiscalYearResponse {
	res := make([]FiscalYearResponse, len(fys))
	for i, fy := range fys {
		res[i] = ToFiscalYearResponse(&fy)
	}
	return res
}
