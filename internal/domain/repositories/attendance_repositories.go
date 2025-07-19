package repositories

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type AttendanceRepository interface {
	Create(attendance models.Attendance) (*models.Attendance, error)
	GetByID(id uuid.UUID) (*models.Attendance, error)
	Update(attendance models.Attendance) (*models.Attendance, error)
	GetByUserAndDate(userID uuid.UUID, date time.Time) (*models.Attendance, error)
	ExistsByUserAndDate(userID uuid.UUID, date time.Time) (bool, error)
	ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Attendance, error)
	ListAllInPeriod(startDate, endDate time.Time) ([]models.Attendance, error)
}
