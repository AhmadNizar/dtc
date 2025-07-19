package persistence

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayrollRepositoryImpl struct {
	db *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) *PayrollRepositoryImpl {
	return &PayrollRepositoryImpl{db}
}

func (pr *PayrollRepositoryImpl) Create(payroll models.Payroll) (*models.Payroll, error) {
	if err := pr.db.Create(&payroll).Error; err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (pr *PayrollRepositoryImpl) Update(payroll models.Payroll) (*models.Payroll, error) {
	if err := pr.db.Save(&payroll).Error; err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (pr *PayrollRepositoryImpl) GetByID(id uuid.UUID) (*models.Payroll, error) {
	var payroll models.Payroll
	if err := pr.db.Preload("Period").First(&payroll, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (pr *PayrollRepositoryImpl) GetByPeriodID(periodID uuid.UUID) (*models.Payroll, error) {
	var payroll models.Payroll
	if err := pr.db.Preload("Period").First(&payroll, "period_id = ?", periodID).Error; err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (pr *PayrollRepositoryImpl) IsAlreadyProcessed(periodID uuid.UUID) (bool, error) {
	var count int64
	if err := pr.db.Model(&models.Payroll{}).Where("period_id = ?", periodID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pr *PayrollRepositoryImpl) ListAll() ([]models.Payroll, error) {
	var payrolls []models.Payroll
	if err := pr.db.Preload("Period").Order("created_at desc").Find(&payrolls).Error; err != nil {
		return nil, err
	}
	return payrolls, nil
}

func (pr *PayrollRepositoryImpl) ListByPeriod(periodID uuid.UUID) ([]models.Payroll, error) {
	var payrolls []models.Payroll
	if err := pr.db.Preload("Period").Where("period_id = ?", periodID).Find(&payrolls).Error; err != nil {
		return nil, err
	}
	return payrolls, nil
}
