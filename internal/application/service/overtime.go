package service

import (
	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type OvertimeServiceImpl struct {
	repo repositories.OvertimeRepository
}

func NewOvertimeService(repo repositories.OvertimeRepository) usecase.Overtime {
	return &OvertimeServiceImpl{repo: repo}
}

func (os *OvertimeServiceImpl) Create(input dto.CreateOvertimeDTO) (*models.Overtime, error) {
	overtime := models.Overtime{
		UserID:    input.UserID,
		Date:      input.Date,
		Hours:     input.Hours,
		Reason:    input.Reason,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.CreatedBy,
		IPAddress: input.IPAddress,
	}

	return os.repo.Create(overtime)
}

func (os *OvertimeServiceImpl) Update(input dto.UpdateOvertimeDTO) (*models.Overtime, error) {
	existing, err := os.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existing.Hours = input.Hours
	existing.UpdatedBy = input.UpdatedBy
	existing.IPAddress = input.IPAddress

	return os.repo.Update(*existing)
}

func (s *OvertimeServiceImpl) GetByID(id uuid.UUID) (*models.Overtime, error) {
	return s.repo.GetByID(id)
}

func (s *OvertimeServiceImpl) ListAll() ([]models.Overtime, error) {
	return s.repo.ListAll()
}
