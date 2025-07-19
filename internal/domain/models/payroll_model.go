package models

import (
	"time"

	"github.com/google/uuid"
)

type Payroll struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PeriodID uuid.UUID `gorm:"type:uuid;not null;unique"`
	Period   AttendancePeriod

	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid"`
	RequestIP string     `gorm:"type:inet"`
	RequestID string
}
