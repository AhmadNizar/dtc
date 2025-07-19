package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Overtime struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_date"`
	Date      time.Time `gorm:"type:date;not null;uniqueIndex:idx_user_date"`
	Hours     float64   `gorm:"type:numeric(5,2);not null"`
	Reason    string
	IPAddress string `gorm:"type:varchar(45)"`

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User          *User `gorm:"foreignKey:UserID;references:ID"`
	CreatedByUser *User `gorm:"foreignKey:CreatedBy;references:ID"`
}
