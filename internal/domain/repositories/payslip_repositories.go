package repositories

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type PayslipRepository interface {
	Create(payslip models.Payslip) (*models.Payslip, error)
	GetByPayrollAndUser(payrollID, userID uuid.UUID) (*models.Payslip, error)
	ListByPayroll(payrollID uuid.UUID) ([]models.Payslip, error)
	SummaryByPayroll(payrollID uuid.UUID) ([]models.EmployeePayslipSummary, error)
}
