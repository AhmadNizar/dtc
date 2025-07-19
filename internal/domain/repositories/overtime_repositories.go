package repositories

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type OvertimeRepository interface {
	ListAll() ([]models.Overtime, error)
	Create(overtime models.Overtime) (*models.Overtime, error)
	GetByID(id uuid.UUID) (*models.Overtime, error)
	Update(overtime models.Overtime) (*models.Overtime, error)
	ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Overtime, error)
	ListAllInPeriod(startDate, endDate time.Time) ([]models.Overtime, error)
}
