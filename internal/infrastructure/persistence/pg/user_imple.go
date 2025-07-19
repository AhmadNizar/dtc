package persistence

import (
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) ListAll() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
