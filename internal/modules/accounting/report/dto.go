package report

import (
	"time"

	"github.com/google/uuid"
)

// General Ledger Statement DTOs
type LedgerTransactionItem struct {
	Date           time.Time  `json:"date"`
	EntryNumber    string     `json:"entryNumber"`
	JournalEntryID uuid.UUID  `json:"journalEntryId"`
	Reference      string     `json:"reference"`
	Description    string     `json:"description"`
	Debit          float64    `json:"debit"`
	Credit         float64    `json:"credit"`
	RunningBalance float64    `json:"runningBalance"`
}

type GeneralLedgerAccountReport struct {
	AccountID         uuid.UUID               `json:"accountId"`
	AccountCode       string                  `json:"accountCode"`
	AccountName       string                  `json:"accountName"`
	AccountType       string                  `json:"accountType"`
	NormalBalance     string                  `json:"normalBalance"`
	OpeningBalance    float64                 `json:"openingBalance"`
	PeriodDebitTotal  float64                 `json:"periodDebitTotal"`
	PeriodCreditTotal float64                 `json:"periodCreditTotal"`
	ClosingBalance    float64                 `json:"closingBalance"`
	Transactions      []LedgerTransactionItem `json:"transactions"`
}

type GeneralLedgerResponse struct {
	StartDate time.Time                    `json:"startDate"`
	EndDate   time.Time                    `json:"endDate"`
	Accounts  []GeneralLedgerAccountReport `json:"accounts"`
}

// Trial Balance DTOs
type TrialBalanceItem struct {
	AccountID     uuid.UUID `json:"accountId"`
	AccountCode   string    `json:"accountCode"`
	AccountName   string    `json:"accountName"`
	AccountType   string    `json:"accountType"`
	NormalBalance string    `json:"normalBalance"`
	DebitTurnover float64   `json:"debitTurnover"`
	CreditTurnover float64  `json:"creditTurnover"`
	NetDebit      float64   `json:"netDebit"`
	NetCredit     float64   `json:"netCredit"`
}

type TrialBalanceResponse struct {
	AsOfDate     time.Time          `json:"asOfDate"`
	TotalDebits  float64            `json:"totalDebits"`
	TotalCredits float64            `json:"totalCredits"`
	Difference   float64            `json:"difference"`
	IsBalanced   bool               `json:"isBalanced"`
	Items        []TrialBalanceItem `json:"items"`
}

// Profit & Loss DTOs
type StatementLineItem struct {
	AccountCode string  `json:"accountCode"`
	AccountName string  `json:"accountName"`
	Amount      float64 `json:"amount"`
}

type ProfitAndLossResponse struct {
	StartDate              time.Time           `json:"startDate"`
	EndDate                time.Time           `json:"endDate"`
	OperatingRevenues      []StatementLineItem `json:"operatingRevenues"`
	TotalRevenue           float64             `json:"totalRevenue"`
	CostOfGoodsSold        []StatementLineItem `json:"costOfGoodsSold"`
	TotalCOGS              float64             `json:"totalCogs"`
	GrossProfit            float64             `json:"grossProfit"`
	OperatingExpenses      []StatementLineItem `json:"operatingExpenses"`
	TotalOperatingExpenses float64             `json:"totalOperatingExpenses"`
	NetIncome              float64             `json:"netIncome"` // GrossProfit - OperatingExpenses
}

// Balance Sheet DTOs
type BalanceSheetSection struct {
	Category string              `json:"category"`
	Items    []StatementLineItem `json:"items"`
	Subtotal float64             `json:"subtotal"`
}

type BalanceSheetResponse struct {
	AsOfDate                  time.Time             `json:"asOfDate"`
	CurrentAssets             BalanceSheetSection   `json:"currentAssets"`
	NonCurrentAssets          BalanceSheetSection   `json:"nonCurrentAssets"`
	TotalAssets               float64               `json:"totalAssets"`
	CurrentLiabilities        BalanceSheetSection   `json:"currentLiabilities"`
	LongTermLiabilities       BalanceSheetSection   `json:"longTermLiabilities"`
	TotalLiabilities          float64               `json:"totalLiabilities"`
	Equity                    BalanceSheetSection   `json:"equity"`
	CurrentPeriodNetIncome    float64               `json:"currentPeriodNetIncome"`
	TotalEquity               float64               `json:"totalEquity"`
	TotalLiabilitiesAndEquity float64               `json:"totalLiabilitiesAndEquity"`
	IsBalanced                bool                  `json:"isBalanced"` // TotalAssets == TotalLiabilitiesAndEquity
}
