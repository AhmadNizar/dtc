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
	UpdatedAt time.Time
	CreatedBy *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid"`
	RequestIP string     `gorm:"type:inet"`
	RequestID string

	User   *User             `gorm:"foreignKey:UserID"`
	Period *AttendancePeriod `gorm:"foreignKey:PeriodID"`
}

func (Payslip) TableName() string {
	return "payslips"
}

type EmployeePayslipSummary struct {
	EmployeeID   uuid.UUID `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	TakeHomePay  float64   `json:"take_home_pay"`
}

type PayslipSummary struct {
	PeriodID    uuid.UUID                `json:"period_id"`
	PeriodName  string                   `json:"period_name"`
	Summaries   []EmployeePayslipSummary `json:"summaries"`
	TotalPayout float64                  `json:"total_payout"`
}
