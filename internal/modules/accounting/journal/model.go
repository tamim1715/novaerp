package journal

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/period"
)

// Journal status lifecycle
const (
	StatusDraft  = "DRAFT"
	StatusPosted = "POSTED"
	StatusVoid   = "VOID"
)

// Source types generating journal transactions
const (
	SourceManual          = "MANUAL"
	SourcePayroll         = "PAYROLL"
	SourceInvoice         = "INVOICE"
	SourcePayment         = "PAYMENT"
	SourceStockAdjustment = "STOCK_ADJUSTMENT"
	SourceReversal        = "REVERSAL"
)

// Partner types linked to journal line items
const (
	PartnerNone     = "NONE"
	PartnerCustomer = "CUSTOMER"
	PartnerVendor   = "VENDOR"
	PartnerEmployee = "EMPLOYEE"
)

// JournalEntry represents a balanced double-entry accounting transaction header.
type JournalEntry struct {
	model.BaseModel

	EntryNumber  string     `gorm:"size:50;uniqueIndex;not null" json:"entryNumber"`
	EntryDate    time.Time  `gorm:"type:date;not null" json:"entryDate"`
	PostingDate  *time.Time `gorm:"type:timestamp" json:"postingDate,omitempty"`
	PeriodID     *uuid.UUID `gorm:"type:uuid;index" json:"periodId,omitempty"`
	Reference    string     `gorm:"size:100" json:"reference"`
	SourceType   string     `gorm:"size:50;default:'MANUAL';not null" json:"sourceType"`
	SourceID     *uuid.UUID `gorm:"type:uuid;index" json:"sourceId,omitempty"`
	Description  string     `gorm:"size:255;not null" json:"description"`
	Status       string     `gorm:"size:20;default:'DRAFT';not null;index" json:"status"` // DRAFT, POSTED, VOID
	TotalDebit   float64    `gorm:"type:numeric(14,2);default:0;not null" json:"totalDebit"`
	TotalCredit  float64    `gorm:"type:numeric(14,2);default:0;not null" json:"totalCredit"`
	PostedBy     *uuid.UUID `gorm:"type:uuid" json:"postedBy,omitempty"`
	VoidReason   string     `gorm:"size:255" json:"voidReason,omitempty"`
	ReversalOfID *uuid.UUID `gorm:"type:uuid" json:"reversalOfId,omitempty"`

	Lines  []JournalEntryLine      `gorm:"foreignKey:JournalEntryID;constraint:OnDelete:CASCADE" json:"lines,omitempty"`
	Period *period.AccountingPeriod `gorm:"foreignKey:PeriodID" json:"period,omitempty"`
}

// JournalEntryLine represents an individual debit or credit line on an account.
type JournalEntryLine struct {
	model.BaseModel

	JournalEntryID uuid.UUID  `gorm:"type:uuid;not null;index" json:"journalEntryId"`
	AccountID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"accountId"`
	Debit          float64    `gorm:"type:numeric(14,2);default:0;not null" json:"debit"`
	Credit         float64    `gorm:"type:numeric(14,2);default:0;not null" json:"credit"`
	Description    string     `gorm:"size:255" json:"description"`
	PartnerType    string     `gorm:"size:50;default:'NONE'" json:"partnerType"` // NONE, CUSTOMER, VENDOR, EMPLOYEE
	PartnerID      *uuid.UUID `gorm:"type:uuid" json:"partnerId,omitempty"`

	Account account.Account `gorm:"foreignKey:AccountID" json:"account,omitempty"`
}
