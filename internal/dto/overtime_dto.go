package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateOvertimeDTO struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Hours     float64   `json:"hours" binding:"required"`
	Reason    string    `json:"reason"`     // Optional, if you add it in model
	IPAddress string    `json:"ip_address"` // Can be filled from context in handler
	CreatedBy uuid.UUID `json:"created_by"` // Usually filled from JWT context, not user-provided
}

type UpdateOvertimeDTO struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Hours     float64   `json:"hours" binding:"required"`
	IPAddress string    `json:"ip_address"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}
