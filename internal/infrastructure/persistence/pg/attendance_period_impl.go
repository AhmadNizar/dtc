package persistence

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendancePeriodRepositoryImpl struct {
	db *gorm.DB
}

func NewAttendancePeriodRepository(db *gorm.DB) *AttendancePeriodRepositoryImpl {
	return &AttendancePeriodRepositoryImpl{db: db}
}

func (apr *AttendancePeriodRepositoryImpl) Create(period models.AttendancePeriod) (*models.AttendancePeriod, error) {
	if err := apr.db.Create(&period).Error; err != nil {
		return nil, err
	}
	return &period, nil
}

func (apr *AttendancePeriodRepositoryImpl) Update(period models.AttendancePeriod) (*models.AttendancePeriod, error) {
	if err := apr.db.Save(&period).Error; err != nil {
		return nil, err
	}
	return &period, nil
}

func (apr *AttendancePeriodRepositoryImpl) Delete(id uuid.UUID) error {
	return apr.db.Delete(&models.AttendancePeriod{}, "id = ?", id).Error
}

func (apr *AttendancePeriodRepositoryImpl) GetByID(id uuid.UUID) (*models.AttendancePeriod, error) {
	var period models.AttendancePeriod
	if err := apr.db.First(&period, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &period, nil
}

func (apr *AttendancePeriodRepositoryImpl) List() ([]models.AttendancePeriod, error) {
	var periods []models.AttendancePeriod
	if err := apr.db.Find(&periods).Error; err != nil {
		return nil, err
	}
	return periods, nil
}

func (apr *AttendancePeriodRepositoryImpl) FindByDate(date time.Time) (*models.AttendancePeriod, error) {
	var period models.AttendancePeriod
	if err := apr.db.Where("? BETWEEN start_date AND end_date", date).First(&period).Error; err != nil {
		return nil, err
	}
	return &period, nil
}

func (apr *AttendancePeriodRepositoryImpl) IsOverlapping(startDate, endDate time.Time) (bool, error) {
	var count int64
	if err := apr.db.Model(&models.AttendancePeriod{}).
		Where("start_date <= ? AND end_date >= ?", endDate, startDate).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
