package persistence

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayslipRepositoryImpl struct {
	db *gorm.DB
}

func NewPayslipRepository(db *gorm.DB) repositories.PayslipRepository {
	return &PayslipRepositoryImpl{db: db}
}

func (psr *PayslipRepositoryImpl) Create(payslip models.Payslip) (*models.Payslip, error) {
	if err := psr.db.Create(&payslip).Error; err != nil {
		return nil, err
	}
	return &payslip, nil
}

func (psr *PayslipRepositoryImpl) GetByPayrollAndUser(payrollID, userID uuid.UUID) (*models.Payslip, error) {
	var payslip models.Payslip
	if err := psr.db.Where("payroll_id = ? AND user_id = ?", payrollID, userID).First(&payslip).Error; err != nil {
		return nil, err
	}
	return &payslip, nil
}

func (psr *PayslipRepositoryImpl) ListByPayroll(payrollID uuid.UUID) ([]models.Payslip, error) {
	var payslips []models.Payslip
	if err := psr.db.Where("payroll_id = ?", payrollID).Find(&payslips).Error; err != nil {
		return nil, err
	}
	return payslips, nil
}

func (r *PayslipRepositoryImpl) SummaryByPayroll(payrollID uuid.UUID) ([]models.EmployeePayslipSummary, error) {
	var summaries []models.EmployeePayslipSummary

	err := r.db.Raw(`
		SELECT 
			u.id as employee_id,
			u.name as employee_name,
			p.net_pay as take_home_pay
		FROM payslips p
		JOIN users u ON u.id = p.user_id
		WHERE p.payroll_id = ?
	`, payrollID).Scan(&summaries).Error

	if err != nil {
		return nil, err
	}

	return summaries, nil
}
