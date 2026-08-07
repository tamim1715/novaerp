package report

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/journal"
	"go.uber.org/zap"
)

type Service interface {
	GetGeneralLedger(ctx context.Context, accountID *uuid.UUID, startDate, endDate time.Time) (*GeneralLedgerResponse, error)
	GetTrialBalance(ctx context.Context, asOfDate time.Time) (*TrialBalanceResponse, error)
	GetProfitAndLoss(ctx context.Context, startDate, endDate time.Time) (*ProfitAndLossResponse, error)
	GetBalanceSheet(ctx context.Context, asOfDate time.Time) (*BalanceSheetResponse, error)
}

type service struct {
	accountRepo account.Repository
	journalRepo journal.Repository
	logger      *zap.Logger
}

func NewService(accountRepo account.Repository, journalRepo journal.Repository, logger *zap.Logger) Service {
	return &service{
		accountRepo: accountRepo,
		journalRepo: journalRepo,
		logger:      logger,
	}
}

func round(val float64) float64 {
	return math.Round(val*100) / 100
}

func (s *service) GetGeneralLedger(ctx context.Context, accountID *uuid.UUID, startDate, endDate time.Time) (*GeneralLedgerResponse, error) {
	var accounts []account.Account
	if accountID != nil {
		acc, err := s.accountRepo.FindByID(ctx, accountID.String())
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, *acc)
	} else {
		all, _, err := s.accountRepo.FindAll(ctx, pagination.PageRequest{Page: 1, Size: 500, SortBy: "code", Order: "asc"}, "")
		if err != nil {
			return nil, err
		}
		accounts = all
	}

	lines, err := s.journalRepo.GetGeneralLedgerLines(ctx, accountID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	linesByAccount := make(map[uuid.UUID][]journal.JournalEntryLine)
	for _, l := range lines {
		linesByAccount[l.AccountID] = append(linesByAccount[l.AccountID], l)
	}

	res := &GeneralLedgerResponse{
		StartDate: startDate,
		EndDate:   endDate,
		Accounts:  make([]GeneralLedgerAccountReport, 0, len(accounts)),
	}

	for _, acc := range accounts {
		openDebit, openCredit, err := s.journalRepo.GetOpeningBalance(ctx, acc.ID, startDate)
		if err != nil {
			return nil, err
		}

		var openingBal float64
		if acc.NormalBalance == account.BalanceDebit {
			openingBal = openDebit - openCredit
		} else {
			openingBal = openCredit - openDebit
		}
		openingBal = round(openingBal)

		runningBal := openingBal
		var periodDebit, periodCredit float64
		var txItems []LedgerTransactionItem

		for _, l := range linesByAccount[acc.ID] {
			periodDebit += l.Debit
			periodCredit += l.Credit

			if acc.NormalBalance == account.BalanceDebit {
				runningBal += (l.Debit - l.Credit)
			} else {
				runningBal += (l.Credit - l.Debit)
			}
			runningBal = round(runningBal)

			entry, _ := s.journalRepo.FindByID(ctx, l.JournalEntryID.String())
			entryNum := ""
			ref := ""
			entryDate := l.CreatedAt
			if entry != nil {
				entryNum = entry.EntryNumber
				ref = entry.Reference
				entryDate = entry.EntryDate
			}

			txItems = append(txItems, LedgerTransactionItem{
				Date:           entryDate,
				EntryNumber:    entryNum,
				JournalEntryID: l.JournalEntryID,
				Reference:      ref,
				Description:    l.Description,
				Debit:          round(l.Debit),
				Credit:         round(l.Credit),
				RunningBalance: runningBal,
			})
		}

		closingBal := runningBal

		// Include account if it has opening balance or transactions
		if openingBal != 0 || len(txItems) > 0 || accountID != nil {
			res.Accounts = append(res.Accounts, GeneralLedgerAccountReport{
				AccountID:         acc.ID,
				AccountCode:       acc.Code,
				AccountName:       acc.Name,
				AccountType:       acc.Type,
				NormalBalance:     acc.NormalBalance,
				OpeningBalance:    openingBal,
				PeriodDebitTotal:  round(periodDebit),
				PeriodCreditTotal: round(periodCredit),
				ClosingBalance:    closingBal,
				Transactions:      txItems,
			})
		}
	}

	return res, nil
}

