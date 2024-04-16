package models

import (
	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TxTag struct {
	BaseModel
	TransactionID v7uuid.UUID `json:"userid" gorm:"NOT NULL;type:uuid"`
	Tag           *Tag        `json:"tag" gorm:"foreignKey:TagID"`
	TagID         v7uuid.UUID `json:"tagid" gorm:"type:uuid"`
	Org           *Org        `json:"org" gorm:"foreignKey:OrgID"`
	OrgID         v7uuid.UUID `json:"orgid" gorm:"NOT NULL;type:uuid"`
}

func (cmp *TxTag) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	return
}
