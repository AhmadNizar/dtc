package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAttendanceDTO struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	IPAddress string    `json:"ip_address"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type UpdateAttendanceDTO struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	IPAddress string    `json:"ip_address"`
}