func (s *service) GetTrialBalance(ctx context.Context, asOfDate time.Time) (*TrialBalanceResponse, error) {
	accounts, _, err := s.accountRepo.FindAll(ctx, pagination.PageRequest{Page: 1, Size: 500, SortBy: "code", Order: "asc"}, "")
	if err != nil {
		return nil, err
	}

	res := &TrialBalanceResponse{
		AsOfDate: asOfDate,
		Items:    make([]TrialBalanceItem, 0, len(accounts)),
	}

	var totalDebits, totalCredits float64

	for _, acc := range accounts {
		debit, credit, err := s.journalRepo.GetOpeningBalance(ctx, acc.ID, asOfDate.AddDate(0, 0, 1))
		if err != nil {
			return nil, err
		}

		if debit == 0 && credit == 0 {
			continue
		}

		var netDebit, netCredit float64
		if debit >= credit {
			netDebit = round(debit - credit)
		} else {
			netCredit = round(credit - debit)
		}

		totalDebits += debit
		totalCredits += credit

		res.Items = append(res.Items, TrialBalanceItem{
			AccountID:      acc.ID,
			AccountCode:    acc.Code,
			AccountName:    acc.Name,
			AccountType:    acc.Type,
			NormalBalance:  acc.NormalBalance,
			DebitTurnover:  round(debit),
			CreditTurnover: round(credit),
			NetDebit:       netDebit,
			NetCredit:      netCredit,
		})
	}

	res.TotalDebits = round(totalDebits)
	res.TotalCredits = round(totalCredits)
	res.Difference = round(math.Abs(res.TotalDebits - res.TotalCredits))
	res.IsBalanced = res.Difference < 0.01

	return res, nil
}

func (s *service) GetProfitAndLoss(ctx context.Context, startDate, endDate time.Time) (*ProfitAndLossResponse, error) {
	lines, err := s.journalRepo.GetGeneralLedgerLines(ctx, nil, startDate, endDate)
	if err != nil {
		return nil, err
	}

	revenueMap := make(map[string]*StatementLineItem)
	cogsMap := make(map[string]*StatementLineItem)
	expenseMap := make(map[string]*StatementLineItem)

	for _, l := range lines {
		if l.Account.ID == uuid.Nil {
			acc, err := s.accountRepo.FindByID(ctx, l.AccountID.String())
			if err == nil && acc != nil {
				l.Account = *acc
			}
		}

		if l.Account.Type == account.TypeRevenue {
			netRev := l.Credit - l.Debit
			if item, exists := revenueMap[l.Account.Code]; exists {
				item.Amount += netRev
			} else {
				revenueMap[l.Account.Code] = &StatementLineItem{
					AccountCode: l.Account.Code,
					AccountName: l.Account.Name,
					Amount:      netRev,
				}
			}
		} else if l.Account.Type == account.TypeExpense {
			netExp := l.Debit - l.Credit
			if strings.HasPrefix(l.Account.Code, "5") || strings.Contains(strings.ToLower(l.Account.Name), "cost of goods") {
				if item, exists := cogsMap[l.Account.Code]; exists {
					item.Amount += netExp
				} else {
					cogsMap[l.Account.Code] = &StatementLineItem{
						AccountCode: l.Account.Code,
						AccountName: l.Account.Name,
						Amount:      netExp,
					}
				}
			} else {
				if item, exists := expenseMap[l.Account.Code]; exists {
					item.Amount += netExp
				} else {
					expenseMap[l.Account.Code] = &StatementLineItem{
						AccountCode: l.Account.Code,
						AccountName: l.Account.Name,
						Amount:      netExp,
					}
				}
			}
		}
	}

	res := &ProfitAndLossResponse{
		StartDate: startDate,
		EndDate:   endDate,
	}

	for _, v := range revenueMap {
		v.Amount = round(v.Amount)
		res.OperatingRevenues = append(res.OperatingRevenues, *v)
		res.TotalRevenue += v.Amount
	}
	for _, v := range cogsMap {
		v.Amount = round(v.Amount)
		res.CostOfGoodsSold = append(res.CostOfGoodsSold, *v)
		res.TotalCOGS += v.Amount
	}
	for _, v := range expenseMap {
		v.Amount = round(v.Amount)
		res.OperatingExpenses = append(res.OperatingExpenses, *v)
		res.TotalOperatingExpenses += v.Amount
	}

	res.TotalRevenue = round(res.TotalRevenue)
	res.TotalCOGS = round(res.TotalCOGS)
	res.GrossProfit = round(res.TotalRevenue - res.TotalCOGS)
	res.TotalOperatingExpenses = round(res.TotalOperatingExpenses)
	res.NetIncome = round(res.GrossProfit - res.TotalOperatingExpenses)

	return res, nil
}

