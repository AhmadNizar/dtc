package persistence

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) repositories.AttendanceRepository {
	return &AttendanceRepositoryImpl{db: db}
}

func (ar *AttendanceRepositoryImpl) Create(data models.Attendance) (*models.Attendance, error) {
	if err := ar.db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (ar *AttendanceRepositoryImpl) GetByID(id uuid.UUID) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := ar.db.First(&attendance, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (ar *AttendanceRepositoryImpl) Update(data models.Attendance) (*models.Attendance, error) {
	if err := ar.db.Save(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (ar *AttendanceRepositoryImpl) GetByUserAndDate(userID uuid.UUID, date time.Time) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := ar.db.Where("user_id = ? AND date = ?", userID, date).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (ar *AttendanceRepositoryImpl) ExistsByUserAndDate(userID uuid.UUID, date time.Time) (bool, error) {
	var count int64
	if err := ar.db.Model(&models.Attendance{}).Where("user_id = ? AND date = ?", userID, date).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *AttendanceRepositoryImpl) ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	if err := ar.db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (ar *AttendanceRepositoryImpl) ListAllInPeriod(startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	if err := ar.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
