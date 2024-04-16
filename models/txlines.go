package models

import (
	"time"

	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Transaction represents a single Quicken transaction
type TxLines struct {
	BaseModel
	TransactionID    v7uuid.UUID         `json:"transactionid" gorm:"NOT NULL;type:uuid"`
	LineNo           uint                `gorm:"not null; serial;AUTO_INCREMENT"`
	TxLineDate       time.Time           `gorm:"not null; default:current_timestamp"`
	TxLineEntryDate  time.Time           `gorm:"not null; default:current_timestamp"`
	TxLineCategory   TransactionCategory `json:"txlinecategory" gorm:"foreignKey:TxLineCategoryID"`
	TxLineCategoryID v7uuid.UUID
	TxLinePayee      *Payee      `json:"txlinepayee"  gorm:"foreignKey:TxLinePayeeID"`
	TxLinePayeeID    v7uuid.UUID `json:"txlinepayeeid" `
	Deposit  float64 `gorm:"type:NUMERIC(18,2); default:0"`
	Payment  float64 `gorm:"type:NUMERIC(18,2); default:0"`
	Currency string  `gorm:"not null"`
	Memo     string  // Optional field for additional notes
}

type TxLinesInput struct {
	BaseModel
	TransactionID    v7uuid.UUID         `json:"transactionid" gorm:"NOT NULL;type:uuid"`
	LineNo           uint                `gorm:"not null; serial;AUTO_INCREMENT"`
	TxLineDate       time.Time           `gorm:"not null; default:current_timestamp"`
	TxLineEntryDate  time.Time           `gorm:"not null; default:current_timestamp"`
	TxLineCategoryID v7uuid.UUID
	TxLinePayeeID    v7uuid.UUID `json:"txlinepayeeid" `
	Deposit  float64 `gorm:"type:NUMERIC(18,2); default:0"`
	Payment  float64 `gorm:"type:NUMERIC(18,2); default:0"`
	Currency string  `gorm:"not null"`
	Memo     string  // Optional field for additional notes
}

func (cmp *TxLinesInput) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	if time.Now().Sub(cmp.TxLineDate).Hours()/24 > 700 {
		cmp.TxLineDate = time.Now()
	}
	return
}
