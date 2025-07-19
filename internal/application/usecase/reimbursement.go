package usecase

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type Reimbursement interface {
	Create(input dto.CreateReimbursementDTO) (*models.Reimbursement, error)
	Update(input dto.UpdateReimbursementDTO) (*models.Reimbursement, error)
	GetByID(id uuid.UUID) (*models.Reimbursement, error)
	Delete(id uuid.UUID) error

	ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Reimbursement, error)
	ListAllInPeriod(startDate, endDate time.Time) ([]models.Reimbursement, error)
}
