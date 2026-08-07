package account

import (
	"errors"

	"gorm.io/gorm"
)

// SeedStandardChartOfAccounts seeds standard GAAP/IFRS accounts if none exist.
func SeedStandardChartOfAccounts(db *gorm.DB) error {
	var count int64
	if err := db.Model(&Account{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil // already seeded
	}

	standardAccounts := []Account{
		// ASSETS (1000s)
		{Code: "1000", Name: "Current Assets", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Parent group for liquid assets"},
		{Code: "1010", Name: "Cash on Hand", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Physical cash and petty cash"},
		{Code: "1020", Name: "Operating Bank Account", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Main operational checking account"},
		{Code: "1200", Name: "Accounts Receivable", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Customer invoices pending payment"},
		{Code: "1300", Name: "Inventory Asset", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Current valuation of merchandise/goods in warehouse"},
		{Code: "1500", Name: "Property, Plant & Equipment", Type: TypeAsset, NormalBalance: BalanceDebit, IsSystem: true, Description: "Fixed capital assets"},

		// LIABILITIES (2000s)
		{Code: "2000", Name: "Current Liabilities", Type: TypeLiability, NormalBalance: BalanceCredit, IsSystem: true, Description: "Parent group for short-term liabilities"},
		{Code: "2010", Name: "Accounts Payable", Type: TypeLiability, NormalBalance: BalanceCredit, IsSystem: true, Description: "Outstanding supplier and vendor bills"},
		{Code: "2100", Name: "Accrued Salaries Payable", Type: TypeLiability, NormalBalance: BalanceCredit, IsSystem: true, Description: "Wages accrued for payroll processing"},
		{Code: "2200", Name: "Sales Tax / VAT Payable", Type: TypeLiability, NormalBalance: BalanceCredit, IsSystem: true, Description: "Taxes collected on sales to be remitted to government"},

		// EQUITY (3000s)
		{Code: "3000", Name: "Owner's Equity / Share Capital", Type: TypeEquity, NormalBalance: BalanceCredit, IsSystem: true, Description: "Capital contributed by shareholders"},
		{Code: "3100", Name: "Retained Earnings", Type: TypeEquity, NormalBalance: BalanceCredit, IsSystem: true, Description: "Accumulated net earnings retained in business"},

		// REVENUE (4000s)
		{Code: "4000", Name: "Operating Sales Revenue", Type: TypeRevenue, NormalBalance: BalanceCredit, IsSystem: true, Description: "Income generated from product sales"},
		{Code: "4100", Name: "Service & Consulting Revenue", Type: TypeRevenue, NormalBalance: BalanceCredit, IsSystem: true, Description: "Income from professional services rendered"},
		{Code: "4200", Name: "Other Income & Discounts Earned", Type: TypeRevenue, NormalBalance: BalanceCredit, IsSystem: true, Description: "Non-operating secondary revenues"},

		// EXPENSES (5000s - 6000s)
		{Code: "5000", Name: "Cost of Goods Sold (COGS)", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "Direct cost of products sold"},
		{Code: "6000", Name: "Salaries & Wages Expense", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "Employee payroll and compensation"},
		{Code: "6100", Name: "Rent & Facilities Expense", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "Office and warehouse lease payments"},
		{Code: "6200", Name: "Utilities Expense (Electricity, Water, Internet)", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "Utility and internet bills"},
		{Code: "6300", Name: "Office Supplies & Maintenance", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "General office supplies and upkeep"},
		{Code: "6400", Name: "Depreciation Expense", Type: TypeExpense, NormalBalance: BalanceDebit, IsSystem: true, Description: "Amortization and asset depreciation"},
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, acc := range standardAccounts {
			if err := tx.Create(&acc).Error; err != nil {
				return errors.New("failed to seed chart of accounts: " + err.Error())
			}
		}
		return nil
	})
}
