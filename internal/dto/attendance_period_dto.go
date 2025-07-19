package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAttendancePeriodDTO struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required,gtfield=StartDate"`
}

type UpdateAttendancePeriodDTO struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}
