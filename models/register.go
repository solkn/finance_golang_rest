package models

import (
	"time"

	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Register represents a collections of transactions

type Register struct {
	BaseModel
	Name         string         `json:"name"`
	BeginDate    time.Time      `json:"time"`
	Transactions []*Transaction `gorm:"foreignKey:RegisterID"`
	UserID       v7uuid.UUID    `json:"userid" gorm:"NOT NULL;type:uuid"`
	OrgID        v7uuid.UUID    `json:"orgid" gorm:"NOT NULL;type:uuid"`
}

func (cmp *Register) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	if time.Now().Sub(cmp.BeginDate).Hours()/24 > 700 {
		cmp.BeginDate = time.Now()
	}
	return
}
