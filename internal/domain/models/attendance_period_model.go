package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type AttendancePeriod struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date;not null"`

	CreatedBy uuid.UUID
	CreatedAt time.Time

	CreatedByUser *User `gorm:"foreignKey:CreatedBy"`
}
