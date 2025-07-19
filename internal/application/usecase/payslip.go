package usecase

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type Payslip interface {
	Create(input dto.CreatePayslipDTO) (*models.Payslip, error)
	Update(input dto.UpdatePayslipDTO) (*models.Payslip, error)
	GetByUserAndPayroll(payrollID, userID uuid.UUID) (*models.Payslip, error)
	ListByPayroll(payrollID uuid.UUID) ([]models.Payslip, error)
	SummaryByPayroll(payrollID uuid.UUID) (*models.PayslipSummary, error)
}
