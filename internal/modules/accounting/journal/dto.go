package journal

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
)

type CreateJournalLineRequest struct {
	AccountID   uuid.UUID  `json:"accountId" binding:"required"`
	Debit       float64    `json:"debit" binding:"min=0"`
	Credit      float64    `json:"credit" binding:"min=0"`
	Description string     `json:"description" binding:"omitempty,max=255"`
	PartnerType string     `json:"partnerType" binding:"omitempty,oneof=NONE CUSTOMER VENDOR EMPLOYEE"`
	PartnerID   *uuid.UUID `json:"partnerId,omitempty"`
}

type CreateJournalEntryRequest struct {
	EntryDate   time.Time                  `json:"entryDate" binding:"required"`
	Reference   string                     `json:"reference" binding:"omitempty,max=100"`
	SourceType  string                     `json:"sourceType" binding:"omitempty,oneof=MANUAL PAYROLL INVOICE PAYMENT STOCK_ADJUSTMENT"`
	SourceID    *uuid.UUID                 `json:"sourceId,omitempty"`
	Description string                     `json:"description" binding:"required,min=3,max=255"`
	AutoPost    bool                       `json:"autoPost"`
	Lines       []CreateJournalLineRequest `json:"lines" binding:"required,min=2,dive"`
}

type VoidJournalEntryRequest struct {
	Reason string `json:"reason" binding:"required,min=3,max=255"`
}

type JournalEntryLineResponse struct {
	ID          uuid.UUID               `json:"id"`
	AccountID   uuid.UUID               `json:"accountId"`
	AccountCode string                  `json:"accountCode,omitempty"`
	AccountName string                  `json:"accountName,omitempty"`
	AccountType string                  `json:"accountType,omitempty"`
	Debit       float64                 `json:"debit"`
	Credit      float64                 `json:"credit"`
	Description string                  `json:"description"`
	PartnerType string                  `json:"partnerType"`
	PartnerID   *uuid.UUID              `json:"partnerId,omitempty"`
	Account     *account.AccountResponse `json:"account,omitempty"`
}

type JournalEntryResponse struct {
	ID           uuid.UUID                  `json:"id"`
	EntryNumber  string                     `json:"entryNumber"`
	EntryDate    time.Time                  `json:"entryDate"`
	PostingDate  *time.Time                 `json:"postingDate,omitempty"`
	PeriodID     *uuid.UUID                 `json:"periodId,omitempty"`
	Reference    string                     `json:"reference"`
	SourceType   string                     `json:"sourceType"`
	SourceID     *uuid.UUID                 `json:"sourceId,omitempty"`
	Description  string                     `json:"description"`
	Status       string                     `json:"status"`
	TotalDebit   float64                    `json:"totalDebit"`
	TotalCredit  float64                    `json:"totalCredit"`
	PostedBy     *uuid.UUID                 `json:"postedBy,omitempty"`
	VoidReason   string                     `json:"voidReason,omitempty"`
	ReversalOfID *uuid.UUID                 `json:"reversalOfId,omitempty"`
	Lines        []JournalEntryLineResponse `json:"lines,omitempty"`
	CreatedAt    time.Time                  `json:"createdAt"`
	UpdatedAt    time.Time                  `json:"updatedAt"`
}
