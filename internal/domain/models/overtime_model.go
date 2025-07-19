package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Overtime struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Date   time.Time `gorm:"type:date;not null"`
	Hours  float64   `gorm:"type:numeric(4,2);not null"` // Max 3

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	IPAddress string

	User          *User `gorm:"foreignKey:UserID"`
	CreatedByUser *User `gorm:"foreignKey:CreatedBy"`
}
