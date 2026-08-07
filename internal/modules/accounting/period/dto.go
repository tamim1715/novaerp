package period

import (
	"time"

	"github.com/google/uuid"
)

type CreateFiscalYearRequest struct {
	Name      string    `json:"name" binding:"required,min=2,max=50"`
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
	AutoCreatePeriods bool `json:"autoCreatePeriods"` // if true, creates 12 monthly accounting periods automatically
}

type FiscalYearResponse struct {
	ID        uuid.UUID                  `json:"id"`
	Name      string                     `json:"name"`
	StartDate time.Time                  `json:"startDate"`
	EndDate   time.Time                  `json:"endDate"`
	IsClosed  bool                       `json:"isClosed"`
	Periods   []AccountingPeriodResponse `json:"periods,omitempty"`
	CreatedAt time.Time                  `json:"createdAt"`
	UpdatedAt time.Time                  `json:"updatedAt"`
}

type AccountingPeriodResponse struct {
	ID           uuid.UUID `json:"id"`
	FiscalYearID uuid.UUID `json:"fiscalYearId"`
	PeriodNumber int       `json:"periodNumber"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
