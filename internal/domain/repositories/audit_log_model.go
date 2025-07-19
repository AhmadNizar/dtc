package repositories

import "github.com/AhmadNizar/dtc/internal/domain/models"

type AuditLogRepository interface {
	Create(log *models.AuditLog) error
	FindByEntity(entityType string, entityID string) ([]models.AuditLog, error)
	ListAll(limit, offset int) ([]models.AuditLog, error)
}
