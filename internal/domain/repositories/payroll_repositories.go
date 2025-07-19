package repositories

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type PayrollRepository interface {
	Create(payroll models.Payroll) (*models.Payroll, error)
	Update(payroll models.Payroll) (*models.Payroll, error)
	GetByID(id uuid.UUID) (*models.Payroll, error)
	GetByPeriodID(periodID uuid.UUID) (*models.Payroll, error)
	ListAll() ([]models.Payroll, error)
	ListByPeriod(periodID uuid.UUID) ([]models.Payroll, error)
	IsAlreadyProcessed(periodID uuid.UUID) (bool, error)
}
