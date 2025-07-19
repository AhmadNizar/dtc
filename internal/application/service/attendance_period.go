package service

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/google/uuid"
)

type AttendancePeriodServiceImpl struct {
	repo repositories.AttendancePeriodRepository
}

func NewAttendancePeriodService(repo repositories.AttendancePeriodRepository) usecase.AttendancePeriod {
	return &AttendancePeriodServiceImpl{repo: repo}
}

func (aps *AttendancePeriodServiceImpl) Create(input dto.CreateAttendancePeriodDTO) (*models.AttendancePeriod, error) {
	model := models.AttendancePeriod{
		ID:        uuid.New(),
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}
	return aps.repo.Create(model)
}

func (aps *AttendancePeriodServiceImpl) Update(input dto.UpdateAttendancePeriodDTO) (*models.AttendancePeriod, error) {
	existing, err := aps.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existing.StartDate = input.StartDate
	existing.EndDate = input.EndDate

	return aps.repo.Update(*existing)
}

func (aps *AttendancePeriodServiceImpl) Delete(id uuid.UUID) error {
	return aps.repo.Delete(id)
}

func (aps *AttendancePeriodServiceImpl) GetByID(id uuid.UUID) (*models.AttendancePeriod, error) {
	return aps.repo.GetByID(id)
}

func (aps *AttendancePeriodServiceImpl) List() ([]models.AttendancePeriod, error) {
	return aps.repo.List()
}

func (aps *AttendancePeriodServiceImpl) GetCurrentPeriod(currentDate time.Time) (*models.AttendancePeriod, error) {
	return aps.repo.FindByDate(currentDate)
}
