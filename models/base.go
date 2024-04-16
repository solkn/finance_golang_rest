package models

import (
	"time"

	v7uuid "github.com/gofrs/uuid"
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID v7uuid.UUID `gorm:"primary_key;type:uuid"`
	// ID        v7uuid.UUID `gorm:"primarykey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	zero, _ := v7uuid.FromString("00000000-0000-0000-0000-000000000000")
	if m.ID == zero {
		m.ID, _ = v7uuid.NewV7()
	}
	return
}
