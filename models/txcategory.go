package models

import (
	v7uuid "github.com/gofrs/uuid"
	_ "gorm.io/gorm"
)

type TransactionCategoryType byte

// Transaction category types
const (
	CATEGORY_TYPE_INCOME   TransactionCategoryType = 1
	CATEGORY_TYPE_EXPENSE  TransactionCategoryType = 2
	CATEGORY_TYPE_TRANSFER TransactionCategoryType = 3
)

// TransactionCategory represents transaction category data stored in database
type TransactionCategory struct {
	BaseModel
	CategoryType   *TxCategoryType `json:"categorytype" gorm:"NOT NULL,foreignKey:CategoryTypeID"`
	CategoryTypeID v7uuid.UUID     `json:"categorytypeid" gorm:"type:uuid;NOT NULL"`
	// ParentCategory   *TransactionCategory `json:"parentcategory" gorm:"NOT NULL,foreignKey:ParentCategoryId"`
	// ParentCategoryId v7uuid.UUID          `json:"parentcategoryid" gorm:"type:uuid;NULL"`
	Name         string      `json:"name" gorm:"NOT NULL"`
	DisplayOrder int32       `json:"displayorder" gorm:"NOT NULL,foreignKey:CategoryTypeID"`
	Icon         int64       `json:"icon"  gorm:"NOT NULL"`
	Color        string      `json:"color"  gorm:"NOT NULL"`
	Hidden       bool        `json:"hidden"  gorm:"NOT NULL"`
	Comment      string      `json:"comment"  gorm:"NOT NULL"`
	UserID       v7uuid.UUID `json:"userid" gorm:"NOT NULL;type:uuid"`
	OrgID        v7uuid.UUID `json:"orgid" gorm:"NOT NULL;type:uuid"`
}
