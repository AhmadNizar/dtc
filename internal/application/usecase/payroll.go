package usecase

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type Payroll interface {
	Create(input dto.CreatePayrollDTO) (*models.Payroll, error)
	Update(input dto.UpdatePayrollDTO) (*models.Payroll, error)
	GetByID(id uuid.UUID) (*models.Payroll, error)
	ListAll() ([]models.Payroll, error)
	ListByPeriod(periodID uuid.UUID) ([]models.Payroll, error)
}
