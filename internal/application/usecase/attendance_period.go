package usecase

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type AttendancePeriod interface {
	Create(input dto.CreateAttendancePeriodDTO) (*models.AttendancePeriod, error)
	Update(input dto.UpdateAttendancePeriodDTO) (*models.AttendancePeriod, error)
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.AttendancePeriod, error)
	List() ([]models.AttendancePeriod, error)
	GetCurrentPeriod(currentDate time.Time) (*models.AttendancePeriod, error)
}
