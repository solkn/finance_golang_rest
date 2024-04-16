package models

import (
	"time"

	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TransactionType int

const (
	TransactionTypeExpense  TransactionType = 0
	TransactionTypeIncome   TransactionType = 1
	TransactionTypeTransfer TransactionType = 2
)

// Transaction represents a single Quicken transaction
type Transaction struct {
	BaseModel
	LineNo                uint                `gorm:"not null; serial;AUTO_INCREMENT"`
	TransactionDate       time.Time           `gorm:"not null; default:current_timestamp"`
	TransactionEntryDate  time.Time           `gorm:"not null; default:current_timestamp"`
	TransactionType       string              `gorm:"not null; default:'E'; type:varchar(2)"`
	TransactionCategory   TransactionCategory `json:"transactioncategory" gorm:"foreignKey:TransactionCategoryID"`
	TransactionCategoryID v7uuid.UUID         `json:"transactioncategoryid" gorm:"NOT NULL;type:uuid"`
	Payee                 *Payee              `json:"payee"  gorm:"foreignKey:PayeeID"`
	PayeeID               v7uuid.UUID         `json:"payeeid" gorm:"NOT NULL;type:uuid"`
	Deposit               float64             `gorm:"type:NUMERIC(18,2); default:0"`
	Payment               float64             `gorm:"type:NUMERIC(18,2); default:0"`
	Currency              string              `gorm:"not null"`
	Memo                  string              // Optional field for additional notes
	ReferenceNumber       string              // Optional field for reference number
	AccountNumber         string              // Optional field for account number
	Clr                   bool                `json:"clr" gorm:"default:false"`
	Status                string              `gorm:"not null;default:Uncleared"` // Default status is "Cleared"
	ReconciledDate        *time.Time          // Optional field for reconciled date
	Tags                  []*TxTag            `gorm:"txtag"  gorm:"foreignKey:TransactionID"`
	TagID                 v7uuid.UUID         `json:"tagid" gorm:"NOT NULL;type:uuid"`
	Register              *Register           `json:"register"`
	RegisterID            v7uuid.UUID         `json:"registerid" gorm:"NOT NULL;type:uuid"`
	TxLines               []*TxLines          `gorm:"txlines"  gorm:"foreignKey:TransactionID"`
	TxLinesID             v7uuid.UUID         `json:"txlinesid"`
}

func (cmp *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	if time.Now().Sub(cmp.TransactionDate).Hours()/24 > 700 {
		cmp.TransactionDate = time.Now()
	}
	return
}
