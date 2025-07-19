package persistence

import (
	"time"

	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReimbursementRepositoryImpl struct {
	db *gorm.DB
}

func NewReimbursementRepository(db *gorm.DB) repositories.ReimbursementRepository {
	return &ReimbursementRepositoryImpl{db: db}
}

func (rr *ReimbursementRepositoryImpl) Create(reimbursement models.Reimbursement) (*models.Reimbursement, error) {
	if err := rr.db.Create(&reimbursement).Error; err != nil {
		return nil, err
	}
	return &reimbursement, nil
}

func (rr *ReimbursementRepositoryImpl) GetByID(id uuid.UUID) (*models.Reimbursement, error) {
	var reimbursement models.Reimbursement
	if err := rr.db.Where("id = ?", id).First(&reimbursement).Error; err != nil {
		return nil, err
	}
	return &reimbursement, nil
}

func (rr *ReimbursementRepositoryImpl) ListByUserInPeriod(userID uuid.UUID, startDate, endDate time.Time) ([]models.Reimbursement, error) {
	var reimbursements []models.Reimbursement
	if err := rr.db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).Find(&reimbursements).Error; err != nil {
		return nil, err
	}
	return reimbursements, nil
}

func (rr *ReimbursementRepositoryImpl) ListAllInPeriod(startDate, endDate time.Time) ([]models.Reimbursement, error) {
	var reimbursements []models.Reimbursement
	if err := rr.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&reimbursements).Error; err != nil {
		return nil, err
	}
	return reimbursements, nil
}

func (rr *ReimbursementRepositoryImpl) Update(reimbursement models.Reimbursement) (*models.Reimbursement, error) {
	if err := rr.db.Save(&reimbursement).Error; err != nil {
		return nil, err
	}
	return &reimbursement, nil
}

func (rr *ReimbursementRepositoryImpl) Delete(id uuid.UUID) error {
	return rr.db.Delete(&models.Reimbursement{}, "id = ?", id).Error
}