func (s *service) GetBalanceSheet(ctx context.Context, asOfDate time.Time) (*BalanceSheetResponse, error) {
	accounts, _, err := s.accountRepo.FindAll(ctx, pagination.PageRequest{Page: 1, Size: 500, SortBy: "code", Order: "asc"}, "")
	if err != nil {
		return nil, err
	}

	res := &BalanceSheetResponse{
		AsOfDate:            asOfDate,
		CurrentAssets:       BalanceSheetSection{Category: "Current Assets"},
		NonCurrentAssets:    BalanceSheetSection{Category: "Non-Current / Fixed Assets"},
		CurrentLiabilities:  BalanceSheetSection{Category: "Current Liabilities"},
		LongTermLiabilities: BalanceSheetSection{Category: "Long-Term Liabilities"},
		Equity:              BalanceSheetSection{Category: "Shareholders' Equity"},
	}

	for _, acc := range accounts {
		debit, credit, err := s.journalRepo.GetOpeningBalance(ctx, acc.ID, asOfDate.AddDate(0, 0, 1))
		if err != nil {
			return nil, err
		}

		if debit == 0 && credit == 0 {
			continue
		}

		switch acc.Type {
		case account.TypeAsset:
			netAsset := round(debit - credit)
			item := StatementLineItem{AccountCode: acc.Code, AccountName: acc.Name, Amount: netAsset}
			if strings.HasPrefix(acc.Code, "15") || strings.Contains(strings.ToLower(acc.Name), "equipment") || strings.Contains(strings.ToLower(acc.Name), "property") {
				res.NonCurrentAssets.Items = append(res.NonCurrentAssets.Items, item)
				res.NonCurrentAssets.Subtotal += netAsset
			} else {
				res.CurrentAssets.Items = append(res.CurrentAssets.Items, item)
				res.CurrentAssets.Subtotal += netAsset
			}
			res.TotalAssets += netAsset

		case account.TypeLiability:
			netLiab := round(credit - debit)
			item := StatementLineItem{AccountCode: acc.Code, AccountName: acc.Name, Amount: netLiab}
			if strings.HasPrefix(acc.Code, "25") || strings.Contains(strings.ToLower(acc.Name), "long-term") {
				res.LongTermLiabilities.Items = append(res.LongTermLiabilities.Items, item)
				res.LongTermLiabilities.Subtotal += netLiab
			} else {
				res.CurrentLiabilities.Items = append(res.CurrentLiabilities.Items, item)
				res.CurrentLiabilities.Subtotal += netLiab
			}
			res.TotalLiabilities += netLiab

		case account.TypeEquity:
			netEq := round(credit - debit)
			item := StatementLineItem{AccountCode: acc.Code, AccountName: acc.Name, Amount: netEq}
			res.Equity.Items = append(res.Equity.Items, item)
			res.Equity.Subtotal += netEq
			res.TotalEquity += netEq
		}
	}

	// Calculate Current Period Net Income from start of year to asOfDate
	yearStart := time.Date(asOfDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	pl, err := s.GetProfitAndLoss(ctx, yearStart, asOfDate)
	if err == nil && pl != nil {
		res.CurrentPeriodNetIncome = pl.NetIncome
		res.TotalEquity += pl.NetIncome
	}

	res.CurrentAssets.Subtotal = round(res.CurrentAssets.Subtotal)
	res.NonCurrentAssets.Subtotal = round(res.NonCurrentAssets.Subtotal)
	res.TotalAssets = round(res.TotalAssets)

	res.CurrentLiabilities.Subtotal = round(res.CurrentLiabilities.Subtotal)
	res.LongTermLiabilities.Subtotal = round(res.LongTermLiabilities.Subtotal)
	res.TotalLiabilities = round(res.TotalLiabilities)

	res.Equity.Subtotal = round(res.Equity.Subtotal)
	res.TotalEquity = round(res.TotalEquity)

	res.TotalLiabilitiesAndEquity = round(res.TotalLiabilities + res.TotalEquity)
	res.IsBalanced = math.Abs(res.TotalAssets-res.TotalLiabilitiesAndEquity) < 0.01

	return res, nil
}
