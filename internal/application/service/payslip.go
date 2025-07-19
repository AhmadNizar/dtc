package service

import (
	"errors"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
)

type PayslipServiceImpl struct {
	repo repositories.PayslipRepository
}

func NewPayslipService(repo repositories.PayslipRepository) usecase.Payslip {
	return &PayslipServiceImpl{repo: repo}
}

func (pls *PayslipServiceImpl) Create(input dto.CreatePayslipDTO) (*models.Payslip, error) {
	payslip := models.Payslip{
		ID:                    uuid.New(), // or use default from GORM
		UserID:                input.UserID,
		PeriodID:              input.PeriodID,
		TotalAttendanceDays:   0, // optionally calculated elsewhere
		TotalAttendanceSalary: input.BasePay,
		TotalOvertimeHours:    0,
		TotalOvertimeSalary:   input.Bonus,
		TotalReimbursement:    0,
		TotalTakeHome:         input.NetPay,
	}
	return pls.repo.Create(payslip)
}

func (pls *PayslipServiceImpl) Update(input dto.UpdatePayslipDTO) (*models.Payslip, error) {
	// First get existing
	existing, err := pls.repo.GetByPayrollAndUser(uuid.Nil, uuid.Nil) // you'll need a new repo method for GetByID or add ID to GetByPayrollAndUser
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("payslip not found")
	}

	// Update fields
	existing.TotalAttendanceSalary = input.BasePay
	existing.TotalOvertimeSalary = input.Bonus
	existing.TotalTakeHome = input.NetPay

	// You might want to add a new repo.Update(payslip) method
	return pls.repo.Create(*existing) // overwrite for now unless update method exists
}

func (pls *PayslipServiceImpl) GetByUserAndPayroll(payrollID, userID uuid.UUID) (*models.Payslip, error) {
	return pls.repo.GetByPayrollAndUser(payrollID, userID)
}

func (pls *PayslipServiceImpl) ListByPayroll(payrollID uuid.UUID) ([]models.Payslip, error) {
	return pls.repo.ListByPayroll(payrollID)
}

func (pls *PayslipServiceImpl) SummaryByPayroll(payrollID uuid.UUID) (*models.PayslipSummary, error) {
	summaries, err := pls.repo.SummaryByPayroll(payrollID)
	if err != nil {
		return nil, err
	}

	var total float64
	for _, s := range summaries {
		total += s.TakeHomePay
	}

	return &models.PayslipSummary{
		PeriodID:    payrollID,
		PeriodName:  "", // optional
		Summaries:   summaries,
		TotalPayout: total,
	}, nil
}
