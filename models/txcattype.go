package models

import (
	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TxCategoryType struct {
	BaseModel
	Title       string      // Title of the event.
	Description string      // Description of the event.
	UserID      v7uuid.UUID `json:"userid" gorm:"NOT NULL;type:uuid"`
	OrgID       v7uuid.UUID `json:"orgid" gorm:"NOT NULL;type:uuid"`
}

func (loc *TxCategoryType) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	if loc.ID == zero {
		loc.ID, _ = v7uuid.NewV7()
	}
	return
}

type CreateTxCategoryInput struct {
	BaseModel
	Title       string // Title of the event.
	Description string // Description of the event.
	UserID      v7uuid.UUID
}

func (loc *CreateTxCategoryInput) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	if loc.ID == zero {
		loc.ID, _ = v7uuid.NewV7()
	}
	return
}

type UpdateTxCategoryInput struct {
	Title       string      // Title of the event.
	Description string      // Description of the event.
	UserID      v7uuid.UUID `json:"-"`
}

type GetTxCategoryOutput struct {
	BaseModel
	Title       string      // Title of the event.
	Description string      // Description of the event.
	UserID      v7uuid.UUID `json:"-"`
}
