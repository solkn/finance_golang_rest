package models

import (
	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	BaseModel
	Name   string      `json:"name" binding:"required"`
	User   *User        `json:"user"  gorm:"foreignKey:UserID"`
	UserID v7uuid.UUID `json:"userid" gorm:"NOT NULL;type:uuid"`
	Org    *Org         `json:"org" gorm:"foreignKey:OrgID"`
	OrgID  v7uuid.UUID `json:"orgid" gorm:"NOT NULL;type:uuid"`
}

func (cmp *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	return
}
