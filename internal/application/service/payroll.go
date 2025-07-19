package service

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type PayrollServiceImpl struct {
	repo repositories.PayrollRepository
}

func NewPayrollService(repo repositories.PayrollRepository) usecase.Payroll {
	return &PayrollServiceImpl{repo: repo}
}

func (pr *PayrollServiceImpl) Create(input dto.CreatePayrollDTO) (*models.Payroll, error) {
	payroll := models.Payroll{
		ID:          uuid.New(),
		UserID:      input.UserID,
		PeriodID:    input.PeriodID,
		BaseSalary:  input.BaseSalary,
		OvertimePay: input.OvertimePay,
		Bonus:       input.Bonus,
		Deductions:  input.Deductions,
		NetPay:      input.NetPay,
		Approved:    *input.Approved,
		GeneratedAt: input.GeneratedAt,
		CreatedBy:   &input.CreatedBy,
		RequestIP:   input.IPAddress,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return pr.repo.Create(payroll)
}

func (pr *PayrollServiceImpl) Update(input dto.UpdatePayrollDTO) (*models.Payroll, error) {
	existing, err := pr.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existing.BaseSalary = input.BaseSalary
	existing.OvertimePay = input.OvertimePay
	existing.Bonus = input.Bonus
	existing.Deductions = input.Deductions
	existing.NetPay = input.NetPay
	existing.Approved = *input.Approved
	existing.UpdatedBy = &input.UpdatedBy
	existing.RequestIP = input.IPAddress
	existing.UpdatedAt = time.Now()

	return pr.repo.Update(*existing)
}

func (pr *PayrollServiceImpl) GetByID(id uuid.UUID) (*models.Payroll, error) {
	return pr.repo.GetByID(id)
}

func (pr *PayrollServiceImpl) ListAll() ([]models.Payroll, error) {
	return pr.repo.ListAll()
}

func (pr *PayrollServiceImpl) ListByPeriod(periodID uuid.UUID) ([]models.Payroll, error) {
	return pr.repo.ListByPeriod(periodID)
}
