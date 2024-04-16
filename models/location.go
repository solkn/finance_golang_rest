package models

import (
	"database/sql"
)

type Location struct {
	BaseModel
	Title       string       // Title of the event.
	Description string       // Description of the event.
	StartTime   sql.NullTime `gorm:"type:TIMESTAMP NULL"` // Start time of the event.
	EndTime     sql.NullTime `gorm:"type:TIMESTAMP NULL"` // End time of the event.
	Street      string       // Street of Location of the event.
}
