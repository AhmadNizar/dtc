package service

import (
	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
)

type AuditLogServiceImpl struct {
	repo repositories.AuditLogRepository
}

func NewAuditLogService(repo repositories.AuditLogRepository) usecase.AuditLog {
	return &AuditLogServiceImpl{repo: repo}
}

func (als *AuditLogServiceImpl) Create(log *models.AuditLog) error {
	return als.repo.Create(log)
}

func (als *AuditLogServiceImpl) FindByEntity(entityType string, entityID string) ([]models.AuditLog, error) {
	return als.repo.FindByEntity(entityType, entityID)
}

func (als *AuditLogServiceImpl) ListAll(limit, offset int) ([]models.AuditLog, error) {
	return als.repo.ListAll(limit, offset)
}
