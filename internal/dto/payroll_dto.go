package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePayrollDTO struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	PeriodID    uuid.UUID `json:"period_id" binding:"required"`
	BaseSalary  float64   `json:"base_salary" binding:"required"`
	OvertimePay float64   `json:"overtime_pay"` // Optional
	Bonus       float64   `json:"bonus"`        // Optional
	Deductions  float64   `json:"deductions"`   // Optional
	NetPay      float64   `json:"net_pay"`      // Optional (can be calculated)
	GeneratedAt time.Time `json:"generated_at"` // Optional: default to now
	Approved    *bool     `json:"approved"`     // Optional approval status
	CreatedBy   uuid.UUID `json:"created_by"`   // Tracking
	IPAddress   string    `json:"ip_address"`   // Tracking
}

type UpdatePayrollDTO struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	BaseSalary  float64   `json:"base_salary"`
	OvertimePay float64   `json:"overtime_pay"`
	Bonus       float64   `json:"bonus"`
	Deductions  float64   `json:"deductions"`
	NetPay      float64   `json:"net_pay"`
	Approved    *bool     `json:"approved"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	IPAddress   string    `json:"ip_address"`
}
