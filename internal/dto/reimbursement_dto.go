package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateReimbursementDTO struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Reason    string    `json:"reason"`
	Date      time.Time `json:"date" binding:"required"`
	CreatedBy uuid.UUID `json:"created_by"`
	IPAddress string    `json:"ip_address"`
}

type UpdateReimbursementDTO struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Reason    string    `json:"reason"`
	Date      time.Time `json:"date" binding:"required"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	IPAddress string    `json:"ip_address"`
}
