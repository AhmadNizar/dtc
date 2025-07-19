package repositories

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type ReimbursementRepository interface {
	Create(reimbursement models.Reimbursement) (*models.Reimbursement, error)
	GetByID(id uuid.UUID) (*models.Reimbursement, error)
	ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Reimbursement, error)
	ListAllInPeriod(startDate, endDate time.Time) ([]models.Reimbursement, error)
	Update(reimbursement models.Reimbursement) (*models.Reimbursement, error)
	Delete(id uuid.UUID) error
}
