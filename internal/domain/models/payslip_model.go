package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Payslip struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID                uuid.UUID `gorm:"not null"`
	PeriodID              uuid.UUID `gorm:"not null"`
	TotalAttendanceDays   int       `gorm:"not null"`
	TotalAttendanceSalary float64   `gorm:"type:numeric(15,2);not null"`
	TotalOvertimeHours    float64   `gorm:"type:numeric(10,2);not null"`
	TotalOvertimeSalary   float64   `gorm:"type:numeric(15,2);not null"`
	TotalReimbursement    float64   `gorm:"type:numeric(15,2);not null"`
	TotalTakeHome         float64   `gorm:"type:numeric(15,2);not null"`

	CreatedAt time.Time

	User   *User             `gorm:"foreignKey:UserID"`
	Period *AttendancePeriod `gorm:"foreignKey:PeriodID"`
}
