package persistence

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OvertimeRepositoryImpl struct {
	db *gorm.DB
}

func NewOvertimeRepository(db *gorm.DB) repositories.OvertimeRepository {
	return &OvertimeRepositoryImpl{db: db}
}

func (or *OvertimeRepositoryImpl) Create(overtime models.Overtime) (*models.Overtime, error) {
	if err := or.db.Create(&overtime).Error; err != nil {
		return nil, err
	}
	return &overtime, nil
}

func (or *OvertimeRepositoryImpl) ListAll() ([]models.Overtime, error) {
	var overtimes []models.Overtime
	if err := or.db.Preload("User").Preload("CreatedByUser").Find(&overtimes).Error; err != nil {
		return nil, err
	}
	return overtimes, nil
}

func (or *OvertimeRepositoryImpl) ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Overtime, error) {
	var overtimes []models.Overtime
	if err := or.db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).Find(&overtimes).Error; err != nil {
		return nil, err
	}
	return overtimes, nil
}

func (or *OvertimeRepositoryImpl) ListAllInPeriod(startDate, endDate time.Time) ([]models.Overtime, error) {
	var overtimes []models.Overtime
	if err := or.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&overtimes).Error; err != nil {
		return nil, err
	}
	return overtimes, nil
}

func (or *OvertimeRepositoryImpl) GetByID(id uuid.UUID) (*models.Overtime, error) {
	var overtime models.Overtime
	if err := or.db.Preload("User").Preload("CreatedByUser").First(&overtime, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &overtime, nil
}

func (or *OvertimeRepositoryImpl) Update(overtime models.Overtime) (*models.Overtime, error) {
	if err := or.db.Save(&overtime).Error; err != nil {
		return nil, err
	}
	return &overtime, nil
}
