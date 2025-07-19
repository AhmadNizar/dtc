package dto

import "github.com/google/uuid"

type CreatePayslipDTO struct {
	UserID     uuid.UUID `json:"user_id" binding:"required"`
	PeriodID   uuid.UUID `json:"period_id" binding:"required"`
	BasePay    float64   `json:"base_pay" binding:"required"`
	Bonus      float64   `json:"bonus"`
	Deductions float64   `json:"deductions"`
	NetPay     float64   `json:"net_pay" binding:"required"`

	CreatedBy uuid.UUID `json:"created_by"`
	IPAddress string    `json:"ip_address"`
}

type UpdatePayslipDTO struct {
	ID         uuid.UUID `json:"id" binding:"required"`
	BasePay    float64   `json:"base_pay" binding:"required"`
	Bonus      float64   `json:"bonus"`
	Deductions float64   `json:"deductions"`
	NetPay     float64   `json:"net_pay" binding:"required"`

	UpdatedBy uuid.UUID `json:"updated_by"`
	IPAddress string    `json:"ip_address"`
}
