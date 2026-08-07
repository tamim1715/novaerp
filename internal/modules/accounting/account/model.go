package account

import (
	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
)

// AccountType represents major accounting classifications.
const (
	TypeAsset     = "ASSET"
	TypeLiability = "LIABILITY"
	TypeEquity    = "EQUITY"
	TypeRevenue   = "REVENUE"
	TypeExpense   = "EXPENSE"
)

// NormalBalance determines whether debits or credits increase the account balance.
const (
	BalanceDebit  = "DEBIT"
	BalanceCredit = "CREDIT"
)

// Account represents a Chart of Accounts (COA) entry.
type Account struct {
	model.BaseModel

	Code          string     `gorm:"size:50;uniqueIndex;not null" json:"code"`
	Name          string     `gorm:"size:150;not null" json:"name"`
	Type          string     `gorm:"size:50;not null;index" json:"type"` // ASSET, LIABILITY, EQUITY, REVENUE, EXPENSE
	NormalBalance string     `gorm:"size:10;not null" json:"normalBalance"` // DEBIT, CREDIT
	ParentID      *uuid.UUID `gorm:"type:uuid;index" json:"parentId,omitempty"`
	Currency      string     `gorm:"size:10;default:'USD'" json:"currency"`
	IsActive      bool       `gorm:"default:true" json:"isActive"`
	IsSystem      bool       `gorm:"default:false" json:"isSystem"`
	Description   string     `gorm:"size:255" json:"description"`

	Parent   *Account  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Account `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}
