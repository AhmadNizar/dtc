package usecase

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type Attendance interface {
	Create(attendance dto.CreateAttendanceDTO) (*models.Attendance, error)
	Update(attendance dto.UpdateAttendanceDTO) (*models.Attendance, error)
	GetByUserAndDate(userID uuid.UUID, date time.Time) (*models.Attendance, error)
	ExistsByUserAndDate(userID uuid.UUID, date time.Time) (bool, error)
	ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Attendance, error)
	ListAllInPeriod(startDate, endDate time.Time) ([]models.Attendance, error)
}
