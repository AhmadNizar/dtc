package models

import (
	"time"

	"github.com/google/uuid"
)

type Payroll struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	User     User      `gorm:"foreignKey:UserID"`
	PeriodID uuid.UUID `gorm:"type:uuid;not null"`
	Period   AttendancePeriod

	BaseSalary  float64   `gorm:"type:numeric(12,2);not null"`
	OvertimePay float64   `gorm:"type:numeric(12,2);not null"`
	Bonus       float64   `gorm:"type:numeric(12,2);not null"`
	Deductions  float64   `gorm:"type:numeric(12,2);not null"`
	NetPay      float64   `gorm:"type:numeric(12,2);not null"`
	Approved    bool      `gorm:"default:false"`
	GeneratedAt time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid"`
	RequestIP string     `gorm:"type:inet"`
	RequestID string
}

func (Payroll) TableName() string {
	return "payrolls"
}
