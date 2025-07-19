package usecase

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type Overtime interface {
	Create(overtime dto.CreateOvertimeDTO) (*models.Overtime, error)
	Update(overtime dto.UpdateOvertimeDTO) (*models.Overtime, error)
	GetByID(id uuid.UUID) (*models.Overtime, error)
	ListAll() ([]models.Overtime, error)
}
