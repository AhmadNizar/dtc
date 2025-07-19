package service

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type AttendanceServiceImpl struct {
	repo repositories.AttendanceRepository
}

func NewAttendanceService(repo repositories.AttendanceRepository) usecase.Attendance {
	return &AttendanceServiceImpl{repo: repo}
}

func (as *AttendanceServiceImpl) Create(attendance dto.CreateAttendanceDTO) (*models.Attendance, error) {
	newAttendanceData := models.Attendance{
		UserID:    attendance.UserID,
		Date:      attendance.Date,
		IPAddress: attendance.IPAddress,
		CreatedBy: attendance.CreatedBy,
	}
	return as.repo.Create(newAttendanceData)
}

func (as *AttendanceServiceImpl) Update(attendance dto.UpdateAttendanceDTO) (*models.Attendance, error) {
	// Optional: fetch existing record for safety/merge check
	existing, err := as.repo.GetByID(attendance.ID)
	if err != nil {
		return nil, err
	}

	// Update only the allowed fields
	existing.Date = attendance.Date
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = attendance.UpdatedBy
	existing.IPAddress = attendance.IPAddress

	return as.repo.Update(*existing)
}

func (as *AttendanceServiceImpl) GetByUserAndDate(userID uuid.UUID, date time.Time) (*models.Attendance, error) {
	return as.repo.GetByUserAndDate(userID, date)
}

func (as *AttendanceServiceImpl) ExistsByUserAndDate(userID uuid.UUID, date time.Time) (bool, error) {
	return as.repo.ExistsByUserAndDate(userID, date)
}

func (as *AttendanceServiceImpl) ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Attendance, error) {
	return as.repo.ListByUserInPeriod(userID, startDate, endDate)
}

func (as *AttendanceServiceImpl) ListAllInPeriod(startDate, endDate time.Time) ([]models.Attendance, error) {
	return as.repo.ListAllInPeriod(startDate, endDate)
}
