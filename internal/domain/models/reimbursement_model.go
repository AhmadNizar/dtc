package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Reimbursement struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Amount      float64   `gorm:"type:numeric(15,2);not null"`
	Description string

	Date      time.Time `gorm:"not null"` // actual reimbursement date
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	IPAddress string

	User          *User `gorm:"foreignKey:UserID"`
	CreatedByUser *User `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser *User `gorm:"foreignKey:UpdatedBy"`
}
