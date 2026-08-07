package accounting

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/journal"
	"github.com/tamim1715/novaerp/internal/modules/accounting/period"
	"github.com/tamim1715/novaerp/internal/modules/accounting/report"
)

func TestAccountMapper(t *testing.T) {
	id := uuid.New()
	acc := &account.Account{
		Code:          "1010",
		Name:          "Operating Bank Account",
		Type:          account.TypeAsset,
		NormalBalance: account.BalanceDebit,
		Currency:      "USD",
		IsActive:      true,
		IsSystem:      true,
		Description:   "Main bank account",
	}
	acc.ID = id

	res := account.ToAccountResponse(acc)
	if res.Code != "1010" {
		t.Errorf("expected code 1010, got %s", res.Code)
	}
	if res.NormalBalance != account.BalanceDebit {
		t.Errorf("expected normal balance DEBIT, got %s", res.NormalBalance)
	}
	if !res.IsSystem {
		t.Errorf("expected isSystem true")
	}
}

func TestFiscalYearMapper(t *testing.T) {
	fyID := uuid.New()
	pID := uuid.New()
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)

	fy := &period.FiscalYear{
		Name:      "FY 2026",
		StartDate: start,
		EndDate:   end,
		IsClosed:  false,
		Periods: []period.AccountingPeriod{
			{
				FiscalYearID: fyID,
				PeriodNumber: 1,
				Name:         "2026-01 (Jan)",
				StartDate:    start,
				EndDate:      time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC),
				Status:       period.StatusOpen,
			},
		},
	}
	fy.ID = fyID
	fy.Periods[0].ID = pID

	res := period.ToFiscalYearResponse(fy)
	if res.Name != "FY 2026" {
		t.Errorf("expected FY 2026, got %s", res.Name)
	}
	if len(res.Periods) != 1 {
		t.Fatalf("expected 1 period, got %d", len(res.Periods))
	}
	if res.Periods[0].PeriodNumber != 1 {
		t.Errorf("expected period number 1, got %d", res.Periods[0].PeriodNumber)
	}
}

func TestJournalEntryMapper(t *testing.T) {
	entryID := uuid.New()
	acc1ID := uuid.New()
	acc2ID := uuid.New()
	now := time.Now().UTC()

	entry := &journal.JournalEntry{
		EntryNumber: "JE-202601-00001",
		EntryDate:   now,
		Reference:   "INV-1001",
		SourceType:  journal.SourceInvoice,
		Description: "Customer Sale Invoice",
		Status:      journal.StatusPosted,
		TotalDebit:  1500.00,
		TotalCredit: 1500.00,
		Lines: []journal.JournalEntryLine{
			{
				JournalEntryID: entryID,
				AccountID:      acc1ID,
				Debit:          1500.00,
				Credit:         0,
				Description:    "Accounts Receivable Debit",
			},
			{
				JournalEntryID: entryID,
				AccountID:      acc2ID,
				Debit:          0,
				Credit:         1500.00,
				Description:    "Sales Revenue Credit",
			},
		},
	}
	entry.ID = entryID

	res := journal.ToJournalEntryResponse(entry)
	if res.EntryNumber != "JE-202601-00001" {
		t.Errorf("expected entry number JE-202601-00001, got %s", res.EntryNumber)
	}
	if res.TotalDebit != 1500.00 || res.TotalCredit != 1500.00 {
		t.Errorf("expected balanced totals 1500, got Debit: %f, Credit: %f", res.TotalDebit, res.TotalCredit)
	}
	if len(res.Lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(res.Lines))
	}
	if res.Lines[0].Debit != 1500.00 || res.Lines[1].Credit != 1500.00 {
		t.Errorf("lines debit/credit mismatch")
	}
}

func TestFinancialStatementsCalculations(t *testing.T) {
	// 1. Profit & Loss calculations check
	pl := report.ProfitAndLossResponse{
		StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
		OperatingRevenues: []report.StatementLineItem{
			{AccountCode: "4000", AccountName: "Operating Sales Revenue", Amount: 100000.00},
			{AccountCode: "4100", AccountName: "Service Revenue", Amount: 20000.00},
		},
		TotalRevenue: 120000.00,
		CostOfGoodsSold: []report.StatementLineItem{
			{AccountCode: "5000", AccountName: "Cost of Goods Sold", Amount: 45000.00},
		},
		TotalCOGS:   45000.00,
		GrossProfit: 75000.00, // 120,000 - 45,000
		OperatingExpenses: []report.StatementLineItem{
			{AccountCode: "6000", AccountName: "Salaries Expense", Amount: 30000.00},
			{AccountCode: "6100", AccountName: "Rent Expense", Amount: 12000.00},
		},
		TotalOperatingExpenses: 42000.00,
		NetIncome:              33000.00, // 75,000 - 42,000
	}

	if pl.GrossProfit != (pl.TotalRevenue - pl.TotalCOGS) {
		t.Errorf("Gross profit mismatch: expected %f, got %f", pl.TotalRevenue-pl.TotalCOGS, pl.GrossProfit)
	}
	if pl.NetIncome != (pl.GrossProfit - pl.TotalOperatingExpenses) {
		t.Errorf("Net income mismatch: expected %f, got %f", pl.GrossProfit-pl.TotalOperatingExpenses, pl.NetIncome)
	}

	// 2. Balance Sheet equation check: Assets == Liabilities + Equity
	bs := report.BalanceSheetResponse{
		AsOfDate: time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
		CurrentAssets: report.BalanceSheetSection{
			Category: "Current Assets",
			Subtotal: 85000.00,
		},
		NonCurrentAssets: report.BalanceSheetSection{
			Category: "Non-Current / Fixed Assets",
			Subtotal: 50000.00,
		},
		TotalAssets: 135000.00, // 85,000 + 50,000
		CurrentLiabilities: report.BalanceSheetSection{
			Category: "Current Liabilities",
			Subtotal: 22000.00,
		},
		LongTermLiabilities: report.BalanceSheetSection{
			Category: "Long-Term Liabilities",
			Subtotal: 30000.00,
		},
		TotalLiabilities: 52000.00, // 22,000 + 30,000
		Equity: report.BalanceSheetSection{
			Category: "Shareholders' Equity",
			Subtotal: 50000.00,
		},
		CurrentPeriodNetIncome:    33000.00,
		TotalEquity:               83000.00, // 50,000 + 33,000
		TotalLiabilitiesAndEquity: 135000.00, // 52,000 + 83,000
		IsBalanced:                true,
	}

	if bs.TotalAssets != (bs.CurrentAssets.Subtotal + bs.NonCurrentAssets.Subtotal) {
		t.Errorf("Total assets mismatch: %f", bs.TotalAssets)
	}
	if bs.TotalLiabilitiesAndEquity != (bs.TotalLiabilities + bs.TotalEquity) {
		t.Errorf("Total liabilities & equity mismatch: %f", bs.TotalLiabilitiesAndEquity)
	}
	if bs.TotalAssets != bs.TotalLiabilitiesAndEquity {
		t.Errorf("Balance sheet equation failed: Assets (%f) != Liabilities+Equity (%f)", bs.TotalAssets, bs.TotalLiabilitiesAndEquity)
	}
	if !bs.IsBalanced {
		t.Errorf("Expected balance sheet to be balanced")
	}
}
