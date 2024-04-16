package models

import (
	v7uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `gorm:"not null;uniqueIndex"`
	Phone     string `json:"phone" binding:"required"`
}

func (cmp *User) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	//zero, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if cmp.ID == zero {
		cmp.ID, _ = v7uuid.NewV7()
	}
	return
}
