package service

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type ReimbursementServiceImpl struct {
	repo repositories.ReimbursementRepository
}

func NewReimbursementService(repo repositories.ReimbursementRepository) usecase.Reimbursement {
	return &ReimbursementServiceImpl{repo: repo}
}

func (rs *ReimbursementServiceImpl) Create(input dto.CreateReimbursementDTO) (*models.Reimbursement, error) {
	reimbursement := &models.Reimbursement{
		UserID:      input.UserID,
		Amount:      input.Amount,
		Description: input.Reason, // map Reason -> Description
		Date:        input.Date,
		CreatedBy:   input.CreatedBy,
		IPAddress:   input.IPAddress,
	}
	return rs.repo.Create(*reimbursement)
}

func (rs *ReimbursementServiceImpl) Update(input dto.UpdateReimbursementDTO) (*models.Reimbursement, error) {
	existing, err := rs.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existing.Amount = input.Amount
	existing.Description = input.Reason // map Reason -> Description
	existing.Date = input.Date
	existing.UpdatedBy = input.UpdatedBy
	existing.IPAddress = input.IPAddress

	return rs.repo.Update(*existing)
}

func (rs *ReimbursementServiceImpl) GetByID(id uuid.UUID) (*models.Reimbursement, error) {
	return rs.repo.GetByID(id)
}

func (rs *ReimbursementServiceImpl) Delete(id uuid.UUID) error {
	return rs.repo.Delete(id)
}

func (rs *ReimbursementServiceImpl) ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Reimbursement, error) {
	return rs.repo.ListByUserInPeriod(userID, startDate, endDate)
}

func (rs *ReimbursementServiceImpl) ListAllInPeriod(startDate, endDate time.Time) ([]models.Reimbursement, error) {
	return rs.repo.ListAllInPeriod(startDate, endDate)
}
