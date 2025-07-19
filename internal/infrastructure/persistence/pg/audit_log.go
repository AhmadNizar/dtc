package persistence

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"gorm.io/gorm"
)

type AuditLogRepositoryImpl struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) repositories.AuditLogRepository {
	return &AuditLogRepositoryImpl{db: db}
}

func (alr *AuditLogRepositoryImpl) Create(log *models.AuditLog) error {
	return alr.db.Create(log).Error
}

func (alr *AuditLogRepositoryImpl) FindByEntity(entityType string, entityID string) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	if err := alr.db.Where("entity_type = ? AND entity_id = ?", entityType, entityID).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (alr *AuditLogRepositoryImpl) ListAll(limit, offset int) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	if err := alr.db.Limit(limit).Offset(offset).Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
