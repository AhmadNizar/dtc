package repositories

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	ListAll() ([]models.User, error)
}
