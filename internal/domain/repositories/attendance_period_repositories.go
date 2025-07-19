package repositories

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type AttendancePeriodRepository interface {
	Create(period models.AttendancePeriod) (*models.AttendancePeriod, error)
	Update(period models.AttendancePeriod) (*models.AttendancePeriod, error) // ✅ Add this
	Delete(id uuid.UUID) error                                               // ✅ Add this
	GetByID(id uuid.UUID) (*models.AttendancePeriod, error)
	List() ([]models.AttendancePeriod, error)
	FindByDate(date time.Time) (*models.AttendancePeriod, error)
	IsOverlapping(startDate, endDate time.Time) (bool, error)
}
