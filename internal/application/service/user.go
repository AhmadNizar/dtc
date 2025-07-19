package service

import (
	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/domain/models"
	"github.com/AhmadNizar/dtc/internal/domain/repositories"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) usecase.User {
	return &UserServiceImpl{repo: repo}
}

func (us *UserServiceImpl) GetByUsername(username string) (*models.User, error) {
	return us.repo.GetByUsername(username)
}

func (us *UserServiceImpl) GetByID(id uuid.UUID) (*models.User, error) {
	return us.repo.GetByID(id)
}

func (us *UserServiceImpl) ListAll() ([]models.User, error) {
	return us.repo.ListAll()
}
