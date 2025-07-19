package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"type:text;not null;check:role IN ('admin', 'employee')"`
	Salary       float64   `gorm:"type:numeric(15,2);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
